// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateSceneKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateSceneKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateSceneKeyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateSceneKeyResponseBody) *GenerateSceneKeyResponse
	GetBody() *GenerateSceneKeyResponseBody
}

type GenerateSceneKeyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateSceneKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateSceneKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateSceneKeyResponse) GoString() string {
	return s.String()
}

func (s *GenerateSceneKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateSceneKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateSceneKeyResponse) GetBody() *GenerateSceneKeyResponseBody {
	return s.Body
}

func (s *GenerateSceneKeyResponse) SetHeaders(v map[string]*string) *GenerateSceneKeyResponse {
	s.Headers = v
	return s
}

func (s *GenerateSceneKeyResponse) SetStatusCode(v int32) *GenerateSceneKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateSceneKeyResponse) SetBody(v *GenerateSceneKeyResponseBody) *GenerateSceneKeyResponse {
	s.Body = v
	return s
}

func (s *GenerateSceneKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
