// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateGroupImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateGroupImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateGroupImageResponse
	GetStatusCode() *int32
	SetBody(v *GenerateGroupImageResponseBody) *GenerateGroupImageResponse
	GetBody() *GenerateGroupImageResponseBody
}

type GenerateGroupImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateGroupImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateGroupImageResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateGroupImageResponse) GoString() string {
	return s.String()
}

func (s *GenerateGroupImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateGroupImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateGroupImageResponse) GetBody() *GenerateGroupImageResponseBody {
	return s.Body
}

func (s *GenerateGroupImageResponse) SetHeaders(v map[string]*string) *GenerateGroupImageResponse {
	s.Headers = v
	return s
}

func (s *GenerateGroupImageResponse) SetStatusCode(v int32) *GenerateGroupImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateGroupImageResponse) SetBody(v *GenerateGroupImageResponseBody) *GenerateGroupImageResponse {
	s.Body = v
	return s
}

func (s *GenerateGroupImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
