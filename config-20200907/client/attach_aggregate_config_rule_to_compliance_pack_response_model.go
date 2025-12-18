// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAggregateConfigRuleToCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachAggregateConfigRuleToCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachAggregateConfigRuleToCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *AttachAggregateConfigRuleToCompliancePackResponseBody) *AttachAggregateConfigRuleToCompliancePackResponse
	GetBody() *AttachAggregateConfigRuleToCompliancePackResponseBody
}

type AttachAggregateConfigRuleToCompliancePackResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachAggregateConfigRuleToCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachAggregateConfigRuleToCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachAggregateConfigRuleToCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) GetBody() *AttachAggregateConfigRuleToCompliancePackResponseBody {
	return s.Body
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) SetHeaders(v map[string]*string) *AttachAggregateConfigRuleToCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) SetStatusCode(v int32) *AttachAggregateConfigRuleToCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) SetBody(v *AttachAggregateConfigRuleToCompliancePackResponseBody) *AttachAggregateConfigRuleToCompliancePackResponse {
	s.Body = v
	return s
}

func (s *AttachAggregateConfigRuleToCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
