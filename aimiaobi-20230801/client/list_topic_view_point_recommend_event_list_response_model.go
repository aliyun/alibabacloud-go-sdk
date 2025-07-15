// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicViewPointRecommendEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTopicViewPointRecommendEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTopicViewPointRecommendEventListResponse
	GetStatusCode() *int32
	SetBody(v *ListTopicViewPointRecommendEventListResponseBody) *ListTopicViewPointRecommendEventListResponse
	GetBody() *ListTopicViewPointRecommendEventListResponseBody
}

type ListTopicViewPointRecommendEventListResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicViewPointRecommendEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicViewPointRecommendEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTopicViewPointRecommendEventListResponse) GoString() string {
	return s.String()
}

func (s *ListTopicViewPointRecommendEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTopicViewPointRecommendEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTopicViewPointRecommendEventListResponse) GetBody() *ListTopicViewPointRecommendEventListResponseBody {
	return s.Body
}

func (s *ListTopicViewPointRecommendEventListResponse) SetHeaders(v map[string]*string) *ListTopicViewPointRecommendEventListResponse {
	s.Headers = v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponse) SetStatusCode(v int32) *ListTopicViewPointRecommendEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponse) SetBody(v *ListTopicViewPointRecommendEventListResponseBody) *ListTopicViewPointRecommendEventListResponse {
	s.Body = v
	return s
}

func (s *ListTopicViewPointRecommendEventListResponse) Validate() error {
	return dara.Validate(s)
}
