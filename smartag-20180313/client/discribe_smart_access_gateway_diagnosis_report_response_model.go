// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscribeSmartAccessGatewayDiagnosisReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiscribeSmartAccessGatewayDiagnosisReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponse
	GetStatusCode() *int32
	SetBody(v *DiscribeSmartAccessGatewayDiagnosisReportResponseBody) *DiscribeSmartAccessGatewayDiagnosisReportResponse
	GetBody() *DiscribeSmartAccessGatewayDiagnosisReportResponseBody
}

type DiscribeSmartAccessGatewayDiagnosisReportResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiscribeSmartAccessGatewayDiagnosisReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponse) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportResponse) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponse) GetBody() *DiscribeSmartAccessGatewayDiagnosisReportResponseBody {
	return s.Body
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponse) SetHeaders(v map[string]*string) *DiscribeSmartAccessGatewayDiagnosisReportResponse {
	s.Headers = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponse) SetStatusCode(v int32) *DiscribeSmartAccessGatewayDiagnosisReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponse) SetBody(v *DiscribeSmartAccessGatewayDiagnosisReportResponseBody) *DiscribeSmartAccessGatewayDiagnosisReportResponse {
	s.Body = v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
