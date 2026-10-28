// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartQuotaUtilizationReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartQuotaUtilizationReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartQuotaUtilizationReportResponse
	GetStatusCode() *int32
	SetBody(v *StartQuotaUtilizationReportResponseBody) *StartQuotaUtilizationReportResponse
	GetBody() *StartQuotaUtilizationReportResponseBody
}

type StartQuotaUtilizationReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartQuotaUtilizationReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartQuotaUtilizationReportResponse) String() string {
	return dara.Prettify(s)
}

func (s StartQuotaUtilizationReportResponse) GoString() string {
	return s.String()
}

func (s *StartQuotaUtilizationReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartQuotaUtilizationReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartQuotaUtilizationReportResponse) GetBody() *StartQuotaUtilizationReportResponseBody {
	return s.Body
}

func (s *StartQuotaUtilizationReportResponse) SetHeaders(v map[string]*string) *StartQuotaUtilizationReportResponse {
	s.Headers = v
	return s
}

func (s *StartQuotaUtilizationReportResponse) SetStatusCode(v int32) *StartQuotaUtilizationReportResponse {
	s.StatusCode = &v
	return s
}

func (s *StartQuotaUtilizationReportResponse) SetBody(v *StartQuotaUtilizationReportResponseBody) *StartQuotaUtilizationReportResponse {
	s.Body = v
	return s
}

func (s *StartQuotaUtilizationReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
