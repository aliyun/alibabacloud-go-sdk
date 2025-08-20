// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRepoBuildRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRepoBuildRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRepoBuildRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListRepoBuildRuleResponseBody) *ListRepoBuildRuleResponse
	GetBody() *ListRepoBuildRuleResponseBody
}

type ListRepoBuildRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRepoBuildRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRepoBuildRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRepoBuildRuleResponse) GoString() string {
	return s.String()
}

func (s *ListRepoBuildRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRepoBuildRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRepoBuildRuleResponse) GetBody() *ListRepoBuildRuleResponseBody {
	return s.Body
}

func (s *ListRepoBuildRuleResponse) SetHeaders(v map[string]*string) *ListRepoBuildRuleResponse {
	s.Headers = v
	return s
}

func (s *ListRepoBuildRuleResponse) SetStatusCode(v int32) *ListRepoBuildRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRepoBuildRuleResponse) SetBody(v *ListRepoBuildRuleResponseBody) *ListRepoBuildRuleResponse {
	s.Body = v
	return s
}

func (s *ListRepoBuildRuleResponse) Validate() error {
	return dara.Validate(s)
}
