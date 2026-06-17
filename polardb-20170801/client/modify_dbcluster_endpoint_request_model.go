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
	// Specifies whether to automatically add new nodes to the endpoint. Valid values:
	//
	// - **Enable**: Automatically adds new nodes.
	//
	// - **Disable**: Does not automatically add new nodes. This is the default value.
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
	// The ID of the cluster endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// pe-**************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The advanced configurations of the cluster endpoint in JSON format. You can set the consistency level, transaction splitting, whether the primary node accepts read requests, the connection pool, and other settings.
	//
	// - Sets the load balancing policy. Format: `{\\"LoadBalancePolicy\\":\\"policy\\"}`. Valid values:
	//
	//   - **0**: Connections-based load balancing (default)
	//
	//   - **1**: Active requests-based load balancing
	//
	// - Specifies whether the primary node accepts read requests. Format: `{\\"MasterAcceptReads\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: The primary node accepts read requests (default).
	//
	//   - **off**: The primary node does not accept read requests.
	//
	// - Enables or disables transaction splitting. Format: `{\\"DistributedTransaction\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: Enables transaction splitting (default).
	//
	//   - **off**: Disables transaction splitting.
	//
	// - Sets the consistency level. Format: `{\\"ConsistLevel\\":\\"level\\"}`. Valid values:
	//
	//   - **0**: Eventual consistency (weak)
	//
	//   - **1**: Session consistency (medium) (default)
	//
	//   - **2**: Global consistency (strong)
	//
	// - Sets the timeout period for a global consistency read. Format: `{\\"ConsistTimeout\\":\\"timeout\\"}`. Valid values: 0 to 60000. Default value: 20. Unit: ms.
	//
	// - Sets the timeout period for a session consistency read. Format: `{\\"ConsistSessionTimeout\\":\\"timeout\\"}`. Valid values: 0 to 60000. Default value: 0. Unit: ms.
	//
	// - Sets the policy for handling timeouts of global or session consistency reads. Format: `{\\"ConsistTimeoutAction\\":\\"policy\\"}`. Valid values:
	//
	//   - **0**: Forwards read requests to the primary node (default).
	//
	//   - **1**: The proxy returns the error message \\`wait replication complete timeout, please retry\\` to the application.
	//
	// - Sets the connection pool type. Format: `{\\"ConnectionPersist\\":\\"type\\"}`. Valid values:
	//
	//   - **off**: Disables the connection pool (default).
	//
	//   - **Session**: Enables the session-level connection pool.
	//
	//   - **Transaction**: Enables the transaction-level connection pool.
	//
	// - Enables or disables parallel query. Format: `{\\"MaxParallelDegree\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: Enables parallel query.
	//
	//   - **off**: Disables parallel query (default).
	//
	// - Enables or disables automatic routing of requests to the row store or column store. Format: `{\\"EnableHtapImci\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: Enables automatic routing.
	//
	//   - **off**: Disables automatic routing (default).
	//
	// - Enables or disables overload protection. Format: `{\\"EnableOverloadThrottle\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: Enables overload protection.
	//
	//   - **off**: Disables overload protection (default).
	//
	// > 	- You can set transaction splitting, whether the primary node accepts read requests, the connection pool, and overload protection only when the read/write mode of the cluster endpoint for PolarDB for MySQL is set to ReadWrite (automatic read/write splitting).
	//
	// >
	//
	// > 	- If the read/write mode of a cluster endpoint for PolarDB for MySQL is **ReadOnly**, both **connections-based*	- and **active requests-based*	- load balancing policies are supported. If the read/write mode is **ReadWrite*	- (automatic read/write splitting), only the **active requests-based*	- load balancing policy is supported.
	//
	// >
	//
	// > 	- You can enable automatic routing to the row store or column store if the read/write mode of the cluster endpoint for PolarDB for MySQL is **ReadWrite*	- (automatic read/write splitting), or if the read/write mode is **ReadOnly*	- and the load balancing policy is **active requests-based**.
	//
	// >
	//
	// > 	- Only PolarDB for MySQL supports global consistency.
	//
	// >
	//
	// > 	- If you set the **ReadWriteMode*	- parameter to **ReadOnly**, you can only set the consistency level to **0**.
	//
	// >
	//
	// > 	- You can set the consistency level, transaction splitting, whether the primary node accepts read requests, and the connection pool at the same time. For example: `{\\"ConsistLevel\\":\\"1\\",\\"DistributedTransaction\\":\\"on\\",\\"ConnectionPersist\\":\\"Session\\",\\"MasterAcceptReads\\":\\"on\\"}`.
	//
	// >
	//
	// > 	- The transaction splitting setting is constrained by the consistency level. For example, you cannot enable transaction splitting if the consistency level is **0*	- (eventual consistency). You can enable transaction splitting if the consistency level is **1*	- (session consistency) or **2*	- (global consistency).
	//
	// example:
	//
	// {"ConsistLevel":"1","DistributedTransaction":"on","MasterAcceptReads":"off","ConnectionPersist": "on"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The nodes to be added to the endpoint for read request distribution. Separate multiple node IDs with commas (,). The original nodes are used by default.
	//
	// > - For PolarDB for MySQL, specify the node IDs.
	//
	// >
	//
	// > - For PolarDB for PostgreSQL and PolarDB for PostgreSQL (Oracle Compatible), specify the node roles, such as `Writer,Reader1,Reader2`.
	//
	// >
	//
	// > - If you set **ReadWriteMode*	- to **ReadOnly**, you can attach only one node. However, if this node fails, the endpoint may be unavailable for up to one hour. Do not use this configuration in a production environment. Select at least two nodes to improve availability.
	//
	// >
	//
	// > - If you set **ReadWriteMode*	- to **ReadWrite**, you must select at least two nodes.
	//
	// >   \\	- For PolarDB for MySQL, you can select any two nodes. If both nodes are read-only nodes, write requests are sent to the primary node.
	//
	// >   \\	- For PolarDB for PostgreSQL and PolarDB for PostgreSQL (Oracle Compatible), you must include the primary node.
	//
	// example:
	//
	// pi-**************,pi-*************
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The policy for handling global consistency timeouts. Valid values:
	//
	// - **0**: Forwards the request to the primary node.
	//
	// - **2**: Degrades the request. If a global consistency read times out, the query is automatically degraded to a regular request. The client does not receive an error message.
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
	// - **ReadWrite**: Read/write (automatic read/write splitting)
	//
	// - **ReadOnly**: Read-only
	//
	// example:
	//
	// ReadWrite
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable the global consistency (high-performance mode) feature for the node. Valid values:
	//
	// - **ON**: Enable
	//
	// - **OFF**: Disable
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
