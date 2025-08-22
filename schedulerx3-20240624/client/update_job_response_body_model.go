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
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3808cf26-dde2-4286-8503-b0a2cd4065a7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
