// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableLayer7CCRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableLayer7CCRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableLayer7CCRuleResponse
	GetStatusCode() *int32
	SetBody(v *DisableLayer7CCRuleResponseBody) *DisableLayer7CCRuleResponse
	GetBody() *DisableLayer7CCRuleResponseBody
}

type DisableLayer7CCRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableLayer7CCRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *DisableLayer7CCRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableLayer7CCRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableLayer7CCRuleResponse) GetBody() *DisableLayer7CCRuleResponseBody {
	return s.Body
}

func (s *DisableLayer7CCRuleResponse) SetHeaders(v map[string]*string) *DisableLayer7CCRuleResponse {
	s.Headers = v
	return s
}

func (s *DisableLayer7CCRuleResponse) SetStatusCode(v int32) *DisableLayer7CCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableLayer7CCRuleResponse) SetBody(v *DisableLayer7CCRuleResponseBody) *DisableLayer7CCRuleResponse {
	s.Body = v
	return s
}

func (s *DisableLayer7CCRuleResponse) Validate() error {
	return dara.Validate(s)
}
