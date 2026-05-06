// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCaCertificateHostnamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientCaCertificateHostnamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientCaCertificateHostnamesResponse
	GetStatusCode() *int32
	SetBody(v *GetClientCaCertificateHostnamesResponseBody) *GetClientCaCertificateHostnamesResponse
	GetBody() *GetClientCaCertificateHostnamesResponseBody
}

type GetClientCaCertificateHostnamesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientCaCertificateHostnamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientCaCertificateHostnamesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientCaCertificateHostnamesResponse) GoString() string {
	return s.String()
}

func (s *GetClientCaCertificateHostnamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientCaCertificateHostnamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientCaCertificateHostnamesResponse) GetBody() *GetClientCaCertificateHostnamesResponseBody {
	return s.Body
}

func (s *GetClientCaCertificateHostnamesResponse) SetHeaders(v map[string]*string) *GetClientCaCertificateHostnamesResponse {
	s.Headers = v
	return s
}

func (s *GetClientCaCertificateHostnamesResponse) SetStatusCode(v int32) *GetClientCaCertificateHostnamesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientCaCertificateHostnamesResponse) SetBody(v *GetClientCaCertificateHostnamesResponseBody) *GetClientCaCertificateHostnamesResponse {
	s.Body = v
	return s
}

func (s *GetClientCaCertificateHostnamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
