// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhatsappUserNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWhatsappUserNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWhatsappUserNameResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWhatsappUserNameResponseBody) *DeleteWhatsappUserNameResponse
	GetBody() *DeleteWhatsappUserNameResponseBody
}

type DeleteWhatsappUserNameResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWhatsappUserNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWhatsappUserNameResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhatsappUserNameResponse) GoString() string {
	return s.String()
}

func (s *DeleteWhatsappUserNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWhatsappUserNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWhatsappUserNameResponse) GetBody() *DeleteWhatsappUserNameResponseBody {
	return s.Body
}

func (s *DeleteWhatsappUserNameResponse) SetHeaders(v map[string]*string) *DeleteWhatsappUserNameResponse {
	s.Headers = v
	return s
}

func (s *DeleteWhatsappUserNameResponse) SetStatusCode(v int32) *DeleteWhatsappUserNameResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWhatsappUserNameResponse) SetBody(v *DeleteWhatsappUserNameResponseBody) *DeleteWhatsappUserNameResponse {
	s.Body = v
	return s
}

func (s *DeleteWhatsappUserNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
