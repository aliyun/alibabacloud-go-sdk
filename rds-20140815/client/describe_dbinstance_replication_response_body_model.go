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
	ReplicationIp           *string `json:"ReplicationIp,omitempty" xml:"ReplicationIp,omitempty"`
	ReplicationPort         *string `json:"ReplicationPort,omitempty" xml:"ReplicationPort,omitempty"`
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

func (s *DescribeDBInstanceReplicationResponseBody) SetExternalReplication(v string) *DescribeDBInstanceReplicationResponseBody {
	s.ExternalReplication = &v
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

func (s *DescribeDBInstanceReplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
