// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPipelinesResponseBody
	GetCode() *string
	SetData(v *ListPipelinesResponseBodyData) *ListPipelinesResponseBody
	GetData() *ListPipelinesResponseBodyData
	SetHttpStatusCode(v int32) *ListPipelinesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPipelinesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPipelinesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPipelinesResponseBody
	GetSuccess() *bool
}

type ListPipelinesResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The paged query result.
	Data *ListPipelinesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPipelinesResponseBody) GetData() *ListPipelinesResponseBodyData {
	return s.Data
}

func (s *ListPipelinesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPipelinesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelinesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPipelinesResponseBody) SetCode(v string) *ListPipelinesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPipelinesResponseBody) SetData(v *ListPipelinesResponseBodyData) *ListPipelinesResponseBody {
	s.Data = v
	return s
}

func (s *ListPipelinesResponseBody) SetHttpStatusCode(v int32) *ListPipelinesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPipelinesResponseBody) SetMessage(v string) *ListPipelinesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPipelinesResponseBody) SetRequestId(v string) *ListPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelinesResponseBody) SetSuccess(v bool) *ListPipelinesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelinesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPipelinesResponseBodyData struct {
	// The list of node information on the current page.
	List []*ListPipelinesResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The cursor for the next page (an opaque cursor that the caller does not need to interpret). A null value indicates that there are no more pages. Otherwise, pass this value as the nextCursor parameter in the next request to retrieve the next page.
	//
	// example:
	//
	// 123
	NextCursor *int64 `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	// The current page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records that match the conditions. For the first page request, the actual total count is returned. For subsequent page requests (when nextCursor is passed in), if totalCount is included in the request, the same value is returned. Otherwise, this field is not returned. The total value is a snapshot taken at the time of the first page query and is not updated in real time as data changes during pagination.
	//
	// example:
	//
	// 105
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListPipelinesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyData) GetList() []*ListPipelinesResponseBodyDataList {
	return s.List
}

func (s *ListPipelinesResponseBodyData) GetNextCursor() *int64 {
	return s.NextCursor
}

func (s *ListPipelinesResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListPipelinesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPipelinesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListPipelinesResponseBodyData) SetList(v []*ListPipelinesResponseBodyDataList) *ListPipelinesResponseBodyData {
	s.List = v
	return s
}

func (s *ListPipelinesResponseBodyData) SetNextCursor(v int64) *ListPipelinesResponseBodyData {
	s.NextCursor = &v
	return s
}

func (s *ListPipelinesResponseBodyData) SetPageNum(v int32) *ListPipelinesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListPipelinesResponseBodyData) SetPageSize(v int32) *ListPipelinesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPipelinesResponseBodyData) SetTotal(v int32) *ListPipelinesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListPipelinesResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelinesResponseBodyDataList struct {
	// The list of user IDs of development owners.
	DevelopOwners []*string `json:"DevelopOwners,omitempty" xml:"DevelopOwners,omitempty" type:"Repeated"`
	// The directory where the node is located.
	//
	// example:
	//
	// /dwd/finance_domain/
	Directory *string `json:"Directory,omitempty" xml:"Directory,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 12121111
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The schedule node ID.
	//
	// example:
	//
	// n_6793582765516849152
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node name.
	//
	// example:
	//
	// test
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The list of user IDs of O&M owners.
	OpsOwners []*string `json:"OpsOwners,omitempty" xml:"OpsOwners,omitempty" type:"Repeated"`
	// The pipeline ID.
	//
	// example:
	//
	// 1450811
	PipelineId *int64 `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The schedule type. Valid values:
	//
	// - 1: periodic scheduling.
	//
	// - 3: manual scheduling.
	//
	// - 5: real-time scheduling.
	//
	// example:
	//
	// 1
	ScheduleType *int32 `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The list of node tag names.
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The node status. Valid values:
	//
	// - DRAFT: draft.
	//
	// - SUBMITTING: being submitted.
	//
	// - SUBMITTED: submitted.
	//
	// - PUBLISHED: published.
	//
	// example:
	//
	// SUBMITTED
	TaskStatus *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// The node type. Valid values:
	//
	// - 0: offline integration.
	//
	// - 1: real-time integration.
	//
	// - 13: data aggregation.
	//
	// - 14: offline unstructured workflow.
	//
	// - 15: real-time unstructured workflow.
	//
	// example:
	//
	// 0
	TaskType *int32 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListPipelinesResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyDataList) GetDevelopOwners() []*string {
	return s.DevelopOwners
}

func (s *ListPipelinesResponseBodyDataList) GetDirectory() *string {
	return s.Directory
}

func (s *ListPipelinesResponseBodyDataList) GetFileId() *int64 {
	return s.FileId
}

func (s *ListPipelinesResponseBodyDataList) GetNodeId() *string {
	return s.NodeId
}

func (s *ListPipelinesResponseBodyDataList) GetNodeName() *string {
	return s.NodeName
}

func (s *ListPipelinesResponseBodyDataList) GetOpsOwners() []*string {
	return s.OpsOwners
}

func (s *ListPipelinesResponseBodyDataList) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *ListPipelinesResponseBodyDataList) GetScheduleType() *int32 {
	return s.ScheduleType
}

func (s *ListPipelinesResponseBodyDataList) GetTags() []*string {
	return s.Tags
}

func (s *ListPipelinesResponseBodyDataList) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ListPipelinesResponseBodyDataList) GetTaskType() *int32 {
	return s.TaskType
}

func (s *ListPipelinesResponseBodyDataList) SetDevelopOwners(v []*string) *ListPipelinesResponseBodyDataList {
	s.DevelopOwners = v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetDirectory(v string) *ListPipelinesResponseBodyDataList {
	s.Directory = &v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetFileId(v int64) *ListPipelinesResponseBodyDataList {
	s.FileId = &v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetNodeId(v string) *ListPipelinesResponseBodyDataList {
	s.NodeId = &v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetNodeName(v string) *ListPipelinesResponseBodyDataList {
	s.NodeName = &v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetOpsOwners(v []*string) *ListPipelinesResponseBodyDataList {
	s.OpsOwners = v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetPipelineId(v int64) *ListPipelinesResponseBodyDataList {
	s.PipelineId = &v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetScheduleType(v int32) *ListPipelinesResponseBodyDataList {
	s.ScheduleType = &v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetTags(v []*string) *ListPipelinesResponseBodyDataList {
	s.Tags = v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetTaskStatus(v string) *ListPipelinesResponseBodyDataList {
	s.TaskStatus = &v
	return s
}

func (s *ListPipelinesResponseBodyDataList) SetTaskType(v int32) *ListPipelinesResponseBodyDataList {
	s.TaskType = &v
	return s
}

func (s *ListPipelinesResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
