// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmFullDuplexCallOperateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallId(v string) *LlmFullDuplexCallOperateRequest
	GetCallId() *string
	SetCommand(v string) *LlmFullDuplexCallOperateRequest
	GetCommand() *string
	SetParam(v string) *LlmFullDuplexCallOperateRequest
	GetParam() *string
}

type LlmFullDuplexCallOperateRequest struct {
	// 通话的唯一回执 ID。可通过 llmSmartCallFullDuplex 接口获取。
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// 动作指令：play / flush / hangup / sendDtmf
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// 扩展参数，JSON 字符串。各 command 参数说明：
	//
	// example:
	//
	// 示例值示例值
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
}

func (s LlmFullDuplexCallOperateRequest) String() string {
	return dara.Prettify(s)
}

func (s LlmFullDuplexCallOperateRequest) GoString() string {
	return s.String()
}

func (s *LlmFullDuplexCallOperateRequest) GetCallId() *string {
	return s.CallId
}

func (s *LlmFullDuplexCallOperateRequest) GetCommand() *string {
	return s.Command
}

func (s *LlmFullDuplexCallOperateRequest) GetParam() *string {
	return s.Param
}

func (s *LlmFullDuplexCallOperateRequest) SetCallId(v string) *LlmFullDuplexCallOperateRequest {
	s.CallId = &v
	return s
}

func (s *LlmFullDuplexCallOperateRequest) SetCommand(v string) *LlmFullDuplexCallOperateRequest {
	s.Command = &v
	return s
}

func (s *LlmFullDuplexCallOperateRequest) SetParam(v string) *LlmFullDuplexCallOperateRequest {
	s.Param = &v
	return s
}

func (s *LlmFullDuplexCallOperateRequest) Validate() error {
	return dara.Validate(s)
}
