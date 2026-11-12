// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeDisksRequest
	GetCategory() *string
	SetDiskChargeType(v string) *DescribeDisksRequest
	GetDiskChargeType() *string
	SetDiskId(v string) *DescribeDisksRequest
	GetDiskId() *string
	SetDiskIds(v string) *DescribeDisksRequest
	GetDiskIds() *string
	SetDiskName(v string) *DescribeDisksRequest
	GetDiskName() *string
	SetDiskType(v string) *DescribeDisksRequest
	GetDiskType() *string
	SetEnsRegionId(v string) *DescribeDisksRequest
	GetEnsRegionId() *string
	SetEnsRegionIds(v string) *DescribeDisksRequest
	GetEnsRegionIds() *string
	SetInstanceId(v string) *DescribeDisksRequest
	GetInstanceId() *string
	SetOrderByParams(v string) *DescribeDisksRequest
	GetOrderByParams() *string
	SetPageNumber(v string) *DescribeDisksRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeDisksRequest
	GetPageSize() *string
	SetSnapshotId(v string) *DescribeDisksRequest
	GetSnapshotId() *string
	SetStatus(v string) *DescribeDisksRequest
	GetStatus() *string
	SetType(v string) *DescribeDisksRequest
	GetType() *string
}

type DescribeDisksRequest struct {
	// The category of the cloud disk. Valid values:
	//
	// - cloud_efficiency: Ultra cloud disk.
	//
	// - cloud_ssd: All-flash cloud disk.
	//
	// - local_hdd: Local HDD.
	//
	// - local_ssd: Local SSD.
	//
	// example:
	//
	// local_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The billing method. Valid values:
	//
	// - prepay: Subscription.
	//
	// - postpay: Pay-as-you-go.
	//
	// example:
	//
	// prepay
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The cloud disk ID.
	//
	// example:
	//
	// d-5soak1gqa507lyfzvz0xo****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The IDs of the cloud disks.
	//
	// example:
	//
	// [\\"d-wz99njena32z90ki****\\"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The name of the cloud disk.
	//
	// example:
	//
	// DiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The purchase method of the cloud disk. Valid values:
	//
	// - ServiceDisk: Created by Edge Node Service.
	//
	// - ResoureDisk: Purchased with an instance.
	//
	// - PostPayDisk: Purchased separately.
	//
	// example:
	//
	// PostPayDisk
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the ENS edge zone.
	//
	// example:
	//
	// cn-kunming-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The list of edge zone IDs.
	//
	// example:
	//
	// ["cn-suzhou-telecom","cn-chengdu-telecom"]
	EnsRegionIds *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
	// The ID of the instance to which the cloud disk is attached.
	//
	// example:
	//
	// i-5t77rb0yoz79m28ku60sx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The sorting parameter. Format example: {"EnsRegionId":"desc"}. By default, the results are sorted in descending order of edge zone ID.
	//
	// example:
	//
	// { "DiskNameSort":"desc","EnsRegionIdSort":"asc" }
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	// The page number of the query list. The minimum value is **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: **50**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-897654321****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The status of the cloud disk. Valid values:
	//
	// - In-use: In use.
	//
	// - Available: Unattached.
	//
	// - Attaching: Attaching.
	//
	// - Detaching: Detaching.
	//
	// - Creating: Creating.
	//
	// - ReIniting: Reinitializing.
	//
	// - Deleting: Releasing.
	//
	// - Deleted: Released.
	//
	// - Expiring: Expiring.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the cloud disk. Valid values:
	//
	// - system: System cloud disk.
	//
	// - data: Data cloud disk.
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDisksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksRequest) GoString() string {
	return s.String()
}

func (s *DescribeDisksRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeDisksRequest) GetDiskChargeType() *string {
	return s.DiskChargeType
}

func (s *DescribeDisksRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDisksRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *DescribeDisksRequest) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeDisksRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *DescribeDisksRequest) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeDisksRequest) GetEnsRegionIds() *string {
	return s.EnsRegionIds
}

func (s *DescribeDisksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDisksRequest) GetOrderByParams() *string {
	return s.OrderByParams
}

func (s *DescribeDisksRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDisksRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDisksRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeDisksRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDisksRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDisksRequest) SetCategory(v string) *DescribeDisksRequest {
	s.Category = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskChargeType(v string) *DescribeDisksRequest {
	s.DiskChargeType = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskId(v string) *DescribeDisksRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskIds(v string) *DescribeDisksRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskName(v string) *DescribeDisksRequest {
	s.DiskName = &v
	return s
}

func (s *DescribeDisksRequest) SetDiskType(v string) *DescribeDisksRequest {
	s.DiskType = &v
	return s
}

func (s *DescribeDisksRequest) SetEnsRegionId(v string) *DescribeDisksRequest {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeDisksRequest) SetEnsRegionIds(v string) *DescribeDisksRequest {
	s.EnsRegionIds = &v
	return s
}

func (s *DescribeDisksRequest) SetInstanceId(v string) *DescribeDisksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDisksRequest) SetOrderByParams(v string) *DescribeDisksRequest {
	s.OrderByParams = &v
	return s
}

func (s *DescribeDisksRequest) SetPageNumber(v string) *DescribeDisksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDisksRequest) SetPageSize(v string) *DescribeDisksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDisksRequest) SetSnapshotId(v string) *DescribeDisksRequest {
	s.SnapshotId = &v
	return s
}

func (s *DescribeDisksRequest) SetStatus(v string) *DescribeDisksRequest {
	s.Status = &v
	return s
}

func (s *DescribeDisksRequest) SetType(v string) *DescribeDisksRequest {
	s.Type = &v
	return s
}

func (s *DescribeDisksRequest) Validate() error {
	return dara.Validate(s)
}
