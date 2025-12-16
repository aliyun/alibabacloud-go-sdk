// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaReviewTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaReviewTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaReviewTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaReviewTasksResponseBody) *ListQuotaReviewTasksResponse
	GetBody() *ListQuotaReviewTasksResponseBody
}

type ListQuotaReviewTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaReviewTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaReviewTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaReviewTasksResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaReviewTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaReviewTasksResponse) GetBody() *ListQuotaReviewTasksResponseBody {
	return s.Body
}

func (s *ListQuotaReviewTasksResponse) SetHeaders(v map[string]*string) *ListQuotaReviewTasksResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaReviewTasksResponse) SetStatusCode(v int32) *ListQuotaReviewTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaReviewTasksResponse) SetBody(v *ListQuotaReviewTasksResponseBody) *ListQuotaReviewTasksResponse {
	s.Body = v
	return s
}

func (s *ListQuotaReviewTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
