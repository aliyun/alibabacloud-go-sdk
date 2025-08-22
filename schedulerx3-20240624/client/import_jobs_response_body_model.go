// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImportJobsResponseBody
	GetCode() *int32
	SetMessage(v string) *ImportJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportJobsResponseBody
	GetSuccess() *bool
}

type ImportJobsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9A48E22F-F30A-5CE5-AC7A-E0FED1B6942E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ImportJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImportJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportJobsResponseBody) SetCode(v int32) *ImportJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ImportJobsResponseBody) SetMessage(v string) *ImportJobsResponseBody {
	s.Message = &v
	return s
}

func (s *ImportJobsResponseBody) SetRequestId(v string) *ImportJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportJobsResponseBody) SetSuccess(v bool) *ImportJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ImportJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
