// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePrivateAccessPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePrivateAccessPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreatePrivateAccessPolicyResponseBody) *CreatePrivateAccessPolicyResponse
	GetBody() *CreatePrivateAccessPolicyResponseBody
}

type CreatePrivateAccessPolicyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePrivateAccessPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivateAccessPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePrivateAccessPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePrivateAccessPolicyResponse) GetBody() *CreatePrivateAccessPolicyResponseBody {
	return s.Body
}

func (s *CreatePrivateAccessPolicyResponse) SetHeaders(v map[string]*string) *CreatePrivateAccessPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreatePrivateAccessPolicyResponse) SetStatusCode(v int32) *CreatePrivateAccessPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePrivateAccessPolicyResponse) SetBody(v *CreatePrivateAccessPolicyResponseBody) *CreatePrivateAccessPolicyResponse {
	s.Body = v
	return s
}

func (s *CreatePrivateAccessPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
