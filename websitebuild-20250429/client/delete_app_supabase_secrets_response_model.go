// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppSupabaseSecretsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppSupabaseSecretsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppSupabaseSecretsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppSupabaseSecretsResponseBody) *DeleteAppSupabaseSecretsResponse
	GetBody() *DeleteAppSupabaseSecretsResponseBody
}

type DeleteAppSupabaseSecretsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppSupabaseSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppSupabaseSecretsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppSupabaseSecretsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppSupabaseSecretsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppSupabaseSecretsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppSupabaseSecretsResponse) GetBody() *DeleteAppSupabaseSecretsResponseBody {
	return s.Body
}

func (s *DeleteAppSupabaseSecretsResponse) SetHeaders(v map[string]*string) *DeleteAppSupabaseSecretsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppSupabaseSecretsResponse) SetStatusCode(v int32) *DeleteAppSupabaseSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppSupabaseSecretsResponse) SetBody(v *DeleteAppSupabaseSecretsResponseBody) *DeleteAppSupabaseSecretsResponse {
	s.Body = v
	return s
}

func (s *DeleteAppSupabaseSecretsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
