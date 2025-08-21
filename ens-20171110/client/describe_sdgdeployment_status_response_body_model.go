// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDGDeploymentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentStatus(v []*DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) *DescribeSDGDeploymentStatusResponseBody
	GetDeploymentStatus() []*DescribeSDGDeploymentStatusResponseBodyDeploymentStatus
	SetPageNumber(v int64) *DescribeSDGDeploymentStatusResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSDGDeploymentStatusResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeSDGDeploymentStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeSDGDeploymentStatusResponseBody
	GetTotalCount() *int64
}

type DescribeSDGDeploymentStatusResponseBody struct {
	// The list of SDG deployment information.
	DeploymentStatus []*DescribeSDGDeploymentStatusResponseBodyDeploymentStatus `json:"DeploymentStatus,omitempty" xml:"DeploymentStatus,omitempty" type:"Repeated"`
	// The page number. Pages start from page 1. Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68B85217-03B8-5141-9216-EA4D7C496B9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried deployment records.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSDGDeploymentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGDeploymentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDGDeploymentStatusResponseBody) GetDeploymentStatus() []*DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	return s.DeploymentStatus
}

func (s *DescribeSDGDeploymentStatusResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSDGDeploymentStatusResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSDGDeploymentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSDGDeploymentStatusResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSDGDeploymentStatusResponseBody) SetDeploymentStatus(v []*DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) *DescribeSDGDeploymentStatusResponseBody {
	s.DeploymentStatus = v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBody) SetPageNumber(v int64) *DescribeSDGDeploymentStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBody) SetPageSize(v int64) *DescribeSDGDeploymentStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBody) SetRequestId(v string) *DescribeSDGDeploymentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBody) SetTotalCount(v int64) *DescribeSDGDeploymentStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSDGDeploymentStatusResponseBodyDeploymentStatus struct {
	BlockRwSplitSize   *int32  `json:"BlockRwSplitSize,omitempty" xml:"BlockRwSplitSize,omitempty"`
	CacheSize          *int32  `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	DiskAccessProtocol *string `json:"DiskAccessProtocol,omitempty" xml:"DiskAccessProtocol,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the AIC instance.
	//
	// example:
	//
	// aic-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The deployment type.
	//
	// Valid values:
	//
	// 	- overlay: read/write splitting.
	//
	// 	- common: common deployment.
	//
	// example:
	//
	// overlay
	MountType *string `json:"MountType,omitempty" xml:"MountType,omitempty"`
	// The deployment phase of the SDG.
	//
	// example:
	//
	// attach
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The deployment status of the SDG.
	//
	// Valid values:
	//
	// 	- sdg_deploying
	//
	// 	- success
	//
	// 	- failed
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the status was last updated.
	//
	// example:
	//
	// 2023-02-17T02:44:31Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GoString() string {
	return s.String()
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetBlockRwSplitSize() *int32 {
	return s.BlockRwSplitSize
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetCacheSize() *int32 {
	return s.CacheSize
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetDiskAccessProtocol() *string {
	return s.DiskAccessProtocol
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetMountType() *string {
	return s.MountType
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetPhase() *string {
	return s.Phase
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetBlockRwSplitSize(v int32) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.BlockRwSplitSize = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetCacheSize(v int32) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.CacheSize = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetDiskAccessProtocol(v string) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.DiskAccessProtocol = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetDiskType(v string) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.DiskType = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetInstanceId(v string) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetMountType(v string) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.MountType = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetPhase(v string) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.Phase = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetRegionId(v string) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.RegionId = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetStatus(v string) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.Status = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) SetUpdateTime(v string) *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSDGDeploymentStatusResponseBodyDeploymentStatus) Validate() error {
	return dara.Validate(s)
}
