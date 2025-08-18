// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWafRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWafRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWafRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWafRuleResponseBody) *UpdateWafRuleResponse
	GetBody() *UpdateWafRuleResponseBody
}

type UpdateWafRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWafRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWafRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWafRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateWafRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWafRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWafRuleResponse) GetBody() *UpdateWafRuleResponseBody {
	return s.Body
}

func (s *UpdateWafRuleResponse) SetHeaders(v map[string]*string) *UpdateWafRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateWafRuleResponse) SetStatusCode(v int32) *UpdateWafRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWafRuleResponse) SetBody(v *UpdateWafRuleResponseBody) *UpdateWafRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateWafRuleResponse) Validate() error {
	return dara.Validate(s)
}
