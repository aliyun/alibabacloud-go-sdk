// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseUpdateVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupabaseUpdateVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupabaseUpdateVersionResponse
	GetStatusCode() *int32
	SetBody(v *GetSupabaseUpdateVersionResponseBody) *GetSupabaseUpdateVersionResponse
	GetBody() *GetSupabaseUpdateVersionResponseBody
}

type GetSupabaseUpdateVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupabaseUpdateVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupabaseUpdateVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseUpdateVersionResponse) GoString() string {
	return s.String()
}

func (s *GetSupabaseUpdateVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupabaseUpdateVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupabaseUpdateVersionResponse) GetBody() *GetSupabaseUpdateVersionResponseBody {
	return s.Body
}

func (s *GetSupabaseUpdateVersionResponse) SetHeaders(v map[string]*string) *GetSupabaseUpdateVersionResponse {
	s.Headers = v
	return s
}

func (s *GetSupabaseUpdateVersionResponse) SetStatusCode(v int32) *GetSupabaseUpdateVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupabaseUpdateVersionResponse) SetBody(v *GetSupabaseUpdateVersionResponseBody) *GetSupabaseUpdateVersionResponse {
	s.Body = v
	return s
}

func (s *GetSupabaseUpdateVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
