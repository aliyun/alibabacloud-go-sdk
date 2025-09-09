// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAddressType(v string) *ModifyDatabaseRequest
	GetActiveAddressType() *string
	SetComment(v string) *ModifyDatabaseRequest
	GetComment() *string
	SetDatabaseId(v string) *ModifyDatabaseRequest
	GetDatabaseId() *string
	SetDatabaseName(v string) *ModifyDatabaseRequest
	GetDatabaseName() *string
	SetDatabasePort(v string) *ModifyDatabaseRequest
	GetDatabasePort() *string
	SetDatabasePrivateAddress(v string) *ModifyDatabaseRequest
	GetDatabasePrivateAddress() *string
	SetDatabasePublicAddress(v string) *ModifyDatabaseRequest
	GetDatabasePublicAddress() *string
	SetInstanceId(v string) *ModifyDatabaseRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *ModifyDatabaseRequest
	GetNetworkDomainId() *string
	SetRegionId(v string) *ModifyDatabaseRequest
	GetRegionId() *string
	SetSourceInstanceId(v string) *ModifyDatabaseRequest
	GetSourceInstanceId() *string
}

type ModifyDatabaseRequest struct {
	// The new address type of the database. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Public
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The new remarks of the database.
	//
	// example:
	//
	// tttttttt
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the database to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The new name of the database.
	//
	// example:
	//
	// pgsql
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The new port of the database.
	//
	// example:
	//
	// 5433
	DatabasePort *string `json:"DatabasePort,omitempty" xml:"DatabasePort,omitempty"`
	// The new internal address of the database. Specify an IPv4 address or a domain name.
	//
	// example:
	//
	// pc-bp169******
	DatabasePrivateAddress *string `json:"DatabasePrivateAddress,omitempty" xml:"DatabasePrivateAddress,omitempty"`
	// The new public address of the database. Specify an IPv4 address or a domain name.
	//
	// example:
	//
	// pgm-uf6c******
	DatabasePublicAddress *string `json:"DatabasePublicAddress,omitempty" xml:"DatabasePublicAddress,omitempty"`
	// The ID of the bastion host that manages the database to modify.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-72137xe5n01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the new network domain for the database.
	//
	// >  You can call the [ListNetworkDomains](https://help.aliyun.com/document_detail/2758827.html) operation to query the network domain ID.
	//
	// example:
	//
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The region ID of the bastion host that manages the database to modify.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the ApsaraDB for RDS instance or PolarDB cluster to modify.
	//
	// > This parameter is required if **Source*	- is set to **Rds*	- or **PolarDB**.
	//
	// example:
	//
	// i-wz99nexqd62z3bvuvpz5
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s ModifyDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseRequest) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ModifyDatabaseRequest) GetComment() *string {
	return s.Comment
}

func (s *ModifyDatabaseRequest) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ModifyDatabaseRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ModifyDatabaseRequest) GetDatabasePort() *string {
	return s.DatabasePort
}

func (s *ModifyDatabaseRequest) GetDatabasePrivateAddress() *string {
	return s.DatabasePrivateAddress
}

func (s *ModifyDatabaseRequest) GetDatabasePublicAddress() *string {
	return s.DatabasePublicAddress
}

func (s *ModifyDatabaseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDatabaseRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ModifyDatabaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDatabaseRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ModifyDatabaseRequest) SetActiveAddressType(v string) *ModifyDatabaseRequest {
	s.ActiveAddressType = &v
	return s
}

func (s *ModifyDatabaseRequest) SetComment(v string) *ModifyDatabaseRequest {
	s.Comment = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDatabaseId(v string) *ModifyDatabaseRequest {
	s.DatabaseId = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDatabaseName(v string) *ModifyDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDatabasePort(v string) *ModifyDatabaseRequest {
	s.DatabasePort = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDatabasePrivateAddress(v string) *ModifyDatabaseRequest {
	s.DatabasePrivateAddress = &v
	return s
}

func (s *ModifyDatabaseRequest) SetDatabasePublicAddress(v string) *ModifyDatabaseRequest {
	s.DatabasePublicAddress = &v
	return s
}

func (s *ModifyDatabaseRequest) SetInstanceId(v string) *ModifyDatabaseRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDatabaseRequest) SetNetworkDomainId(v string) *ModifyDatabaseRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *ModifyDatabaseRequest) SetRegionId(v string) *ModifyDatabaseRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDatabaseRequest) SetSourceInstanceId(v string) *ModifyDatabaseRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *ModifyDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
