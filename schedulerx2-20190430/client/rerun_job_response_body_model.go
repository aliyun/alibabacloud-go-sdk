// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RerunJobResponseBody
	GetCode() *int32
	SetMessage(v string) *RerunJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *RerunJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RerunJobResponseBody
	GetSuccess() *bool
}

type RerunJobResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// jobId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
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

func (s RerunJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RerunJobResponseBody) GoString() string {
	return s.String()
}

func (s *RerunJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RerunJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RerunJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RerunJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RerunJobResponseBody) SetCode(v int32) *RerunJobResponseBody {
	s.Code = &v
	return s
}

func (s *RerunJobResponseBody) SetMessage(v string) *RerunJobResponseBody {
	s.Message = &v
	return s
}

func (s *RerunJobResponseBody) SetRequestId(v string) *RerunJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *RerunJobResponseBody) SetSuccess(v bool) *RerunJobResponseBody {
	s.Success = &v
	return s
}

func (s *RerunJobResponseBody) Validate() error {
	return dara.Validate(s)
}
