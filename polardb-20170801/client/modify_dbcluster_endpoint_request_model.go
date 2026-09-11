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
	// Specifies whether new nodes are automatically added to the endpoint. Valid values:
	//
	// - **Enable**: New nodes are automatically added.
	//
	// - **Disable**: New nodes are not automatically added. (Default)
	//
	// example:
	//
	// Enable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-*************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the endpoint.
	//
	// example:
	//
	// test
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	// The cluster endpoint ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pe-**************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The advanced configuration of the cluster endpoint in JSON format. You can configure the consistency level, transaction splitting, whether the primary node accepts read requests, and connection pooling.
	//
	// 	- To set the load balancing policy, use the format `{\\"LoadBalancePolicy\\":\\"Load balancing policy\\"}`. Valid values:
	//
	//     	- **0**: connection-based load balancing (default)
	//
	//     	- **1**: active-request-based load balancing
	//
	// 	- To set whether the primary node accepts read requests, use the format `{\\"MasterAcceptReads\\":\\"Whether the primary node accepts reads\\"}`. Valid values:
	//
	//     	- **on**: The primary node accepts read requests. (Default)
	//
	//     	- **off**: The primary node does not accept read requests.
	//
	// 	- To set transaction splitting, use the format `{\\"DistributedTransaction\\":\\"Transaction splitting\\"}`. Valid values:
	//
	//     	- **on**: Transaction splitting is enabled. (Default)
	//
	//     	- **off**: Transaction splitting is disabled.
	//
	// 	- To set the consistency level, use the format `{\\"ConsistLevel\\":\\"Consistency level\\"}`. Valid values:
	//
	//     	- **0**: eventual consistency (weak)
	//
	//     	- **1**: session consistency (medium) (default)
	//
	//     	- **2**: global consistency (strong)
	//
	// 	- To set the global consistency read timeout period, use the format `{\\"ConsistTimeout\\":\\"Global consistency read timeout\\"}`. Valid values: 0 to 60000. Default value: 20. Unit: ms.
	//
	// 	- To set the session consistency read timeout period, use the format `{\\"ConsistSessionTimeout\\":\\"Session consistency read timeout\\"}`. Valid values: 0 to 60000. Default value: 0. Unit: ms.
	//
	// 	- To set the global (or session) consistency read timeout policy, use the format `{\\"ConsistTimeoutAction\\":\\"Global consistency read timeout policy\\"}`. Valid values:
	//
	//     	- **0**: Forward the read request to the primary node. (Default)
	//
	//     	- **1**: The proxy returns the error message `wait replication complete timeout, please retry` to the application.
	//
	// 	- To set the connection pool, use the format `{\\"ConnectionPersist\\":\\"Connection pool\\"}`. Valid values:
	//
	//     	- **off**: The connection pool is disabled. (Default)
	//
	//     	- **Session**: The session-level connection pool is enabled.
	//
	//     	- **Transaction**: The transaction-level connection pooling is enabled.
	//
	// 	- To set parallel query, use the format `{\\"MaxParallelDegree\\":\\"Parallel query\\"}`. Valid values:
	//
	//     	- **on**: Parallel query is enabled.
	//
	//     	- **off**: Parallel query is disabled. (Default)
	//
	// 	- To set automatic request distribution among row offload reads from primary nodes, use the format `{\\"EnableHtapImci\\":\\"Automatic request distribution among row store and column store\\"}`. Valid values:
	//
	//     	- **on**: Automatic request distribution among row offload reads from primary nodes is enabled.
	//
	//     	- **off**: Automatic request distribution among row offload reads from primary nodes is disabled. (Default)
	//
	//
	// 	- To set whether to enable overload protection, use the format `{\\"EnableOverloadThrottle\\":\\"Whether to enable overload protection\\"}`. Valid values:
	//
	//     	- **on**: Overload protection is enabled.
	//
	//     	- **off**: Overload protection is disabled. (Default)
	//
	// 	- To set node weights, use the format `{\\"NodesWeight\\":{\\"Node ID\\":\\"Weight value\\"}}`.
	//
	// > 	- Transaction splitting, whether the primary node accepts read requests, connection pooling, and overload protection can be configured only when the read/write mode of the PolarDB for MySQL cluster endpoint is **ReadWrite (automatic read/write splitting)**.
	//
	// > 	- When the read/write mode of the PolarDB for MySQL cluster endpoint is **ReadOnly**, both **connection-based load balancing*	- and **active-request-based load balancing*	- policies are supported. The **ReadWrite (automatic read/write splitting)*	- mode supports only the **active-request-based load balancing*	- policy.
	//
	// > 	- Automatic request distribution among row offload reads from primary nodes can be configured when the read/write mode of the PolarDB for MySQL cluster endpoint is **ReadWrite (automatic read/write splitting)**, or when the read/write mode is **ReadOnly*	- and the load balancing policy is **active-request-based load balancing**.
	//
	// > 	- Only PolarDB for MySQL supports setting the consistency level to global consistency.
	//
	// > 	- If **ReadWriteMode*	- is set to **ReadOnly**, the consistency level can only be set to **0**.
	//
	// > 	- You can configure the consistency level, transaction splitting, whether the primary node accepts read requests, and connection pooling at the same time. Example: `{\\"ConsistLevel\\":\\"1\\",\\"DistributedTransaction\\":\\"on\\",\\"ConnectionPersist\\":\\"Session\\",\\"MasterAcceptReads\\":\\"on\\"}`.
	//
	// > 	- Transaction splitting is constrained by the consistency level. For example, transaction splitting cannot be enabled when the consistency level is **0**. Transaction splitting can be enabled when the consistency level is **1*	- or **2**.
	//
	// example:
	//
	// {"ConsistLevel":"1","DistributedTransaction":"on","MasterAcceptReads":"off","ConnectionPersist": "on"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The read load nodes to add to the endpoint. Separate multiple nodes with commas (,). Default value: the existing nodes.
	//
	// > 	- For PolarDB for MySQL, specify node IDs.
	//
	// > 	- For PolarDB for PostgreSQL and PolarDB for PostgreSQL (Compatible with Oracle), specify node role names, such as `Writer,Reader1,Reader2`.
	//
	// > 	- If **ReadWriteMode*	- is set to **ReadOnly**, you can mount only one node. However, if this node fails, the endpoint may be unavailable for up to 1 hour. Do not use this configuration in production environments. Select at least 2 nodes to improve availability.
	//
	// > 	- If **ReadWriteMode*	- is set to **ReadWrite**, select at least 2 nodes.
	//
	//     	- PolarDB for MySQL allows you to select any two nodes. If both nodes are read-only nodes, write requests are forwarded to the primary node.
	//
	//     	- PolarDB for PostgreSQL and PolarDB for PostgreSQL (Compatible with Oracle) require the primary node to be included.
	//
	// example:
	//
	// pi-**************,pi-*************
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The timeout policy for global consistency. Valid values:
	//
	// - **0**: Send the request to the primary node.
	//
	// - **2**: Timeout degradation. When the global consistency read times out, the query is automatically degraded to regular requests, and the client does not receive an error message.
	//
	// example:
	//
	// 0
	PolarSccTimeoutAction *string `json:"PolarSccTimeoutAction,omitempty" xml:"PolarSccTimeoutAction,omitempty"`
	// The timeout period for global consistency.
	//
	// example:
	//
	// 100
	PolarSccWaitTimeout *string `json:"PolarSccWaitTimeout,omitempty" xml:"PolarSccWaitTimeout,omitempty"`
	// The read/write mode. Valid values:
	//
	// - **ReadWrite**: read/write (automatic read/write splitting)
	//
	// - **ReadOnly**: read-only
	//
	// example:
	//
	// ReadWrite
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable the global consistency (high-performance mode) feature for the node. Valid values:
	//
	// - **ON**: Enabled.
	//
	// - **OFF**: Disabled.
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
