// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalAgentReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHistoricalAgentReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHistoricalAgentReportResponse
	GetStatusCode() *int32
	SetBody(v *ListHistoricalAgentReportResponseBody) *ListHistoricalAgentReportResponse
	GetBody() *ListHistoricalAgentReportResponseBody
}

type ListHistoricalAgentReportResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHistoricalAgentReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHistoricalAgentReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentReportResponse) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHistoricalAgentReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHistoricalAgentReportResponse) GetBody() *ListHistoricalAgentReportResponseBody {
	return s.Body
}

func (s *ListHistoricalAgentReportResponse) SetHeaders(v map[string]*string) *ListHistoricalAgentReportResponse {
	s.Headers = v
	return s
}

func (s *ListHistoricalAgentReportResponse) SetStatusCode(v int32) *ListHistoricalAgentReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHistoricalAgentReportResponse) SetBody(v *ListHistoricalAgentReportResponseBody) *ListHistoricalAgentReportResponse {
	s.Body = v
	return s
}

func (s *ListHistoricalAgentReportResponse) Validate() error {
	return dara.Validate(s)
}
