// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMySQLAllSessionAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetMySQLAllSessionAsyncResponseBody
	GetCode() *int64
	SetData(v *GetMySQLAllSessionAsyncResponseBodyData) *GetMySQLAllSessionAsyncResponseBody
	GetData() *GetMySQLAllSessionAsyncResponseBodyData
	SetMessage(v string) *GetMySQLAllSessionAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMySQLAllSessionAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMySQLAllSessionAsyncResponseBody
	GetSuccess() *bool
}

type GetMySQLAllSessionAsyncResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetMySQLAllSessionAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
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

func (s GetMySQLAllSessionAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMySQLAllSessionAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *GetMySQLAllSessionAsyncResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetMySQLAllSessionAsyncResponseBody) GetData() *GetMySQLAllSessionAsyncResponseBodyData {
	return s.Data
}

func (s *GetMySQLAllSessionAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMySQLAllSessionAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMySQLAllSessionAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMySQLAllSessionAsyncResponseBody) SetCode(v int64) *GetMySQLAllSessionAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBody) SetData(v *GetMySQLAllSessionAsyncResponseBodyData) *GetMySQLAllSessionAsyncResponseBody {
	s.Data = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBody) SetMessage(v string) *GetMySQLAllSessionAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBody) SetRequestId(v string) *GetMySQLAllSessionAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBody) SetSuccess(v bool) *GetMySQLAllSessionAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMySQLAllSessionAsyncResponseBodyData struct {
	// Indicates whether the asynchronous request was complete. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Complete *bool `json:"Complete,omitempty" xml:"Complete,omitempty"`
	// Indicates whether the asynchronous request failed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Fail *bool `json:"Fail,omitempty" xml:"Fail,omitempty"`
	// Indicates whether the asynchronous request was complete. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsFinish *bool `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	// The ID of the asynchronous request.
	//
	// example:
	//
	// async__507044db6c4eadfa2dab9b084e80****
	ResultId *string `json:"ResultId,omitempty" xml:"ResultId,omitempty"`
	// The session data.
	SessionData *GetMySQLAllSessionAsyncResponseBodyDataSessionData `json:"SessionData,omitempty" xml:"SessionData,omitempty" type:"Struct"`
	// The state of the asynchronous request. Valid values:
	//
	// 	- **RUNNING**
	//
	// 	- **SUCCESS**
	//
	// 	- **FAIL**
	//
	// example:
	//
	// SUCCESS
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the asynchronous request was made. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1660100753556
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMySQLAllSessionAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMySQLAllSessionAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) GetComplete() *bool {
	return s.Complete
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) GetFail() *bool {
	return s.Fail
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) GetResultId() *string {
	return s.ResultId
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) GetSessionData() *GetMySQLAllSessionAsyncResponseBodyDataSessionData {
	return s.SessionData
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) GetState() *string {
	return s.State
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) SetComplete(v bool) *GetMySQLAllSessionAsyncResponseBodyData {
	s.Complete = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) SetFail(v bool) *GetMySQLAllSessionAsyncResponseBodyData {
	s.Fail = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) SetIsFinish(v bool) *GetMySQLAllSessionAsyncResponseBodyData {
	s.IsFinish = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) SetResultId(v string) *GetMySQLAllSessionAsyncResponseBodyData {
	s.ResultId = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) SetSessionData(v *GetMySQLAllSessionAsyncResponseBodyDataSessionData) *GetMySQLAllSessionAsyncResponseBodyData {
	s.SessionData = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) SetState(v string) *GetMySQLAllSessionAsyncResponseBodyData {
	s.State = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) SetTimestamp(v int64) *GetMySQLAllSessionAsyncResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetMySQLAllSessionAsyncResponseBodyDataSessionData struct {
	// The total number of active sessions.
	//
	// example:
	//
	// 10
	ActiveSessionCount *int64 `json:"ActiveSessionCount,omitempty" xml:"ActiveSessionCount,omitempty"`
	// The sessions that are counted by client IP address.
	ClientStats []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats `json:"ClientStats,omitempty" xml:"ClientStats,omitempty" type:"Repeated"`
	// The sessions that are counted by database.
	DbStats []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats `json:"DbStats,omitempty" xml:"DbStats,omitempty" type:"Repeated"`
	// The maximum execution duration of an active session. Unit: seconds.
	//
	// example:
	//
	// 6
	MaxActiveTime *int64 `json:"MaxActiveTime,omitempty" xml:"MaxActiveTime,omitempty"`
	// The sessions.
	SessionList []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList `json:"SessionList,omitempty" xml:"SessionList,omitempty" type:"Repeated"`
	// The time when the session was queried. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1659581514000020
	TimeStamp *int64 `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
	// The total number of sessions.
	//
	// example:
	//
	// 988
	TotalSessionCount *int64 `json:"TotalSessionCount,omitempty" xml:"TotalSessionCount,omitempty"`
	// The sessions that are counted by database account.
	UserStats []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats `json:"UserStats,omitempty" xml:"UserStats,omitempty" type:"Repeated"`
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionData) String() string {
	return dara.Prettify(s)
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionData) GoString() string {
	return s.String()
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) GetActiveSessionCount() *int64 {
	return s.ActiveSessionCount
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) GetClientStats() []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats {
	return s.ClientStats
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) GetDbStats() []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats {
	return s.DbStats
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) GetMaxActiveTime() *int64 {
	return s.MaxActiveTime
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) GetSessionList() []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	return s.SessionList
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) GetTimeStamp() *int64 {
	return s.TimeStamp
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) GetTotalSessionCount() *int64 {
	return s.TotalSessionCount
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) GetUserStats() []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats {
	return s.UserStats
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) SetActiveSessionCount(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionData {
	s.ActiveSessionCount = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) SetClientStats(v []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) *GetMySQLAllSessionAsyncResponseBodyDataSessionData {
	s.ClientStats = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) SetDbStats(v []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) *GetMySQLAllSessionAsyncResponseBodyDataSessionData {
	s.DbStats = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) SetMaxActiveTime(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionData {
	s.MaxActiveTime = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) SetSessionList(v []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) *GetMySQLAllSessionAsyncResponseBodyDataSessionData {
	s.SessionList = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) SetTimeStamp(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionData {
	s.TimeStamp = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) SetTotalSessionCount(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionData {
	s.TotalSessionCount = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) SetUserStats(v []*GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) *GetMySQLAllSessionAsyncResponseBodyDataSessionData {
	s.UserStats = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionData) Validate() error {
	return dara.Validate(s)
}

type GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats struct {
	// The number of active sessions that belong to the client IP address.
	//
	// >  If the type of the command executed in the session is Query or Execute and the session in the transaction is not terminated, the session is active.
	//
	// example:
	//
	// 1
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 47.100.XX.XX
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The IDs of the sessions that belong to the client IP address.
	ThreadIdList []*int64 `json:"ThreadIdList,omitempty" xml:"ThreadIdList,omitempty" type:"Repeated"`
	// The total number of sessions that belong to the client IP address.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The database accounts to which the sessions belong.
	UserList []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) String() string {
	return dara.Prettify(s)
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) GoString() string {
	return s.String()
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) GetActiveCount() *int64 {
	return s.ActiveCount
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) GetKey() *string {
	return s.Key
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) GetThreadIdList() []*int64 {
	return s.ThreadIdList
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) GetUserList() []*string {
	return s.UserList
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) SetActiveCount(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats {
	s.ActiveCount = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) SetKey(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats {
	s.Key = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) SetThreadIdList(v []*int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats {
	s.ThreadIdList = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) SetTotalCount(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats {
	s.TotalCount = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) SetUserList(v []*string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats {
	s.UserList = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataClientStats) Validate() error {
	return dara.Validate(s)
}

type GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats struct {
	// The number of active sessions of the database.
	//
	// >  If the type of the command executed in the session is Query or Execute and the session in the transaction is not terminated, the session is active.
	//
	// example:
	//
	// 1
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The database name.
	//
	// example:
	//
	// dbTest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The IDs of the sessions of the database.
	ThreadIdList []*int64 `json:"ThreadIdList,omitempty" xml:"ThreadIdList,omitempty" type:"Repeated"`
	// The total number of sessions of the database.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The database accounts to which the sessions belong.
	UserList []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) String() string {
	return dara.Prettify(s)
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) GoString() string {
	return s.String()
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) GetActiveCount() *int64 {
	return s.ActiveCount
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) GetKey() *string {
	return s.Key
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) GetThreadIdList() []*int64 {
	return s.ThreadIdList
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) GetUserList() []*string {
	return s.UserList
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) SetActiveCount(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats {
	s.ActiveCount = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) SetKey(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats {
	s.Key = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) SetThreadIdList(v []*int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats {
	s.ThreadIdList = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) SetTotalCount(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats {
	s.TotalCount = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) SetUserList(v []*string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats {
	s.UserList = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataDbStats) Validate() error {
	return dara.Validate(s)
}

type GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList struct {
	// The IP address of the client.
	//
	// example:
	//
	// 47.100.XX.XX
	Client *string `json:"Client,omitempty" xml:"Client,omitempty"`
	// The type of the command executed in the session.
	//
	// example:
	//
	// Query
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The database name.
	//
	// example:
	//
	// dbTest
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 14521783
	SessionId *int64 `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The SQL template ID.
	//
	// >  This parameter is returned only when you use a PolarDB-X 2.0 instance.
	//
	// example:
	//
	// a7cac1a9
	SqlTemplateId *string `json:"SqlTemplateId,omitempty" xml:"SqlTemplateId,omitempty"`
	// The SQL statement executed in the session.
	//
	// example:
	//
	// INSERT INTO ...
	SqlText *string `json:"SqlText,omitempty" xml:"SqlText,omitempty"`
	// The status of the session.
	//
	// example:
	//
	// starting
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The execution duration of the session. Unit: seconds.
	//
	// example:
	//
	// 6
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The duration of the transaction. Unit: seconds.
	//
	// example:
	//
	// 6
	TrxDuration *int64 `json:"TrxDuration,omitempty" xml:"TrxDuration,omitempty"`
	// The ID of the transaction to which the session belongs.
	//
	// example:
	//
	// 754300775132
	TrxId *string `json:"TrxId,omitempty" xml:"TrxId,omitempty"`
	// The username of the database account.
	//
	// example:
	//
	// testUser
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The alias of the IP address of the client.
	//
	// example:
	//
	// master-shanghai
	UserClientAlias *string `json:"UserClientAlias,omitempty" xml:"UserClientAlias,omitempty"`
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) String() string {
	return dara.Prettify(s)
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GoString() string {
	return s.String()
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetClient() *string {
	return s.Client
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetCommand() *string {
	return s.Command
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetDbName() *string {
	return s.DbName
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetSessionId() *int64 {
	return s.SessionId
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetSqlTemplateId() *string {
	return s.SqlTemplateId
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetSqlText() *string {
	return s.SqlText
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetState() *string {
	return s.State
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetTime() *int64 {
	return s.Time
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetTrxDuration() *int64 {
	return s.TrxDuration
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetTrxId() *string {
	return s.TrxId
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetUser() *string {
	return s.User
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) GetUserClientAlias() *string {
	return s.UserClientAlias
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetClient(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.Client = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetCommand(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.Command = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetDbName(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.DbName = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetSessionId(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.SessionId = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetSqlTemplateId(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.SqlTemplateId = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetSqlText(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.SqlText = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetState(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.State = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetTime(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.Time = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetTrxDuration(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.TrxDuration = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetTrxId(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.TrxId = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetUser(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.User = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) SetUserClientAlias(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList {
	s.UserClientAlias = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataSessionList) Validate() error {
	return dara.Validate(s)
}

type GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats struct {
	// The number of active sessions within the account.
	//
	// >  If the type of the command executed in the session is Query or Execute and the session in the transaction is not terminated, the session is active.
	//
	// example:
	//
	// 1
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The database account.
	//
	// example:
	//
	// testUser
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The IDs of the sessions within the account.
	ThreadIdList []*int64 `json:"ThreadIdList,omitempty" xml:"ThreadIdList,omitempty" type:"Repeated"`
	// The total number of sessions within the account.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The database accounts to which the sessions belong.
	UserList []*string `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) String() string {
	return dara.Prettify(s)
}

func (s GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) GoString() string {
	return s.String()
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) GetActiveCount() *int64 {
	return s.ActiveCount
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) GetKey() *string {
	return s.Key
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) GetThreadIdList() []*int64 {
	return s.ThreadIdList
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) GetUserList() []*string {
	return s.UserList
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) SetActiveCount(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats {
	s.ActiveCount = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) SetKey(v string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats {
	s.Key = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) SetThreadIdList(v []*int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats {
	s.ThreadIdList = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) SetTotalCount(v int64) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats {
	s.TotalCount = &v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) SetUserList(v []*string) *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats {
	s.UserList = v
	return s
}

func (s *GetMySQLAllSessionAsyncResponseBodyDataSessionDataUserStats) Validate() error {
	return dara.Validate(s)
}
