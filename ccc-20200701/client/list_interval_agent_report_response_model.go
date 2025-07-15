// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalAgentReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntervalAgentReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntervalAgentReportResponse
	GetStatusCode() *int32
	SetBody(v *ListIntervalAgentReportResponseBody) *ListIntervalAgentReportResponse
	GetBody() *ListIntervalAgentReportResponseBody
}

type ListIntervalAgentReportResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntervalAgentReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntervalAgentReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentReportResponse) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntervalAgentReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntervalAgentReportResponse) GetBody() *ListIntervalAgentReportResponseBody {
	return s.Body
}

func (s *ListIntervalAgentReportResponse) SetHeaders(v map[string]*string) *ListIntervalAgentReportResponse {
	s.Headers = v
	return s
}

func (s *ListIntervalAgentReportResponse) SetStatusCode(v int32) *ListIntervalAgentReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntervalAgentReportResponse) SetBody(v *ListIntervalAgentReportResponseBody) *ListIntervalAgentReportResponse {
	s.Body = v
	return s
}

func (s *ListIntervalAgentReportResponse) Validate() error {
	return dara.Validate(s)
}
