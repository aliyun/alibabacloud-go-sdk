// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBootAndAntiUninstallPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBootAndAntiUninstallPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBootAndAntiUninstallPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBootAndAntiUninstallPolicyResponseBody) *UpdateBootAndAntiUninstallPolicyResponse
	GetBody() *UpdateBootAndAntiUninstallPolicyResponseBody
}

type UpdateBootAndAntiUninstallPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBootAndAntiUninstallPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBootAndAntiUninstallPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBootAndAntiUninstallPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateBootAndAntiUninstallPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBootAndAntiUninstallPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBootAndAntiUninstallPolicyResponse) GetBody() *UpdateBootAndAntiUninstallPolicyResponseBody {
	return s.Body
}

func (s *UpdateBootAndAntiUninstallPolicyResponse) SetHeaders(v map[string]*string) *UpdateBootAndAntiUninstallPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponse) SetStatusCode(v int32) *UpdateBootAndAntiUninstallPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponse) SetBody(v *UpdateBootAndAntiUninstallPolicyResponseBody) *UpdateBootAndAntiUninstallPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateBootAndAntiUninstallPolicyResponse) Validate() error {
	return dara.Validate(s)
}
