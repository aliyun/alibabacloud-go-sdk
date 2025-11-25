// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTaskCodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateTaskCodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateTaskCodesResponse
	GetStatusCode() *int32
	SetBody(v *GenerateTaskCodesResponseBody) *GenerateTaskCodesResponse
	GetBody() *GenerateTaskCodesResponseBody
}

type GenerateTaskCodesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTaskCodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTaskCodesResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateTaskCodesResponse) GoString() string {
	return s.String()
}

func (s *GenerateTaskCodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateTaskCodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateTaskCodesResponse) GetBody() *GenerateTaskCodesResponseBody {
	return s.Body
}

func (s *GenerateTaskCodesResponse) SetHeaders(v map[string]*string) *GenerateTaskCodesResponse {
	s.Headers = v
	return s
}

func (s *GenerateTaskCodesResponse) SetStatusCode(v int32) *GenerateTaskCodesResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTaskCodesResponse) SetBody(v *GenerateTaskCodesResponseBody) *GenerateTaskCodesResponse {
	s.Body = v
	return s
}

func (s *GenerateTaskCodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
