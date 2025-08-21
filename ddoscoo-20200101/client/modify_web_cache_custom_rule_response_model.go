// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCacheCustomRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyWebCacheCustomRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyWebCacheCustomRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyWebCacheCustomRuleResponseBody) *ModifyWebCacheCustomRuleResponse
	GetBody() *ModifyWebCacheCustomRuleResponseBody
}

type ModifyWebCacheCustomRuleResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyWebCacheCustomRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyWebCacheCustomRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCacheCustomRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheCustomRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyWebCacheCustomRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyWebCacheCustomRuleResponse) GetBody() *ModifyWebCacheCustomRuleResponseBody {
	return s.Body
}

func (s *ModifyWebCacheCustomRuleResponse) SetHeaders(v map[string]*string) *ModifyWebCacheCustomRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyWebCacheCustomRuleResponse) SetStatusCode(v int32) *ModifyWebCacheCustomRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyWebCacheCustomRuleResponse) SetBody(v *ModifyWebCacheCustomRuleResponseBody) *ModifyWebCacheCustomRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyWebCacheCustomRuleResponse) Validate() error {
	return dara.Validate(s)
}
