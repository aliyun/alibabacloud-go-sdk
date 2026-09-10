// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBwmMigrationTaskWriterWorkflowListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetBwmMigrationTaskWriterWorkflowListResponseBodyData) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetData() []*GetBwmMigrationTaskWriterWorkflowListResponseBodyData
	SetEmpty(v bool) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetEmpty() *bool
	SetErrCode(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetErrMessage() *string
	SetNotEmpty(v bool) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetNotEmpty() *bool
	SetPageIndex(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetTotalCount() *int32
	SetTotalPages(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBody
	GetTotalPages() *int32
}

type GetBwmMigrationTaskWriterWorkflowListResponseBody struct {
	// The response data.
	Data []*GetBwmMigrationTaskWriterWorkflowListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Indicates whether the result is empty.
	Empty *bool `json:"empty,omitempty" xml:"empty,omitempty"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Indicates whether the result is not empty.
	NotEmpty *bool `json:"notEmpty,omitempty" xml:"notEmpty,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4C467B38-3910-4477-9B0B-6963D83B4E72
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure. If the call fails, check errCode and errMessage for details.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 10
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s GetBwmMigrationTaskWriterWorkflowListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationTaskWriterWorkflowListResponseBody) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetData() []*GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	return s.Data
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetEmpty() *bool {
	return s.Empty
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetNotEmpty() *bool {
	return s.NotEmpty
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetData(v []*GetBwmMigrationTaskWriterWorkflowListResponseBodyData) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.Data = v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetEmpty(v bool) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.Empty = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetErrCode(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetErrMessage(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetNotEmpty(v bool) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.NotEmpty = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetPageIndex(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetPageSize(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetRequestId(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetSuccess(v bool) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.Success = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetTotalCount(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) SetTotalPages(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBwmMigrationTaskWriterWorkflowListResponseBodyData struct {
	// The cron expression.
	//
	// example:
	//
	// 	- 	- 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The database primary key ID.
	//
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// The submit failure error message.
	//
	// example:
	//
	// Submit failed
	SubmitDetail *string `json:"submitDetail,omitempty" xml:"submitDetail,omitempty"`
	// Filter by status. Valid values:
	//
	// - WRT_INIT: Submit not started.
	//
	// - WRT_RUN: Submitting.
	//
	// - WRT_SUCC: All submitted successfully.
	//
	// - WRT_FAIL: All submissions failed.
	//
	// - WRT_PART_FAIL: Some submissions failed.
	//
	// - DPY_SUCC: Published successfully.
	//
	// - DPY_FAIL: Publish failed.
	//
	// example:
	//
	// WRT_INIT
	SubmitStatus *string `json:"submitStatus,omitempty" xml:"submitStatus,omitempty"`
	// The workflow ID written to the target side.
	//
	// example:
	//
	// target_wf-12345
	TargetWorkflowId *string `json:"targetWorkflowId,omitempty" xml:"targetWorkflowId,omitempty"`
	// The workflow name on the target side.
	//
	// example:
	//
	// target_workflow_name
	TargetWorkflowName *string `json:"targetWorkflowName,omitempty" xml:"targetWorkflowName,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 5
	TaskNodeCount *int32 `json:"taskNodeCount,omitempty" xml:"taskNodeCount,omitempty"`
	// The actual workflow ID.
	//
	// example:
	//
	// wf-12345
	WorkflowId *string `json:"workflowId,omitempty" xml:"workflowId,omitempty"`
	// The workflow name.
	//
	// example:
	//
	// workflow_name
	WorkflowName *string `json:"workflowName,omitempty" xml:"workflowName,omitempty"`
}

func (s GetBwmMigrationTaskWriterWorkflowListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GetCron() *string {
	return s.Cron
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GetId() *int32 {
	return s.Id
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GetSubmitDetail() *string {
	return s.SubmitDetail
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GetSubmitStatus() *string {
	return s.SubmitStatus
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GetTargetWorkflowId() *string {
	return s.TargetWorkflowId
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GetTargetWorkflowName() *string {
	return s.TargetWorkflowName
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GetTaskNodeCount() *int32 {
	return s.TaskNodeCount
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GetWorkflowId() *string {
	return s.WorkflowId
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) GetWorkflowName() *string {
	return s.WorkflowName
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) SetCron(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	s.Cron = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) SetId(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) SetSubmitDetail(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	s.SubmitDetail = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) SetSubmitStatus(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	s.SubmitStatus = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) SetTargetWorkflowId(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	s.TargetWorkflowId = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) SetTargetWorkflowName(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	s.TargetWorkflowName = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) SetTaskNodeCount(v int32) *GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	s.TaskNodeCount = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) SetWorkflowId(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	s.WorkflowId = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) SetWorkflowName(v string) *GetBwmMigrationTaskWriterWorkflowListResponseBodyData {
	s.WorkflowName = &v
	return s
}

func (s *GetBwmMigrationTaskWriterWorkflowListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
