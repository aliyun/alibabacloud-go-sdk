// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileUploadParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateFileUploadParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateFileUploadParamsResponse
	GetStatusCode() *int32
	SetBody(v *GenerateFileUploadParamsResponseBody) *GenerateFileUploadParamsResponse
	GetBody() *GenerateFileUploadParamsResponseBody
}

type GenerateFileUploadParamsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateFileUploadParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateFileUploadParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileUploadParamsResponse) GoString() string {
	return s.String()
}

func (s *GenerateFileUploadParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateFileUploadParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateFileUploadParamsResponse) GetBody() *GenerateFileUploadParamsResponseBody {
	return s.Body
}

func (s *GenerateFileUploadParamsResponse) SetHeaders(v map[string]*string) *GenerateFileUploadParamsResponse {
	s.Headers = v
	return s
}

func (s *GenerateFileUploadParamsResponse) SetStatusCode(v int32) *GenerateFileUploadParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateFileUploadParamsResponse) SetBody(v *GenerateFileUploadParamsResponseBody) *GenerateFileUploadParamsResponse {
	s.Body = v
	return s
}

func (s *GenerateFileUploadParamsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
