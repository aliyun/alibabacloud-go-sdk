// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GenerateUploadUrlResponseBody) *GenerateUploadUrlResponse
	GetBody() *GenerateUploadUrlResponseBody
}

type GenerateUploadUrlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateUploadUrlResponse) GetBody() *GenerateUploadUrlResponseBody {
	return s.Body
}

func (s *GenerateUploadUrlResponse) SetHeaders(v map[string]*string) *GenerateUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadUrlResponse) SetStatusCode(v int32) *GenerateUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadUrlResponse) SetBody(v *GenerateUploadUrlResponseBody) *GenerateUploadUrlResponse {
	s.Body = v
	return s
}

func (s *GenerateUploadUrlResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
