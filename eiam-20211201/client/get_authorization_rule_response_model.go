// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthorizationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAuthorizationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAuthorizationRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetAuthorizationRuleResponseBody) *GetAuthorizationRuleResponse
	GetBody() *GetAuthorizationRuleResponseBody
}

type GetAuthorizationRuleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAuthorizationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAuthorizationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAuthorizationRuleResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAuthorizationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAuthorizationRuleResponse) GetBody() *GetAuthorizationRuleResponseBody {
	return s.Body
}

func (s *GetAuthorizationRuleResponse) SetHeaders(v map[string]*string) *GetAuthorizationRuleResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationRuleResponse) SetStatusCode(v int32) *GetAuthorizationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAuthorizationRuleResponse) SetBody(v *GetAuthorizationRuleResponseBody) *GetAuthorizationRuleResponse {
	s.Body = v
	return s
}

func (s *GetAuthorizationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
