// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroup(v *CreateResourceGroupResponseBodyResourceGroup) *CreateResourceGroupResponseBody
	GetResourceGroup() *CreateResourceGroupResponseBodyResourceGroup
}

type CreateResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource group.
	ResourceGroup *CreateResourceGroupResponseBodyResourceGroup `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty" type:"Struct"`
}

func (s CreateResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateResourceGroupResponseBody) GetResourceGroup() *CreateResourceGroupResponseBodyResourceGroup {
	return s.ResourceGroup
}

func (s *CreateResourceGroupResponseBody) SetRequestId(v string) *CreateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceGroupResponseBody) SetResourceGroup(v *CreateResourceGroupResponseBodyResourceGroup) *CreateResourceGroupResponseBody {
	s.ResourceGroup = v
	return s
}

func (s *CreateResourceGroupResponseBody) Validate() error {
	if s.ResourceGroup != nil {
		if err := s.ResourceGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateResourceGroupResponseBodyResourceGroup struct {
	// The ID of the Alibaba Cloud account to which the resource group belongs.
	//
	// example:
	//
	// 151266687691****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The time when the resource group was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-05T14:39:13+08:00
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
	RegionStatuses *CreateResourceGroupResponseBodyResourceGroupRegionStatuses `json:"RegionStatuses,omitempty" xml:"RegionStatuses,omitempty" type:"Struct"`
	// The status of the resource group. Valid values:
	//
	// 	- Creating: The resource group is being created.
	//
	// 	- OK: The resource group is created.
	//
	// example:
	//
	// Creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateResourceGroupResponseBodyResourceGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupResponseBodyResourceGroup) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBodyResourceGroup) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateResourceGroupResponseBodyResourceGroup) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateResourceGroupResponseBodyResourceGroup) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateResourceGroupResponseBodyResourceGroup) GetId() *string {
	return s.Id
}

func (s *CreateResourceGroupResponseBodyResourceGroup) GetName() *string {
	return s.Name
}

func (s *CreateResourceGroupResponseBodyResourceGroup) GetRegionStatuses() *CreateResourceGroupResponseBodyResourceGroupRegionStatuses {
	return s.RegionStatuses
}

func (s *CreateResourceGroupResponseBodyResourceGroup) GetStatus() *string {
	return s.Status
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetAccountId(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.AccountId = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetCreateDate(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.CreateDate = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetDisplayName(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.DisplayName = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetId(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.Id = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetName(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.Name = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetRegionStatuses(v *CreateResourceGroupResponseBodyResourceGroupRegionStatuses) *CreateResourceGroupResponseBodyResourceGroup {
	s.RegionStatuses = v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) SetStatus(v string) *CreateResourceGroupResponseBodyResourceGroup {
	s.Status = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroup) Validate() error {
	if s.RegionStatuses != nil {
		if err := s.RegionStatuses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateResourceGroupResponseBodyResourceGroupRegionStatuses struct {
	RegionStatus []*CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty" type:"Repeated"`
}

func (s CreateResourceGroupResponseBodyResourceGroupRegionStatuses) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupResponseBodyResourceGroupRegionStatuses) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatuses) GetRegionStatus() []*CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	return s.RegionStatus
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatuses) SetRegionStatus(v []*CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) *CreateResourceGroupResponseBodyResourceGroupRegionStatuses {
	s.RegionStatus = v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatuses) Validate() error {
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

type CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status   *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GoString() string {
	return s.String()
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) GetStatus() *string {
	return s.Status
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetRegionId(v string) *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.RegionId = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) SetStatus(v string) *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus {
	s.Status = &v
	return s
}

func (s *CreateResourceGroupResponseBodyResourceGroupRegionStatusesRegionStatus) Validate() error {
	return dara.Validate(s)
}
