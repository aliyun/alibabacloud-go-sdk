// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMqttTicketResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMqttTicketResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMqttTicketResponse
	GetStatusCode() *int32
	SetBody(v *GetMqttTicketResponseBody) *GetMqttTicketResponse
	GetBody() *GetMqttTicketResponseBody
}

type GetMqttTicketResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMqttTicketResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMqttTicketResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMqttTicketResponse) GoString() string {
	return s.String()
}

func (s *GetMqttTicketResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMqttTicketResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMqttTicketResponse) GetBody() *GetMqttTicketResponseBody {
	return s.Body
}

func (s *GetMqttTicketResponse) SetHeaders(v map[string]*string) *GetMqttTicketResponse {
	s.Headers = v
	return s
}

func (s *GetMqttTicketResponse) SetStatusCode(v int32) *GetMqttTicketResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMqttTicketResponse) SetBody(v *GetMqttTicketResponseBody) *GetMqttTicketResponse {
	s.Body = v
	return s
}

func (s *GetMqttTicketResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
