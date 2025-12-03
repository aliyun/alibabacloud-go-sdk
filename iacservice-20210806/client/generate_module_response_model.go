// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateModuleResponse
	GetStatusCode() *int32
	SetBody(v *GenerateModuleResponseBody) *GenerateModuleResponse
	GetBody() *GenerateModuleResponseBody
}

type GenerateModuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateModuleResponse) GoString() string {
	return s.String()
}

func (s *GenerateModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateModuleResponse) GetBody() *GenerateModuleResponseBody {
	return s.Body
}

func (s *GenerateModuleResponse) SetHeaders(v map[string]*string) *GenerateModuleResponse {
	s.Headers = v
	return s
}

func (s *GenerateModuleResponse) SetStatusCode(v int32) *GenerateModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateModuleResponse) SetBody(v *GenerateModuleResponseBody) *GenerateModuleResponse {
	s.Body = v
	return s
}

func (s *GenerateModuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
