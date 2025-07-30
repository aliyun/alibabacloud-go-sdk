// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportEdsAgentInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportEdsAgentInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportEdsAgentInfoResponse
	GetStatusCode() *int32
	SetBody(v *ReportEdsAgentInfoResponseBody) *ReportEdsAgentInfoResponse
	GetBody() *ReportEdsAgentInfoResponseBody
}

type ReportEdsAgentInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportEdsAgentInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportEdsAgentInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportEdsAgentInfoResponse) GoString() string {
	return s.String()
}

func (s *ReportEdsAgentInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportEdsAgentInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportEdsAgentInfoResponse) GetBody() *ReportEdsAgentInfoResponseBody {
	return s.Body
}

func (s *ReportEdsAgentInfoResponse) SetHeaders(v map[string]*string) *ReportEdsAgentInfoResponse {
	s.Headers = v
	return s
}

func (s *ReportEdsAgentInfoResponse) SetStatusCode(v int32) *ReportEdsAgentInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportEdsAgentInfoResponse) SetBody(v *ReportEdsAgentInfoResponseBody) *ReportEdsAgentInfoResponse {
	s.Body = v
	return s
}

func (s *ReportEdsAgentInfoResponse) Validate() error {
	return dara.Validate(s)
}
