// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *UploadPublicKeyResponseBody) *UploadPublicKeyResponse
	GetBody() *UploadPublicKeyResponseBody
}

type UploadPublicKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *UploadPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadPublicKeyResponse) GetBody() *UploadPublicKeyResponseBody {
	return s.Body
}

func (s *UploadPublicKeyResponse) SetHeaders(v map[string]*string) *UploadPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *UploadPublicKeyResponse) SetStatusCode(v int32) *UploadPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadPublicKeyResponse) SetBody(v *UploadPublicKeyResponseBody) *UploadPublicKeyResponse {
	s.Body = v
	return s
}

func (s *UploadPublicKeyResponse) Validate() error {
	return dara.Validate(s)
}
