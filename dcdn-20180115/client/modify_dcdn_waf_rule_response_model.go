// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDcdnWafRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDcdnWafRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDcdnWafRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDcdnWafRuleResponseBody) *ModifyDcdnWafRuleResponse
	GetBody() *ModifyDcdnWafRuleResponseBody
}

type ModifyDcdnWafRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDcdnWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDcdnWafRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDcdnWafRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyDcdnWafRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDcdnWafRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDcdnWafRuleResponse) GetBody() *ModifyDcdnWafRuleResponseBody {
	return s.Body
}

func (s *ModifyDcdnWafRuleResponse) SetHeaders(v map[string]*string) *ModifyDcdnWafRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyDcdnWafRuleResponse) SetStatusCode(v int32) *ModifyDcdnWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDcdnWafRuleResponse) SetBody(v *ModifyDcdnWafRuleResponseBody) *ModifyDcdnWafRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyDcdnWafRuleResponse) Validate() error {
	return dara.Validate(s)
}
