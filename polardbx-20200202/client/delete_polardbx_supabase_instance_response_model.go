// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolardbxSupabaseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolardbxSupabaseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolardbxSupabaseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolardbxSupabaseInstanceResponseBody) *DeletePolardbxSupabaseInstanceResponse
	GetBody() *DeletePolardbxSupabaseInstanceResponseBody
}

type DeletePolardbxSupabaseInstanceResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolardbxSupabaseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolardbxSupabaseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolardbxSupabaseInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeletePolardbxSupabaseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolardbxSupabaseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolardbxSupabaseInstanceResponse) GetBody() *DeletePolardbxSupabaseInstanceResponseBody {
	return s.Body
}

func (s *DeletePolardbxSupabaseInstanceResponse) SetHeaders(v map[string]*string) *DeletePolardbxSupabaseInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponse) SetStatusCode(v int32) *DeletePolardbxSupabaseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponse) SetBody(v *DeletePolardbxSupabaseInstanceResponseBody) *DeletePolardbxSupabaseInstanceResponse {
	s.Body = v
	return s
}

func (s *DeletePolardbxSupabaseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
