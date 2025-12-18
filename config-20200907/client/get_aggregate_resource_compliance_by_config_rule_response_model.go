// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceByConfigRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateResourceComplianceByConfigRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateResourceComplianceByConfigRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateResourceComplianceByConfigRuleResponseBody) *GetAggregateResourceComplianceByConfigRuleResponse
	GetBody() *GetAggregateResourceComplianceByConfigRuleResponseBody
}

type GetAggregateResourceComplianceByConfigRuleResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateResourceComplianceByConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateResourceComplianceByConfigRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceByConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) GetBody() *GetAggregateResourceComplianceByConfigRuleResponseBody {
	return s.Body
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceByConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceByConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) SetBody(v *GetAggregateResourceComplianceByConfigRuleResponseBody) *GetAggregateResourceComplianceByConfigRuleResponse {
	s.Body = v
	return s
}

func (s *GetAggregateResourceComplianceByConfigRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
