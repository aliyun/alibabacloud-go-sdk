// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicRecommendEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTopicRecommendEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTopicRecommendEventListResponse
	GetStatusCode() *int32
	SetBody(v *ListTopicRecommendEventListResponseBody) *ListTopicRecommendEventListResponse
	GetBody() *ListTopicRecommendEventListResponseBody
}

type ListTopicRecommendEventListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicRecommendEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicRecommendEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTopicRecommendEventListResponse) GoString() string {
	return s.String()
}

func (s *ListTopicRecommendEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTopicRecommendEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTopicRecommendEventListResponse) GetBody() *ListTopicRecommendEventListResponseBody {
	return s.Body
}

func (s *ListTopicRecommendEventListResponse) SetHeaders(v map[string]*string) *ListTopicRecommendEventListResponse {
	s.Headers = v
	return s
}

func (s *ListTopicRecommendEventListResponse) SetStatusCode(v int32) *ListTopicRecommendEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicRecommendEventListResponse) SetBody(v *ListTopicRecommendEventListResponseBody) *ListTopicRecommendEventListResponse {
	s.Body = v
	return s
}

func (s *ListTopicRecommendEventListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
