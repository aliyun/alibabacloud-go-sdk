// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTablesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListTablesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListTablesResponseBody
	GetMessage() *string
	SetPageResult(v *ListTablesResponseBodyPageResult) *ListTablesResponseBody
	GetPageResult() *ListTablesResponseBodyPageResult
	SetRequestId(v string) *ListTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListTablesResponseBody
	GetSuccess() *bool
}

type ListTablesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP response code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paged query result.
	PageResult *ListTablesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTablesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListTablesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTablesResponseBody) GetPageResult() *ListTablesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListTablesResponseBody) SetCode(v string) *ListTablesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTablesResponseBody) SetHttpStatusCode(v int32) *ListTablesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListTablesResponseBody) SetMessage(v string) *ListTablesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTablesResponseBody) SetPageResult(v *ListTablesResponseBodyPageResult) *ListTablesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetSuccess(v bool) *ListTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTablesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListTablesResponseBodyPageResult struct {
	// The table list.
	TableList []*ListTablesResponseBodyPageResultTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTablesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyPageResult) GetTableList() []*ListTablesResponseBodyPageResultTableList {
	return s.TableList
}

func (s *ListTablesResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTablesResponseBodyPageResult) SetTableList(v []*ListTablesResponseBodyPageResultTableList) *ListTablesResponseBodyPageResult {
	s.TableList = v
	return s
}

func (s *ListTablesResponseBodyPageResult) SetTotalCount(v int32) *ListTablesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListTablesResponseBodyPageResult) Validate() error {
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

type ListTablesResponseBodyPageResultTableList struct {
	// The asset inventory tags.
	AssetTagList []*string `json:"AssetTagList,omitempty" xml:"AssetTagList,omitempty" type:"Repeated"`
	// The business unit ID.
	//
	// example:
	//
	// 2011
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The business unit name.
	//
	// example:
	//
	// LD_test01
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// The table comment.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator ID.
	//
	// example:
	//
	// 30011211
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The data domain ID.
	//
	// example:
	//
	// 211
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
	// The subject area name.
	//
	// example:
	//
	// 211
	DataDomainName *string `json:"DataDomainName,omitempty" xml:"DataDomainName,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 3301
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The table display name.
	//
	// example:
	//
	// Student
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environment. Valid values: dev and prod.
	//
	// example:
	//
	// dev
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The file ID used to navigate to the ODM model.
	//
	// example:
	//
	// 2
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The table GUID.
	//
	// example:
	//
	// 10011
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// Specifies whether the project or business version to which the table belongs is in Basic mode.
	IsBasicMode *bool `json:"IsBasicMode,omitempty" xml:"IsBasicMode,omitempty"`
	// Specifies whether the table is a partitioned table.
	IsPartitionTable *bool `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	// The time when the table DDL was last modified.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	LastDdlTime *string `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	// The time when the table data was last updated.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	LastDmlTime *string `json:"LastDmlTime,omitempty" xml:"LastDmlTime,omitempty"`
	// The time when the data was last viewed.
	//
	// example:
	//
	// 2025-06-30 00:00:00
	LastQueryTime *string `json:"LastQueryTime,omitempty" xml:"LastQueryTime,omitempty"`
	// The lifecycle of the table, in days.
	//
	// example:
	//
	// 30
	LifeCycle *int64 `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	// The table name.
	//
	// example:
	//
	// t_test01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user ID of the table owner in Dataphin. This parameter may be empty if the owner is not bound.
	//
	// example:
	//
	// 30011211
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parent model ID used to navigate to the ODM model.
	//
	// example:
	//
	// 1
	ParentModelId *string `json:"ParentModelId,omitempty" xml:"ParentModelId,omitempty"`
	// The logical project ID.
	//
	// example:
	//
	// 1011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The logical project name.
	//
	// example:
	//
	// testPrj
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The security classification.
	//
	// example:
	//
	// 1
	SecurityLevel *int64 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// The abbreviation of the security classification name.
	//
	// example:
	//
	// High
	SecurityLevelAbbreviation *string `json:"SecurityLevelAbbreviation,omitempty" xml:"SecurityLevelAbbreviation,omitempty"`
	// The security classification name.
	//
	// example:
	//
	// High
	SecurityLevelName *string `json:"SecurityLevelName,omitempty" xml:"SecurityLevelName,omitempty"`
	// The storage type. Valid values: HIVE and MAX_COMPUTE.
	//
	// example:
	//
	// HIVE
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// The dynamic properties of the real-time metatable.
	StreamTableConfig []*ListTablesResponseBodyPageResultTableListStreamTableConfig `json:"StreamTableConfig,omitempty" xml:"StreamTableConfig,omitempty" type:"Repeated"`
	// The total storage size of the table, in bytes.
	//
	// example:
	//
	// 10241024
	TableSizeInBytes *int64 `json:"TableSizeInBytes,omitempty" xml:"TableSizeInBytes,omitempty"`
	// The number of visits in the last 30 days.
	//
	// example:
	//
	// 22
	VisitCount30d *int64 `json:"VisitCount30d,omitempty" xml:"VisitCount30d,omitempty"`
}

func (s ListTablesResponseBodyPageResultTableList) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyPageResultTableList) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyPageResultTableList) GetAssetTagList() []*string {
	return s.AssetTagList
}

func (s *ListTablesResponseBodyPageResultTableList) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListTablesResponseBodyPageResultTableList) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *ListTablesResponseBodyPageResultTableList) GetComment() *string {
	return s.Comment
}

func (s *ListTablesResponseBodyPageResultTableList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTablesResponseBodyPageResultTableList) GetCreator() *string {
	return s.Creator
}

func (s *ListTablesResponseBodyPageResultTableList) GetDataDomainId() *int64 {
	return s.DataDomainId
}

func (s *ListTablesResponseBodyPageResultTableList) GetDataDomainName() *string {
	return s.DataDomainName
}

func (s *ListTablesResponseBodyPageResultTableList) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *ListTablesResponseBodyPageResultTableList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListTablesResponseBodyPageResultTableList) GetEnv() *string {
	return s.Env
}

func (s *ListTablesResponseBodyPageResultTableList) GetFileId() *string {
	return s.FileId
}

func (s *ListTablesResponseBodyPageResultTableList) GetGuid() *string {
	return s.Guid
}

func (s *ListTablesResponseBodyPageResultTableList) GetIsBasicMode() *bool {
	return s.IsBasicMode
}

func (s *ListTablesResponseBodyPageResultTableList) GetIsPartitionTable() *bool {
	return s.IsPartitionTable
}

func (s *ListTablesResponseBodyPageResultTableList) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *ListTablesResponseBodyPageResultTableList) GetLastDmlTime() *string {
	return s.LastDmlTime
}

func (s *ListTablesResponseBodyPageResultTableList) GetLastQueryTime() *string {
	return s.LastQueryTime
}

func (s *ListTablesResponseBodyPageResultTableList) GetLifeCycle() *int64 {
	return s.LifeCycle
}

func (s *ListTablesResponseBodyPageResultTableList) GetName() *string {
	return s.Name
}

func (s *ListTablesResponseBodyPageResultTableList) GetOwner() *string {
	return s.Owner
}

func (s *ListTablesResponseBodyPageResultTableList) GetParentModelId() *string {
	return s.ParentModelId
}

func (s *ListTablesResponseBodyPageResultTableList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListTablesResponseBodyPageResultTableList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListTablesResponseBodyPageResultTableList) GetSecurityLevel() *int64 {
	return s.SecurityLevel
}

func (s *ListTablesResponseBodyPageResultTableList) GetSecurityLevelAbbreviation() *string {
	return s.SecurityLevelAbbreviation
}

func (s *ListTablesResponseBodyPageResultTableList) GetSecurityLevelName() *string {
	return s.SecurityLevelName
}

func (s *ListTablesResponseBodyPageResultTableList) GetStorageType() *string {
	return s.StorageType
}

func (s *ListTablesResponseBodyPageResultTableList) GetStreamTableConfig() []*ListTablesResponseBodyPageResultTableListStreamTableConfig {
	return s.StreamTableConfig
}

func (s *ListTablesResponseBodyPageResultTableList) GetTableSizeInBytes() *int64 {
	return s.TableSizeInBytes
}

func (s *ListTablesResponseBodyPageResultTableList) GetVisitCount30d() *int64 {
	return s.VisitCount30d
}

func (s *ListTablesResponseBodyPageResultTableList) SetAssetTagList(v []*string) *ListTablesResponseBodyPageResultTableList {
	s.AssetTagList = v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetBizUnitId(v int64) *ListTablesResponseBodyPageResultTableList {
	s.BizUnitId = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetBizUnitName(v string) *ListTablesResponseBodyPageResultTableList {
	s.BizUnitName = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetComment(v string) *ListTablesResponseBodyPageResultTableList {
	s.Comment = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetCreateTime(v string) *ListTablesResponseBodyPageResultTableList {
	s.CreateTime = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetCreator(v string) *ListTablesResponseBodyPageResultTableList {
	s.Creator = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetDataDomainId(v int64) *ListTablesResponseBodyPageResultTableList {
	s.DataDomainId = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetDataDomainName(v string) *ListTablesResponseBodyPageResultTableList {
	s.DataDomainName = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetDataSourceId(v int64) *ListTablesResponseBodyPageResultTableList {
	s.DataSourceId = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetDisplayName(v string) *ListTablesResponseBodyPageResultTableList {
	s.DisplayName = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetEnv(v string) *ListTablesResponseBodyPageResultTableList {
	s.Env = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetFileId(v string) *ListTablesResponseBodyPageResultTableList {
	s.FileId = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetGuid(v string) *ListTablesResponseBodyPageResultTableList {
	s.Guid = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetIsBasicMode(v bool) *ListTablesResponseBodyPageResultTableList {
	s.IsBasicMode = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetIsPartitionTable(v bool) *ListTablesResponseBodyPageResultTableList {
	s.IsPartitionTable = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetLastDdlTime(v string) *ListTablesResponseBodyPageResultTableList {
	s.LastDdlTime = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetLastDmlTime(v string) *ListTablesResponseBodyPageResultTableList {
	s.LastDmlTime = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetLastQueryTime(v string) *ListTablesResponseBodyPageResultTableList {
	s.LastQueryTime = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetLifeCycle(v int64) *ListTablesResponseBodyPageResultTableList {
	s.LifeCycle = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetName(v string) *ListTablesResponseBodyPageResultTableList {
	s.Name = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetOwner(v string) *ListTablesResponseBodyPageResultTableList {
	s.Owner = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetParentModelId(v string) *ListTablesResponseBodyPageResultTableList {
	s.ParentModelId = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetProjectId(v int64) *ListTablesResponseBodyPageResultTableList {
	s.ProjectId = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetProjectName(v string) *ListTablesResponseBodyPageResultTableList {
	s.ProjectName = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetSecurityLevel(v int64) *ListTablesResponseBodyPageResultTableList {
	s.SecurityLevel = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetSecurityLevelAbbreviation(v string) *ListTablesResponseBodyPageResultTableList {
	s.SecurityLevelAbbreviation = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetSecurityLevelName(v string) *ListTablesResponseBodyPageResultTableList {
	s.SecurityLevelName = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetStorageType(v string) *ListTablesResponseBodyPageResultTableList {
	s.StorageType = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetStreamTableConfig(v []*ListTablesResponseBodyPageResultTableListStreamTableConfig) *ListTablesResponseBodyPageResultTableList {
	s.StreamTableConfig = v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetTableSizeInBytes(v int64) *ListTablesResponseBodyPageResultTableList {
	s.TableSizeInBytes = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) SetVisitCount30d(v int64) *ListTablesResponseBodyPageResultTableList {
	s.VisitCount30d = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableList) Validate() error {
	if s.StreamTableConfig != nil {
		for _, item := range s.StreamTableConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTablesResponseBodyPageResultTableListStreamTableConfig struct {
	// The configuration item.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the configuration item.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTablesResponseBodyPageResultTableListStreamTableConfig) String() string {
	return dara.Prettify(s)
}

func (s ListTablesResponseBodyPageResultTableListStreamTableConfig) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBodyPageResultTableListStreamTableConfig) GetKey() *string {
	return s.Key
}

func (s *ListTablesResponseBodyPageResultTableListStreamTableConfig) GetValue() *string {
	return s.Value
}

func (s *ListTablesResponseBodyPageResultTableListStreamTableConfig) SetKey(v string) *ListTablesResponseBodyPageResultTableListStreamTableConfig {
	s.Key = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableListStreamTableConfig) SetValue(v string) *ListTablesResponseBodyPageResultTableListStreamTableConfig {
	s.Value = &v
	return s
}

func (s *ListTablesResponseBodyPageResultTableListStreamTableConfig) Validate() error {
	return dara.Validate(s)
}
