// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResolverRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResolverRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResolverRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResolverRuleResponseBody) *DeleteResolverRuleResponse
	GetBody() *DeleteResolverRuleResponseBody
}

type DeleteResolverRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResolverRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResolverRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResolverRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteResolverRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResolverRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResolverRuleResponse) GetBody() *DeleteResolverRuleResponseBody {
	return s.Body
}

func (s *DeleteResolverRuleResponse) SetHeaders(v map[string]*string) *DeleteResolverRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteResolverRuleResponse) SetStatusCode(v int32) *DeleteResolverRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResolverRuleResponse) SetBody(v *DeleteResolverRuleResponseBody) *DeleteResolverRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteResolverRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
