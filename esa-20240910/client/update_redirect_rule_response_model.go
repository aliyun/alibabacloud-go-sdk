// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRedirectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRedirectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRedirectRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRedirectRuleResponseBody) *UpdateRedirectRuleResponse
	GetBody() *UpdateRedirectRuleResponseBody
}

type UpdateRedirectRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRedirectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRedirectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRedirectRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRedirectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRedirectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRedirectRuleResponse) GetBody() *UpdateRedirectRuleResponseBody {
	return s.Body
}

func (s *UpdateRedirectRuleResponse) SetHeaders(v map[string]*string) *UpdateRedirectRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRedirectRuleResponse) SetStatusCode(v int32) *UpdateRedirectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRedirectRuleResponse) SetBody(v *UpdateRedirectRuleResponseBody) *UpdateRedirectRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateRedirectRuleResponse) Validate() error {
	return dara.Validate(s)
}
