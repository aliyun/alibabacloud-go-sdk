// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateJobResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateJobResponseBody
	GetSuccess() *bool
}

type UpdateJobResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned only if an error occurs.
	//
	// example:
	//
	// job type is java className can not be blank
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateJobResponseBody) SetCode(v int32) *UpdateJobResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateJobResponseBody) SetMessage(v string) *UpdateJobResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateJobResponseBody) SetRequestId(v string) *UpdateJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobResponseBody) SetSuccess(v bool) *UpdateJobResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateJobResponseBody) Validate() error {
	return dara.Validate(s)
}
