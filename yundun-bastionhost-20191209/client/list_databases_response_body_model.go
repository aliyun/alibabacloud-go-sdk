// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabases(v []*ListDatabasesResponseBodyDatabases) *ListDatabasesResponseBody
	GetDatabases() []*ListDatabasesResponseBodyDatabases
	SetRequestId(v string) *ListDatabasesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDatabasesResponseBody
	GetTotalCount() *int64
}

type ListDatabasesResponseBody struct {
	// The databases returned.
	Databases []*ListDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of databases returned.
	//
	// example:
	//
	// 4
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBody) GetDatabases() []*ListDatabasesResponseBodyDatabases {
	return s.Databases
}

func (s *ListDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabasesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatabasesResponseBody) SetDatabases(v []*ListDatabasesResponseBodyDatabases) *ListDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *ListDatabasesResponseBody) SetRequestId(v string) *ListDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesResponseBody) SetTotalCount(v int64) *ListDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatabasesResponseBody) Validate() error {
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

type ListDatabasesResponseBodyDatabases struct {
	// The address type of the database. Valid values:
	//
	// 	- **Public**
	//
	// 	- **Private**
	//
	// example:
	//
	// Public
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
	// 9
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The name of the database.
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
	DatabasePort *int32 `json:"DatabasePort,omitempty" xml:"DatabasePort,omitempty"`
	// The internal address of the database. The value is a domain name or an IP address.
	//
	// example:
	//
	// rm-wz973w7******
	DatabasePrivateAddress *string `json:"DatabasePrivateAddress,omitempty" xml:"DatabasePrivateAddress,omitempty"`
	// The public address of the database. The value is a domain name or an IP address.
	//
	// example:
	//
	// rm-uf65n2******
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
	// 8
	NetworkDomainId *string `json:"NetworkDomainId,omitempty" xml:"NetworkDomainId,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **Local**: on-premises database.
	//
	// 	- **Rds**: ApsaraDB for RDS instance.
	//
	// 	- **PolarDB**: PolarDB cluster
	//
	// example:
	//
	// Local
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The ID of the ApsaraDB for RDS instance or PolarDB cluster.
	//
	// > No value is returned for this parameter if **Source*	- is set to **Local**.
	//
	// example:
	//
	// i-wz9ejupczf41******
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The region ID of the ApsaraDB for RDS instance or PolarDB cluster.
	//
	// example:
	//
	// cn-hangzhou
	SourceInstanceRegionId *string `json:"SourceInstanceRegionId,omitempty" xml:"SourceInstanceRegionId,omitempty"`
	// The status of the database. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **Release**
	//
	// example:
	//
	// Normal
	SourceInstanceState *string `json:"SourceInstanceState,omitempty" xml:"SourceInstanceState,omitempty"`
}

func (s ListDatabasesResponseBodyDatabases) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBodyDatabases) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *ListDatabasesResponseBodyDatabases) GetComment() *string {
	return s.Comment
}

func (s *ListDatabasesResponseBodyDatabases) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListDatabasesResponseBodyDatabases) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *ListDatabasesResponseBodyDatabases) GetDatabasePort() *int32 {
	return s.DatabasePort
}

func (s *ListDatabasesResponseBodyDatabases) GetDatabasePrivateAddress() *string {
	return s.DatabasePrivateAddress
}

func (s *ListDatabasesResponseBodyDatabases) GetDatabasePublicAddress() *string {
	return s.DatabasePublicAddress
}

func (s *ListDatabasesResponseBodyDatabases) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDatabasesResponseBodyDatabases) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *ListDatabasesResponseBodyDatabases) GetSource() *string {
	return s.Source
}

func (s *ListDatabasesResponseBodyDatabases) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *ListDatabasesResponseBodyDatabases) GetSourceInstanceRegionId() *string {
	return s.SourceInstanceRegionId
}

func (s *ListDatabasesResponseBodyDatabases) GetSourceInstanceState() *string {
	return s.SourceInstanceState
}

func (s *ListDatabasesResponseBodyDatabases) SetActiveAddressType(v string) *ListDatabasesResponseBodyDatabases {
	s.ActiveAddressType = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetComment(v string) *ListDatabasesResponseBodyDatabases {
	s.Comment = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetDatabaseId(v string) *ListDatabasesResponseBodyDatabases {
	s.DatabaseId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetDatabaseName(v string) *ListDatabasesResponseBodyDatabases {
	s.DatabaseName = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetDatabasePort(v int32) *ListDatabasesResponseBodyDatabases {
	s.DatabasePort = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetDatabasePrivateAddress(v string) *ListDatabasesResponseBodyDatabases {
	s.DatabasePrivateAddress = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetDatabasePublicAddress(v string) *ListDatabasesResponseBodyDatabases {
	s.DatabasePublicAddress = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetDatabaseType(v string) *ListDatabasesResponseBodyDatabases {
	s.DatabaseType = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetNetworkDomainId(v string) *ListDatabasesResponseBodyDatabases {
	s.NetworkDomainId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetSource(v string) *ListDatabasesResponseBodyDatabases {
	s.Source = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetSourceInstanceId(v string) *ListDatabasesResponseBodyDatabases {
	s.SourceInstanceId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetSourceInstanceRegionId(v string) *ListDatabasesResponseBodyDatabases {
	s.SourceInstanceRegionId = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) SetSourceInstanceState(v string) *ListDatabasesResponseBodyDatabases {
	s.SourceInstanceState = &v
	return s
}

func (s *ListDatabasesResponseBodyDatabases) Validate() error {
	return dara.Validate(s)
}
