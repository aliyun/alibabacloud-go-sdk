// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDeviceCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateDeviceCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateDeviceCodeResponse
	GetStatusCode() *int32
	SetBody(v *GenerateDeviceCodeResponseBody) *GenerateDeviceCodeResponse
	GetBody() *GenerateDeviceCodeResponseBody
}

type GenerateDeviceCodeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateDeviceCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateDeviceCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateDeviceCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateDeviceCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateDeviceCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateDeviceCodeResponse) GetBody() *GenerateDeviceCodeResponseBody {
	return s.Body
}

func (s *GenerateDeviceCodeResponse) SetHeaders(v map[string]*string) *GenerateDeviceCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateDeviceCodeResponse) SetStatusCode(v int32) *GenerateDeviceCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateDeviceCodeResponse) SetBody(v *GenerateDeviceCodeResponseBody) *GenerateDeviceCodeResponse {
	s.Body = v
	return s
}

func (s *GenerateDeviceCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
