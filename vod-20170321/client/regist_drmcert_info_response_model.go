// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegistDRMCertInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegistDRMCertInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegistDRMCertInfoResponse
	GetStatusCode() *int32
	SetBody(v *RegistDRMCertInfoResponseBody) *RegistDRMCertInfoResponse
	GetBody() *RegistDRMCertInfoResponseBody
}

type RegistDRMCertInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegistDRMCertInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegistDRMCertInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s RegistDRMCertInfoResponse) GoString() string {
	return s.String()
}

func (s *RegistDRMCertInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegistDRMCertInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegistDRMCertInfoResponse) GetBody() *RegistDRMCertInfoResponseBody {
	return s.Body
}

func (s *RegistDRMCertInfoResponse) SetHeaders(v map[string]*string) *RegistDRMCertInfoResponse {
	s.Headers = v
	return s
}

func (s *RegistDRMCertInfoResponse) SetStatusCode(v int32) *RegistDRMCertInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *RegistDRMCertInfoResponse) SetBody(v *RegistDRMCertInfoResponseBody) *RegistDRMCertInfoResponse {
	s.Body = v
	return s
}

func (s *RegistDRMCertInfoResponse) Validate() error {
	return dara.Validate(s)
}
