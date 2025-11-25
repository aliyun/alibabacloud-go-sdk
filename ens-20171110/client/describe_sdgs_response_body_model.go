// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSDGsResponseBody
	GetRequestId() *string
	SetSDGs(v []*DescribeSDGsResponseBodySDGs) *DescribeSDGsResponseBody
	GetSDGs() []*DescribeSDGsResponseBodySDGs
}

type DescribeSDGsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3703C4AC-9396-458C-8F25-1D701334D309
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SDGs.
	SDGs []*DescribeSDGsResponseBodySDGs `json:"SDGs,omitempty" xml:"SDGs,omitempty" type:"Repeated"`
}

func (s DescribeSDGsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDGsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSDGsResponseBody) GetSDGs() []*DescribeSDGsResponseBodySDGs {
	return s.SDGs
}

func (s *DescribeSDGsResponseBody) SetRequestId(v string) *DescribeSDGsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDGsResponseBody) SetSDGs(v []*DescribeSDGsResponseBodySDGs) *DescribeSDGsResponseBody {
	s.SDGs = v
	return s
}

func (s *DescribeSDGsResponseBody) Validate() error {
	if s.SDGs != nil {
		for _, item := range s.SDGs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSDGsResponseBodySDGs struct {
	// The IDs of available edge nodes.
	AvaliableRegionIds []*DescribeSDGsResponseBodySDGsAvaliableRegionIds `json:"AvaliableRegionIds,omitempty" xml:"AvaliableRegionIds,omitempty" type:"Repeated"`
	BillingCycle       *string                                           `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// example:
	//
	// Open
	BillingType      *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	CreationDiskType *string `json:"CreationDiskType,omitempty" xml:"CreationDiskType,omitempty"`
	// The ID of the instance on which the SDG is created.
	//
	// example:
	//
	// aic-5x20dyeos****
	CreationInstanceId *string `json:"CreationInstanceId,omitempty" xml:"CreationInstanceId,omitempty"`
	// The ID of the node on which the SDG is created.
	//
	// example:
	//
	// cn-hangzhou-26
	CreationRegionId *string `json:"CreationRegionId,omitempty" xml:"CreationRegionId,omitempty"`
	// The time when the SDG was first created.
	//
	// example:
	//
	// 2023-02-27 15:07:21
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The deployment information.
	DeployedInstanceIds []*DescribeSDGsResponseBodySDGsDeployedInstanceIds `json:"DeployedInstanceIds,omitempty" xml:"DeployedInstanceIds,omitempty" type:"Repeated"`
	// The description of the SDG.
	//
	// example:
	//
	// Testing SDGs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the source SDG from which you want to create an SDG. The value of this parameter is the value of the **FromSDGId*	- parameter that you need to specify when you call the [CreateSDG](https://help.aliyun.com/document_detail/608128.html) operation.
	//
	// example:
	//
	// sdg-xxxxx
	ParentSDGId *string `json:"ParentSDGId,omitempty" xml:"ParentSDGId,omitempty"`
	// example:
	//
	// 100
	PerformanceLevel *int64 `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The ID of the SDG.
	//
	// example:
	//
	// sdg-30e1fdba7196bc****
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
	// The size of the SDG. Unit: GB.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The status of the SDG creation. Valid values:
	//
	// 	- **sdg_making**
	//
	// 	- **sdg_saving**
	//
	// 	- **failed**
	//
	// 	- **success**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the SDG was last updated.
	//
	// example:
	//
	// 2023-02-27 16:04:39
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSDGsResponseBodySDGs) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGsResponseBodySDGs) GoString() string {
	return s.String()
}

func (s *DescribeSDGsResponseBodySDGs) GetAvaliableRegionIds() []*DescribeSDGsResponseBodySDGsAvaliableRegionIds {
	return s.AvaliableRegionIds
}

func (s *DescribeSDGsResponseBodySDGs) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeSDGsResponseBodySDGs) GetBillingType() *string {
	return s.BillingType
}

func (s *DescribeSDGsResponseBodySDGs) GetCreationDiskType() *string {
	return s.CreationDiskType
}

func (s *DescribeSDGsResponseBodySDGs) GetCreationInstanceId() *string {
	return s.CreationInstanceId
}

func (s *DescribeSDGsResponseBodySDGs) GetCreationRegionId() *string {
	return s.CreationRegionId
}

func (s *DescribeSDGsResponseBodySDGs) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSDGsResponseBodySDGs) GetDeployedInstanceIds() []*DescribeSDGsResponseBodySDGsDeployedInstanceIds {
	return s.DeployedInstanceIds
}

func (s *DescribeSDGsResponseBodySDGs) GetDescription() *string {
	return s.Description
}

func (s *DescribeSDGsResponseBodySDGs) GetParentSDGId() *string {
	return s.ParentSDGId
}

func (s *DescribeSDGsResponseBodySDGs) GetPerformanceLevel() *int64 {
	return s.PerformanceLevel
}

func (s *DescribeSDGsResponseBodySDGs) GetSDGId() *string {
	return s.SDGId
}

func (s *DescribeSDGsResponseBodySDGs) GetSize() *int64 {
	return s.Size
}

func (s *DescribeSDGsResponseBodySDGs) GetStatus() *string {
	return s.Status
}

func (s *DescribeSDGsResponseBodySDGs) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSDGsResponseBodySDGs) SetAvaliableRegionIds(v []*DescribeSDGsResponseBodySDGsAvaliableRegionIds) *DescribeSDGsResponseBodySDGs {
	s.AvaliableRegionIds = v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetBillingCycle(v string) *DescribeSDGsResponseBodySDGs {
	s.BillingCycle = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetBillingType(v string) *DescribeSDGsResponseBodySDGs {
	s.BillingType = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetCreationDiskType(v string) *DescribeSDGsResponseBodySDGs {
	s.CreationDiskType = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetCreationInstanceId(v string) *DescribeSDGsResponseBodySDGs {
	s.CreationInstanceId = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetCreationRegionId(v string) *DescribeSDGsResponseBodySDGs {
	s.CreationRegionId = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetCreationTime(v string) *DescribeSDGsResponseBodySDGs {
	s.CreationTime = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetDeployedInstanceIds(v []*DescribeSDGsResponseBodySDGsDeployedInstanceIds) *DescribeSDGsResponseBodySDGs {
	s.DeployedInstanceIds = v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetDescription(v string) *DescribeSDGsResponseBodySDGs {
	s.Description = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetParentSDGId(v string) *DescribeSDGsResponseBodySDGs {
	s.ParentSDGId = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetPerformanceLevel(v int64) *DescribeSDGsResponseBodySDGs {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetSDGId(v string) *DescribeSDGsResponseBodySDGs {
	s.SDGId = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetSize(v int64) *DescribeSDGsResponseBodySDGs {
	s.Size = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetStatus(v string) *DescribeSDGsResponseBodySDGs {
	s.Status = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) SetUpdateTime(v string) *DescribeSDGsResponseBodySDGs {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGs) Validate() error {
	if s.AvaliableRegionIds != nil {
		for _, item := range s.AvaliableRegionIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeployedInstanceIds != nil {
		for _, item := range s.DeployedInstanceIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSDGsResponseBodySDGsAvaliableRegionIds struct {
	// The time when the SDG was created on the node.
	//
	// example:
	//
	// 2023-02-27 15:13:26
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-hangzhou-26
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// mock-clone_snapshot_id
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The status of the SDG on the node. Valid values:
	//
	// 	- **sdg_making**
	//
	// 	- **sdg_saving**
	//
	// 	- **sdg_copying**
	//
	// 	- **failed**
	//
	// 	- **success**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSDGsResponseBodySDGsAvaliableRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGsResponseBodySDGsAvaliableRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeSDGsResponseBodySDGsAvaliableRegionIds) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSDGsResponseBodySDGsAvaliableRegionIds) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSDGsResponseBodySDGsAvaliableRegionIds) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSDGsResponseBodySDGsAvaliableRegionIds) GetStatus() *string {
	return s.Status
}

func (s *DescribeSDGsResponseBodySDGsAvaliableRegionIds) SetCreationTime(v string) *DescribeSDGsResponseBodySDGsAvaliableRegionIds {
	s.CreationTime = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGsAvaliableRegionIds) SetRegionId(v string) *DescribeSDGsResponseBodySDGsAvaliableRegionIds {
	s.RegionId = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGsAvaliableRegionIds) SetSnapshotId(v string) *DescribeSDGsResponseBodySDGsAvaliableRegionIds {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGsAvaliableRegionIds) SetStatus(v string) *DescribeSDGsResponseBodySDGsAvaliableRegionIds {
	s.Status = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGsAvaliableRegionIds) Validate() error {
	return dara.Validate(s)
}

type DescribeSDGsResponseBodySDGsDeployedInstanceIds struct {
	// The time when the SDG was deployed on the instance.
	//
	// example:
	//
	// 2023-02-27 16:48:43
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The deployment type of the SDG. Valid values:
	//
	// 	- common: common deployment.
	//
	// 	- overlay: read/write splitting deployment.
	//
	// example:
	//
	// overlay
	DeploymentType *string `json:"DeploymentType,omitempty" xml:"DeploymentType,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// aic-5x20dyeos****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The deployment status. Valid values:
	//
	// 	- **sdg_deploying**
	//
	// 	- **failed**
	//
	// 	- **success**
	//
	// example:
	//
	// sdg_deploying
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSDGsResponseBodySDGsDeployedInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGsResponseBodySDGsDeployedInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeSDGsResponseBodySDGsDeployedInstanceIds) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSDGsResponseBodySDGsDeployedInstanceIds) GetDeploymentType() *string {
	return s.DeploymentType
}

func (s *DescribeSDGsResponseBodySDGsDeployedInstanceIds) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSDGsResponseBodySDGsDeployedInstanceIds) GetStatus() *string {
	return s.Status
}

func (s *DescribeSDGsResponseBodySDGsDeployedInstanceIds) SetCreationTime(v string) *DescribeSDGsResponseBodySDGsDeployedInstanceIds {
	s.CreationTime = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGsDeployedInstanceIds) SetDeploymentType(v string) *DescribeSDGsResponseBodySDGsDeployedInstanceIds {
	s.DeploymentType = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGsDeployedInstanceIds) SetInstanceId(v string) *DescribeSDGsResponseBodySDGsDeployedInstanceIds {
	s.InstanceId = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGsDeployedInstanceIds) SetStatus(v string) *DescribeSDGsResponseBodySDGsDeployedInstanceIds {
	s.Status = &v
	return s
}

func (s *DescribeSDGsResponseBodySDGsDeployedInstanceIds) Validate() error {
	return dara.Validate(s)
}
