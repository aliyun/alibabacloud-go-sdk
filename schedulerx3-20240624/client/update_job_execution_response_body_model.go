// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateJobExecutionResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateJobExecutionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateJobExecutionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateJobExecutionResponseBody
	GetSuccess() *bool
}

type UpdateJobExecutionResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 27B1345D-5F71-5972-8E4C-AABA6C6232F0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateJobExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobExecutionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateJobExecutionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateJobExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateJobExecutionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateJobExecutionResponseBody) SetCode(v int32) *UpdateJobExecutionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateJobExecutionResponseBody) SetMessage(v string) *UpdateJobExecutionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateJobExecutionResponseBody) SetRequestId(v string) *UpdateJobExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateJobExecutionResponseBody) SetSuccess(v bool) *UpdateJobExecutionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateJobExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
