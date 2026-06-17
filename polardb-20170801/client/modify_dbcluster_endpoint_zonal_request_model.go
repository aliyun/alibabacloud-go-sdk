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
	// - **Enable**: yes
	//
	// - **Disable**: no (default)
	//
	// example:
	//
	// Enable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// A client token to ensure the idempotence of the request. The client generates the value, but you must make sure that the value is unique among different requests. The token is case-sensitive and can contain up to 64 ASCII characters.
	//
	// example:
	//
	// 6000170000591aed949d0f******************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
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
	// pe-****************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The advanced configurations of the cluster endpoint, which are specified in the JSON format. You can set the consistency level, transaction splitting, whether the primary node accepts read requests, the connection pool, and more.
	//
	// - To set the load balancing policy, use the format `{\\"LoadBalancePolicy\\":\\"policy\\"}`. Valid values:
	//
	//   - **0**: connection-based load balancing (default)
	//
	//   - **1**: active request-based load balancing
	//
	// - To specify whether the primary node accepts read requests, use the format `{\\"MasterAcceptReads\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: The primary node accepts read requests (default).
	//
	//   - **off**: The primary node does not accept read requests.
	//
	// - To configure transaction splitting, use the format `{\\"DistributedTransaction\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: enables transaction splitting (default)
	//
	//   - **off**: disables transaction splitting
	//
	// - To set the consistency level, use the format `{\\"ConsistLevel\\":\\"level\\"}`. Valid values:
	//
	//   - **0**: eventual consistency (weak)
	//
	//   - **1**: session consistency (medium) (default)
	//
	//   - **2**: global consistency (strong)
	//
	// - To set the timeout period for a global consistency read, use the format `{\\"ConsistTimeout\\":\\"timeout\\"}`. Valid values: 0 to 60000. Default value: 20. Unit: ms.
	//
	// - To set the timeout period for a session consistency read, use the format `{\\"ConsistSessionTimeout\\":\\"timeout\\"}`. Valid values: 0 to 60000. Default value: 0. Unit: ms.
	//
	// - To set the policy for a global or session consistency read timeout, use the format `{\\"ConsistTimeoutAction\\":\\"policy\\"}`. Valid values:
	//
	//   - **0**: Sends the read request to the primary node (default).
	//
	//   - **1**: The proxy returns a \\`wait replication complete timeout, please retry\\` error message to the application.
	//
	// - To configure the connection pool, use the format `{\\"ConnectionPersist\\":\\"pool_type\\"}`. Valid values:
	//
	//   - **off**: disables the connection pool (default)
	//
	//   - **Session**: enables the session-level connection pool
	//
	//   - **Transaction**: enables transaction-level connection pooling
	//
	// - To configure parallel queries, use the format `{\\"MaxParallelDegree\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: enables parallel queries
	//
	//   - **off**: disables parallel queries (default)
	//
	// - To configure automatic routing between row store and column store, use the format `{\\"EnableHtapImci\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: enables automatic routing between row store and column store
	//
	//   - **off**: disables automatic routing between row store and column store (default)
	//
	// - To specify whether to enable overload protection, use the format `{\\"EnableOverloadThrottle\\":\\"value\\"}`. Valid values:
	//
	//   - **on**: enables overload protection
	//
	//   - **off**: disables overload protection (default)
	//
	// > 	- You can configure transaction splitting, specify whether the primary node accepts read requests, configure the connection pool, and enable overload protection only when the read/write mode of the cluster endpoint for PolarDB for MySQL is set to **ReadWrite*	- (automatic read/write splitting).
	//
	// >
	//
	// > 	- If the read/write mode of a cluster endpoint for PolarDB for MySQL is **ReadOnly**, both connection-based and active request-based load balancing policies are supported. If the read/write mode is **ReadWrite*	- (automatic read/write splitting), only the active request-based load balancing policy is supported.
	//
	// >
	//
	// > 	- You can configure automatic routing between row store and column store when the read/write mode of the cluster endpoint for PolarDB for MySQL is **ReadWrite*	- (automatic read/write splitting), or when the read/write mode is **ReadOnly*	- and the load balancing policy is active request-based.
	//
	// >
	//
	// > 	- Only PolarDB for MySQL supports the global consistency level.
	//
	// >
	//
	// > 	- If you set **ReadWriteMode*	- to **ReadOnly**, you can set the consistency level only to **0**.
	//
	// >
	//
	// > 	- You can set the consistency level, transaction splitting, whether the primary node accepts read requests, and the connection pool at the same time. For example: `{\\"ConsistLevel\\":\\"1\\",\\"DistributedTransaction\\":\\"on\\",\\"ConnectionPersist\\":\\"Session\\",\\"MasterAcceptReads\\":\\"on\\"}`.
	//
	// >
	//
	// > 	- The transaction splitting setting is constrained by the consistency level. For example, if the consistency level is **0**, you cannot enable transaction splitting. If the consistency level is **1*	- or **2**, you can enable transaction splitting.
	//
	// example:
	//
	// {\\"DistributedTransaction\\":\\"off\\",\\"ConsistLevel\\":\\"0\\",\\"LoadBalanceStrategy\\":\\"load\\",\\"MasterAcceptReads\\":\\"on\\"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The read-only nodes to add to the endpoint. Separate multiple node IDs with commas (,). If you do not specify this parameter, the original nodes are retained.
	//
	// > - For PolarDB for MySQL, specify the node IDs.
	//
	// >
	//
	// > - For PolarDB for PostgreSQL and PolarDB for PostgreSQL (compatible with Oracle), specify the node role names, such as `Writer,Reader1,Reader2`.
	//
	// >
	//
	// > - If you set **ReadWriteMode*	- to **ReadOnly**, you can attach only one node. However, if this node fails, the endpoint might be unavailable for up to one hour. Do not use this configuration in a production environment. Select at least two nodes to improve availability.
	//
	// >
	//
	// > - If you set **ReadWriteMode*	- to **ReadWrite**, you must select at least two nodes.
	//
	// >   \\	- For PolarDB for MySQL, you can select any two nodes. If both nodes are read-only nodes, write requests are sent to the primary node.
	//
	// >   \\	- For PolarDB for PostgreSQL and PolarDB for PostgreSQL (compatible with Oracle), you must include the primary node.
	//
	// example:
	//
	// pi-**************,pi-*************
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The policy for a global consistency timeout. Valid values:
	//
	// - **0**: Sends the request to the primary node.
	//
	// - **2**: Timeout degradation. If a global consistency read times out, the query is automatically downgraded to a regular request, and the client does not receive an error message.
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
	// - **ReadWrite**: read-write (automatic read/write splitting)
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
	// - **ON**: enables the feature
	//
	// - **OFF**: disables the feature
	//
	// example:
	//
	// OFF
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
