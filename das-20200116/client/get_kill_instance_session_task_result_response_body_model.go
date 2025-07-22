// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKillInstanceSessionTaskResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetKillInstanceSessionTaskResultResponseBody
	GetCode() *int64
	SetData(v *GetKillInstanceSessionTaskResultResponseBodyData) *GetKillInstanceSessionTaskResultResponseBody
	GetData() *GetKillInstanceSessionTaskResultResponseBodyData
	SetMessage(v string) *GetKillInstanceSessionTaskResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKillInstanceSessionTaskResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKillInstanceSessionTaskResultResponseBody
	GetSuccess() *bool
}

type GetKillInstanceSessionTaskResultResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetKillInstanceSessionTaskResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, Successful is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKillInstanceSessionTaskResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKillInstanceSessionTaskResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetKillInstanceSessionTaskResultResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetKillInstanceSessionTaskResultResponseBody) GetData() *GetKillInstanceSessionTaskResultResponseBodyData {
	return s.Data
}

func (s *GetKillInstanceSessionTaskResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKillInstanceSessionTaskResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKillInstanceSessionTaskResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKillInstanceSessionTaskResultResponseBody) SetCode(v int64) *GetKillInstanceSessionTaskResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBody) SetData(v *GetKillInstanceSessionTaskResultResponseBodyData) *GetKillInstanceSessionTaskResultResponseBody {
	s.Data = v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBody) SetMessage(v string) *GetKillInstanceSessionTaskResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBody) SetRequestId(v string) *GetKillInstanceSessionTaskResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBody) SetSuccess(v bool) *GetKillInstanceSessionTaskResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetKillInstanceSessionTaskResultResponseBodyData struct {
	// The number of ignored sessions, including sessions of the accounts that are specified by IgnoredUsers, sessions of internal O\\&M accounts of Alibaba Cloud, and **Binlog Dump*	- sessions.
	//
	// example:
	//
	// 9
	IgnoredUserSessionCount *int64 `json:"IgnoredUserSessionCount,omitempty" xml:"IgnoredUserSessionCount,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze1jdv45i7l6****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of sessions that failed to be terminated.
	//
	// example:
	//
	// 0
	KillFailCount *int64 `json:"KillFailCount,omitempty" xml:"KillFailCount,omitempty"`
	// The number of sessions that were terminated.
	//
	// example:
	//
	// 100
	KillSuccessCount *int64 `json:"KillSuccessCount,omitempty" xml:"KillSuccessCount,omitempty"`
	// The node ID.
	//
	// >  This parameter is returned only if the instance is a PolarDB for MySQL cluster.
	//
	// example:
	//
	// pi-bp1h12rv501cv****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The details of the task that terminated sessions.
	Result []*GetKillInstanceSessionTaskResultResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The session IDs.
	//
	// >  If all sessions are terminated, the IDs of all sessions on the instance or node are returned.
	Sessions []*int64 `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// The task ID.
	//
	// example:
	//
	// f77d535b45405bd462b21caa3ee8****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The state of the task that terminates sessions.
	//
	// 	- **RUNNING**: The task is in progress.
	//
	// 	- **SUCCESS**: The task is successful.
	//
	// 	- **FAILURE**: The task failed.
	//
	// 	- **ERROR**: Other errors occur.
	//
	// example:
	//
	// SUCCESS
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 164882191396****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetKillInstanceSessionTaskResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKillInstanceSessionTaskResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetIgnoredUserSessionCount() *int64 {
	return s.IgnoredUserSessionCount
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetKillFailCount() *int64 {
	return s.KillFailCount
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetKillSuccessCount() *int64 {
	return s.KillSuccessCount
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetResult() []*GetKillInstanceSessionTaskResultResponseBodyDataResult {
	return s.Result
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetSessions() []*int64 {
	return s.Sessions
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetTaskState() *string {
	return s.TaskState
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetIgnoredUserSessionCount(v int64) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.IgnoredUserSessionCount = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetInstanceId(v string) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetKillFailCount(v int64) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.KillFailCount = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetKillSuccessCount(v int64) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.KillSuccessCount = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetNodeId(v string) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetResult(v []*GetKillInstanceSessionTaskResultResponseBodyDataResult) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.Result = v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetSessions(v []*int64) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.Sessions = v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetTaskId(v string) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetTaskState(v string) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.TaskState = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) SetUserId(v string) *GetKillInstanceSessionTaskResultResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetKillInstanceSessionTaskResultResponseBodyDataResult struct {
	// Indicates whether the session is active.
	//
	// > If the type of the command is Query or Execute and the session in the transaction is not terminated, the session is active.
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The type of the command executed in the session.
	//
	// example:
	//
	// Sleep
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// dbTest
	Db *string `json:"Db,omitempty" xml:"Db,omitempty"`
	// The IP address and port number of the host that initiated the session.
	//
	// example:
	//
	// 100.104.XX.XX:23428
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 8357518
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The SQL statement executed in the session.
	//
	// example:
	//
	// SELECT sleep(60)
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The description of the session when the session was terminated.
	//
	// 	- **SESSION_KILLED**: The session is terminated.
	//
	// 	- **SESSION_EXPIRED**: The session has expired.
	//
	// 	- **SESSION_NO_PERMISSION**: The account used to terminate the session has insufficient permissions.
	//
	// 	- **SESSION_ACCOUNT_ERROR**: The account or password used to terminate the session is invalid.
	//
	// 	- **SESSION_IGNORED_USER**: The session of the account does not need to be terminated.
	//
	// 	- **SESSION_INTERNAL_USER_OR_COMMAND**: The session is a session initiated by or a command run by an Alibaba Cloud O\\&M account.
	//
	// 	- **SESSION_KILL_TASK_TIMEOUT**: Timeout occurs when the session is terminated.
	//
	// 	- **SESSION_OTHER_ERROR**: Other errors occurred.
	//
	// example:
	//
	// SESSION_KILLED
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The status of the session.
	//
	// example:
	//
	// Sending data
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The ID of the subtask that terminates the session.
	//
	// example:
	//
	// task_d9e94107-6116-4ac3-b874-81466aff****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The execution duration. Unit: seconds.
	//
	// example:
	//
	// 1
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The account of the database.
	//
	// example:
	//
	// testUser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GetKillInstanceSessionTaskResultResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetKillInstanceSessionTaskResultResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetActive() *bool {
	return s.Active
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetCommand() *string {
	return s.Command
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetDb() *string {
	return s.Db
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetHost() *string {
	return s.Host
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetId() *int64 {
	return s.Id
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetInfo() *string {
	return s.Info
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetReason() *string {
	return s.Reason
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetState() *string {
	return s.State
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetTaskId() *string {
	return s.TaskId
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetTime() *int64 {
	return s.Time
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) GetUser() *string {
	return s.User
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetActive(v bool) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.Active = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetCommand(v string) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.Command = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetDb(v string) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.Db = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetHost(v string) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.Host = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetId(v int64) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetInfo(v string) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.Info = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetReason(v string) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.Reason = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetState(v string) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.State = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetTaskId(v string) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.TaskId = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetTime(v int64) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.Time = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) SetUser(v string) *GetKillInstanceSessionTaskResultResponseBodyDataResult {
	s.User = &v
	return s
}

func (s *GetKillInstanceSessionTaskResultResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
