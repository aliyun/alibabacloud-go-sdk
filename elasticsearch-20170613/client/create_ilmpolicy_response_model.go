// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateILMPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateILMPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateILMPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateILMPolicyResponseBody) *CreateILMPolicyResponse
	GetBody() *CreateILMPolicyResponseBody
}

type CreateILMPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateILMPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateILMPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateILMPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateILMPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateILMPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateILMPolicyResponse) GetBody() *CreateILMPolicyResponseBody {
	return s.Body
}

func (s *CreateILMPolicyResponse) SetHeaders(v map[string]*string) *CreateILMPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateILMPolicyResponse) SetStatusCode(v int32) *CreateILMPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateILMPolicyResponse) SetBody(v *CreateILMPolicyResponseBody) *CreateILMPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateILMPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
