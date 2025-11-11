// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileUrlByKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateFileUrlByKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateFileUrlByKeyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateFileUrlByKeyResponseBody) *GenerateFileUrlByKeyResponse
	GetBody() *GenerateFileUrlByKeyResponseBody
}

type GenerateFileUrlByKeyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateFileUrlByKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateFileUrlByKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileUrlByKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateFileUrlByKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateFileUrlByKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateFileUrlByKeyResponse) GetBody() *GenerateFileUrlByKeyResponseBody {
	return s.Body
}

func (s *GenerateFileUrlByKeyResponse) SetHeaders(v map[string]*string) *GenerateFileUrlByKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateFileUrlByKeyResponse) SetStatusCode(v int32) *GenerateFileUrlByKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateFileUrlByKeyResponse) SetBody(v *GenerateFileUrlByKeyResponseBody) *GenerateFileUrlByKeyResponse {
	s.Body = v
	return s
}

func (s *GenerateFileUrlByKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
