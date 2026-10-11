// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEgressProxyRuleGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEgressProxyRuleGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEgressProxyRuleGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListEgressProxyRuleGroupsResponseBody) *ListEgressProxyRuleGroupsResponse
	GetBody() *ListEgressProxyRuleGroupsResponseBody
}

type ListEgressProxyRuleGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEgressProxyRuleGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEgressProxyRuleGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEgressProxyRuleGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListEgressProxyRuleGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEgressProxyRuleGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEgressProxyRuleGroupsResponse) GetBody() *ListEgressProxyRuleGroupsResponseBody {
	return s.Body
}

func (s *ListEgressProxyRuleGroupsResponse) SetHeaders(v map[string]*string) *ListEgressProxyRuleGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListEgressProxyRuleGroupsResponse) SetStatusCode(v int32) *ListEgressProxyRuleGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEgressProxyRuleGroupsResponse) SetBody(v *ListEgressProxyRuleGroupsResponseBody) *ListEgressProxyRuleGroupsResponse {
	s.Body = v
	return s
}

func (s *ListEgressProxyRuleGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
