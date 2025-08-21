// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceSDGStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentStatus(v []*DescribeInstanceSDGStatusResponseBodyDeploymentStatus) *DescribeInstanceSDGStatusResponseBody
	GetDeploymentStatus() []*DescribeInstanceSDGStatusResponseBodyDeploymentStatus
	SetPageNumber(v int64) *DescribeInstanceSDGStatusResponseBody
	GetPageNumber() *int64
	SetPageSize(v string) *DescribeInstanceSDGStatusResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeInstanceSDGStatusResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeInstanceSDGStatusResponseBody
	GetTotalCount() *string
}

type DescribeInstanceSDGStatusResponseBody struct {
	// The deployment information of the SDGs.
	DeploymentStatus []*DescribeInstanceSDGStatusResponseBodyDeploymentStatus `json:"DeploymentStatus,omitempty" xml:"DeploymentStatus,omitempty" type:"Repeated"`
	// The number of the page to return. Pages start from page 1. Default value: 1
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
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C0003E8B-B930-4F59-ADC0-0E209A9012A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried deployment records.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceSDGStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSDGStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSDGStatusResponseBody) GetDeploymentStatus() []*DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	return s.DeploymentStatus
}

func (s *DescribeInstanceSDGStatusResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInstanceSDGStatusResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeInstanceSDGStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceSDGStatusResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeInstanceSDGStatusResponseBody) SetDeploymentStatus(v []*DescribeInstanceSDGStatusResponseBodyDeploymentStatus) *DescribeInstanceSDGStatusResponseBody {
	s.DeploymentStatus = v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBody) SetPageNumber(v int64) *DescribeInstanceSDGStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBody) SetPageSize(v string) *DescribeInstanceSDGStatusResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBody) SetRequestId(v string) *DescribeInstanceSDGStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBody) SetTotalCount(v string) *DescribeInstanceSDGStatusResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceSDGStatusResponseBodyDeploymentStatus struct {
	BlockRwSplitSize   *int32  `json:"BlockRwSplitSize,omitempty" xml:"BlockRwSplitSize,omitempty"`
	CacheSize          *int32  `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	DiskAccessProtocol *string `json:"DiskAccessProtocol,omitempty" xml:"DiskAccessProtocol,omitempty"`
	DiskType           *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-guangzhou-26
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the AIC instance.
	//
	// example:
	//
	// aic-xxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The deployment type of the SDG.
	//
	// example:
	//
	// shared
	MountType *string `json:"MountType,omitempty" xml:"MountType,omitempty"`
	// Deployment Phase
	//
	// example:
	//
	// attach
	Phase *string `json:"Phase,omitempty" xml:"Phase,omitempty"`
	// The ID of the SDG.
	//
	// example:
	//
	// sdg-xxxxxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
	// The deployment status of the SDG.
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

func (s DescribeInstanceSDGStatusResponseBodyDeploymentStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GoString() string {
	return s.String()
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetBlockRwSplitSize() *int32 {
	return s.BlockRwSplitSize
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetCacheSize() *int32 {
	return s.CacheSize
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetDiskAccessProtocol() *string {
	return s.DiskAccessProtocol
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetMountType() *string {
	return s.MountType
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetPhase() *string {
	return s.Phase
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetSDGId() *string {
	return s.SDGId
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetBlockRwSplitSize(v int32) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.BlockRwSplitSize = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetCacheSize(v int32) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.CacheSize = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetDiskAccessProtocol(v string) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.DiskAccessProtocol = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetDiskType(v string) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetEnsRegionId(v string) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetInstanceId(v string) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetMountType(v string) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.MountType = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetPhase(v string) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.Phase = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetSDGId(v string) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.SDGId = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetStatus(v string) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.Status = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) SetUpdateTime(v string) *DescribeInstanceSDGStatusResponseBodyDeploymentStatus {
	s.UpdateTime = &v
	return s
}

func (s *DescribeInstanceSDGStatusResponseBodyDeploymentStatus) Validate() error {
	return dara.Validate(s)
}
