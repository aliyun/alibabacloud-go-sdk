// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseConfigsForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySupabaseConfigsForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySupabaseConfigsForAdminResponse
	GetStatusCode() *int32
	SetBody(v *QuerySupabaseConfigsForAdminResponseBody) *QuerySupabaseConfigsForAdminResponse
	GetBody() *QuerySupabaseConfigsForAdminResponseBody
}

type QuerySupabaseConfigsForAdminResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySupabaseConfigsForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySupabaseConfigsForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseConfigsForAdminResponse) GoString() string {
	return s.String()
}

func (s *QuerySupabaseConfigsForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySupabaseConfigsForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySupabaseConfigsForAdminResponse) GetBody() *QuerySupabaseConfigsForAdminResponseBody {
	return s.Body
}

func (s *QuerySupabaseConfigsForAdminResponse) SetHeaders(v map[string]*string) *QuerySupabaseConfigsForAdminResponse {
	s.Headers = v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponse) SetStatusCode(v int32) *QuerySupabaseConfigsForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponse) SetBody(v *QuerySupabaseConfigsForAdminResponseBody) *QuerySupabaseConfigsForAdminResponse {
	s.Body = v
	return s
}

func (s *QuerySupabaseConfigsForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
