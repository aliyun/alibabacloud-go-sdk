// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateSupabaseForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateSupabaseForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateSupabaseForAdminResponse
	GetStatusCode() *int32
	SetBody(v *AllocateSupabaseForAdminResponseBody) *AllocateSupabaseForAdminResponse
	GetBody() *AllocateSupabaseForAdminResponseBody
}

type AllocateSupabaseForAdminResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateSupabaseForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateSupabaseForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateSupabaseForAdminResponse) GoString() string {
	return s.String()
}

func (s *AllocateSupabaseForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateSupabaseForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateSupabaseForAdminResponse) GetBody() *AllocateSupabaseForAdminResponseBody {
	return s.Body
}

func (s *AllocateSupabaseForAdminResponse) SetHeaders(v map[string]*string) *AllocateSupabaseForAdminResponse {
	s.Headers = v
	return s
}

func (s *AllocateSupabaseForAdminResponse) SetStatusCode(v int32) *AllocateSupabaseForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateSupabaseForAdminResponse) SetBody(v *AllocateSupabaseForAdminResponseBody) *AllocateSupabaseForAdminResponse {
	s.Body = v
	return s
}

func (s *AllocateSupabaseForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
