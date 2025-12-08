// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCartoonizedImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateCartoonizedImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateCartoonizedImageResponse
	GetStatusCode() *int32
	SetBody(v *GenerateCartoonizedImageResponseBody) *GenerateCartoonizedImageResponse
	GetBody() *GenerateCartoonizedImageResponseBody
}

type GenerateCartoonizedImageResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCartoonizedImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCartoonizedImageResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateCartoonizedImageResponse) GoString() string {
	return s.String()
}

func (s *GenerateCartoonizedImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateCartoonizedImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateCartoonizedImageResponse) GetBody() *GenerateCartoonizedImageResponseBody {
	return s.Body
}

func (s *GenerateCartoonizedImageResponse) SetHeaders(v map[string]*string) *GenerateCartoonizedImageResponse {
	s.Headers = v
	return s
}

func (s *GenerateCartoonizedImageResponse) SetStatusCode(v int32) *GenerateCartoonizedImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCartoonizedImageResponse) SetBody(v *GenerateCartoonizedImageResponseBody) *GenerateCartoonizedImageResponse {
	s.Body = v
	return s
}

func (s *GenerateCartoonizedImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
