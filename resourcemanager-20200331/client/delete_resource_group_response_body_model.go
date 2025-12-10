// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroup(v *DeleteResourceGroupResponseBodyResourceGroup) *DeleteResourceGroupResponseBody
	GetResourceGroup() *DeleteResourceGroupResponseBodyResourceGroup
}

type DeleteResourceGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1C488B66-B819-4D14-8711-C4EAAA13AC01
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource group.
	ResourceGroup *DeleteResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
}

func (s DeleteResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceGroupResponseBody) GetResourceGroup() *DeleteResourceGroupResponseBodyResourceGroup {
	return s.ResourceGroup
}

func (s *DeleteResourceGroupResponseBody) SetRequestId(v string) *DeleteResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceGroupResponseBody) SetResourceGroup(v *DeleteResourceGroupResponseBodyResourceGroup) *DeleteResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

func (s *DeleteResourceGroupResponseBody) Validate() error {
	if s.ResourceGroup != nil {
		if err := s.ResourceGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteResourceGroupResponseBodyResourceGroup struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	//
	// example:
	//
	// 123456789****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2015-01-23T12:33:18Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The display name of the resource group.
	//
	// example:
	//
	// my-project
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-9gLOoK****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The unique identifier of the resource group.
	//
	// example:
	//
	// my-project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the resource group in all regions.
	RegionStatuses *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" type:"Struct"`
	// The status of the resource group. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// 	- PendingDelete: The resource group is waiting to be deleted.
	//
	// example:
	//
	// PendingDelete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteResourceGroupResponseBodyResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) GetId() *string {
	return s.Id
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) GetName() *string {
	return s.Name
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) GetRegionStatuses() *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses {
	return s.RegionStatuses
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) GetStatus() *string {
	return s.Status
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetAccountId(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.AccountId = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetCreateDate(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetDisplayName(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetId(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetName(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetRegionStatuses(v *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) *DeleteResourceGroupResponseBodyResourceGroup {
	s.RegionStatuses = v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) SetStatus(v string) *DeleteResourceGroupResponseBodyResourceGroup {
	s.Status = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroup) Validate() error {
	if s.RegionStatuses != nil {
		if err := s.RegionStatuses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteResourceGroupResponseBodyResourceGroupRegionStatuses struct {
	RegionStatus []*DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" type:"Repeated"`
}

func (s DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) GetRegionStatus() []*DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	return s.RegionStatus
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) SetRegionStatus(v []*DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses {
	s.RegionStatus = v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatuses) Validate() error {
	if s.RegionStatus != nil {
		for _, item := range s.RegionStatus {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus struct {
	// The region ID.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the resource group. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// 	- PendingDelete: The resource group is waiting to be deleted.
	//
	// 	- Deleting: The resource group is being deleted.
	//
	// example:
	//
	// PendingDelete
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GetStatus() *string {
	return s.Status
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetRegionId(v string) *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.RegionId = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetStatus(v string) *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.Status = &v
	return s
}

func (s *DeleteResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) Validate() error {
	return dara.Validate(s)
}
