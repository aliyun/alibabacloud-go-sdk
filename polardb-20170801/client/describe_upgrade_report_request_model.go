// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreationCategory(v string) *DescribeUpgradeReportRequest
	GetCreationCategory() *string
	SetDBType(v string) *DescribeUpgradeReportRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeUpgradeReportRequest
	GetDBVersion() *string
	SetOwnerAccount(v string) *DescribeUpgradeReportRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUpgradeReportRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeUpgradeReportRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeUpgradeReportRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeUpgradeReportRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeUpgradeReportRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUpgradeReportRequest
	GetResourceOwnerId() *int64
	SetSourceDBClusterId(v string) *DescribeUpgradeReportRequest
	GetSourceDBClusterId() *string
	SetStatus(v string) *DescribeUpgradeReportRequest
	GetStatus() *string
	SetTaskId(v string) *DescribeUpgradeReportRequest
	GetTaskId() *string
	SetType(v string) *DescribeUpgradeReportRequest
	GetType() *string
}

type DescribeUpgradeReportRequest struct {
	// example:
	//
	// Normal
	CreationCategory *string `json:"CreationCategory,omitempty" xml:"CreationCategory,omitempty"`
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// example:
	//
	// 8.0
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// rm-2zen5pe5vi56447d0
	SourceDBClusterId *string `json:"SourceDBClusterId,omitempty" xml:"SourceDBClusterId,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2321321
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// orca
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeUpgradeReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeReportRequest) GetCreationCategory() *string {
	return s.CreationCategory
}

func (s *DescribeUpgradeReportRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeUpgradeReportRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeUpgradeReportRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUpgradeReportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUpgradeReportRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeUpgradeReportRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUpgradeReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUpgradeReportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUpgradeReportRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUpgradeReportRequest) GetSourceDBClusterId() *string {
	return s.SourceDBClusterId
}

func (s *DescribeUpgradeReportRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeUpgradeReportRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeUpgradeReportRequest) GetType() *string {
	return s.Type
}

func (s *DescribeUpgradeReportRequest) SetCreationCategory(v string) *DescribeUpgradeReportRequest {
	s.CreationCategory = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetDBType(v string) *DescribeUpgradeReportRequest {
	s.DBType = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetDBVersion(v string) *DescribeUpgradeReportRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetOwnerAccount(v string) *DescribeUpgradeReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetOwnerId(v int64) *DescribeUpgradeReportRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetPageNumber(v int32) *DescribeUpgradeReportRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetPageSize(v int32) *DescribeUpgradeReportRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetRegionId(v string) *DescribeUpgradeReportRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetResourceOwnerAccount(v string) *DescribeUpgradeReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetResourceOwnerId(v int64) *DescribeUpgradeReportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetSourceDBClusterId(v string) *DescribeUpgradeReportRequest {
	s.SourceDBClusterId = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetStatus(v string) *DescribeUpgradeReportRequest {
	s.Status = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetTaskId(v string) *DescribeUpgradeReportRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeUpgradeReportRequest) SetType(v string) *DescribeUpgradeReportRequest {
	s.Type = &v
	return s
}

func (s *DescribeUpgradeReportRequest) Validate() error {
	return dara.Validate(s)
}
