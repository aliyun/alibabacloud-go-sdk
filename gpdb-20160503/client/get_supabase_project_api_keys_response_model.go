// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSupabaseProjectApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSupabaseProjectApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *GetSupabaseProjectApiKeysResponseBody) *GetSupabaseProjectApiKeysResponse
	GetBody() *GetSupabaseProjectApiKeysResponseBody
}

type GetSupabaseProjectApiKeysResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSupabaseProjectApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSupabaseProjectApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectApiKeysResponse) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSupabaseProjectApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSupabaseProjectApiKeysResponse) GetBody() *GetSupabaseProjectApiKeysResponseBody {
	return s.Body
}

func (s *GetSupabaseProjectApiKeysResponse) SetHeaders(v map[string]*string) *GetSupabaseProjectApiKeysResponse {
	s.Headers = v
	return s
}

func (s *GetSupabaseProjectApiKeysResponse) SetStatusCode(v int32) *GetSupabaseProjectApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSupabaseProjectApiKeysResponse) SetBody(v *GetSupabaseProjectApiKeysResponseBody) *GetSupabaseProjectApiKeysResponse {
	s.Body = v
	return s
}

func (s *GetSupabaseProjectApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
