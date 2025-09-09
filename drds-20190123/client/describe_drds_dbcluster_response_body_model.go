// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDbInstance(v *DescribeDrdsDBClusterResponseBodyDbInstance) *DescribeDrdsDBClusterResponseBody
	GetDbInstance() *DescribeDrdsDBClusterResponseBodyDbInstance
	SetRequestId(v string) *DescribeDrdsDBClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsDBClusterResponseBody
	GetSuccess() *bool
}

type DescribeDrdsDBClusterResponseBody struct {
	// The details of each PolarDB cluster.
	DbInstance *DescribeDrdsDBClusterResponseBodyDbInstance `json:"DbInstance,omitempty" xml:"DbInstance,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 60A77FD6-0DE4-4A34-B6FB-9C2673******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBody) GetDbInstance() *DescribeDrdsDBClusterResponseBodyDbInstance {
	return s.DbInstance
}

func (s *DescribeDrdsDBClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsDBClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsDBClusterResponseBody) SetDbInstance(v *DescribeDrdsDBClusterResponseBodyDbInstance) *DescribeDrdsDBClusterResponseBody {
	s.DbInstance = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBody) SetRequestId(v string) *DescribeDrdsDBClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBody) SetSuccess(v bool) *DescribeDrdsDBClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDBClusterResponseBodyDbInstance struct {
	// The ID of the PolarDB cluster.
	//
	// example:
	//
	// pc-*****************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The status of the PolarDB instance.
	//
	// example:
	//
	// 1
	DBInstanceStatus *string `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty"`
	// The information about the nodes in the PolarDB Cluster.
	DBNodes *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes `json:"DBNodes,omitempty" xml:"DBNodes,omitempty" type:"Struct"`
	// The type of storage used by the DRDS database.
	//
	// example:
	//
	// POLARDB
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The endpoint of the PolarDB read /write splitting endpoint
	Endpoints *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Struct"`
	// The type of the DRDS database storage engine.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The version of the DRDS database storage engine.
	//
	// example:
	//
	// 8.0
	EngineVersion *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	// The time when the PolarDB cluster expires.
	//
	// example:
	//
	// 2019-09-27 11:22:33
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The network type of the PolarDB cluster.
	//
	// example:
	//
	// VPC
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The billing method of the PolarDB cluster.
	//
	// example:
	//
	// Postpaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The PolarDB access port.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of RDS instance. PolarDB cluster does not support this parameter.
	//
	// example:
	//
	// ignore
	RdsInstType *string `json:"RdsInstType,omitempty" xml:"RdsInstType,omitempty"`
	// This parameter specifies the Read mode when the database storage type is PolarDB.
	//
	// Valid values:
	//
	// 	- **DEFAULT**: the default mode (that is, all read traffic is sent to the PolarDB read /write node).
	//
	// 	- **CUSTOM**: Custom mode (you can customize the ratio of traffic sent to read /write nodes and read-only nodes).
	//
	// 	- **BALANCE**: read balancing mode (the read traffic is automatically distributed by the read load module of the PolarDB cluster, which can also be understood as the read traffic being evenly distributed to each node).
	//
	// example:
	//
	// CUSTOM
	ReadMode *string `json:"ReadMode,omitempty" xml:"ReadMode,omitempty"`
	// The number of days remaining on the PolarDB for MySQL instance.
	//
	// example:
	//
	// 0
	RemainDays *string `json:"RemainDays,omitempty" xml:"RemainDays,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstance) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetDBInstanceStatus() *string {
	return s.DBInstanceStatus
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetDBNodes() *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes {
	return s.DBNodes
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetEndpoints() *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints {
	return s.Endpoints
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetPayType() *string {
	return s.PayType
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetRdsInstType() *string {
	return s.RdsInstType
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetReadMode() *string {
	return s.ReadMode
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) GetRemainDays() *string {
	return s.RemainDays
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBInstanceId(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBInstanceStatus(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDBNodes(v *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DBNodes = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetDbInstType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEndpoints(v *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Endpoints = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEngine(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetEngineVersion(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetExpireTime(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetNetworkType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetPayType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetPort(v int32) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.Port = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetRdsInstType(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.RdsInstType = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetReadMode(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.ReadMode = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) SetRemainDays(v string) *DescribeDrdsDBClusterResponseBodyDbInstance {
	s.RemainDays = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes struct {
	DBNode []*DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode `json:"DBNode,omitempty" xml:"DBNode,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) GetDBNode() []*DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	return s.DBNode
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) SetDBNode(v []*DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes {
	s.DBNode = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode struct {
	// The ID of the node in the apsaradb for PolarDB cluster.
	//
	// example:
	//
	// pi-***************
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The role of a node in the apsaradb for PolarDB cluster. Valid values:
	//
	// 	- **Reader**
	//
	// 	- **Writer**
	//
	// example:
	//
	// Reader
	DBNodeRole *string `json:"DBNodeRole,omitempty" xml:"DBNodeRole,omitempty"`
	// The status of the nodes in the PolarDB cluster.
	//
	// example:
	//
	// Running
	DBNodeStatus *string `json:"DBNodeStatus,omitempty" xml:"DBNodeStatus,omitempty"`
	// The ID of the zone where the node of the PolarDB cluster resides.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) GetDBNodeRole() *string {
	return s.DBNodeRole
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) GetDBNodeStatus() *string {
	return s.DBNodeStatus
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeRole(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeRole = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetDBNodeStatus(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.DBNodeStatus = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) SetZoneId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode {
	s.ZoneId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceDBNodesDBNode) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints struct {
	Endpoint []*DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) GetEndpoint() []*DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	return s.Endpoint
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) SetEndpoint(v []*DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints {
	s.Endpoint = v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpoints) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint struct {
	// The ID of the PolarDB connection address.
	//
	// example:
	//
	// pe-*****************
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The ID list of the nodes in the PolarDB connection string. Separate multiple nodes with commas (,).
	//
	// example:
	//
	// pi-*****************,pi-*****************
	NodeIds *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// The read ratio of this connection address managed by the DRDS database.
	//
	// example:
	//
	// 85
	ReadWeight *int32 `json:"ReadWeight,omitempty" xml:"ReadWeight,omitempty"`
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) GetEndpointId() *string {
	return s.EndpointId
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) GetNodeIds() *string {
	return s.NodeIds
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) GetReadWeight() *int32 {
	return s.ReadWeight
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetEndpointId(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.EndpointId = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetNodeIds(v string) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.NodeIds = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) SetReadWeight(v int32) *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint {
	s.ReadWeight = &v
	return s
}

func (s *DescribeDrdsDBClusterResponseBodyDbInstanceEndpointsEndpoint) Validate() error {
	return dara.Validate(s)
}
