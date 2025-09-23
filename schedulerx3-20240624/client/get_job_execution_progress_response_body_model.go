// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobExecutionProgressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetJobExecutionProgressResponseBody
	GetCode() *int32
	SetData(v *GetJobExecutionProgressResponseBodyData) *GetJobExecutionProgressResponseBody
	GetData() *GetJobExecutionProgressResponseBodyData
	SetMessage(v string) *GetJobExecutionProgressResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJobExecutionProgressResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetJobExecutionProgressResponseBody
	GetSuccess() *bool
}

type GetJobExecutionProgressResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetJobExecutionProgressResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter format error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9B57FDD7-ABBE-5030-B348-86EB9943DB59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetJobExecutionProgressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetJobExecutionProgressResponseBody) GetData() *GetJobExecutionProgressResponseBodyData {
	return s.Data
}

func (s *GetJobExecutionProgressResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJobExecutionProgressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobExecutionProgressResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJobExecutionProgressResponseBody) SetCode(v int32) *GetJobExecutionProgressResponseBody {
	s.Code = &v
	return s
}

func (s *GetJobExecutionProgressResponseBody) SetData(v *GetJobExecutionProgressResponseBodyData) *GetJobExecutionProgressResponseBody {
	s.Data = v
	return s
}

func (s *GetJobExecutionProgressResponseBody) SetMessage(v string) *GetJobExecutionProgressResponseBody {
	s.Message = &v
	return s
}

func (s *GetJobExecutionProgressResponseBody) SetRequestId(v string) *GetJobExecutionProgressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobExecutionProgressResponseBody) SetSuccess(v bool) *GetJobExecutionProgressResponseBody {
	s.Success = &v
	return s
}

func (s *GetJobExecutionProgressResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetJobExecutionProgressResponseBodyData struct {
	// example:
	//
	// 1758594961000
	EndTime          *string                                                    `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	JobDescription   *string                                                    `json:"JobDescription,omitempty" xml:"JobDescription,omitempty"`
	RootProgress     *GetJobExecutionProgressResponseBodyDataRootProgress       `json:"RootProgress,omitempty" xml:"RootProgress,omitempty" type:"Struct"`
	ShardingProgress []*GetJobExecutionProgressResponseBodyDataShardingProgress `json:"ShardingProgress,omitempty" xml:"ShardingProgress,omitempty" type:"Repeated"`
	// example:
	//
	// 1758506761000
	StartTime      *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskProgress   []*GetJobExecutionProgressResponseBodyDataTaskProgress   `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty" type:"Repeated"`
	TotalProgress  *GetJobExecutionProgressResponseBodyDataTotalProgress    `json:"TotalProgress,omitempty" xml:"TotalProgress,omitempty" type:"Struct"`
	WorkerProgress []*GetJobExecutionProgressResponseBodyDataWorkerProgress `json:"WorkerProgress,omitempty" xml:"WorkerProgress,omitempty" type:"Repeated"`
}

func (s GetJobExecutionProgressResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *GetJobExecutionProgressResponseBodyData) GetJobDescription() *string {
	return s.JobDescription
}

func (s *GetJobExecutionProgressResponseBodyData) GetRootProgress() *GetJobExecutionProgressResponseBodyDataRootProgress {
	return s.RootProgress
}

func (s *GetJobExecutionProgressResponseBodyData) GetShardingProgress() []*GetJobExecutionProgressResponseBodyDataShardingProgress {
	return s.ShardingProgress
}

func (s *GetJobExecutionProgressResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *GetJobExecutionProgressResponseBodyData) GetTaskProgress() []*GetJobExecutionProgressResponseBodyDataTaskProgress {
	return s.TaskProgress
}

func (s *GetJobExecutionProgressResponseBodyData) GetTotalProgress() *GetJobExecutionProgressResponseBodyDataTotalProgress {
	return s.TotalProgress
}

func (s *GetJobExecutionProgressResponseBodyData) GetWorkerProgress() []*GetJobExecutionProgressResponseBodyDataWorkerProgress {
	return s.WorkerProgress
}

func (s *GetJobExecutionProgressResponseBodyData) SetEndTime(v string) *GetJobExecutionProgressResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetJobDescription(v string) *GetJobExecutionProgressResponseBodyData {
	s.JobDescription = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetRootProgress(v *GetJobExecutionProgressResponseBodyDataRootProgress) *GetJobExecutionProgressResponseBodyData {
	s.RootProgress = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetShardingProgress(v []*GetJobExecutionProgressResponseBodyDataShardingProgress) *GetJobExecutionProgressResponseBodyData {
	s.ShardingProgress = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetStartTime(v string) *GetJobExecutionProgressResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetTaskProgress(v []*GetJobExecutionProgressResponseBodyDataTaskProgress) *GetJobExecutionProgressResponseBodyData {
	s.TaskProgress = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetTotalProgress(v *GetJobExecutionProgressResponseBodyDataTotalProgress) *GetJobExecutionProgressResponseBodyData {
	s.TotalProgress = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) SetWorkerProgress(v []*GetJobExecutionProgressResponseBodyDataWorkerProgress) *GetJobExecutionProgressResponseBodyData {
	s.WorkerProgress = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetJobExecutionProgressResponseBodyDataRootProgress struct {
	// example:
	//
	// 2
	Finished *int64 `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataRootProgress) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataRootProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataRootProgress) GetFinished() *int64 {
	return s.Finished
}

func (s *GetJobExecutionProgressResponseBodyDataRootProgress) GetTotal() *int64 {
	return s.Total
}

func (s *GetJobExecutionProgressResponseBodyDataRootProgress) SetFinished(v int64) *GetJobExecutionProgressResponseBodyDataRootProgress {
	s.Finished = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataRootProgress) SetTotal(v int64) *GetJobExecutionProgressResponseBodyDataRootProgress {
	s.Total = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataRootProgress) Validate() error {
	return dara.Validate(s)
}

type GetJobExecutionProgressResponseBodyDataShardingProgress struct {
	// id
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1306189481388277762
	JobExecutionId *string `json:"JobExecutionId,omitempty" xml:"JobExecutionId,omitempty"`
	// example:
	//
	// 2,4,6,8,10
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 5
	Status     *int32                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusType *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType `json:"StatusType,omitempty" xml:"StatusType,omitempty" type:"Struct"`
	// example:
	//
	// http://192.168.1.9:9999/
	WorkerAddr *string `json:"WorkerAddr,omitempty" xml:"WorkerAddr,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataShardingProgress) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataShardingProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) GetId() *int64 {
	return s.Id
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) GetJobExecutionId() *string {
	return s.JobExecutionId
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) GetResult() *string {
	return s.Result
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) GetStatus() *int32 {
	return s.Status
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) GetStatusType() *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType {
	return s.StatusType
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) GetWorkerAddr() *string {
	return s.WorkerAddr
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetId(v int64) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.Id = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetJobExecutionId(v string) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.JobExecutionId = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetResult(v string) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.Result = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetStatus(v int32) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.Status = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetStatusType(v *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.StatusType = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) SetWorkerAddr(v string) *GetJobExecutionProgressResponseBodyDataShardingProgress {
	s.WorkerAddr = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgress) Validate() error {
	return dara.Validate(s)
}

type GetJobExecutionProgressResponseBodyDataShardingProgressStatusType struct {
	// example:
	//
	// 5
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// TaskStatus.FAILED
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// -
	Tips map[string]*string `json:"Tips,omitempty" xml:"Tips,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) GetCode() *string {
	return s.Code
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) GetName() *string {
	return s.Name
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) GetTips() map[string]*string {
	return s.Tips
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) SetCode(v string) *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType {
	s.Code = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) SetName(v string) *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType {
	s.Name = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) SetTips(v map[string]*string) *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType {
	s.Tips = v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataShardingProgressStatusType) Validate() error {
	return dara.Validate(s)
}

type GetJobExecutionProgressResponseBodyDataTaskProgress struct {
	// example:
	//
	// 100
	Failed *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// example:
	//
	// calendar_test_2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	Pulled *int32 `json:"Pulled,omitempty" xml:"Pulled,omitempty"`
	// example:
	//
	// 100
	Queue *int32 `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// example:
	//
	// 1
	Running *int32 `json:"Running,omitempty" xml:"Running,omitempty"`
	// example:
	//
	// 100
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1000
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataTaskProgress) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataTaskProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) GetFailed() *int32 {
	return s.Failed
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) GetName() *string {
	return s.Name
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) GetPulled() *int32 {
	return s.Pulled
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) GetQueue() *int32 {
	return s.Queue
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) GetRunning() *int32 {
	return s.Running
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) GetSuccess() *int32 {
	return s.Success
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) GetTotal() *int32 {
	return s.Total
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetFailed(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Failed = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetName(v string) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Name = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetPulled(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Pulled = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetQueue(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Queue = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetRunning(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Running = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetSuccess(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Success = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) SetTotal(v int32) *GetJobExecutionProgressResponseBodyDataTaskProgress {
	s.Total = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTaskProgress) Validate() error {
	return dara.Validate(s)
}

type GetJobExecutionProgressResponseBodyDataTotalProgress struct {
	// example:
	//
	// 15
	Finished *int64 `json:"Finished,omitempty" xml:"Finished,omitempty"`
	// example:
	//
	// 30
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataTotalProgress) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataTotalProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataTotalProgress) GetFinished() *int64 {
	return s.Finished
}

func (s *GetJobExecutionProgressResponseBodyDataTotalProgress) GetTotal() *int64 {
	return s.Total
}

func (s *GetJobExecutionProgressResponseBodyDataTotalProgress) SetFinished(v int64) *GetJobExecutionProgressResponseBodyDataTotalProgress {
	s.Finished = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTotalProgress) SetTotal(v int64) *GetJobExecutionProgressResponseBodyDataTotalProgress {
	s.Total = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataTotalProgress) Validate() error {
	return dara.Validate(s)
}

type GetJobExecutionProgressResponseBodyDataWorkerProgress struct {
	// example:
	//
	// 20
	Failed *int32 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// example:
	//
	// 20
	Pulled *int32 `json:"Pulled,omitempty" xml:"Pulled,omitempty"`
	// example:
	//
	// 20
	Queue *int32 `json:"Queue,omitempty" xml:"Queue,omitempty"`
	// example:
	//
	// 20
	Running *int32 `json:"Running,omitempty" xml:"Running,omitempty"`
	// example:
	//
	// 20
	Success *int32 `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 1a0e97fb17244665327205402dbd6d
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// example:
	//
	// 10.10.116.53:61941
	WorkerAddr *string `json:"WorkerAddr,omitempty" xml:"WorkerAddr,omitempty"`
}

func (s GetJobExecutionProgressResponseBodyDataWorkerProgress) String() string {
	return dara.Prettify(s)
}

func (s GetJobExecutionProgressResponseBodyDataWorkerProgress) GoString() string {
	return s.String()
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) GetFailed() *int32 {
	return s.Failed
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) GetPulled() *int32 {
	return s.Pulled
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) GetQueue() *int32 {
	return s.Queue
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) GetRunning() *int32 {
	return s.Running
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) GetSuccess() *int32 {
	return s.Success
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) GetTotal() *int32 {
	return s.Total
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) GetTraceId() *string {
	return s.TraceId
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) GetWorkerAddr() *string {
	return s.WorkerAddr
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetFailed(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Failed = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetPulled(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Pulled = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetQueue(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Queue = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetRunning(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Running = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetSuccess(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Success = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetTotal(v int32) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.Total = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetTraceId(v string) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.TraceId = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) SetWorkerAddr(v string) *GetJobExecutionProgressResponseBodyDataWorkerProgress {
	s.WorkerAddr = &v
	return s
}

func (s *GetJobExecutionProgressResponseBodyDataWorkerProgress) Validate() error {
	return dara.Validate(s)
}
