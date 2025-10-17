// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResourceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeClusterResourceDetailResponseBody
	GetCode() *int32
	SetData(v *DescribeClusterResourceDetailResponseBodyData) *DescribeClusterResourceDetailResponseBody
	GetData() *DescribeClusterResourceDetailResponseBodyData
	SetRequestId(v string) *DescribeClusterResourceDetailResponseBody
	GetRequestId() *string
}

type DescribeClusterResourceDetailResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried resource usage.
	Data *DescribeClusterResourceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterResourceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeClusterResourceDetailResponseBody) GetData() *DescribeClusterResourceDetailResponseBodyData {
	return s.Data
}

func (s *DescribeClusterResourceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterResourceDetailResponseBody) SetCode(v int32) *DescribeClusterResourceDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBody) SetData(v *DescribeClusterResourceDetailResponseBodyData) *DescribeClusterResourceDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeClusterResourceDetailResponseBody) SetRequestId(v string) *DescribeClusterResourceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClusterResourceDetailResponseBodyData struct {
	// The amount of reserved computing resources. Unit: AnalyticDB compute units (ACUs). Valid values: 0 to 4096. The value must be in increments of 16 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	//
	// example:
	//
	// 16ACU
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// amv-adbxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The amount of idle reserved computing resources. Unit: ACUs. Valid values: 0 to 4096. The value must be in increments of 16 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	//
	// example:
	//
	// 0ACU
	FreeComputeResource *string `json:"FreeComputeResource,omitempty" xml:"FreeComputeResource,omitempty"`
	// The resource groups.
	ResourceGroupList []*DescribeClusterResourceDetailResponseBodyDataResourceGroupList `json:"ResourceGroupList,omitempty" xml:"ResourceGroupList,omitempty" type:"Repeated"`
	// The amount of reserved storage resources. Unit: ACUs. Valid values: 0 to 2064. The value must be in increments of 24 ACUs. Each ACU is equivalent to 1 core and 4 GB memory.
	//
	// example:
	//
	// 24ACU
	StorageResource *string `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
}

func (s DescribeClusterResourceDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailResponseBodyData) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *DescribeClusterResourceDetailResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeClusterResourceDetailResponseBodyData) GetFreeComputeResource() *string {
	return s.FreeComputeResource
}

func (s *DescribeClusterResourceDetailResponseBodyData) GetResourceGroupList() []*DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	return s.ResourceGroupList
}

func (s *DescribeClusterResourceDetailResponseBodyData) GetStorageResource() *string {
	return s.StorageResource
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetComputeResource(v string) *DescribeClusterResourceDetailResponseBodyData {
	s.ComputeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetDBClusterId(v string) *DescribeClusterResourceDetailResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetFreeComputeResource(v string) *DescribeClusterResourceDetailResponseBodyData {
	s.FreeComputeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetResourceGroupList(v []*DescribeClusterResourceDetailResponseBodyDataResourceGroupList) *DescribeClusterResourceDetailResponseBodyData {
	s.ResourceGroupList = v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyData) SetStorageResource(v string) *DescribeClusterResourceDetailResponseBodyData {
	s.StorageResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyData) Validate() error {
	if s.ResourceGroupList != nil {
		for _, item := range s.ResourceGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClusterResourceDetailResponseBodyDataResourceGroupList struct {
	// A reserved parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	ClusterMode *string `json:"ClusterMode,omitempty" xml:"ClusterMode,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// 0
	ClusterSizeResource *string `json:"ClusterSizeResource,omitempty" xml:"ClusterSizeResource,omitempty"`
	// Indicates whether the preemptible instance feature is enabled for the resource group. After the preemptible instance feature is enabled, you are charged for resources at a lower unit price but the resources are probably released. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// The True value is returned only for job resource groups.
	//
	// example:
	//
	// true
	EnableSpot *bool `json:"EnableSpot,omitempty" xml:"EnableSpot,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// 0
	MaxClusterCount *int32 `json:"MaxClusterCount,omitempty" xml:"MaxClusterCount,omitempty"`
	// The maximum amount of reserved computing resources.
	//
	// example:
	//
	// 128ACU
	MaxComputeResource *string `json:"MaxComputeResource,omitempty" xml:"MaxComputeResource,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// 0
	MinClusterCount *int32 `json:"MinClusterCount,omitempty" xml:"MinClusterCount,omitempty"`
	// The minimum amount of reserved computing resources.
	//
	// example:
	//
	// 16ACU
	MinComputeResource *string `json:"MinComputeResource,omitempty" xml:"MinComputeResource,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// 17
	PoolId *int64 `json:"PoolId,omitempty" xml:"PoolId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// testadb
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The type of the resource group.
	//
	// example:
	//
	// interactive
	PoolType *string `json:"PoolType,omitempty" xml:"PoolType,omitempty"`
	// The user of the resource group.
	//
	// example:
	//
	// user1
	PoolUsers *string `json:"PoolUsers,omitempty" xml:"PoolUsers,omitempty"`
	// A reserved parameter.
	//
	// example:
	//
	// 0
	RunningClusterCount *int32 `json:"RunningClusterCount,omitempty" xml:"RunningClusterCount,omitempty"`
	// The status of the resource group. Valid values:
	//
	// 	- **running**
	//
	// 	- **deleting**
	//
	// 	- **scaling**
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeClusterResourceDetailResponseBodyDataResourceGroupList) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetClusterMode() *string {
	return s.ClusterMode
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetClusterSizeResource() *string {
	return s.ClusterSizeResource
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetEnableSpot() *bool {
	return s.EnableSpot
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetMaxClusterCount() *int32 {
	return s.MaxClusterCount
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetMaxComputeResource() *string {
	return s.MaxComputeResource
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetMinClusterCount() *int32 {
	return s.MinClusterCount
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetMinComputeResource() *string {
	return s.MinComputeResource
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetPoolId() *int64 {
	return s.PoolId
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetPoolName() *string {
	return s.PoolName
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetPoolType() *string {
	return s.PoolType
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetPoolUsers() *string {
	return s.PoolUsers
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetRunningClusterCount() *int32 {
	return s.RunningClusterCount
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) GetStatus() *string {
	return s.Status
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetClusterMode(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.ClusterMode = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetClusterSizeResource(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.ClusterSizeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetEnableSpot(v bool) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.EnableSpot = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetMaxClusterCount(v int32) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.MaxClusterCount = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetMaxComputeResource(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.MaxComputeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetMinClusterCount(v int32) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.MinClusterCount = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetMinComputeResource(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.MinComputeResource = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetPoolId(v int64) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.PoolId = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetPoolName(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.PoolName = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetPoolType(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.PoolType = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetPoolUsers(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.PoolUsers = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetRunningClusterCount(v int32) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.RunningClusterCount = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) SetStatus(v string) *DescribeClusterResourceDetailResponseBodyDataResourceGroupList {
	s.Status = &v
	return s
}

func (s *DescribeClusterResourceDetailResponseBodyDataResourceGroupList) Validate() error {
	return dara.Validate(s)
}
