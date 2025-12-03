// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTLSCipherPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTLSCipherPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTLSCipherPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateTLSCipherPolicyResponseBody) *CreateTLSCipherPolicyResponse
	GetBody() *CreateTLSCipherPolicyResponseBody
}

type CreateTLSCipherPolicyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTLSCipherPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTLSCipherPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTLSCipherPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateTLSCipherPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTLSCipherPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTLSCipherPolicyResponse) GetBody() *CreateTLSCipherPolicyResponseBody {
	return s.Body
}

func (s *CreateTLSCipherPolicyResponse) SetHeaders(v map[string]*string) *CreateTLSCipherPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateTLSCipherPolicyResponse) SetStatusCode(v int32) *CreateTLSCipherPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTLSCipherPolicyResponse) SetBody(v *CreateTLSCipherPolicyResponseBody) *CreateTLSCipherPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateTLSCipherPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
