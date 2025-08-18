// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRewriteUrlRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRewriteUrlRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRewriteUrlRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRewriteUrlRuleResponseBody) *DeleteRewriteUrlRuleResponse
	GetBody() *DeleteRewriteUrlRuleResponseBody
}

type DeleteRewriteUrlRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRewriteUrlRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRewriteUrlRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRewriteUrlRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRewriteUrlRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRewriteUrlRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRewriteUrlRuleResponse) GetBody() *DeleteRewriteUrlRuleResponseBody {
	return s.Body
}

func (s *DeleteRewriteUrlRuleResponse) SetHeaders(v map[string]*string) *DeleteRewriteUrlRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRewriteUrlRuleResponse) SetStatusCode(v int32) *DeleteRewriteUrlRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRewriteUrlRuleResponse) SetBody(v *DeleteRewriteUrlRuleResponseBody) *DeleteRewriteUrlRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteRewriteUrlRuleResponse) Validate() error {
	return dara.Validate(s)
}
