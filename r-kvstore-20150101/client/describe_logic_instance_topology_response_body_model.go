// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogicInstanceTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeLogicInstanceTopologyResponseBody
	GetInstanceId() *string
	SetRedisProxyList(v *DescribeLogicInstanceTopologyResponseBodyRedisProxyList) *DescribeLogicInstanceTopologyResponseBody
	GetRedisProxyList() *DescribeLogicInstanceTopologyResponseBodyRedisProxyList
	SetRedisShardList(v *DescribeLogicInstanceTopologyResponseBodyRedisShardList) *DescribeLogicInstanceTopologyResponseBody
	GetRedisShardList() *DescribeLogicInstanceTopologyResponseBodyRedisShardList
	SetRequestId(v string) *DescribeLogicInstanceTopologyResponseBody
	GetRequestId() *string
}

type DescribeLogicInstanceTopologyResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The detailed proxy information, including information about proxy nodes.
	RedisProxyList *DescribeLogicInstanceTopologyResponseBodyRedisProxyList `json:"RedisProxyList,omitempty" xml:"RedisProxyList,omitempty" type:"Struct"`
	// Details of data shards, which includes node information such as NodeInfo.
	RedisShardList *DescribeLogicInstanceTopologyResponseBodyRedisShardList `json:"RedisShardList,omitempty" xml:"RedisShardList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 794120D1-E0CF-4713-BAE4-EBAEA04506AF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogicInstanceTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeLogicInstanceTopologyResponseBody) GetRedisProxyList() *DescribeLogicInstanceTopologyResponseBodyRedisProxyList {
	return s.RedisProxyList
}

func (s *DescribeLogicInstanceTopologyResponseBody) GetRedisShardList() *DescribeLogicInstanceTopologyResponseBodyRedisShardList {
	return s.RedisShardList
}

func (s *DescribeLogicInstanceTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogicInstanceTopologyResponseBody) SetInstanceId(v string) *DescribeLogicInstanceTopologyResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBody) SetRedisProxyList(v *DescribeLogicInstanceTopologyResponseBodyRedisProxyList) *DescribeLogicInstanceTopologyResponseBody {
	s.RedisProxyList = v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBody) SetRedisShardList(v *DescribeLogicInstanceTopologyResponseBodyRedisShardList) *DescribeLogicInstanceTopologyResponseBody {
	s.RedisShardList = v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBody) SetRequestId(v string) *DescribeLogicInstanceTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLogicInstanceTopologyResponseBodyRedisProxyList struct {
	NodeInfo []*DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisProxyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisProxyList) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyList) GetNodeInfo() []*DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	return s.NodeInfo
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyList) SetNodeInfo(v []*DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) *DescribeLogicInstanceTopologyResponseBodyRedisProxyList {
	s.NodeInfo = v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyList) Validate() error {
	return dara.Validate(s)
}

type DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo struct {
	// The bandwidth throttling of the node. Unit: MB/s.
	//
	// example:
	//
	// 96
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The storage capacity of the node. Unit: MB.
	//
	// example:
	//
	// 5120
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The maximum number of connections.
	//
	// example:
	//
	// 320000
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// r-bp10noxlhcoim2****-proxy-3#542****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node type. Valid values:
	//
	// 	- **proxy**: proxy node
	//
	// 	- **db**: data node
	//
	// example:
	//
	// proxy
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) GetCapacity() *string {
	return s.Capacity
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) GetConnection() *string {
	return s.Connection
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetBandwidth(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetCapacity(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.Capacity = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetConnection(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.Connection = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetNodeId(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) SetNodeType(v string) *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisProxyListNodeInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeLogicInstanceTopologyResponseBodyRedisShardList struct {
	NodeInfo []*DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo `json:"NodeInfo,omitempty" xml:"NodeInfo,omitempty" type:"Repeated"`
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisShardList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisShardList) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardList) GetNodeInfo() []*DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	return s.NodeInfo
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardList) SetNodeInfo(v []*DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) *DescribeLogicInstanceTopologyResponseBodyRedisShardList {
	s.NodeInfo = v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardList) Validate() error {
	return dara.Validate(s)
}

type DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo struct {
	// The bandwidth throttling of the node. Unit: MB/s.
	//
	// example:
	//
	// 96
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The storage capacity of the node. Unit: MB.
	//
	// example:
	//
	// 2048
	Capacity *string `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The maximum number of connections.
	//
	// example:
	//
	// 10000
	Connection *string `json:"Connection,omitempty" xml:"Connection,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// r-bp10noxlhcoim2****-db-0#688****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The node type. Valid values:
	//
	// 	- **proxy**: proxy node
	//
	// 	- **db**: data node
	//
	// example:
	//
	// db
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// 子实例类型，返回值：
	//
	// 	- **master**：主节点类型。
	//
	// 	- **readonly**：只读实例类型。
	//
	// example:
	//
	// master
	SubInstanceType *string `json:"SubInstanceType,omitempty" xml:"SubInstanceType,omitempty"`
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) GoString() string {
	return s.String()
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) GetCapacity() *string {
	return s.Capacity
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) GetConnection() *string {
	return s.Connection
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) GetNodeType() *string {
	return s.NodeType
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) GetSubInstanceType() *string {
	return s.SubInstanceType
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetBandwidth(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetCapacity(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.Capacity = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetConnection(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.Connection = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetNodeId(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.NodeId = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetNodeType(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.NodeType = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) SetSubInstanceType(v string) *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo {
	s.SubInstanceType = &v
	return s
}

func (s *DescribeLogicInstanceTopologyResponseBodyRedisShardListNodeInfo) Validate() error {
	return dara.Validate(s)
}
