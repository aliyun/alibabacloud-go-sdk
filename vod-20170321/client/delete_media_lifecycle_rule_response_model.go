// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaLifecycleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMediaLifecycleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMediaLifecycleRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMediaLifecycleRuleResponseBody) *DeleteMediaLifecycleRuleResponse
	GetBody() *DeleteMediaLifecycleRuleResponseBody
}

type DeleteMediaLifecycleRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMediaLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMediaLifecycleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteMediaLifecycleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMediaLifecycleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMediaLifecycleRuleResponse) GetBody() *DeleteMediaLifecycleRuleResponseBody {
	return s.Body
}

func (s *DeleteMediaLifecycleRuleResponse) SetHeaders(v map[string]*string) *DeleteMediaLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteMediaLifecycleRuleResponse) SetStatusCode(v int32) *DeleteMediaLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMediaLifecycleRuleResponse) SetBody(v *DeleteMediaLifecycleRuleResponseBody) *DeleteMediaLifecycleRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteMediaLifecycleRuleResponse) Validate() error {
	return dara.Validate(s)
}
