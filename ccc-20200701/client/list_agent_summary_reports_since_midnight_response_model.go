// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSummaryReportsSinceMidnightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAgentSummaryReportsSinceMidnightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAgentSummaryReportsSinceMidnightResponse
	GetStatusCode() *int32
	SetBody(v *ListAgentSummaryReportsSinceMidnightResponseBody) *ListAgentSummaryReportsSinceMidnightResponse
	GetBody() *ListAgentSummaryReportsSinceMidnightResponseBody
}

type ListAgentSummaryReportsSinceMidnightResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAgentSummaryReportsSinceMidnightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAgentSummaryReportsSinceMidnightResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSummaryReportsSinceMidnightResponse) GoString() string {
	return s.String()
}

func (s *ListAgentSummaryReportsSinceMidnightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAgentSummaryReportsSinceMidnightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAgentSummaryReportsSinceMidnightResponse) GetBody() *ListAgentSummaryReportsSinceMidnightResponseBody {
	return s.Body
}

func (s *ListAgentSummaryReportsSinceMidnightResponse) SetHeaders(v map[string]*string) *ListAgentSummaryReportsSinceMidnightResponse {
	s.Headers = v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponse) SetStatusCode(v int32) *ListAgentSummaryReportsSinceMidnightResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponse) SetBody(v *ListAgentSummaryReportsSinceMidnightResponseBody) *ListAgentSummaryReportsSinceMidnightResponse {
	s.Body = v
	return s
}

func (s *ListAgentSummaryReportsSinceMidnightResponse) Validate() error {
	return dara.Validate(s)
}
