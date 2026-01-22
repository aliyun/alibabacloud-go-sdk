// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCredentialReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateCredentialReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateCredentialReportResponse
	GetStatusCode() *int32
	SetBody(v *GenerateCredentialReportResponseBody) *GenerateCredentialReportResponse
	GetBody() *GenerateCredentialReportResponseBody
}

type GenerateCredentialReportResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCredentialReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCredentialReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateCredentialReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateCredentialReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateCredentialReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateCredentialReportResponse) GetBody() *GenerateCredentialReportResponseBody {
	return s.Body
}

func (s *GenerateCredentialReportResponse) SetHeaders(v map[string]*string) *GenerateCredentialReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateCredentialReportResponse) SetStatusCode(v int32) *GenerateCredentialReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCredentialReportResponse) SetBody(v *GenerateCredentialReportResponseBody) *GenerateCredentialReportResponse {
	s.Body = v
	return s
}

func (s *GenerateCredentialReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
