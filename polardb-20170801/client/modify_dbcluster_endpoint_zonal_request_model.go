// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterEndpointZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAddNewNodes(v string) *ModifyDBClusterEndpointZonalRequest
	GetAutoAddNewNodes() *string
	SetClientToken(v string) *ModifyDBClusterEndpointZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *ModifyDBClusterEndpointZonalRequest
	GetDBClusterId() *string
	SetDBEndpointDescription(v string) *ModifyDBClusterEndpointZonalRequest
	GetDBEndpointDescription() *string
	SetDBEndpointId(v string) *ModifyDBClusterEndpointZonalRequest
	GetDBEndpointId() *string
	SetEndpointConfig(v string) *ModifyDBClusterEndpointZonalRequest
	GetEndpointConfig() *string
	SetNodes(v string) *ModifyDBClusterEndpointZonalRequest
	GetNodes() *string
	SetOwnerAccount(v string) *ModifyDBClusterEndpointZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBClusterEndpointZonalRequest
	GetOwnerId() *int64
	SetPolarSccTimeoutAction(v string) *ModifyDBClusterEndpointZonalRequest
	GetPolarSccTimeoutAction() *string
	SetPolarSccWaitTimeout(v string) *ModifyDBClusterEndpointZonalRequest
	GetPolarSccWaitTimeout() *string
	SetReadWriteMode(v string) *ModifyDBClusterEndpointZonalRequest
	GetReadWriteMode() *string
	SetResourceOwnerAccount(v string) *ModifyDBClusterEndpointZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterEndpointZonalRequest
	GetResourceOwnerId() *int64
	SetSccMode(v string) *ModifyDBClusterEndpointZonalRequest
	GetSccMode() *string
}

type ModifyDBClusterEndpointZonalRequest struct {
	// Specifies whether to automatically add new nodes to this endpoint. Valid values:
	//
	// - **Enable**: Yes.
	//
	// - **Disable**: No (default).
	//
	// example:
	//
	// Enable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and is case-sensitive. The value can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 6000170000591aed949d0f******************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The edge cloud cluster ID. You can call DescribeDBClustersZonal to query the ID. The ID of a common PolarDB cluster returned by DescribeDBClusters is not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The endpoint name.
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
	// pe-****************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The advanced configurations of the cluster endpoint. The value is in JSON format. You can configure the consistency level, transaction splitting, whether to offload offload reads from primary node, and the connection pool.
	//
	// 	- To set the load balancing policy, use the format `{"LoadBalancePolicy":"Load balancing policy"}`. Valid values:
	//
	//     	- **0**: load balancing based on the number of connections (default).
	//
	//     	- **1**: load balancing based on the number of active requests.
	//
	// 	- To set whether to offload offload reads from primary node, use the format `{"MasterAcceptReads":"Whether to offload offload reads from primary node"}`. Valid values:
	//
	//     	- **on**: The primary node accepts reads (default).
	//
	//     	- **off**: Reads are offloaded from the primary node.
	//
	// 	- To set transaction splitting, use the format `{"DistributedTransaction":"Transaction splitting"}`. Valid values:
	//
	//     	- **on**: enables transaction splitting (default).
	//
	//     	- **off**: disables transaction splitting.
	//
	// 	- To set the consistency level, use the format `{"ConsistLevel":"Consistency level"}`. Valid values:
	//
	//     	- **0**: eventual consistency (weak).
	//
	//     	- **1**: session consistency (medium) (default).
	//
	//     	- **2**: global consistency (strong).
	//
	// 	- To set the global consistency read timeout period, use the format `{"ConsistTimeout":"Global consistency read timeout period"}`. Valid values: 0 to 60000. Default value: 20. Unit: milliseconds.
	//
	// 	- To set the session consistency read timeout period, use the format `{"ConsistSessionTimeout":"Session consistency read timeout period"}`. Valid values: 0 to 60000. Default value: 0. Unit: milliseconds.
	//
	// 	- To set the global (or session) consistency read timeout policy, use the format `{"ConsistTimeoutAction":"Global consistency read timeout policy"}`. Valid values:
	//
	//     	- **0**: forwards read requests to the primary node (default).
	//
	//     	- **1**: The proxy returns an error message `wait replication complete timeout, please retry` to the application.
	//
	// 	- To set the connection pool, use the format `{"ConnectionPersist":"Connection pool"}`. Valid values:
	//
	//     	- **off**: disables the connection pool (default).
	//
	//     	- **Session**: enables the session-level connection pool.
	//
	//     	- **Transaction**: enables transaction-level connection pooling.
	//
	// 	- To set parallel query, use the format `{"MaxParallelDegree":"Parallel query"}`. Valid values:
	//
	//     	- **on**: enables parallel query.
	//
	//     	- **off**: disables parallel query (default).
	//
	// 	- To set automatic request distribution between row store and column store nodes, use the format `{"EnableHtapImci":"Automatic request distribution between row store and column store nodes"}`. Valid values:
	//
	//     	- **on**: enables automatic request distribution between row store and column store nodes.
	//
	//     	- **off**: disables automatic request distribution between row store and column store nodes (default).
	//
	//
	// 	- To set whether to enable overload protection, use the format `{"EnableOverloadThrottle":"Whether to enable overload protection"}`. Valid values:
	//
	//     	- **on**: enables overload protection.
	//
	//     	- **off**: disables overload protection (default).
	//
	// > 	- You can configure transaction splitting, whether to offload offload reads from primary node, the connection pool, and whether to enable overload protection only when the read/write mode of the PolarDB for MySQL cluster endpoint is set to **Read/write (automatic read/write splitting)**.
	//
	// > 	- When the read/write mode of the PolarDB for MySQL cluster endpoint is set to **Read-only**, both **load balancing based on the number of connections*	- and **load balancing based on the number of active requests*	- are supported. Cluster endpoints in **Read/write (automatic read/write splitting)*	- mode support only **load balancing based on the number of active requests**.
	//
	// > 	- You can configure automatic request distribution between row store and column store nodes when the read/write mode of the PolarDB for MySQL cluster endpoint is set to **Read/write (automatic read/write splitting)**, or when the read/write mode is set to **Read-only*	- and the load balancing policy is set to **load balancing based on the number of active requests**.
	//
	// > 	- Only PolarDB for MySQL supports setting the consistency level to global consistency.
	//
	// > 	- If **ReadWriteMode*	- is set to **ReadOnly**, the consistency level can only be set to **0**.
	//
	// > 	- You can configure the consistency level, transaction splitting, whether to offload offload reads from primary node, and the connection pool at the same time. Example: `{"ConsistLevel":"1","DistributedTransaction":"on","ConnectionPersist":"Session","MasterAcceptReads":"on"}`.
	//
	// > 	- The transaction splitting configuration is restricted by the consistency level configuration. For example, if the consistency level is set to **0**, transaction splitting cannot be enabled. If the consistency level is set to **1*	- or **2**, transaction splitting can be enabled.
	//
	// example:
	//
	// {\\"DistributedTransaction\\":\\"off\\",\\"ConsistLevel\\":\\"0\\",\\"LoadBalanceStrategy\\":\\"load\\",\\"MasterAcceptReads\\":\\"on\\"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The read-only nodes to be added to the destination endpoint. Separate multiple nodes with commas (,). Default value: the original nodes.
	//
	// > 	- For PolarDB for MySQL, you must specify node IDs.
	//
	// > 	- For PolarDB for PostgreSQL and PolarDB for PostgreSQL (Compatible with Oracle), you must specify node role names, such as `Writer,Reader1,Reader2`.
	//
	// > 	- If **ReadWriteMode*	- is set to **ReadOnly**, you can mount only one node. However, if this node fails, the endpoint may be unavailable for up to 1 hour. Do not use this configuration in production environments. Therefore, select at least two nodes to improve availability.
	//
	// > 	- If **ReadWriteMode*	- is set to **ReadWrite**, you must select at least two nodes.
	//
	//     	- For PolarDB for MySQL, you can select any two nodes. If both nodes are read-only nodes, write requests are forwarded to the primary node.
	//
	//     	- For PolarDB for PostgreSQL and PolarDB for PostgreSQL (Compatible with Oracle), the primary node must be included.
	//
	// example:
	//
	// pi-**************,pi-*************
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The global consistency timeout policy. Valid values:
	//
	// - **0**: send the request to the primary node.
	//
	// - **2**: timeout degradation. If a global consistency read times out, the current query operation is automatically downgraded to regular requests, and the client does not receive an error message.
	//
	// example:
	//
	// 0
	PolarSccTimeoutAction *string `json:"PolarSccTimeoutAction,omitempty" xml:"PolarSccTimeoutAction,omitempty"`
	// The global consistency timeout period.
	//
	// example:
	//
	// 100
	PolarSccWaitTimeout *string `json:"PolarSccWaitTimeout,omitempty" xml:"PolarSccWaitTimeout,omitempty"`
	// The read/write mode. Valid values:
	//
	// - **ReadWrite**: read/write (automatic read/write splitting).
	//
	// - **ReadOnly**: read-only.
	//
	// example:
	//
	// ReadWrite
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable the global consistency (high-performance mode) feature for the node. Valid values: on (enabled) and off (disabled).
	//
	// example:
	//
	// off
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
}

func (s ModifyDBClusterEndpointZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterEndpointZonalRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterEndpointZonalRequest) GetAutoAddNewNodes() *string {
	return s.AutoAddNewNodes
}

func (s *ModifyDBClusterEndpointZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDBClusterEndpointZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterEndpointZonalRequest) GetDBEndpointDescription() *string {
	return s.DBEndpointDescription
}

func (s *ModifyDBClusterEndpointZonalRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *ModifyDBClusterEndpointZonalRequest) GetEndpointConfig() *string {
	return s.EndpointConfig
}

func (s *ModifyDBClusterEndpointZonalRequest) GetNodes() *string {
	return s.Nodes
}

func (s *ModifyDBClusterEndpointZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBClusterEndpointZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBClusterEndpointZonalRequest) GetPolarSccTimeoutAction() *string {
	return s.PolarSccTimeoutAction
}

func (s *ModifyDBClusterEndpointZonalRequest) GetPolarSccWaitTimeout() *string {
	return s.PolarSccWaitTimeout
}

func (s *ModifyDBClusterEndpointZonalRequest) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *ModifyDBClusterEndpointZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBClusterEndpointZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterEndpointZonalRequest) GetSccMode() *string {
	return s.SccMode
}

func (s *ModifyDBClusterEndpointZonalRequest) SetAutoAddNewNodes(v string) *ModifyDBClusterEndpointZonalRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetClientToken(v string) *ModifyDBClusterEndpointZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetDBClusterId(v string) *ModifyDBClusterEndpointZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetDBEndpointDescription(v string) *ModifyDBClusterEndpointZonalRequest {
	s.DBEndpointDescription = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetDBEndpointId(v string) *ModifyDBClusterEndpointZonalRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetEndpointConfig(v string) *ModifyDBClusterEndpointZonalRequest {
	s.EndpointConfig = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetNodes(v string) *ModifyDBClusterEndpointZonalRequest {
	s.Nodes = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetOwnerAccount(v string) *ModifyDBClusterEndpointZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetOwnerId(v int64) *ModifyDBClusterEndpointZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetPolarSccTimeoutAction(v string) *ModifyDBClusterEndpointZonalRequest {
	s.PolarSccTimeoutAction = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetPolarSccWaitTimeout(v string) *ModifyDBClusterEndpointZonalRequest {
	s.PolarSccWaitTimeout = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetReadWriteMode(v string) *ModifyDBClusterEndpointZonalRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetResourceOwnerAccount(v string) *ModifyDBClusterEndpointZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetResourceOwnerId(v int64) *ModifyDBClusterEndpointZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) SetSccMode(v string) *ModifyDBClusterEndpointZonalRequest {
	s.SccMode = &v
	return s
}

func (s *ModifyDBClusterEndpointZonalRequest) Validate() error {
	return dara.Validate(s)
}
