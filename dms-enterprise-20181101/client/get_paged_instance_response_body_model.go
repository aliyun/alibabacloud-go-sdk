// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPagedInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPagedInstanceResponseBodyData) *GetPagedInstanceResponseBody
	GetData() *GetPagedInstanceResponseBodyData
	SetErrorCode(v string) *GetPagedInstanceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPagedInstanceResponseBody
	GetErrorMessage() *string
	SetPageIndex(v int64) *GetPagedInstanceResponseBody
	GetPageIndex() *int64
	SetPageSize(v int64) *GetPagedInstanceResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetPagedInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPagedInstanceResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *GetPagedInstanceResponseBody
	GetTotal() *int64
	SetTraceId(v string) *GetPagedInstanceResponseBody
	GetTraceId() *string
}

type GetPagedInstanceResponseBody struct {
	Data *GetPagedInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 1
	PageIndex *int64 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0a06e1e316757357507896067d3780
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s GetPagedInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPagedInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPagedInstanceResponseBody) GetData() *GetPagedInstanceResponseBodyData {
	return s.Data
}

func (s *GetPagedInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPagedInstanceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPagedInstanceResponseBody) GetPageIndex() *int64 {
	return s.PageIndex
}

func (s *GetPagedInstanceResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetPagedInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPagedInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPagedInstanceResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *GetPagedInstanceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *GetPagedInstanceResponseBody) SetData(v *GetPagedInstanceResponseBodyData) *GetPagedInstanceResponseBody {
	s.Data = v
	return s
}

func (s *GetPagedInstanceResponseBody) SetErrorCode(v string) *GetPagedInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPagedInstanceResponseBody) SetErrorMessage(v string) *GetPagedInstanceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPagedInstanceResponseBody) SetPageIndex(v int64) *GetPagedInstanceResponseBody {
	s.PageIndex = &v
	return s
}

func (s *GetPagedInstanceResponseBody) SetPageSize(v int64) *GetPagedInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetPagedInstanceResponseBody) SetRequestId(v string) *GetPagedInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPagedInstanceResponseBody) SetSuccess(v bool) *GetPagedInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *GetPagedInstanceResponseBody) SetTotal(v int64) *GetPagedInstanceResponseBody {
	s.Total = &v
	return s
}

func (s *GetPagedInstanceResponseBody) SetTraceId(v string) *GetPagedInstanceResponseBody {
	s.TraceId = &v
	return s
}

func (s *GetPagedInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPagedInstanceResponseBodyData struct {
	Instance []*GetPagedInstanceResponseBodyDataInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s GetPagedInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPagedInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPagedInstanceResponseBodyData) GetInstance() []*GetPagedInstanceResponseBodyDataInstance {
	return s.Instance
}

func (s *GetPagedInstanceResponseBodyData) SetInstance(v []*GetPagedInstanceResponseBodyDataInstance) *GetPagedInstanceResponseBodyData {
	s.Instance = v
	return s
}

func (s *GetPagedInstanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetPagedInstanceResponseBodyDataInstance struct {
	// example:
	//
	// 2023-05-14 16:00:57
	BusinessTime *string `json:"BusinessTime,omitempty" xml:"BusinessTime,omitempty"`
	// example:
	//
	// 3
	CheckStatus *int64 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// example:
	//
	// 33753
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// example:
	//
	// false
	Delete *string `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// example:
	//
	// 2023-04-23 10:23:20
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2023-03-28 10:50:45
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2023-04-18 15:28:16
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 6851
	HistoryDagId *int64 `json:"HistoryDagId,omitempty" xml:"HistoryDagId,omitempty"`
	// example:
	//
	// 24271
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// {”nodes":[11694,11695]"}
	LastRunningContext *string `json:"LastRunningContext,omitempty" xml:"LastRunningContext,omitempty"`
	Msg                *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TaskType *int64 `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 3406
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// 1
	TriggerType *int64 `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetPagedInstanceResponseBodyDataInstance) String() string {
	return dara.Prettify(s)
}

func (s GetPagedInstanceResponseBodyDataInstance) GoString() string {
	return s.String()
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetBusinessTime() *string {
	return s.BusinessTime
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetCheckStatus() *int64 {
	return s.CheckStatus
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetDagId() *int64 {
	return s.DagId
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetDelete() *string {
	return s.Delete
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetEndTime() *string {
	return s.EndTime
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetHistoryDagId() *int64 {
	return s.HistoryDagId
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetId() *int64 {
	return s.Id
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetLastRunningContext() *string {
	return s.LastRunningContext
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetMsg() *string {
	return s.Msg
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetStatus() *int64 {
	return s.Status
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetTaskType() *int64 {
	return s.TaskType
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetTenantId() *string {
	return s.TenantId
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetTriggerType() *int64 {
	return s.TriggerType
}

func (s *GetPagedInstanceResponseBodyDataInstance) GetVersion() *string {
	return s.Version
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetBusinessTime(v string) *GetPagedInstanceResponseBodyDataInstance {
	s.BusinessTime = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetCheckStatus(v int64) *GetPagedInstanceResponseBodyDataInstance {
	s.CheckStatus = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetDagId(v int64) *GetPagedInstanceResponseBodyDataInstance {
	s.DagId = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetDelete(v string) *GetPagedInstanceResponseBodyDataInstance {
	s.Delete = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetEndTime(v string) *GetPagedInstanceResponseBodyDataInstance {
	s.EndTime = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetGmtCreate(v string) *GetPagedInstanceResponseBodyDataInstance {
	s.GmtCreate = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetGmtModified(v string) *GetPagedInstanceResponseBodyDataInstance {
	s.GmtModified = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetHistoryDagId(v int64) *GetPagedInstanceResponseBodyDataInstance {
	s.HistoryDagId = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetId(v int64) *GetPagedInstanceResponseBodyDataInstance {
	s.Id = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetLastRunningContext(v string) *GetPagedInstanceResponseBodyDataInstance {
	s.LastRunningContext = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetMsg(v string) *GetPagedInstanceResponseBodyDataInstance {
	s.Msg = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetStatus(v int64) *GetPagedInstanceResponseBodyDataInstance {
	s.Status = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetTaskType(v int64) *GetPagedInstanceResponseBodyDataInstance {
	s.TaskType = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetTenantId(v string) *GetPagedInstanceResponseBodyDataInstance {
	s.TenantId = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetTriggerType(v int64) *GetPagedInstanceResponseBodyDataInstance {
	s.TriggerType = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) SetVersion(v string) *GetPagedInstanceResponseBodyDataInstance {
	s.Version = &v
	return s
}

func (s *GetPagedInstanceResponseBodyDataInstance) Validate() error {
	return dara.Validate(s)
}
