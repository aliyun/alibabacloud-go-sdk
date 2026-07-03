// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolardbxSupabaseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolardbxSupabaseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolardbxSupabaseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolardbxSupabaseInstanceResponseBody) *CreatePolardbxSupabaseInstanceResponse
	GetBody() *CreatePolardbxSupabaseInstanceResponseBody
}

type CreatePolardbxSupabaseInstanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolardbxSupabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolardbxSupabaseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolardbxSupabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreatePolardbxSupabaseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolardbxSupabaseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolardbxSupabaseInstanceResponse) GetBody() *CreatePolardbxSupabaseInstanceResponseBody {
	return s.Body
}

func (s *CreatePolardbxSupabaseInstanceResponse) SetHeaders(v map[string]*string) *CreatePolardbxSupabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponse) SetStatusCode(v int32) *CreatePolardbxSupabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponse) SetBody(v *CreatePolardbxSupabaseInstanceResponseBody) *CreatePolardbxSupabaseInstanceResponse {
	s.Body = v
	return s
}

func (s *CreatePolardbxSupabaseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
