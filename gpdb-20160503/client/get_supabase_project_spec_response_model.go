// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectSpecResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupabaseProjectSpecResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupabaseProjectSpecResponse
	GetStatusCode() *int32
	SetBody(v *GetSupabaseProjectSpecResponseBody) *GetSupabaseProjectSpecResponse
	GetBody() *GetSupabaseProjectSpecResponseBody
}

type GetSupabaseProjectSpecResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupabaseProjectSpecResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupabaseProjectSpecResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectSpecResponse) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectSpecResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupabaseProjectSpecResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupabaseProjectSpecResponse) GetBody() *GetSupabaseProjectSpecResponseBody {
	return s.Body
}

func (s *GetSupabaseProjectSpecResponse) SetHeaders(v map[string]*string) *GetSupabaseProjectSpecResponse {
	s.Headers = v
	return s
}

func (s *GetSupabaseProjectSpecResponse) SetStatusCode(v int32) *GetSupabaseProjectSpecResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupabaseProjectSpecResponse) SetBody(v *GetSupabaseProjectSpecResponseBody) *GetSupabaseProjectSpecResponse {
	s.Body = v
	return s
}

func (s *GetSupabaseProjectSpecResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
