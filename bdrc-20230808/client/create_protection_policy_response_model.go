// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtectionPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProtectionPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProtectionPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateProtectionPolicyResponseBody) *CreateProtectionPolicyResponse
	GetBody() *CreateProtectionPolicyResponseBody
}

type CreateProtectionPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProtectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProtectionPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProtectionPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateProtectionPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProtectionPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProtectionPolicyResponse) GetBody() *CreateProtectionPolicyResponseBody {
	return s.Body
}

func (s *CreateProtectionPolicyResponse) SetHeaders(v map[string]*string) *CreateProtectionPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateProtectionPolicyResponse) SetStatusCode(v int32) *CreateProtectionPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProtectionPolicyResponse) SetBody(v *CreateProtectionPolicyResponseBody) *CreateProtectionPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateProtectionPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
