// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorReportComponentSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorReportComponentSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorReportComponentSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorReportComponentSummaryResponseBody) *GetDoctorReportComponentSummaryResponse
	GetBody() *GetDoctorReportComponentSummaryResponseBody
}

type GetDoctorReportComponentSummaryResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorReportComponentSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorReportComponentSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorReportComponentSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorReportComponentSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorReportComponentSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorReportComponentSummaryResponse) GetBody() *GetDoctorReportComponentSummaryResponseBody {
	return s.Body
}

func (s *GetDoctorReportComponentSummaryResponse) SetHeaders(v map[string]*string) *GetDoctorReportComponentSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorReportComponentSummaryResponse) SetStatusCode(v int32) *GetDoctorReportComponentSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorReportComponentSummaryResponse) SetBody(v *GetDoctorReportComponentSummaryResponseBody) *GetDoctorReportComponentSummaryResponse {
	s.Body = v
	return s
}

func (s *GetDoctorReportComponentSummaryResponse) Validate() error {
	return dara.Validate(s)
}
