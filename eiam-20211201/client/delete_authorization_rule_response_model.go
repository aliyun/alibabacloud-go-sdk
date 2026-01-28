// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAuthorizationRuleResponseBody) *DeleteAuthorizationRuleResponse
	GetBody() *DeleteAuthorizationRuleResponseBody
}

type DeleteAuthorizationRuleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAuthorizationRuleResponse) GetBody() *DeleteAuthorizationRuleResponseBody {
	return s.Body
}

func (s *DeleteAuthorizationRuleResponse) SetHeaders(v map[string]*string) *DeleteAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteAuthorizationRuleResponse) SetStatusCode(v int32) *DeleteAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAuthorizationRuleResponse) SetBody(v *DeleteAuthorizationRuleResponseBody) *DeleteAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
