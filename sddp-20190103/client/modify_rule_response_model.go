// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRuleResponseBody) *ModifyRuleResponse
	GetBody() *ModifyRuleResponseBody
}

type ModifyRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRuleResponse) GetBody() *ModifyRuleResponseBody {
	return s.Body
}

func (s *ModifyRuleResponse) SetHeaders(v map[string]*string) *ModifyRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyRuleResponse) SetStatusCode(v int32) *ModifyRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRuleResponse) SetBody(v *ModifyRuleResponseBody) *ModifyRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyRuleResponse) Validate() error {
	return dara.Validate(s)
}
