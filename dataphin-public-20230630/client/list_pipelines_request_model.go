// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v *ListPipelinesRequestContext) *ListPipelinesRequest
	GetContext() *ListPipelinesRequestContext
	SetListCommand(v *ListPipelinesRequestListCommand) *ListPipelinesRequest
	GetListCommand() *ListPipelinesRequestListCommand
	SetOpTenantId(v int64) *ListPipelinesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListPipelinesRequest
	GetOpUserId() *string
}

type ListPipelinesRequest struct {
	// The request context.
	//
	// This parameter is required.
	Context *ListPipelinesRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The query parameters.
	//
	// This parameter is required.
	ListCommand *ListPipelinesRequestListCommand `json:"ListCommand,omitempty" xml:"ListCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListPipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) GetContext() *ListPipelinesRequestContext {
	return s.Context
}

func (s *ListPipelinesRequest) GetListCommand() *ListPipelinesRequestListCommand {
	return s.ListCommand
}

func (s *ListPipelinesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListPipelinesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListPipelinesRequest) SetContext(v *ListPipelinesRequestContext) *ListPipelinesRequest {
	s.Context = v
	return s
}

func (s *ListPipelinesRequest) SetListCommand(v *ListPipelinesRequestListCommand) *ListPipelinesRequest {
	s.ListCommand = v
	return s
}

func (s *ListPipelinesRequest) SetOpTenantId(v int64) *ListPipelinesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListPipelinesRequest) SetOpUserId(v string) *ListPipelinesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListPipelinesRequest) Validate() error {
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	if s.ListCommand != nil {
		if err := s.ListCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPipelinesRequestContext struct {
	// The environment identifier. Valid values: DEV: development environment. PROD: production environment. Default value: PROD.
	//
	// example:
	//
	// PROD
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7128268454335680
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListPipelinesRequestContext) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesRequestContext) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequestContext) GetEnv() *string {
	return s.Env
}

func (s *ListPipelinesRequestContext) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListPipelinesRequestContext) SetEnv(v string) *ListPipelinesRequestContext {
	s.Env = &v
	return s
}

func (s *ListPipelinesRequestContext) SetProjectId(v int64) *ListPipelinesRequestContext {
	s.ProjectId = &v
	return s
}

func (s *ListPipelinesRequestContext) Validate() error {
	return dara.Validate(s)
}

type ListPipelinesRequestListCommand struct {
	// The user IDs of the creators used for filtering. If this parameter is left empty, no filtering is performed. If multiple IDs are specified, they are evaluated by using the OR operator.
	CreatorList []*string `json:"CreatorList,omitempty" xml:"CreatorList,omitempty" type:"Repeated"`
	// The user IDs of the development owners used for filtering. If this parameter is left empty, no filtering is performed. If multiple IDs are specified, they are evaluated by using the OR operator.
	DevelopOwnerList []*string `json:"DevelopOwnerList,omitempty" xml:"DevelopOwnerList,omitempty" type:"Repeated"`
	// The full paths of the directories to query. If this parameter is left empty, the root directory is queried.
	Directories []*string `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// Specifies whether an exact match is performed on the node name. Default value: false.
	ExactMatch *bool `json:"ExactMatch,omitempty" xml:"ExactMatch,omitempty"`
	// The keywords used to filter task names. This parameter is optional. If this parameter is left empty, tasks are not filtered by name. If exact match is enabled, this parameter specifies a list of full names. If fuzzy match is used, this parameter specifies a list of keywords. If multiple keywords are specified, they are evaluated by using the OR operator.
	Keywords []*string `json:"Keywords,omitempty" xml:"Keywords,omitempty" type:"Repeated"`
	// The cursor for the next page. Leave this parameter empty for the first page. For subsequent pages, pass the NextCursor value returned from the previous page.
	//
	// example:
	//
	// 123
	NextCursor *int64 `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	// The user IDs of the O&M owners used for filtering. If this parameter is left empty, no filtering is performed. If multiple IDs are specified, they are evaluated by using the OR operator.
	OpsOwnerList []*string `json:"OpsOwnerList,omitempty" xml:"OpsOwnerList,omitempty" type:"Repeated"`
	// The page number. Default value: 1. The value starts from 1.
	//
	// example:
	//
	// 3
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The task types. Default value: [0]. Valid values: 0: offline integration. 1: real-time integration. 13: data ingestion. 14: offline unstructured workflow. 15: real-time unstructured workflow. 16: online unstructured workflow. If you pass null or an empty list, the default value [0] is used.
	PipelineTypeList []*int32 `json:"PipelineTypeList,omitempty" xml:"PipelineTypeList,omitempty" type:"Repeated"`
	// Specifies whether to recursively query subdirectories. Default value: false.
	//
	// example:
	//
	// true
	Recursive *bool `json:"Recursive,omitempty" xml:"Recursive,omitempty"`
	// The schedule types used for filtering. If this parameter is left empty, no filtering is performed. Valid values: 1: periodic scheduling. 3: manual scheduling. 5: real-time scheduling. 7: online workflow.
	ScheduleTypeList []*int32 `json:"ScheduleTypeList,omitempty" xml:"ScheduleTypeList,omitempty" type:"Repeated"`
	// The submit statuses used for filtering. If this parameter is left empty, no filtering is performed. Valid values: DRAFT: draft. SUBMITTING: submitting. SUBMITTED: submitted. PUBLISHED: published.
	SubmitStatusList []*string `json:"SubmitStatusList,omitempty" xml:"SubmitStatusList,omitempty" type:"Repeated"`
	// The tag names used for filtering. If this parameter is left empty, no filtering is performed. If multiple names are specified, they are evaluated by using the OR operator.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 1233
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPipelinesRequestListCommand) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesRequestListCommand) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequestListCommand) GetCreatorList() []*string {
	return s.CreatorList
}

func (s *ListPipelinesRequestListCommand) GetDevelopOwnerList() []*string {
	return s.DevelopOwnerList
}

func (s *ListPipelinesRequestListCommand) GetDirectories() []*string {
	return s.Directories
}

func (s *ListPipelinesRequestListCommand) GetExactMatch() *bool {
	return s.ExactMatch
}

func (s *ListPipelinesRequestListCommand) GetKeywords() []*string {
	return s.Keywords
}

func (s *ListPipelinesRequestListCommand) GetNextCursor() *int64 {
	return s.NextCursor
}

func (s *ListPipelinesRequestListCommand) GetOpsOwnerList() []*string {
	return s.OpsOwnerList
}

func (s *ListPipelinesRequestListCommand) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListPipelinesRequestListCommand) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPipelinesRequestListCommand) GetPipelineTypeList() []*int32 {
	return s.PipelineTypeList
}

func (s *ListPipelinesRequestListCommand) GetRecursive() *bool {
	return s.Recursive
}

func (s *ListPipelinesRequestListCommand) GetScheduleTypeList() []*int32 {
	return s.ScheduleTypeList
}

func (s *ListPipelinesRequestListCommand) GetSubmitStatusList() []*string {
	return s.SubmitStatusList
}

func (s *ListPipelinesRequestListCommand) GetTagList() []*string {
	return s.TagList
}

func (s *ListPipelinesRequestListCommand) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPipelinesRequestListCommand) SetCreatorList(v []*string) *ListPipelinesRequestListCommand {
	s.CreatorList = v
	return s
}

func (s *ListPipelinesRequestListCommand) SetDevelopOwnerList(v []*string) *ListPipelinesRequestListCommand {
	s.DevelopOwnerList = v
	return s
}

func (s *ListPipelinesRequestListCommand) SetDirectories(v []*string) *ListPipelinesRequestListCommand {
	s.Directories = v
	return s
}

func (s *ListPipelinesRequestListCommand) SetExactMatch(v bool) *ListPipelinesRequestListCommand {
	s.ExactMatch = &v
	return s
}

func (s *ListPipelinesRequestListCommand) SetKeywords(v []*string) *ListPipelinesRequestListCommand {
	s.Keywords = v
	return s
}

func (s *ListPipelinesRequestListCommand) SetNextCursor(v int64) *ListPipelinesRequestListCommand {
	s.NextCursor = &v
	return s
}

func (s *ListPipelinesRequestListCommand) SetOpsOwnerList(v []*string) *ListPipelinesRequestListCommand {
	s.OpsOwnerList = v
	return s
}

func (s *ListPipelinesRequestListCommand) SetPageNum(v int32) *ListPipelinesRequestListCommand {
	s.PageNum = &v
	return s
}

func (s *ListPipelinesRequestListCommand) SetPageSize(v int32) *ListPipelinesRequestListCommand {
	s.PageSize = &v
	return s
}

func (s *ListPipelinesRequestListCommand) SetPipelineTypeList(v []*int32) *ListPipelinesRequestListCommand {
	s.PipelineTypeList = v
	return s
}

func (s *ListPipelinesRequestListCommand) SetRecursive(v bool) *ListPipelinesRequestListCommand {
	s.Recursive = &v
	return s
}

func (s *ListPipelinesRequestListCommand) SetScheduleTypeList(v []*int32) *ListPipelinesRequestListCommand {
	s.ScheduleTypeList = v
	return s
}

func (s *ListPipelinesRequestListCommand) SetSubmitStatusList(v []*string) *ListPipelinesRequestListCommand {
	s.SubmitStatusList = v
	return s
}

func (s *ListPipelinesRequestListCommand) SetTagList(v []*string) *ListPipelinesRequestListCommand {
	s.TagList = v
	return s
}

func (s *ListPipelinesRequestListCommand) SetTotalCount(v int32) *ListPipelinesRequestListCommand {
	s.TotalCount = &v
	return s
}

func (s *ListPipelinesRequestListCommand) Validate() error {
	return dara.Validate(s)
}
