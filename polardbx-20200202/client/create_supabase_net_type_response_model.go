// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSupabaseNetTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSupabaseNetTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSupabaseNetTypeResponse
	GetStatusCode() *int32
	SetBody(v *CreateSupabaseNetTypeResponseBody) *CreateSupabaseNetTypeResponse
	GetBody() *CreateSupabaseNetTypeResponseBody
}

type CreateSupabaseNetTypeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSupabaseNetTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSupabaseNetTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSupabaseNetTypeResponse) GoString() string {
	return s.String()
}

func (s *CreateSupabaseNetTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSupabaseNetTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSupabaseNetTypeResponse) GetBody() *CreateSupabaseNetTypeResponseBody {
	return s.Body
}

func (s *CreateSupabaseNetTypeResponse) SetHeaders(v map[string]*string) *CreateSupabaseNetTypeResponse {
	s.Headers = v
	return s
}

func (s *CreateSupabaseNetTypeResponse) SetStatusCode(v int32) *CreateSupabaseNetTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSupabaseNetTypeResponse) SetBody(v *CreateSupabaseNetTypeResponseBody) *CreateSupabaseNetTypeResponse {
	s.Body = v
	return s
}

func (s *CreateSupabaseNetTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
