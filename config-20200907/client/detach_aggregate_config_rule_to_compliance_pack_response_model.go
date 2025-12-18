// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAggregateConfigRuleToCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachAggregateConfigRuleToCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachAggregateConfigRuleToCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *DetachAggregateConfigRuleToCompliancePackResponseBody) *DetachAggregateConfigRuleToCompliancePackResponse
	GetBody() *DetachAggregateConfigRuleToCompliancePackResponseBody
}

type DetachAggregateConfigRuleToCompliancePackResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachAggregateConfigRuleToCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachAggregateConfigRuleToCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachAggregateConfigRuleToCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) GetBody() *DetachAggregateConfigRuleToCompliancePackResponseBody {
	return s.Body
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) SetHeaders(v map[string]*string) *DetachAggregateConfigRuleToCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) SetStatusCode(v int32) *DetachAggregateConfigRuleToCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) SetBody(v *DetachAggregateConfigRuleToCompliancePackResponseBody) *DetachAggregateConfigRuleToCompliancePackResponse {
	s.Body = v
	return s
}

func (s *DetachAggregateConfigRuleToCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
