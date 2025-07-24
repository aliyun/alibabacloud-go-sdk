// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateServicePolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateServicePolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateServicePolicyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateServicePolicyResponseBody) *GenerateServicePolicyResponse
	GetBody() *GenerateServicePolicyResponseBody
}

type GenerateServicePolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateServicePolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateServicePolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateServicePolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateServicePolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateServicePolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateServicePolicyResponse) GetBody() *GenerateServicePolicyResponseBody {
	return s.Body
}

func (s *GenerateServicePolicyResponse) SetHeaders(v map[string]*string) *GenerateServicePolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateServicePolicyResponse) SetStatusCode(v int32) *GenerateServicePolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateServicePolicyResponse) SetBody(v *GenerateServicePolicyResponseBody) *GenerateServicePolicyResponse {
	s.Body = v
	return s
}

func (s *GenerateServicePolicyResponse) Validate() error {
	return dara.Validate(s)
}
