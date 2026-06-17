// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBClusterEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoAddNewNodes(v string) *CreateDBClusterEndpointRequest
	GetAutoAddNewNodes() *string
	SetClientToken(v string) *CreateDBClusterEndpointRequest
	GetClientToken() *string
	SetDBClusterId(v string) *CreateDBClusterEndpointRequest
	GetDBClusterId() *string
	SetDBEndpointDescription(v string) *CreateDBClusterEndpointRequest
	GetDBEndpointDescription() *string
	SetEndpointConfig(v string) *CreateDBClusterEndpointRequest
	GetEndpointConfig() *string
	SetEndpointType(v string) *CreateDBClusterEndpointRequest
	GetEndpointType() *string
	SetNodes(v string) *CreateDBClusterEndpointRequest
	GetNodes() *string
	SetOwnerAccount(v string) *CreateDBClusterEndpointRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateDBClusterEndpointRequest
	GetOwnerId() *int64
	SetPolarFsInstanceId(v string) *CreateDBClusterEndpointRequest
	GetPolarFsInstanceId() *string
	SetPolarSccTimeoutAction(v string) *CreateDBClusterEndpointRequest
	GetPolarSccTimeoutAction() *string
	SetPolarSccWaitTimeout(v string) *CreateDBClusterEndpointRequest
	GetPolarSccWaitTimeout() *string
	SetReadWriteMode(v string) *CreateDBClusterEndpointRequest
	GetReadWriteMode() *string
	SetResourceOwnerAccount(v string) *CreateDBClusterEndpointRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateDBClusterEndpointRequest
	GetResourceOwnerId() *int64
	SetSccMode(v string) *CreateDBClusterEndpointRequest
	GetSccMode() *string
	SetVPCId(v string) *CreateDBClusterEndpointRequest
	GetVPCId() *string
	SetVSwitchId(v string) *CreateDBClusterEndpointRequest
	GetVSwitchId() *string
}

type CreateDBClusterEndpointRequest struct {
	// Specifies whether to automatically add new nodes to the endpoint. Valid values:
	//
	// - **Enable**: Automatically adds new nodes to the endpoint.
	//
	// - **Disable*	- (default): Does not automatically add new nodes to the endpoint.
	//
	// example:
	//
	// Disable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// A client-generated token to ensure the idempotence of the request. The token must be unique, case-sensitive, and a maximum of 64 ASCII characters.
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
	// The name of the custom cluster endpoint.
	//
	// example:
	//
	// test
	DBEndpointDescription *string `json:"DBEndpointDescription,omitempty" xml:"DBEndpointDescription,omitempty"`
	// The advanced configurations for the custom cluster endpoint, specified as a JSON string. You can configure features such as consistency level, transaction splitting, whether the primary node accepts read requests, connection pool, and load balancing policy.
	//
	// - Specifies the load balancing policy. The format is {"LoadBalancePolicy":"policy"}. Valid values:
	//
	//   - **0**: connections-based load balancing (default).
	//
	//   - **1**: active requests-based load balancing.
	//
	// - Specifies the consistency level. The format is `{"ConsistLevel":"level"}`. Valid values:
	//
	//   - **0**: eventual consistency.
	//
	//   - **1**: session consistency (default).
	//
	//   - **2**: global consistency.
	//
	// - Specifies whether to enable transaction splitting. The format is `{"DistributedTransaction":"status"}`. Valid values:
	//
	//   - **on**: enables transaction splitting (default).
	//
	//   - **off**: disables transaction splitting.
	//
	// - Specifies whether the primary node accepts read requests. The format is `{"MasterAcceptReads":"status"}`. Valid values:
	//
	//   - **on**: The primary node accepts read requests.
	//
	//   - **off**: The primary node does not accept read requests (default).
	//
	// - Specifies whether to enable the connection pool. The format is `{"ConnectionPersist":"status"}`. Valid values:
	//
	//   - **off**: disables the connection pool (default).
	//
	//   - **Session**: enables the session-level connection pool.
	//
	//   - **Transaction**: enables the transaction-level connection pool.
	//
	// - Specifies the degree of parallelism for a parallel query. The format is {"MaxParallelDegree":"degree"}. Valid values:
	//
	//   - A specific integer that specifies the degree of parallelism. For example: "MaxParallelDegree":"2".
	//
	//   - **off**: disables parallel query (default).
	//
	// - Specifies whether to enable automatic routing between row store and column store. The format is {"EnableHtapImci":"status"}. Valid values:
	//
	//   - **on**: enables automatic routing.
	//
	//   - **off**: disables automatic routing (default).
	//
	// - Specifies whether to enable overload protection. The format is {"EnableOverloadThrottle":"status"}. Valid values:
	//
	//   - **on**: enables overload protection.
	//
	//   - **off**: disables overload protection (default).
	//
	// > 	- You can configure transaction splitting, whether the primary node accepts read requests, the connection pool, and overload protection only for a PolarDB for MySQL endpoint in **ReadWrite*	- (automatic read/write splitting) mode.
	//
	// >
	//
	// > 	- A PolarDB for MySQL cluster endpoint in **ReadOnly*	- mode supports both **connections-based load balancing*	- and **active requests-based load balancing**. An endpoint in **ReadWrite*	- (automatic read/write splitting) mode supports only **active requests-based load balancing**.
	//
	// >
	//
	// > 	- You can enable automatic routing between row store and column store if the read/write mode of the cluster endpoint for PolarDB for MySQL is **ReadWrite*	- (automatic read/write splitting), or if the read/write mode is **ReadOnly*	- and the load balancing policy is **active requests-based load balancing**.
	//
	// >
	//
	// > 	- Only PolarDB for MySQL clusters support global consistency.
	//
	// >
	//
	// > 	- If you set **ReadWriteMode*	- to **ReadOnly**, the consistency level must be **0*	- (eventual consistency).
	//
	// >
	//
	// > 	- You can configure the consistency level, transaction splitting, whether the primary node accepts read requests, and the connection pool at the same time. Example: `{"ConsistLevel":"1","DistributedTransaction":"on","ConnectionPersist":"Session","MasterAcceptReads":"on"}`.
	//
	// >
	//
	// > 	- The setting for transaction splitting depends on the consistency level. For example, if you set the consistency level to **0*	- (eventual consistency), you cannot enable transaction splitting. If you set the consistency level to **1*	- (session consistency) or **2*	- (global consistency), you can enable transaction splitting.
	//
	// example:
	//
	// {"ConsistLevel": "1","DistributedTransaction": "on"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The type of the custom cluster endpoint. Set the value to **Custom**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The nodes to associate with the endpoint. Separate multiple node IDs with a comma (,). If you omit this parameter, all nodes in the cluster are added by default.
	//
	// > - For a PolarDB for MySQL cluster, specify the node IDs.
	//
	// >
	//
	// > - For a PolarDB for PostgreSQL cluster or a PolarDB for PostgreSQL (compatible with Oracle) cluster, specify the roles of the nodes, such as `Writer,Reader1,Reader2`.
	//
	// >
	//
	// > - If you set **ReadWriteMode*	- to **ReadOnly**, you can associate only one node with the endpoint. If this node fails, the endpoint may be unavailable for up to 1 hour. This configuration is not recommended for a production environment. To improve availability, associate at least two nodes with the endpoint.
	//
	// >
	//
	// > - If you set **ReadWriteMode*	- to **ReadWrite**, you must associate at least two nodes with the endpoint.
	//
	// >   \\	- For a PolarDB for MySQL cluster, you can select any two nodes. If both nodes are read-only nodes, write requests are routed to the primary node.
	//
	// >   \\	- For a PolarDB for PostgreSQL cluster or a PolarDB for PostgreSQL (compatible with Oracle) cluster, the primary node must be included.
	//
	// example:
	//
	// pi-**********,pi-*********
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// pfs-test****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The policy for handling global consistency read timeouts. Valid values:
	//
	// - **0**: Send the request to the primary node.
	//
	// - **2**: Downgrade to a regular request. If a global consistency read times out, the query is automatically downgraded, and the client does not receive an error.
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
	// - **ReadWrite**: read/write (automatic read/write splitting).
	//
	// - **ReadOnly*	- (default): read-only.
	//
	// example:
	//
	// ReadOnly
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable global consistency (high-performance mode). Valid values:
	//
	// - **ON**: Enables the feature.
	//
	// - **OFF**: Disables the feature.
	//
	// example:
	//
	// on
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
	// example:
	//
	// vpc-2zehr7ghqovftils0****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// example:
	//
	// vsw-2ze775gnf7jn33ua****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateDBClusterEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDBClusterEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateDBClusterEndpointRequest) GetAutoAddNewNodes() *string {
	return s.AutoAddNewNodes
}

func (s *CreateDBClusterEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDBClusterEndpointRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateDBClusterEndpointRequest) GetDBEndpointDescription() *string {
	return s.DBEndpointDescription
}

func (s *CreateDBClusterEndpointRequest) GetEndpointConfig() *string {
	return s.EndpointConfig
}

func (s *CreateDBClusterEndpointRequest) GetEndpointType() *string {
	return s.EndpointType
}

func (s *CreateDBClusterEndpointRequest) GetNodes() *string {
	return s.Nodes
}

func (s *CreateDBClusterEndpointRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateDBClusterEndpointRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDBClusterEndpointRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *CreateDBClusterEndpointRequest) GetPolarSccTimeoutAction() *string {
	return s.PolarSccTimeoutAction
}

func (s *CreateDBClusterEndpointRequest) GetPolarSccWaitTimeout() *string {
	return s.PolarSccWaitTimeout
}

func (s *CreateDBClusterEndpointRequest) GetReadWriteMode() *string {
	return s.ReadWriteMode
}

func (s *CreateDBClusterEndpointRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateDBClusterEndpointRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateDBClusterEndpointRequest) GetSccMode() *string {
	return s.SccMode
}

func (s *CreateDBClusterEndpointRequest) GetVPCId() *string {
	return s.VPCId
}

func (s *CreateDBClusterEndpointRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateDBClusterEndpointRequest) SetAutoAddNewNodes(v string) *CreateDBClusterEndpointRequest {
	s.AutoAddNewNodes = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetClientToken(v string) *CreateDBClusterEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetDBClusterId(v string) *CreateDBClusterEndpointRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetDBEndpointDescription(v string) *CreateDBClusterEndpointRequest {
	s.DBEndpointDescription = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetEndpointConfig(v string) *CreateDBClusterEndpointRequest {
	s.EndpointConfig = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetEndpointType(v string) *CreateDBClusterEndpointRequest {
	s.EndpointType = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetNodes(v string) *CreateDBClusterEndpointRequest {
	s.Nodes = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetOwnerAccount(v string) *CreateDBClusterEndpointRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetOwnerId(v int64) *CreateDBClusterEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetPolarFsInstanceId(v string) *CreateDBClusterEndpointRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetPolarSccTimeoutAction(v string) *CreateDBClusterEndpointRequest {
	s.PolarSccTimeoutAction = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetPolarSccWaitTimeout(v string) *CreateDBClusterEndpointRequest {
	s.PolarSccWaitTimeout = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetReadWriteMode(v string) *CreateDBClusterEndpointRequest {
	s.ReadWriteMode = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetResourceOwnerAccount(v string) *CreateDBClusterEndpointRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetResourceOwnerId(v int64) *CreateDBClusterEndpointRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetSccMode(v string) *CreateDBClusterEndpointRequest {
	s.SccMode = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetVPCId(v string) *CreateDBClusterEndpointRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) SetVSwitchId(v string) *CreateDBClusterEndpointRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBClusterEndpointRequest) Validate() error {
	return dara.Validate(s)
}
