// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRewriteUrlRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRewriteUrlRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRewriteUrlRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListRewriteUrlRulesResponseBody) *ListRewriteUrlRulesResponse
	GetBody() *ListRewriteUrlRulesResponseBody
}

type ListRewriteUrlRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRewriteUrlRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRewriteUrlRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRewriteUrlRulesResponse) GoString() string {
	return s.String()
}

func (s *ListRewriteUrlRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRewriteUrlRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRewriteUrlRulesResponse) GetBody() *ListRewriteUrlRulesResponseBody {
	return s.Body
}

func (s *ListRewriteUrlRulesResponse) SetHeaders(v map[string]*string) *ListRewriteUrlRulesResponse {
	s.Headers = v
	return s
}

func (s *ListRewriteUrlRulesResponse) SetStatusCode(v int32) *ListRewriteUrlRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRewriteUrlRulesResponse) SetBody(v *ListRewriteUrlRulesResponseBody) *ListRewriteUrlRulesResponse {
	s.Body = v
	return s
}

func (s *ListRewriteUrlRulesResponse) Validate() error {
	return dara.Validate(s)
}
