// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSupabaseNetTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSupabaseNetTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSupabaseNetTypeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSupabaseNetTypeResponseBody) *DeleteSupabaseNetTypeResponse
	GetBody() *DeleteSupabaseNetTypeResponseBody
}

type DeleteSupabaseNetTypeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSupabaseNetTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSupabaseNetTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSupabaseNetTypeResponse) GoString() string {
	return s.String()
}

func (s *DeleteSupabaseNetTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSupabaseNetTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSupabaseNetTypeResponse) GetBody() *DeleteSupabaseNetTypeResponseBody {
	return s.Body
}

func (s *DeleteSupabaseNetTypeResponse) SetHeaders(v map[string]*string) *DeleteSupabaseNetTypeResponse {
	s.Headers = v
	return s
}

func (s *DeleteSupabaseNetTypeResponse) SetStatusCode(v int32) *DeleteSupabaseNetTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSupabaseNetTypeResponse) SetBody(v *DeleteSupabaseNetTypeResponseBody) *DeleteSupabaseNetTypeResponse {
	s.Body = v
	return s
}

func (s *DeleteSupabaseNetTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
