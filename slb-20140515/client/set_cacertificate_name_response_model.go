// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCACertificateNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCACertificateNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCACertificateNameResponse
	GetStatusCode() *int32
	SetBody(v *SetCACertificateNameResponseBody) *SetCACertificateNameResponse
	GetBody() *SetCACertificateNameResponseBody
}

type SetCACertificateNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCACertificateNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCACertificateNameResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCACertificateNameResponse) GoString() string {
	return s.String()
}

func (s *SetCACertificateNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCACertificateNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCACertificateNameResponse) GetBody() *SetCACertificateNameResponseBody {
	return s.Body
}

func (s *SetCACertificateNameResponse) SetHeaders(v map[string]*string) *SetCACertificateNameResponse {
	s.Headers = v
	return s
}

func (s *SetCACertificateNameResponse) SetStatusCode(v int32) *SetCACertificateNameResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCACertificateNameResponse) SetBody(v *SetCACertificateNameResponseBody) *SetCACertificateNameResponse {
	s.Body = v
	return s
}

func (s *SetCACertificateNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
