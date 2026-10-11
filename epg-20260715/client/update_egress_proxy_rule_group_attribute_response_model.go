// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyRuleGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEgressProxyRuleGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEgressProxyRuleGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEgressProxyRuleGroupAttributeResponseBody) *UpdateEgressProxyRuleGroupAttributeResponse
	GetBody() *UpdateEgressProxyRuleGroupAttributeResponseBody
}

type UpdateEgressProxyRuleGroupAttributeResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEgressProxyRuleGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEgressProxyRuleGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyRuleGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyRuleGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEgressProxyRuleGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEgressProxyRuleGroupAttributeResponse) GetBody() *UpdateEgressProxyRuleGroupAttributeResponseBody {
	return s.Body
}

func (s *UpdateEgressProxyRuleGroupAttributeResponse) SetHeaders(v map[string]*string) *UpdateEgressProxyRuleGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeResponse) SetStatusCode(v int32) *UpdateEgressProxyRuleGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeResponse) SetBody(v *UpdateEgressProxyRuleGroupAttributeResponseBody) *UpdateEgressProxyRuleGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateEgressProxyRuleGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
