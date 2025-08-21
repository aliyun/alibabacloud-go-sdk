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
	// The category of the disk.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: all-flash disk.
	//
	// 	- local_hdd: local HDD.
	//
	// 	- local_ssd: local SSD.
	//
	// example:
	//
	// local_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The billing method.
	//
	// 	- prePay: subscription.
	//
	// 	- postpay: pay-as-you-go.
	//
	// example:
	//
	// prepay
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-5soak1gqa507lyfzvz0xo****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// [\\"d-wz99njena32z90ki****\\"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The name of the disk.
	//
	// example:
	//
	// DiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The purchase method of the disk. Valid values:
	//
	// 	- ServiceDisk: The disk is purchased when ENS is activated.
	//
	// 	- ResoureDisk: The disk is purchased when the instance is created.
	//
	// 	- PostPayDisk: The disk is separately purchased.
	//
	// example:
	//
	// PostPayDisk
	DiskType *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-kunming-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The node information.
	//
	// example:
	//
	// ["cn-suzhou-telecom","cn-chengdu-telecom"]
	EnsRegionIds *string `json:"EnsRegionIds,omitempty" xml:"EnsRegionIds,omitempty"`
	// The ID of the instance to which the disk is attached.
	//
	// example:
	//
	// i-5t77rb0yoz79m28ku60sx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The order in which you want to sort the returned data. Example: {"EnsRegionId":"desc"}. By default, the nodes are sorted by IDs in descending order.
	//
	// example:
	//
	// { "DiskNameSort":"desc","EnsRegionIdSort":"asc" }
	OrderByParams *string `json:"OrderByParams,omitempty" xml:"OrderByParams,omitempty"`
	// The number of the page to return. Pages start from page **1**.
	//
	// Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: **50**.
	//
	// Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the snapshot.
	//
	// example:
	//
	// s-897654321****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The status of the disk. Valid values:
	//
	// 	- In-use: The disk is in use.
	//
	// 	- Available: The disk can be attached.
	//
	// 	- Attaching: The disk is being attached.
	//
	// 	- Detaching: The disk is being detached.
	//
	// 	- Creating: The disk is being created.
	//
	// 	- ReIniting: The disk is being reset.
	//
	// 	- Deleting: The disk is being released.
	//
	// 	- Deleted: The disk is released.
	//
	// 	- Expiring: The disk is about to expire.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the disk. Valid values:
	//
	// 	- system: system disk.
	//
	// 	- data: data disk.
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
