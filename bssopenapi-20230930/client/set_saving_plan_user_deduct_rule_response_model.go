// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSavingPlanUserDeductRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSavingPlanUserDeductRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSavingPlanUserDeductRuleResponse
	GetStatusCode() *int32
	SetBody(v *SetSavingPlanUserDeductRuleResponseBody) *SetSavingPlanUserDeductRuleResponse
	GetBody() *SetSavingPlanUserDeductRuleResponseBody
}

type SetSavingPlanUserDeductRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSavingPlanUserDeductRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSavingPlanUserDeductRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSavingPlanUserDeductRuleResponse) GoString() string {
	return s.String()
}

func (s *SetSavingPlanUserDeductRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSavingPlanUserDeductRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSavingPlanUserDeductRuleResponse) GetBody() *SetSavingPlanUserDeductRuleResponseBody {
	return s.Body
}

func (s *SetSavingPlanUserDeductRuleResponse) SetHeaders(v map[string]*string) *SetSavingPlanUserDeductRuleResponse {
	s.Headers = v
	return s
}

func (s *SetSavingPlanUserDeductRuleResponse) SetStatusCode(v int32) *SetSavingPlanUserDeductRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSavingPlanUserDeductRuleResponse) SetBody(v *SetSavingPlanUserDeductRuleResponseBody) *SetSavingPlanUserDeductRuleResponse {
	s.Body = v
	return s
}

func (s *SetSavingPlanUserDeductRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
