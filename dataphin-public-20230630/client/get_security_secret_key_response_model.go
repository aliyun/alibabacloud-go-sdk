// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecuritySecretKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecuritySecretKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecuritySecretKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetSecuritySecretKeyResponseBody) *GetSecuritySecretKeyResponse
	GetBody() *GetSecuritySecretKeyResponseBody
}

type GetSecuritySecretKeyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecuritySecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecuritySecretKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecuritySecretKeyResponse) GoString() string {
	return s.String()
}

func (s *GetSecuritySecretKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecuritySecretKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecuritySecretKeyResponse) GetBody() *GetSecuritySecretKeyResponseBody {
	return s.Body
}

func (s *GetSecuritySecretKeyResponse) SetHeaders(v map[string]*string) *GetSecuritySecretKeyResponse {
	s.Headers = v
	return s
}

func (s *GetSecuritySecretKeyResponse) SetStatusCode(v int32) *GetSecuritySecretKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecuritySecretKeyResponse) SetBody(v *GetSecuritySecretKeyResponseBody) *GetSecuritySecretKeyResponse {
	s.Body = v
	return s
}

func (s *GetSecuritySecretKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
