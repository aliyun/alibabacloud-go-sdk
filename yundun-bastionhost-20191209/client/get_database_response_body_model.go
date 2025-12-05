// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v *GetDatabaseResponseBodyDatabase) *GetDatabaseResponseBody
	GetDatabase() *GetDatabaseResponseBodyDatabase
	SetRequestId(v string) *GetDatabaseResponseBody
	GetRequestId() *string
}

type GetDatabaseResponseBody struct {
	// The returned detailed information about the database.
	Database *GetDatabaseResponseBodyDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4D72B883-9D15-5B05-B987-DFD10EB1FFB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBody) GetDatabase() *GetDatabaseResponseBodyDatabase {
	return s.Database
}

func (s *GetDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatabaseResponseBody) SetDatabase(v *GetDatabaseResponseBodyDatabase) *GetDatabaseResponseBody {
	s.Database = v
	return s
}

func (s *GetDatabaseResponseBody) SetRequestId(v string) *GetDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseResponseBody) Validate() error {
	if s.Database != nil {
		if err := s.Database.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDatabaseResponseBodyDatabase struct {
	// The address type of the database. Valid values:
	//
	// 	- Public
	//
	// 	- Private
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
	// 22
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The port of the database.
	//
	// example:
	//
	// 3306
	DatabasePort *int64 `json:"DatabasePort,omitempty" xml:"DatabasePort,omitempty"`
	// The internal endpoint of the database.
	//
	// example:
	//
	// rm-bp1zq******
	DatabasePrivateAddress *string `json:"DatabasePrivateAddress,omitempty" xml:"DatabasePrivateAddress,omitempty"`
	// The public endpoint of the database.
	//
	// example:
	//
	// rm-uf65******
	DatabasePublicAddress *string `json:"DatabasePublicAddress,omitempty" xml:"DatabasePublicAddress,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **mysql**
	//
	// 	- **sqlserver**
	//
	// 	- **postgresql**
	//
	// 	- **oracle**
	//
	// example:
	//
	// MySQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The ID of the network domain to which the database belongs.
	//
	// example:
	//
	// 45
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
	// > If **Source*	- is set to **Local**, this parameter is empty.
	//
	// example:
	//
	// i-wz9527ob0e0nftcsffke
	SourceInstanceId *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	// The region ID of the ApsaraDB RDS instance or PolarDB cluster.
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

func (s GetDatabaseResponseBodyDatabase) String() string {
	return dara.Prettify(s)
}

func (s GetDatabaseResponseBodyDatabase) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBodyDatabase) GetActiveAddressType() *string {
	return s.ActiveAddressType
}

func (s *GetDatabaseResponseBodyDatabase) GetComment() *string {
	return s.Comment
}

func (s *GetDatabaseResponseBodyDatabase) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *GetDatabaseResponseBodyDatabase) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetDatabaseResponseBodyDatabase) GetDatabasePort() *int64 {
	return s.DatabasePort
}

func (s *GetDatabaseResponseBodyDatabase) GetDatabasePrivateAddress() *string {
	return s.DatabasePrivateAddress
}

func (s *GetDatabaseResponseBodyDatabase) GetDatabasePublicAddress() *string {
	return s.DatabasePublicAddress
}

func (s *GetDatabaseResponseBodyDatabase) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *GetDatabaseResponseBodyDatabase) GetNetworkDomainId() *string {
	return s.NetworkDomainId
}

func (s *GetDatabaseResponseBodyDatabase) GetSource() *string {
	return s.Source
}

func (s *GetDatabaseResponseBodyDatabase) GetSourceInstanceId() *string {
	return s.SourceInstanceId
}

func (s *GetDatabaseResponseBodyDatabase) GetSourceInstanceRegionId() *string {
	return s.SourceInstanceRegionId
}

func (s *GetDatabaseResponseBodyDatabase) GetSourceInstanceState() *string {
	return s.SourceInstanceState
}

func (s *GetDatabaseResponseBodyDatabase) SetActiveAddressType(v string) *GetDatabaseResponseBodyDatabase {
	s.ActiveAddressType = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetComment(v string) *GetDatabaseResponseBodyDatabase {
	s.Comment = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDatabaseId(v string) *GetDatabaseResponseBodyDatabase {
	s.DatabaseId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDatabaseName(v string) *GetDatabaseResponseBodyDatabase {
	s.DatabaseName = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDatabasePort(v int64) *GetDatabaseResponseBodyDatabase {
	s.DatabasePort = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDatabasePrivateAddress(v string) *GetDatabaseResponseBodyDatabase {
	s.DatabasePrivateAddress = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDatabasePublicAddress(v string) *GetDatabaseResponseBodyDatabase {
	s.DatabasePublicAddress = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetDatabaseType(v string) *GetDatabaseResponseBodyDatabase {
	s.DatabaseType = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetNetworkDomainId(v string) *GetDatabaseResponseBodyDatabase {
	s.NetworkDomainId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSource(v string) *GetDatabaseResponseBodyDatabase {
	s.Source = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSourceInstanceId(v string) *GetDatabaseResponseBodyDatabase {
	s.SourceInstanceId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSourceInstanceRegionId(v string) *GetDatabaseResponseBodyDatabase {
	s.SourceInstanceRegionId = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) SetSourceInstanceState(v string) *GetDatabaseResponseBodyDatabase {
	s.SourceInstanceState = &v
	return s
}

func (s *GetDatabaseResponseBodyDatabase) Validate() error {
	return dara.Validate(s)
}
