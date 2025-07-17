// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRunsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListPipelineRunsResponseBodyPagingInfo) *ListPipelineRunsResponseBody
	GetPagingInfo() *ListPipelineRunsResponseBodyPagingInfo
	SetRequestId(v string) *ListPipelineRunsResponseBody
	GetRequestId() *string
}

type ListPipelineRunsResponseBody struct {
	// The pagination information.
	PagingInfo *ListPipelineRunsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPipelineRunsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBody) GetPagingInfo() *ListPipelineRunsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListPipelineRunsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineRunsResponseBody) SetPagingInfo(v *ListPipelineRunsResponseBodyPagingInfo) *ListPipelineRunsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListPipelineRunsResponseBody) SetRequestId(v string) *ListPipelineRunsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineRunsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPipelineRunsResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The processes.
	PipelineRuns []*ListPipelineRunsResponseBodyPagingInfoPipelineRuns `json:"PipelineRuns,omitempty" xml:"PipelineRuns,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2524
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPipelineRunsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPipelineRunsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPipelineRunsResponseBodyPagingInfo) GetPipelineRuns() []*ListPipelineRunsResponseBodyPagingInfoPipelineRuns {
	return s.PipelineRuns
}

func (s *ListPipelineRunsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPipelineRunsResponseBodyPagingInfo) SetPageNumber(v int32) *ListPipelineRunsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfo) SetPageSize(v int32) *ListPipelineRunsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfo) SetPipelineRuns(v []*ListPipelineRunsResponseBodyPagingInfoPipelineRuns) *ListPipelineRunsResponseBodyPagingInfo {
	s.PipelineRuns = v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfo) SetTotalCount(v int32) *ListPipelineRunsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListPipelineRunsResponseBodyPagingInfoPipelineRuns struct {
	// The time when the process was created. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1702736654000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator of the process.
	//
	// example:
	//
	// 110755000425XXXX
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 097c73fe-ed6e-4fb1-b109-a5d59e46cd58
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message returned during the stage.
	//
	// example:
	//
	// Error message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The time when the process was modified. This value is a UNIX timestamp.
	//
	// example:
	//
	// 1702736654000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 70199
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The stages of the process.
	Stages []*ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages `json:"Stages,omitempty" xml:"Stages,omitempty" type:"Repeated"`
	// The status of the process.
	//
	// Valid values:
	//
	// 	- Init
	//
	// 	- Running
	//
	// 	- Success
	//
	// 	- Fail
	//
	// 	- Termination
	//
	// 	- Cancel
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListPipelineRunsResponseBodyPagingInfoPipelineRuns) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsResponseBodyPagingInfoPipelineRuns) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) GetCreator() *string {
	return s.Creator
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) GetId() *string {
	return s.Id
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) GetMessage() *string {
	return s.Message
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) GetStages() []*ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages {
	return s.Stages
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) GetStatus() *string {
	return s.Status
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) SetCreateTime(v int64) *ListPipelineRunsResponseBodyPagingInfoPipelineRuns {
	s.CreateTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) SetCreator(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRuns {
	s.Creator = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) SetId(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRuns {
	s.Id = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) SetMessage(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRuns {
	s.Message = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) SetModifyTime(v int64) *ListPipelineRunsResponseBodyPagingInfoPipelineRuns {
	s.ModifyTime = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) SetProjectId(v int64) *ListPipelineRunsResponseBodyPagingInfoPipelineRuns {
	s.ProjectId = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) SetStages(v []*ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) *ListPipelineRunsResponseBodyPagingInfoPipelineRuns {
	s.Stages = v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) SetStatus(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRuns {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRuns) Validate() error {
	return dara.Validate(s)
}

type ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages struct {
	// The code of the stage.
	//
	// example:
	//
	// DEV_CHECK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The description of the stage.
	//
	// example:
	//
	// Check before going online to development
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The additional information about the stage.
	Detail map[string]interface{} `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The error message returned during the stage.
	//
	// example:
	//
	// Error message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The name of the stage.
	//
	// example:
	//
	// Check before going online to development
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The status of the stage.
	//
	// Valid values:
	//
	// 	- Init
	//
	// 	- Running
	//
	// 	- Success
	//
	// 	- Fail
	//
	// 	- Termination
	//
	// 	- Cancel
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The step number of the stage.
	//
	// example:
	//
	// 1
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
	// The type of the stage. This parameter indicates the operation type in the stage.
	//
	// Valid values:
	//
	// 	- Deploy
	//
	// 	- Check
	//
	// 	- Offline
	//
	// 	- Build
	//
	// 	- Delete
	//
	// example:
	//
	// Check
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) GoString() string {
	return s.String()
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) GetCode() *string {
	return s.Code
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) GetDescription() *string {
	return s.Description
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) GetDetail() map[string]interface{} {
	return s.Detail
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) GetMessage() *string {
	return s.Message
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) GetName() *string {
	return s.Name
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) GetStatus() *string {
	return s.Status
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) GetStep() *int32 {
	return s.Step
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) GetType() *string {
	return s.Type
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) SetCode(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages {
	s.Code = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) SetDescription(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages {
	s.Description = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) SetDetail(v map[string]interface{}) *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages {
	s.Detail = v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) SetMessage(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages {
	s.Message = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) SetName(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages {
	s.Name = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) SetStatus(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages {
	s.Status = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) SetStep(v int32) *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages {
	s.Step = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) SetType(v string) *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages {
	s.Type = &v
	return s
}

func (s *ListPipelineRunsResponseBodyPagingInfoPipelineRunsStages) Validate() error {
	return dara.Validate(s)
}
