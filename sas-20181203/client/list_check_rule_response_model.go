// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckRuleResponseBody) *ListCheckRuleResponse
	GetBody() *ListCheckRuleResponseBody
}

type ListCheckRuleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckRuleResponse) GoString() string {
	return s.String()
}

func (s *ListCheckRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckRuleResponse) GetBody() *ListCheckRuleResponseBody {
	return s.Body
}

func (s *ListCheckRuleResponse) SetHeaders(v map[string]*string) *ListCheckRuleResponse {
	s.Headers = v
	return s
}

func (s *ListCheckRuleResponse) SetStatusCode(v int32) *ListCheckRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckRuleResponse) SetBody(v *ListCheckRuleResponseBody) *ListCheckRuleResponse {
	s.Body = v
	return s
}

func (s *ListCheckRuleResponse) Validate() error {
	return dara.Validate(s)
}
