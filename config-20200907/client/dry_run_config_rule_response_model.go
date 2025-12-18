// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunConfigRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DryRunConfigRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DryRunConfigRuleResponse
	GetStatusCode() *int32
	SetBody(v *DryRunConfigRuleResponseBody) *DryRunConfigRuleResponse
	GetBody() *DryRunConfigRuleResponseBody
}

type DryRunConfigRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DryRunConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DryRunConfigRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DryRunConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *DryRunConfigRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DryRunConfigRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DryRunConfigRuleResponse) GetBody() *DryRunConfigRuleResponseBody {
	return s.Body
}

func (s *DryRunConfigRuleResponse) SetHeaders(v map[string]*string) *DryRunConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *DryRunConfigRuleResponse) SetStatusCode(v int32) *DryRunConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DryRunConfigRuleResponse) SetBody(v *DryRunConfigRuleResponseBody) *DryRunConfigRuleResponse {
	s.Body = v
	return s
}

func (s *DryRunConfigRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
