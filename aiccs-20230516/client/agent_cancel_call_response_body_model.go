// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentCancelCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *AgentCancelCallResponseBody
	GetCode() *int64
	SetMessage(v string) *AgentCancelCallResponseBody
	GetMessage() *string
	SetModel(v *AgentCancelCallResponseBodyModel) *AgentCancelCallResponseBody
	GetModel() *AgentCancelCallResponseBodyModel
	SetRequestId(v string) *AgentCancelCallResponseBody
	GetRequestId() *string
	SetSuccess(v string) *AgentCancelCallResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *AgentCancelCallResponseBody
	GetTimestamp() *int64
}

type AgentCancelCallResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AgentCancelCallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AgentCancelCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AgentCancelCallResponseBody) GoString() string {
	return s.String()
}

func (s *AgentCancelCallResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AgentCancelCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AgentCancelCallResponseBody) GetModel() *AgentCancelCallResponseBodyModel {
	return s.Model
}

func (s *AgentCancelCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentCancelCallResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AgentCancelCallResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AgentCancelCallResponseBody) SetCode(v int64) *AgentCancelCallResponseBody {
	s.Code = &v
	return s
}

func (s *AgentCancelCallResponseBody) SetMessage(v string) *AgentCancelCallResponseBody {
	s.Message = &v
	return s
}

func (s *AgentCancelCallResponseBody) SetModel(v *AgentCancelCallResponseBodyModel) *AgentCancelCallResponseBody {
	s.Model = v
	return s
}

func (s *AgentCancelCallResponseBody) SetRequestId(v string) *AgentCancelCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgentCancelCallResponseBody) SetSuccess(v string) *AgentCancelCallResponseBody {
	s.Success = &v
	return s
}

func (s *AgentCancelCallResponseBody) SetTimestamp(v int64) *AgentCancelCallResponseBody {
	s.Timestamp = &v
	return s
}

func (s *AgentCancelCallResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentCancelCallResponseBodyModel struct {
	// 错误手机列表
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s AgentCancelCallResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s AgentCancelCallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AgentCancelCallResponseBodyModel) GetUnHandleNumbers() []*string {
	return s.UnHandleNumbers
}

func (s *AgentCancelCallResponseBodyModel) SetUnHandleNumbers(v []*string) *AgentCancelCallResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

func (s *AgentCancelCallResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
