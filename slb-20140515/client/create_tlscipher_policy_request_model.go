// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTLSCipherPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphers(v []*string) *CreateTLSCipherPolicyRequest
	GetCiphers() []*string
	SetName(v string) *CreateTLSCipherPolicyRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateTLSCipherPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTLSCipherPolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTLSCipherPolicyRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTLSCipherPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTLSCipherPolicyRequest
	GetResourceOwnerId() *int64
	SetTLSVersions(v []*string) *CreateTLSCipherPolicyRequest
	GetTLSVersions() []*string
}

type CreateTLSCipherPolicyRequest struct {
	// The supported cipher suites, which are determined by the TLS protocol version. You can specify at most 32 cipher suites.
	//
	// TLS 1.0 and TLS 1.1 support the following cipher suites:
	//
	// 	- **ECDHE-ECDSA-AES128-SHA**
	//
	// 	- **ECDHE-ECDSA-AES256-SHA**
	//
	// 	- **ECDHE-RSA-AES128-SHA**
	//
	// 	- **ECDHE-RSA-AES256-SHA**
	//
	// 	- **AES128-SHA**
	//
	// 	- **AES256-SHA**
	//
	// 	- **DES-CBC3-SHA**
	//
	// TLS 1.2 supports the following cipher suites:
	//
	// 	- **ECDHE-ECDSA-AES128-SHA**
	//
	// 	- **ECDHE-ECDSA-AES256-SHA**
	//
	// 	- **ECDHE-RSA-AES128-SHA**
	//
	// 	- **ECDHE-RSA-AES256-SHA**
	//
	// 	- **AES128-SHA**
	//
	// 	- **AES256-SHA**
	//
	// 	- **DES-CBC3-SHA**
	//
	// 	- **ECDHE-ECDSA-AES128-GCM-SHA256**
	//
	// 	- **ECDHE-ECDSA-AES256-GCM-SHA384**
	//
	// 	- **ECDHE-ECDSA-AES128-SHA256**
	//
	// 	- **ECDHE-ECDSA-AES256-SHA384**
	//
	// 	- **ECDHE-RSA-AES128-GCM-SHA256**
	//
	// 	- **ECDHE-RSA-AES256-GCM-SHA384**
	//
	// 	- **ECDHE-RSA-AES128-SHA256**
	//
	// 	- **ECDHE-RSA-AES256-SHA384**
	//
	// 	- **AES128-GCM-SHA256**
	//
	// 	- **AES256-GCM-SHA384**
	//
	// 	- **AES128-SHA256**
	//
	// 	- **AES256-SHA256**
	//
	// TLS 1.3 supports the following cipher suites:
	//
	// 	- **TLS_AES_128_GCM_SHA256**
	//
	// 	- **TLS_AES_256_GCM_SHA384**
	//
	// 	- **TLS_CHACHA20_POLY1305_SHA256**
	//
	// 	- **TLS_AES_128_CCM_SHA256**
	//
	// 	- **TLS_AES_128_CCM_8_SHA256**
	//
	// This parameter is required.
	//
	// example:
	//
	// AES256-SHA256
	Ciphers []*string `json:"Ciphers,omitempty" xml:"Ciphers,omitempty" type:"Repeated"`
	// The name of the TLS policy. The name must be 2 to 200 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// TLSPolicy-test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Server Load Balancer (SLB) instance is created.
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
	// The version of the TLS protocol. Valid values: **TLSv1.0**, **TLSv1.1**, **TLSv1.2**, and **TLSv1.3**. You can specify at most four TLS versions.
	//
	// This parameter is required.
	//
	// example:
	//
	// TLSv1.0
	TLSVersions []*string `json:"TLSVersions,omitempty" xml:"TLSVersions,omitempty" type:"Repeated"`
}

func (s CreateTLSCipherPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTLSCipherPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateTLSCipherPolicyRequest) GetCiphers() []*string {
	return s.Ciphers
}

func (s *CreateTLSCipherPolicyRequest) GetName() *string {
	return s.Name
}

func (s *CreateTLSCipherPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTLSCipherPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTLSCipherPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTLSCipherPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTLSCipherPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTLSCipherPolicyRequest) GetTLSVersions() []*string {
	return s.TLSVersions
}

func (s *CreateTLSCipherPolicyRequest) SetCiphers(v []*string) *CreateTLSCipherPolicyRequest {
	s.Ciphers = v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetName(v string) *CreateTLSCipherPolicyRequest {
	s.Name = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetOwnerAccount(v string) *CreateTLSCipherPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetOwnerId(v int64) *CreateTLSCipherPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetRegionId(v string) *CreateTLSCipherPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetResourceOwnerAccount(v string) *CreateTLSCipherPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetResourceOwnerId(v int64) *CreateTLSCipherPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTLSCipherPolicyRequest) SetTLSVersions(v []*string) *CreateTLSCipherPolicyRequest {
	s.TLSVersions = v
	return s
}

func (s *CreateTLSCipherPolicyRequest) Validate() error {
	return dara.Validate(s)
}
