// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDefenseRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDefenseRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDefenseRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDefenseRuleResponseBody) *ModifyDefenseRuleResponse
	GetBody() *ModifyDefenseRuleResponseBody
}

type ModifyDefenseRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDefenseRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDefenseRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDefenseRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyDefenseRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDefenseRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDefenseRuleResponse) GetBody() *ModifyDefenseRuleResponseBody {
	return s.Body
}

func (s *ModifyDefenseRuleResponse) SetHeaders(v map[string]*string) *ModifyDefenseRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyDefenseRuleResponse) SetStatusCode(v int32) *ModifyDefenseRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDefenseRuleResponse) SetBody(v *ModifyDefenseRuleResponseBody) *ModifyDefenseRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyDefenseRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
