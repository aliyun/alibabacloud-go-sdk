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
	// - Disable: New nodes are not automatically added to the endpoint. This is the default value.
	//
	// example:
	//
	// Disable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// The client token that is used to ensure the idempotence of the request and prevent duplicate submissions. You can generate the value of this parameter on the client side. The value must be unique across different requests and is case-sensitive. The value can contain up to 64 ASCII characters.
	//
	// example:
	//
	// 6000170000591aed949d0f******************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the PolarDB on ENS cluster. You can call the DescribeDBClustersZonal operation to obtain the ID. Standard PolarDB cluster IDs are not supported. Usage notes: This operation applies only to PolarDB on ENS clusters. It outlines the concepts of edge clouds, activation methods, and how to obtain the edge cluster ID, noting that the CreateDBClusterZonal operation is unavailable and clusters must be created through alternative methods.
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
	// The advanced configurations of the cluster endpoint. The value is a JSON string. Currently, settings for the consistency level, transaction splitting, offload reads from primary node, and connection pool are supported.
	//
	// 	- Configure the load balancing policy. Format: {"LoadBalancePolicy":"load balancing policy"}. Valid values:
	//
	//    	- **0**: load balancing based on the number of connections. This is the default value.
	//
	//    	- **1**: load balancing based on the number of active requests.
	//
	// 	- Configure the consistency level. Format: `{"ConsistLevel":"consistency level"}`. Valid values:
	//
	//     	- **0**: eventual consistency.
	//
	//     	- **1**: session consistency. This is the default value.
	//
	//     	- **2**: global consistency.
	//
	// 	- Configure transaction splitting. Format: `{"DistributedTransaction":"transaction splitting"}`. Valid values:
	//
	//     	- **on**: enables transaction splitting. This is the default value.
	//
	//     	- **off**: disables transaction splitting.
	//
	// 	- Configure whether the primary node accepts read requests. Format: `{"MasterAcceptReads":"offload reads from primary node"}`. Valid values:
	//
	//     	- **on**: The primary node accepts read requests.
	//
	//     	- **off**: The primary node does not accept read requests. This is the default value.
	//
	// 	- Configure the connection pool. Format: `{"ConnectionPersist":"connection pool"}`. Valid values:
	//
	//     	- **off**: disables the connection pool. This is the default value.
	//
	//     	- **Session**: enables the session-level connection pool.
	//
	//     	- **Transaction**: enables the transaction-level connection pool.
	//
	// 	- Configure parallel query. Format: {"MaxParallelDegree":"parallel query"}. Valid values:
	//
	//     	- A specific number of concurrent queries. Example: "MaxParallelDegree":"2".
	//
	//     	- **off**: disables parallel query. This is the default value.
	//
	// 	- Configure automatic request distribution among row store and column store nodes. Format: {"EnableHtapImci":"automatic request distribution among row store and column store nodes"}. Valid values:
	//
	//     	- **on**: enables automatic request distribution among row store and column store nodes.
	//
	//     	- **off**: disables automatic request distribution among row store and column store nodes. This is the default value.
	//
	// 	- Configure whether to enable overload protection. Format: {"EnableOverloadThrottle":"overload protection"}. Valid values:
	//
	//    	- **on**: enables overload protection.
	//
	//    	- **off**: disables overload protection. This is the default value.
	//
	// > 	- Transaction splitting, whether the primary node accepts read requests, the connection pool, and overload protection can be configured only when the read/write mode of the PolarDB for MySQL cluster endpoint is set to **read and write (automatic read/write splitting)**.
	//
	// >	- When the read/write mode of the PolarDB for MySQL cluster endpoint is set to **read-only**, both **load balancing based on the number of connections*	- and **load balancing based on the number of active requests*	- are supported. Cluster endpoints in the **read and write (automatic read/write splitting) mode*	- only support the **load balancing based on the number of active requests*	- policy.
	//
	// >	- Automatic request distribution among row store and column store nodes can be configured when the read/write mode of the PolarDB for MySQL cluster endpoint is set to **read and write (automatic read/write splitting)**, or when the read/write mode is set to **read-only*	- and the load balancing policy is set to **load balancing based on the number of active requests**.
	//
	// >	- Only PolarDB for MySQL supports setting the consistency level to global consistency.
	//
	// >	- If the ReadWriteMode parameter is set to **ReadOnly**, the consistency level can only be set to **0**.
	//
	// >	- You can configure the consistency level, transaction splitting, whether the primary node accepts read requests, and the connection pool at the same time. Example: `{"ConsistLevel":"1","DistributedTransaction":"on","ConnectionPersist":"Session","MasterAcceptReads":"on"}`.
	//
	// >	- The configuration of transaction splitting is constrained by the consistency level. For example, if the consistency level is set to **0**, transaction splitting cannot be enabled. If the consistency level is set to **1*	- or **2**, transaction splitting can be enabled.
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
	// The read-only nodes to add to the endpoint. Separate multiple node IDs with commas (,). By default, all nodes are added.
	//
	// example:
	//
	// pi-**************,pi-*************
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The timeout policy for global consistency. Valid values:
	//
	// - 0: sends the request to the primary node.
	//
	// - 2: degrades the request upon timeout. If a global consistency read times out, the query is automatically downgraded to regular requests, and the client does not receive an error message.
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
	// - ON: enables the feature.
	//
	// - OFF: disables the feature.
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
