// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCacheRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCacheRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCacheRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetCacheRuleResponseBody) *GetCacheRuleResponse
	GetBody() *GetCacheRuleResponseBody
}

type GetCacheRuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCacheRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCacheRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCacheRuleResponse) GoString() string {
	return s.String()
}

func (s *GetCacheRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCacheRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCacheRuleResponse) GetBody() *GetCacheRuleResponseBody {
	return s.Body
}

func (s *GetCacheRuleResponse) SetHeaders(v map[string]*string) *GetCacheRuleResponse {
	s.Headers = v
	return s
}

func (s *GetCacheRuleResponse) SetStatusCode(v int32) *GetCacheRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCacheRuleResponse) SetBody(v *GetCacheRuleResponseBody) *GetCacheRuleResponse {
	s.Body = v
	return s
}

func (s *GetCacheRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
