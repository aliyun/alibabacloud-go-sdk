// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeSDGResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSDGResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSDGResponseBody
	GetRequestId() *string
	SetSDGs(v []*DescribeSDGResponseBodySDGs) *DescribeSDGResponseBody
	GetSDGs() []*DescribeSDGResponseBodySDGs
	SetTotalCount(v int32) *DescribeSDGResponseBody
	GetTotalCount() *int32
}

type DescribeSDGResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F3B261DD-3858-4D3C-877D-303ADF374600
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the SDGs.
	SDGs []*DescribeSDGResponseBodySDGs `json:"SDGs,omitempty" xml:"SDGs,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 49
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDGResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSDGResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSDGResponseBody) GetSDGs() []*DescribeSDGResponseBodySDGs {
	return s.SDGs
}

func (s *DescribeSDGResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSDGResponseBody) SetPageNumber(v int32) *DescribeSDGResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSDGResponseBody) SetPageSize(v int32) *DescribeSDGResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSDGResponseBody) SetRequestId(v string) *DescribeSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDGResponseBody) SetSDGs(v []*DescribeSDGResponseBodySDGs) *DescribeSDGResponseBody {
	s.SDGs = v
	return s
}

func (s *DescribeSDGResponseBody) SetTotalCount(v int32) *DescribeSDGResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSDGResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSDGResponseBodySDGs struct {
	// SDGs that have snapshots.
	AvaliableRegionIds []*DescribeSDGResponseBodySDGsAvaliableRegionIds `json:"AvaliableRegionIds,omitempty" xml:"AvaliableRegionIds,omitempty" type:"Repeated"`
	BillingCycle       *string                                          `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	CreationDiskType   *string                                          `json:"CreationDiskType,omitempty" xml:"CreationDiskType,omitempty"`
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
	// The time when the first SDG in the node was created.
	//
	// example:
	//
	// 2023-02-27 15:07:21
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
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
	ParentSDGId      *string `json:"ParentSDGId,omitempty" xml:"ParentSDGId,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The preload information.
	PreloadInfos []*DescribeSDGResponseBodySDGsPreloadInfos `json:"PreloadInfos,omitempty" xml:"PreloadInfos,omitempty" type:"Repeated"`
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

func (s DescribeSDGResponseBodySDGs) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGResponseBodySDGs) GoString() string {
	return s.String()
}

func (s *DescribeSDGResponseBodySDGs) GetAvaliableRegionIds() []*DescribeSDGResponseBodySDGsAvaliableRegionIds {
	return s.AvaliableRegionIds
}

func (s *DescribeSDGResponseBodySDGs) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeSDGResponseBodySDGs) GetCreationDiskType() *string {
	return s.CreationDiskType
}

func (s *DescribeSDGResponseBodySDGs) GetCreationInstanceId() *string {
	return s.CreationInstanceId
}

func (s *DescribeSDGResponseBodySDGs) GetCreationRegionId() *string {
	return s.CreationRegionId
}

func (s *DescribeSDGResponseBodySDGs) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSDGResponseBodySDGs) GetDescription() *string {
	return s.Description
}

func (s *DescribeSDGResponseBodySDGs) GetParentSDGId() *string {
	return s.ParentSDGId
}

func (s *DescribeSDGResponseBodySDGs) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeSDGResponseBodySDGs) GetPreloadInfos() []*DescribeSDGResponseBodySDGsPreloadInfos {
	return s.PreloadInfos
}

func (s *DescribeSDGResponseBodySDGs) GetSDGId() *string {
	return s.SDGId
}

func (s *DescribeSDGResponseBodySDGs) GetSize() *int64 {
	return s.Size
}

func (s *DescribeSDGResponseBodySDGs) GetStatus() *string {
	return s.Status
}

func (s *DescribeSDGResponseBodySDGs) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSDGResponseBodySDGs) SetAvaliableRegionIds(v []*DescribeSDGResponseBodySDGsAvaliableRegionIds) *DescribeSDGResponseBodySDGs {
	s.AvaliableRegionIds = v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetBillingCycle(v string) *DescribeSDGResponseBodySDGs {
	s.BillingCycle = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetCreationDiskType(v string) *DescribeSDGResponseBodySDGs {
	s.CreationDiskType = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetCreationInstanceId(v string) *DescribeSDGResponseBodySDGs {
	s.CreationInstanceId = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetCreationRegionId(v string) *DescribeSDGResponseBodySDGs {
	s.CreationRegionId = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetCreationTime(v string) *DescribeSDGResponseBodySDGs {
	s.CreationTime = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetDescription(v string) *DescribeSDGResponseBodySDGs {
	s.Description = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetParentSDGId(v string) *DescribeSDGResponseBodySDGs {
	s.ParentSDGId = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetPerformanceLevel(v string) *DescribeSDGResponseBodySDGs {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetPreloadInfos(v []*DescribeSDGResponseBodySDGsPreloadInfos) *DescribeSDGResponseBodySDGs {
	s.PreloadInfos = v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetSDGId(v string) *DescribeSDGResponseBodySDGs {
	s.SDGId = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetSize(v int64) *DescribeSDGResponseBodySDGs {
	s.Size = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetStatus(v string) *DescribeSDGResponseBodySDGs {
	s.Status = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) SetUpdateTime(v string) *DescribeSDGResponseBodySDGs {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSDGResponseBodySDGs) Validate() error {
	return dara.Validate(s)
}

type DescribeSDGResponseBodySDGsAvaliableRegionIds struct {
	// The time when the SDG was created on the node.
	//
	// example:
	//
	// 2023-02-27 15:13:26
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-hangzhou-26
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// sp-517qu0tgznrg622he7nf4wd7n
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

func (s DescribeSDGResponseBodySDGsAvaliableRegionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGResponseBodySDGsAvaliableRegionIds) GoString() string {
	return s.String()
}

func (s *DescribeSDGResponseBodySDGsAvaliableRegionIds) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSDGResponseBodySDGsAvaliableRegionIds) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSDGResponseBodySDGsAvaliableRegionIds) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeSDGResponseBodySDGsAvaliableRegionIds) GetStatus() *string {
	return s.Status
}

func (s *DescribeSDGResponseBodySDGsAvaliableRegionIds) SetCreationTime(v string) *DescribeSDGResponseBodySDGsAvaliableRegionIds {
	s.CreationTime = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsAvaliableRegionIds) SetRegionId(v string) *DescribeSDGResponseBodySDGsAvaliableRegionIds {
	s.RegionId = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsAvaliableRegionIds) SetSnapshotId(v string) *DescribeSDGResponseBodySDGsAvaliableRegionIds {
	s.SnapshotId = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsAvaliableRegionIds) SetStatus(v string) *DescribeSDGResponseBodySDGsAvaliableRegionIds {
	s.Status = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsAvaliableRegionIds) Validate() error {
	return dara.Validate(s)
}

type DescribeSDGResponseBodySDGsPreloadInfos struct {
	// The time when the SDG was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-16T06:18:40Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DiskType     *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The namespace.
	//
	// example:
	//
	// test-20000
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The number of redundant replicas to quickly respond to shared mounts.
	//
	// example:
	//
	// 2
	RedundantNum *int32 `json:"RedundantNum,omitempty" xml:"RedundantNum,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The time when the status was last updated.
	//
	// example:
	//
	// 2021-01-22T08:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSDGResponseBodySDGsPreloadInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGResponseBodySDGsPreloadInfos) GoString() string {
	return s.String()
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) GetRedundantNum() *int32 {
	return s.RedundantNum
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) SetCreationTime(v string) *DescribeSDGResponseBodySDGsPreloadInfos {
	s.CreationTime = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) SetDiskType(v string) *DescribeSDGResponseBodySDGsPreloadInfos {
	s.DiskType = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) SetNamespace(v string) *DescribeSDGResponseBodySDGsPreloadInfos {
	s.Namespace = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) SetRedundantNum(v int32) *DescribeSDGResponseBodySDGsPreloadInfos {
	s.RedundantNum = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) SetRegionId(v string) *DescribeSDGResponseBodySDGsPreloadInfos {
	s.RegionId = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) SetUpdateTime(v string) *DescribeSDGResponseBodySDGsPreloadInfos {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSDGResponseBodySDGsPreloadInfos) Validate() error {
	return dara.Validate(s)
}
