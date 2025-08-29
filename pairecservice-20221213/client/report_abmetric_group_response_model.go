// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportABMetricGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportABMetricGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportABMetricGroupResponse
	GetStatusCode() *int32
	SetBody(v *ReportABMetricGroupResponseBody) *ReportABMetricGroupResponse
	GetBody() *ReportABMetricGroupResponseBody
}

type ReportABMetricGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportABMetricGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportABMetricGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportABMetricGroupResponse) GoString() string {
	return s.String()
}

func (s *ReportABMetricGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportABMetricGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportABMetricGroupResponse) GetBody() *ReportABMetricGroupResponseBody {
	return s.Body
}

func (s *ReportABMetricGroupResponse) SetHeaders(v map[string]*string) *ReportABMetricGroupResponse {
	s.Headers = v
	return s
}

func (s *ReportABMetricGroupResponse) SetStatusCode(v int32) *ReportABMetricGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportABMetricGroupResponse) SetBody(v *ReportABMetricGroupResponseBody) *ReportABMetricGroupResponse {
	s.Body = v
	return s
}

func (s *ReportABMetricGroupResponse) Validate() error {
	return dara.Validate(s)
}
