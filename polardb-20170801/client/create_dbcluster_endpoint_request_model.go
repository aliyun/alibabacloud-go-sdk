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
	// 	- **Enable**: New nodes are automatically added to the endpoint.
	//
	// 	- **Disable*	- (default): New nodes are not automatically added to the endpoint.
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
	// The advanced settings of the cluster endpoint. The value must be in JSON format. You can set the consistency level, transaction splitting, offload reads from primary node, and connection pool.
	//
	// 	- Set the load balancing policy. Format: `{"LoadBalancePolicy":"load balancing policy"}`. Valid values:
	//
	//    	- **0**: load balancing based on the number of connections (default).
	//
	//    	- **1**: load balancing based on the number of active requests.
	//
	// 	- Set the consistency level. Format: `{"ConsistLevel":"consistency level"}`. Valid values:
	//
	//     	- **0**: eventual consistency.
	//
	//     	- **1**: session consistency (default).
	//
	//     	- **2**: global consistency.
	//
	// 	- Set transaction splitting. Format: `{"DistributedTransaction":"transaction splitting"}`. Valid values:
	//
	//     	- **on**: enables transaction splitting (default).
	//
	//     	- **off**: shutdown transaction splitting.
	//
	// 	- Set whether the primary node accepts read requests. Format: `{"MasterAcceptReads":"offload reads from primary node"}`. Valid values:
	//
	//     	- **on**: The primary node accepts read requests.
	//
	//     	- **off**: The primary node does not accept read requests (default).
	//
	// 	- Set the connection pool. Format: `{"ConnectionPersist":"connection pool"}`. Valid values:
	//
	//     	- **off**: shutdown the connection pool (default).
	//
	//     	- **Session**: enables the session-level connection pool.
	//
	//     	- **Transaction**: enables transaction-level connection pooling.
	//
	// 	- Set parallel query. Format: `{"MaxParallelDegree":"parallel query"}`. Valid values:
	//
	//     	- A specific concurrent count. Example: `"MaxParallelDegree":"2"`.
	//
	//     	- **off**: shutdown parallel query (default).
	//
	// 	- Set automatic request distribution among row store and column store nodes. Format: `{"EnableHtapImci":"automatic request distribution"}`. Valid values:
	//
	//     	- **on**: enables the automatic request distribution feature among row store and column store nodes.
	//
	//     	- **off**: shutdown the automatic request distribution feature among row store and column store nodes (default).
	//
	// 	- Set whether to enable overload protection. Format: `{"EnableOverloadThrottle":"whether to enable overload protection"}`. Valid values:
	//
	//    	- **on**: enables overload protection.
	//
	//    	- **off**: shutdown overload protection (default).
	//
	// > 	- Transaction splitting, offloading reads from the primary node, connection pool, and overload protection can be set only when the read/write pattern of the PolarDB for MySQL cluster endpoint is set to **read and write (automatic read/write splitting)**.
	//
	// >	- When the read/write pattern of the PolarDB for MySQL cluster endpoint is set to **read-only**, both **load balancing based on the number of connections*	- and **load balancing based on the number of active requests*	- are supported. For cluster endpoints in the **read and write (automatic read/write splitting) pattern**, only **load balancing based on the number of active requests*	- is supported.
	//
	// >	- Automatic request distribution among row store and column store nodes can be set when the read/write pattern of the PolarDB for MySQL cluster endpoint is set to **read and write (automatic read/write splitting)**, or when the read/write pattern is set to **read-only*	- and the load balancing policy is set to **load balancing based on the number of active requests**.
	//
	// >	- Only PolarDB for MySQL supports setting the consistency level to global consistency.
	//
	// >	- If the parameter **ReadWriteMode*	- is set to **ReadOnly**, the consistency level can only be set to **0**.
	//
	// >	- You can set the consistency level, transaction splitting, offloading reads from the primary node, and connection pool at the same time. Example: `{"ConsistLevel":"1","DistributedTransaction":"on","ConnectionPersist":"Session","MasterAcceptReads":"on"}`.
	//
	// >	- The transaction splitting settings are constrained by the consistency level settings. For example, if the consistency level is set to **0**, transaction splitting cannot be enabled. If the consistency level is set to **1*	- or **2**, transaction splitting can be enabled.
	//
	// > Note: Clusters with DeployUnit=Advanced (hot standby/multi-master) may not support some advanced endpoint settings. If the OperationDenied.InstanceType error occurs, remove EndpointConfig and retry, or confirm the compatibility of the target cluster deployment type with the selected settings.
	//
	// example:
	//
	// {"ConsistLevel": "1","DistributedTransaction": "on"}
	EndpointConfig *string `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty"`
	// The type of the custom cluster endpoint. Valid values:
	//
	// - **Custom**: custom cluster endpoint.
	//
	// - **DynamoDB**: DynamoDB endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// The read-only nodes to be added to the endpoint. Separate multiple nodes with commas (,). Default value: all nodes.
	//
	// > 	- For PolarDB for MySQL, specify node IDs.
	//
	// >	- For PolarDB for PostgreSQL and PolarDB for PostgreSQL (Compatible with Oracle), specify node role names, such as `Writer,Reader1,Reader2`.
	//
	// >	- If **ReadWriteMode*	- is set to **ReadOnly**, you can add only one node. However, if this node fails, the endpoint may be unavailable for up to 1 hour. Do not use this configuration in production environments. Select at least two nodes to improve availability.
	//
	// >	- If **ReadWriteMode*	- is set to **ReadWrite**, select at least two nodes.
	//
	//     	- For PolarDB for MySQL, you can select any two nodes. If both nodes are read-only nodes, write requests are forwarded to the primary node.
	//
	//     	- For PolarDB for PostgreSQL and PolarDB for PostgreSQL (Compatible with Oracle), the primary node must be included.
	//
	// example:
	//
	// pi-**********,pi-*********
	Nodes        *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The PolarFS instance ID.
	//
	// example:
	//
	// pfs-test****
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The timeout policy for global consistency. Valid values:
	//
	// - **0**: sends the request to the primary node.
	//
	// - **2**: degrades the request upon timeout. If a global consistency read times out, the query is automatically downgraded to regular requests, and the client does not receive an error message.
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
	// - **ReadWrite**: read and write (automatic read/write splitting).
	//
	// - **ReadOnly*	- (default): read-only.
	//
	// example:
	//
	// ReadOnly
	ReadWriteMode        *string `json:"ReadWriteMode,omitempty" xml:"ReadWriteMode,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable the global consistency (high-performance mode) feature for the node. Valid values:
	//
	// - **ON**: Enabled.
	//
	// - **OFF**: shutdown
	//
	// example:
	//
	// on
	SccMode *string `json:"SccMode,omitempty" xml:"SccMode,omitempty"`
	// vpc-2zehr7ghqovft******
	//
	// Virtual private cloud identifier.
	//
	// example:
	//
	// vpc-2zehr7ghqovftils0****
	VPCId *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	// The vSwitch ID.
	//
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
