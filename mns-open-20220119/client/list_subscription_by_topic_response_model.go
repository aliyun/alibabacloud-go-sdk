// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionByTopicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubscriptionByTopicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubscriptionByTopicResponse
	GetStatusCode() *int32
	SetBody(v *ListSubscriptionByTopicResponseBody) *ListSubscriptionByTopicResponse
	GetBody() *ListSubscriptionByTopicResponseBody
}

type ListSubscriptionByTopicResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubscriptionByTopicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubscriptionByTopicResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionByTopicResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionByTopicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubscriptionByTopicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubscriptionByTopicResponse) GetBody() *ListSubscriptionByTopicResponseBody {
	return s.Body
}

func (s *ListSubscriptionByTopicResponse) SetHeaders(v map[string]*string) *ListSubscriptionByTopicResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionByTopicResponse) SetStatusCode(v int32) *ListSubscriptionByTopicResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionByTopicResponse) SetBody(v *ListSubscriptionByTopicResponseBody) *ListSubscriptionByTopicResponse {
	s.Body = v
	return s
}

func (s *ListSubscriptionByTopicResponse) Validate() error {
	return dara.Validate(s)
}
