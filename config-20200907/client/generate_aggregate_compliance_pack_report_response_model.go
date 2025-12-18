// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAggregateCompliancePackReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAggregateCompliancePackReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAggregateCompliancePackReportResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAggregateCompliancePackReportResponseBody) *GenerateAggregateCompliancePackReportResponse
	GetBody() *GenerateAggregateCompliancePackReportResponseBody
}

type GenerateAggregateCompliancePackReportResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAggregateCompliancePackReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAggregateCompliancePackReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAggregateCompliancePackReportResponse) GoString() string {
	return s.String()
}

func (s *GenerateAggregateCompliancePackReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAggregateCompliancePackReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAggregateCompliancePackReportResponse) GetBody() *GenerateAggregateCompliancePackReportResponseBody {
	return s.Body
}

func (s *GenerateAggregateCompliancePackReportResponse) SetHeaders(v map[string]*string) *GenerateAggregateCompliancePackReportResponse {
	s.Headers = v
	return s
}

func (s *GenerateAggregateCompliancePackReportResponse) SetStatusCode(v int32) *GenerateAggregateCompliancePackReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAggregateCompliancePackReportResponse) SetBody(v *GenerateAggregateCompliancePackReportResponseBody) *GenerateAggregateCompliancePackReportResponse {
	s.Body = v
	return s
}

func (s *GenerateAggregateCompliancePackReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
