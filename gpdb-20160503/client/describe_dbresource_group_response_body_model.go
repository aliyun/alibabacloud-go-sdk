// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDBResourceGroupResponseBody
	GetRequestId() *string
	SetResourceGroupItems(v *DescribeDBResourceGroupResponseBodyResourceGroupItems) *DescribeDBResourceGroupResponseBody
	GetResourceGroupItems() *DescribeDBResourceGroupResponseBodyResourceGroupItems
}

type DescribeDBResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId          *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupItems *DescribeDBResourceGroupResponseBodyResourceGroupItems `json:"ResourceGroupItems,omitempty" xml:"ResourceGroupItems,omitempty" type:"Struct"`
}

func (s DescribeDBResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBResourceGroupResponseBody) GetResourceGroupItems() *DescribeDBResourceGroupResponseBodyResourceGroupItems {
	return s.ResourceGroupItems
}

func (s *DescribeDBResourceGroupResponseBody) SetRequestId(v string) *DescribeDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) SetResourceGroupItems(v *DescribeDBResourceGroupResponseBodyResourceGroupItems) *DescribeDBResourceGroupResponseBody {
	s.ResourceGroupItems = v
	return s
}

func (s *DescribeDBResourceGroupResponseBody) Validate() error {
	if s.ResourceGroupItems != nil {
		if err := s.ResourceGroupItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBResourceGroupResponseBodyResourceGroupItems struct {
	ResourceGroupItem []*DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem `json:"ResourceGroupItem,omitempty" xml:"ResourceGroupItem,omitempty" type:"Repeated"`
}

func (s DescribeDBResourceGroupResponseBodyResourceGroupItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyResourceGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItems) GetResourceGroupItem() []*DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem {
	return s.ResourceGroupItem
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItems) SetResourceGroupItem(v []*DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) *DescribeDBResourceGroupResponseBodyResourceGroupItems {
	s.ResourceGroupItem = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItems) Validate() error {
	if s.ResourceGroupItem != nil {
		for _, item := range s.ResourceGroupItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem struct {
	ResourceGroupConfig *string                                                                         `json:"ResourceGroupConfig,omitempty" xml:"ResourceGroupConfig,omitempty"`
	ResourceGroupName   *string                                                                         `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	RoleList            *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Struct"`
}

func (s DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) GetResourceGroupConfig() *string {
	return s.ResourceGroupConfig
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) GetRoleList() *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList {
	return s.RoleList
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) SetResourceGroupConfig(v string) *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem {
	s.ResourceGroupConfig = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) SetResourceGroupName(v string) *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) SetRoleList(v *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList) *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem {
	s.RoleList = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem) Validate() error {
	if s.RoleList != nil {
		if err := s.RoleList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList struct {
	Role []*string `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList) GoString() string {
	return s.String()
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList) GetRole() []*string {
	return s.Role
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList) SetRole(v []*string) *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList {
	s.Role = v
	return s
}

func (s *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList) Validate() error {
	return dara.Validate(s)
}
