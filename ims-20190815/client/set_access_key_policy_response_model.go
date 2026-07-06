// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccessKeyPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetAccessKeyPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetAccessKeyPolicyResponse
	GetStatusCode() *int32
	SetBody(v *SetAccessKeyPolicyResponseBody) *SetAccessKeyPolicyResponse
	GetBody() *SetAccessKeyPolicyResponseBody
}

type SetAccessKeyPolicyResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetAccessKeyPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetAccessKeyPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s SetAccessKeyPolicyResponse) GoString() string {
	return s.String()
}

func (s *SetAccessKeyPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetAccessKeyPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetAccessKeyPolicyResponse) GetBody() *SetAccessKeyPolicyResponseBody {
	return s.Body
}

func (s *SetAccessKeyPolicyResponse) SetHeaders(v map[string]*string) *SetAccessKeyPolicyResponse {
	s.Headers = v
	return s
}

func (s *SetAccessKeyPolicyResponse) SetStatusCode(v int32) *SetAccessKeyPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetAccessKeyPolicyResponse) SetBody(v *SetAccessKeyPolicyResponseBody) *SetAccessKeyPolicyResponse {
	s.Body = v
	return s
}

func (s *SetAccessKeyPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
