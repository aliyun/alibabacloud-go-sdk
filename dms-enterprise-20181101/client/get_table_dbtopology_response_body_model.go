// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableDBTopologyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBTopology(v *GetTableDBTopologyResponseBodyDBTopology) *GetTableDBTopologyResponseBody
	GetDBTopology() *GetTableDBTopologyResponseBodyDBTopology
	SetErrorCode(v string) *GetTableDBTopologyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetTableDBTopologyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetTableDBTopologyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableDBTopologyResponseBody
	GetSuccess() *bool
}

type GetTableDBTopologyResponseBody struct {
	// The topology of the data table.
	DBTopology *GetTableDBTopologyResponseBodyDBTopology `json:"DBTopology,omitempty" xml:"DBTopology,omitempty" type:"Struct"`
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
	// 853F7FD4-D922-4EFB-931C-D253EF159E06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableDBTopologyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableDBTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBody) GetDBTopology() *GetTableDBTopologyResponseBodyDBTopology {
	return s.DBTopology
}

func (s *GetTableDBTopologyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTableDBTopologyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetTableDBTopologyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableDBTopologyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableDBTopologyResponseBody) SetDBTopology(v *GetTableDBTopologyResponseBodyDBTopology) *GetTableDBTopologyResponseBody {
	s.DBTopology = v
	return s
}

func (s *GetTableDBTopologyResponseBody) SetErrorCode(v string) *GetTableDBTopologyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTableDBTopologyResponseBody) SetErrorMessage(v string) *GetTableDBTopologyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetTableDBTopologyResponseBody) SetRequestId(v string) *GetTableDBTopologyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableDBTopologyResponseBody) SetSuccess(v bool) *GetTableDBTopologyResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableDBTopologyResponseBody) Validate() error {
	if s.DBTopology != nil {
		if err := s.DBTopology.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTableDBTopologyResponseBodyDBTopology struct {
	// The data sources.
	DataSourceList []*GetTableDBTopologyResponseBodyDBTopologyDataSourceList `json:"DataSourceList,omitempty" xml:"DataSourceList,omitempty" type:"Repeated"`
	// The GUID of the table in DMS.
	//
	// example:
	//
	// IDB_L_9032.db-test.yuyang_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The name of the table.
	//
	// >
	//
	// 	- If a logical table is queried, the name of the logical table is returned.
	//
	// 	- If a physical table is queried, the name of the physical table is returned.
	//
	// example:
	//
	// yuyang_test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableDBTopologyResponseBodyDBTopology) String() string {
	return dara.Prettify(s)
}

func (s GetTableDBTopologyResponseBodyDBTopology) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBodyDBTopology) GetDataSourceList() []*GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	return s.DataSourceList
}

func (s *GetTableDBTopologyResponseBodyDBTopology) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetTableDBTopologyResponseBodyDBTopology) GetTableName() *string {
	return s.TableName
}

func (s *GetTableDBTopologyResponseBodyDBTopology) SetDataSourceList(v []*GetTableDBTopologyResponseBodyDBTopologyDataSourceList) *GetTableDBTopologyResponseBodyDBTopology {
	s.DataSourceList = v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopology) SetTableGuid(v string) *GetTableDBTopologyResponseBodyDBTopology {
	s.TableGuid = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopology) SetTableName(v string) *GetTableDBTopologyResponseBodyDBTopology {
	s.TableName = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopology) Validate() error {
	if s.DataSourceList != nil {
		for _, item := range s.DataSourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableDBTopologyResponseBodyDBTopologyDataSourceList struct {
	// The physical databases.
	DatabaseList []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList `json:"DatabaseList,omitempty" xml:"DatabaseList,omitempty" type:"Repeated"`
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The endpoint of the data source.
	//
	// example:
	//
	// xxx.mysql.polardb.rds.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The port that is used to connect to the data source.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The system ID (SID) of the data source.
	//
	// example:
	//
	// def
	Sid *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceList) String() string {
	return dara.Prettify(s)
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceList) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) GetDatabaseList() []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	return s.DatabaseList
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) GetDbType() *string {
	return s.DbType
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) GetHost() *string {
	return s.Host
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) GetPort() *int32 {
	return s.Port
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) GetSid() *string {
	return s.Sid
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetDatabaseList(v []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.DatabaseList = v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetDbType(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.DbType = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetHost(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.Host = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetPort(v int32) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.Port = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) SetSid(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceList {
	s.Sid = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceList) Validate() error {
	if s.DatabaseList != nil {
		for _, item := range s.DatabaseList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList struct {
	// The ID of the database.
	//
	// example:
	//
	// 489347
	DbId *string `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// db-test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The type of the database. For more information about the valid values of this parameter, see [DbType parameter](https://help.aliyun.com/document_detail/198106.html).
	//
	// example:
	//
	// MySQL
	DbType *string `json:"DbType,omitempty" xml:"DbType,omitempty"`
	// The type of the environment to which the database belongs. Valid values:
	//
	// 	- **product**: production environment
	//
	// 	- **dev**: development environment
	//
	// 	- **pre**: pre-release environment
	//
	// 	- **test**: test environment
	//
	// 	- **sit**: system integration testing (SIT) environment
	//
	// 	- **uat**: user acceptance testing (UAT) environment
	//
	// 	- **pet**: stress testing environment
	//
	// 	- **stag**: staging environment
	//
	// > For more information, see [Change the environment type of an instance](https://help.aliyun.com/document_detail/163309.html).
	//
	// example:
	//
	// pre
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The physical tables.
	TableList []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Repeated"`
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) String() string {
	return dara.Prettify(s)
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) GetDbId() *string {
	return s.DbId
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) GetDbName() *string {
	return s.DbName
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) GetDbType() *string {
	return s.DbType
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) GetEnvType() *string {
	return s.EnvType
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) GetTableList() []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList {
	return s.TableList
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetDbId(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.DbId = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetDbName(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.DbName = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetDbType(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.DbType = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetEnvType(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.EnvType = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) SetTableList(v []*GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList {
	s.TableList = v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseList) Validate() error {
	if s.TableList != nil {
		for _, item := range s.TableList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList struct {
	// The ID of the table.
	//
	// example:
	//
	// NORMAL
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The name of the physical table.
	//
	// example:
	//
	// 151977812
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The type of the table. This is a reserved parameter.
	//
	// example:
	//
	// yuyang_test_0000
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) String() string {
	return dara.Prettify(s)
}

func (s GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) GoString() string {
	return s.String()
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) GetTableId() *string {
	return s.TableId
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) GetTableName() *string {
	return s.TableName
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) GetTableType() *string {
	return s.TableType
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) SetTableId(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList {
	s.TableId = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) SetTableName(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList {
	s.TableName = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) SetTableType(v string) *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList {
	s.TableType = &v
	return s
}

func (s *GetTableDBTopologyResponseBodyDBTopologyDataSourceListDatabaseListTableList) Validate() error {
	return dara.Validate(s)
}
