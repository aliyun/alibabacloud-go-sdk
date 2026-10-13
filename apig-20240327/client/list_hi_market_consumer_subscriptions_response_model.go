// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHiMarketConsumerSubscriptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHiMarketConsumerSubscriptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHiMarketConsumerSubscriptionsResponse
	GetStatusCode() *int32
	SetBody(v *ListHiMarketConsumerSubscriptionsResponseBody) *ListHiMarketConsumerSubscriptionsResponse
	GetBody() *ListHiMarketConsumerSubscriptionsResponseBody
}

type ListHiMarketConsumerSubscriptionsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHiMarketConsumerSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHiMarketConsumerSubscriptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumerSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumerSubscriptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHiMarketConsumerSubscriptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHiMarketConsumerSubscriptionsResponse) GetBody() *ListHiMarketConsumerSubscriptionsResponseBody {
	return s.Body
}

func (s *ListHiMarketConsumerSubscriptionsResponse) SetHeaders(v map[string]*string) *ListHiMarketConsumerSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponse) SetStatusCode(v int32) *ListHiMarketConsumerSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponse) SetBody(v *ListHiMarketConsumerSubscriptionsResponseBody) *ListHiMarketConsumerSubscriptionsResponse {
	s.Body = v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
