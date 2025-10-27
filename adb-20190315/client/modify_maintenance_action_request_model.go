// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaintenanceActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *ModifyMaintenanceActionRequest
	GetIds() *string
	SetOwnerAccount(v string) *ModifyMaintenanceActionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyMaintenanceActionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyMaintenanceActionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyMaintenanceActionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyMaintenanceActionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyMaintenanceActionRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ModifyMaintenanceActionRequest
	GetSwitchTime() *string
}

type ModifyMaintenanceActionRequest struct {
	// The ID of the pending O\\&M event. You can specify multiple IDs to batch change the switchover time. Separate multiple IDs with commas (,).
	//
	// > - You can call the [DescribeMaintenanceAction](https://help.aliyun.com/document_detail/271738.html) operation to query the information about pending O\\&M events, including the event ID.
	//
	// > - You can change the switchover time only for pending O\\&M events. The switchover time of historical O\\&M events cannot be changed. For more information about the status of pending and historical O\\&M events, see [DescribeMaintenanceAction](https://help.aliyun.com/document_detail/271738.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 11111
	Ids          *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the pending O\\&M event occurs.
	//
	// > - You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-4690g37929XXXX
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The point in time when you want the system to perform operations on the pending O\\&M event. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-07-09T22:00:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s ModifyMaintenanceActionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaintenanceActionRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaintenanceActionRequest) GetIds() *string {
	return s.Ids
}

func (s *ModifyMaintenanceActionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyMaintenanceActionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyMaintenanceActionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyMaintenanceActionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyMaintenanceActionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyMaintenanceActionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyMaintenanceActionRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyMaintenanceActionRequest) SetIds(v string) *ModifyMaintenanceActionRequest {
	s.Ids = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetOwnerAccount(v string) *ModifyMaintenanceActionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetOwnerId(v int64) *ModifyMaintenanceActionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetRegionId(v string) *ModifyMaintenanceActionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetResourceGroupId(v string) *ModifyMaintenanceActionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetResourceOwnerAccount(v string) *ModifyMaintenanceActionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetResourceOwnerId(v int64) *ModifyMaintenanceActionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) SetSwitchTime(v string) *ModifyMaintenanceActionRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyMaintenanceActionRequest) Validate() error {
	return dara.Validate(s)
}
