// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalSkillGroupReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHistoricalSkillGroupReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHistoricalSkillGroupReportResponse
	GetStatusCode() *int32
	SetBody(v *ListHistoricalSkillGroupReportResponseBody) *ListHistoricalSkillGroupReportResponse
	GetBody() *ListHistoricalSkillGroupReportResponseBody
}

type ListHistoricalSkillGroupReportResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHistoricalSkillGroupReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHistoricalSkillGroupReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalSkillGroupReportResponse) GoString() string {
	return s.String()
}

func (s *ListHistoricalSkillGroupReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHistoricalSkillGroupReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHistoricalSkillGroupReportResponse) GetBody() *ListHistoricalSkillGroupReportResponseBody {
	return s.Body
}

func (s *ListHistoricalSkillGroupReportResponse) SetHeaders(v map[string]*string) *ListHistoricalSkillGroupReportResponse {
	s.Headers = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponse) SetStatusCode(v int32) *ListHistoricalSkillGroupReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHistoricalSkillGroupReportResponse) SetBody(v *ListHistoricalSkillGroupReportResponseBody) *ListHistoricalSkillGroupReportResponse {
	s.Body = v
	return s
}

func (s *ListHistoricalSkillGroupReportResponse) Validate() error {
	return dara.Validate(s)
}
