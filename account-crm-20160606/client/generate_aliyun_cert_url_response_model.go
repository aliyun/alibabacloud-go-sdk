// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAliyunCertUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAliyunCertUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAliyunCertUrlResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAliyunCertUrlResponseBody) *GenerateAliyunCertUrlResponse
	GetBody() *GenerateAliyunCertUrlResponseBody
}

type GenerateAliyunCertUrlResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAliyunCertUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAliyunCertUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAliyunCertUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateAliyunCertUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAliyunCertUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAliyunCertUrlResponse) GetBody() *GenerateAliyunCertUrlResponseBody {
	return s.Body
}

func (s *GenerateAliyunCertUrlResponse) SetHeaders(v map[string]*string) *GenerateAliyunCertUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateAliyunCertUrlResponse) SetStatusCode(v int32) *GenerateAliyunCertUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAliyunCertUrlResponse) SetBody(v *GenerateAliyunCertUrlResponseBody) *GenerateAliyunCertUrlResponse {
	s.Body = v
	return s
}

func (s *GenerateAliyunCertUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
