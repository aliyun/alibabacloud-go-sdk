// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SearchDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SearchDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SearchDatabaseResponseBody
	GetRequestId() *string
	SetSearchDatabaseList(v *SearchDatabaseResponseBodySearchDatabaseList) *SearchDatabaseResponseBody
	GetSearchDatabaseList() *SearchDatabaseResponseBodySearchDatabaseList
	SetSuccess(v bool) *SearchDatabaseResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *SearchDatabaseResponseBody
	GetTotalCount() *int64
}

type SearchDatabaseResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E0D21075-CD3E-4D98-8264-FD8AD04A63B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the databases.
	SearchDatabaseList *SearchDatabaseResponseBodySearchDatabaseList `json:"SearchDatabaseList,omitempty" xml:"SearchDatabaseList,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SearchDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SearchDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchDatabaseResponseBody) GetSearchDatabaseList() *SearchDatabaseResponseBodySearchDatabaseList {
	return s.SearchDatabaseList
}

func (s *SearchDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchDatabaseResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchDatabaseResponseBody) SetErrorCode(v string) *SearchDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SearchDatabaseResponseBody) SetErrorMessage(v string) *SearchDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SearchDatabaseResponseBody) SetRequestId(v string) *SearchDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchDatabaseResponseBody) SetSearchDatabaseList(v *SearchDatabaseResponseBodySearchDatabaseList) *SearchDatabaseResponseBody {
	s.SearchDatabaseList = v
	return s
}

func (s *SearchDatabaseResponseBody) SetSuccess(v bool) *SearchDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *SearchDatabaseResponseBody) SetTotalCount(v int64) *SearchDatabaseResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchDatabaseResponseBody) Validate() error {
	if s.SearchDatabaseList != nil {
		if err := s.SearchDatabaseList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchDatabaseResponseBodySearchDatabaseList struct {
	SearchDatabase []*SearchDatabaseResponseBodySearchDatabaseListSearchDatabase `json:"SearchDatabase,omitempty" xml:"SearchDatabase,omitempty" type:"Repeated"`
}

func (s SearchDatabaseResponseBodySearchDatabaseList) String() string {
	return dara.Prettify(s)
}

func (s SearchDatabaseResponseBodySearchDatabaseList) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBodySearchDatabaseList) GetSearchDatabase() []*SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	return s.SearchDatabase
}

func (s *SearchDatabaseResponseBodySearchDatabaseList) SetSearchDatabase(v []*SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) *SearchDatabaseResponseBodySearchDatabaseList {
	s.SearchDatabase = v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseList) Validate() error {
	if s.SearchDatabase != nil {
		for _, item := range s.SearchDatabase {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchDatabaseResponseBodySearchDatabaseListSearchDatabase struct {
	// The alias of the database.
	//
	// example:
	//
	// test_rds
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The name of the catalog to which the database belongs.
	//
	// > If the type of the database engine is PostgreSQL, the name of the database is displayed.
	//
	// example:
	//
	// dmstest
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The ID of the database.
	//
	// example:
	//
	// 2528****
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// The name of the data link for cross-database queries.
	//
	// example:
	//
	// datalink_name
	DatalinkName *string `json:"DatalinkName,omitempty" xml:"DatalinkName,omitempty"`
	// The type of the database engine.
	//
	// example:
	//
	// mysql
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The ID of the user who assumes the database administrator (DBA) role.
	//
	// example:
	//
	// 10****
	DbaId *string `json:"DbaId,omitempty" xml:"DbaId,omitempty"`
	// The encoding method of the database.
	//
	// example:
	//
	// utf8
	Encoding *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	// The environment type of the database. For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// test
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The endpoint of the instance in which the database resides.
	//
	// example:
	//
	// rm-xxxx.mysql.rds.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// Indicates whether the database is a logical database. Valid values:
	//
	// 	- **true**: The database is a logical database.
	//
	// 	- **false**: The database is not a logical database.
	//
	// example:
	//
	// false
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The IDs of the owners of the databases.
	OwnerIdList *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The nicknames of the database owners.
	OwnerNameList *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The port of the instance in which the database resides.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used to search for the database.
	//
	// example:
	//
	// test@xxx.xxx.xxx.xxx:3306
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
	// The system ID (SID) of the instance in which the database resides.
	//
	// example:
	//
	// testSid
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) String() string {
	return dara.Prettify(s)
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetAlias() *string {
	return s.Alias
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetCatalogName() *string {
	return s.CatalogName
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetDatalinkName() *string {
	return s.DatalinkName
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetDbType() *string {
	return s.DbType
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetDbaId() *string {
	return s.DbaId
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetEncoding() *string {
	return s.Encoding
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetEnvType() *string {
	return s.EnvType
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetHost() *string {
	return s.Host
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetLogic() *bool {
	return s.Logic
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetOwnerIdList() *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList {
	return s.OwnerIdList
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetOwnerNameList() *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList {
	return s.OwnerNameList
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetPort() *int32 {
	return s.Port
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetSchemaName() *string {
	return s.SchemaName
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) GetSid() *string {
	return s.Sid
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetAlias(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Alias = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetCatalogName(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.CatalogName = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetDatabaseId(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.DatabaseId = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetDatalinkName(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.DatalinkName = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetDbType(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.DbType = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetDbaId(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.DbaId = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetEncoding(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Encoding = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetEnvType(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.EnvType = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetHost(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Host = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetLogic(v bool) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Logic = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetOwnerIdList(v *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.OwnerIdList = v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetOwnerNameList(v *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.OwnerNameList = v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetPort(v int32) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Port = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetSchemaName(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.SchemaName = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetSearchName(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.SearchName = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) SetSid(v string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase {
	s.Sid = &v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabase) Validate() error {
	if s.OwnerIdList != nil {
		if err := s.OwnerIdList.Validate(); err != nil {
			return err
		}
	}
	if s.OwnerNameList != nil {
		if err := s.OwnerNameList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) SetOwnerIds(v []*string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) SetOwnerNames(v []*string) *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *SearchDatabaseResponseBodySearchDatabaseListSearchDatabaseOwnerNameList) Validate() error {
	return dara.Validate(s)
}
