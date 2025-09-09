// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolicyUserScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPolicyUserScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPolicyUserScopeResponse
	GetStatusCode() *int32
	SetBody(v *GetPolicyUserScopeResponseBody) *GetPolicyUserScopeResponse
	GetBody() *GetPolicyUserScopeResponseBody
}

type GetPolicyUserScopeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPolicyUserScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPolicyUserScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPolicyUserScopeResponse) GoString() string {
	return s.String()
}

func (s *GetPolicyUserScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPolicyUserScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPolicyUserScopeResponse) GetBody() *GetPolicyUserScopeResponseBody {
	return s.Body
}

func (s *GetPolicyUserScopeResponse) SetHeaders(v map[string]*string) *GetPolicyUserScopeResponse {
	s.Headers = v
	return s
}

func (s *GetPolicyUserScopeResponse) SetStatusCode(v int32) *GetPolicyUserScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPolicyUserScopeResponse) SetBody(v *GetPolicyUserScopeResponseBody) *GetPolicyUserScopeResponse {
	s.Body = v
	return s
}

func (s *GetPolicyUserScopeResponse) Validate() error {
	return dara.Validate(s)
}
