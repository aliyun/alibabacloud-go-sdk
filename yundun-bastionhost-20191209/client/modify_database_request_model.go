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
	SetProjectId(v int64) *ModifyDatabaseRequest
	GetProjectId() *int64
	SetRegionId(v string) *ModifyDatabaseRequest
	GetRegionId() *string
	SetSourceInstanceId(v string) *ModifyDatabaseRequest
	GetSourceInstanceId() *string
}

type ModifyDatabaseRequest struct {
	// The address type of the database after the modification. Valid values:
	//
	// - **Public**: public network address
	//
	// - **Private**: private network address
	//
	// example:
	//
	// Public
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the database after the modification.
	//
	// example:
	//
	// tttttttt
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the database to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The name of the database instance after the modification.
	//
	// example:
	//
	// pgsql
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The port of the database after the modification.
	//
	// example:
	//
	// 5433
	DatabasePort *string `json:"DatabasePort,omitempty" xml:"DatabasePort,omitempty"`
	// The private network address of the database after the modification. IPv4 addresses and domain name formats are supported.
	//
	// example:
	//
	// 192.168.XX.XX
	DatabasePrivateAddress *string `json:"DatabasePrivateAddress,omitempty" xml:"DatabasePrivateAddress,omitempty"`
	// The public address of the database after the modification. Both IPv4 addresses and domain names are supported.
	//
	// example:
	//
	// www.example.com
	DatabasePublicAddress *string `json:"DatabasePublicAddress,omitempty" xml:"DatabasePublicAddress,omitempty"`
	// The ID of the Bastionhost instance to which the database to be modified belongs.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-72137xe5n01
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the network domain to which the database belongs after the modification.
	//
	// > You can call the [ListNetworkDomains ](https://help.aliyun.com/document_detail/2758827.html) operation to obtain this parameter.
	//
	// example:
	//
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance to which the database to be modified belongs.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the database to be modified.
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

func (s *ModifyDatabaseRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *ModifyDatabaseRequest) SetProjectId(v int64) *ModifyDatabaseRequest {
	s.ProjectId = &v
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
