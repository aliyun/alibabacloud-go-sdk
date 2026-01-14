// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateAgentStatusResponseBody
	GetCode() *int64
	SetMessage(v string) *UpdateAgentStatusResponseBody
	GetMessage() *string
	SetModel(v map[string]interface{}) *UpdateAgentStatusResponseBody
	GetModel() map[string]interface{}
	SetRequestId(v string) *UpdateAgentStatusResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateAgentStatusResponseBody
	GetSuccess() *string
	SetTimestamp(v int64) *UpdateAgentStatusResponseBody
	GetTimestamp() *int64
}

type UpdateAgentStatusResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   map[string]interface{} `json:"Model,omitempty" xml:"Model,omitempty"`
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

func (s UpdateAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentStatusResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateAgentStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentStatusResponseBody) GetModel() map[string]interface{} {
	return s.Model
}

func (s *UpdateAgentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentStatusResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateAgentStatusResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *UpdateAgentStatusResponseBody) SetCode(v int64) *UpdateAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetMessage(v string) *UpdateAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetModel(v map[string]interface{}) *UpdateAgentStatusResponseBody {
	s.Model = v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetRequestId(v string) *UpdateAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetSuccess(v string) *UpdateAgentStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentStatusResponseBody) SetTimestamp(v int64) *UpdateAgentStatusResponseBody {
	s.Timestamp = &v
	return s
}

func (s *UpdateAgentStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
