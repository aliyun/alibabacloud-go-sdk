// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSecurityPolicyResponseBody
	GetRequestId() *string
	SetSecurityPolicyId(v string) *CreateSecurityPolicyResponseBody
	GetSecurityPolicyId() *string
}

type CreateSecurityPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security policy ID.
	//
	// example:
	//
	// scp-bp1bpn0kn9****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
}

func (s CreateSecurityPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecurityPolicyResponseBody) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *CreateSecurityPolicyResponseBody) SetRequestId(v string) *CreateSecurityPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) SetSecurityPolicyId(v string) *CreateSecurityPolicyResponseBody {
	s.SecurityPolicyId = &v
	return s
}

func (s *CreateSecurityPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
