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
}

type CreateDBClusterEndpointRequest struct {
	// Specifies whether to enable automatic association of newly added nodes with the cluster endpoint. Valid values:
	//
	// 	- **Enable**
	//
	// 	- **Disable*	- (default)
	//
	// example:
	//
	// Disable
	AutoAddNewNodes *string `json:"AutoAddNewNodes,omitempty" xml:"AutoAddNewNodes,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. The token is case-sensitive.
	//
	// example:
	//
	// 6000170000591aed949d0f******************
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of cluster.
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
	// The advanced configurations of the cluster endpoint. You must specify the configurations in the JSON format. You can specify the configurations of the following attributes: consistency level, transaction splitting, connection pool, and offload reads from the primary node.
	//
	// 	- Specify the consistency level in the format of `{"ConsistLevel":"Consistency level"}`. Default value: 1. Valid values:
	//
	//     	- **0**: eventual consistency.
	//
	//     	- **1**: session consistency.
	//
	//     	- **2**: global consistency.
	//
	// 	- Specify the transaction splitting feature in the format of `{"DistributedTransaction":"Transaction splitting"}`. Valid values:
	//
	//     	- **on**: enables the transaction splitting feature. By default, the feature is enabled.
	//
	//     	- **off**: disables the transaction splitting feature.
	//
	// 	- Specify the offload reads from the primary node feature in the format of `{"MasterAcceptReads":"Offload reads from the primary node"}`. Default value: off. Valid values:
	//
	//     	- **on**: The primary node accepts read requests.
	//
	//     	- **off**: The primary node does not accept read requests.
	//
	// 	- Specify the connection pool in the format of `{"ConnectionPersist":"Connection pool"}`. Default value: off. Valid values:
	//
	//     	- **off**: disables the connection pool.
	//
	//     	- **Session**: enables the session-level connection pool.
	//
	//     	- **Transaction**: enables the transaction-level connection pool.
	//
	// >- You can specify the transaction splitting, connection pool, and offload reads from the primary node features for a PolarDB for MySQL cluster only if ReadWriteMode is set to ReadWrite for the cluster endpoint.
	//
	// >- Only PolarDB for MySQL supports global consistency.
	//
	// >- If the **ReadWriteMode*	- parameter is set to **ReadOnly**, the consistency level must be **0**.
	//
	// >- You can use one record to specify the consistency level, transaction splitting, connection pool, and offload reads from the primary node features, such as `{"ConsistLevel":"1","DistributedTransaction":"on","ConnectionPersist":"Session","MasterAcceptReads":"on"}`.
	//
	// >- The transaction splitting settings are restricted by the consistency level settings. For example, if you set the consistency level to **0**, transaction splitting cannot be enabled. If you set the consistency level to **1*	- or **2**, transaction splitting can be enabled.
	//
	// example:
	//
	// {"ConsistLevel": "1","DistributedTransaction": "on"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The type of the cluster endpoint. Set the value to **Custom**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The reader nodes that you want to associate with the endpoint. If you want to specify multiple reader nodes, separate the reader nodes with commas (,). If you do not specify this parameter, all nodes are used.
	//
	// >- You need to specify the node IDs for a PolarDB for MySQL cluster.
	//
	// >- You need to specify the role name of each node for a PolarDB for PostgreSQL cluster or a PolarDB for PostgreSQL (Compatible with Oracle) cluster. Example: `Writer, Reader1, Reader2`.
	//
	// >- If you set **ReadWriteMode*	- to **ReadOnly**, you can associate only one node with the endpoint. If the only node becomes faulty, the cluster endpoint may be unavailable for up to 1 hour. We recommend that you associate more than one node with the cluster endpoint in production environments. We recommend that you associate at least two nodes with the cluster endpoint to improve service availability.
	//
	// >- If you set **ReadWriteMode*	- to **ReadWrite**, you need to associate at least two nodes with the cluster endpoint.
	//
	// >- PolarDB for MySQL does not impose a limit on the types of the two nodes. If the nodes are read-only nodes, write requests are forwarded to the primary node.
	//
	// >- PolarDB for PostgreSQL and PolarDB for PostgreSQL (compatible with Oracle) require a primary node.
	//
	// example:
	//
	// pi-**********,pi-*********
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Global consistency timeout strategy. The value range is as follows:
	//
	// - **0**: Send the request to the primary node
	//
	// - **2**: Timeout degradation, when a global consistency read times out, the query operation will automatically degrade to an inconsistent read, and the client will not receive an error message
	//
	// example:
	//
	// 0
	PolarSccTimeoutAction *string `json:"PolarSccTimeoutAction,omitempty" xml:"PolarSccTimeoutAction,omitempty"`
	// Global consistency timeout
	//
	// example:
	//
	// 100
	PolarSccWaitTimeout *string `json:"PolarSccWaitTimeout,omitempty" xml:"PolarSccWaitTimeout,omitempty"`
	// The read/write mode. Valid values:
	//
	// 	- **ReadWrite**: receives and forwards read and write requests. Automatic read/write splitting is enabled.
	//
	// 	- **ReadOnly**: The account has the read-only permissions on the database.
	//
	// Default value: **ReadOnly**.
	//
	// example:
	//
	// ReadOnly
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

func (s *CreateDBClusterEndpointRequest) Validate() error {
	return dara.Validate(s)
}
