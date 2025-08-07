// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAddNewNodes(v string) *ModifyDBClusterEndpointRequest
	GetAutoAddNewNodes() *string
	SetDBClusterId(v string) *ModifyDBClusterEndpointRequest
	GetDBClusterId() *string
	SetDBEndpointDescription(v string) *ModifyDBClusterEndpointRequest
	GetDBEndpointDescription() *string
	SetDBEndpointId(v string) *ModifyDBClusterEndpointRequest
	GetDBEndpointId() *string
	SetEndpointConfig(v string) *ModifyDBClusterEndpointRequest
	GetEndpointConfig() *string
	SetNodes(v string) *ModifyDBClusterEndpointRequest
	GetNodes() *string
	SetOwnerAccount(v string) *ModifyDBClusterEndpointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterEndpointRequest
	GetOwnerId() *int64
	SetPolarSccTimeoutAction(v string) *ModifyDBClusterEndpointRequest
	GetPolarSccTimeoutAction() *string
	SetPolarSccWaitTimeout(v string) *ModifyDBClusterEndpointRequest
	GetPolarSccWaitTimeout() *string
	SetReadWriteMode(v string) *ModifyDBClusterEndpointRequest
	GetReadWriteMode() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterEndpointRequest
	GetResourceOwnerId() *int64
	SetSccMode(v string) *ModifyDBClusterEndpointRequest
	GetSccMode() *string
}

type ModifyDBClusterEndpointRequest struct {
	// Specifies whether to enable automatic association of newly added nodes with the cluster endpoint. Valid values:
	//
	// 	- **Enable**: enables automatic association of newly added nodes with the cluster endpoint.
	//
	// 	- **Disable*	- (default): disables automatic association of newly added nodes with the cluster endpoint.
	//
	// example:
	//
	// Enable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the custom cluster endpoint.
	//
	// example:
	//
	// test
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	// The ID of the endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// pe-**************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The advanced configurations of the cluster endpoint, which are in the JSON format. You can configure the consistency level, transaction splitting, and connection pool settings, and specify whether the primary node accepts read requests.
	//
	// 	- The load balancing policy. Format: `{\\"LoadBalancePolicy\\":\\"Load balancing policy\\"}`. Valid values:
	//
	//     	- **0*	- (default): connections-based load balancing
	//
	//     	- **1**: active requests-based load balancing
	//
	// 	- Specifies whether to allow the primary node to accept read requests. Format: `{\\"MasterAcceptReads\\":\\"Specification about whether to allow the primary node to accept read requests\\"}`. Valid values:
	//
	//     	- **on*	- (default): allows the primary node to accept read requests.
	//
	//     	- **off**: does not allow the primary node to accept read requests.
	//
	// 	- Specifies whether to enable the transaction splitting feature. Format: `{\\"DistributedTransaction\\":\\"Specification about whether to enable the transaction splitting feature\\"}`. Valid values:
	//
	//     	- **on*	- (default): enables the transaction splitting feature.
	//
	//     	- **off**: disables the transaction splitting feature.
	//
	// 	- The consistency level. Format: `{\\"ConsistLevel\\":\\"Consistency level\\"}`. Valid values:
	//
	//     	- **0**: eventual consistency (weak)
	//
	//     	- **1*	- (default): session consistency (medium)
	//
	//     	- **2**: global consistency (strong)
	//
	// 	- The global consistency timeout. Format: `{\\"ConsistTimeout\\":\\"Global consistency timeout\\"}`. Valid values: 0 to 60,000. Default value: 20. Unit: ms.
	//
	// 	- The session consistency timeout. Format: `{\\"ConsistSessionTimeout\\":\\"Session consistency timeout\\"}`. Valid values: 0 to 60,000. Default value: 0. Unit: ms.
	//
	// 	- The global (or session) consistency timeout policy. Format: `{\\"ConsistTimeoutAction\\":\\"Consistency timeout policy\\"}`. Valid values:
	//
	//     	- **0*	- (default): PolarProxy sends read requests to the primary node.
	//
	//     	- **1**: PolarProxy returns the "wait replication complete timeout, please retry" error message to the application.
	//
	// 	- Specifies whether to enable the connection pool feature. Format: `{\\"ConnectionPersist\\":\\"Specification about whether to enable the connection pool feature\\"}`. Valid values:
	//
	//     	- **off*	- (default): disables the connection pool feature.
	//
	//     	- **Session**: enables the session-level connection pool.
	//
	//     	- **Transaction**: enables the transaction-level connection pool.
	//
	// 	- Specifies whether to enable the parallel query feature. Format: `{\\"MaxParallelDegree\\":\\"Specification about whether to enable the parallel query feature\\"}`. Valid values:
	//
	//     	- **on**: enables the parallel query feature.
	//
	//     	- **off*	- (default): disables the parallel query feature.
	//
	// 	- Specifies whether to enable the automatic request distribution among row store and column store nodes feature. Format: `{\\"EnableHtapImci\\":\\"Specification about whether to enable automatic request distribution among row store and column store nodes feature\\"}`. Valid values:
	//
	//     	- **on**: enables the automatic request distribution among row store and column store nodes feature.
	//
	//     	- **off*	- (default): disables the automatic request distribution among row store and column store nodes feature.
	//
	// 	- Specifies whether to enable the overload protection feature. Format: `{\\"EnableOverloadThrottle\\":\\"Specification about whether to enable the overload protection feature\\"}`. Valid values:
	//
	//     	- **on**: enables the overload protection feature.
	//
	//     	- **off*	- (default): disables the overload protection feature.
	//
	// >
	//
	// 	- You can configure the transaction splitting, connection pool, and overload protection settings, and specify whether the primary node accepts read requests settings for the cluster endpoint of a PolarDB for MySQL cluster only if ReadWriteMode of the cluster endpoint is set to Read and Write (Automatic Read/Write Splitting).
	//
	// 	- If ReadWriteMode of the cluster endpoint of a PolarDB for MySQL cluster is set to **Read-only**, you can specify the **Connections-based Load Balancing*	- or **Active Request-based Load Balancing*	- policy for the cluster endpoint. If ReadWriteMode of the cluster endpoint of a PolarDB for MySQL cluster is set to **Read/Write (Automatic Read/Write Splitting)**, you can specify only the **Active Request-based Load Balancing*	- policy for the cluster endpoint.
	//
	// 	- You can enable automatic request distribution among column store and row store nodes for the cluster endpoint of a PolarDB for MySQL cluster if ReadWriteMode of the cluster endpoint is set to **Read and Write (Automatic Read/Write Splitting)**, or if the ReadWriteMode of the cluster endpoint is set to **Read-only*	- and the load balancing policy is set to **Active requests-based load balancing**.
	//
	// 	- Only PolarDB for MySQL supports global consistency.
	//
	// 	- You can set the consistency level of the cluster endpoint of a PolarDB for MySQL cluster only to **0*	- if **ReadWriteMode*	- of the cluster endpoint is set to **ReadOnly**.
	//
	// 	- You can configure the settings for the consistency level, transaction splitting, and connection pool features, and specify whether the primary node accepts read requests settings at a time. Example: `{\\"ConsistLevel\\":\\"1\\",\\"DistributedTransaction\\":\\"on\\",\\"ConnectionPersist\\":\\"Session\\",\\"MasterAcceptReads\\":\\"on\\"}`.
	//
	// 	- The configuration for transaction splitting is limited by the configuration for the consistency level. For example, if you set the consistency level to **0**, you cannot enable transaction splitting. If you set the consistency level to **1*	- or **2**, you can enable transaction splitting.
	//
	// example:
	//
	// {"ConsistLevel":"1","DistributedTransaction":"on","MasterAcceptReads":"off","ConnectionPersist": "on"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The reader nodes to be associated with the endpoint. If you need to specify multiple reader nodes, separate the reader nodes with commas (,). If you do not specify this parameter, the predefined nodes are used by default.
	//
	// >
	//
	// 	- You must specify the node ID for each PolarDB for MySQL cluster.
	//
	// 	- You must specify the role name of each node for each PolarDB for PostgreSQL or PolarDB for Oracle cluster. Example: `Writer,Reader1,Reader2`.
	//
	// 	- If you set **ReadWriteMode*	- to **ReadOnly**, only one node can be associated with the cluster endpoint. If the only node becomes faulty, the cluster endpoint may be unavailable for up to an hour. We recommend that you do not associate only one node with the cluster endpoint in production environments. We recommend that you associate at least two nodes with the cluster endpoint to improve service availability.
	//
	// 	- If you set **ReadWriteMode*	- to **ReadWrite**, you must associate at least two nodes with the cluster endpoint.
	//
	//     	- No limits are imposed on the two nodes that you select for each PolarDB for MySQL cluster. If the two nodes are read-only nodes, write requests are forwarded to the primary node.
	//
	//     	- The following limit applies to PolarDB for PostgreSQL and PolarDB for Oracle clusters: One of the selected nodes must be the primary node.
	//
	// example:
	//
	// pi-**************,pi-*************
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Global consistency timeout policy. Valid values:
	//
	// 	- **0**: sends the request to the primary node.
	//
	// 	- **2**: downgrades the consistency level of a query to inconsistent read when a global consistent read in the query times out. No error message is returned to the client.
	//
	// example:
	//
	// 0
	PolarSccTimeoutAction *string `json:"PolarSccTimeoutAction,omitempty" xml:"PolarSccTimeoutAction,omitempty"`
	// Global consistency timeout.
	//
	// example:
	//
	// 100
	PolarSccWaitTimeout *string `json:"PolarSccWaitTimeout,omitempty" xml:"PolarSccWaitTimeout,omitempty"`
	// The read/write mode. Valid values:
	//
	// 	- **ReadWrite**: The cluster endpoint handles read and write requests. Automatic read/write splitting is enabled.
	//
	// 	- **ReadOnly**: The cluster endpoint handles read-only requests.
	//
	// example:
	//
	// ReadWrite
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable the global consistency (high-performance mode) feature for the nodes. Valid values:
	//
	// 	- **ON**
	//
	// 	- **OFF**
	//
	// example:
	//
	// on
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
}

func (s ModifyDBClusterEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterEndpointRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointRequest) GetAutoAddNewNodes() *string {
	return s.AutoAddNewNodes
}

func (s *ModifyDBClusterEndpointRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterEndpointRequest) GetDBEndpointDescription() *string {
	return s.DBEndpointDescription
}

func (s *ModifyDBClusterEndpointRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *ModifyDBClusterEndpointRequest) GetEndpointConfig() *string {
	return s.EndpointConfig
}

func (s *ModifyDBClusterEndpointRequest) GetNodes() *string {
	return s.Nodes
}

func (s *ModifyDBClusterEndpointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterEndpointRequest) GetPolarSccTimeoutAction() *string {
	return s.PolarSccTimeoutAction
}

func (s *ModifyDBClusterEndpointRequest) GetPolarSccWaitTimeout() *string {
	return s.PolarSccWaitTimeout
}

func (s *ModifyDBClusterEndpointRequest) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *ModifyDBClusterEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterEndpointRequest) GetSccMode() *string {
	return s.SccMode
}

func (s *ModifyDBClusterEndpointRequest) SetAutoAddNewNodes(v string) *ModifyDBClusterEndpointRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetDBClusterId(v string) *ModifyDBClusterEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetDBEndpointDescription(v string) *ModifyDBClusterEndpointRequest {
	s.DBEndpointDescription = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetDBEndpointId(v string) *ModifyDBClusterEndpointRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetEndpointConfig(v string) *ModifyDBClusterEndpointRequest {
	s.EndpointConfig = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetNodes(v string) *ModifyDBClusterEndpointRequest {
	s.Nodes = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetOwnerAccount(v string) *ModifyDBClusterEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetOwnerId(v int64) *ModifyDBClusterEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetPolarSccTimeoutAction(v string) *ModifyDBClusterEndpointRequest {
	s.PolarSccTimeoutAction = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetPolarSccWaitTimeout(v string) *ModifyDBClusterEndpointRequest {
	s.PolarSccWaitTimeout = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetReadWriteMode(v string) *ModifyDBClusterEndpointRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetResourceOwnerId(v int64) *ModifyDBClusterEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) SetSccMode(v string) *ModifyDBClusterEndpointRequest {
	s.SccMode = &v
	return s
}

func (s *ModifyDBClusterEndpointRequest) Validate() error {
	return dara.Validate(s)
}
