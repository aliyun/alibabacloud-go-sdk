// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConsumerGroupSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConsumerGroupSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConsumerGroupSubscriptionResponseBody) *DeleteConsumerGroupSubscriptionResponse
	GetBody() *DeleteConsumerGroupSubscriptionResponseBody
}

type DeleteConsumerGroupSubscriptionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConsumerGroupSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConsumerGroupSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConsumerGroupSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConsumerGroupSubscriptionResponse) GetBody() *DeleteConsumerGroupSubscriptionResponseBody {
	return s.Body
}

func (s *DeleteConsumerGroupSubscriptionResponse) SetHeaders(v map[string]*string) *DeleteConsumerGroupSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponse) SetStatusCode(v int32) *DeleteConsumerGroupSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponse) SetBody(v *DeleteConsumerGroupSubscriptionResponseBody) *DeleteConsumerGroupSubscriptionResponse {
	s.Body = v
	return s
}

func (s *DeleteConsumerGroupSubscriptionResponse) Validate() error {
	return dara.Validate(s)
}
