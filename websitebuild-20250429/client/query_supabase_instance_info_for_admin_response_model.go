// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseInstanceInfoForAdminResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySupabaseInstanceInfoForAdminResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySupabaseInstanceInfoForAdminResponse
	GetStatusCode() *int32
	SetBody(v *QuerySupabaseInstanceInfoForAdminResponseBody) *QuerySupabaseInstanceInfoForAdminResponse
	GetBody() *QuerySupabaseInstanceInfoForAdminResponseBody
}

type QuerySupabaseInstanceInfoForAdminResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySupabaseInstanceInfoForAdminResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySupabaseInstanceInfoForAdminResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseInstanceInfoForAdminResponse) GoString() string {
	return s.String()
}

func (s *QuerySupabaseInstanceInfoForAdminResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySupabaseInstanceInfoForAdminResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySupabaseInstanceInfoForAdminResponse) GetBody() *QuerySupabaseInstanceInfoForAdminResponseBody {
	return s.Body
}

func (s *QuerySupabaseInstanceInfoForAdminResponse) SetHeaders(v map[string]*string) *QuerySupabaseInstanceInfoForAdminResponse {
	s.Headers = v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponse) SetStatusCode(v int32) *QuerySupabaseInstanceInfoForAdminResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponse) SetBody(v *QuerySupabaseInstanceInfoForAdminResponseBody) *QuerySupabaseInstanceInfoForAdminResponse {
	s.Body = v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
