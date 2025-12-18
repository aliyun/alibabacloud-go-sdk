// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompliancePackReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCompliancePackReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCompliancePackReportResponse
	GetStatusCode() *int32
	SetBody(v *GetCompliancePackReportResponseBody) *GetCompliancePackReportResponse
	GetBody() *GetCompliancePackReportResponseBody
}

type GetCompliancePackReportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCompliancePackReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCompliancePackReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCompliancePackReportResponse) GoString() string {
	return s.String()
}

func (s *GetCompliancePackReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCompliancePackReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCompliancePackReportResponse) GetBody() *GetCompliancePackReportResponseBody {
	return s.Body
}

func (s *GetCompliancePackReportResponse) SetHeaders(v map[string]*string) *GetCompliancePackReportResponse {
	s.Headers = v
	return s
}

func (s *GetCompliancePackReportResponse) SetStatusCode(v int32) *GetCompliancePackReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCompliancePackReportResponse) SetBody(v *GetCompliancePackReportResponseBody) *GetCompliancePackReportResponse {
	s.Body = v
	return s
}

func (s *GetCompliancePackReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
