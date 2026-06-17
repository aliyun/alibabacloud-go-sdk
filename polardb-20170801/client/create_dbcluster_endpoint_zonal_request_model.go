// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterEndpointZonalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAddNewNodes(v string) *CreateDBClusterEndpointZonalRequest
	GetAutoAddNewNodes() *string
	SetClientToken(v string) *CreateDBClusterEndpointZonalRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateDBClusterEndpointZonalRequest
	GetDBClusterId() *string
	SetDBEndpointDescription(v string) *CreateDBClusterEndpointZonalRequest
	GetDBEndpointDescription() *string
	SetEndpointConfig(v string) *CreateDBClusterEndpointZonalRequest
	GetEndpointConfig() *string
	SetEndpointType(v string) *CreateDBClusterEndpointZonalRequest
	GetEndpointType() *string
	SetNodes(v string) *CreateDBClusterEndpointZonalRequest
	GetNodes() *string
	SetOwnerAccount(v string) *CreateDBClusterEndpointZonalRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBClusterEndpointZonalRequest
	GetOwnerId() *int64
	SetPolarSccTimeoutAction(v string) *CreateDBClusterEndpointZonalRequest
	GetPolarSccTimeoutAction() *string
	SetPolarSccWaitTimeout(v string) *CreateDBClusterEndpointZonalRequest
	GetPolarSccWaitTimeout() *string
	SetReadWriteMode(v string) *CreateDBClusterEndpointZonalRequest
	GetReadWriteMode() *string
	SetResourceOwnerAccount(v string) *CreateDBClusterEndpointZonalRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBClusterEndpointZonalRequest
	GetResourceOwnerId() *int64
	SetSccMode(v string) *CreateDBClusterEndpointZonalRequest
	GetSccMode() *string
}

type CreateDBClusterEndpointZonalRequest struct {
	// Specifies whether to automatically add new nodes to this endpoint. Valid values:
	//
	// - Enable: New nodes are automatically added to this endpoint.
	//
	// - Disable: New nodes are not automatically added to this endpoint. This is the default value.
	//
	// example:
	//
	// Disable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// A client token that is used to ensure the idempotence of the request. The client generates the value, which must be unique among different requests. The token is case-sensitive and can be up to 64 ASCII characters in length.
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
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the custom cluster endpoint.
	//
	// example:
	//
	// test
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	// The advanced configurations of the cluster endpoint, specified in the JSON format. This parameter supports settings for consistency level, transaction splitting, offloading reads from the primary node, and the connection pool.
	//
	// - Sets the load balancing policy. The format is {"LoadBalancePolicy":"policy"}. Valid values:
	//
	//   - **0**: Connections-based load balancing. This is the default value.
	//
	//   - **1**: Active requests-based load balancing.
	//
	// - Sets the consistency level. The format is `{"ConsistLevel":"level"}`. Valid values:
	//
	//   - **0**: Eventual consistency.
	//
	//   - **1**: Session consistency. This is the default value.
	//
	//   - **2**: Global consistency.
	//
	// - Sets transaction splitting. The format is `{"DistributedTransaction":"on/off"}`. Valid values:
	//
	//   - **on**: Enables transaction splitting. This is the default value.
	//
	//   - **off**: Disables transaction splitting.
	//
	// - Specifies whether the primary node accepts read requests. The format is `{"MasterAcceptReads":"on/off"}`. Valid values:
	//
	//   - **on**: The primary node accepts read requests.
	//
	//   - **off**: The primary node does not accept read requests. This is the default value.
	//
	// - Sets the connection pool. The format is `{"ConnectionPersist":"type"}`. Valid values:
	//
	//   - **off**: Disables the connection pool. This is the default value.
	//
	//   - **Session**: Enables the session-level connection pool.
	//
	//   - **Transaction**: Enables transaction-level connection pooling.
	//
	// - Sets parallel query. The format is {"MaxParallelDegree":"degree"}. Valid values:
	//
	//   - A specific number of concurrent queries. Example: "MaxParallelDegree":"2".
	//
	//   - **off**: Disables parallel query. This is the default value.
	//
	// - Sets automatic routing for row store and column store. The format is {"EnableHtapImci":"on/off"}. Valid values:
	//
	//   - **on**: Enables automatic routing for row store and column store.
	//
	//   - **off**: Disables automatic routing for row store and column store. This is the default value.
	//
	// - Specifies whether to enable overload protection. The format is {"EnableOverloadThrottle":"on/off"}. Valid values:
	//
	//   - **on**: Enables overload protection.
	//
	//   - **off**: Disables overload protection. This is the default value.
	//
	// > 	- You can set transaction splitting, specify whether the primary node accepts read requests, configure the connection pool, and enable overload protection only when the read/write mode of the cluster endpoint for a PolarDB for MySQL cluster is set to \\*\\*ReadWrite\\*\\	- (automatic read/write splitting).
	//
	// >
	//
	// > 	- When the read/write mode of the cluster endpoint for a PolarDB for MySQL cluster is set to **ReadOnly**, both **connections-based load balancing*	- and **active requests-based load balancing*	- are supported. When the read/write mode is set to **ReadWrite*	- (automatic read/write splitting), only **active requests-based load balancing*	- is supported.
	//
	// >
	//
	// > 	- You can enable automatic routing for row store and column store when the read/write mode of the cluster endpoint for a PolarDB for MySQL cluster is set to **ReadWrite*	- (automatic read/write splitting), or when the read/write mode is set to **ReadOnly*	- and the load balancing policy is set to **active requests-based load balancing**.
	//
	// >
	//
	// > 	- Only PolarDB for MySQL supports setting the consistency level to global consistency.
	//
	// >
	//
	// > 	- If you set the **ReadWriteMode*	- parameter to **ReadOnly**, you can only set the consistency level to **0**.
	//
	// >
	//
	// > 	- You can configure the consistency level, transaction splitting, whether the primary node accepts reads, and the connection pool at the same time. For example: `{"ConsistLevel":"1","DistributedTransaction":"on","ConnectionPersist":"Session","MasterAcceptReads":"on"}`.
	//
	// >
	//
	// > 	- The transaction splitting setting is constrained by the consistency level. For example, if the consistency level is set to **0**, you cannot enable transaction splitting. If the consistency level is set to **1*	- or **2**, you can enable transaction splitting.
	//
	// example:
	//
	// {"ConsistLevel": "1","DistributedTransaction": "on"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The type of the custom cluster endpoint. The value is fixed to **Custom**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The read-only nodes to be added to the endpoint. Separate multiple node IDs with commas (,). By default, all nodes are added.
	//
	// example:
	//
	// pi-**************,pi-*************
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The policy for handling global consistency timeouts. Valid values:
	//
	// - 0: Sends the request to the primary node.
	//
	// - 2: Degrades to regular requests. If a global consistency read times out, the query is automatically degraded to a regular request, and the client does not receive an error message.
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
	// - ReadWrite: read and write (automatic read/write splitting).
	//
	// - ReadOnly: read-only. This is the default value.
	//
	// example:
	//
	// ReadOnly
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable the global consistency (high-performance mode) feature for the node. Valid values:
	//
	// - ON: Enables the feature.
	//
	// - OFF: Disables the feature.
	//
	// example:
	//
	// on
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
}

func (s CreateDBClusterEndpointZonalRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterEndpointZonalRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointZonalRequest) GetAutoAddNewNodes() *string {
	return s.AutoAddNewNodes
}

func (s *CreateDBClusterEndpointZonalRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBClusterEndpointZonalRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBClusterEndpointZonalRequest) GetDBEndpointDescription() *string {
	return s.DBEndpointDescription
}

func (s *CreateDBClusterEndpointZonalRequest) GetEndpointConfig() *string {
	return s.EndpointConfig
}

func (s *CreateDBClusterEndpointZonalRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateDBClusterEndpointZonalRequest) GetNodes() *string {
	return s.Nodes
}

func (s *CreateDBClusterEndpointZonalRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBClusterEndpointZonalRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBClusterEndpointZonalRequest) GetPolarSccTimeoutAction() *string {
	return s.PolarSccTimeoutAction
}

func (s *CreateDBClusterEndpointZonalRequest) GetPolarSccWaitTimeout() *string {
	return s.PolarSccWaitTimeout
}

func (s *CreateDBClusterEndpointZonalRequest) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *CreateDBClusterEndpointZonalRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBClusterEndpointZonalRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBClusterEndpointZonalRequest) GetSccMode() *string {
	return s.SccMode
}

func (s *CreateDBClusterEndpointZonalRequest) SetAutoAddNewNodes(v string) *CreateDBClusterEndpointZonalRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetClientToken(v string) *CreateDBClusterEndpointZonalRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetDBClusterId(v string) *CreateDBClusterEndpointZonalRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetDBEndpointDescription(v string) *CreateDBClusterEndpointZonalRequest {
	s.DBEndpointDescription = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetEndpointConfig(v string) *CreateDBClusterEndpointZonalRequest {
	s.EndpointConfig = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetEndpointType(v string) *CreateDBClusterEndpointZonalRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetNodes(v string) *CreateDBClusterEndpointZonalRequest {
	s.Nodes = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetOwnerAccount(v string) *CreateDBClusterEndpointZonalRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetOwnerId(v int64) *CreateDBClusterEndpointZonalRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetPolarSccTimeoutAction(v string) *CreateDBClusterEndpointZonalRequest {
	s.PolarSccTimeoutAction = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetPolarSccWaitTimeout(v string) *CreateDBClusterEndpointZonalRequest {
	s.PolarSccWaitTimeout = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetReadWriteMode(v string) *CreateDBClusterEndpointZonalRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetResourceOwnerAccount(v string) *CreateDBClusterEndpointZonalRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetResourceOwnerId(v int64) *CreateDBClusterEndpointZonalRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) SetSccMode(v string) *CreateDBClusterEndpointZonalRequest {
	s.SccMode = &v
	return s
}

func (s *CreateDBClusterEndpointZonalRequest) Validate() error {
	return dara.Validate(s)
}
