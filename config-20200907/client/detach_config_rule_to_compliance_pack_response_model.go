// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachConfigRuleToCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachConfigRuleToCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachConfigRuleToCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *DetachConfigRuleToCompliancePackResponseBody) *DetachConfigRuleToCompliancePackResponse
	GetBody() *DetachConfigRuleToCompliancePackResponseBody
}

type DetachConfigRuleToCompliancePackResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachConfigRuleToCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachConfigRuleToCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachConfigRuleToCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *DetachConfigRuleToCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachConfigRuleToCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachConfigRuleToCompliancePackResponse) GetBody() *DetachConfigRuleToCompliancePackResponseBody {
	return s.Body
}

func (s *DetachConfigRuleToCompliancePackResponse) SetHeaders(v map[string]*string) *DetachConfigRuleToCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponse) SetStatusCode(v int32) *DetachConfigRuleToCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponse) SetBody(v *DetachConfigRuleToCompliancePackResponseBody) *DetachConfigRuleToCompliancePackResponse {
	s.Body = v
	return s
}

func (s *DetachConfigRuleToCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
