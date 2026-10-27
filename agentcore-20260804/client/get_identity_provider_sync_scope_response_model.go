// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSyncScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIdentityProviderSyncScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIdentityProviderSyncScopeResponse
	GetStatusCode() *int32
	SetBody(v *GetIdentityProviderSyncScopeResponseBody) *GetIdentityProviderSyncScopeResponse
	GetBody() *GetIdentityProviderSyncScopeResponseBody
}

type GetIdentityProviderSyncScopeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIdentityProviderSyncScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityProviderSyncScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSyncScopeResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSyncScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIdentityProviderSyncScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIdentityProviderSyncScopeResponse) GetBody() *GetIdentityProviderSyncScopeResponseBody {
	return s.Body
}

func (s *GetIdentityProviderSyncScopeResponse) SetHeaders(v map[string]*string) *GetIdentityProviderSyncScopeResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityProviderSyncScopeResponse) SetStatusCode(v int32) *GetIdentityProviderSyncScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityProviderSyncScopeResponse) SetBody(v *GetIdentityProviderSyncScopeResponseBody) *GetIdentityProviderSyncScopeResponse {
	s.Body = v
	return s
}

func (s *GetIdentityProviderSyncScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
