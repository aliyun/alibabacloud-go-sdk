// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClientCaCertificateHostnamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetClientCaCertificateHostnamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetClientCaCertificateHostnamesResponse
	GetStatusCode() *int32
	SetBody(v *SetClientCaCertificateHostnamesResponseBody) *SetClientCaCertificateHostnamesResponse
	GetBody() *SetClientCaCertificateHostnamesResponseBody
}

type SetClientCaCertificateHostnamesResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetClientCaCertificateHostnamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetClientCaCertificateHostnamesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetClientCaCertificateHostnamesResponse) GoString() string {
	return s.String()
}

func (s *SetClientCaCertificateHostnamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetClientCaCertificateHostnamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetClientCaCertificateHostnamesResponse) GetBody() *SetClientCaCertificateHostnamesResponseBody {
	return s.Body
}

func (s *SetClientCaCertificateHostnamesResponse) SetHeaders(v map[string]*string) *SetClientCaCertificateHostnamesResponse {
	s.Headers = v
	return s
}

func (s *SetClientCaCertificateHostnamesResponse) SetStatusCode(v int32) *SetClientCaCertificateHostnamesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetClientCaCertificateHostnamesResponse) SetBody(v *SetClientCaCertificateHostnamesResponseBody) *SetClientCaCertificateHostnamesResponse {
	s.Body = v
	return s
}

func (s *SetClientCaCertificateHostnamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
