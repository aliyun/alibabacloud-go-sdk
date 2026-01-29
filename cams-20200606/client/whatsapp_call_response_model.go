// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhatsappCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *WhatsappCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *WhatsappCallResponse
	GetStatusCode() *int32
	SetBody(v *WhatsappCallResponseBody) *WhatsappCallResponse
	GetBody() *WhatsappCallResponseBody
}

type WhatsappCallResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WhatsappCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s WhatsappCallResponse) String() string {
	return dara.Prettify(s)
}

func (s WhatsappCallResponse) GoString() string {
	return s.String()
}

func (s *WhatsappCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *WhatsappCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *WhatsappCallResponse) GetBody() *WhatsappCallResponseBody {
	return s.Body
}

func (s *WhatsappCallResponse) SetHeaders(v map[string]*string) *WhatsappCallResponse {
	s.Headers = v
	return s
}

func (s *WhatsappCallResponse) SetStatusCode(v int32) *WhatsappCallResponse {
	s.StatusCode = &v
	return s
}

func (s *WhatsappCallResponse) SetBody(v *WhatsappCallResponseBody) *WhatsappCallResponse {
	s.Body = v
	return s
}

func (s *WhatsappCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
