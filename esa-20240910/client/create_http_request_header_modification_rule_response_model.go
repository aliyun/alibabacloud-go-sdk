// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpRequestHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHttpRequestHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHttpRequestHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateHttpRequestHeaderModificationRuleResponseBody) *CreateHttpRequestHeaderModificationRuleResponse
	GetBody() *CreateHttpRequestHeaderModificationRuleResponseBody
}

type CreateHttpRequestHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHttpRequestHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHttpRequestHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpRequestHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateHttpRequestHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHttpRequestHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHttpRequestHeaderModificationRuleResponse) GetBody() *CreateHttpRequestHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *CreateHttpRequestHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *CreateHttpRequestHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleResponse) SetStatusCode(v int32) *CreateHttpRequestHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleResponse) SetBody(v *CreateHttpRequestHeaderModificationRuleResponseBody) *CreateHttpRequestHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleResponse) Validate() error {
	return dara.Validate(s)
}
