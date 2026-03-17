// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddACLRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddACLRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddACLRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddACLRuleResponseBody) *AddACLRuleResponse
	GetBody() *AddACLRuleResponseBody
}

type AddACLRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddACLRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddACLRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddACLRuleResponse) GoString() string {
	return s.String()
}

func (s *AddACLRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddACLRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddACLRuleResponse) GetBody() *AddACLRuleResponseBody {
	return s.Body
}

func (s *AddACLRuleResponse) SetHeaders(v map[string]*string) *AddACLRuleResponse {
	s.Headers = v
	return s
}

func (s *AddACLRuleResponse) SetStatusCode(v int32) *AddACLRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddACLRuleResponse) SetBody(v *AddACLRuleResponseBody) *AddACLRuleResponse {
	s.Body = v
	return s
}

func (s *AddACLRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
