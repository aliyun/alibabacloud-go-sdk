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
	// Specifies whether to automatically add new nodes to the endpoint. Valid values:
	//
	// - Enable: New nodes are automatically added to the endpoint.
	//
	// - Disable (default): New nodes are not automatically added to the endpoint.
	//
	// example:
	//
	// Disable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and is case-sensitive. It can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 6000170000591aed949d0f******************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the ENS edge cloud cluster. You can call DescribeDBClustersZonal to obtain the ID. Standard PolarDB cluster IDs are not applicable. Operation description: Prerequisites for edge cloud clusters: This API applies only to ENS edge cloud clusters and is not applicable to standard PolarDB clusters. It provides an overview of edge clouds, activation methods, and instructions on how to obtain the edge cluster ID, noting that the CreateDBClusterZonal API is currently unavailable and clusters must be created through alternative methods.
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
	// The advanced configurations of the cluster endpoint. The value is a JSON string. Supported settings include the consistency level, transaction splitting, offload reads from primary node, and connection pool.
	//
	// 	- Set the load balancing policy. Format: `{"LoadBalancePolicy":"Load balancing policy"}`. Valid values:
	//
	//    	- **0**: Load balancing based on the number of connections (default).
	//
	//    	- **1**: Load balancing based on the number of active requests.
	//
	// 	- Set the consistency level. Format: `{"ConsistLevel":"Consistency level"}`. Valid values:
	//
	//     	- **0**: Eventual consistency.
	//
	//     	- **1**: Session consistency (default).
	//
	//     	- **2**: Global consistency.
	//
	// 	- Set transaction splitting. Format: `{"DistributedTransaction":"Transaction splitting"}`. Valid values:
	//
	//     	- **on**: Enable transaction splitting (default).
	//
	//     	- **off**: Disable transaction splitting.
	//
	// 	- Set whether the primary node accepts read requests. Format: `{"MasterAcceptReads":"offload reads from primary node"}`. Valid values:
	//
	//     	- **on**: The primary node accepts read requests.
	//
	//     	- **off**: The primary node does not accept read requests (default).
	//
	// 	- Set the connection pool. Format: `{"ConnectionPersist":"Connection pool"}`. Valid values:
	//
	//     	- **off**: Disable the connection pool (default).
	//
	//     	- **Session**: Enable the session-level connection pool.
	//
	//     	- **Transaction**: Enable the transaction-level connection pool.
	//
	// 	- Set parallel query. Format: `{"MaxParallelDegree":"Parallel query"}`. Valid values:
	//
	//     	- A specific degree of parallelism. Example: `"MaxParallelDegree":"2"`.
	//
	//     	- **off**: Disable parallel query (default).
	//
	// 	- Set automatic request distribution among row store and column store nodes. Format: `{"EnableHtapImci":"Automatic request distribution among row store and column store nodes"}`. Valid values:
	//
	//     	- **on**: Enable automatic request distribution among row store and column store nodes.
	//
	//     	- **off**: Disable automatic request distribution among row store and column store nodes (default).
	//
	// 	- Set whether to enable overload protection. Format: `{"EnableOverloadThrottle":"Whether to enable overload protection"}`. Valid values:
	//
	//    	- **on**: Enable overload protection.
	//
	//    	- **off**: Disable overload protection (default).
	//
	// > 	- Transaction splitting, offloading reads from the primary node, the connection pool, and overload protection can be configured only when the read/write pattern of the PolarDB for MySQL cluster endpoint is set to **Read and write (automatic read/write splitting)**.
	//
	// >	- When the read/write pattern of the PolarDB for MySQL cluster endpoint is set to **Read-only**, both **load balancing based on the number of connections*	- and **load balancing based on the number of active requests*	- are supported. Cluster endpoints in the **read and write (automatic read/write splitting) pattern*	- support only **load balancing based on the number of active requests**.
	//
	// >	- Automatic request distribution among row store and column store nodes can be configured when the read/write pattern of the PolarDB for MySQL cluster endpoint is set to **Read and write (automatic read/write splitting)**, or when the read/write pattern is set to **Read-only*	- and the load balancing policy is set to **load balancing based on the number of active requests**.
	//
	// >	- Only PolarDB for MySQL supports setting the consistency level to global consistency.
	//
	// >	- If the **ReadWriteMode*	- parameter is set to **ReadOnly**, the consistency level can only be set to **0**.
	//
	// >	- You can apply settings for the consistency level, transaction splitting, offloading reads from the primary node, and the connection pool at the same time. Example: `{"ConsistLevel":"1","DistributedTransaction":"on","ConnectionPersist":"Session","MasterAcceptReads":"on"}`.
	//
	// >	- The configuration of transaction splitting is restricted by the consistency level. For example, if the consistency level is set to **0**, transaction splitting cannot be enabled. If the consistency level is set to **1*	- or **2**, transaction splitting can be enabled.
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
	// The nodes that handle read requests to be added to the target endpoint. Separate multiple nodes with commas (,). By default, all nodes are added.
	//
	// example:
	//
	// pi-**************,pi-*************
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The timeout policy for global consistency. Valid values:
	//
	// - 0: Send the request to the primary node.
	//
	// - 2: Timeout degradation. If a global consistency read times out, the query is automatically downgraded to regular requests, and the client does not receive an error message.
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
	// - ReadWrite: Read and write (automatic read/write splitting).
	//
	// - ReadOnly (default): Read-only.
	//
	// example:
	//
	// ReadOnly
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable the global consistency (high-performance mode) feature for the node. Valid values:
	//
	// - ON: Enabled.
	//
	// - OFF: Disabled.
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
