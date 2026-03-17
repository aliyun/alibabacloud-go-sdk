// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyACLRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyACLRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyACLRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyACLRuleResponseBody) *ModifyACLRuleResponse
	GetBody() *ModifyACLRuleResponseBody
}

type ModifyACLRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyACLRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyACLRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyACLRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyACLRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyACLRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyACLRuleResponse) GetBody() *ModifyACLRuleResponseBody {
	return s.Body
}

func (s *ModifyACLRuleResponse) SetHeaders(v map[string]*string) *ModifyACLRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyACLRuleResponse) SetStatusCode(v int32) *ModifyACLRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyACLRuleResponse) SetBody(v *ModifyACLRuleResponseBody) *ModifyACLRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyACLRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
