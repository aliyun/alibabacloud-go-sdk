// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUserFbAcIssueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportUserFbAcIssueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportUserFbAcIssueResponse
	GetStatusCode() *int32
	SetBody(v *ReportUserFbAcIssueResponseBody) *ReportUserFbAcIssueResponse
	GetBody() *ReportUserFbAcIssueResponseBody
}

type ReportUserFbAcIssueResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportUserFbAcIssueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportUserFbAcIssueResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbAcIssueResponse) GoString() string {
	return s.String()
}

func (s *ReportUserFbAcIssueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportUserFbAcIssueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportUserFbAcIssueResponse) GetBody() *ReportUserFbAcIssueResponseBody {
	return s.Body
}

func (s *ReportUserFbAcIssueResponse) SetHeaders(v map[string]*string) *ReportUserFbAcIssueResponse {
	s.Headers = v
	return s
}

func (s *ReportUserFbAcIssueResponse) SetStatusCode(v int32) *ReportUserFbAcIssueResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportUserFbAcIssueResponse) SetBody(v *ReportUserFbAcIssueResponseBody) *ReportUserFbAcIssueResponse {
	s.Body = v
	return s
}

func (s *ReportUserFbAcIssueResponse) Validate() error {
	return dara.Validate(s)
}
