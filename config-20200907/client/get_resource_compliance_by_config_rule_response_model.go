// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceByConfigRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceComplianceByConfigRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceComplianceByConfigRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceComplianceByConfigRuleResponseBody) *GetResourceComplianceByConfigRuleResponse
	GetBody() *GetResourceComplianceByConfigRuleResponseBody
}

type GetResourceComplianceByConfigRuleResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceComplianceByConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceComplianceByConfigRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceByConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceByConfigRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceComplianceByConfigRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceComplianceByConfigRuleResponse) GetBody() *GetResourceComplianceByConfigRuleResponseBody {
	return s.Body
}

func (s *GetResourceComplianceByConfigRuleResponse) SetHeaders(v map[string]*string) *GetResourceComplianceByConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponse) SetStatusCode(v int32) *GetResourceComplianceByConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponse) SetBody(v *GetResourceComplianceByConfigRuleResponseBody) *GetResourceComplianceByConfigRuleResponse {
	s.Body = v
	return s
}

func (s *GetResourceComplianceByConfigRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
