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
	// The databases returned.
	Databases []*ListDatabasesForUserResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A8A665B9-8550-4942-9DEE-73198051856B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of databases returned.
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
	return dara.Validate(s)
}

type ListDatabasesForUserResponseBodyDatabases struct {
	// The address type of the database. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Private
	ActiveAddressType *string `json:"ActiveAddressType,omitempty" xml:"ActiveAddressType,omitempty"`
	// The remarks of the database.
	//
	// example:
	//
	// comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The database ID.
	//
	// example:
	//
	// 36
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The database name.
	//
	// example:
	//
	// MySQL56
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The database port.
	//
	// example:
	//
	// 3306
	DatabasePort *int64 `json:"DatabasePort,omitempty" xml:"DatabasePort,omitempty"`
	// The internal endpoint of the database. The value is a domain name or an IP address.
	//
	// example:
	//
	// rm-wz97******
	DatabasePrivateAddress *string `json:"DatabasePrivateAddress,omitempty" xml:"DatabasePrivateAddress,omitempty"`
	// The public endpoint of the database. The value is a domain name or an IP address.
	//
	// example:
	//
	// rm-uf65******
	DatabasePublicAddress *string `json:"DatabasePublicAddress,omitempty" xml:"DatabasePublicAddress,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **Oracle**
	//
	// 	- **PostgreSQL**
	//
	// 	- **SQLServer**
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The ID of the network domain where the database resides.
	//
	// example:
	//
	// 2
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The database type. Valid values:
	//
	// 	- **Local**: on-premises database.
	//
	// 	- **Rds**: ApsaraDB RDS instance.
	//
	// 	- **PolarDB**: PolarDB cluster.
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ApsaraDB RDS instance or PolarDB cluster.
	//
	// >  No value is returned for this parameter if **Source*	- is set to **Local**.
	//
	// example:
	//
	// i-wz9fv2hwux78x9h1pmje
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
