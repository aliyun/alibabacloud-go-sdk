// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaLifecycleRuleResponseBody) *ListMediaLifecycleRuleResponse
	GetBody() *ListMediaLifecycleRuleResponseBody
}

type ListMediaLifecycleRuleResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *ListMediaLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaLifecycleRuleResponse) GetBody() *ListMediaLifecycleRuleResponseBody {
	return s.Body
}

func (s *ListMediaLifecycleRuleResponse) SetHeaders(v map[string]*string) *ListMediaLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *ListMediaLifecycleRuleResponse) SetStatusCode(v int32) *ListMediaLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaLifecycleRuleResponse) SetBody(v *ListMediaLifecycleRuleResponseBody) *ListMediaLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *ListMediaLifecycleRuleResponse) Validate() error {
	return dara.Validate(s)
}
