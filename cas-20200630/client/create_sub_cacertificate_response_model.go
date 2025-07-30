// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCACertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSubCACertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSubCACertificateResponse
	GetStatusCode() *int32
	SetBody(v *CreateSubCACertificateResponseBody) *CreateSubCACertificateResponse
	GetBody() *CreateSubCACertificateResponseBody
}

type CreateSubCACertificateResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSubCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSubCACertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCACertificateResponse) GoString() string {
	return s.String()
}

func (s *CreateSubCACertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSubCACertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSubCACertificateResponse) GetBody() *CreateSubCACertificateResponseBody {
	return s.Body
}

func (s *CreateSubCACertificateResponse) SetHeaders(v map[string]*string) *CreateSubCACertificateResponse {
	s.Headers = v
	return s
}

func (s *CreateSubCACertificateResponse) SetStatusCode(v int32) *CreateSubCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubCACertificateResponse) SetBody(v *CreateSubCACertificateResponseBody) *CreateSubCACertificateResponse {
	s.Body = v
	return s
}

func (s *CreateSubCACertificateResponse) Validate() error {
	return dara.Validate(s)
}
