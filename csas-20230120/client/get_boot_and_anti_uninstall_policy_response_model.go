// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBootAndAntiUninstallPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBootAndAntiUninstallPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBootAndAntiUninstallPolicyResponse
	GetStatusCode() *int32
	SetBody(v *GetBootAndAntiUninstallPolicyResponseBody) *GetBootAndAntiUninstallPolicyResponse
	GetBody() *GetBootAndAntiUninstallPolicyResponseBody
}

type GetBootAndAntiUninstallPolicyResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBootAndAntiUninstallPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBootAndAntiUninstallPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBootAndAntiUninstallPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetBootAndAntiUninstallPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBootAndAntiUninstallPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBootAndAntiUninstallPolicyResponse) GetBody() *GetBootAndAntiUninstallPolicyResponseBody {
	return s.Body
}

func (s *GetBootAndAntiUninstallPolicyResponse) SetHeaders(v map[string]*string) *GetBootAndAntiUninstallPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponse) SetStatusCode(v int32) *GetBootAndAntiUninstallPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponse) SetBody(v *GetBootAndAntiUninstallPolicyResponseBody) *GetBootAndAntiUninstallPolicyResponse {
	s.Body = v
	return s
}

func (s *GetBootAndAntiUninstallPolicyResponse) Validate() error {
	return dara.Validate(s)
}
