// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFileProtectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFileProtectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFileProtectRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListFileProtectRuleResponseBody) *ListFileProtectRuleResponse
	GetBody() *ListFileProtectRuleResponseBody
}

type ListFileProtectRuleResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileProtectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileProtectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFileProtectRuleResponse) GoString() string {
	return s.String()
}

func (s *ListFileProtectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFileProtectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFileProtectRuleResponse) GetBody() *ListFileProtectRuleResponseBody {
	return s.Body
}

func (s *ListFileProtectRuleResponse) SetHeaders(v map[string]*string) *ListFileProtectRuleResponse {
	s.Headers = v
	return s
}

func (s *ListFileProtectRuleResponse) SetStatusCode(v int32) *ListFileProtectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileProtectRuleResponse) SetBody(v *ListFileProtectRuleResponseBody) *ListFileProtectRuleResponse {
	s.Body = v
	return s
}

func (s *ListFileProtectRuleResponse) Validate() error {
	return dara.Validate(s)
}
