// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyToRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachPolicyToRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachPolicyToRoleResponse
	GetStatusCode() *int32
	SetBody(v *AttachPolicyToRoleResponseBody) *AttachPolicyToRoleResponse
	GetBody() *AttachPolicyToRoleResponseBody
}

type AttachPolicyToRoleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachPolicyToRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachPolicyToRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyToRoleResponse) GoString() string {
	return s.String()
}

func (s *AttachPolicyToRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachPolicyToRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachPolicyToRoleResponse) GetBody() *AttachPolicyToRoleResponseBody {
	return s.Body
}

func (s *AttachPolicyToRoleResponse) SetHeaders(v map[string]*string) *AttachPolicyToRoleResponse {
	s.Headers = v
	return s
}

func (s *AttachPolicyToRoleResponse) SetStatusCode(v int32) *AttachPolicyToRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachPolicyToRoleResponse) SetBody(v *AttachPolicyToRoleResponseBody) *AttachPolicyToRoleResponse {
	s.Body = v
	return s
}

func (s *AttachPolicyToRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
