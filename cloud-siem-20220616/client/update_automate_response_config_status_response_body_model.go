// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAutomateResponseConfigStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateAutomateResponseConfigStatusResponseBody
	GetCode() *int32
	SetData(v string) *UpdateAutomateResponseConfigStatusResponseBody
	GetData() *string
	SetMessage(v string) *UpdateAutomateResponseConfigStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAutomateResponseConfigStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAutomateResponseConfigStatusResponseBody
	GetSuccess() *bool
}

type UpdateAutomateResponseConfigStatusResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAutomateResponseConfigStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAutomateResponseConfigStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetCode(v int32) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetData(v string) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetMessage(v string) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetRequestId(v string) *UpdateAutomateResponseConfigStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) SetSuccess(v bool) *UpdateAutomateResponseConfigStatusResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAutomateResponseConfigStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
