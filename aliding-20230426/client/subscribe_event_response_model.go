// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubscribeEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubscribeEventResponse
	GetStatusCode() *int32
	SetBody(v *SubscribeEventResponseBody) *SubscribeEventResponse
	GetBody() *SubscribeEventResponseBody
}

type SubscribeEventResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscribeEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscribeEventResponse) String() string {
	return dara.Prettify(s)
}

func (s SubscribeEventResponse) GoString() string {
	return s.String()
}

func (s *SubscribeEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubscribeEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubscribeEventResponse) GetBody() *SubscribeEventResponseBody {
	return s.Body
}

func (s *SubscribeEventResponse) SetHeaders(v map[string]*string) *SubscribeEventResponse {
	s.Headers = v
	return s
}

func (s *SubscribeEventResponse) SetStatusCode(v int32) *SubscribeEventResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscribeEventResponse) SetBody(v *SubscribeEventResponseBody) *SubscribeEventResponse {
	s.Body = v
	return s
}

func (s *SubscribeEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
