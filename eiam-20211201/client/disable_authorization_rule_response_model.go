// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *DisableAuthorizationRuleResponseBody) *DisableAuthorizationRuleResponse
	GetBody() *DisableAuthorizationRuleResponseBody
}

type DisableAuthorizationRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAuthorizationRuleResponse) GetBody() *DisableAuthorizationRuleResponseBody {
	return s.Body
}

func (s *DisableAuthorizationRuleResponse) SetHeaders(v map[string]*string) *DisableAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableAuthorizationRuleResponse) SetStatusCode(v int32) *DisableAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAuthorizationRuleResponse) SetBody(v *DisableAuthorizationRuleResponseBody) *DisableAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *DisableAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
