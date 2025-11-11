// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateUploadConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateUploadConfigResponse
	GetStatusCode() *int32
	SetBody(v *GenerateUploadConfigResponseBody) *GenerateUploadConfigResponse
	GetBody() *GenerateUploadConfigResponseBody
}

type GenerateUploadConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateUploadConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateUploadConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadConfigResponse) GoString() string {
	return s.String()
}

func (s *GenerateUploadConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateUploadConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateUploadConfigResponse) GetBody() *GenerateUploadConfigResponseBody {
	return s.Body
}

func (s *GenerateUploadConfigResponse) SetHeaders(v map[string]*string) *GenerateUploadConfigResponse {
	s.Headers = v
	return s
}

func (s *GenerateUploadConfigResponse) SetStatusCode(v int32) *GenerateUploadConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateUploadConfigResponse) SetBody(v *GenerateUploadConfigResponseBody) *GenerateUploadConfigResponse {
	s.Body = v
	return s
}

func (s *GenerateUploadConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
