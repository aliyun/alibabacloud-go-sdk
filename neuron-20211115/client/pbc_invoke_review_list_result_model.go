// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcInvokeReviewListResult interface {
	dara.Model
	String() string
	GoString() string
	SetPbcInvokeReviews(v []*PbcInvokeReview) *PbcInvokeReviewListResult
	GetPbcInvokeReviews() []*PbcInvokeReview
	SetRequestId(v string) *PbcInvokeReviewListResult
	GetRequestId() *string
}

type PbcInvokeReviewListResult struct {
	PbcInvokeReviews []*PbcInvokeReview `json:"pbcInvokeReviews,omitempty" xml:"pbcInvokeReviews,omitempty" type:"Repeated"`
	RequestId        *string            `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PbcInvokeReviewListResult) String() string {
	return dara.Prettify(s)
}

func (s PbcInvokeReviewListResult) GoString() string {
	return s.String()
}

func (s *PbcInvokeReviewListResult) GetPbcInvokeReviews() []*PbcInvokeReview {
	return s.PbcInvokeReviews
}

func (s *PbcInvokeReviewListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcInvokeReviewListResult) SetPbcInvokeReviews(v []*PbcInvokeReview) *PbcInvokeReviewListResult {
	s.PbcInvokeReviews = v
	return s
}

func (s *PbcInvokeReviewListResult) SetRequestId(v string) *PbcInvokeReviewListResult {
	s.RequestId = &v
	return s
}

func (s *PbcInvokeReviewListResult) Validate() error {
	if s.PbcInvokeReviews != nil {
		for _, item := range s.PbcInvokeReviews {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
