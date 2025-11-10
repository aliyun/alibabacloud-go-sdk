// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerGroupSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConsumerGroupSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConsumerGroupSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *GetConsumerGroupSubscriptionResponseBody) *GetConsumerGroupSubscriptionResponse
	GetBody() *GetConsumerGroupSubscriptionResponseBody
}

type GetConsumerGroupSubscriptionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConsumerGroupSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConsumerGroupSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerGroupSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *GetConsumerGroupSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConsumerGroupSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConsumerGroupSubscriptionResponse) GetBody() *GetConsumerGroupSubscriptionResponseBody {
	return s.Body
}

func (s *GetConsumerGroupSubscriptionResponse) SetHeaders(v map[string]*string) *GetConsumerGroupSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *GetConsumerGroupSubscriptionResponse) SetStatusCode(v int32) *GetConsumerGroupSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConsumerGroupSubscriptionResponse) SetBody(v *GetConsumerGroupSubscriptionResponseBody) *GetConsumerGroupSubscriptionResponse {
	s.Body = v
	return s
}

func (s *GetConsumerGroupSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
