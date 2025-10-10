// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemSecurityPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListSystemSecurityPoliciesResponseBody
	GetRequestId() *string
	SetSecurityPolicies(v []*ListSystemSecurityPoliciesResponseBodySecurityPolicies) *ListSystemSecurityPoliciesResponseBody
	GetSecurityPolicies() []*ListSystemSecurityPoliciesResponseBodySecurityPolicies
}

type ListSystemSecurityPoliciesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 593B0448-D13E-4C56-AC0D-FDF0FDE0E9A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The security policies.
	SecurityPolicies []*ListSystemSecurityPoliciesResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
}

func (s ListSystemSecurityPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemSecurityPoliciesResponseBody) GetSecurityPolicies() []*ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	return s.SecurityPolicies
}

func (s *ListSystemSecurityPoliciesResponseBody) SetRequestId(v string) *ListSystemSecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetSecurityPolicies(v []*ListSystemSecurityPoliciesResponseBodySecurityPolicies) *ListSystemSecurityPoliciesResponseBody {
	s.SecurityPolicies = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSystemSecurityPoliciesResponseBodySecurityPolicies struct {
	// The supported cipher suite.
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The ID of the security policy.
	//
	// example:
	//
	// spy-n0kn923****
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The supported TLS protocol versions.
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s ListSystemSecurityPoliciesResponseBodySecurityPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponseBodySecurityPolicies) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) GetCiphers() []*string {
	return s.Ciphers
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) GetSecurityPolicyId() *string {
	return s.SecurityPolicyId
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) GetTLSVersions() []*string {
	return s.TLSVersions
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetCiphers(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.Ciphers = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetTLSVersions(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.TLSVersions = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) Validate() error {
	return dara.Validate(s)
}
