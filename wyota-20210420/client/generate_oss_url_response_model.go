// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOssUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateOssUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateOssUrlResponse
	GetStatusCode() *int32
	SetBody(v *GenerateOssUrlResponseBody) *GenerateOssUrlResponse
	GetBody() *GenerateOssUrlResponseBody
}

type GenerateOssUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateOssUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateOssUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateOssUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateOssUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateOssUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateOssUrlResponse) GetBody() *GenerateOssUrlResponseBody {
	return s.Body
}

func (s *GenerateOssUrlResponse) SetHeaders(v map[string]*string) *GenerateOssUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateOssUrlResponse) SetStatusCode(v int32) *GenerateOssUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateOssUrlResponse) SetBody(v *GenerateOssUrlResponseBody) *GenerateOssUrlResponse {
	s.Body = v
	return s
}

func (s *GenerateOssUrlResponse) Validate() error {
	return dara.Validate(s)
}
