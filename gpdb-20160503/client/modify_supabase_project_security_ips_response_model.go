// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseProjectSecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySupabaseProjectSecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySupabaseProjectSecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *ModifySupabaseProjectSecurityIpsResponseBody) *ModifySupabaseProjectSecurityIpsResponse
	GetBody() *ModifySupabaseProjectSecurityIpsResponseBody
}

type ModifySupabaseProjectSecurityIpsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySupabaseProjectSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySupabaseProjectSecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseProjectSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifySupabaseProjectSecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySupabaseProjectSecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySupabaseProjectSecurityIpsResponse) GetBody() *ModifySupabaseProjectSecurityIpsResponseBody {
	return s.Body
}

func (s *ModifySupabaseProjectSecurityIpsResponse) SetHeaders(v map[string]*string) *ModifySupabaseProjectSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifySupabaseProjectSecurityIpsResponse) SetStatusCode(v int32) *ModifySupabaseProjectSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySupabaseProjectSecurityIpsResponse) SetBody(v *ModifySupabaseProjectSecurityIpsResponseBody) *ModifySupabaseProjectSecurityIpsResponse {
	s.Body = v
	return s
}

func (s *ModifySupabaseProjectSecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
