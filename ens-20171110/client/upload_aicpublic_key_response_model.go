// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAICPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadAICPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadAICPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *UploadAICPublicKeyResponseBody) *UploadAICPublicKeyResponse
	GetBody() *UploadAICPublicKeyResponseBody
}

type UploadAICPublicKeyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadAICPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadAICPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadAICPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *UploadAICPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadAICPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadAICPublicKeyResponse) GetBody() *UploadAICPublicKeyResponseBody {
	return s.Body
}

func (s *UploadAICPublicKeyResponse) SetHeaders(v map[string]*string) *UploadAICPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *UploadAICPublicKeyResponse) SetStatusCode(v int32) *UploadAICPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadAICPublicKeyResponse) SetBody(v *UploadAICPublicKeyResponseBody) *UploadAICPublicKeyResponse {
	s.Body = v
	return s
}

func (s *UploadAICPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
