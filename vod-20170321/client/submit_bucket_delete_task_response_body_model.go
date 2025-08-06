// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBucketDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitBucketDeleteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitBucketDeleteTaskResponseBody
	GetSuccess() *bool
}

type SubmitBucketDeleteTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitBucketDeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitBucketDeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBucketDeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitBucketDeleteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitBucketDeleteTaskResponseBody) SetRequestId(v string) *SubmitBucketDeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitBucketDeleteTaskResponseBody) SetSuccess(v bool) *SubmitBucketDeleteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitBucketDeleteTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
