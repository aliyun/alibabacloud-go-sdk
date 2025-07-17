// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDesensitizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDesensitizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDesensitizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListDesensitizationRuleResponseBody) *ListDesensitizationRuleResponse
	GetBody() *ListDesensitizationRuleResponseBody
}

type ListDesensitizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDesensitizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDesensitizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDesensitizationRuleResponse) GoString() string {
	return s.String()
}

func (s *ListDesensitizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDesensitizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDesensitizationRuleResponse) GetBody() *ListDesensitizationRuleResponseBody {
	return s.Body
}

func (s *ListDesensitizationRuleResponse) SetHeaders(v map[string]*string) *ListDesensitizationRuleResponse {
	s.Headers = v
	return s
}

func (s *ListDesensitizationRuleResponse) SetStatusCode(v int32) *ListDesensitizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDesensitizationRuleResponse) SetBody(v *ListDesensitizationRuleResponseBody) *ListDesensitizationRuleResponse {
	s.Body = v
	return s
}

func (s *ListDesensitizationRuleResponse) Validate() error {
	return dara.Validate(s)
}
