// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMongoDBCurrentOpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetMongoDBCurrentOpResponseBody
	GetCode() *int64
	SetData(v *GetMongoDBCurrentOpResponseBodyData) *GetMongoDBCurrentOpResponseBody
	GetData() *GetMongoDBCurrentOpResponseBodyData
	SetMessage(v string) *GetMongoDBCurrentOpResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMongoDBCurrentOpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMongoDBCurrentOpResponseBody
	GetSuccess() *bool
}

type GetMongoDBCurrentOpResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the sessions.
	Data *GetMongoDBCurrentOpResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. Otherwise, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FC6C0929-29E1-59FD-8DFE-70D9D41E****
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

func (s GetMongoDBCurrentOpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMongoDBCurrentOpResponseBody) GoString() string {
	return s.String()
}

func (s *GetMongoDBCurrentOpResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetMongoDBCurrentOpResponseBody) GetData() *GetMongoDBCurrentOpResponseBodyData {
	return s.Data
}

func (s *GetMongoDBCurrentOpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMongoDBCurrentOpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMongoDBCurrentOpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMongoDBCurrentOpResponseBody) SetCode(v int64) *GetMongoDBCurrentOpResponseBody {
	s.Code = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBody) SetData(v *GetMongoDBCurrentOpResponseBodyData) *GetMongoDBCurrentOpResponseBody {
	s.Data = v
	return s
}

func (s *GetMongoDBCurrentOpResponseBody) SetMessage(v string) *GetMongoDBCurrentOpResponseBody {
	s.Message = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBody) SetRequestId(v string) *GetMongoDBCurrentOpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBody) SetSuccess(v bool) *GetMongoDBCurrentOpResponseBody {
	s.Success = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMongoDBCurrentOpResponseBodyData struct {
	// The sessions.
	SessionList []*GetMongoDBCurrentOpResponseBodyDataSessionList `json:"SessionList,omitempty" xml:"SessionList,omitempty" type:"Repeated"`
	// The statistics on the sessions.
	SessionStat *GetMongoDBCurrentOpResponseBodyDataSessionStat `json:"SessionStat,omitempty" xml:"SessionStat,omitempty" type:"Struct"`
	// The time when the database sessions were returned. The value is in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1692029584428
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s GetMongoDBCurrentOpResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMongoDBCurrentOpResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMongoDBCurrentOpResponseBodyData) GetSessionList() []*GetMongoDBCurrentOpResponseBodyDataSessionList {
	return s.SessionList
}

func (s *GetMongoDBCurrentOpResponseBodyData) GetSessionStat() *GetMongoDBCurrentOpResponseBodyDataSessionStat {
	return s.SessionStat
}

func (s *GetMongoDBCurrentOpResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetMongoDBCurrentOpResponseBodyData) SetSessionList(v []*GetMongoDBCurrentOpResponseBodyDataSessionList) *GetMongoDBCurrentOpResponseBodyData {
	s.SessionList = v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyData) SetSessionStat(v *GetMongoDBCurrentOpResponseBodyDataSessionStat) *GetMongoDBCurrentOpResponseBodyData {
	s.SessionStat = v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyData) SetTimestamp(v int64) *GetMongoDBCurrentOpResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyData) Validate() error {
	if s.SessionList != nil {
		for _, item := range s.SessionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SessionStat != nil {
		if err := s.SessionStat.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMongoDBCurrentOpResponseBodyDataSessionList struct {
	// Indicates whether the operation is active. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 219.143.177.4:52324
	Client *string `json:"Client,omitempty" xml:"Client,omitempty"`
	// The document that contains the complete command object associated with the operation.
	//
	// example:
	//
	// "command" : {
	//
	//   "find" : "items",
	//
	//   "filter" : {
	//
	//     "sku" : 1403978
	//
	//   },
	//
	//   ...
	//
	//   "$db" : "test"
	//
	// }
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The connection ID.
	//
	// example:
	//
	// 66378736
	ConnectionId *int64 `json:"ConnectionId,omitempty" xml:"ConnectionId,omitempty"`
	// The description of the connection.
	//
	// example:
	//
	// conn1013858
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The driver for MongoDB.
	//
	// example:
	//
	// mongo-java-driver|legacy@3.11.2
	Driver *string `json:"Driver,omitempty" xml:"Driver,omitempty"`
	// The host.
	//
	// example:
	//
	// a79******.cloud.et15:3328
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Indicates whether the operation is marked as terminated.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	KillPending *bool `json:"KillPending,omitempty" xml:"KillPending,omitempty"`
	// The namespace.
	//
	// example:
	//
	// admin.cmd
	Ns *string `json:"Ns,omitempty" xml:"Ns,omitempty"`
	// The type of the operation.
	//
	// example:
	//
	// update
	Op *string `json:"Op,omitempty" xml:"Op,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// 14508
	OpId *string `json:"OpId,omitempty" xml:"OpId,omitempty"`
	// The architecture of the operating system.
	//
	// example:
	//
	// amd64
	OsArch *string `json:"OsArch,omitempty" xml:"OsArch,omitempty"`
	// The name of the operating system.
	//
	// example:
	//
	// Linux
	OsName *string `json:"OsName,omitempty" xml:"OsName,omitempty"`
	// The type of the operating system.
	//
	// example:
	//
	// Linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The description of the execution plan.
	//
	// example:
	//
	// None
	PlanSummary *string `json:"PlanSummary,omitempty" xml:"PlanSummary,omitempty"`
	// The platform.
	//
	// example:
	//
	// Java/Alibaba/1.8.0_152-b5
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The duration of the operation. Unit: seconds.
	//
	// example:
	//
	// 5
	SecsRunning *int64 `json:"SecsRunning,omitempty" xml:"SecsRunning,omitempty"`
	// The ID of the data shard.
	//
	// >  This parameter is returned for sharded cluster instances.
	//
	// example:
	//
	// d-bp1689995b78****
	Shard *string `json:"Shard,omitempty" xml:"Shard,omitempty"`
}

func (s GetMongoDBCurrentOpResponseBodyDataSessionList) String() string {
	return dara.Prettify(s)
}

func (s GetMongoDBCurrentOpResponseBodyDataSessionList) GoString() string {
	return s.String()
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetActive() *bool {
	return s.Active
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetClient() *string {
	return s.Client
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetCommand() *string {
	return s.Command
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetConnectionId() *int64 {
	return s.ConnectionId
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetDesc() *string {
	return s.Desc
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetDriver() *string {
	return s.Driver
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetHost() *string {
	return s.Host
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetKillPending() *bool {
	return s.KillPending
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetNs() *string {
	return s.Ns
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetOp() *string {
	return s.Op
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetOpId() *string {
	return s.OpId
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetOsArch() *string {
	return s.OsArch
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetOsName() *string {
	return s.OsName
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetOsType() *string {
	return s.OsType
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetPlanSummary() *string {
	return s.PlanSummary
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetPlatform() *string {
	return s.Platform
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetSecsRunning() *int64 {
	return s.SecsRunning
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) GetShard() *string {
	return s.Shard
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetActive(v bool) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Active = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetClient(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Client = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetCommand(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Command = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetConnectionId(v int64) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.ConnectionId = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetDesc(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Desc = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetDriver(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Driver = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetHost(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Host = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetKillPending(v bool) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.KillPending = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetNs(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Ns = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetOp(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Op = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetOpId(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.OpId = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetOsArch(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.OsArch = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetOsName(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.OsName = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetOsType(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.OsType = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetPlanSummary(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.PlanSummary = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetPlatform(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Platform = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetSecsRunning(v int64) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.SecsRunning = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) SetShard(v string) *GetMongoDBCurrentOpResponseBodyDataSessionList {
	s.Shard = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionList) Validate() error {
	return dara.Validate(s)
}

type GetMongoDBCurrentOpResponseBodyDataSessionStat struct {
	// The number of active sessions.
	//
	// example:
	//
	// 0
	ActiveCount *int64 `json:"ActiveCount,omitempty" xml:"ActiveCount,omitempty"`
	// The statistics on the IP addresses of the clients.
	ClientStats map[string]*DataSessionStatClientStatsValue `json:"ClientStats,omitempty" xml:"ClientStats,omitempty"`
	// The statistics on the namespaces.
	DbStats map[string]*DataSessionStatDbStatsValue `json:"DbStats,omitempty" xml:"DbStats,omitempty"`
	// The longest duration of a session. Unit: seconds.
	//
	// example:
	//
	// 0
	LongestSecsRunning *int64 `json:"LongestSecsRunning,omitempty" xml:"LongestSecsRunning,omitempty"`
	// The total number of sessions.
	//
	// example:
	//
	// 55
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetMongoDBCurrentOpResponseBodyDataSessionStat) String() string {
	return dara.Prettify(s)
}

func (s GetMongoDBCurrentOpResponseBodyDataSessionStat) GoString() string {
	return s.String()
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) GetActiveCount() *int64 {
	return s.ActiveCount
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) GetClientStats() map[string]*DataSessionStatClientStatsValue {
	return s.ClientStats
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) GetDbStats() map[string]*DataSessionStatDbStatsValue {
	return s.DbStats
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) GetLongestSecsRunning() *int64 {
	return s.LongestSecsRunning
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) SetActiveCount(v int64) *GetMongoDBCurrentOpResponseBodyDataSessionStat {
	s.ActiveCount = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) SetClientStats(v map[string]*DataSessionStatClientStatsValue) *GetMongoDBCurrentOpResponseBodyDataSessionStat {
	s.ClientStats = v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) SetDbStats(v map[string]*DataSessionStatDbStatsValue) *GetMongoDBCurrentOpResponseBodyDataSessionStat {
	s.DbStats = v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) SetLongestSecsRunning(v int64) *GetMongoDBCurrentOpResponseBodyDataSessionStat {
	s.LongestSecsRunning = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) SetTotalCount(v int64) *GetMongoDBCurrentOpResponseBodyDataSessionStat {
	s.TotalCount = &v
	return s
}

func (s *GetMongoDBCurrentOpResponseBodyDataSessionStat) Validate() error {
	return dara.Validate(s)
}
