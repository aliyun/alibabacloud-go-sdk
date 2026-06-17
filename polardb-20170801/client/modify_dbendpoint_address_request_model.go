// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBEndpointAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionStringPrefix(v string) *ModifyDBEndpointAddressRequest
	GetConnectionStringPrefix() *string
	SetDBClusterId(v string) *ModifyDBEndpointAddressRequest
	GetDBClusterId() *string
	SetDBEndpointId(v string) *ModifyDBEndpointAddressRequest
	GetDBEndpointId() *string
	SetNetType(v string) *ModifyDBEndpointAddressRequest
	GetNetType() *string
	SetOwnerAccount(v string) *ModifyDBEndpointAddressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBEndpointAddressRequest
	GetOwnerId() *int64
	SetPort(v string) *ModifyDBEndpointAddressRequest
	GetPort() *string
	SetPrivateZoneAddressPrefix(v string) *ModifyDBEndpointAddressRequest
	GetPrivateZoneAddressPrefix() *string
	SetPrivateZoneName(v string) *ModifyDBEndpointAddressRequest
	GetPrivateZoneName() *string
	SetResourceOwnerAccount(v string) *ModifyDBEndpointAddressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBEndpointAddressRequest
	GetResourceOwnerId() *int64
}

type ModifyDBEndpointAddressRequest struct {
	// The new connection string prefix. The prefix must meet the following requirements:
	//
	// - It can contain only lowercase letters, digits, and hyphens (-).
	//
	// - It must start with a letter and end with a letter or a digit.
	//
	// - It must be 6 to 30 characters in length.
	//
	// example:
	//
	// example
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of all clusters in your account, including cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the connection address.
	//
	// > You can call the [DescribeDBClusterEndpoints](https://help.aliyun.com/document_detail/98205.html) operation to query the ID of a connection address.
	//
	// example:
	//
	// pe-****************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The network type of the connection address. Valid values:
	//
	// - **Public**: public network
	//
	// - **Private**: private network
	//
	// <props="china">
	//
	// - **Inner**: classic network
	//
	//
	//
	// <props="china">
	//
	// Only PolarDB for MySQL clusters support the classic network type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port number. The valid range is 3000 to 5999.
	//
	// > - This parameter is supported only for PolarDB for MySQL clusters. If you do not specify this parameter, the port defaults to 3306.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The prefix of the private domain name. The prefix must meet the following requirements:
	//
	// - It can contain only lowercase letters, digits, and hyphens (-).
	//
	// - It must start with a letter and end with a letter or a digit.
	//
	// - It must be 6 to 30 characters in length.
	//
	// > 	- You can bind a private domain name to each private endpoint of a PolarDB cluster. This domain name is effective only in the specified VPC within the current region. The private domain name is managed by PrivateZone and is mapped to the built-in private endpoint of the cluster through a CNAME record. This feature incurs a small fee. For more information, see [Pricing](https://help.aliyun.com/document_detail/71338.html).
	//
	// >
	//
	// > 	- This parameter is valid only when **NetType is set to Private**.
	//
	// example:
	//
	// aliyundoc
	PrivateZoneAddressPrefix *string `json:"PrivateZoneAddressPrefix,omitempty" xml:"PrivateZoneAddressPrefix,omitempty"`
	// The private zone name.
	//
	// > This parameter is valid only when **NetType is set to Private**.
	//
	// example:
	//
	// aliyundoc.com
	PrivateZoneName      *string `json:"PrivateZoneName,omitempty" xml:"PrivateZoneName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBEndpointAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBEndpointAddressRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBEndpointAddressRequest) GetConnectionStringPrefix() *string {
	return s.ConnectionStringPrefix
}

func (s *ModifyDBEndpointAddressRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBEndpointAddressRequest) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *ModifyDBEndpointAddressRequest) GetNetType() *string {
	return s.NetType
}

func (s *ModifyDBEndpointAddressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBEndpointAddressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBEndpointAddressRequest) GetPort() *string {
	return s.Port
}

func (s *ModifyDBEndpointAddressRequest) GetPrivateZoneAddressPrefix() *string {
	return s.PrivateZoneAddressPrefix
}

func (s *ModifyDBEndpointAddressRequest) GetPrivateZoneName() *string {
	return s.PrivateZoneName
}

func (s *ModifyDBEndpointAddressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBEndpointAddressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBEndpointAddressRequest) SetConnectionStringPrefix(v string) *ModifyDBEndpointAddressRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetDBClusterId(v string) *ModifyDBEndpointAddressRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetDBEndpointId(v string) *ModifyDBEndpointAddressRequest {
	s.DBEndpointId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetNetType(v string) *ModifyDBEndpointAddressRequest {
	s.NetType = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetOwnerAccount(v string) *ModifyDBEndpointAddressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetOwnerId(v int64) *ModifyDBEndpointAddressRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetPort(v string) *ModifyDBEndpointAddressRequest {
	s.Port = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetPrivateZoneAddressPrefix(v string) *ModifyDBEndpointAddressRequest {
	s.PrivateZoneAddressPrefix = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetPrivateZoneName(v string) *ModifyDBEndpointAddressRequest {
	s.PrivateZoneName = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetResourceOwnerAccount(v string) *ModifyDBEndpointAddressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) SetResourceOwnerId(v int64) *ModifyDBEndpointAddressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBEndpointAddressRequest) Validate() error {
	return dara.Validate(s)
}
