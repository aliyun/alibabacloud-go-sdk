// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateRamPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateRamPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateRamPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateRamPolicyResponseBody) *GenerateRamPolicyResponse
	GetBody() *GenerateRamPolicyResponseBody
}

type GenerateRamPolicyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateRamPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateRamPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateRamPolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateRamPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateRamPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateRamPolicyResponse) GetBody() *GenerateRamPolicyResponseBody {
	return s.Body
}

func (s *GenerateRamPolicyResponse) SetHeaders(v map[string]*string) *GenerateRamPolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateRamPolicyResponse) SetStatusCode(v int32) *GenerateRamPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateRamPolicyResponse) SetBody(v *GenerateRamPolicyResponseBody) *GenerateRamPolicyResponse {
	s.Body = v
	return s
}

func (s *GenerateRamPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
