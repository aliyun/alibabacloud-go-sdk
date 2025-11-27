// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBInstanceTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDBInstanceTopologyResponseBody
	GetCode() *string
	SetData(v *GetDBInstanceTopologyResponseBodyData) *GetDBInstanceTopologyResponseBody
	GetData() *GetDBInstanceTopologyResponseBodyData
	SetMessage(v string) *GetDBInstanceTopologyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDBInstanceTopologyResponseBody
	GetRequestId() *string
}

type GetDBInstanceTopologyResponseBody struct {
	// An internal parameter. You can ignore this parameter.
	//
	// example:
	//
	// None
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details about the topology.
	Data *GetDBInstanceTopologyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// An internal parameter. You can ignore this parameter.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7430AB1A-6D49-5B6D-B9E5-920250076074
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDBInstanceTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDBInstanceTopologyResponseBody) GetData() *GetDBInstanceTopologyResponseBodyData {
	return s.Data
}

func (s *GetDBInstanceTopologyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDBInstanceTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDBInstanceTopologyResponseBody) SetCode(v string) *GetDBInstanceTopologyResponseBody {
	s.Code = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBody) SetData(v *GetDBInstanceTopologyResponseBodyData) *GetDBInstanceTopologyResponseBody {
	s.Data = v
	return s
}

func (s *GetDBInstanceTopologyResponseBody) SetMessage(v string) *GetDBInstanceTopologyResponseBody {
	s.Message = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBody) SetRequestId(v string) *GetDBInstanceTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDBInstanceTopologyResponseBodyData struct {
	// The network connection information of the instance.
	Connections []*GetDBInstanceTopologyResponseBodyDataConnections `json:"Connections,omitempty" xml:"Connections,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// example:
	//
	// rm-m5ezban**********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The queried nodes.
	Nodes []*GetDBInstanceTopologyResponseBodyDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s GetDBInstanceTopologyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceTopologyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponseBodyData) GetConnections() []*GetDBInstanceTopologyResponseBodyDataConnections {
	return s.Connections
}

func (s *GetDBInstanceTopologyResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *GetDBInstanceTopologyResponseBodyData) GetNodes() []*GetDBInstanceTopologyResponseBodyDataNodes {
	return s.Nodes
}

func (s *GetDBInstanceTopologyResponseBodyData) SetConnections(v []*GetDBInstanceTopologyResponseBodyDataConnections) *GetDBInstanceTopologyResponseBodyData {
	s.Connections = v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyData) SetDBInstanceName(v string) *GetDBInstanceTopologyResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyData) SetNodes(v []*GetDBInstanceTopologyResponseBodyDataNodes) *GetDBInstanceTopologyResponseBodyData {
	s.Nodes = v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyData) Validate() error {
	if s.Connections != nil {
		for _, item := range s.Connections {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDBInstanceTopologyResponseBodyDataConnections struct {
	// The endpoint that is used to connect to the database instance.
	//
	// example:
	//
	// rm-m5ezban**********.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-m5ezban**********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// 	- **vpc**
	//
	// 	- **public**
	//
	// example:
	//
	// vpc
	NetType *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-qingdao-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetDBInstanceTopologyResponseBodyDataConnections) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceTopologyResponseBodyDataConnections) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) GetNetType() *string {
	return s.NetType
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) SetConnectionString(v string) *GetDBInstanceTopologyResponseBodyDataConnections {
	s.ConnectionString = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) SetDBInstanceName(v string) *GetDBInstanceTopologyResponseBodyDataConnections {
	s.DBInstanceName = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) SetNetType(v string) *GetDBInstanceTopologyResponseBodyDataConnections {
	s.NetType = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) SetZoneId(v string) *GetDBInstanceTopologyResponseBodyDataConnections {
	s.ZoneId = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataConnections) Validate() error {
	return dara.Validate(s)
}

type GetDBInstanceTopologyResponseBodyDataNodes struct {
	// The ID of the instance.
	//
	// example:
	//
	// rm-m5ezban**********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the dedicated cluster.
	//
	// > : If the instance does not reside in the specified dedicated cluster, no value is returned.
	//
	// example:
	//
	// dhg-4n*****
	DedicatedHostGroupId *string `json:"DedicatedHostGroupId,omitempty" xml:"DedicatedHostGroupId,omitempty"`
	// The host ID of the instance in the dedicated cluster.
	//
	// > : If the instance does not reside in the specified dedicated cluster, no value is returned.
	//
	// example:
	//
	// i-bpxxxxxxx
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The ID of the instance.
	//
	// > : The value \\*\\*-1\\*\\	- is returned for an instance that does not reside in a dedicated cluster.
	//
	// example:
	//
	// 349054
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The type of the node. The following result is returned:
	//
	// 	- **Master**: a primary node
	//
	// 	- **Slave**: a secondary node
	//
	// example:
	//
	// master
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The zone ID of the instance.
	//
	// example:
	//
	// cn-qingdao-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetDBInstanceTopologyResponseBodyDataNodes) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceTopologyResponseBodyDataNodes) GoString() string {
	return s.String()
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) GetDedicatedHostGroupId() *string {
	return s.DedicatedHostGroupId
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) GetRole() *string {
	return s.Role
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetDBInstanceName(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.DBInstanceName = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetDedicatedHostGroupId(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.DedicatedHostGroupId = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetDedicatedHostId(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.DedicatedHostId = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetNodeId(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.NodeId = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetRole(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.Role = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) SetZoneId(v string) *GetDBInstanceTopologyResponseBodyDataNodes {
	s.ZoneId = &v
	return s
}

func (s *GetDBInstanceTopologyResponseBodyDataNodes) Validate() error {
	return dara.Validate(s)
}
