// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTopicSubscriptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTopicSubscriptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTopicSubscriptionsResponse
	GetStatusCode() *int32
	SetBody(v *ListTopicSubscriptionsResponseBody) *ListTopicSubscriptionsResponse
	GetBody() *ListTopicSubscriptionsResponseBody
}

type ListTopicSubscriptionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTopicSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTopicSubscriptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTopicSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ListTopicSubscriptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTopicSubscriptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTopicSubscriptionsResponse) GetBody() *ListTopicSubscriptionsResponseBody {
	return s.Body
}

func (s *ListTopicSubscriptionsResponse) SetHeaders(v map[string]*string) *ListTopicSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ListTopicSubscriptionsResponse) SetStatusCode(v int32) *ListTopicSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTopicSubscriptionsResponse) SetBody(v *ListTopicSubscriptionsResponseBody) *ListTopicSubscriptionsResponse {
	s.Body = v
	return s
}

func (s *ListTopicSubscriptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
