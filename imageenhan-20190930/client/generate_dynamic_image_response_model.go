// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDynamicImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDynamicImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDynamicImageResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDynamicImageResponseBody) *GenerateDynamicImageResponse
	GetBody() *GenerateDynamicImageResponseBody
}

type GenerateDynamicImageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDynamicImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDynamicImageResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDynamicImageResponse) GoString() string {
	return s.String()
}

func (s *GenerateDynamicImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDynamicImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDynamicImageResponse) GetBody() *GenerateDynamicImageResponseBody {
	return s.Body
}

func (s *GenerateDynamicImageResponse) SetHeaders(v map[string]*string) *GenerateDynamicImageResponse {
	s.Headers = v
	return s
}

func (s *GenerateDynamicImageResponse) SetStatusCode(v int32) *GenerateDynamicImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDynamicImageResponse) SetBody(v *GenerateDynamicImageResponseBody) *GenerateDynamicImageResponse {
	s.Body = v
	return s
}

func (s *GenerateDynamicImageResponse) Validate() error {
	return dara.Validate(s)
}
