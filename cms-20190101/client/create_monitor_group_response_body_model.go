// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateMonitorGroupResponseBody
	GetCode() *int32
	SetGroupId(v int64) *CreateMonitorGroupResponseBody
	GetGroupId() *int64
	SetMessage(v string) *CreateMonitorGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMonitorGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateMonitorGroupResponseBody
	GetSuccess() *bool
}

type CreateMonitorGroupResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 1234567
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 83C89BA6-ABD4-4398-A175-83E86C47A001
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

func (s CreateMonitorGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateMonitorGroupResponseBody) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateMonitorGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMonitorGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMonitorGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateMonitorGroupResponseBody) SetCode(v int32) *CreateMonitorGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) SetGroupId(v int64) *CreateMonitorGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) SetMessage(v string) *CreateMonitorGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) SetRequestId(v string) *CreateMonitorGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) SetSuccess(v bool) *CreateMonitorGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMonitorGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
