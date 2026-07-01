// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachNetworkInterfaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *AttachNetworkInterfaceRequest
	GetInstanceId() *string
	SetNetworkCardIndex(v int32) *AttachNetworkInterfaceRequest
	GetNetworkCardIndex() *int32
	SetNetworkInterfaceId(v string) *AttachNetworkInterfaceRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *AttachNetworkInterfaceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AttachNetworkInterfaceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AttachNetworkInterfaceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AttachNetworkInterfaceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AttachNetworkInterfaceRequest
	GetResourceOwnerId() *int64
	SetTrunkNetworkInstanceId(v string) *AttachNetworkInterfaceRequest
	GetTrunkNetworkInstanceId() *string
	SetWaitForNetworkConfigurationReady(v bool) *AttachNetworkInterfaceRequest
	GetWaitForNetworkConfigurationReady() *bool
}

type AttachNetworkInterfaceRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp16qstyvxj9gpqw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The index of the network card specified for the ENI.
	//
	// >Valid values of NetworkCardIndex depend on the instance family. If the instance type does not support network cards, you cannot specify this parameter. If the instance type supports network cards, see [Instance families](https://help.aliyun.com/document_detail/25378.html) for valid values.
	//
	// example:
	//
	// 0
	NetworkCardIndex *int32 `json:"NetworkCardIndex,omitempty" xml:"NetworkCardIndex,omitempty"`
	// The network interface controller (NIC) ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp17pdijfczax1huji****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the trunk ENI.
	//
	// >This parameter is not yet available.
	//
	// example:
	//
	// eni-f8zapqwj1v1j4ia3****
	TrunkNetworkInstanceId *string `json:"TrunkNetworkInstanceId,omitempty" xml:"TrunkNetworkInstanceId,omitempty"`
	// >This parameter is deprecated.
	//
	// example:
	//
	// null
	WaitForNetworkConfigurationReady *bool `json:"WaitForNetworkConfigurationReady,omitempty" xml:"WaitForNetworkConfigurationReady,omitempty"`
}

func (s AttachNetworkInterfaceRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachNetworkInterfaceRequest) GoString() string {
	return s.String()
}

func (s *AttachNetworkInterfaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AttachNetworkInterfaceRequest) GetNetworkCardIndex() *int32 {
	return s.NetworkCardIndex
}

func (s *AttachNetworkInterfaceRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *AttachNetworkInterfaceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AttachNetworkInterfaceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AttachNetworkInterfaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachNetworkInterfaceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AttachNetworkInterfaceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AttachNetworkInterfaceRequest) GetTrunkNetworkInstanceId() *string {
	return s.TrunkNetworkInstanceId
}

func (s *AttachNetworkInterfaceRequest) GetWaitForNetworkConfigurationReady() *bool {
	return s.WaitForNetworkConfigurationReady
}

func (s *AttachNetworkInterfaceRequest) SetInstanceId(v string) *AttachNetworkInterfaceRequest {
	s.InstanceId = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetNetworkCardIndex(v int32) *AttachNetworkInterfaceRequest {
	s.NetworkCardIndex = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetNetworkInterfaceId(v string) *AttachNetworkInterfaceRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetOwnerAccount(v string) *AttachNetworkInterfaceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetOwnerId(v int64) *AttachNetworkInterfaceRequest {
	s.OwnerId = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetRegionId(v string) *AttachNetworkInterfaceRequest {
	s.RegionId = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetResourceOwnerAccount(v string) *AttachNetworkInterfaceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetResourceOwnerId(v int64) *AttachNetworkInterfaceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetTrunkNetworkInstanceId(v string) *AttachNetworkInterfaceRequest {
	s.TrunkNetworkInstanceId = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) SetWaitForNetworkConfigurationReady(v bool) *AttachNetworkInterfaceRequest {
	s.WaitForNetworkConfigurationReady = &v
	return s
}

func (s *AttachNetworkInterfaceRequest) Validate() error {
	return dara.Validate(s)
}
