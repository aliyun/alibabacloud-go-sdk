// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpRequestHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHttpRequestHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHttpRequestHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHttpRequestHeaderModificationRuleResponseBody) *UpdateHttpRequestHeaderModificationRuleResponse
	GetBody() *UpdateHttpRequestHeaderModificationRuleResponseBody
}

type UpdateHttpRequestHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHttpRequestHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHttpRequestHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpRequestHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateHttpRequestHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHttpRequestHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHttpRequestHeaderModificationRuleResponse) GetBody() *UpdateHttpRequestHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *UpdateHttpRequestHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *UpdateHttpRequestHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleResponse) SetStatusCode(v int32) *UpdateHttpRequestHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleResponse) SetBody(v *UpdateHttpRequestHeaderModificationRuleResponseBody) *UpdateHttpRequestHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleResponse) Validate() error {
	return dara.Validate(s)
}
