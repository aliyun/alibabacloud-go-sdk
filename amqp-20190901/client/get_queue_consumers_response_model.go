// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueueConsumersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQueueConsumersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQueueConsumersResponse
	GetStatusCode() *int32
	SetBody(v *GetQueueConsumersResponseBody) *GetQueueConsumersResponse
	GetBody() *GetQueueConsumersResponseBody
}

type GetQueueConsumersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQueueConsumersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQueueConsumersResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQueueConsumersResponse) GoString() string {
	return s.String()
}

func (s *GetQueueConsumersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQueueConsumersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQueueConsumersResponse) GetBody() *GetQueueConsumersResponseBody {
	return s.Body
}

func (s *GetQueueConsumersResponse) SetHeaders(v map[string]*string) *GetQueueConsumersResponse {
	s.Headers = v
	return s
}

func (s *GetQueueConsumersResponse) SetStatusCode(v int32) *GetQueueConsumersResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueueConsumersResponse) SetBody(v *GetQueueConsumersResponseBody) *GetQueueConsumersResponse {
	s.Body = v
	return s
}

func (s *GetQueueConsumersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
