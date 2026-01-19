// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendNotificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendNotificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendNotificationResponse
	GetStatusCode() *int32
	SetBody(v *SendNotificationResponseBody) *SendNotificationResponse
	GetBody() *SendNotificationResponseBody
}

type SendNotificationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendNotificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendNotificationResponse) String() string {
	return dara.Prettify(s)
}

func (s SendNotificationResponse) GoString() string {
	return s.String()
}

func (s *SendNotificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendNotificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendNotificationResponse) GetBody() *SendNotificationResponseBody {
	return s.Body
}

func (s *SendNotificationResponse) SetHeaders(v map[string]*string) *SendNotificationResponse {
	s.Headers = v
	return s
}

func (s *SendNotificationResponse) SetStatusCode(v int32) *SendNotificationResponse {
	s.StatusCode = &v
	return s
}

func (s *SendNotificationResponse) SetBody(v *SendNotificationResponseBody) *SendNotificationResponse {
	s.Body = v
	return s
}

func (s *SendNotificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
