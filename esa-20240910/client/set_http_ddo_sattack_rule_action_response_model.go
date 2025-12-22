// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackRuleActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetHttpDDoSAttackRuleActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetHttpDDoSAttackRuleActionResponse
	GetStatusCode() *int32
	SetBody(v *SetHttpDDoSAttackRuleActionResponseBody) *SetHttpDDoSAttackRuleActionResponse
	GetBody() *SetHttpDDoSAttackRuleActionResponseBody
}

type SetHttpDDoSAttackRuleActionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetHttpDDoSAttackRuleActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetHttpDDoSAttackRuleActionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackRuleActionResponse) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackRuleActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetHttpDDoSAttackRuleActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetHttpDDoSAttackRuleActionResponse) GetBody() *SetHttpDDoSAttackRuleActionResponseBody {
	return s.Body
}

func (s *SetHttpDDoSAttackRuleActionResponse) SetHeaders(v map[string]*string) *SetHttpDDoSAttackRuleActionResponse {
	s.Headers = v
	return s
}

func (s *SetHttpDDoSAttackRuleActionResponse) SetStatusCode(v int32) *SetHttpDDoSAttackRuleActionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetHttpDDoSAttackRuleActionResponse) SetBody(v *SetHttpDDoSAttackRuleActionResponseBody) *SetHttpDDoSAttackRuleActionResponse {
	s.Body = v
	return s
}

func (s *SetHttpDDoSAttackRuleActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
