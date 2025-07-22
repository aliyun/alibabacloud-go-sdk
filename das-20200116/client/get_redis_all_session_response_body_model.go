// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRedisAllSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetRedisAllSessionResponseBody
	GetCode() *int64
	SetData(v *GetRedisAllSessionResponseBodyData) *GetRedisAllSessionResponseBody
	GetData() *GetRedisAllSessionResponseBodyData
	SetMessage(v string) *GetRedisAllSessionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRedisAllSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRedisAllSessionResponseBody
	GetSuccess() *bool
}

type GetRedisAllSessionResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The session data.
	Data *GetRedisAllSessionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 40C6E9AF-6C23-5614-AA83-34344CC6****
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

func (s GetRedisAllSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRedisAllSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRedisAllSessionResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetRedisAllSessionResponseBody) GetData() *GetRedisAllSessionResponseBodyData {
	return s.Data
}

func (s *GetRedisAllSessionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRedisAllSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRedisAllSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRedisAllSessionResponseBody) SetCode(v int64) *GetRedisAllSessionResponseBody {
	s.Code = &v
	return s
}

func (s *GetRedisAllSessionResponseBody) SetData(v *GetRedisAllSessionResponseBodyData) *GetRedisAllSessionResponseBody {
	s.Data = v
	return s
}

func (s *GetRedisAllSessionResponseBody) SetMessage(v string) *GetRedisAllSessionResponseBody {
	s.Message = &v
	return s
}

func (s *GetRedisAllSessionResponseBody) SetRequestId(v string) *GetRedisAllSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRedisAllSessionResponseBody) SetSuccess(v bool) *GetRedisAllSessionResponseBody {
	s.Success = &v
	return s
}

func (s *GetRedisAllSessionResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRedisAllSessionResponseBodyData struct {
	// The information about the sessions.
	Sessions []*GetRedisAllSessionResponseBodyDataSessions `json:"Sessions,omitempty" xml:"Sessions,omitempty" type:"Repeated"`
	// The statistics on the access source.
	SourceStats []*GetRedisAllSessionResponseBodyDataSourceStats `json:"SourceStats,omitempty" xml:"SourceStats,omitempty" type:"Repeated"`
	// The time when the instance sessions were returned. The value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1660100753556
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The total number of sessions.
	//
	// example:
	//
	// 2
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetRedisAllSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRedisAllSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRedisAllSessionResponseBodyData) GetSessions() []*GetRedisAllSessionResponseBodyDataSessions {
	return s.Sessions
}

func (s *GetRedisAllSessionResponseBodyData) GetSourceStats() []*GetRedisAllSessionResponseBodyDataSourceStats {
	return s.SourceStats
}

func (s *GetRedisAllSessionResponseBodyData) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetRedisAllSessionResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetRedisAllSessionResponseBodyData) SetSessions(v []*GetRedisAllSessionResponseBodyDataSessions) *GetRedisAllSessionResponseBodyData {
	s.Sessions = v
	return s
}

func (s *GetRedisAllSessionResponseBodyData) SetSourceStats(v []*GetRedisAllSessionResponseBodyDataSourceStats) *GetRedisAllSessionResponseBodyData {
	s.SourceStats = v
	return s
}

func (s *GetRedisAllSessionResponseBodyData) SetTimestamp(v int64) *GetRedisAllSessionResponseBodyData {
	s.Timestamp = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyData) SetTotal(v int64) *GetRedisAllSessionResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetRedisAllSessionResponseBodyDataSessions struct {
	// The IP address and port number of the client.
	//
	// example:
	//
	// 172.16.XX.XX:53458
	Addr *string `json:"Addr,omitempty" xml:"Addr,omitempty"`
	// The connection duration of the session. Unit: seconds.
	//
	// example:
	//
	// 12
	Age *string `json:"Age,omitempty" xml:"Age,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 172.16.XX.XX
	Client *string `json:"Client,omitempty" xml:"Client,omitempty"`
	// The alias of the client.
	//
	// example:
	//
	// prod ip
	ClientDesc *string `json:"ClientDesc,omitempty" xml:"ClientDesc,omitempty"`
	// The command that was last run.
	//
	// example:
	//
	// PING
	Cmd *string `json:"Cmd,omitempty" xml:"Cmd,omitempty"`
	// The ID of the database that the client is using.
	//
	// example:
	//
	// 0
	Db *int64 `json:"Db,omitempty" xml:"Db,omitempty"`
	// The file descriptor event. Valid values:
	//
	// 	- **r**: Client sockets are readable in the event loop.
	//
	// 	- **w**: Client sockets are writable in the event loop.
	//
	// example:
	//
	// r
	Events *string `json:"Events,omitempty" xml:"Events,omitempty"`
	// The file descriptor that is used by sockets.
	//
	// example:
	//
	// 73
	Fd *int64 `json:"Fd,omitempty" xml:"Fd,omitempty"`
	// The client flag. Valid values:
	//
	// 	- **A**: The connection needs to be closed at the earliest opportunity.
	//
	// 	- **b**: The client is waiting for blocked events.
	//
	// 	- **c**: The connection is closed after all replies are written.
	//
	// 	- **d**: The monitored keys have been modified, and the `EXEC` command is about to fail.
	//
	// 	- **i**: The client is waiting for VM I/O operations. This value is no longer used.
	//
	// 	- **M**: The client is the primary node.
	//
	// 	- **N**: No special flags are configured.
	//
	// 	- **O**: The client is in monitor mode.
	//
	// 	- **r**: The client is a cluster node in read-only mode.
	//
	// 	- **S**: The client is a replica node in normal mode.
	//
	// 	- **u**: The client is not blocked.
	//
	// 	- **U**: The client is connected by using UNIX domain sockets.
	//
	// 	- **x**: The client is executing a transaction.
	//
	// example:
	//
	// N
	Flags *string `json:"Flags,omitempty" xml:"Flags,omitempty"`
	// The client ID.
	//
	// example:
	//
	// 9080586
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The duration during which the session is in the idle state. Unit: seconds.
	//
	// example:
	//
	// 8
	Idle *int64 `json:"Idle,omitempty" xml:"Idle,omitempty"`
	// The number of commands in `MULTI` or `EXEC`.
	//
	// example:
	//
	// -1
	Multi *int64 `json:"Multi,omitempty" xml:"Multi,omitempty"`
	// The name of the client.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node ID.
	//
	// example:
	//
	// r-2zemyfd1sh1u2i****-proxy-14#1679****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The size of the fixed output buffer. Unit: bytes.
	//
	// example:
	//
	// 0
	Obl *int64 `json:"Obl,omitempty" xml:"Obl,omitempty"`
	// The number of objects contained in the output list.
	//
	// example:
	//
	// 0
	Oll *int64 `json:"Oll,omitempty" xml:"Oll,omitempty"`
	// The size of the output buffer. Unit: bytes.
	//
	// example:
	//
	// 0
	Omem *int64 `json:"Omem,omitempty" xml:"Omem,omitempty"`
	// The number of subscriptions that match the pattern.
	//
	// example:
	//
	// 0
	Psub *int64 `json:"Psub,omitempty" xml:"Psub,omitempty"`
	// The size of the input buffer. Unit: bytes.
	//
	// example:
	//
	// 0
	Qbuf *int64 `json:"Qbuf,omitempty" xml:"Qbuf,omitempty"`
	// The remaining size of the input buffer. Unit: bytes.
	//
	// example:
	//
	// 0
	QbufFree *int64 `json:"QbufFree,omitempty" xml:"QbufFree,omitempty"`
	// The number of subscribed channels.
	//
	// example:
	//
	// 0
	Sub *int64 `json:"Sub,omitempty" xml:"Sub,omitempty"`
}

func (s GetRedisAllSessionResponseBodyDataSessions) String() string {
	return dara.Prettify(s)
}

func (s GetRedisAllSessionResponseBodyDataSessions) GoString() string {
	return s.String()
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetAddr() *string {
	return s.Addr
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetAge() *string {
	return s.Age
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetClient() *string {
	return s.Client
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetClientDesc() *string {
	return s.ClientDesc
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetCmd() *string {
	return s.Cmd
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetDb() *int64 {
	return s.Db
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetEvents() *string {
	return s.Events
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetFd() *int64 {
	return s.Fd
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetFlags() *string {
	return s.Flags
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetId() *int64 {
	return s.Id
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetIdle() *int64 {
	return s.Idle
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetMulti() *int64 {
	return s.Multi
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetName() *string {
	return s.Name
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetNodeId() *string {
	return s.NodeId
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetObl() *int64 {
	return s.Obl
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetOll() *int64 {
	return s.Oll
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetOmem() *int64 {
	return s.Omem
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetPsub() *int64 {
	return s.Psub
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetQbuf() *int64 {
	return s.Qbuf
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetQbufFree() *int64 {
	return s.QbufFree
}

func (s *GetRedisAllSessionResponseBodyDataSessions) GetSub() *int64 {
	return s.Sub
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetAddr(v string) *GetRedisAllSessionResponseBodyDataSessions {
	s.Addr = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetAge(v string) *GetRedisAllSessionResponseBodyDataSessions {
	s.Age = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetClient(v string) *GetRedisAllSessionResponseBodyDataSessions {
	s.Client = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetClientDesc(v string) *GetRedisAllSessionResponseBodyDataSessions {
	s.ClientDesc = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetCmd(v string) *GetRedisAllSessionResponseBodyDataSessions {
	s.Cmd = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetDb(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Db = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetEvents(v string) *GetRedisAllSessionResponseBodyDataSessions {
	s.Events = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetFd(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Fd = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetFlags(v string) *GetRedisAllSessionResponseBodyDataSessions {
	s.Flags = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetId(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Id = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetIdle(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Idle = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetMulti(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Multi = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetName(v string) *GetRedisAllSessionResponseBodyDataSessions {
	s.Name = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetNodeId(v string) *GetRedisAllSessionResponseBodyDataSessions {
	s.NodeId = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetObl(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Obl = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetOll(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Oll = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetOmem(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Omem = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetPsub(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Psub = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetQbuf(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Qbuf = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetQbufFree(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.QbufFree = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) SetSub(v int64) *GetRedisAllSessionResponseBodyDataSessions {
	s.Sub = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSessions) Validate() error {
	return dara.Validate(s)
}

type GetRedisAllSessionResponseBodyDataSourceStats struct {
	// The total number of sessions from the access source.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// The client IDs.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	// The access source.
	//
	// example:
	//
	// 172.16.XX.XX
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s GetRedisAllSessionResponseBodyDataSourceStats) String() string {
	return dara.Prettify(s)
}

func (s GetRedisAllSessionResponseBodyDataSourceStats) GoString() string {
	return s.String()
}

func (s *GetRedisAllSessionResponseBodyDataSourceStats) GetCount() *string {
	return s.Count
}

func (s *GetRedisAllSessionResponseBodyDataSourceStats) GetIds() []*int64 {
	return s.Ids
}

func (s *GetRedisAllSessionResponseBodyDataSourceStats) GetKey() *string {
	return s.Key
}

func (s *GetRedisAllSessionResponseBodyDataSourceStats) SetCount(v string) *GetRedisAllSessionResponseBodyDataSourceStats {
	s.Count = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSourceStats) SetIds(v []*int64) *GetRedisAllSessionResponseBodyDataSourceStats {
	s.Ids = v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSourceStats) SetKey(v string) *GetRedisAllSessionResponseBodyDataSourceStats {
	s.Key = &v
	return s
}

func (s *GetRedisAllSessionResponseBodyDataSourceStats) Validate() error {
	return dara.Validate(s)
}
