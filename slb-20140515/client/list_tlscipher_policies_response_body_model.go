// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTLSCipherPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsTruncated(v bool) *ListTLSCipherPoliciesResponseBody
	GetIsTruncated() *bool
	SetNextToken(v string) *ListTLSCipherPoliciesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTLSCipherPoliciesResponseBody
	GetRequestId() *string
	SetTLSCipherPolicies(v []*ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) *ListTLSCipherPoliciesResponseBody
	GetTLSCipherPolicies() []*ListTLSCipherPoliciesResponseBodyTLSCipherPolicies
	SetTotalCount(v int32) *ListTLSCipherPoliciesResponseBody
	GetTotalCount() *int32
}

type ListTLSCipherPoliciesResponseBody struct {
	// Indicates whether the current page is the last page. Valid values:
	//
	// 	- **true**: The current page is the last page.
	//
	// 	- **false**: The current page is not the last page.
	//
	// example:
	//
	// false
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If **NextToken*	- is empty, it indicates that no next query is to be sent.
	//
	// 	- If **NextToken*	- is not empty, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of TLS policies.
	TLSCipherPolicies []*ListTLSCipherPoliciesResponseBodyTLSCipherPolicies `json:"TLSCipherPolicies,omitempty" xml:"TLSCipherPolicies,omitempty" type:"Repeated"`
	// The total number of TLS policies returned.
	//
	// example:
	//
	// 1000
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTLSCipherPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTLSCipherPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *ListTLSCipherPoliciesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTLSCipherPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTLSCipherPoliciesResponseBody) GetTLSCipherPolicies() []*ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	return s.TLSCipherPolicies
}

func (s *ListTLSCipherPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTLSCipherPoliciesResponseBody) SetIsTruncated(v bool) *ListTLSCipherPoliciesResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBody) SetNextToken(v string) *ListTLSCipherPoliciesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBody) SetRequestId(v string) *ListTLSCipherPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBody) SetTLSCipherPolicies(v []*ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) *ListTLSCipherPoliciesResponseBody {
	s.TLSCipherPolicies = v
	return s
}

func (s *ListTLSCipherPoliciesResponseBody) SetTotalCount(v int32) *ListTLSCipherPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBody) Validate() error {
	if s.TLSCipherPolicies != nil {
		for _, item := range s.TLSCipherPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTLSCipherPoliciesResponseBodyTLSCipherPolicies struct {
	// The cipher suites supported by the TLS version.
	//
	// TLS 1.0 and TLS 1.1 support the following cipher suites:
	//
	// 	- ECDHE-ECDSA-AES128-SHA
	//
	// 	- ECDHE-ECDSA-AES256-SHA
	//
	// 	- ECDHE-RSA-AES128-SHA
	//
	// 	- ECDHE-RSA-AES256-SHA
	//
	// 	- AES128-SHA AES256-SHA
	//
	// 	- DES-CBC3-SHA
	//
	// TLS 1.2 supports the following cipher suites:
	//
	// 	- ECDHE-ECDSA-AES128-SHA
	//
	// 	- ECDHE-ECDSA-AES256-SHA
	//
	// 	- ECDHE-RSA-AES128-SHA
	//
	// 	- ECDHE-RSA-AES256-SHA
	//
	// 	- AES128-SHA AES256-SHA
	//
	// 	- DES-CBC3-SHA
	//
	// 	- ECDHE-ECDSA-AES128-GCM-SHA256
	//
	// 	- ECDHE-ECDSA-AES256-GCM-SHA384
	//
	// 	- ECDHE-ECDSA-AES128-SHA256
	//
	// 	- ECDHE-ECDSA-AES256-SHA384
	//
	// 	- ECDHE-RSA-AES128-GCM-SHA256
	//
	// 	- ECDHE-RSA-AES256-GCM-SHA384
	//
	// 	- ECDHE-RSA-AES128-SHA256
	//
	// 	- ECDHE-RSA-AES256-SHA384
	//
	// 	- AES128-GCM-SHA256
	//
	// 	- AES256-GCM-SHA384
	//
	// 	- AES128-SHA256 AES256-SHA256
	//
	// TLS 1.3 supports the following cipher suites:
	//
	// 	- TLS_AES_128_GCM_SHA256
	//
	// 	- TLS_AES_256_GCM_SHA384
	//
	// 	- TLS_CHACHA20_POLY1305_SHA256
	//
	// 	- TLS_AES_128_CCM_SHA256
	//
	// 	- TLS_AES_128_CCM_8_SHA256
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The timestamp generated when the TLS policy is created.
	//
	// example:
	//
	// 1608273800000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the TLS policy.
	//
	// example:
	//
	// tls-bp17elso1h323r****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the TLS policy.
	//
	// example:
	//
	// TLSPolicy-test****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The list of associated listeners.
	RelateListeners []*ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners `json:"RelateListeners,omitempty" xml:"RelateListeners,omitempty" type:"Repeated"`
	// The status of the TLS policy. Valid values:
	//
	// 	- **configuring**: The TLS policy is being configured.
	//
	// 	- **normal**: The TLS policy works as expected.
	//
	// example:
	//
	// normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The version of the TLS protocol.
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) GetCiphers() []*string {
	return s.Ciphers
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) GetName() *string {
	return s.Name
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) GetRelateListeners() []*ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners {
	return s.RelateListeners
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) GetStatus() *string {
	return s.Status
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) GetTLSVersions() []*string {
	return s.TLSVersions
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetCiphers(v []*string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.Ciphers = v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetCreateTime(v int64) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.CreateTime = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetInstanceId(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.InstanceId = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetName(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.Name = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetRelateListeners(v []*ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.RelateListeners = v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetStatus(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.Status = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) SetTLSVersions(v []*string) *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies {
	s.TLSVersions = v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPolicies) Validate() error {
	if s.RelateListeners != nil {
		for _, item := range s.RelateListeners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners struct {
	// The ID of the CLB instance.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The listening port. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The listening protocol. Valid values:
	//
	// 	- **TCP**
	//
	// 	- **UDP**
	//
	// 	- **HTTP**
	//
	// 	- **HTTPS**
	//
	// example:
	//
	// HTTPS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) String() string {
	return dara.Prettify(s)
}

func (s ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) GoString() string {
	return s.String()
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) GetPort() *int32 {
	return s.Port
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) GetProtocol() *string {
	return s.Protocol
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) SetLoadBalancerId(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners {
	s.LoadBalancerId = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) SetPort(v int32) *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners {
	s.Port = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) SetProtocol(v string) *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners {
	s.Protocol = &v
	return s
}

func (s *ListTLSCipherPoliciesResponseBodyTLSCipherPoliciesRelateListeners) Validate() error {
	return dara.Validate(s)
}
