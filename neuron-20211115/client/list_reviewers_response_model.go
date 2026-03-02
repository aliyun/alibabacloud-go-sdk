// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReviewersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReviewersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReviewersResponse
	GetStatusCode() *int32
	SetBody(v *ReviewerListResult) *ListReviewersResponse
	GetBody() *ReviewerListResult
}

type ListReviewersResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReviewerListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReviewersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReviewersResponse) GoString() string {
	return s.String()
}

func (s *ListReviewersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReviewersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReviewersResponse) GetBody() *ReviewerListResult {
	return s.Body
}

func (s *ListReviewersResponse) SetHeaders(v map[string]*string) *ListReviewersResponse {
	s.Headers = v
	return s
}

func (s *ListReviewersResponse) SetStatusCode(v int32) *ListReviewersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReviewersResponse) SetBody(v *ReviewerListResult) *ListReviewersResponse {
	s.Body = v
	return s
}

func (s *ListReviewersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
