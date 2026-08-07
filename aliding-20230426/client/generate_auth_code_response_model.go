// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAuthCodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAuthCodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAuthCodeResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAuthCodeResponseBody) *GenerateAuthCodeResponse
	GetBody() *GenerateAuthCodeResponseBody
}

type GenerateAuthCodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAuthCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAuthCodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAuthCodeResponse) GoString() string {
	return s.String()
}

func (s *GenerateAuthCodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAuthCodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAuthCodeResponse) GetBody() *GenerateAuthCodeResponseBody {
	return s.Body
}

func (s *GenerateAuthCodeResponse) SetHeaders(v map[string]*string) *GenerateAuthCodeResponse {
	s.Headers = v
	return s
}

func (s *GenerateAuthCodeResponse) SetStatusCode(v int32) *GenerateAuthCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAuthCodeResponse) SetBody(v *GenerateAuthCodeResponseBody) *GenerateAuthCodeResponse {
	s.Body = v
	return s
}

func (s *GenerateAuthCodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
