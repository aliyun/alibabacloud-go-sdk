// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCacheRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCacheRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateCacheRuleResponseBody) *CreateCacheRuleResponse
	GetBody() *CreateCacheRuleResponseBody
}

type CreateCacheRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCacheRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCacheRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateCacheRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCacheRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCacheRuleResponse) GetBody() *CreateCacheRuleResponseBody {
	return s.Body
}

func (s *CreateCacheRuleResponse) SetHeaders(v map[string]*string) *CreateCacheRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateCacheRuleResponse) SetStatusCode(v int32) *CreateCacheRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCacheRuleResponse) SetBody(v *CreateCacheRuleResponseBody) *CreateCacheRuleResponse {
	s.Body = v
	return s
}

func (s *CreateCacheRuleResponse) Validate() error {
	return dara.Validate(s)
}
