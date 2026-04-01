// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceReplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExternalReplication(v string) *DescribeDBInstanceReplicationResponseBody
	GetExternalReplication() *string
	SetGtidExecuted(v string) *DescribeDBInstanceReplicationResponseBody
	GetGtidExecuted() *string
	SetImportStatus(v string) *DescribeDBInstanceReplicationResponseBody
	GetImportStatus() *string
	SetReplicationDelay(v string) *DescribeDBInstanceReplicationResponseBody
	GetReplicationDelay() *string
	SetReplicationErrorMessage(v string) *DescribeDBInstanceReplicationResponseBody
	GetReplicationErrorMessage() *string
	SetReplicationIp(v string) *DescribeDBInstanceReplicationResponseBody
	GetReplicationIp() *string
	SetReplicationPort(v string) *DescribeDBInstanceReplicationResponseBody
	GetReplicationPort() *string
	SetReplicationSource(v string) *DescribeDBInstanceReplicationResponseBody
	GetReplicationSource() *string
	SetReplicationState(v string) *DescribeDBInstanceReplicationResponseBody
	GetReplicationState() *string
	SetRequestId(v string) *DescribeDBInstanceReplicationResponseBody
	GetRequestId() *string
	SetSlaveStatusList(v []*DescribeDBInstanceReplicationResponseBodySlaveStatusList) *DescribeDBInstanceReplicationResponseBody
	GetSlaveStatusList() []*DescribeDBInstanceReplicationResponseBodySlaveStatusList
}

type DescribeDBInstanceReplicationResponseBody struct {
	// Indicates whether the native replication mods is enabled. Valid values:
	//
	// 	- **ON**
	//
	// 	- **OFF**
	//
	// example:
	//
	// ON
	ExternalReplication *string `json:"ExternalReplication,omitempty" xml:"ExternalReplication,omitempty"`
	// example:
	//
	// bd2a34b9-8b8d-11ef-8917-00163e1298b9:1-20567
	GtidExecuted *string `json:"GtidExecuted,omitempty" xml:"GtidExecuted,omitempty"`
	// COMPLETED: 导入完成，INIT: 初始化，IMPORTING: 正在导入
	//
	// example:
	//
	// COMPLETED
	ImportStatus *string `json:"ImportStatus,omitempty" xml:"ImportStatus,omitempty"`
	// The replication latency. Unit: seconds.
	//
	// example:
	//
	// 0
	ReplicationDelay *string `json:"ReplicationDelay,omitempty" xml:"ReplicationDelay,omitempty"`
	// The replication error message.
	//
	// example:
	//
	// Got fatal error 1236 from master when reading data from binary log...
	ReplicationErrorMessage *string `json:"ReplicationErrorMessage,omitempty" xml:"ReplicationErrorMessage,omitempty"`
	// example:
	//
	// 192.168.10.x
	ReplicationIp *string `json:"ReplicationIp,omitempty" xml:"ReplicationIp,omitempty"`
	// example:
	//
	// 3306
	ReplicationPort *string `json:"ReplicationPort,omitempty" xml:"ReplicationPort,omitempty"`
	// The source of the native replication.
	//
	// example:
	//
	// 192.168.x.x
	ReplicationSource *string `json:"ReplicationSource,omitempty" xml:"ReplicationSource,omitempty"`
	// The current replication status. Valid values:
	//
	// 	- **Running**
	//
	// 	- **Connecting**
	//
	// 	- **Stopped**
	//
	// 	- **Error**
	//
	// example:
	//
	// Running
	//
	// Connecting
	//
	// Stopped
	//
	// Error
	ReplicationState *string `json:"ReplicationState,omitempty" xml:"ReplicationState,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 847BA085-B377-4BFA-8267-F82345ECE1D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// [{"SlaveIoRunning": "Yes", "SlaveSqlRunning": "Yes", "SecondsBehindMaster": 0, "SlaveIoState": "Waiting for master to send event", "SlaveSqlRunningState": "Slave has read all relay log; waiting for more updates", "ExecutedGtidSet": "bd2a34b9-8b8d-11ef-8917-00163e1298b9:1-20567", "MasterHost": "192.168.10.100", "MasterUser": "repl_user", "MasterUuid": "bd2a34b9-8b8d-11ef-8917-00163e1298b9", "LastErrno": 0, "LastSqlErrno": 0, "LastIoErrno": 0, "LastSqlError": "", "LastIoError": "", "ChannelName": "my_test_channel", "ReplicateDoDb": "test_db,test_db_1", "ReplicateIgnoreDb": "information_schema,performance_schema", "ReplicateDoTable": "test_table,test_table_1", "ReplicateIgnoreTable": "temp_table,temp_table_1", "ReplicateWildDoTable": "test_table.%", "ReplicateWildIgnoreTable": "temp_table.%"}]
	SlaveStatusList []*DescribeDBInstanceReplicationResponseBodySlaveStatusList `json:"SlaveStatusList,omitempty" xml:"SlaveStatusList,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceReplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceReplicationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceReplicationResponseBody) GetExternalReplication() *string {
	return s.ExternalReplication
}

func (s *DescribeDBInstanceReplicationResponseBody) GetGtidExecuted() *string {
	return s.GtidExecuted
}

func (s *DescribeDBInstanceReplicationResponseBody) GetImportStatus() *string {
	return s.ImportStatus
}

func (s *DescribeDBInstanceReplicationResponseBody) GetReplicationDelay() *string {
	return s.ReplicationDelay
}

func (s *DescribeDBInstanceReplicationResponseBody) GetReplicationErrorMessage() *string {
	return s.ReplicationErrorMessage
}

func (s *DescribeDBInstanceReplicationResponseBody) GetReplicationIp() *string {
	return s.ReplicationIp
}

func (s *DescribeDBInstanceReplicationResponseBody) GetReplicationPort() *string {
	return s.ReplicationPort
}

func (s *DescribeDBInstanceReplicationResponseBody) GetReplicationSource() *string {
	return s.ReplicationSource
}

func (s *DescribeDBInstanceReplicationResponseBody) GetReplicationState() *string {
	return s.ReplicationState
}

func (s *DescribeDBInstanceReplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceReplicationResponseBody) GetSlaveStatusList() []*DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	return s.SlaveStatusList
}

func (s *DescribeDBInstanceReplicationResponseBody) SetExternalReplication(v string) *DescribeDBInstanceReplicationResponseBody {
	s.ExternalReplication = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetGtidExecuted(v string) *DescribeDBInstanceReplicationResponseBody {
	s.GtidExecuted = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetImportStatus(v string) *DescribeDBInstanceReplicationResponseBody {
	s.ImportStatus = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetReplicationDelay(v string) *DescribeDBInstanceReplicationResponseBody {
	s.ReplicationDelay = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetReplicationErrorMessage(v string) *DescribeDBInstanceReplicationResponseBody {
	s.ReplicationErrorMessage = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetReplicationIp(v string) *DescribeDBInstanceReplicationResponseBody {
	s.ReplicationIp = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetReplicationPort(v string) *DescribeDBInstanceReplicationResponseBody {
	s.ReplicationPort = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetReplicationSource(v string) *DescribeDBInstanceReplicationResponseBody {
	s.ReplicationSource = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetReplicationState(v string) *DescribeDBInstanceReplicationResponseBody {
	s.ReplicationState = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetRequestId(v string) *DescribeDBInstanceReplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) SetSlaveStatusList(v []*DescribeDBInstanceReplicationResponseBodySlaveStatusList) *DescribeDBInstanceReplicationResponseBody {
	s.SlaveStatusList = v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBody) Validate() error {
	if s.SlaveStatusList != nil {
		for _, item := range s.SlaveStatusList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceReplicationResponseBodySlaveStatusList struct {
	// example:
	//
	// my_test_channel
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// example:
	//
	// bd2a34b9-8b8d-11ef-8917-00163e1298b9:1-20567
	ExecutedGtidSet *string `json:"ExecutedGtidSet,omitempty" xml:"ExecutedGtidSet,omitempty"`
	// 0表示无错误，其他值表示具体的错误代码
	//
	// example:
	//
	// 0
	LastErrno *int32 `json:"LastErrno,omitempty" xml:"LastErrno,omitempty"`
	// 0表示无错误，其他值表示IO线程的错误代码
	//
	// example:
	//
	// 0
	LastIoErrno *int32 `json:"LastIoErrno,omitempty" xml:"LastIoErrno,omitempty"`
	// IO线程的错误信息描述
	LastIoError *string `json:"LastIoError,omitempty" xml:"LastIoError,omitempty"`
	// 0表示无错误，其他值表示SQL线程的错误代码
	//
	// example:
	//
	// 0
	LastSqlErrno *int32 `json:"LastSqlErrno,omitempty" xml:"LastSqlErrno,omitempty"`
	// SQL线程的错误信息描述
	LastSqlError *string `json:"LastSqlError,omitempty" xml:"LastSqlError,omitempty"`
	// example:
	//
	// 192.168.10.100
	MasterHost *string `json:"MasterHost,omitempty" xml:"MasterHost,omitempty"`
	// example:
	//
	// repl_user
	MasterUser *string `json:"MasterUser,omitempty" xml:"MasterUser,omitempty"`
	// example:
	//
	// bd2a34b9-8b8d-11ef-8917-00163e1298b9
	MasterUuid *string `json:"MasterUuid,omitempty" xml:"MasterUuid,omitempty"`
	// example:
	//
	// test_db,test_db_1
	ReplicateDoDb *string `json:"ReplicateDoDb,omitempty" xml:"ReplicateDoDb,omitempty"`
	// example:
	//
	// test_table,test_table_1
	ReplicateDoTable *string `json:"ReplicateDoTable,omitempty" xml:"ReplicateDoTable,omitempty"`
	// example:
	//
	// information_schema,performance_schema
	ReplicateIgnoreDb *string `json:"ReplicateIgnoreDb,omitempty" xml:"ReplicateIgnoreDb,omitempty"`
	// example:
	//
	// temp_table,temp_table_1
	ReplicateIgnoreTable *string `json:"ReplicateIgnoreTable,omitempty" xml:"ReplicateIgnoreTable,omitempty"`
	// example:
	//
	// test_table.%
	ReplicateWildDoTable *string `json:"ReplicateWildDoTable,omitempty" xml:"ReplicateWildDoTable,omitempty"`
	// example:
	//
	// temp_table.%
	ReplicateWildIgnoreTable *string `json:"ReplicateWildIgnoreTable,omitempty" xml:"ReplicateWildIgnoreTable,omitempty"`
	// example:
	//
	// 0
	SecondsBehindMaster *int32 `json:"SecondsBehindMaster,omitempty" xml:"SecondsBehindMaster,omitempty"`
	// Yes: 运行中，No: 已停止
	//
	// example:
	//
	// Yes
	SlaveIoRunning *string `json:"SlaveIoRunning,omitempty" xml:"SlaveIoRunning,omitempty"`
	// example:
	//
	// Waiting for master to send event
	SlaveIoState *string `json:"SlaveIoState,omitempty" xml:"SlaveIoState,omitempty"`
	// Yes: 运行中，No: 已停止
	//
	// example:
	//
	// Yes
	SlaveSqlRunning *string `json:"SlaveSqlRunning,omitempty" xml:"SlaveSqlRunning,omitempty"`
	// example:
	//
	// Slave has read all relay log; waiting for more updates
	SlaveSqlRunningState *string `json:"SlaveSqlRunningState,omitempty" xml:"SlaveSqlRunningState,omitempty"`
}

func (s DescribeDBInstanceReplicationResponseBodySlaveStatusList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceReplicationResponseBodySlaveStatusList) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetChannelName() *string {
	return s.ChannelName
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetExecutedGtidSet() *string {
	return s.ExecutedGtidSet
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetLastErrno() *int32 {
	return s.LastErrno
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetLastIoErrno() *int32 {
	return s.LastIoErrno
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetLastIoError() *string {
	return s.LastIoError
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetLastSqlErrno() *int32 {
	return s.LastSqlErrno
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetLastSqlError() *string {
	return s.LastSqlError
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetMasterHost() *string {
	return s.MasterHost
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetMasterUser() *string {
	return s.MasterUser
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetMasterUuid() *string {
	return s.MasterUuid
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetReplicateDoDb() *string {
	return s.ReplicateDoDb
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetReplicateDoTable() *string {
	return s.ReplicateDoTable
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetReplicateIgnoreDb() *string {
	return s.ReplicateIgnoreDb
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetReplicateIgnoreTable() *string {
	return s.ReplicateIgnoreTable
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetReplicateWildDoTable() *string {
	return s.ReplicateWildDoTable
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetReplicateWildIgnoreTable() *string {
	return s.ReplicateWildIgnoreTable
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetSecondsBehindMaster() *int32 {
	return s.SecondsBehindMaster
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetSlaveIoRunning() *string {
	return s.SlaveIoRunning
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetSlaveIoState() *string {
	return s.SlaveIoState
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetSlaveSqlRunning() *string {
	return s.SlaveSqlRunning
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) GetSlaveSqlRunningState() *string {
	return s.SlaveSqlRunningState
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetChannelName(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.ChannelName = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetExecutedGtidSet(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.ExecutedGtidSet = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetLastErrno(v int32) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.LastErrno = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetLastIoErrno(v int32) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.LastIoErrno = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetLastIoError(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.LastIoError = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetLastSqlErrno(v int32) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.LastSqlErrno = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetLastSqlError(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.LastSqlError = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetMasterHost(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.MasterHost = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetMasterUser(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.MasterUser = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetMasterUuid(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.MasterUuid = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetReplicateDoDb(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.ReplicateDoDb = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetReplicateDoTable(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.ReplicateDoTable = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetReplicateIgnoreDb(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.ReplicateIgnoreDb = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetReplicateIgnoreTable(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.ReplicateIgnoreTable = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetReplicateWildDoTable(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.ReplicateWildDoTable = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetReplicateWildIgnoreTable(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.ReplicateWildIgnoreTable = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetSecondsBehindMaster(v int32) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.SecondsBehindMaster = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetSlaveIoRunning(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.SlaveIoRunning = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetSlaveIoState(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.SlaveIoState = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetSlaveSqlRunning(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.SlaveSqlRunning = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) SetSlaveSqlRunningState(v string) *DescribeDBInstanceReplicationResponseBodySlaveStatusList {
	s.SlaveSqlRunningState = &v
	return s
}

func (s *DescribeDBInstanceReplicationResponseBodySlaveStatusList) Validate() error {
	return dara.Validate(s)
}
