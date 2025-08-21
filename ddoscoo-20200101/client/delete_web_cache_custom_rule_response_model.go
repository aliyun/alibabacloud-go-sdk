// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCacheCustomRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWebCacheCustomRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWebCacheCustomRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWebCacheCustomRuleResponseBody) *DeleteWebCacheCustomRuleResponse
	GetBody() *DeleteWebCacheCustomRuleResponseBody
}

type DeleteWebCacheCustomRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWebCacheCustomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWebCacheCustomRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCacheCustomRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteWebCacheCustomRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWebCacheCustomRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWebCacheCustomRuleResponse) GetBody() *DeleteWebCacheCustomRuleResponseBody {
	return s.Body
}

func (s *DeleteWebCacheCustomRuleResponse) SetHeaders(v map[string]*string) *DeleteWebCacheCustomRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteWebCacheCustomRuleResponse) SetStatusCode(v int32) *DeleteWebCacheCustomRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWebCacheCustomRuleResponse) SetBody(v *DeleteWebCacheCustomRuleResponseBody) *DeleteWebCacheCustomRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteWebCacheCustomRuleResponse) Validate() error {
	return dara.Validate(s)
}
