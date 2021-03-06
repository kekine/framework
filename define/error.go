package define

import (
	"encoding/json"
	"fmt"
)

const (
	// ErrnoSuccess 成功
	ErrnoSuccess int = 0

	// ErrnoFailure 失败
	ErrnoFailure int = 1
)

var (
	// ErrSuccess 成功
	ErrSuccess = &MyError{Errno: ErrnoSuccess, Errdesc: "success"}

	// ErrFailure 失败
	ErrFailure = &MyError{Errno: ErrnoFailure, Errdesc: "failure"}

	// ErrSignature 签名
	ErrSignature = &MyError{Errno: ErrnoFailure, Errdesc: "signature"}

	// ErrDisconnect 断开连接
	ErrDisconnect = &MyError{Errno: ErrnoFailure, Errdesc: "disconnect"}

	// ErrUnknownMainCmd 未知主命令
	ErrUnknownMainCmd = &MyError{Errno: ErrnoFailure, Errdesc: "unknown main cmd"}

	// ErrUnknownSubCmd 未知子命令
	ErrUnknownSubCmd = &MyError{Errno: ErrnoFailure, Errdesc: "unknown sub cmd"}

	// ErrRepeatRegisterService 重复注册服务
	ErrRepeatRegisterService = &MyError{Errno: ErrnoFailure, Errdesc: "repeat register service"}

	// ErrNotExistService 不存在该服务
	ErrNotExistService = &MyError{Errno: ErrnoFailure, Errdesc: "not exist service"}

	// ErrServiceAlreadyOpen 服务已经开启
	ErrServiceAlreadyOpen = &MyError{Errno: ErrnoFailure, Errdesc: "service already open"}

	// ErrServiceAlreadyShut 服务已经关闭
	ErrServiceAlreadyShut = &MyError{Errno: ErrnoFailure, Errdesc: "service already shut"}
)

// MyError 错误
type MyError struct {
	Errno   int    `json:",omitempty"` // 错误码
	Errdesc string `json:",omitempty"` // 错误描述
}

func (m *MyError) Error() string {
	return fmt.Sprintf(`{"Errno":%d,"Errdesc":"%s"}`, m.Errno, m.Errdesc)
}

// CheckError 检查错误
func CheckError(data []byte) error {
	me := &MyError{}

	if err := json.Unmarshal(data, me); err != nil {
		return err
	}

	if me.Errno != ErrnoSuccess {
		return me
	}

	return nil
}
