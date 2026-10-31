// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStreamTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetStreamTaskResponseBody
	GetCode() *string
	SetData(v *GetStreamTaskResponseBodyData) *GetStreamTaskResponseBody
	GetData() *GetStreamTaskResponseBodyData
	SetHttpStatusCode(v int32) *GetStreamTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetStreamTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetStreamTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetStreamTaskResponseBody
	GetSuccess() *bool
}

type GetStreamTaskResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The basic information and running information of the real-time task instance. The fileId can be used to subsequently query operation records and historical running instances.
	Data *GetStreamTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The details of the backend response exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStreamTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStreamTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetStreamTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetStreamTaskResponseBody) GetData() *GetStreamTaskResponseBodyData {
	return s.Data
}

func (s *GetStreamTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetStreamTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetStreamTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStreamTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStreamTaskResponseBody) SetCode(v string) *GetStreamTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetStreamTaskResponseBody) SetData(v *GetStreamTaskResponseBodyData) *GetStreamTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetStreamTaskResponseBody) SetHttpStatusCode(v int32) *GetStreamTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetStreamTaskResponseBody) SetMessage(v string) *GetStreamTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetStreamTaskResponseBody) SetRequestId(v string) *GetStreamTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStreamTaskResponseBody) SetSuccess(v bool) *GetStreamTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetStreamTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStreamTaskResponseBodyData struct {
	// The time corresponding to the latest read position. The format is yyyy-MM-dd HH:mm:ss. If no consumption metrics are available, the value is null.
	//
	// example:
	//
	// 2026-10-15 13:45:52
	ConsumeLocationDate *string `json:"ConsumeLocationDate,omitempty" xml:"ConsumeLocationDate,omitempty"`
	// The database engine version of the compute DPI engine used to run the node.
	//
	// example:
	//
	// 1.15.3
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The ID of the task file. This corresponds to the FileId request parameter in ListStreamOpsRecords and ListStreamInstanceRecords. It is not a nodeId or a template file ID.
	//
	// example:
	//
	// 7333610201479808
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the node to which the real-time task belongs.
	//
	// example:
	//
	// rt_mysql_doris
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the node. The ID starts with n_ and identifies a development node. It is different from the task file ID.
	//
	// example:
	//
	// n_8442437168737681408
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The ID of the user who last operated the task.
	//
	// example:
	//
	// 30001011
	OperatorId *string `json:"OperatorId,omitempty" xml:"OperatorId,omitempty"`
	// The display name of the user who last operated the task.
	//
	// example:
	//
	// Alice
	OperatorName *string `json:"OperatorName,omitempty" xml:"OperatorName,omitempty"`
	// The ID of the Dataphin project to which the task belongs.
	//
	// example:
	//
	// 7299592831043008
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project mode. Valid values: BASIC for the basic mode, and DEV_PROD for the development-production isolation mode.
	//
	// example:
	//
	// BASIC
	ProjectMode *string `json:"ProjectMode,omitempty" xml:"ProjectMode,omitempty"`
	// The name of the Dataphin project to which the task belongs.
	//
	// example:
	//
	// flink_k8s
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The display information of the resource queue used by the task, including the queue name and ID.
	//
	// example:
	//
	// Dedicated for real-time development (1052061)
	ResourceQueue *string `json:"ResourceQueue,omitempty" xml:"ResourceQueue,omitempty"`
	// The running status of the task. The original TaskStatusEnum enumeration name is returned. Common real-time statuses include: STREAMING_STARTING (starting), STREAMING_RUNNING (running), STREAMING_PAUSING (pausing), STREAMING_PAUSED (paused), STREAMING_RESUMING (resuming), STREAMING_STOPPING (stopping), STREAMING_STOPPED (stopped), STREAMING_INIT (initial status), STREAMING_FAILED (running exception), STREAMING_START_FAILED (start exception), STREAMING_RESUME_FAILED (resume exception), and STREAMING_FINISHED (completed). Other known statuses also return the original enumeration name. If the status cannot be identified, UNKNOWN is returned. If the original status is empty, null is returned.
	//
	// example:
	//
	// STREAMING_RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Instance ID of the real-time task instance. Instance ID starts with t_ and corresponds to the TaskId request parameter in GetStreamTask.
	//
	// example:
	//
	// t_8442437168737681408_20261014_8442447717177229312
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task, such as REAL_TIME_PIPELINE (real-time integration), FLINK_SQL (SQL task), FLINK_TEMPLATE_SQL (template task), and FLINK_DATASTREAM (DataStream task).
	//
	// example:
	//
	// REAL_TIME_PIPELINE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStreamTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStreamTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStreamTaskResponseBodyData) GetConsumeLocationDate() *string {
	return s.ConsumeLocationDate
}

func (s *GetStreamTaskResponseBodyData) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *GetStreamTaskResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *GetStreamTaskResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetStreamTaskResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *GetStreamTaskResponseBodyData) GetOperatorId() *string {
	return s.OperatorId
}

func (s *GetStreamTaskResponseBodyData) GetOperatorName() *string {
	return s.OperatorName
}

func (s *GetStreamTaskResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetStreamTaskResponseBodyData) GetProjectMode() *string {
	return s.ProjectMode
}

func (s *GetStreamTaskResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetStreamTaskResponseBodyData) GetResourceQueue() *string {
	return s.ResourceQueue
}

func (s *GetStreamTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetStreamTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetStreamTaskResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetStreamTaskResponseBodyData) SetConsumeLocationDate(v string) *GetStreamTaskResponseBodyData {
	s.ConsumeLocationDate = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetEngineVersion(v string) *GetStreamTaskResponseBodyData {
	s.EngineVersion = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetFileId(v string) *GetStreamTaskResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetName(v string) *GetStreamTaskResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetNodeId(v string) *GetStreamTaskResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetOperatorId(v string) *GetStreamTaskResponseBodyData {
	s.OperatorId = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetOperatorName(v string) *GetStreamTaskResponseBodyData {
	s.OperatorName = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetProjectId(v int64) *GetStreamTaskResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetProjectMode(v string) *GetStreamTaskResponseBodyData {
	s.ProjectMode = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetProjectName(v string) *GetStreamTaskResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetResourceQueue(v string) *GetStreamTaskResponseBodyData {
	s.ResourceQueue = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetStatus(v string) *GetStreamTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetTaskId(v string) *GetStreamTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) SetType(v string) *GetStreamTaskResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetStreamTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
