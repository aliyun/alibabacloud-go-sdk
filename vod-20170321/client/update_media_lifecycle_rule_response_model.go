// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaLifecycleRuleResponseBody) *UpdateMediaLifecycleRuleResponse
	GetBody() *UpdateMediaLifecycleRuleResponseBody
}

type UpdateMediaLifecycleRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaLifecycleRuleResponse) GetBody() *UpdateMediaLifecycleRuleResponseBody {
	return s.Body
}

func (s *UpdateMediaLifecycleRuleResponse) SetHeaders(v map[string]*string) *UpdateMediaLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaLifecycleRuleResponse) SetStatusCode(v int32) *UpdateMediaLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaLifecycleRuleResponse) SetBody(v *UpdateMediaLifecycleRuleResponseBody) *UpdateMediaLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaLifecycleRuleResponse) Validate() error {
	return dara.Validate(s)
}
