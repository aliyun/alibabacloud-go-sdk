// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueueConsumersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueueConsumersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueueConsumersResponse
	GetStatusCode() *int32
	SetBody(v *ListQueueConsumersResponseBody) *ListQueueConsumersResponse
	GetBody() *ListQueueConsumersResponseBody
}

type ListQueueConsumersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueueConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueueConsumersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueueConsumersResponse) GoString() string {
	return s.String()
}

func (s *ListQueueConsumersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueueConsumersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueueConsumersResponse) GetBody() *ListQueueConsumersResponseBody {
	return s.Body
}

func (s *ListQueueConsumersResponse) SetHeaders(v map[string]*string) *ListQueueConsumersResponse {
	s.Headers = v
	return s
}

func (s *ListQueueConsumersResponse) SetStatusCode(v int32) *ListQueueConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueueConsumersResponse) SetBody(v *ListQueueConsumersResponseBody) *ListQueueConsumersResponse {
	s.Body = v
	return s
}

func (s *ListQueueConsumersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
