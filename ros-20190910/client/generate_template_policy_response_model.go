// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTemplatePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateTemplatePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateTemplatePolicyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateTemplatePolicyResponseBody) *GenerateTemplatePolicyResponse
	GetBody() *GenerateTemplatePolicyResponseBody
}

type GenerateTemplatePolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateTemplatePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateTemplatePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateTemplatePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateTemplatePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateTemplatePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateTemplatePolicyResponse) GetBody() *GenerateTemplatePolicyResponseBody {
	return s.Body
}

func (s *GenerateTemplatePolicyResponse) SetHeaders(v map[string]*string) *GenerateTemplatePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateTemplatePolicyResponse) SetStatusCode(v int32) *GenerateTemplatePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateTemplatePolicyResponse) SetBody(v *GenerateTemplatePolicyResponseBody) *GenerateTemplatePolicyResponse {
	s.Body = v
	return s
}

func (s *GenerateTemplatePolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
