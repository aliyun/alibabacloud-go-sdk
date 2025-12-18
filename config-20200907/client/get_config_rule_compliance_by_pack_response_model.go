// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigRuleComplianceByPackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConfigRuleComplianceByPackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConfigRuleComplianceByPackResponse
	GetStatusCode() *int32
	SetBody(v *GetConfigRuleComplianceByPackResponseBody) *GetConfigRuleComplianceByPackResponse
	GetBody() *GetConfigRuleComplianceByPackResponseBody
}

type GetConfigRuleComplianceByPackResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConfigRuleComplianceByPackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConfigRuleComplianceByPackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConfigRuleComplianceByPackResponse) GoString() string {
	return s.String()
}

func (s *GetConfigRuleComplianceByPackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConfigRuleComplianceByPackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConfigRuleComplianceByPackResponse) GetBody() *GetConfigRuleComplianceByPackResponseBody {
	return s.Body
}

func (s *GetConfigRuleComplianceByPackResponse) SetHeaders(v map[string]*string) *GetConfigRuleComplianceByPackResponse {
	s.Headers = v
	return s
}

func (s *GetConfigRuleComplianceByPackResponse) SetStatusCode(v int32) *GetConfigRuleComplianceByPackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConfigRuleComplianceByPackResponse) SetBody(v *GetConfigRuleComplianceByPackResponseBody) *GetConfigRuleComplianceByPackResponse {
	s.Body = v
	return s
}

func (s *GetConfigRuleComplianceByPackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
