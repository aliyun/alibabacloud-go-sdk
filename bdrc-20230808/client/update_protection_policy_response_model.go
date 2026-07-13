// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProtectionPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateProtectionPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateProtectionPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateProtectionPolicyResponseBody) *UpdateProtectionPolicyResponse
	GetBody() *UpdateProtectionPolicyResponseBody
}

type UpdateProtectionPolicyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateProtectionPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateProtectionPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateProtectionPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateProtectionPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateProtectionPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateProtectionPolicyResponse) GetBody() *UpdateProtectionPolicyResponseBody {
	return s.Body
}

func (s *UpdateProtectionPolicyResponse) SetHeaders(v map[string]*string) *UpdateProtectionPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateProtectionPolicyResponse) SetStatusCode(v int32) *UpdateProtectionPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateProtectionPolicyResponse) SetBody(v *UpdateProtectionPolicyResponseBody) *UpdateProtectionPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateProtectionPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
