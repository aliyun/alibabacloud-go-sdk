// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGeneratePresignedUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GeneratePresignedUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GeneratePresignedUrlResponse
	GetStatusCode() *int32
	SetBody(v *GeneratePresignedUrlResponseBody) *GeneratePresignedUrlResponse
	GetBody() *GeneratePresignedUrlResponseBody
}

type GeneratePresignedUrlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GeneratePresignedUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GeneratePresignedUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GeneratePresignedUrlResponse) GoString() string {
	return s.String()
}

func (s *GeneratePresignedUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GeneratePresignedUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GeneratePresignedUrlResponse) GetBody() *GeneratePresignedUrlResponseBody {
	return s.Body
}

func (s *GeneratePresignedUrlResponse) SetHeaders(v map[string]*string) *GeneratePresignedUrlResponse {
	s.Headers = v
	return s
}

func (s *GeneratePresignedUrlResponse) SetStatusCode(v int32) *GeneratePresignedUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GeneratePresignedUrlResponse) SetBody(v *GeneratePresignedUrlResponseBody) *GeneratePresignedUrlResponse {
	s.Body = v
	return s
}

func (s *GeneratePresignedUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
