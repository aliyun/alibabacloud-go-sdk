// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWebReviewPointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWebReviewPointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWebReviewPointsResponse
	GetStatusCode() *int32
	SetBody(v *ListWebReviewPointsResponseBody) *ListWebReviewPointsResponse
	GetBody() *ListWebReviewPointsResponseBody
}

type ListWebReviewPointsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWebReviewPointsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWebReviewPointsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWebReviewPointsResponse) GoString() string {
	return s.String()
}

func (s *ListWebReviewPointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWebReviewPointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWebReviewPointsResponse) GetBody() *ListWebReviewPointsResponseBody {
	return s.Body
}

func (s *ListWebReviewPointsResponse) SetHeaders(v map[string]*string) *ListWebReviewPointsResponse {
	s.Headers = v
	return s
}

func (s *ListWebReviewPointsResponse) SetStatusCode(v int32) *ListWebReviewPointsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWebReviewPointsResponse) SetBody(v *ListWebReviewPointsResponseBody) *ListWebReviewPointsResponse {
	s.Body = v
	return s
}

func (s *ListWebReviewPointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
