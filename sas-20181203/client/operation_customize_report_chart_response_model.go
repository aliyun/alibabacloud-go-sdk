// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperationCustomizeReportChartResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperationCustomizeReportChartResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperationCustomizeReportChartResponse
	GetStatusCode() *int32
	SetBody(v *OperationCustomizeReportChartResponseBody) *OperationCustomizeReportChartResponse
	GetBody() *OperationCustomizeReportChartResponseBody
}

type OperationCustomizeReportChartResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperationCustomizeReportChartResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperationCustomizeReportChartResponse) String() string {
	return dara.Prettify(s)
}

func (s OperationCustomizeReportChartResponse) GoString() string {
	return s.String()
}

func (s *OperationCustomizeReportChartResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperationCustomizeReportChartResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperationCustomizeReportChartResponse) GetBody() *OperationCustomizeReportChartResponseBody {
	return s.Body
}

func (s *OperationCustomizeReportChartResponse) SetHeaders(v map[string]*string) *OperationCustomizeReportChartResponse {
	s.Headers = v
	return s
}

func (s *OperationCustomizeReportChartResponse) SetStatusCode(v int32) *OperationCustomizeReportChartResponse {
	s.StatusCode = &v
	return s
}

func (s *OperationCustomizeReportChartResponse) SetBody(v *OperationCustomizeReportChartResponseBody) *OperationCustomizeReportChartResponse {
	s.Body = v
	return s
}

func (s *OperationCustomizeReportChartResponse) Validate() error {
	return dara.Validate(s)
}
