// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackLibraryReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FeedbackLibraryReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FeedbackLibraryReviewResponse
	GetStatusCode() *int32
	SetBody(v *LibraryReview) *FeedbackLibraryReviewResponse
	GetBody() *LibraryReview
}

type FeedbackLibraryReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LibraryReview     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FeedbackLibraryReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s FeedbackLibraryReviewResponse) GoString() string {
	return s.String()
}

func (s *FeedbackLibraryReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FeedbackLibraryReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FeedbackLibraryReviewResponse) GetBody() *LibraryReview {
	return s.Body
}

func (s *FeedbackLibraryReviewResponse) SetHeaders(v map[string]*string) *FeedbackLibraryReviewResponse {
	s.Headers = v
	return s
}

func (s *FeedbackLibraryReviewResponse) SetStatusCode(v int32) *FeedbackLibraryReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *FeedbackLibraryReviewResponse) SetBody(v *LibraryReview) *FeedbackLibraryReviewResponse {
	s.Body = v
	return s
}

func (s *FeedbackLibraryReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
