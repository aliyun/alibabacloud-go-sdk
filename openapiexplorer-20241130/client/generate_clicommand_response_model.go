// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCLICommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateCLICommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateCLICommandResponse
	GetStatusCode() *int32
	SetBody(v *GenerateCLICommandResponseBody) *GenerateCLICommandResponse
	GetBody() *GenerateCLICommandResponseBody
}

type GenerateCLICommandResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateCLICommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateCLICommandResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateCLICommandResponse) GoString() string {
	return s.String()
}

func (s *GenerateCLICommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateCLICommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateCLICommandResponse) GetBody() *GenerateCLICommandResponseBody {
	return s.Body
}

func (s *GenerateCLICommandResponse) SetHeaders(v map[string]*string) *GenerateCLICommandResponse {
	s.Headers = v
	return s
}

func (s *GenerateCLICommandResponse) SetStatusCode(v int32) *GenerateCLICommandResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateCLICommandResponse) SetBody(v *GenerateCLICommandResponseBody) *GenerateCLICommandResponse {
	s.Body = v
	return s
}

func (s *GenerateCLICommandResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
