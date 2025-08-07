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
	// The prefix of the new endpoint. The prefix must meet the following requirements:
	//
	// 	- It can contain lowercase letters, digits, and hyphens (-).
	//
	// 	- It must start with a letter and end with a digit or a letter.
	//
	// 	- It must be 6 to 30 characters in length.
	//
	// example:
	//
	// example
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty"`
	// The ID of cluster.
	//
	// > You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of the clusters that belong to your Alibaba Cloud account, such as cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-***************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of the endpoint.
	//
	// > You can call the [DescribeDBClusterEndpoints](https://help.aliyun.com/document_detail/98205.html) operation to query endpoint IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pe-****************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// The network type of the endpoint. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	NetType      *string `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The port number. Valid values: 3000 to 5999.
	//
	// > This parameter is valid only for PolarDB for MySQL clusters. If you leave this parameter empty, the default port 3306 is used.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The prefix of the private domain name. The prefix must meet the following requirements:
	//
	// 	- The prefix can contain lowercase letters, digits, and hyphens (-).
	//
	// 	- The prefix must start with a letter and end with a digit or a letter.
	//
	// 	- The prefix must be 6 to 30 characters in length.
	//
	// >- You can bind each internal endpoint of PolarDB to a private domain name. The private domain name takes effect only in the specified virtual private clouds (VPCs) in the current region. Private domain names are managed by using PrivateZone. You can use the CNAME record of PrivateZone to map domain names to PolarDB. You are charged a small fee for this feature. For more information, see [Pricing](https://help.aliyun.com/document_detail/71338.html).
	//
	// >- This parameter takes effect only if you set **NetType*	- to Private.
	//
	// example:
	//
	// aliyundoc
	PrivateZoneAddressPrefix *string `json:"PrivateZoneAddressPrefix,omitempty" xml:"PrivateZoneAddressPrefix,omitempty"`
	// The name of the private zone.
	//
	// > This parameter takes effect only when **NetType*	- is set to Private.
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
