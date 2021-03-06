package game

import (
	"fmt"
	"net/http"
	"sort"
	"sync"

	"github.com/panshiqu/framework/define"
)

var tins TableManager

// TableManager 桌子管理
type TableManager struct {
	count  int           // 计数
	mutex  sync.Mutex    // 加锁
	tables []*TableFrame // 桌子
}

// TrySitDown 尝试坐下
func (t *TableManager) TrySitDown(userItem *UserItem) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	for {
		sort.Sort(TableFrameSlice(t.tables))

		// 只要有桌子椅子就能坐下，这里不关心桌子状态
		if len(t.tables) == 0 || t.tables[0].UserCount() == define.CG.UserPerTable {
			t.AddTableFrame()
			continue
		}

		t.tables[0].SitDown(userItem)
		break
	}
}

// AddTableFrame 增加桌子
func (t *TableManager) AddTableFrame() {
	t.count++

	tableFrame := &TableFrame{
		id:    t.count,
		users: make([]*UserItem, define.CG.UserPerTable),
	}

	t.tables = append(t.tables, tableFrame)
}

// TableFrameSlice 排序
type TableFrameSlice []*TableFrame

func (t TableFrameSlice) Len() int {
	return len(t)
}
func (t TableFrameSlice) Less(i, j int) bool {
	if t[i].TableStatus() != t[j].TableStatus() {
		return t[i].TableStatus() < t[j].TableStatus()
	} else if c1, c2 := t[i].UserCount(), t[j].UserCount(); c1 != c2 {
		switch {
		case c2 == define.CG.UserPerTable:
			return true
		case c1 == define.CG.UserPerTable:
			return false
		default:
			return c1 > c2
		}
	}

	return t[i].TableID() < t[j].TableID()
}
func (t TableFrameSlice) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Monitor 监视器
func (t *TableManager) Monitor(w http.ResponseWriter, r *http.Request) {
	t.mutex.Lock()
	fmt.Fprintln(w, "tables:")
	for _, v := range t.tables {
		fmt.Fprintln(w, v)
	}
	t.mutex.Unlock()
}
