// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeReviewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCodeReviewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCodeReviewsResponse
	GetStatusCode() *int32
	SetBody(v *ListCodeReviewsResponseBody) *ListCodeReviewsResponse
	GetBody() *ListCodeReviewsResponseBody
}

type ListCodeReviewsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCodeReviewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCodeReviewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCodeReviewsResponse) GoString() string {
	return s.String()
}

func (s *ListCodeReviewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCodeReviewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCodeReviewsResponse) GetBody() *ListCodeReviewsResponseBody {
	return s.Body
}

func (s *ListCodeReviewsResponse) SetHeaders(v map[string]*string) *ListCodeReviewsResponse {
	s.Headers = v
	return s
}

func (s *ListCodeReviewsResponse) SetStatusCode(v int32) *ListCodeReviewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCodeReviewsResponse) SetBody(v *ListCodeReviewsResponseBody) *ListCodeReviewsResponse {
	s.Body = v
	return s
}

func (s *ListCodeReviewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
