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
	// The unique receipt ID of the call. You can obtain this ID by calling the LlmSmartCallFullDuplex operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	CallId *string `json:"CallId,omitempty" xml:"CallId,omitempty"`
	// The action command: play / flush / hangup / sendDtmf.
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The extension parameter, a JSON character string. The metric description for each command:
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
