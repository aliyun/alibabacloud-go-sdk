// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmartAccessGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAclIds(v string) *DescribeSmartAccessGatewaysRequest
	GetAclIds() *string
	SetAssociatedCcnId(v string) *DescribeSmartAccessGatewaysRequest
	GetAssociatedCcnId() *string
	SetAssociatedCcnName(v string) *DescribeSmartAccessGatewaysRequest
	GetAssociatedCcnName() *string
	SetBusinessState(v string) *DescribeSmartAccessGatewaysRequest
	GetBusinessState() *string
	SetCanAssociateQos(v bool) *DescribeSmartAccessGatewaysRequest
	GetCanAssociateQos() *bool
	SetHardwareType(v string) *DescribeSmartAccessGatewaysRequest
	GetHardwareType() *string
	SetInstanceType(v string) *DescribeSmartAccessGatewaysRequest
	GetInstanceType() *string
	SetName(v string) *DescribeSmartAccessGatewaysRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeSmartAccessGatewaysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSmartAccessGatewaysRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSmartAccessGatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSmartAccessGatewaysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSmartAccessGatewaysRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSmartAccessGatewaysRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewaysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSmartAccessGatewaysRequest
	GetResourceOwnerId() *int64
	SetSerialNumber(v string) *DescribeSmartAccessGatewaysRequest
	GetSerialNumber() *string
	SetSmartAGId(v string) *DescribeSmartAccessGatewaysRequest
	GetSmartAGId() *string
	SetSmartAGIds(v []*string) *DescribeSmartAccessGatewaysRequest
	GetSmartAGIds() []*string
	SetSoftwareVersion(v string) *DescribeSmartAccessGatewaysRequest
	GetSoftwareVersion() *string
	SetStatus(v string) *DescribeSmartAccessGatewaysRequest
	GetStatus() *string
	SetUnboundAclIds(v string) *DescribeSmartAccessGatewaysRequest
	GetUnboundAclIds() *string
	SetVersionComparator(v string) *DescribeSmartAccessGatewaysRequest
	GetVersionComparator() *string
}

type DescribeSmartAccessGatewaysRequest struct {
	// The instance ID of the access control instance attached to the Smart Access Gateway instance.
	//
	// example:
	//
	// acl-xhwhyuo43l0n****
	AclIds *string `json:"AclIds,omitempty" xml:"AclIds,omitempty"`
	// The instance ID of the Cloud Connect Network (CCN) attached to the Smart Access Gateway instance.
	//
	// example:
	//
	// ccn-bxuau4ezctt****
	AssociatedCcnId *string `json:"AssociatedCcnId,omitempty" xml:"AssociatedCcnId,omitempty"`
	// The name of the associated Cloud Connect Network (CCN) instance. The name must be 2 to 100 characters in length. It must start with a letter or a Chinese character, and can contain digits, underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// ccn
	AssociatedCcnName *string `json:"AssociatedCcnName,omitempty" xml:"AssociatedCcnName,omitempty"`
	// The service status of the Smart Access Gateway instance. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **Arrearage**: Locked due to overdue payment.
	//
	// example:
	//
	// Normal
	BusinessState *string `json:"BusinessState,omitempty" xml:"BusinessState,omitempty"`
	// Specifies whether the Smart Access Gateway instance can be associated with a quality of service (QoS) policy instance. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// This parameter is used to query Smart Access Gateway instances based on whether they can be associated with a QoS policy instance.
	//
	// example:
	//
	// false
	CanAssociateQos *bool `json:"CanAssociateQos,omitempty" xml:"CanAssociateQos,omitempty"`
	// The device specification of the Smart Access Gateway device. Valid values:
	//
	// - **sag-1000**.
	//
	// - **sag-100wm**.
	//
	// example:
	//
	// sag-100wm
	HardwareType *string `json:"HardwareType,omitempty" xml:"HardwareType,omitempty"`
	// The type of the Smart Access Gateway instance. Valid values:
	//
	// - **software**: SAG app instance.
	//
	// - **hardware**: Smart Access Gateway hardware instance.
	//
	// example:
	//
	// hardware
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the Smart Access Gateway instance.
	//
	// The name must be 2 to 128 characters in length, and can contain digits, underscores (_), and hyphens (-). It must start with a letter or a Chinese character.
	//
	// example:
	//
	// nametest
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page in a paging query. Default value: **10**. Valid values: **1*	- to **50**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region to which the Smart Access Gateway instance belongs.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the Smart Access Gateway instance belongs.
	//
	// example:
	//
	// rg-acfm2iu4fnc****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The serial number of the Smart Access Gateway device.
	//
	// example:
	//
	// sag32a3****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The instance ID of the Smart Access Gateway.
	//
	// example:
	//
	// sag-6z21oj0vjjrx6s****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The list of Smart Access Gateway instance IDs.
	SmartAGIds []*string `json:"SmartAGIds,omitempty" xml:"SmartAGIds,omitempty" type:"Repeated"`
	// The software version running on the Smart Access Gateway device.
	//
	// example:
	//
	// 2.3.0.0
	SoftwareVersion *string `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`
	// The status of the Smart Access Gateway instance. Valid values:
	//
	// - **Ordered**: Ordered.
	//
	// - **Delivered**: Delivered.
	//
	// - **Received**: Received.
	//
	// - **Returning**: Returning.
	//
	// - **Active**: Active.
	//
	// - **Init**: Initializing.
	//
	// - **Offline**: Offline.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the access control instance. This parameter is used to query Smart Access Gateway instances that are not associated with the specified access control instance. Separate multiple access control instance IDs with commas (,).
	//
	// example:
	//
	// acl-sjfbgngj****
	UnboundAclIds *string `json:"UnboundAclIds,omitempty" xml:"UnboundAclIds,omitempty"`
	// The version filter. Valid values:
	//
	// - **greater**: Greater than the current version.
	//
	// - **less**: Less than the current version.
	//
	// - **equals**: Equal to the current version.
	//
	// example:
	//
	// equals
	VersionComparator *string `json:"VersionComparator,omitempty" xml:"VersionComparator,omitempty"`
}

func (s DescribeSmartAccessGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmartAccessGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmartAccessGatewaysRequest) GetAclIds() *string {
	return s.AclIds
}

func (s *DescribeSmartAccessGatewaysRequest) GetAssociatedCcnId() *string {
	return s.AssociatedCcnId
}

func (s *DescribeSmartAccessGatewaysRequest) GetAssociatedCcnName() *string {
	return s.AssociatedCcnName
}

func (s *DescribeSmartAccessGatewaysRequest) GetBusinessState() *string {
	return s.BusinessState
}

func (s *DescribeSmartAccessGatewaysRequest) GetCanAssociateQos() *bool {
	return s.CanAssociateQos
}

func (s *DescribeSmartAccessGatewaysRequest) GetHardwareType() *string {
	return s.HardwareType
}

func (s *DescribeSmartAccessGatewaysRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeSmartAccessGatewaysRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSmartAccessGatewaysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSmartAccessGatewaysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSmartAccessGatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSmartAccessGatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSmartAccessGatewaysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSmartAccessGatewaysRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSmartAccessGatewaysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSmartAccessGatewaysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSmartAccessGatewaysRequest) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeSmartAccessGatewaysRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSmartAccessGatewaysRequest) GetSmartAGIds() []*string {
	return s.SmartAGIds
}

func (s *DescribeSmartAccessGatewaysRequest) GetSoftwareVersion() *string {
	return s.SoftwareVersion
}

func (s *DescribeSmartAccessGatewaysRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeSmartAccessGatewaysRequest) GetUnboundAclIds() *string {
	return s.UnboundAclIds
}

func (s *DescribeSmartAccessGatewaysRequest) GetVersionComparator() *string {
	return s.VersionComparator
}

func (s *DescribeSmartAccessGatewaysRequest) SetAclIds(v string) *DescribeSmartAccessGatewaysRequest {
	s.AclIds = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetAssociatedCcnId(v string) *DescribeSmartAccessGatewaysRequest {
	s.AssociatedCcnId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetAssociatedCcnName(v string) *DescribeSmartAccessGatewaysRequest {
	s.AssociatedCcnName = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetBusinessState(v string) *DescribeSmartAccessGatewaysRequest {
	s.BusinessState = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetCanAssociateQos(v bool) *DescribeSmartAccessGatewaysRequest {
	s.CanAssociateQos = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetHardwareType(v string) *DescribeSmartAccessGatewaysRequest {
	s.HardwareType = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetInstanceType(v string) *DescribeSmartAccessGatewaysRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetName(v string) *DescribeSmartAccessGatewaysRequest {
	s.Name = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetOwnerAccount(v string) *DescribeSmartAccessGatewaysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetOwnerId(v int64) *DescribeSmartAccessGatewaysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetPageNumber(v int32) *DescribeSmartAccessGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetPageSize(v int32) *DescribeSmartAccessGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetRegionId(v string) *DescribeSmartAccessGatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetResourceGroupId(v string) *DescribeSmartAccessGatewaysRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetResourceOwnerAccount(v string) *DescribeSmartAccessGatewaysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetResourceOwnerId(v int64) *DescribeSmartAccessGatewaysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetSerialNumber(v string) *DescribeSmartAccessGatewaysRequest {
	s.SerialNumber = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetSmartAGId(v string) *DescribeSmartAccessGatewaysRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetSmartAGIds(v []*string) *DescribeSmartAccessGatewaysRequest {
	s.SmartAGIds = v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetSoftwareVersion(v string) *DescribeSmartAccessGatewaysRequest {
	s.SoftwareVersion = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetStatus(v string) *DescribeSmartAccessGatewaysRequest {
	s.Status = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetUnboundAclIds(v string) *DescribeSmartAccessGatewaysRequest {
	s.UnboundAclIds = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) SetVersionComparator(v string) *DescribeSmartAccessGatewaysRequest {
	s.VersionComparator = &v
	return s
}

func (s *DescribeSmartAccessGatewaysRequest) Validate() error {
	return dara.Validate(s)
}
