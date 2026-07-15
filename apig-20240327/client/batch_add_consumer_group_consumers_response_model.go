// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddConsumerGroupConsumersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchAddConsumerGroupConsumersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchAddConsumerGroupConsumersResponse
	GetStatusCode() *int32
	SetBody(v *BatchAddConsumerGroupConsumersResponseBody) *BatchAddConsumerGroupConsumersResponse
	GetBody() *BatchAddConsumerGroupConsumersResponseBody
}

type BatchAddConsumerGroupConsumersResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddConsumerGroupConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddConsumerGroupConsumersResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchAddConsumerGroupConsumersResponse) GoString() string {
	return s.String()
}

func (s *BatchAddConsumerGroupConsumersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchAddConsumerGroupConsumersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchAddConsumerGroupConsumersResponse) GetBody() *BatchAddConsumerGroupConsumersResponseBody {
	return s.Body
}

func (s *BatchAddConsumerGroupConsumersResponse) SetHeaders(v map[string]*string) *BatchAddConsumerGroupConsumersResponse {
	s.Headers = v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponse) SetStatusCode(v int32) *BatchAddConsumerGroupConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponse) SetBody(v *BatchAddConsumerGroupConsumersResponseBody) *BatchAddConsumerGroupConsumersResponse {
	s.Body = v
	return s
}

func (s *BatchAddConsumerGroupConsumersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
