// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotaUtilizationReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQuotaUtilizationReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQuotaUtilizationReportResponse
	GetStatusCode() *int32
	SetBody(v *ListQuotaUtilizationReportResponseBody) *ListQuotaUtilizationReportResponse
	GetBody() *ListQuotaUtilizationReportResponseBody
}

type ListQuotaUtilizationReportResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaUtilizationReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQuotaUtilizationReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQuotaUtilizationReportResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaUtilizationReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQuotaUtilizationReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQuotaUtilizationReportResponse) GetBody() *ListQuotaUtilizationReportResponseBody {
	return s.Body
}

func (s *ListQuotaUtilizationReportResponse) SetHeaders(v map[string]*string) *ListQuotaUtilizationReportResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaUtilizationReportResponse) SetStatusCode(v int32) *ListQuotaUtilizationReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaUtilizationReportResponse) SetBody(v *ListQuotaUtilizationReportResponseBody) *ListQuotaUtilizationReportResponse {
	s.Body = v
	return s
}

func (s *ListQuotaUtilizationReportResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
