// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v int64) *ListOperationTaskRequest
	GetCheckId() *int64
	SetCurrentPage(v int32) *ListOperationTaskRequest
	GetCurrentPage() *int32
	SetLang(v string) *ListOperationTaskRequest
	GetLang() *string
	SetOperationTaskInstances(v []*ListOperationTaskRequestOperationTaskInstances) *ListOperationTaskRequest
	GetOperationTaskInstances() []*ListOperationTaskRequestOperationTaskInstances
	SetPageSize(v int32) *ListOperationTaskRequest
	GetPageSize() *int32
	SetStatuses(v []*string) *ListOperationTaskRequest
	GetStatuses() []*string
	SetTaskId(v string) *ListOperationTaskRequest
	GetTaskId() *string
	SetType(v string) *ListOperationTaskRequest
	GetType() *string
}

type ListOperationTaskRequest struct {
	// The ID of the check item.
	//
	// > You can call the [ListCheckResult](~~ListCheckResult~~) API to obtain the check item ID.
	//
	// example:
	//
	// 76
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The page number to display in a paginated query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Set the language type for the request and response messages. The default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The instance information of the operation tasks to be queried.
	OperationTaskInstances []*ListOperationTaskRequestOperationTaskInstances `json:"OperationTaskInstances,omitempty" xml:"OperationTaskInstances,omitempty" type:"Repeated"`
	// The maximum number of items to display per page in a paginated query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A list of task statuses. Multiple statuses should be separated by commas (,). Values:
	//
	// - **INIT**: Initialization
	//
	// - **FAIL**: Processing failed
	//
	// - **THROTTLING**: Repairing concurrently
	//
	// - **IN_BACKUP**: Backing up
	//
	// - **BACKED_UP**: Backed up
	//
	// - **BACKUP_FAIL**: Backup failed
	//
	// - **REPAIRING**: Repairing
	//
	// - **REPAIR_SUCCESS**: Repair succeeded
	//
	// - **REPAIR_FAIL**: Repair failed
	//
	// - **REPAIR_SUCCESS_VERIFYING**: Verifying repair success
	//
	// - **REPAIR_SUCCESS_UNVERIFIED**: Repair success verification failed
	//
	// - **REPAIR_SUCCESS_VERIFIED**: Repair success verified
	//
	// - **REPAIR_RE_EXECUTE**: Re-executing repair
	//
	// - **ROLL_BACKING**: Rolling back
	//
	// - **ROLL_BACKED**: Rolled back
	//
	// - **ROLL_BACK_FAIL**: Rollback failed
	//
	// - **ROLL_BACK_INIT**: Rollback initiated
	//
	// - **ROLL_BACK_VERIFYING**: Verifying rollback success
	//
	// - **ROLL_BACK_UNVERIFIED**: Rollback success verification failed
	//
	// - **ROLL_BACK_VERIFIED**: Rollback success verified
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
	// The ID of the task to be queried.
	//
	// example:
	//
	// d7b2acf8d362742123e4a84e1bf8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of operation task to be queried:
	//
	// - REPAIR: Repair
	//
	// - ROLLBACK: Rollback
	//
	// example:
	//
	// REPAIR
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListOperationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationTaskRequest) GoString() string {
	return s.String()
}

func (s *ListOperationTaskRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *ListOperationTaskRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListOperationTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *ListOperationTaskRequest) GetOperationTaskInstances() []*ListOperationTaskRequestOperationTaskInstances {
	return s.OperationTaskInstances
}

func (s *ListOperationTaskRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOperationTaskRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListOperationTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListOperationTaskRequest) GetType() *string {
	return s.Type
}

func (s *ListOperationTaskRequest) SetCheckId(v int64) *ListOperationTaskRequest {
	s.CheckId = &v
	return s
}

func (s *ListOperationTaskRequest) SetCurrentPage(v int32) *ListOperationTaskRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListOperationTaskRequest) SetLang(v string) *ListOperationTaskRequest {
	s.Lang = &v
	return s
}

func (s *ListOperationTaskRequest) SetOperationTaskInstances(v []*ListOperationTaskRequestOperationTaskInstances) *ListOperationTaskRequest {
	s.OperationTaskInstances = v
	return s
}

func (s *ListOperationTaskRequest) SetPageSize(v int32) *ListOperationTaskRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationTaskRequest) SetStatuses(v []*string) *ListOperationTaskRequest {
	s.Statuses = v
	return s
}

func (s *ListOperationTaskRequest) SetTaskId(v string) *ListOperationTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ListOperationTaskRequest) SetType(v string) *ListOperationTaskRequest {
	s.Type = &v
	return s
}

func (s *ListOperationTaskRequest) Validate() error {
	return dara.Validate(s)
}

type ListOperationTaskRequestOperationTaskInstances struct {
	// The instance ID of the server.
	//
	// example:
	//
	// am-bp1vy2yaj1nvx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Cloud asset vendor.
	//
	// example:
	//
	// ALIYUN
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListOperationTaskRequestOperationTaskInstances) String() string {
	return dara.Prettify(s)
}

func (s ListOperationTaskRequestOperationTaskInstances) GoString() string {
	return s.String()
}

func (s *ListOperationTaskRequestOperationTaskInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationTaskRequestOperationTaskInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOperationTaskRequestOperationTaskInstances) GetVendor() *string {
	return s.Vendor
}

func (s *ListOperationTaskRequestOperationTaskInstances) SetInstanceId(v string) *ListOperationTaskRequestOperationTaskInstances {
	s.InstanceId = &v
	return s
}

func (s *ListOperationTaskRequestOperationTaskInstances) SetRegionId(v string) *ListOperationTaskRequestOperationTaskInstances {
	s.RegionId = &v
	return s
}

func (s *ListOperationTaskRequestOperationTaskInstances) SetVendor(v string) *ListOperationTaskRequestOperationTaskInstances {
	s.Vendor = &v
	return s
}

func (s *ListOperationTaskRequestOperationTaskInstances) Validate() error {
	return dara.Validate(s)
}
