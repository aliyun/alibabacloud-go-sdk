// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAppPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAppPolicyResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAppPolicyResponseBody) *UpdateAppPolicyResponse
	GetBody() *UpdateAppPolicyResponseBody
}

type UpdateAppPolicyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAppPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAppPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppPolicyResponse) GoString() string {
	return s.String()
}

func (s *UpdateAppPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAppPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAppPolicyResponse) GetBody() *UpdateAppPolicyResponseBody {
	return s.Body
}

func (s *UpdateAppPolicyResponse) SetHeaders(v map[string]*string) *UpdateAppPolicyResponse {
	s.Headers = v
	return s
}

func (s *UpdateAppPolicyResponse) SetStatusCode(v int32) *UpdateAppPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAppPolicyResponse) SetBody(v *UpdateAppPolicyResponseBody) *UpdateAppPolicyResponse {
	s.Body = v
	return s
}

func (s *UpdateAppPolicyResponse) Validate() error {
	return dara.Validate(s)
}
