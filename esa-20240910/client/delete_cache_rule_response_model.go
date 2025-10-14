// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCacheRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCacheRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCacheRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCacheRuleResponseBody) *DeleteCacheRuleResponse
	GetBody() *DeleteCacheRuleResponseBody
}

type DeleteCacheRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCacheRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCacheRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCacheRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteCacheRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCacheRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCacheRuleResponse) GetBody() *DeleteCacheRuleResponseBody {
	return s.Body
}

func (s *DeleteCacheRuleResponse) SetHeaders(v map[string]*string) *DeleteCacheRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteCacheRuleResponse) SetStatusCode(v int32) *DeleteCacheRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCacheRuleResponse) SetBody(v *DeleteCacheRuleResponseBody) *DeleteCacheRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteCacheRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
