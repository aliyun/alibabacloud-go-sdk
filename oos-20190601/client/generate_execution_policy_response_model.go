// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateExecutionPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateExecutionPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateExecutionPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GenerateExecutionPolicyResponseBody) *GenerateExecutionPolicyResponse
	GetBody() *GenerateExecutionPolicyResponseBody
}

type GenerateExecutionPolicyResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateExecutionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateExecutionPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateExecutionPolicyResponse) GoString() string {
	return s.String()
}

func (s *GenerateExecutionPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateExecutionPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateExecutionPolicyResponse) GetBody() *GenerateExecutionPolicyResponseBody {
	return s.Body
}

func (s *GenerateExecutionPolicyResponse) SetHeaders(v map[string]*string) *GenerateExecutionPolicyResponse {
	s.Headers = v
	return s
}

func (s *GenerateExecutionPolicyResponse) SetStatusCode(v int32) *GenerateExecutionPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateExecutionPolicyResponse) SetBody(v *GenerateExecutionPolicyResponseBody) *GenerateExecutionPolicyResponse {
	s.Body = v
	return s
}

func (s *GenerateExecutionPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
