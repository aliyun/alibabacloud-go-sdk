// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntervalSkillGroupReportResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntervalSkillGroupReportResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntervalSkillGroupReportResponse
	GetStatusCode() *int32
	SetBody(v *ListIntervalSkillGroupReportResponseBody) *ListIntervalSkillGroupReportResponse
	GetBody() *ListIntervalSkillGroupReportResponseBody
}

type ListIntervalSkillGroupReportResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntervalSkillGroupReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntervalSkillGroupReportResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntervalSkillGroupReportResponse) GoString() string {
	return s.String()
}

func (s *ListIntervalSkillGroupReportResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntervalSkillGroupReportResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntervalSkillGroupReportResponse) GetBody() *ListIntervalSkillGroupReportResponseBody {
	return s.Body
}

func (s *ListIntervalSkillGroupReportResponse) SetHeaders(v map[string]*string) *ListIntervalSkillGroupReportResponse {
	s.Headers = v
	return s
}

func (s *ListIntervalSkillGroupReportResponse) SetStatusCode(v int32) *ListIntervalSkillGroupReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntervalSkillGroupReportResponse) SetBody(v *ListIntervalSkillGroupReportResponseBody) *ListIntervalSkillGroupReportResponse {
	s.Body = v
	return s
}

func (s *ListIntervalSkillGroupReportResponse) Validate() error {
	return dara.Validate(s)
}
