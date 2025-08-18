// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCacheRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCacheRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCacheRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCacheRuleResponseBody) *UpdateCacheRuleResponse
	GetBody() *UpdateCacheRuleResponseBody
}

type UpdateCacheRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCacheRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCacheRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCacheRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateCacheRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCacheRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCacheRuleResponse) GetBody() *UpdateCacheRuleResponseBody {
	return s.Body
}

func (s *UpdateCacheRuleResponse) SetHeaders(v map[string]*string) *UpdateCacheRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateCacheRuleResponse) SetStatusCode(v int32) *UpdateCacheRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCacheRuleResponse) SetBody(v *UpdateCacheRuleResponseBody) *UpdateCacheRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateCacheRuleResponse) Validate() error {
	return dara.Validate(s)
}
