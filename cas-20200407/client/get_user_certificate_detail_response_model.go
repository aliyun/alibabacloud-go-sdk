// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCertificateDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserCertificateDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserCertificateDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetUserCertificateDetailResponseBody) *GetUserCertificateDetailResponse
	GetBody() *GetUserCertificateDetailResponseBody
}

type GetUserCertificateDetailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserCertificateDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserCertificateDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserCertificateDetailResponse) GoString() string {
	return s.String()
}

func (s *GetUserCertificateDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserCertificateDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserCertificateDetailResponse) GetBody() *GetUserCertificateDetailResponseBody {
	return s.Body
}

func (s *GetUserCertificateDetailResponse) SetHeaders(v map[string]*string) *GetUserCertificateDetailResponse {
	s.Headers = v
	return s
}

func (s *GetUserCertificateDetailResponse) SetStatusCode(v int32) *GetUserCertificateDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserCertificateDetailResponse) SetBody(v *GetUserCertificateDetailResponseBody) *GetUserCertificateDetailResponse {
	s.Body = v
	return s
}

func (s *GetUserCertificateDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
