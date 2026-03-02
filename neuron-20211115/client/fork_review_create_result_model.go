// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForkReviewCreateResult interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ForkReviewCreateResult
	GetId() *int64
	SetRequestId(v string) *ForkReviewCreateResult
	GetRequestId() *string
}

type ForkReviewCreateResult struct {
	Id        *int64  `json:"id,omitempty" xml:"id,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ForkReviewCreateResult) String() string {
	return dara.Prettify(s)
}

func (s ForkReviewCreateResult) GoString() string {
	return s.String()
}

func (s *ForkReviewCreateResult) GetId() *int64 {
	return s.Id
}

func (s *ForkReviewCreateResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ForkReviewCreateResult) SetId(v int64) *ForkReviewCreateResult {
	s.Id = &v
	return s
}

func (s *ForkReviewCreateResult) SetRequestId(v string) *ForkReviewCreateResult {
	s.RequestId = &v
	return s
}

func (s *ForkReviewCreateResult) Validate() error {
	return dara.Validate(s)
}
