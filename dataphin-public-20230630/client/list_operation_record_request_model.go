// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListCommand(v *ListOperationRecordRequestListCommand) *ListOperationRecordRequest
	GetListCommand() *ListOperationRecordRequestListCommand
	SetOpTenantId(v int64) *ListOperationRecordRequest
	GetOpTenantId() *int64
}

type ListOperationRecordRequest struct {
	// The query command.
	//
	// This parameter is required.
	ListCommand *ListOperationRecordRequestListCommand `json:"ListCommand,omitempty" xml:"ListCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListOperationRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationRecordRequest) GoString() string {
	return s.String()
}

func (s *ListOperationRecordRequest) GetListCommand() *ListOperationRecordRequestListCommand {
	return s.ListCommand
}

func (s *ListOperationRecordRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListOperationRecordRequest) SetListCommand(v *ListOperationRecordRequestListCommand) *ListOperationRecordRequest {
	s.ListCommand = v
	return s
}

func (s *ListOperationRecordRequest) SetOpTenantId(v int64) *ListOperationRecordRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListOperationRecordRequest) Validate() error {
	if s.ListCommand != nil {
		if err := s.ListCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListOperationRecordRequestListCommand struct {
	// The end of the start time range. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-12-31 23:59:59
	BeginTimeEnd *string `json:"BeginTimeEnd,omitempty" xml:"BeginTimeEnd,omitempty"`
	// The beginning of the start time range. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-01-01 00:00:00
	BeginTimeStart *string `json:"BeginTimeStart,omitempty" xml:"BeginTimeStart,omitempty"`
	// The keyword for code search.
	//
	// example:
	//
	// select
	CodeContent *string `json:"CodeContent,omitempty" xml:"CodeContent,omitempty"`
	// The list of code types.
	CodeType []*int32 `json:"CodeType,omitempty" xml:"CodeType,omitempty" type:"Repeated"`
	// The list of duration ranges.
	Duration []*int32 `json:"Duration,omitempty" xml:"Duration,omitempty" type:"Repeated"`
	// The script name.
	//
	// example:
	//
	// 测试脚本
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The list of object types.
	ObjectType []*string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 131211211
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sort type. Valid values: 0 (start time ascending), 1 (start time descending), 2 (object name).
	//
	// example:
	//
	// 1
	SortType *int32 `json:"SortType,omitempty" xml:"SortType,omitempty"`
	// The list of task statuses.
	Status []*int32 `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The list of executor IDs.
	UserIds []*string `json:"UserIds,omitempty" xml:"UserIds,omitempty" type:"Repeated"`
}

func (s ListOperationRecordRequestListCommand) String() string {
	return dara.Prettify(s)
}

func (s ListOperationRecordRequestListCommand) GoString() string {
	return s.String()
}

func (s *ListOperationRecordRequestListCommand) GetBeginTimeEnd() *string {
	return s.BeginTimeEnd
}

func (s *ListOperationRecordRequestListCommand) GetBeginTimeStart() *string {
	return s.BeginTimeStart
}

func (s *ListOperationRecordRequestListCommand) GetCodeContent() *string {
	return s.CodeContent
}

func (s *ListOperationRecordRequestListCommand) GetCodeType() []*int32 {
	return s.CodeType
}

func (s *ListOperationRecordRequestListCommand) GetDuration() []*int32 {
	return s.Duration
}

func (s *ListOperationRecordRequestListCommand) GetFileName() *string {
	return s.FileName
}

func (s *ListOperationRecordRequestListCommand) GetObjectType() []*string {
	return s.ObjectType
}

func (s *ListOperationRecordRequestListCommand) GetPage() *int32 {
	return s.Page
}

func (s *ListOperationRecordRequestListCommand) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOperationRecordRequestListCommand) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListOperationRecordRequestListCommand) GetSortType() *int32 {
	return s.SortType
}

func (s *ListOperationRecordRequestListCommand) GetStatus() []*int32 {
	return s.Status
}

func (s *ListOperationRecordRequestListCommand) GetUserIds() []*string {
	return s.UserIds
}

func (s *ListOperationRecordRequestListCommand) SetBeginTimeEnd(v string) *ListOperationRecordRequestListCommand {
	s.BeginTimeEnd = &v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetBeginTimeStart(v string) *ListOperationRecordRequestListCommand {
	s.BeginTimeStart = &v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetCodeContent(v string) *ListOperationRecordRequestListCommand {
	s.CodeContent = &v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetCodeType(v []*int32) *ListOperationRecordRequestListCommand {
	s.CodeType = v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetDuration(v []*int32) *ListOperationRecordRequestListCommand {
	s.Duration = v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetFileName(v string) *ListOperationRecordRequestListCommand {
	s.FileName = &v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetObjectType(v []*string) *ListOperationRecordRequestListCommand {
	s.ObjectType = v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetPage(v int32) *ListOperationRecordRequestListCommand {
	s.Page = &v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetPageSize(v int32) *ListOperationRecordRequestListCommand {
	s.PageSize = &v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetProjectId(v int64) *ListOperationRecordRequestListCommand {
	s.ProjectId = &v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetSortType(v int32) *ListOperationRecordRequestListCommand {
	s.SortType = &v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetStatus(v []*int32) *ListOperationRecordRequestListCommand {
	s.Status = v
	return s
}

func (s *ListOperationRecordRequestListCommand) SetUserIds(v []*string) *ListOperationRecordRequestListCommand {
	s.UserIds = v
	return s
}

func (s *ListOperationRecordRequestListCommand) Validate() error {
	return dara.Validate(s)
}
