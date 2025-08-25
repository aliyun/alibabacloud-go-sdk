// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageWithTextAndImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateImageWithTextAndImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateImageWithTextAndImageResponse
	GetStatusCode() *int32
	SetBody(v *GenerateImageWithTextAndImageResponseBody) *GenerateImageWithTextAndImageResponse
	GetBody() *GenerateImageWithTextAndImageResponseBody
}

type GenerateImageWithTextAndImageResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateImageWithTextAndImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateImageWithTextAndImageResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageWithTextAndImageResponse) GoString() string {
	return s.String()
}

func (s *GenerateImageWithTextAndImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateImageWithTextAndImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateImageWithTextAndImageResponse) GetBody() *GenerateImageWithTextAndImageResponseBody {
	return s.Body
}

func (s *GenerateImageWithTextAndImageResponse) SetHeaders(v map[string]*string) *GenerateImageWithTextAndImageResponse {
	s.Headers = v
	return s
}

func (s *GenerateImageWithTextAndImageResponse) SetStatusCode(v int32) *GenerateImageWithTextAndImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateImageWithTextAndImageResponse) SetBody(v *GenerateImageWithTextAndImageResponseBody) *GenerateImageWithTextAndImageResponse {
	s.Body = v
	return s
}

func (s *GenerateImageWithTextAndImageResponse) Validate() error {
	return dara.Validate(s)
}
