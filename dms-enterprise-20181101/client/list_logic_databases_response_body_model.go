// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicDatabasesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListLogicDatabasesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListLogicDatabasesResponseBody
	GetErrorMessage() *string
	SetLogicDatabaseList(v *ListLogicDatabasesResponseBodyLogicDatabaseList) *ListLogicDatabasesResponseBody
	GetLogicDatabaseList() *ListLogicDatabasesResponseBodyLogicDatabaseList
	SetRequestId(v string) *ListLogicDatabasesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLogicDatabasesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListLogicDatabasesResponseBody
	GetTotalCount() *int64
}

type ListLogicDatabasesResponseBody struct {
	// The error code that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The details of logical databases.
	LogicDatabaseList *ListLogicDatabasesResponseBodyLogicDatabaseList `json:"LogicDatabaseList,omitempty" xml:"LogicDatabaseList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 8068AF82-8A1A-592C-AC2E-6B75338BAB87
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of logical databases.
	//
	// example:
	//
	// 7
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLogicDatabasesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLogicDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListLogicDatabasesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListLogicDatabasesResponseBody) GetLogicDatabaseList() *ListLogicDatabasesResponseBodyLogicDatabaseList {
	return s.LogicDatabaseList
}

func (s *ListLogicDatabasesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLogicDatabasesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLogicDatabasesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLogicDatabasesResponseBody) SetErrorCode(v string) *ListLogicDatabasesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetErrorMessage(v string) *ListLogicDatabasesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetLogicDatabaseList(v *ListLogicDatabasesResponseBodyLogicDatabaseList) *ListLogicDatabasesResponseBody {
	s.LogicDatabaseList = v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetRequestId(v string) *ListLogicDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetSuccess(v bool) *ListLogicDatabasesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLogicDatabasesResponseBody) SetTotalCount(v int64) *ListLogicDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLogicDatabasesResponseBody) Validate() error {
	if s.LogicDatabaseList != nil {
		if err := s.LogicDatabaseList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLogicDatabasesResponseBodyLogicDatabaseList struct {
	LogicDatabase []*ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase `json:"LogicDatabase,omitempty" xml:"LogicDatabase,omitempty" type:"Repeated"`
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseList) String() string {
	return dara.Prettify(s)
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseList) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseList) GetLogicDatabase() []*ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	return s.LogicDatabase
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseList) SetLogicDatabase(v []*ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) *ListLogicDatabasesResponseBodyLogicDatabaseList {
	s.LogicDatabase = v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseList) Validate() error {
	if s.LogicDatabase != nil {
		for _, item := range s.LogicDatabase {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase struct {
	// The alias of the logical database.
	//
	// example:
	//
	// logic_db_alias
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The ID of the logical database.
	//
	// example:
	//
	// 1***
	DatabaseId *string `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	// Logical database sub-ID list.
	DatabaseIds *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty" type:"Struct"`
	// The type of the logical database. For more information about the valid values of this parameter, see [DbType parameter](https://www.alibabacloud.com/help/en/data-management-service/latest/dbtype-parameter).
	//
	// example:
	//
	// polardb
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the logical database belongs. Valid values:
	//
	// - **product**: production environment
	//
	// - **dev**: development environment
	//
	// - **pre**: staging environment
	//
	// - **test**: test environment
	//
	// - **sit**: system integration testing (SIT) environment
	//
	// - **uat**: user acceptance testing (UAT) environment
	//
	// - **pet**: stress testing environment
	//
	// - **stag**: STAG environment
	//
	// example:
	//
	// test
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// Indicates whether the database is a logical database. The return value is true.
	//
	// example:
	//
	// true
	Logic *bool `json:"Logic,omitempty" xml:"Logic,omitempty"`
	// The IDs of the owners of the logical database.
	OwnerIdList *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList `json:"OwnerIdList,omitempty" xml:"OwnerIdList,omitempty" type:"Struct"`
	// The names of the owners of the logical database.
	OwnerNameList *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList `json:"OwnerNameList,omitempty" xml:"OwnerNameList,omitempty" type:"Struct"`
	// The name of the logical database.
	//
	// example:
	//
	// logic_db
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name that is used to search for the logical database.
	//
	// > We recommend that you do not use this parameter for business development. The format of the parameter value may be modified in later versions.
	//
	// example:
	//
	// logic_db[logic_db_alias]
	SearchName *string `json:"SearchName,omitempty" xml:"SearchName,omitempty"`
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) String() string {
	return dara.Prettify(s)
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetAlias() *string {
	return s.Alias
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetDatabaseId() *string {
	return s.DatabaseId
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetDatabaseIds() *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds {
	return s.DatabaseIds
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetDbType() *string {
	return s.DbType
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetEnvType() *string {
	return s.EnvType
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetLogic() *bool {
	return s.Logic
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetOwnerIdList() *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList {
	return s.OwnerIdList
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetOwnerNameList() *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList {
	return s.OwnerNameList
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetSchemaName() *string {
	return s.SchemaName
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) GetSearchName() *string {
	return s.SearchName
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetAlias(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.Alias = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetDatabaseId(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.DatabaseId = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetDatabaseIds(v *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.DatabaseIds = v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetDbType(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.DbType = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetEnvType(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.EnvType = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetLogic(v bool) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.Logic = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetOwnerIdList(v *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.OwnerIdList = v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetOwnerNameList(v *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.OwnerNameList = v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetSchemaName(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.SchemaName = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) SetSearchName(v string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase {
	s.SearchName = &v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabase) Validate() error {
	if s.DatabaseIds != nil {
		if err := s.DatabaseIds.Validate(); err != nil {
			return err
		}
	}
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

type ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds struct {
	DatabaseIds []*int64 `json:"DatabaseIds,omitempty" xml:"DatabaseIds,omitempty" type:"Repeated"`
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds) String() string {
	return dara.Prettify(s)
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds) GetDatabaseIds() []*int64 {
	return s.DatabaseIds
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds) SetDatabaseIds(v []*int64) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds {
	s.DatabaseIds = v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseDatabaseIds) Validate() error {
	return dara.Validate(s)
}

type ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList struct {
	OwnerIds []*string `json:"OwnerIds,omitempty" xml:"OwnerIds,omitempty" type:"Repeated"`
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) String() string {
	return dara.Prettify(s)
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) GetOwnerIds() []*string {
	return s.OwnerIds
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) SetOwnerIds(v []*string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList {
	s.OwnerIds = v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerIdList) Validate() error {
	return dara.Validate(s)
}

type ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList struct {
	OwnerNames []*string `json:"OwnerNames,omitempty" xml:"OwnerNames,omitempty" type:"Repeated"`
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) String() string {
	return dara.Prettify(s)
}

func (s ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) GetOwnerNames() []*string {
	return s.OwnerNames
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) SetOwnerNames(v []*string) *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList {
	s.OwnerNames = v
	return s
}

func (s *ListLogicDatabasesResponseBodyLogicDatabaseListLogicDatabaseOwnerNameList) Validate() error {
	return dara.Validate(s)
}
