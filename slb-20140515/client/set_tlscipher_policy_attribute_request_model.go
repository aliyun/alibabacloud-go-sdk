// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetTLSCipherPolicyAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphers(v []*string) *SetTLSCipherPolicyAttributeRequest
	GetCiphers() []*string
	SetName(v string) *SetTLSCipherPolicyAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *SetTLSCipherPolicyAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetTLSCipherPolicyAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetTLSCipherPolicyAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetTLSCipherPolicyAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetTLSCipherPolicyAttributeRequest
	GetResourceOwnerId() *int64
	SetTLSCipherPolicyId(v string) *SetTLSCipherPolicyAttributeRequest
	GetTLSCipherPolicyId() *string
	SetTLSVersions(v []*string) *SetTLSCipherPolicyAttributeRequest
	GetTLSVersions() []*string
}

type SetTLSCipherPolicyAttributeRequest struct {
	// The cipher suites supported by the TLS version.
	//
	// The specified cipher suites must be supported by at least one TLS protocol version that you specify. For example, if you set the TLSVersions parameter to TLSv1.3, you must specify cipher suites that are supported by this protocol version.
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
	//
	// This parameter is required.
	//
	// example:
	//
	// DES-CBC3-SHA
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The name of the TLS policy. The name must be 2 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// tls-policy*****-test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Server Load Balancer (SLB) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the TLS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// tls-bp1lp2076qx4e******bridp
	TLSCipherPolicyId *string `json:"TLSCipherPolicyId,omitempty" xml:"TLSCipherPolicyId,omitempty"`
	// The version of the TLS protocol. Valid values: **TLSv1.0**, **TLSv1.1**, **TLSv1.2**, and **TLSv1.3**.
	//
	// This parameter is required.
	//
	// example:
	//
	// TLSv1.0
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s SetTLSCipherPolicyAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s SetTLSCipherPolicyAttributeRequest) GoString() string {
	return s.String()
}

func (s *SetTLSCipherPolicyAttributeRequest) GetCiphers() []*string {
	return s.Ciphers
}

func (s *SetTLSCipherPolicyAttributeRequest) GetName() *string {
	return s.Name
}

func (s *SetTLSCipherPolicyAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetTLSCipherPolicyAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetTLSCipherPolicyAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetTLSCipherPolicyAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetTLSCipherPolicyAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetTLSCipherPolicyAttributeRequest) GetTLSCipherPolicyId() *string {
	return s.TLSCipherPolicyId
}

func (s *SetTLSCipherPolicyAttributeRequest) GetTLSVersions() []*string {
	return s.TLSVersions
}

func (s *SetTLSCipherPolicyAttributeRequest) SetCiphers(v []*string) *SetTLSCipherPolicyAttributeRequest {
	s.Ciphers = v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetName(v string) *SetTLSCipherPolicyAttributeRequest {
	s.Name = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetOwnerAccount(v string) *SetTLSCipherPolicyAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetOwnerId(v int64) *SetTLSCipherPolicyAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetRegionId(v string) *SetTLSCipherPolicyAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetResourceOwnerAccount(v string) *SetTLSCipherPolicyAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetResourceOwnerId(v int64) *SetTLSCipherPolicyAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetTLSCipherPolicyId(v string) *SetTLSCipherPolicyAttributeRequest {
	s.TLSCipherPolicyId = &v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) SetTLSVersions(v []*string) *SetTLSCipherPolicyAttributeRequest {
	s.TLSVersions = v
	return s
}

func (s *SetTLSCipherPolicyAttributeRequest) Validate() error {
	return dara.Validate(s)
}
