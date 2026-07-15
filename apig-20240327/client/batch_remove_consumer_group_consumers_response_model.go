// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemoveConsumerGroupConsumersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchRemoveConsumerGroupConsumersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchRemoveConsumerGroupConsumersResponse
	GetStatusCode() *int32
	SetBody(v *BatchRemoveConsumerGroupConsumersResponseBody) *BatchRemoveConsumerGroupConsumersResponse
	GetBody() *BatchRemoveConsumerGroupConsumersResponseBody
}

type BatchRemoveConsumerGroupConsumersResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRemoveConsumerGroupConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRemoveConsumerGroupConsumersResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveConsumerGroupConsumersResponse) GoString() string {
	return s.String()
}

func (s *BatchRemoveConsumerGroupConsumersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchRemoveConsumerGroupConsumersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchRemoveConsumerGroupConsumersResponse) GetBody() *BatchRemoveConsumerGroupConsumersResponseBody {
	return s.Body
}

func (s *BatchRemoveConsumerGroupConsumersResponse) SetHeaders(v map[string]*string) *BatchRemoveConsumerGroupConsumersResponse {
	s.Headers = v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponse) SetStatusCode(v int32) *BatchRemoveConsumerGroupConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponse) SetBody(v *BatchRemoveConsumerGroupConsumersResponseBody) *BatchRemoveConsumerGroupConsumersResponse {
	s.Body = v
	return s
}

func (s *BatchRemoveConsumerGroupConsumersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
