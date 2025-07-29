// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEnableJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BatchEnableJobsResponseBody
	GetCode() *int32
	SetMessage(v string) *BatchEnableJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchEnableJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchEnableJobsResponseBody
	GetSuccess() *bool
}

type BatchEnableJobsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned additional information.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the jobs were enabled at a time. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchEnableJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchEnableJobsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchEnableJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BatchEnableJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchEnableJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchEnableJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchEnableJobsResponseBody) SetCode(v int32) *BatchEnableJobsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchEnableJobsResponseBody) SetMessage(v string) *BatchEnableJobsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchEnableJobsResponseBody) SetRequestId(v string) *BatchEnableJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchEnableJobsResponseBody) SetSuccess(v bool) *BatchEnableJobsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchEnableJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
