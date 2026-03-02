// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLibraryReviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLibraryReviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLibraryReviewResponse
	GetStatusCode() *int32
	SetBody(v *LibraryReview) *GetLibraryReviewResponse
	GetBody() *LibraryReview
}

type GetLibraryReviewResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LibraryReview     `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLibraryReviewResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLibraryReviewResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryReviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLibraryReviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLibraryReviewResponse) GetBody() *LibraryReview {
	return s.Body
}

func (s *GetLibraryReviewResponse) SetHeaders(v map[string]*string) *GetLibraryReviewResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryReviewResponse) SetStatusCode(v int32) *GetLibraryReviewResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLibraryReviewResponse) SetBody(v *LibraryReview) *GetLibraryReviewResponse {
	s.Body = v
	return s
}

func (s *GetLibraryReviewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
