// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAgentResponseBody
	GetCode() *string
	SetData(v bool) *UpdateAgentResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAgentResponseBody
	GetSuccess() *bool
}

type UpdateAgentResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message, if any.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. The caller can use this field to determine whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false/null: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAgentResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAgentResponseBody) SetCode(v string) *UpdateAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAgentResponseBody) SetData(v bool) *UpdateAgentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAgentResponseBody) SetMessage(v string) *UpdateAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAgentResponseBody) SetRequestId(v string) *UpdateAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAgentResponseBody) SetSuccess(v bool) *UpdateAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
