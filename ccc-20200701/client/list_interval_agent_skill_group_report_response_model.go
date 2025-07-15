// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalAgentSkillGroupReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntervalAgentSkillGroupReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntervalAgentSkillGroupReportResponse
	GetStatusCode() *int32
	SetBody(v *ListIntervalAgentSkillGroupReportResponseBody) *ListIntervalAgentSkillGroupReportResponse
	GetBody() *ListIntervalAgentSkillGroupReportResponseBody
}

type ListIntervalAgentSkillGroupReportResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntervalAgentSkillGroupReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntervalAgentSkillGroupReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalAgentSkillGroupReportResponse) GoString() string {
	return s.String()
}

func (s *ListIntervalAgentSkillGroupReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntervalAgentSkillGroupReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntervalAgentSkillGroupReportResponse) GetBody() *ListIntervalAgentSkillGroupReportResponseBody {
	return s.Body
}

func (s *ListIntervalAgentSkillGroupReportResponse) SetHeaders(v map[string]*string) *ListIntervalAgentSkillGroupReportResponse {
	s.Headers = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponse) SetStatusCode(v int32) *ListIntervalAgentSkillGroupReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponse) SetBody(v *ListIntervalAgentSkillGroupReportResponseBody) *ListIntervalAgentSkillGroupReportResponse {
	s.Body = v
	return s
}

func (s *ListIntervalAgentSkillGroupReportResponse) Validate() error {
	return dara.Validate(s)
}
