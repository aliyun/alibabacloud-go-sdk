// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLibraryReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLibraryReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLibraryReviewResponse
	GetStatusCode() *int32
	SetBody(v *LibraryReview) *CreateLibraryReviewResponse
	GetBody() *LibraryReview
}

type CreateLibraryReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LibraryReview     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLibraryReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLibraryReviewResponse) GoString() string {
	return s.String()
}

func (s *CreateLibraryReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLibraryReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLibraryReviewResponse) GetBody() *LibraryReview {
	return s.Body
}

func (s *CreateLibraryReviewResponse) SetHeaders(v map[string]*string) *CreateLibraryReviewResponse {
	s.Headers = v
	return s
}

func (s *CreateLibraryReviewResponse) SetStatusCode(v int32) *CreateLibraryReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLibraryReviewResponse) SetBody(v *LibraryReview) *CreateLibraryReviewResponse {
	s.Body = v
	return s
}

func (s *CreateLibraryReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
