// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyRuleGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEgressProxyRuleGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEgressProxyRuleGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEgressProxyRuleGroupResponseBody) *DeleteEgressProxyRuleGroupResponse
	GetBody() *DeleteEgressProxyRuleGroupResponseBody
}

type DeleteEgressProxyRuleGroupResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEgressProxyRuleGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEgressProxyRuleGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyRuleGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyRuleGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEgressProxyRuleGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEgressProxyRuleGroupResponse) GetBody() *DeleteEgressProxyRuleGroupResponseBody {
	return s.Body
}

func (s *DeleteEgressProxyRuleGroupResponse) SetHeaders(v map[string]*string) *DeleteEgressProxyRuleGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteEgressProxyRuleGroupResponse) SetStatusCode(v int32) *DeleteEgressProxyRuleGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEgressProxyRuleGroupResponse) SetBody(v *DeleteEgressProxyRuleGroupResponseBody) *DeleteEgressProxyRuleGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteEgressProxyRuleGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
