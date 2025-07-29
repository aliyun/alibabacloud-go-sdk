// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BatchDeleteJobsResponseBody
	GetCode() *int32
	SetMessage(v string) *BatchDeleteJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDeleteJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchDeleteJobsResponseBody
	GetSuccess() *bool
}

type BatchDeleteJobsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether multiple jobs were deleted at a time. Valid values:
	//
	// 	- **true**: Multiple jobs were deleted at a time.
	//
	// 	- **false**: Multiple jobs were not deleted at a time.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeleteJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BatchDeleteJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDeleteJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDeleteJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDeleteJobsResponseBody) SetCode(v int32) *BatchDeleteJobsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteJobsResponseBody) SetMessage(v string) *BatchDeleteJobsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteJobsResponseBody) SetRequestId(v string) *BatchDeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteJobsResponseBody) SetSuccess(v bool) *BatchDeleteJobsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeleteJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
