// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignQualityRuleOfAllRuleScopeSchedulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignQualityRuleOfAllRuleScopeSchedulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignQualityRuleOfAllRuleScopeSchedulesResponse
	GetStatusCode() *int32
	SetBody(v *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) *AssignQualityRuleOfAllRuleScopeSchedulesResponse
	GetBody() *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody
}

type AssignQualityRuleOfAllRuleScopeSchedulesResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignQualityRuleOfAllRuleScopeSchedulesResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignQualityRuleOfAllRuleScopeSchedulesResponse) GoString() string {
	return s.String()
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponse) GetBody() *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody {
	return s.Body
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponse) SetHeaders(v map[string]*string) *AssignQualityRuleOfAllRuleScopeSchedulesResponse {
	s.Headers = v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponse) SetStatusCode(v int32) *AssignQualityRuleOfAllRuleScopeSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponse) SetBody(v *AssignQualityRuleOfAllRuleScopeSchedulesResponseBody) *AssignQualityRuleOfAllRuleScopeSchedulesResponse {
	s.Body = v
	return s
}

func (s *AssignQualityRuleOfAllRuleScopeSchedulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
