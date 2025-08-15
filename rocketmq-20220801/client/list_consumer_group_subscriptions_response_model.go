// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupSubscriptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumerGroupSubscriptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumerGroupSubscriptionsResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumerGroupSubscriptionsResponseBody) *ListConsumerGroupSubscriptionsResponse
	GetBody() *ListConsumerGroupSubscriptionsResponseBody
}

type ListConsumerGroupSubscriptionsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerGroupSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerGroupSubscriptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupSubscriptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumerGroupSubscriptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumerGroupSubscriptionsResponse) GetBody() *ListConsumerGroupSubscriptionsResponseBody {
	return s.Body
}

func (s *ListConsumerGroupSubscriptionsResponse) SetHeaders(v map[string]*string) *ListConsumerGroupSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponse) SetStatusCode(v int32) *ListConsumerGroupSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponse) SetBody(v *ListConsumerGroupSubscriptionsResponseBody) *ListConsumerGroupSubscriptionsResponse {
	s.Body = v
	return s
}

func (s *ListConsumerGroupSubscriptionsResponse) Validate() error {
	return dara.Validate(s)
}
