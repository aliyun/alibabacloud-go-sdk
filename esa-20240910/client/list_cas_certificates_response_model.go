// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCasCertificatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCasCertificatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCasCertificatesResponse
	GetStatusCode() *int32
	SetBody(v *ListCasCertificatesResponseBody) *ListCasCertificatesResponse
	GetBody() *ListCasCertificatesResponseBody
}

type ListCasCertificatesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCasCertificatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCasCertificatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCasCertificatesResponse) GoString() string {
	return s.String()
}

func (s *ListCasCertificatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCasCertificatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCasCertificatesResponse) GetBody() *ListCasCertificatesResponseBody {
	return s.Body
}

func (s *ListCasCertificatesResponse) SetHeaders(v map[string]*string) *ListCasCertificatesResponse {
	s.Headers = v
	return s
}

func (s *ListCasCertificatesResponse) SetStatusCode(v int32) *ListCasCertificatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCasCertificatesResponse) SetBody(v *ListCasCertificatesResponseBody) *ListCasCertificatesResponse {
	s.Body = v
	return s
}

func (s *ListCasCertificatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
