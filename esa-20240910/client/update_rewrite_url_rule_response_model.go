// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRewriteUrlRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRewriteUrlRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRewriteUrlRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRewriteUrlRuleResponseBody) *UpdateRewriteUrlRuleResponse
	GetBody() *UpdateRewriteUrlRuleResponseBody
}

type UpdateRewriteUrlRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRewriteUrlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRewriteUrlRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRewriteUrlRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRewriteUrlRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRewriteUrlRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRewriteUrlRuleResponse) GetBody() *UpdateRewriteUrlRuleResponseBody {
	return s.Body
}

func (s *UpdateRewriteUrlRuleResponse) SetHeaders(v map[string]*string) *UpdateRewriteUrlRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRewriteUrlRuleResponse) SetStatusCode(v int32) *UpdateRewriteUrlRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRewriteUrlRuleResponse) SetBody(v *UpdateRewriteUrlRuleResponseBody) *UpdateRewriteUrlRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateRewriteUrlRuleResponse) Validate() error {
	return dara.Validate(s)
}
