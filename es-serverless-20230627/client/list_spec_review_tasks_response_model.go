// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSpecReviewTasksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSpecReviewTasksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSpecReviewTasksResponse
	GetStatusCode() *int32
	SetBody(v *ListSpecReviewTasksResponseBody) *ListSpecReviewTasksResponse
	GetBody() *ListSpecReviewTasksResponseBody
}

type ListSpecReviewTasksResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSpecReviewTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSpecReviewTasksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSpecReviewTasksResponse) GoString() string {
	return s.String()
}

func (s *ListSpecReviewTasksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSpecReviewTasksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSpecReviewTasksResponse) GetBody() *ListSpecReviewTasksResponseBody {
	return s.Body
}

func (s *ListSpecReviewTasksResponse) SetHeaders(v map[string]*string) *ListSpecReviewTasksResponse {
	s.Headers = v
	return s
}

func (s *ListSpecReviewTasksResponse) SetStatusCode(v int32) *ListSpecReviewTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSpecReviewTasksResponse) SetBody(v *ListSpecReviewTasksResponseBody) *ListSpecReviewTasksResponse {
	s.Body = v
	return s
}

func (s *ListSpecReviewTasksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
