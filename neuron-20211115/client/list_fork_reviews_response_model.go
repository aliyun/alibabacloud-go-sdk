// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForkReviewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListForkReviewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListForkReviewsResponse
	GetStatusCode() *int32
	SetBody(v *ForkReviewListResult) *ListForkReviewsResponse
	GetBody() *ForkReviewListResult
}

type ListForkReviewsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForkReviewListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListForkReviewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListForkReviewsResponse) GoString() string {
	return s.String()
}

func (s *ListForkReviewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListForkReviewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListForkReviewsResponse) GetBody() *ForkReviewListResult {
	return s.Body
}

func (s *ListForkReviewsResponse) SetHeaders(v map[string]*string) *ListForkReviewsResponse {
	s.Headers = v
	return s
}

func (s *ListForkReviewsResponse) SetStatusCode(v int32) *ListForkReviewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListForkReviewsResponse) SetBody(v *ForkReviewListResult) *ListForkReviewsResponse {
	s.Body = v
	return s
}

func (s *ListForkReviewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
