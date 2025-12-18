// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateCompliancePackReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateCompliancePackReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateCompliancePackReportResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateCompliancePackReportResponseBody) *GetAggregateCompliancePackReportResponse
	GetBody() *GetAggregateCompliancePackReportResponseBody
}

type GetAggregateCompliancePackReportResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateCompliancePackReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateCompliancePackReportResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateCompliancePackReportResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateCompliancePackReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateCompliancePackReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateCompliancePackReportResponse) GetBody() *GetAggregateCompliancePackReportResponseBody {
	return s.Body
}

func (s *GetAggregateCompliancePackReportResponse) SetHeaders(v map[string]*string) *GetAggregateCompliancePackReportResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateCompliancePackReportResponse) SetStatusCode(v int32) *GetAggregateCompliancePackReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateCompliancePackReportResponse) SetBody(v *GetAggregateCompliancePackReportResponseBody) *GetAggregateCompliancePackReportResponse {
	s.Body = v
	return s
}

func (s *GetAggregateCompliancePackReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
