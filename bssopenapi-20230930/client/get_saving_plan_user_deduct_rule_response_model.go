// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanUserDeductRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSavingPlanUserDeductRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSavingPlanUserDeductRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetSavingPlanUserDeductRuleResponseBody) *GetSavingPlanUserDeductRuleResponse
	GetBody() *GetSavingPlanUserDeductRuleResponseBody
}

type GetSavingPlanUserDeductRuleResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavingPlanUserDeductRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavingPlanUserDeductRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanUserDeductRuleResponse) GoString() string {
	return s.String()
}

func (s *GetSavingPlanUserDeductRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSavingPlanUserDeductRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSavingPlanUserDeductRuleResponse) GetBody() *GetSavingPlanUserDeductRuleResponseBody {
	return s.Body
}

func (s *GetSavingPlanUserDeductRuleResponse) SetHeaders(v map[string]*string) *GetSavingPlanUserDeductRuleResponse {
	s.Headers = v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponse) SetStatusCode(v int32) *GetSavingPlanUserDeductRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponse) SetBody(v *GetSavingPlanUserDeductRuleResponseBody) *GetSavingPlanUserDeductRuleResponse {
	s.Body = v
	return s
}

func (s *GetSavingPlanUserDeductRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
