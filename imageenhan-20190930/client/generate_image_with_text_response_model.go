// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageWithTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateImageWithTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateImageWithTextResponse
	GetStatusCode() *int32
	SetBody(v *GenerateImageWithTextResponseBody) *GenerateImageWithTextResponse
	GetBody() *GenerateImageWithTextResponseBody
}

type GenerateImageWithTextResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateImageWithTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateImageWithTextResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageWithTextResponse) GoString() string {
	return s.String()
}

func (s *GenerateImageWithTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateImageWithTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateImageWithTextResponse) GetBody() *GenerateImageWithTextResponseBody {
	return s.Body
}

func (s *GenerateImageWithTextResponse) SetHeaders(v map[string]*string) *GenerateImageWithTextResponse {
	s.Headers = v
	return s
}

func (s *GenerateImageWithTextResponse) SetStatusCode(v int32) *GenerateImageWithTextResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateImageWithTextResponse) SetBody(v *GenerateImageWithTextResponseBody) *GenerateImageWithTextResponse {
	s.Body = v
	return s
}

func (s *GenerateImageWithTextResponse) Validate() error {
	return dara.Validate(s)
}
