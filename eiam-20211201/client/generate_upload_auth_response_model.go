// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadAuthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateUploadAuthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateUploadAuthResponse
	GetStatusCode() *int32
	SetBody(v *GenerateUploadAuthResponseBody) *GenerateUploadAuthResponse
	GetBody() *GenerateUploadAuthResponseBody
}

type GenerateUploadAuthResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadAuthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadAuthResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadAuthResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadAuthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateUploadAuthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateUploadAuthResponse) GetBody() *GenerateUploadAuthResponseBody {
	return s.Body
}

func (s *GenerateUploadAuthResponse) SetHeaders(v map[string]*string) *GenerateUploadAuthResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadAuthResponse) SetStatusCode(v int32) *GenerateUploadAuthResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadAuthResponse) SetBody(v *GenerateUploadAuthResponseBody) *GenerateUploadAuthResponse {
	s.Body = v
	return s
}

func (s *GenerateUploadAuthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
