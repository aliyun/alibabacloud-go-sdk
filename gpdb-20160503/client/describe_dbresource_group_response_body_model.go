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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried resource group information.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItem struct {
	// The configurations of the resource group.
	//
	// >
	//
	// 	- CpuRateLimit: the percentage of CPU resources that are available for the resource group. Unit: %.
	//
	// 	- MemoryLimit: the percentage of memory resources that are available for the resource group. Unit: %.
	//
	// 	- MemorySharedQuota: the percentage of memory resources shared among transactions that are submitted to the resource group. Unit: %. Default value: 80.
	//
	// 	- MemorySpillRatio: the memory spill ratio for memory-intensive transactions. When the memory that is used by memory-intensive transactions reaches this value, data is spilled to disks. Unit: %. Default value: 0.
	//
	// 	- Concurrency: the maximum number of concurrent transactions or parallel queries that are allowed for a resource group. Default value: 20.
	//
	// example:
	//
	// {"CpuRateLimit":"10","MemoryLimit":"12","MemorySharedQuota":"20","MemorySpillRatio":"75","Concurrency":"3"}
	ResourceGroupConfig *string `json:"ResourceGroupConfig,omitempty" xml:"ResourceGroupConfig,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// testgroup
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The roles.
	RoleList *DescribeDBResourceGroupResponseBodyResourceGroupItemsResourceGroupItemRoleList `json:"RoleList,omitempty" xml:"RoleList,omitempty" type:"Struct"`
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
	return dara.Validate(s)
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
