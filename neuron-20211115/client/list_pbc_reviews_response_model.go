// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPbcReviewsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPbcReviewsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPbcReviewsResponse
	GetStatusCode() *int32
	SetBody(v *ListPbcReviewsResponseBody) *ListPbcReviewsResponse
	GetBody() *ListPbcReviewsResponseBody
}

type ListPbcReviewsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPbcReviewsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPbcReviewsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPbcReviewsResponse) GoString() string {
	return s.String()
}

func (s *ListPbcReviewsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPbcReviewsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPbcReviewsResponse) GetBody() *ListPbcReviewsResponseBody {
	return s.Body
}

func (s *ListPbcReviewsResponse) SetHeaders(v map[string]*string) *ListPbcReviewsResponse {
	s.Headers = v
	return s
}

func (s *ListPbcReviewsResponse) SetStatusCode(v int32) *ListPbcReviewsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPbcReviewsResponse) SetBody(v *ListPbcReviewsResponseBody) *ListPbcReviewsResponse {
	s.Body = v
	return s
}

func (s *ListPbcReviewsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
