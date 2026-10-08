// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchTaskQuery(v *ListBatchTasksRequestBatchTaskQuery) *ListBatchTasksRequest
	GetBatchTaskQuery() *ListBatchTasksRequestBatchTaskQuery
	SetOpTenantId(v int64) *ListBatchTasksRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListBatchTasksRequest
	GetOpUserId() *string
}

type ListBatchTasksRequest struct {
	// This parameter is required.
	BatchTaskQuery *ListBatchTasksRequestBatchTaskQuery `json:"BatchTaskQuery,omitempty" xml:"BatchTaskQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListBatchTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTasksRequest) GoString() string {
	return s.String()
}

func (s *ListBatchTasksRequest) GetBatchTaskQuery() *ListBatchTasksRequestBatchTaskQuery {
	return s.BatchTaskQuery
}

func (s *ListBatchTasksRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListBatchTasksRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListBatchTasksRequest) SetBatchTaskQuery(v *ListBatchTasksRequestBatchTaskQuery) *ListBatchTasksRequest {
	s.BatchTaskQuery = v
	return s
}

func (s *ListBatchTasksRequest) SetOpTenantId(v int64) *ListBatchTasksRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListBatchTasksRequest) SetOpUserId(v string) *ListBatchTasksRequest {
	s.OpUserId = &v
	return s
}

func (s *ListBatchTasksRequest) Validate() error {
	if s.BatchTaskQuery != nil {
		if err := s.BatchTaskQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBatchTasksRequestBatchTaskQuery struct {
	ConditionScheduleEnable *bool `json:"ConditionScheduleEnable,omitempty" xml:"ConditionScheduleEnable,omitempty"`
	// example:
	//
	// 1785930337435
	CreateBeginTime *int64 `json:"CreateBeginTime,omitempty" xml:"CreateBeginTime,omitempty"`
	// example:
	//
	// 1788608737435
	CreateEndTime       *int64    `json:"CreateEndTime,omitempty" xml:"CreateEndTime,omitempty"`
	DevelopOwnerList    []*string `json:"DevelopOwnerList,omitempty" xml:"DevelopOwnerList,omitempty" type:"Repeated"`
	DirectoryList       []*string `json:"DirectoryList,omitempty" xml:"DirectoryList,omitempty" type:"Repeated"`
	IncludeSubDirectory *bool     `json:"IncludeSubDirectory,omitempty" xml:"IncludeSubDirectory,omitempty"`
	// example:
	//
	// dwd_order
	Keyword              *string   `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	LastSubmitStatusList []*string `json:"LastSubmitStatusList,omitempty" xml:"LastSubmitStatusList,omitempty" type:"Repeated"`
	LockUserList         []*string `json:"LockUserList,omitempty" xml:"LockUserList,omitempty" type:"Repeated"`
	// example:
	//
	// 1785930337435
	ModifiedBeginTime *int64 `json:"ModifiedBeginTime,omitempty" xml:"ModifiedBeginTime,omitempty"`
	// example:
	//
	// 1788608737435
	ModifiedEndTime     *int64    `json:"ModifiedEndTime,omitempty" xml:"ModifiedEndTime,omitempty"`
	NodeStatusList      []*int32  `json:"NodeStatusList,omitempty" xml:"NodeStatusList,omitempty" type:"Repeated"`
	OpsOwnerList        []*string `json:"OpsOwnerList,omitempty" xml:"OpsOwnerList,omitempty" type:"Repeated"`
	OutputTableNameList []*string `json:"OutputTableNameList,omitempty" xml:"OutputTableNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7086194564164288
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Published *bool  `json:"Published,omitempty" xml:"Published,omitempty"`
	// example:
	//
	// 7305621095333696
	RefCodeTemplateId        *string   `json:"RefCodeTemplateId,omitempty" xml:"RefCodeTemplateId,omitempty"`
	ScheduleIntervalTypeList []*string `json:"ScheduleIntervalTypeList,omitempty" xml:"ScheduleIntervalTypeList,omitempty" type:"Repeated"`
	TaskStatusList           []*int32  `json:"TaskStatusList,omitempty" xml:"TaskStatusList,omitempty" type:"Repeated"`
	TaskTagList              []*string `json:"TaskTagList,omitempty" xml:"TaskTagList,omitempty" type:"Repeated"`
	TaskTypeList             []*int32  `json:"TaskTypeList,omitempty" xml:"TaskTypeList,omitempty" type:"Repeated"`
}

func (s ListBatchTasksRequestBatchTaskQuery) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTasksRequestBatchTaskQuery) GoString() string {
	return s.String()
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetConditionScheduleEnable() *bool {
	return s.ConditionScheduleEnable
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetCreateBeginTime() *int64 {
	return s.CreateBeginTime
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetCreateEndTime() *int64 {
	return s.CreateEndTime
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetDevelopOwnerList() []*string {
	return s.DevelopOwnerList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetDirectoryList() []*string {
	return s.DirectoryList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetIncludeSubDirectory() *bool {
	return s.IncludeSubDirectory
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetLastSubmitStatusList() []*string {
	return s.LastSubmitStatusList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetLockUserList() []*string {
	return s.LockUserList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetModifiedBeginTime() *int64 {
	return s.ModifiedBeginTime
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetModifiedEndTime() *int64 {
	return s.ModifiedEndTime
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetNodeStatusList() []*int32 {
	return s.NodeStatusList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetOpsOwnerList() []*string {
	return s.OpsOwnerList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetOutputTableNameList() []*string {
	return s.OutputTableNameList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetPage() *int32 {
	return s.Page
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetPublished() *bool {
	return s.Published
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetRefCodeTemplateId() *string {
	return s.RefCodeTemplateId
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetScheduleIntervalTypeList() []*string {
	return s.ScheduleIntervalTypeList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetTaskStatusList() []*int32 {
	return s.TaskStatusList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetTaskTagList() []*string {
	return s.TaskTagList
}

func (s *ListBatchTasksRequestBatchTaskQuery) GetTaskTypeList() []*int32 {
	return s.TaskTypeList
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetConditionScheduleEnable(v bool) *ListBatchTasksRequestBatchTaskQuery {
	s.ConditionScheduleEnable = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetCreateBeginTime(v int64) *ListBatchTasksRequestBatchTaskQuery {
	s.CreateBeginTime = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetCreateEndTime(v int64) *ListBatchTasksRequestBatchTaskQuery {
	s.CreateEndTime = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetDevelopOwnerList(v []*string) *ListBatchTasksRequestBatchTaskQuery {
	s.DevelopOwnerList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetDirectoryList(v []*string) *ListBatchTasksRequestBatchTaskQuery {
	s.DirectoryList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetIncludeSubDirectory(v bool) *ListBatchTasksRequestBatchTaskQuery {
	s.IncludeSubDirectory = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetKeyword(v string) *ListBatchTasksRequestBatchTaskQuery {
	s.Keyword = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetLastSubmitStatusList(v []*string) *ListBatchTasksRequestBatchTaskQuery {
	s.LastSubmitStatusList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetLockUserList(v []*string) *ListBatchTasksRequestBatchTaskQuery {
	s.LockUserList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetModifiedBeginTime(v int64) *ListBatchTasksRequestBatchTaskQuery {
	s.ModifiedBeginTime = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetModifiedEndTime(v int64) *ListBatchTasksRequestBatchTaskQuery {
	s.ModifiedEndTime = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetNodeStatusList(v []*int32) *ListBatchTasksRequestBatchTaskQuery {
	s.NodeStatusList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetOpsOwnerList(v []*string) *ListBatchTasksRequestBatchTaskQuery {
	s.OpsOwnerList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetOutputTableNameList(v []*string) *ListBatchTasksRequestBatchTaskQuery {
	s.OutputTableNameList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetPage(v int32) *ListBatchTasksRequestBatchTaskQuery {
	s.Page = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetPageSize(v int32) *ListBatchTasksRequestBatchTaskQuery {
	s.PageSize = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetProjectId(v int64) *ListBatchTasksRequestBatchTaskQuery {
	s.ProjectId = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetPublished(v bool) *ListBatchTasksRequestBatchTaskQuery {
	s.Published = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetRefCodeTemplateId(v string) *ListBatchTasksRequestBatchTaskQuery {
	s.RefCodeTemplateId = &v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetScheduleIntervalTypeList(v []*string) *ListBatchTasksRequestBatchTaskQuery {
	s.ScheduleIntervalTypeList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetTaskStatusList(v []*int32) *ListBatchTasksRequestBatchTaskQuery {
	s.TaskStatusList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetTaskTagList(v []*string) *ListBatchTasksRequestBatchTaskQuery {
	s.TaskTagList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) SetTaskTypeList(v []*int32) *ListBatchTasksRequestBatchTaskQuery {
	s.TaskTypeList = v
	return s
}

func (s *ListBatchTasksRequestBatchTaskQuery) Validate() error {
	return dara.Validate(s)
}
