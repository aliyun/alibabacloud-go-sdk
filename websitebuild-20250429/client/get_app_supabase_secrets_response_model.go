// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSupabaseSecretsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppSupabaseSecretsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppSupabaseSecretsResponse
	GetStatusCode() *int32
	SetBody(v *GetAppSupabaseSecretsResponseBody) *GetAppSupabaseSecretsResponse
	GetBody() *GetAppSupabaseSecretsResponseBody
}

type GetAppSupabaseSecretsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppSupabaseSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppSupabaseSecretsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppSupabaseSecretsResponse) GoString() string {
	return s.String()
}

func (s *GetAppSupabaseSecretsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppSupabaseSecretsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppSupabaseSecretsResponse) GetBody() *GetAppSupabaseSecretsResponseBody {
	return s.Body
}

func (s *GetAppSupabaseSecretsResponse) SetHeaders(v map[string]*string) *GetAppSupabaseSecretsResponse {
	s.Headers = v
	return s
}

func (s *GetAppSupabaseSecretsResponse) SetStatusCode(v int32) *GetAppSupabaseSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppSupabaseSecretsResponse) SetBody(v *GetAppSupabaseSecretsResponseBody) *GetAppSupabaseSecretsResponse {
	s.Body = v
	return s
}

func (s *GetAppSupabaseSecretsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
