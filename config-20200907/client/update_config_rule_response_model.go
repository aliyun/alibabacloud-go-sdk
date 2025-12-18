// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConfigRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConfigRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConfigRuleResponseBody) *UpdateConfigRuleResponse
	GetBody() *UpdateConfigRuleResponseBody
}

type UpdateConfigRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConfigRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConfigRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateConfigRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConfigRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConfigRuleResponse) GetBody() *UpdateConfigRuleResponseBody {
	return s.Body
}

func (s *UpdateConfigRuleResponse) SetHeaders(v map[string]*string) *UpdateConfigRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateConfigRuleResponse) SetStatusCode(v int32) *UpdateConfigRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConfigRuleResponse) SetBody(v *UpdateConfigRuleResponseBody) *UpdateConfigRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateConfigRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
