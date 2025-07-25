// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceWebTerminalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckInfo(v string) *CheckInstanceWebTerminalRequest
	GetCheckInfo() *string
}

type CheckInstanceWebTerminalRequest struct {
	// example:
	//
	// wss://pai-dlc-proxy-cn-shanghai.aliyun.com/terminal/t1157703270994901/dlcmjzjt1dxbmx4h/dlcmjzjt1dxbmx4h-worker-0?Token=******
	CheckInfo *string `json:"CheckInfo,omitempty" xml:"CheckInfo,omitempty"`
}

func (s CheckInstanceWebTerminalRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceWebTerminalRequest) GoString() string {
	return s.String()
}

func (s *CheckInstanceWebTerminalRequest) GetCheckInfo() *string {
	return s.CheckInfo
}

func (s *CheckInstanceWebTerminalRequest) SetCheckInfo(v string) *CheckInstanceWebTerminalRequest {
	s.CheckInfo = &v
	return s
}

func (s *CheckInstanceWebTerminalRequest) Validate() error {
	return dara.Validate(s)
}
