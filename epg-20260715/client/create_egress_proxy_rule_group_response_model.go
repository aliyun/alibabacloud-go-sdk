// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyRuleGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEgressProxyRuleGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEgressProxyRuleGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateEgressProxyRuleGroupResponseBody) *CreateEgressProxyRuleGroupResponse
	GetBody() *CreateEgressProxyRuleGroupResponseBody
}

type CreateEgressProxyRuleGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEgressProxyRuleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEgressProxyRuleGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyRuleGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyRuleGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEgressProxyRuleGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEgressProxyRuleGroupResponse) GetBody() *CreateEgressProxyRuleGroupResponseBody {
	return s.Body
}

func (s *CreateEgressProxyRuleGroupResponse) SetHeaders(v map[string]*string) *CreateEgressProxyRuleGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateEgressProxyRuleGroupResponse) SetStatusCode(v int32) *CreateEgressProxyRuleGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEgressProxyRuleGroupResponse) SetBody(v *CreateEgressProxyRuleGroupResponseBody) *CreateEgressProxyRuleGroupResponse {
	s.Body = v
	return s
}

func (s *CreateEgressProxyRuleGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
