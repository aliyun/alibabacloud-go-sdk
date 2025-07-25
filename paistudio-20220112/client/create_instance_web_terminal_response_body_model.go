// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceWebTerminalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInstanceWebTerminalResponseBody
	GetRequestId() *string
	SetWebTerminalId(v string) *CreateInstanceWebTerminalResponseBody
	GetWebTerminalId() *string
}

type CreateInstanceWebTerminalResponseBody struct {
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// wss://pai-dlc-proxy-cn-shanghai.aliyun.com/terminal/t1157703270994901/dlcmjzjt1dxbmx4h/dlcmjzjt1dxbmx4h-worker-0?Token=******
	WebTerminalId *string `json:"WebTerminalId,omitempty" xml:"WebTerminalId,omitempty"`
}

func (s CreateInstanceWebTerminalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceWebTerminalResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceWebTerminalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceWebTerminalResponseBody) GetWebTerminalId() *string {
	return s.WebTerminalId
}

func (s *CreateInstanceWebTerminalResponseBody) SetRequestId(v string) *CreateInstanceWebTerminalResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceWebTerminalResponseBody) SetWebTerminalId(v string) *CreateInstanceWebTerminalResponseBody {
	s.WebTerminalId = &v
	return s
}

func (s *CreateInstanceWebTerminalResponseBody) Validate() error {
	return dara.Validate(s)
}
