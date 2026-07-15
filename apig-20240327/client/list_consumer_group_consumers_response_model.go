// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupConsumersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumerGroupConsumersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumerGroupConsumersResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumerGroupConsumersResponseBody) *ListConsumerGroupConsumersResponse
	GetBody() *ListConsumerGroupConsumersResponseBody
}

type ListConsumerGroupConsumersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumerGroupConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumerGroupConsumersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupConsumersResponse) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupConsumersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumerGroupConsumersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumerGroupConsumersResponse) GetBody() *ListConsumerGroupConsumersResponseBody {
	return s.Body
}

func (s *ListConsumerGroupConsumersResponse) SetHeaders(v map[string]*string) *ListConsumerGroupConsumersResponse {
	s.Headers = v
	return s
}

func (s *ListConsumerGroupConsumersResponse) SetStatusCode(v int32) *ListConsumerGroupConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumerGroupConsumersResponse) SetBody(v *ListConsumerGroupConsumersResponseBody) *ListConsumerGroupConsumersResponse {
	s.Body = v
	return s
}

func (s *ListConsumerGroupConsumersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
