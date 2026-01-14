// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentRecoverCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *AgentRecoverCallResponseBody
	GetCode() *int64
	SetMessage(v string) *AgentRecoverCallResponseBody
	GetMessage() *string
	SetModel(v *AgentRecoverCallResponseBodyModel) *AgentRecoverCallResponseBody
	GetModel() *AgentRecoverCallResponseBodyModel
	SetRequestId(v string) *AgentRecoverCallResponseBody
	GetRequestId() *string
	SetSuccess(v string) *AgentRecoverCallResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *AgentRecoverCallResponseBody
	GetTimestamp() *int64
}

type AgentRecoverCallResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AgentRecoverCallResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 8EFC6D10-307B-1ECA-A8C6-7CBDF776AAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 坐席标签
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1683440860035
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AgentRecoverCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AgentRecoverCallResponseBody) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *AgentRecoverCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AgentRecoverCallResponseBody) GetModel() *AgentRecoverCallResponseBodyModel {
	return s.Model
}

func (s *AgentRecoverCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentRecoverCallResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *AgentRecoverCallResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AgentRecoverCallResponseBody) SetCode(v int64) *AgentRecoverCallResponseBody {
	s.Code = &v
	return s
}

func (s *AgentRecoverCallResponseBody) SetMessage(v string) *AgentRecoverCallResponseBody {
	s.Message = &v
	return s
}

func (s *AgentRecoverCallResponseBody) SetModel(v *AgentRecoverCallResponseBodyModel) *AgentRecoverCallResponseBody {
	s.Model = v
	return s
}

func (s *AgentRecoverCallResponseBody) SetRequestId(v string) *AgentRecoverCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgentRecoverCallResponseBody) SetSuccess(v string) *AgentRecoverCallResponseBody {
	s.Success = &v
	return s
}

func (s *AgentRecoverCallResponseBody) SetTimestamp(v int64) *AgentRecoverCallResponseBody {
	s.Timestamp = &v
	return s
}

func (s *AgentRecoverCallResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentRecoverCallResponseBodyModel struct {
	// 错误手机列表
	UnHandleNumbers []*string `json:"UnHandleNumbers,omitempty" xml:"UnHandleNumbers,omitempty" type:"Repeated"`
}

func (s AgentRecoverCallResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s AgentRecoverCallResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AgentRecoverCallResponseBodyModel) GetUnHandleNumbers() []*string {
	return s.UnHandleNumbers
}

func (s *AgentRecoverCallResponseBodyModel) SetUnHandleNumbers(v []*string) *AgentRecoverCallResponseBodyModel {
	s.UnHandleNumbers = v
	return s
}

func (s *AgentRecoverCallResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
