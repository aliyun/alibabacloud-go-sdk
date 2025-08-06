// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAppPolicyToIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachAppPolicyToIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachAppPolicyToIdentityResponse
	GetStatusCode() *int32
	SetBody(v *AttachAppPolicyToIdentityResponseBody) *AttachAppPolicyToIdentityResponse
	GetBody() *AttachAppPolicyToIdentityResponseBody
}

type AttachAppPolicyToIdentityResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachAppPolicyToIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachAppPolicyToIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachAppPolicyToIdentityResponse) GoString() string {
	return s.String()
}

func (s *AttachAppPolicyToIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachAppPolicyToIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachAppPolicyToIdentityResponse) GetBody() *AttachAppPolicyToIdentityResponseBody {
	return s.Body
}

func (s *AttachAppPolicyToIdentityResponse) SetHeaders(v map[string]*string) *AttachAppPolicyToIdentityResponse {
	s.Headers = v
	return s
}

func (s *AttachAppPolicyToIdentityResponse) SetStatusCode(v int32) *AttachAppPolicyToIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachAppPolicyToIdentityResponse) SetBody(v *AttachAppPolicyToIdentityResponseBody) *AttachAppPolicyToIdentityResponse {
	s.Body = v
	return s
}

func (s *AttachAppPolicyToIdentityResponse) Validate() error {
	return dara.Validate(s)
}
