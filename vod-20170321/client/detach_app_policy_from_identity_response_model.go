// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachAppPolicyFromIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachAppPolicyFromIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachAppPolicyFromIdentityResponse
	GetStatusCode() *int32
	SetBody(v *DetachAppPolicyFromIdentityResponseBody) *DetachAppPolicyFromIdentityResponse
	GetBody() *DetachAppPolicyFromIdentityResponseBody
}

type DetachAppPolicyFromIdentityResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachAppPolicyFromIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachAppPolicyFromIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachAppPolicyFromIdentityResponse) GoString() string {
	return s.String()
}

func (s *DetachAppPolicyFromIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachAppPolicyFromIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachAppPolicyFromIdentityResponse) GetBody() *DetachAppPolicyFromIdentityResponseBody {
	return s.Body
}

func (s *DetachAppPolicyFromIdentityResponse) SetHeaders(v map[string]*string) *DetachAppPolicyFromIdentityResponse {
	s.Headers = v
	return s
}

func (s *DetachAppPolicyFromIdentityResponse) SetStatusCode(v int32) *DetachAppPolicyFromIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachAppPolicyFromIdentityResponse) SetBody(v *DetachAppPolicyFromIdentityResponseBody) *DetachAppPolicyFromIdentityResponse {
	s.Body = v
	return s
}

func (s *DetachAppPolicyFromIdentityResponse) Validate() error {
	return dara.Validate(s)
}
