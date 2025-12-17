// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSanityCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSanityCheckTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateSanityCheckTaskResponseBody
	GetTaskId() *string
}

type CreateSanityCheckTaskResponseBody struct {
	// example:
	//
	// BEBDF2EE-642E-5992-8907-D2011A7ACEFE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 917479ff-c869-49ea-908e-ae85bd987bc0
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateSanityCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSanityCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSanityCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSanityCheckTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateSanityCheckTaskResponseBody) SetRequestId(v string) *CreateSanityCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSanityCheckTaskResponseBody) SetTaskId(v string) *CreateSanityCheckTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateSanityCheckTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
