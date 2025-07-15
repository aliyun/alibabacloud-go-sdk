// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupSummaryReportsSinceMidnightResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillGroupSummaryReportsSinceMidnightResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillGroupSummaryReportsSinceMidnightResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillGroupSummaryReportsSinceMidnightResponseBody) *ListSkillGroupSummaryReportsSinceMidnightResponse
	GetBody() *ListSkillGroupSummaryReportsSinceMidnightResponseBody
}

type ListSkillGroupSummaryReportsSinceMidnightResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillGroupSummaryReportsSinceMidnightResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupSummaryReportsSinceMidnightResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponse) GetBody() *ListSkillGroupSummaryReportsSinceMidnightResponseBody {
	return s.Body
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponse) SetHeaders(v map[string]*string) *ListSkillGroupSummaryReportsSinceMidnightResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponse) SetStatusCode(v int32) *ListSkillGroupSummaryReportsSinceMidnightResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponse) SetBody(v *ListSkillGroupSummaryReportsSinceMidnightResponseBody) *ListSkillGroupSummaryReportsSinceMidnightResponse {
	s.Body = v
	return s
}

func (s *ListSkillGroupSummaryReportsSinceMidnightResponse) Validate() error {
	return dara.Validate(s)
}
