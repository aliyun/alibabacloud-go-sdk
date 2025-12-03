// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *CreateDomainExtensionRequest
	GetDomain() *string
	SetListenerPort(v int32) *CreateDomainExtensionRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *CreateDomainExtensionRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *CreateDomainExtensionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDomainExtensionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateDomainExtensionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateDomainExtensionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDomainExtensionRequest
	GetResourceOwnerId() *int64
	SetServerCertificateId(v string) *CreateDomainExtensionRequest
	GetServerCertificateId() *string
}

type CreateDomainExtensionRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// *.example1.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The frontend port that is used by the HTTPS listener of the SLB instance.
	//
	// Valid values: **1 to 65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1o94dp5i6earrxxxxxx
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Server Load Balancer (SLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the certificate used by the domain name.
	//
	// example:
	//
	// 123157xxxxxxx_166f820xxxxxx_1714763408_709981xxxx
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s CreateDomainExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainExtensionRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainExtensionRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateDomainExtensionRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *CreateDomainExtensionRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *CreateDomainExtensionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDomainExtensionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDomainExtensionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDomainExtensionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDomainExtensionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDomainExtensionRequest) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *CreateDomainExtensionRequest) SetDomain(v string) *CreateDomainExtensionRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetListenerPort(v int32) *CreateDomainExtensionRequest {
	s.ListenerPort = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetLoadBalancerId(v string) *CreateDomainExtensionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetOwnerAccount(v string) *CreateDomainExtensionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetOwnerId(v int64) *CreateDomainExtensionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetRegionId(v string) *CreateDomainExtensionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetResourceOwnerAccount(v string) *CreateDomainExtensionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetResourceOwnerId(v int64) *CreateDomainExtensionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDomainExtensionRequest) SetServerCertificateId(v string) *CreateDomainExtensionRequest {
	s.ServerCertificateId = &v
	return s
}

func (s *CreateDomainExtensionRequest) Validate() error {
	return dara.Validate(s)
}
