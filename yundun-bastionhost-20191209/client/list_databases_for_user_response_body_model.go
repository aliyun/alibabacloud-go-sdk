// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*ListDatabasesForUserResponseBodyDatabases) *ListDatabasesForUserResponseBody
	GetDatabases() []*ListDatabasesForUserResponseBodyDatabases
	SetRequestId(v string) *ListDatabasesForUserResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDatabasesForUserResponseBody
	GetTotalCount() *int64
}

type ListDatabasesForUserResponseBody struct {
	// The list of databases.
	Databases []*ListDatabasesForUserResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A8A665B9-8550-4942-9DEE-73198051856B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of databases.
	//
	// example:
	//
	// 15
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatabasesForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesForUserResponseBody) GetDatabases() []*ListDatabasesForUserResponseBodyDatabases {
	return s.Databases
}

func (s *ListDatabasesForUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabasesForUserResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatabasesForUserResponseBody) SetDatabases(v []*ListDatabasesForUserResponseBodyDatabases) *ListDatabasesForUserResponseBody {
	s.Databases = v
	return s
}

func (s *ListDatabasesForUserResponseBody) SetRequestId(v string) *ListDatabasesForUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesForUserResponseBody) SetTotalCount(v int64) *ListDatabasesForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabasesForUserResponseBody) Validate() error {
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

type ListDatabasesForUserResponseBodyDatabases struct {
	// The address type of the database. Valid values:
	//
	// - **Public**: public endpoint
	//
	// - **Private**: private endpoint in the private network
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
	// The instance ID of the database.
	//
	// example:
	//
	// 36
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The name of the database instance.
	//
	// example:
	//
	// MySQL56
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The port of the database.
	//
	// example:
	//
	// 3306
	DatabasePort *int64 `json:"DatabasePort,omitempty" xml:"DatabasePort,omitempty"`
	// The private endpoint of the database in the private network, which can be a domain name or an IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	DatabasePrivateAddress *string `json:"DatabasePrivateAddress,omitempty" xml:"DatabasePrivateAddress,omitempty"`
	// The public address of the database, which can be a domain name or an IP address.
	//
	// example:
	//
	// www.example.com
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
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The ID of the project to which the database belongs.
	//
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The source of the database. Valid values:
	//
	// - **Local**: local database instance
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
	// i-wz9fv2hwux78x9******
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
}

func (s ListDatabasesForUserResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesForUserResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetComment() *string {
	return s.Comment
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetDatabasePort() *int64 {
	return s.DatabasePort
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetDatabasePrivateAddress() *string {
	return s.DatabasePrivateAddress
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetDatabasePublicAddress() *string {
	return s.DatabasePublicAddress
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetSource() *string {
	return s.Source
}

func (s *ListDatabasesForUserResponseBodyDatabases) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetActiveAddressType(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.ActiveAddressType = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetComment(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.Comment = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetDatabaseId(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetDatabaseName(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.DatabaseName = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetDatabasePort(v int64) *ListDatabasesForUserResponseBodyDatabases {
	s.DatabasePort = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetDatabasePrivateAddress(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.DatabasePrivateAddress = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetDatabasePublicAddress(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.DatabasePublicAddress = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetDatabaseType(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.DatabaseType = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetNetworkDomainId(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.NetworkDomainId = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetProjectId(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.ProjectId = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetProjectName(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.ProjectName = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetSource(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.Source = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) SetSourceInstanceId(v string) *ListDatabasesForUserResponseBodyDatabases {
	s.SourceInstanceId = &v
	return s
}

func (s *ListDatabasesForUserResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}
