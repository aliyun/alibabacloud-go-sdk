// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLayer7CCRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLayer7CCRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLayer7CCRuleResponse
	GetStatusCode() *int32
	SetBody(v *AddLayer7CCRuleResponseBody) *AddLayer7CCRuleResponse
	GetBody() *AddLayer7CCRuleResponseBody
}

type AddLayer7CCRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLayer7CCRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLayer7CCRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLayer7CCRuleResponse) GoString() string {
	return s.String()
}

func (s *AddLayer7CCRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLayer7CCRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLayer7CCRuleResponse) GetBody() *AddLayer7CCRuleResponseBody {
	return s.Body
}

func (s *AddLayer7CCRuleResponse) SetHeaders(v map[string]*string) *AddLayer7CCRuleResponse {
	s.Headers = v
	return s
}

func (s *AddLayer7CCRuleResponse) SetStatusCode(v int32) *AddLayer7CCRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLayer7CCRuleResponse) SetBody(v *AddLayer7CCRuleResponseBody) *AddLayer7CCRuleResponse {
	s.Body = v
	return s
}

func (s *AddLayer7CCRuleResponse) Validate() error {
	return dara.Validate(s)
}
