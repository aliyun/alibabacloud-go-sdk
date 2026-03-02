// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForkReviewListResult interface {
	dara.Model
	String() string
	GoString() string
	SetForkReviews(v []*ForkReview) *ForkReviewListResult
	GetForkReviews() []*ForkReview
	SetRequestId(v string) *ForkReviewListResult
	GetRequestId() *string
	SetTotal(v int32) *ForkReviewListResult
	GetTotal() *int32
}

type ForkReviewListResult struct {
	ForkReviews []*ForkReview `json:"forkReviews,omitempty" xml:"forkReviews,omitempty" type:"Repeated"`
	RequestId   *string       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 24
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ForkReviewListResult) String() string {
	return dara.Prettify(s)
}

func (s ForkReviewListResult) GoString() string {
	return s.String()
}

func (s *ForkReviewListResult) GetForkReviews() []*ForkReview {
	return s.ForkReviews
}

func (s *ForkReviewListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ForkReviewListResult) GetTotal() *int32 {
	return s.Total
}

func (s *ForkReviewListResult) SetForkReviews(v []*ForkReview) *ForkReviewListResult {
	s.ForkReviews = v
	return s
}

func (s *ForkReviewListResult) SetRequestId(v string) *ForkReviewListResult {
	s.RequestId = &v
	return s
}

func (s *ForkReviewListResult) SetTotal(v int32) *ForkReviewListResult {
	s.Total = &v
	return s
}

func (s *ForkReviewListResult) Validate() error {
	if s.ForkReviews != nil {
		for _, item := range s.ForkReviews {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
