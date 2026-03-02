// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLibraryReviewListResult interface {
	dara.Model
	String() string
	GoString() string
	SetLibraryReviews(v []*LibraryReview) *LibraryReviewListResult
	GetLibraryReviews() []*LibraryReview
	SetRequestId(v string) *LibraryReviewListResult
	GetRequestId() *string
}

type LibraryReviewListResult struct {
	LibraryReviews []*LibraryReview `json:"libraryReviews,omitempty" xml:"libraryReviews,omitempty" type:"Repeated"`
	RequestId      *string          `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s LibraryReviewListResult) String() string {
	return dara.Prettify(s)
}

func (s LibraryReviewListResult) GoString() string {
	return s.String()
}

func (s *LibraryReviewListResult) GetLibraryReviews() []*LibraryReview {
	return s.LibraryReviews
}

func (s *LibraryReviewListResult) GetRequestId() *string {
	return s.RequestId
}

func (s *LibraryReviewListResult) SetLibraryReviews(v []*LibraryReview) *LibraryReviewListResult {
	s.LibraryReviews = v
	return s
}

func (s *LibraryReviewListResult) SetRequestId(v string) *LibraryReviewListResult {
	s.RequestId = &v
	return s
}

func (s *LibraryReviewListResult) Validate() error {
	if s.LibraryReviews != nil {
		for _, item := range s.LibraryReviews {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
