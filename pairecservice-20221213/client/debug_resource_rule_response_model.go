// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugResourceRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DebugResourceRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DebugResourceRuleResponse
	GetStatusCode() *int32
	SetBody(v *DebugResourceRuleResponseBody) *DebugResourceRuleResponse
	GetBody() *DebugResourceRuleResponseBody
}

type DebugResourceRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DebugResourceRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DebugResourceRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DebugResourceRuleResponse) GoString() string {
	return s.String()
}

func (s *DebugResourceRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DebugResourceRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DebugResourceRuleResponse) GetBody() *DebugResourceRuleResponseBody {
	return s.Body
}

func (s *DebugResourceRuleResponse) SetHeaders(v map[string]*string) *DebugResourceRuleResponse {
	s.Headers = v
	return s
}

func (s *DebugResourceRuleResponse) SetStatusCode(v int32) *DebugResourceRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DebugResourceRuleResponse) SetBody(v *DebugResourceRuleResponseBody) *DebugResourceRuleResponse {
	s.Body = v
	return s
}

func (s *DebugResourceRuleResponse) Validate() error {
	return dara.Validate(s)
}
