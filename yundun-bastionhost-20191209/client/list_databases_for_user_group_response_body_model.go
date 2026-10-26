// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesForUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*ListDatabasesForUserGroupResponseBodyDatabases) *ListDatabasesForUserGroupResponseBody
	GetDatabases() []*ListDatabasesForUserGroupResponseBodyDatabases
	SetRequestId(v string) *ListDatabasesForUserGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDatabasesForUserGroupResponseBody
	GetTotalCount() *int64
}

type ListDatabasesForUserGroupResponseBody struct {
	// The queried databases.
	Databases []*ListDatabasesForUserGroupResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried databases.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatabasesForUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesForUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesForUserGroupResponseBody) GetDatabases() []*ListDatabasesForUserGroupResponseBodyDatabases {
	return s.Databases
}

func (s *ListDatabasesForUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabasesForUserGroupResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatabasesForUserGroupResponseBody) SetDatabases(v []*ListDatabasesForUserGroupResponseBodyDatabases) *ListDatabasesForUserGroupResponseBody {
	s.Databases = v
	return s
}

func (s *ListDatabasesForUserGroupResponseBody) SetRequestId(v string) *ListDatabasesForUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBody) SetTotalCount(v int64) *ListDatabasesForUserGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBody) Validate() error {
	if s.Databases != nil {
		for _, item := range s.Databases {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDatabasesForUserGroupResponseBodyDatabases struct {
	// The address type of the database. Valid values:
	//
	// - **Public**: public network address
	//
	// - **Private**: private network address
	//
	// example:
	//
	// Private
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The description of the database instance.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The total number of queried database accounts.
	//
	// example:
	//
	// 2
	DatabaseAccountCount *int64 `json:"DatabaseAccountCount,omitempty" xml:"DatabaseAccountCount,omitempty"`
	// The instance ID of the database to which the database account belongs.
	//
	// example:
	//
	// 2
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The name of the database instance.
	//
	// example:
	//
	// MySQL0
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The port of the database.
	//
	// example:
	//
	// 3306
	DatabasePort *int64 `json:"DatabasePort,omitempty" xml:"DatabasePort,omitempty"`
	// The private network address of the database, which can be a domain name or an IP address.
	//
	// example:
	//
	// rm-bp1******
	DatabasePrivateAddress *string `json:"DatabasePrivateAddress,omitempty" xml:"DatabasePrivateAddress,omitempty"`
	// The public address of the database, which can be a domain name or an IP address.
	//
	// example:
	//
	// rm-uf65******
	DatabasePublicAddress *string `json:"DatabasePublicAddress,omitempty" xml:"DatabasePublicAddress,omitempty"`
	// The type of the database. Valid values:
	//
	// - **MySQL**
	//
	// - **Oracle**
	//
	// - **PostgreSQL**
	//
	// - **SQLServer**
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The ID of the network domain to which the database belongs.
	//
	// example:
	//
	// 5
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The source of the database. Valid values:
	//
	// - **Local**: self-managed database instance
	//
	// - **Rds**: ApsaraDB RDS instance
	//
	// - **PolarDB**: PolarDB instance
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The instance ID of the ApsaraDB RDS or PolarDB instance that corresponds to the database.
	//
	// > If **Source*	- is set to **Local**, this parameter is empty.
	//
	// example:
	//
	// i-wz9c7mjxywmdmqk7q6e4
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s ListDatabasesForUserGroupResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesForUserGroupResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetComment() *string {
	return s.Comment
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetDatabaseAccountCount() *int64 {
	return s.DatabaseAccountCount
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetDatabasePort() *int64 {
	return s.DatabasePort
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetDatabasePrivateAddress() *string {
	return s.DatabasePrivateAddress
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetDatabasePublicAddress() *string {
	return s.DatabasePublicAddress
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetSource() *string {
	return s.Source
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetActiveAddressType(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.ActiveAddressType = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetComment(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.Comment = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetDatabaseAccountCount(v int64) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.DatabaseAccountCount = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetDatabaseId(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetDatabaseName(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.DatabaseName = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetDatabasePort(v int64) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.DatabasePort = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetDatabasePrivateAddress(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.DatabasePrivateAddress = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetDatabasePublicAddress(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.DatabasePublicAddress = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetDatabaseType(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.DatabaseType = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetNetworkDomainId(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.NetworkDomainId = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetProjectId(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.ProjectId = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetProjectName(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.ProjectName = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetSource(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.Source = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) SetSourceInstanceId(v string) *ListDatabasesForUserGroupResponseBodyDatabases {
	s.SourceInstanceId = &v
	return s
}

func (s *ListDatabasesForUserGroupResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}
