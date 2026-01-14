// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSystemSecurityPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSystemSecurityPoliciesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSystemSecurityPoliciesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSystemSecurityPoliciesResponseBody
	GetRequestId() *string
	SetSecurityPolicies(v []*ListSystemSecurityPoliciesResponseBodySecurityPolicies) *ListSystemSecurityPoliciesResponseBody
	GetSecurityPolicies() []*ListSystemSecurityPoliciesResponseBodySecurityPolicies
	SetTotalCount(v int32) *ListSystemSecurityPoliciesResponseBody
	GetTotalCount() *int32
}

type ListSystemSecurityPoliciesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A9B4E54C-9CCD-4002-91A9-D38C6C209192
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of TLS security policies.
	SecurityPolicies []*ListSystemSecurityPoliciesResponseBodySecurityPolicies `json:"SecurityPolicies,omitempty" xml:"SecurityPolicies,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSystemSecurityPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSystemSecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSystemSecurityPoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSystemSecurityPoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSystemSecurityPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSystemSecurityPoliciesResponseBody) GetSecurityPolicies() []*ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	return s.SecurityPolicies
}

func (s *ListSystemSecurityPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSystemSecurityPoliciesResponseBody) SetPageNumber(v int32) *ListSystemSecurityPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetPageSize(v int32) *ListSystemSecurityPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetRequestId(v string) *ListSystemSecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetSecurityPolicies(v []*ListSystemSecurityPoliciesResponseBodySecurityPolicies) *ListSystemSecurityPoliciesResponseBody {
	s.SecurityPolicies = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) SetTotalCount(v int32) *ListSystemSecurityPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBody) Validate() error {
	if s.SecurityPolicies != nil {
		for _, item := range s.SecurityPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSystemSecurityPoliciesResponseBodySecurityPolicies struct {
	// The supported cipher suites. The value of this parameter is determined by the value of **TLSVersions**.
	//
	// The specified cipher suites must be supported by at least one value of **TLSVersions**. For example, if you set TLSVersions to **TLSv1.3**, you must specify cipher suites that are supported by **TLSv1.3**.
	//
	// 	- Valid values when TLSVersions is set to **TLSv1.0*	- or **TLSv1.1**:
	//
	//     	- ECDHE-ECDSA-AES128-SHA
	//
	//     	- ECDHE-ECDSA-AES256-SHA
	//
	//     	- ECDHE-RSA-AES128-SHA
	//
	//     	- ECDHE-RSA-AES256-SHA
	//
	//     	- AES128-SHA
	//
	//     	- AES256-SHA
	//
	//     	- DES-CBC3-SHA
	//
	// 	- Valid values when TLSVersions is set to **TLSv1.2**:
	//
	//     	- ECDHE-ECDSA-AES128-SHA
	//
	//     	- ECDHE-ECDSA-AES256-SHA
	//
	//     	- ECDHE-RSA-AES128-SHA
	//
	//     	- ECDHE-RSA-AES256-SHA
	//
	//     	- AES128-SHA
	//
	//     	- AES256-SHA
	//
	//     	- DES-CBC3-SHA
	//
	//     	- ECDHE-ECDSA-AES128-GCM-SHA256
	//
	//     	- ECDHE-ECDSA-AES256-GCM-SHA384
	//
	//     	- ECDHE-ECDSA-AES128-SHA256
	//
	//     	- ECDHE-ECDSA-AES256-SHA384
	//
	//     	- ECDHE-RSA-AES128-GCM-SHA256
	//
	//     	- ECDHE-RSA-AES256-GCM-SHA384
	//
	//     	- ECDHE-RSA-AES128-SHA256
	//
	//     	- ECDHE-RSA-AES256-SHA384
	//
	//     	- AES128-GCM-SHA256
	//
	//     	- AES256-GCM-SHA384
	//
	//     	- AES128-SHA256
	//
	//     	- AES256-SHA256
	//
	// 	- Valid values when TLSVersions is set to **TLSv1.3**:
	//
	//     	- TLS_AES_128_GCM_SHA256
	//
	//     	- TLS_AES_256_GCM_SHA384
	//
	//     	- TLS_CHACHA20_POLY1305_SHA256
	//
	//     	- TLS_AES_128_CCM_SHA256
	//
	//     	- TLS_AES_128_CCM_8_SHA256
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The ID of the TLS security policy.
	//
	// example:
	//
	// tls_cipher_policy_1_1
	SecurityPolicyId *string `json:"SecurityPolicyId,omitempty" xml:"SecurityPolicyId,omitempty"`
	// The supported TLS versions. Valid values: **TLSv1.0**, **TLSv1.1**, **TLSv1.2**, and **TLSv1.3**.
	TlsVersions []*string `json:"TlsVersions,omitempty" xml:"TlsVersions,omitempty" type:"Repeated"`
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

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) GetTlsVersions() []*string {
	return s.TlsVersions
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetCiphers(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.Ciphers = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetSecurityPolicyId(v string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.SecurityPolicyId = &v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) SetTlsVersions(v []*string) *ListSystemSecurityPoliciesResponseBodySecurityPolicies {
	s.TlsVersions = v
	return s
}

func (s *ListSystemSecurityPoliciesResponseBodySecurityPolicies) Validate() error {
	return dara.Validate(s)
}
