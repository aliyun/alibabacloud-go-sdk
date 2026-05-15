// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillGroupAndAgentStatusSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillGroupAndAgentStatusSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillGroupAndAgentStatusSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillGroupAndAgentStatusSummaryResponseBody) *GetSkillGroupAndAgentStatusSummaryResponse
	GetBody() *GetSkillGroupAndAgentStatusSummaryResponseBody
}

type GetSkillGroupAndAgentStatusSummaryResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillGroupAndAgentStatusSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillGroupAndAgentStatusSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillGroupAndAgentStatusSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetSkillGroupAndAgentStatusSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillGroupAndAgentStatusSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillGroupAndAgentStatusSummaryResponse) GetBody() *GetSkillGroupAndAgentStatusSummaryResponseBody {
	return s.Body
}

func (s *GetSkillGroupAndAgentStatusSummaryResponse) SetHeaders(v map[string]*string) *GetSkillGroupAndAgentStatusSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponse) SetStatusCode(v int32) *GetSkillGroupAndAgentStatusSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponse) SetBody(v *GetSkillGroupAndAgentStatusSummaryResponseBody) *GetSkillGroupAndAgentStatusSummaryResponse {
	s.Body = v
	return s
}

func (s *GetSkillGroupAndAgentStatusSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
