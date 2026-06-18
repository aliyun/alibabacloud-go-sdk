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
	SetDbType(v string) *DescribeActiveOperationTasksRequest
	GetDbType() *string
	SetInsName(v string) *DescribeActiveOperationTasksRequest
	GetInsName() *string
	SetPageNumber(v int64) *DescribeActiveOperationTasksRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeActiveOperationTasksRequest
	GetPageSize() *int64
	SetProductId(v string) *DescribeActiveOperationTasksRequest
	GetProductId() *string
	SetRegion(v string) *DescribeActiveOperationTasksRequest
	GetRegion() *string
	SetRegionId(v string) *DescribeActiveOperationTasksRequest
	GetRegionId() *string
	SetStatus(v int64) *DescribeActiveOperationTasksRequest
	GetStatus() *int64
	SetTaskType(v string) *DescribeActiveOperationTasksRequest
	GetTaskType() *string
}

type DescribeActiveOperationTasksRequest struct {
	// Specifies whether cancellation is allowed. Valid values:
	//
	// - **-1**: All O&M events can be canceled.
	//
	// - **0**: Not allowed.
	//
	// - **1**: Allowed.
	//
	// example:
	//
	// -1
	AllowCancel *int64 `json:"AllowCancel,omitempty" xml:"AllowCancel,omitempty"`
	// Specifies whether the time can be modified. Valid values:
	//
	// - **-1**: The time of all O&M events can be modified.
	//
	// - **0**: Not allowed.
	//
	// - **1**: Allowed.
	//
	// example:
	//
	// -1
	AllowChange *int64 `json:"AllowChange,omitempty" xml:"AllowChange,omitempty"`
	// The change level. Default value: all. Valid values:
	//
	// - **all**: All levels of exception recovery and system O&M.
	//
	// - **S0**: Exception recovery.
	//
	// - **S1**: System O&M.
	//
	// example:
	//
	// all
	ChangeLevel *string `json:"ChangeLevel,omitempty" xml:"ChangeLevel,omitempty"`
	// The database type. Set this parameter to polarx.
	//
	// example:
	//
	// polarx
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The instance name. Default value: empty string.
	//
	// example:
	//
	// pxc-xxxxx
	InsName *string `json:"InsName,omitempty" xml:"InsName,omitempty"`
	// The page number. The value must be greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 25.
	//
	// example:
	//
	// 25
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product name. Default value: polarx.
	//
	// example:
	//
	// polarx
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task status. Valid values:
	//
	// - **-1**: All pending and running tasks.
	//
	// - **3**: Pending.
	//
	// - **4**: Running.
	//
	// example:
	//
	// -1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The task type. A value of all indicates that all tasks of the user are queried. Set this parameter to all.
	//
	// example:
	//
	// all
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

func (s *DescribeActiveOperationTasksRequest) GetDbType() *string {
	return s.DbType
}

func (s *DescribeActiveOperationTasksRequest) GetInsName() *string {
	return s.InsName
}

func (s *DescribeActiveOperationTasksRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeActiveOperationTasksRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeActiveOperationTasksRequest) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeActiveOperationTasksRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeActiveOperationTasksRequest) GetRegionId() *string {
	return s.RegionId
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

func (s *DescribeActiveOperationTasksRequest) SetDbType(v string) *DescribeActiveOperationTasksRequest {
	s.DbType = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetInsName(v string) *DescribeActiveOperationTasksRequest {
	s.InsName = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageNumber(v int64) *DescribeActiveOperationTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetPageSize(v int64) *DescribeActiveOperationTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetProductId(v string) *DescribeActiveOperationTasksRequest {
	s.ProductId = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetRegion(v string) *DescribeActiveOperationTasksRequest {
	s.Region = &v
	return s
}

func (s *DescribeActiveOperationTasksRequest) SetRegionId(v string) *DescribeActiveOperationTasksRequest {
	s.RegionId = &v
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
