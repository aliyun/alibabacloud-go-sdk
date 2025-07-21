// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportUserFbIssueResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportUserFbIssueResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportUserFbIssueResponse
	GetStatusCode() *int32
	SetBody(v *ReportUserFbIssueResponseBody) *ReportUserFbIssueResponse
	GetBody() *ReportUserFbIssueResponseBody
}

type ReportUserFbIssueResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportUserFbIssueResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportUserFbIssueResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportUserFbIssueResponse) GoString() string {
	return s.String()
}

func (s *ReportUserFbIssueResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportUserFbIssueResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportUserFbIssueResponse) GetBody() *ReportUserFbIssueResponseBody {
	return s.Body
}

func (s *ReportUserFbIssueResponse) SetHeaders(v map[string]*string) *ReportUserFbIssueResponse {
	s.Headers = v
	return s
}

func (s *ReportUserFbIssueResponse) SetStatusCode(v int32) *ReportUserFbIssueResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportUserFbIssueResponse) SetBody(v *ReportUserFbIssueResponseBody) *ReportUserFbIssueResponse {
	s.Body = v
	return s
}

func (s *ReportUserFbIssueResponse) Validate() error {
	return dara.Validate(s)
}
