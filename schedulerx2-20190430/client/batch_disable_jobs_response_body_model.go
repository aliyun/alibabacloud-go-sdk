// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDisableJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BatchDisableJobsResponseBody
	GetCode() *int32
	SetMessage(v string) *BatchDisableJobsResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchDisableJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchDisableJobsResponseBody
	GetSuccess() *bool
}

type BatchDisableJobsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// A message that provides more information about the response.
	//
	// example:
	//
	// disable failed jobs=[99341]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 71BCC0E3-64B2-4B63-A870-AFB64EBCB5A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the operation was successful. Valid values:
	//
	// - **true**: The operation was successful.
	//
	// - **false**: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDisableJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchDisableJobsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDisableJobsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BatchDisableJobsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchDisableJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchDisableJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchDisableJobsResponseBody) SetCode(v int32) *BatchDisableJobsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDisableJobsResponseBody) SetMessage(v string) *BatchDisableJobsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDisableJobsResponseBody) SetRequestId(v string) *BatchDisableJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDisableJobsResponseBody) SetSuccess(v bool) *BatchDisableJobsResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDisableJobsResponseBody) Validate() error {
	return dara.Validate(s)
}
