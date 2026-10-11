// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyRuleGroupAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEgressProxyRuleGroupAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEgressProxyRuleGroupAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetEgressProxyRuleGroupAttributeResponseBody) *GetEgressProxyRuleGroupAttributeResponse
	GetBody() *GetEgressProxyRuleGroupAttributeResponseBody
}

type GetEgressProxyRuleGroupAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEgressProxyRuleGroupAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEgressProxyRuleGroupAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyRuleGroupAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetEgressProxyRuleGroupAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEgressProxyRuleGroupAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEgressProxyRuleGroupAttributeResponse) GetBody() *GetEgressProxyRuleGroupAttributeResponseBody {
	return s.Body
}

func (s *GetEgressProxyRuleGroupAttributeResponse) SetHeaders(v map[string]*string) *GetEgressProxyRuleGroupAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponse) SetStatusCode(v int32) *GetEgressProxyRuleGroupAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponse) SetBody(v *GetEgressProxyRuleGroupAttributeResponseBody) *GetEgressProxyRuleGroupAttributeResponse {
	s.Body = v
	return s
}

func (s *GetEgressProxyRuleGroupAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
