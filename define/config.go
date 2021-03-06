package define

// ConfigManager 管理服务器配置
type ConfigManager struct {
	PprofIP  string
	ListenIP string
}

// ConfigProxy 代理服务器配置
type ConfigProxy struct {
	ID       int
	DialIP   string
	PprofIP  string
	ListenIP string
}

// ConfigLogin 登陆服务器配置
type ConfigLogin struct {
	ID       int
	DBIP     string
	DialIP   string
	PprofIP  string
	ListenIP string
}

// ConfigGame 游戏服务器配置
type ConfigGame struct {
	ID           int
	DBIP         string
	DialIP       string
	PprofIP      string
	ListenIP     string
	UserPerTable int // 用户数量每桌
}

// CG 游戏配置
var CG ConfigGame

// ConfigDB 数据库服务器配置
type ConfigDB struct {
	PprofIP  string
	ListenIP string
}
