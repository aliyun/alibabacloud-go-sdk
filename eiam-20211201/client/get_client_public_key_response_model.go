// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClientPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClientPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetClientPublicKeyResponseBody) *GetClientPublicKeyResponse
	GetBody() *GetClientPublicKeyResponseBody
}

type GetClientPublicKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClientPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClientPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClientPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *GetClientPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClientPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClientPublicKeyResponse) GetBody() *GetClientPublicKeyResponseBody {
	return s.Body
}

func (s *GetClientPublicKeyResponse) SetHeaders(v map[string]*string) *GetClientPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *GetClientPublicKeyResponse) SetStatusCode(v int32) *GetClientPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClientPublicKeyResponse) SetBody(v *GetClientPublicKeyResponseBody) *GetClientPublicKeyResponse {
	s.Body = v
	return s
}

func (s *GetClientPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
