// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLibraryReviewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLibraryReviewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLibraryReviewsResponse
	GetStatusCode() *int32
	SetBody(v *LibraryReviewListResult) *ListLibraryReviewsResponse
	GetBody() *LibraryReviewListResult
}

type ListLibraryReviewsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LibraryReviewListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLibraryReviewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLibraryReviewsResponse) GoString() string {
	return s.String()
}

func (s *ListLibraryReviewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLibraryReviewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLibraryReviewsResponse) GetBody() *LibraryReviewListResult {
	return s.Body
}

func (s *ListLibraryReviewsResponse) SetHeaders(v map[string]*string) *ListLibraryReviewsResponse {
	s.Headers = v
	return s
}

func (s *ListLibraryReviewsResponse) SetStatusCode(v int32) *ListLibraryReviewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLibraryReviewsResponse) SetBody(v *LibraryReviewListResult) *ListLibraryReviewsResponse {
	s.Body = v
	return s
}

func (s *ListLibraryReviewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
