// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateImageTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateImageTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateImageTaskResponse
	GetStatusCode() *int32
	SetBody(v *GenerateImageTaskResponseBody) *GenerateImageTaskResponse
	GetBody() *GenerateImageTaskResponseBody
}

type GenerateImageTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateImageTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateImageTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateImageTaskResponse) GoString() string {
	return s.String()
}

func (s *GenerateImageTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateImageTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateImageTaskResponse) GetBody() *GenerateImageTaskResponseBody {
	return s.Body
}

func (s *GenerateImageTaskResponse) SetHeaders(v map[string]*string) *GenerateImageTaskResponse {
	s.Headers = v
	return s
}

func (s *GenerateImageTaskResponse) SetStatusCode(v int32) *GenerateImageTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateImageTaskResponse) SetBody(v *GenerateImageTaskResponseBody) *GenerateImageTaskResponse {
	s.Body = v
	return s
}

func (s *GenerateImageTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
