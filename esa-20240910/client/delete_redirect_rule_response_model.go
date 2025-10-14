// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRedirectRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRedirectRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRedirectRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRedirectRuleResponseBody) *DeleteRedirectRuleResponse
	GetBody() *DeleteRedirectRuleResponseBody
}

type DeleteRedirectRuleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRedirectRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRedirectRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRedirectRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRedirectRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRedirectRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRedirectRuleResponse) GetBody() *DeleteRedirectRuleResponseBody {
	return s.Body
}

func (s *DeleteRedirectRuleResponse) SetHeaders(v map[string]*string) *DeleteRedirectRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRedirectRuleResponse) SetStatusCode(v int32) *DeleteRedirectRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRedirectRuleResponse) SetBody(v *DeleteRedirectRuleResponseBody) *DeleteRedirectRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteRedirectRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
