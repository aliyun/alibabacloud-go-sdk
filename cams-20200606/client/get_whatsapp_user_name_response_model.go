// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhatsappUserNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWhatsappUserNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWhatsappUserNameResponse
	GetStatusCode() *int32
	SetBody(v *GetWhatsappUserNameResponseBody) *GetWhatsappUserNameResponse
	GetBody() *GetWhatsappUserNameResponseBody
}

type GetWhatsappUserNameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWhatsappUserNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWhatsappUserNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWhatsappUserNameResponse) GoString() string {
	return s.String()
}

func (s *GetWhatsappUserNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWhatsappUserNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWhatsappUserNameResponse) GetBody() *GetWhatsappUserNameResponseBody {
	return s.Body
}

func (s *GetWhatsappUserNameResponse) SetHeaders(v map[string]*string) *GetWhatsappUserNameResponse {
	s.Headers = v
	return s
}

func (s *GetWhatsappUserNameResponse) SetStatusCode(v int32) *GetWhatsappUserNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWhatsappUserNameResponse) SetBody(v *GetWhatsappUserNameResponseBody) *GetWhatsappUserNameResponse {
	s.Body = v
	return s
}

func (s *GetWhatsappUserNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
