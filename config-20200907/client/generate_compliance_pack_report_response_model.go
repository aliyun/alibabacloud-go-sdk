// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCompliancePackReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateCompliancePackReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateCompliancePackReportResponse
	GetStatusCode() *int32
	SetBody(v *GenerateCompliancePackReportResponseBody) *GenerateCompliancePackReportResponse
	GetBody() *GenerateCompliancePackReportResponseBody
}

type GenerateCompliancePackReportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCompliancePackReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCompliancePackReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateCompliancePackReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateCompliancePackReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateCompliancePackReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateCompliancePackReportResponse) GetBody() *GenerateCompliancePackReportResponseBody {
	return s.Body
}

func (s *GenerateCompliancePackReportResponse) SetHeaders(v map[string]*string) *GenerateCompliancePackReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateCompliancePackReportResponse) SetStatusCode(v int32) *GenerateCompliancePackReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCompliancePackReportResponse) SetBody(v *GenerateCompliancePackReportResponseBody) *GenerateCompliancePackReportResponse {
	s.Body = v
	return s
}

func (s *GenerateCompliancePackReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
