// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseAuthConfigsForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySupabaseAuthConfigsForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySupabaseAuthConfigsForAdminResponse
	GetStatusCode() *int32
	SetBody(v *QuerySupabaseAuthConfigsForAdminResponseBody) *QuerySupabaseAuthConfigsForAdminResponse
	GetBody() *QuerySupabaseAuthConfigsForAdminResponseBody
}

type QuerySupabaseAuthConfigsForAdminResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySupabaseAuthConfigsForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySupabaseAuthConfigsForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseAuthConfigsForAdminResponse) GoString() string {
	return s.String()
}

func (s *QuerySupabaseAuthConfigsForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySupabaseAuthConfigsForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySupabaseAuthConfigsForAdminResponse) GetBody() *QuerySupabaseAuthConfigsForAdminResponseBody {
	return s.Body
}

func (s *QuerySupabaseAuthConfigsForAdminResponse) SetHeaders(v map[string]*string) *QuerySupabaseAuthConfigsForAdminResponse {
	s.Headers = v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponse) SetStatusCode(v int32) *QuerySupabaseAuthConfigsForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponse) SetBody(v *QuerySupabaseAuthConfigsForAdminResponseBody) *QuerySupabaseAuthConfigsForAdminResponse {
	s.Body = v
	return s
}

func (s *QuerySupabaseAuthConfigsForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
