// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAddressType(v string) *CreateDatabaseRequest
	GetActiveAddressType() *string
	SetComment(v string) *CreateDatabaseRequest
	GetComment() *string
	SetDatabaseName(v string) *CreateDatabaseRequest
	GetDatabaseName() *string
	SetDatabasePort(v int32) *CreateDatabaseRequest
	GetDatabasePort() *int32
	SetDatabasePrivateAddress(v string) *CreateDatabaseRequest
	GetDatabasePrivateAddress() *string
	SetDatabasePublicAddress(v string) *CreateDatabaseRequest
	GetDatabasePublicAddress() *string
	SetDatabaseType(v string) *CreateDatabaseRequest
	GetDatabaseType() *string
	SetInstanceId(v string) *CreateDatabaseRequest
	GetInstanceId() *string
	SetNetworkDomainId(v string) *CreateDatabaseRequest
	GetNetworkDomainId() *string
	SetPolarDBEndpointType(v string) *CreateDatabaseRequest
	GetPolarDBEndpointType() *string
	SetRegionId(v string) *CreateDatabaseRequest
	GetRegionId() *string
	SetSource(v string) *CreateDatabaseRequest
	GetSource() *string
	SetSourceInstanceId(v string) *CreateDatabaseRequest
	GetSourceInstanceId() *string
	SetSourceInstanceRegionId(v string) *CreateDatabaseRequest
	GetSourceInstanceRegionId() *string
}

type CreateDatabaseRequest struct {
	// The address type of the database to add. Valid values:
	//
	// 	- Public
	//
	// 	- Private
	//
	// This parameter is required.
	//
	// example:
	//
	// Public
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The remarks of the database to add. The remarks can be up to 500 characters in length.
	//
	// example:
	//
	// cpp
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The name of the database to add. This parameter is required if Source is set to **Local**.
	//
	// example:
	//
	// Oracle
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The port of the database. This parameter is required if Source is set to **Local**.
	//
	// example:
	//
	// 5433
	DatabasePort *int32 `json:"DatabasePort,omitempty" xml:"DatabasePort,omitempty"`
	// The internal IP address of the database. Specify an IPv4 address or a domain name.
	//
	// >  This parameter is required if ActiveAddressType is set to Private.
	//
	// example:
	//
	// pgm-uf6o******
	DatabasePrivateAddress *string `json:"DatabasePrivateAddress,omitempty" xml:"DatabasePrivateAddress,omitempty"`
	// The public IP address of the database. Specify an IPv4 address or a domain name.
	//
	// >  This parameter is required if ActiveAddressType is set to Public.
	//
	// example:
	//
	// rm-uf65251k51******
	DatabasePublicAddress *string `json:"DatabasePublicAddress,omitempty" xml:"DatabasePublicAddress,omitempty"`
	// The type of the database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **Oracle**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The bastion host ID.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-7mz2g5hu20e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain to which the database to add belongs.
	//
	// >  You can call the [ListNetworkDomains](https://help.aliyun.com/document_detail/2758827.html) operation to query the network domain ID.
	//
	// example:
	//
	// 1
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The endpoint type of the PolarDB database. This parameter is required if Source is set to PolarDB. Valid values:
	//
	// 	- Cluster
	//
	// 	- Primary
	//
	// example:
	//
	// Cluster
	PolarDBEndpointType *string `json:"PolarDBEndpointType,omitempty" xml:"PolarDBEndpointType,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the database to add. Valid values:
	//
	// 	- Local: on-premises database.
	//
	// 	- Rds: ApsaraDB RDS instance.
	//
	// 	- PolarDB: PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The instance ID of the database to add.
	//
	// > This parameter is required if **Source*	- is set to **Rds*	- or **PolarDB**.
	//
	// example:
	//
	// i-bp19ienyt0yax748****
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The region ID of the database to add.
	//
	// >  This parameter is required if **Source*	- is set to **Rds*	- or **PolarDB**.
	//
	// example:
	//
	// cn-shanghai
	SourceInstanceRegionId *string `json:"SourceInstanceRegionId,omitempty" xml:"SourceInstanceRegionId,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *CreateDatabaseRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDatabaseRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *CreateDatabaseRequest) GetDatabasePort() *int32 {
	return s.DatabasePort
}

func (s *CreateDatabaseRequest) GetDatabasePrivateAddress() *string {
	return s.DatabasePrivateAddress
}

func (s *CreateDatabaseRequest) GetDatabasePublicAddress() *string {
	return s.DatabasePublicAddress
}

func (s *CreateDatabaseRequest) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *CreateDatabaseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDatabaseRequest) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *CreateDatabaseRequest) GetPolarDBEndpointType() *string {
	return s.PolarDBEndpointType
}

func (s *CreateDatabaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDatabaseRequest) GetSource() *string {
	return s.Source
}

func (s *CreateDatabaseRequest) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *CreateDatabaseRequest) GetSourceInstanceRegionId() *string {
	return s.SourceInstanceRegionId
}

func (s *CreateDatabaseRequest) SetActiveAddressType(v string) *CreateDatabaseRequest {
	s.ActiveAddressType = &v
	return s
}

func (s *CreateDatabaseRequest) SetComment(v string) *CreateDatabaseRequest {
	s.Comment = &v
	return s
}

func (s *CreateDatabaseRequest) SetDatabaseName(v string) *CreateDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateDatabaseRequest) SetDatabasePort(v int32) *CreateDatabaseRequest {
	s.DatabasePort = &v
	return s
}

func (s *CreateDatabaseRequest) SetDatabasePrivateAddress(v string) *CreateDatabaseRequest {
	s.DatabasePrivateAddress = &v
	return s
}

func (s *CreateDatabaseRequest) SetDatabasePublicAddress(v string) *CreateDatabaseRequest {
	s.DatabasePublicAddress = &v
	return s
}

func (s *CreateDatabaseRequest) SetDatabaseType(v string) *CreateDatabaseRequest {
	s.DatabaseType = &v
	return s
}

func (s *CreateDatabaseRequest) SetInstanceId(v string) *CreateDatabaseRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDatabaseRequest) SetNetworkDomainId(v string) *CreateDatabaseRequest {
	s.NetworkDomainId = &v
	return s
}

func (s *CreateDatabaseRequest) SetPolarDBEndpointType(v string) *CreateDatabaseRequest {
	s.PolarDBEndpointType = &v
	return s
}

func (s *CreateDatabaseRequest) SetRegionId(v string) *CreateDatabaseRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDatabaseRequest) SetSource(v string) *CreateDatabaseRequest {
	s.Source = &v
	return s
}

func (s *CreateDatabaseRequest) SetSourceInstanceId(v string) *CreateDatabaseRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *CreateDatabaseRequest) SetSourceInstanceRegionId(v string) *CreateDatabaseRequest {
	s.SourceInstanceRegionId = &v
	return s
}

func (s *CreateDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
