// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateBucketDeleteTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TerminateBucketDeleteTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TerminateBucketDeleteTaskResponseBody
	GetSuccess() *bool
}

type TerminateBucketDeleteTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TerminateBucketDeleteTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TerminateBucketDeleteTaskResponseBody) GoString() string {
	return s.String()
}

func (s *TerminateBucketDeleteTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TerminateBucketDeleteTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TerminateBucketDeleteTaskResponseBody) SetRequestId(v string) *TerminateBucketDeleteTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *TerminateBucketDeleteTaskResponseBody) SetSuccess(v bool) *TerminateBucketDeleteTaskResponseBody {
	s.Success = &v
	return s
}

func (s *TerminateBucketDeleteTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
