// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCertificateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCertificateDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetCertificateDetailResponseBody) *GetCertificateDetailResponse
	GetBody() *GetCertificateDetailResponseBody
}

type GetCertificateDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCertificateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetCertificateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCertificateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCertificateDetailResponse) GetBody() *GetCertificateDetailResponseBody {
	return s.Body
}

func (s *GetCertificateDetailResponse) SetHeaders(v map[string]*string) *GetCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetCertificateDetailResponse) SetStatusCode(v int32) *GetCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCertificateDetailResponse) SetBody(v *GetCertificateDetailResponseBody) *GetCertificateDetailResponse {
	s.Body = v
	return s
}

func (s *GetCertificateDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
