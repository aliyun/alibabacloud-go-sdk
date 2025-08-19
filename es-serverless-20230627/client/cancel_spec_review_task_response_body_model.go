// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSpecReviewTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelSpecReviewTaskResponseBody
	GetRequestId() *string
	SetResult(v bool) *CancelSpecReviewTaskResponseBody
	GetResult() *bool
}

type CancelSpecReviewTaskResponseBody struct {
	// example:
	//
	// 1B64F3E0-25D5-5043-B5C8-4FF22BB12CCD
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CancelSpecReviewTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelSpecReviewTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelSpecReviewTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelSpecReviewTaskResponseBody) GetResult() *bool {
	return s.Result
}

func (s *CancelSpecReviewTaskResponseBody) SetRequestId(v string) *CancelSpecReviewTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelSpecReviewTaskResponseBody) SetResult(v bool) *CancelSpecReviewTaskResponseBody {
	s.Result = &v
	return s
}

func (s *CancelSpecReviewTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
