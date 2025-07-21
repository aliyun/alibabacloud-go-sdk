// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *GetPublicKeyResponseBody) *GetPublicKeyResponse
	GetBody() *GetPublicKeyResponseBody
}

type GetPublicKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *GetPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPublicKeyResponse) GetBody() *GetPublicKeyResponseBody {
	return s.Body
}

func (s *GetPublicKeyResponse) SetHeaders(v map[string]*string) *GetPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *GetPublicKeyResponse) SetStatusCode(v int32) *GetPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPublicKeyResponse) SetBody(v *GetPublicKeyResponseBody) *GetPublicKeyResponse {
	s.Body = v
	return s
}

func (s *GetPublicKeyResponse) Validate() error {
	return dara.Validate(s)
}
