// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySupabaseSecurityIPListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySupabaseSecurityIPListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySupabaseSecurityIPListResponse
	GetStatusCode() *int32
	SetBody(v *ModifySupabaseSecurityIPListResponseBody) *ModifySupabaseSecurityIPListResponse
	GetBody() *ModifySupabaseSecurityIPListResponseBody
}

type ModifySupabaseSecurityIPListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySupabaseSecurityIPListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySupabaseSecurityIPListResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySupabaseSecurityIPListResponse) GoString() string {
	return s.String()
}

func (s *ModifySupabaseSecurityIPListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySupabaseSecurityIPListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySupabaseSecurityIPListResponse) GetBody() *ModifySupabaseSecurityIPListResponseBody {
	return s.Body
}

func (s *ModifySupabaseSecurityIPListResponse) SetHeaders(v map[string]*string) *ModifySupabaseSecurityIPListResponse {
	s.Headers = v
	return s
}

func (s *ModifySupabaseSecurityIPListResponse) SetStatusCode(v int32) *ModifySupabaseSecurityIPListResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySupabaseSecurityIPListResponse) SetBody(v *ModifySupabaseSecurityIPListResponseBody) *ModifySupabaseSecurityIPListResponse {
	s.Body = v
	return s
}

func (s *ModifySupabaseSecurityIPListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
