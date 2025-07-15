// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHistoricalAgentSkillGroupReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHistoricalAgentSkillGroupReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHistoricalAgentSkillGroupReportResponse
	GetStatusCode() *int32
	SetBody(v *ListHistoricalAgentSkillGroupReportResponseBody) *ListHistoricalAgentSkillGroupReportResponse
	GetBody() *ListHistoricalAgentSkillGroupReportResponseBody
}

type ListHistoricalAgentSkillGroupReportResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHistoricalAgentSkillGroupReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHistoricalAgentSkillGroupReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHistoricalAgentSkillGroupReportResponse) GoString() string {
	return s.String()
}

func (s *ListHistoricalAgentSkillGroupReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHistoricalAgentSkillGroupReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHistoricalAgentSkillGroupReportResponse) GetBody() *ListHistoricalAgentSkillGroupReportResponseBody {
	return s.Body
}

func (s *ListHistoricalAgentSkillGroupReportResponse) SetHeaders(v map[string]*string) *ListHistoricalAgentSkillGroupReportResponse {
	s.Headers = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponse) SetStatusCode(v int32) *ListHistoricalAgentSkillGroupReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponse) SetBody(v *ListHistoricalAgentSkillGroupReportResponseBody) *ListHistoricalAgentSkillGroupReportResponse {
	s.Body = v
	return s
}

func (s *ListHistoricalAgentSkillGroupReportResponse) Validate() error {
	return dara.Validate(s)
}
