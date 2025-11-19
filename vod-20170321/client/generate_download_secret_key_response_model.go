// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDownloadSecretKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDownloadSecretKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDownloadSecretKeyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDownloadSecretKeyResponseBody) *GenerateDownloadSecretKeyResponse
	GetBody() *GenerateDownloadSecretKeyResponseBody
}

type GenerateDownloadSecretKeyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDownloadSecretKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDownloadSecretKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDownloadSecretKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateDownloadSecretKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDownloadSecretKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDownloadSecretKeyResponse) GetBody() *GenerateDownloadSecretKeyResponseBody {
	return s.Body
}

func (s *GenerateDownloadSecretKeyResponse) SetHeaders(v map[string]*string) *GenerateDownloadSecretKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateDownloadSecretKeyResponse) SetStatusCode(v int32) *GenerateDownloadSecretKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDownloadSecretKeyResponse) SetBody(v *GenerateDownloadSecretKeyResponseBody) *GenerateDownloadSecretKeyResponse {
	s.Body = v
	return s
}

func (s *GenerateDownloadSecretKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
