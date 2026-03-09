package vars

import (
	"github.com/zsqmiller/open-im-sdk/v3/sdk_struct"
	"sync/atomic"
)

type StatMsg struct {
	CostTime    int64
	ReceiveTime int64
	Msg         *sdk_struct.MsgStruct
}

var (
	SendMsgCount     atomic.Int64
	RecvMsgConsuming chan *StatMsg
)
