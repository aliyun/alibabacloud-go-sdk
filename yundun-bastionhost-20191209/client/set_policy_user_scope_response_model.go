// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPolicyUserScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPolicyUserScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPolicyUserScopeResponse
	GetStatusCode() *int32
	SetBody(v *SetPolicyUserScopeResponseBody) *SetPolicyUserScopeResponse
	GetBody() *SetPolicyUserScopeResponseBody
}

type SetPolicyUserScopeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPolicyUserScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPolicyUserScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPolicyUserScopeResponse) GoString() string {
	return s.String()
}

func (s *SetPolicyUserScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPolicyUserScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPolicyUserScopeResponse) GetBody() *SetPolicyUserScopeResponseBody {
	return s.Body
}

func (s *SetPolicyUserScopeResponse) SetHeaders(v map[string]*string) *SetPolicyUserScopeResponse {
	s.Headers = v
	return s
}

func (s *SetPolicyUserScopeResponse) SetStatusCode(v int32) *SetPolicyUserScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPolicyUserScopeResponse) SetBody(v *SetPolicyUserScopeResponseBody) *SetPolicyUserScopeResponse {
	s.Body = v
	return s
}

func (s *SetPolicyUserScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
