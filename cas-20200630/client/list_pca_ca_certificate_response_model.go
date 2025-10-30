// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPcaCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPcaCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPcaCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *ListPcaCaCertificateResponseBody) *ListPcaCaCertificateResponse
	GetBody() *ListPcaCaCertificateResponseBody
}

type ListPcaCaCertificateResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPcaCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPcaCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPcaCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *ListPcaCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPcaCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPcaCaCertificateResponse) GetBody() *ListPcaCaCertificateResponseBody {
	return s.Body
}

func (s *ListPcaCaCertificateResponse) SetHeaders(v map[string]*string) *ListPcaCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *ListPcaCaCertificateResponse) SetStatusCode(v int32) *ListPcaCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPcaCaCertificateResponse) SetBody(v *ListPcaCaCertificateResponseBody) *ListPcaCaCertificateResponse {
	s.Body = v
	return s
}

func (s *ListPcaCaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
