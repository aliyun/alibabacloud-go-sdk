// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportSampleConsistencyJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportSampleConsistencyJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportSampleConsistencyJobResponse
	GetStatusCode() *int32
	SetBody(v *ReportSampleConsistencyJobResponseBody) *ReportSampleConsistencyJobResponse
	GetBody() *ReportSampleConsistencyJobResponseBody
}

type ReportSampleConsistencyJobResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportSampleConsistencyJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportSampleConsistencyJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportSampleConsistencyJobResponse) GoString() string {
	return s.String()
}

func (s *ReportSampleConsistencyJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportSampleConsistencyJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportSampleConsistencyJobResponse) GetBody() *ReportSampleConsistencyJobResponseBody {
	return s.Body
}

func (s *ReportSampleConsistencyJobResponse) SetHeaders(v map[string]*string) *ReportSampleConsistencyJobResponse {
	s.Headers = v
	return s
}

func (s *ReportSampleConsistencyJobResponse) SetStatusCode(v int32) *ReportSampleConsistencyJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportSampleConsistencyJobResponse) SetBody(v *ReportSampleConsistencyJobResponseBody) *ReportSampleConsistencyJobResponse {
	s.Body = v
	return s
}

func (s *ReportSampleConsistencyJobResponse) Validate() error {
	return dara.Validate(s)
}
