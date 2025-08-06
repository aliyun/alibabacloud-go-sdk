// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMediaLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMediaLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateMediaLifecycleRuleResponseBody) *CreateMediaLifecycleRuleResponse
	GetBody() *CreateMediaLifecycleRuleResponseBody
}

type CreateMediaLifecycleRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMediaLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMediaLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateMediaLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMediaLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMediaLifecycleRuleResponse) GetBody() *CreateMediaLifecycleRuleResponseBody {
	return s.Body
}

func (s *CreateMediaLifecycleRuleResponse) SetHeaders(v map[string]*string) *CreateMediaLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateMediaLifecycleRuleResponse) SetStatusCode(v int32) *CreateMediaLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMediaLifecycleRuleResponse) SetBody(v *CreateMediaLifecycleRuleResponseBody) *CreateMediaLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *CreateMediaLifecycleRuleResponse) Validate() error {
	return dara.Validate(s)
}
