// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhatsappUserNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWhatsappUserNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWhatsappUserNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWhatsappUserNameResponseBody) *UpdateWhatsappUserNameResponse
	GetBody() *UpdateWhatsappUserNameResponseBody
}

type UpdateWhatsappUserNameResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWhatsappUserNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWhatsappUserNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhatsappUserNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhatsappUserNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWhatsappUserNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWhatsappUserNameResponse) GetBody() *UpdateWhatsappUserNameResponseBody {
	return s.Body
}

func (s *UpdateWhatsappUserNameResponse) SetHeaders(v map[string]*string) *UpdateWhatsappUserNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateWhatsappUserNameResponse) SetStatusCode(v int32) *UpdateWhatsappUserNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWhatsappUserNameResponse) SetBody(v *UpdateWhatsappUserNameResponseBody) *UpdateWhatsappUserNameResponse {
	s.Body = v
	return s
}

func (s *UpdateWhatsappUserNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
