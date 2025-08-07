// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActiveOperationTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowCancel(v int64) *DescribeActiveOperationTasksRequest
	GetAllowCancel() *int64
	SetAllowChange(v int64) *DescribeActiveOperationTasksRequest
	GetAllowChange() *int64
	SetChangeLevel(v string) *DescribeActiveOperationTasksRequest
	GetChangeLevel() *string
	SetDBClusterId(v string) *DescribeActiveOperationTasksRequest
	GetDBClusterId() *string
	SetDBType(v string) *DescribeActiveOperationTasksRequest
	GetDBType() *string
	SetOwnerAccount(v string) *DescribeActiveOperationTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActiveOperationTasksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeActiveOperationTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeActiveOperationTasksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeActiveOperationTasksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeActiveOperationTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActiveOperationTasksRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeActiveOperationTasksRequest
	GetSecurityToken() *string
	SetStatus(v int64) *DescribeActiveOperationTasksRequest
	GetStatus() *int64
	SetTaskType(v string) *DescribeActiveOperationTasksRequest
	GetTaskType() *string
}

type DescribeActiveOperationTasksRequest struct {
	// example:
	//
	// -1
	AllowCancel *int64 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// example:
	//
	// -1
	AllowChange *int64 `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// example:
	//
	// all
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// example:
	//
	// pc-3ns***********d5d
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// MySQL
	DBType       *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the page to return. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// -1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// DatabaseProxyUpgrading
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s DescribeActiveOperationTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActiveOperationTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeActiveOperationTasksRequest) GetAllowCancel() *int64 {
	return s.AllowCancel
}

func (s *DescribeActiveOperationTasksRequest) GetAllowChange() *int64 {
	return s.AllowChange
}

func (s *DescribeActiveOperationTasksRequest) GetChangeLevel() *string {
	return s.ChangeLevel
}

func (s *DescribeActiveOperationTasksRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeActiveOperationTasksRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeActiveOperationTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActiveOperationTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActiveOperationTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeActiveOperationTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeActiveOperationTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeActiveOperationTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActiveOperationTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActiveOperationTasksRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeActiveOperationTasksRequest) GetStatus() *int64 {
	return s.Status
}

func (s *DescribeActiveOperationTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeActiveOperationTasksRequest) SetAllowCancel(v int64) *DescribeActiveOperationTasksRequest {
	s.AllowCancel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetAllowChange(v int64) *DescribeActiveOperationTasksRequest {
	s.AllowChange = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetChangeLevel(v string) *DescribeActiveOperationTasksRequest {
	s.ChangeLevel = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetDBClusterId(v string) *DescribeActiveOperationTasksRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetDBType(v string) *DescribeActiveOperationTasksRequest {
	s.DBType = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetOwnerAccount(v string) *DescribeActiveOperationTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetOwnerId(v int64) *DescribeActiveOperationTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageNumber(v int32) *DescribeActiveOperationTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageSize(v int32) *DescribeActiveOperationTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetRegionId(v string) *DescribeActiveOperationTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetResourceOwnerAccount(v string) *DescribeActiveOperationTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetResourceOwnerId(v int64) *DescribeActiveOperationTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetSecurityToken(v string) *DescribeActiveOperationTasksRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetStatus(v int64) *DescribeActiveOperationTasksRequest {
	s.Status = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetTaskType(v string) *DescribeActiveOperationTasksRequest {
	s.TaskType = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) Validate() error {
	return dara.Validate(s)
}
