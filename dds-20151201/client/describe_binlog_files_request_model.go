// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBinlogFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBinlogId(v string) *DescribeBinlogFilesRequest
	GetBinlogId() *string
	SetDBInstanceId(v string) *DescribeBinlogFilesRequest
	GetDBInstanceId() *string
	SetDestRegion(v string) *DescribeBinlogFilesRequest
	GetDestRegion() *string
	SetEndTime(v string) *DescribeBinlogFilesRequest
	GetEndTime() *string
	SetNodeId(v string) *DescribeBinlogFilesRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeBinlogFilesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeBinlogFilesRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeBinlogFilesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeBinlogFilesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeBinlogFilesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeBinlogFilesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeBinlogFilesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeBinlogFilesRequest
	GetResourceOwnerId() *int64
	SetSrcRegion(v string) *DescribeBinlogFilesRequest
	GetSrcRegion() *string
	SetStartTime(v string) *DescribeBinlogFilesRequest
	GetStartTime() *string
}

type DescribeBinlogFilesRequest struct {
	BinlogId *string `json:"BinlogId,omitempty" xml:"BinlogId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dds-wz9ca592fc637a54
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	DestRegion   *string `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-04-17T05:50:28.914Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// d-uf696817a1b5d9f4
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-zhangjiakou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SrcRegion            *string `json:"SrcRegion,omitempty" xml:"SrcRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2021-03-23T06:24:21.425Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeBinlogFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBinlogFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBinlogFilesRequest) GetBinlogId() *string {
	return s.BinlogId
}

func (s *DescribeBinlogFilesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeBinlogFilesRequest) GetDestRegion() *string {
	return s.DestRegion
}

func (s *DescribeBinlogFilesRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeBinlogFilesRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeBinlogFilesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeBinlogFilesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeBinlogFilesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeBinlogFilesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeBinlogFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBinlogFilesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeBinlogFilesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeBinlogFilesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeBinlogFilesRequest) GetSrcRegion() *string {
	return s.SrcRegion
}

func (s *DescribeBinlogFilesRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeBinlogFilesRequest) SetBinlogId(v string) *DescribeBinlogFilesRequest {
	s.BinlogId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetDBInstanceId(v string) *DescribeBinlogFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetDestRegion(v string) *DescribeBinlogFilesRequest {
	s.DestRegion = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetEndTime(v string) *DescribeBinlogFilesRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetNodeId(v string) *DescribeBinlogFilesRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetOwnerAccount(v string) *DescribeBinlogFilesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetOwnerId(v int64) *DescribeBinlogFilesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetPageNumber(v int64) *DescribeBinlogFilesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetPageSize(v int64) *DescribeBinlogFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetRegionId(v string) *DescribeBinlogFilesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetResourceGroupId(v string) *DescribeBinlogFilesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetResourceOwnerAccount(v string) *DescribeBinlogFilesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetResourceOwnerId(v int64) *DescribeBinlogFilesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetSrcRegion(v string) *DescribeBinlogFilesRequest {
	s.SrcRegion = &v
	return s
}

func (s *DescribeBinlogFilesRequest) SetStartTime(v string) *DescribeBinlogFilesRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeBinlogFilesRequest) Validate() error {
	return dara.Validate(s)
}
