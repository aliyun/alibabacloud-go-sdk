// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachConfigRuleToCompliancePackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachConfigRuleToCompliancePackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachConfigRuleToCompliancePackResponse
	GetStatusCode() *int32
	SetBody(v *AttachConfigRuleToCompliancePackResponseBody) *AttachConfigRuleToCompliancePackResponse
	GetBody() *AttachConfigRuleToCompliancePackResponseBody
}

type AttachConfigRuleToCompliancePackResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachConfigRuleToCompliancePackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachConfigRuleToCompliancePackResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachConfigRuleToCompliancePackResponse) GoString() string {
	return s.String()
}

func (s *AttachConfigRuleToCompliancePackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachConfigRuleToCompliancePackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachConfigRuleToCompliancePackResponse) GetBody() *AttachConfigRuleToCompliancePackResponseBody {
	return s.Body
}

func (s *AttachConfigRuleToCompliancePackResponse) SetHeaders(v map[string]*string) *AttachConfigRuleToCompliancePackResponse {
	s.Headers = v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponse) SetStatusCode(v int32) *AttachConfigRuleToCompliancePackResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponse) SetBody(v *AttachConfigRuleToCompliancePackResponseBody) *AttachConfigRuleToCompliancePackResponse {
	s.Body = v
	return s
}

func (s *AttachConfigRuleToCompliancePackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
