// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTableResponseBody
	GetCode() *string
	SetData(v *GetTableResponseBodyData) *GetTableResponseBody
	GetData() *GetTableResponseBodyData
	SetHttpStatusCode(v int32) *GetTableResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTableResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTableResponseBody
	GetSuccess() *bool
}

type GetTableResponseBody struct {
	// example:
	//
	// OK
	Code *string                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTableResponseBody) GetData() *GetTableResponseBodyData {
	return s.Data
}

func (s *GetTableResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTableResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTableResponseBody) SetCode(v string) *GetTableResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableResponseBody) SetData(v *GetTableResponseBodyData) *GetTableResponseBody {
	s.Data = v
	return s
}

func (s *GetTableResponseBody) SetHttpStatusCode(v int32) *GetTableResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTableResponseBody) SetMessage(v string) *GetTableResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableResponseBody) SetRequestId(v string) *GetTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableResponseBody) SetSuccess(v bool) *GetTableResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTableResponseBodyData struct {
	AssetTags []*string `json:"AssetTags,omitempty" xml:"AssetTags,omitempty" type:"Repeated"`
	// example:
	//
	// 2011
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// LD_test01
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 30011211
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 211
	DataDomainId *int64 `json:"DataDomainId,omitempty" xml:"DataDomainId,omitempty"`
	// example:
	//
	// 课程域
	DataDomainName *string `json:"DataDomainName,omitempty" xml:"DataDomainName,omitempty"`
	// example:
	//
	// 3301
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// 学生
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// dev
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 2
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// dp_ds_table.300023201.7311626611751680256.load_test.abc
	Guid             *string                                 `json:"Guid,omitempty" xml:"Guid,omitempty"`
	Instructions     []*GetTableResponseBodyDataInstructions `json:"Instructions,omitempty" xml:"Instructions,omitempty" type:"Repeated"`
	IsBasicMode      *bool                                   `json:"IsBasicMode,omitempty" xml:"IsBasicMode,omitempty"`
	IsPartitionTable *bool                                   `json:"IsPartitionTable,omitempty" xml:"IsPartitionTable,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	LastDdlTime *string `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	LastDmlTime *string `json:"LastDmlTime,omitempty" xml:"LastDmlTime,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	LastQueryTime *string `json:"LastQueryTime,omitempty" xml:"LastQueryTime,omitempty"`
	// example:
	//
	// 30
	LifeCycle *int64 `json:"LifeCycle,omitempty" xml:"LifeCycle,omitempty"`
	// example:
	//
	// t_test01
	Name    *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// example:
	//
	// 30011211
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	ParentModelId *string `json:"ParentModelId,omitempty" xml:"ParentModelId,omitempty"`
	// example:
	//
	// 1011
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// testPrj
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1
	SecurityLevel *int64 `json:"SecurityLevel,omitempty" xml:"SecurityLevel,omitempty"`
	// example:
	//
	// 高
	SecurityLevelAbbreviation *string `json:"SecurityLevelAbbreviation,omitempty" xml:"SecurityLevelAbbreviation,omitempty"`
	// example:
	//
	// 高级
	SecurityLevelName *string                                    `json:"SecurityLevelName,omitempty" xml:"SecurityLevelName,omitempty"`
	SimpleNodeInfos   []*GetTableResponseBodyDataSimpleNodeInfos `json:"SimpleNodeInfos,omitempty" xml:"SimpleNodeInfos,omitempty" type:"Repeated"`
	// example:
	//
	// HIVE
	StorageType       *string                                      `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	StreamTableConfig []*GetTableResponseBodyDataStreamTableConfig `json:"StreamTableConfig,omitempty" xml:"StreamTableConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 10241024
	TableSizeInBytes *int64 `json:"TableSizeInBytes,omitempty" xml:"TableSizeInBytes,omitempty"`
	// example:
	//
	// 22
	VisitCount30d *int64 `json:"VisitCount30d,omitempty" xml:"VisitCount30d,omitempty"`
}

func (s GetTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyData) GetAssetTags() []*string {
	return s.AssetTags
}

func (s *GetTableResponseBodyData) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetTableResponseBodyData) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetTableResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *GetTableResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetTableResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *GetTableResponseBodyData) GetDataDomainId() *int64 {
	return s.DataDomainId
}

func (s *GetTableResponseBodyData) GetDataDomainName() *string {
	return s.DataDomainName
}

func (s *GetTableResponseBodyData) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *GetTableResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTableResponseBodyData) GetEnv() *string {
	return s.Env
}

func (s *GetTableResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *GetTableResponseBodyData) GetGuid() *string {
	return s.Guid
}

func (s *GetTableResponseBodyData) GetInstructions() []*GetTableResponseBodyDataInstructions {
	return s.Instructions
}

func (s *GetTableResponseBodyData) GetIsBasicMode() *bool {
	return s.IsBasicMode
}

func (s *GetTableResponseBodyData) GetIsPartitionTable() *bool {
	return s.IsPartitionTable
}

func (s *GetTableResponseBodyData) GetLastDdlTime() *string {
	return s.LastDdlTime
}

func (s *GetTableResponseBodyData) GetLastDmlTime() *string {
	return s.LastDmlTime
}

func (s *GetTableResponseBodyData) GetLastQueryTime() *string {
	return s.LastQueryTime
}

func (s *GetTableResponseBodyData) GetLifeCycle() *int64 {
	return s.LifeCycle
}

func (s *GetTableResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTableResponseBodyData) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *GetTableResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetTableResponseBodyData) GetParentModelId() *string {
	return s.ParentModelId
}

func (s *GetTableResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetTableResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTableResponseBodyData) GetSecurityLevel() *int64 {
	return s.SecurityLevel
}

func (s *GetTableResponseBodyData) GetSecurityLevelAbbreviation() *string {
	return s.SecurityLevelAbbreviation
}

func (s *GetTableResponseBodyData) GetSecurityLevelName() *string {
	return s.SecurityLevelName
}

func (s *GetTableResponseBodyData) GetSimpleNodeInfos() []*GetTableResponseBodyDataSimpleNodeInfos {
	return s.SimpleNodeInfos
}

func (s *GetTableResponseBodyData) GetStorageType() *string {
	return s.StorageType
}

func (s *GetTableResponseBodyData) GetStreamTableConfig() []*GetTableResponseBodyDataStreamTableConfig {
	return s.StreamTableConfig
}

func (s *GetTableResponseBodyData) GetTableSizeInBytes() *int64 {
	return s.TableSizeInBytes
}

func (s *GetTableResponseBodyData) GetVisitCount30d() *int64 {
	return s.VisitCount30d
}

func (s *GetTableResponseBodyData) SetAssetTags(v []*string) *GetTableResponseBodyData {
	s.AssetTags = v
	return s
}

func (s *GetTableResponseBodyData) SetBizUnitId(v int64) *GetTableResponseBodyData {
	s.BizUnitId = &v
	return s
}

func (s *GetTableResponseBodyData) SetBizUnitName(v string) *GetTableResponseBodyData {
	s.BizUnitName = &v
	return s
}

func (s *GetTableResponseBodyData) SetComment(v string) *GetTableResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetTableResponseBodyData) SetCreateTime(v string) *GetTableResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetTableResponseBodyData) SetCreator(v string) *GetTableResponseBodyData {
	s.Creator = &v
	return s
}

func (s *GetTableResponseBodyData) SetDataDomainId(v int64) *GetTableResponseBodyData {
	s.DataDomainId = &v
	return s
}

func (s *GetTableResponseBodyData) SetDataDomainName(v string) *GetTableResponseBodyData {
	s.DataDomainName = &v
	return s
}

func (s *GetTableResponseBodyData) SetDataSourceId(v int64) *GetTableResponseBodyData {
	s.DataSourceId = &v
	return s
}

func (s *GetTableResponseBodyData) SetDisplayName(v string) *GetTableResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetTableResponseBodyData) SetEnv(v string) *GetTableResponseBodyData {
	s.Env = &v
	return s
}

func (s *GetTableResponseBodyData) SetFileId(v string) *GetTableResponseBodyData {
	s.FileId = &v
	return s
}

func (s *GetTableResponseBodyData) SetGuid(v string) *GetTableResponseBodyData {
	s.Guid = &v
	return s
}

func (s *GetTableResponseBodyData) SetInstructions(v []*GetTableResponseBodyDataInstructions) *GetTableResponseBodyData {
	s.Instructions = v
	return s
}

func (s *GetTableResponseBodyData) SetIsBasicMode(v bool) *GetTableResponseBodyData {
	s.IsBasicMode = &v
	return s
}

func (s *GetTableResponseBodyData) SetIsPartitionTable(v bool) *GetTableResponseBodyData {
	s.IsPartitionTable = &v
	return s
}

func (s *GetTableResponseBodyData) SetLastDdlTime(v string) *GetTableResponseBodyData {
	s.LastDdlTime = &v
	return s
}

func (s *GetTableResponseBodyData) SetLastDmlTime(v string) *GetTableResponseBodyData {
	s.LastDmlTime = &v
	return s
}

func (s *GetTableResponseBodyData) SetLastQueryTime(v string) *GetTableResponseBodyData {
	s.LastQueryTime = &v
	return s
}

func (s *GetTableResponseBodyData) SetLifeCycle(v int64) *GetTableResponseBodyData {
	s.LifeCycle = &v
	return s
}

func (s *GetTableResponseBodyData) SetName(v string) *GetTableResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTableResponseBodyData) SetNodeIds(v []*string) *GetTableResponseBodyData {
	s.NodeIds = v
	return s
}

func (s *GetTableResponseBodyData) SetOwner(v string) *GetTableResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetTableResponseBodyData) SetParentModelId(v string) *GetTableResponseBodyData {
	s.ParentModelId = &v
	return s
}

func (s *GetTableResponseBodyData) SetProjectId(v int64) *GetTableResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetTableResponseBodyData) SetProjectName(v string) *GetTableResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetTableResponseBodyData) SetSecurityLevel(v int64) *GetTableResponseBodyData {
	s.SecurityLevel = &v
	return s
}

func (s *GetTableResponseBodyData) SetSecurityLevelAbbreviation(v string) *GetTableResponseBodyData {
	s.SecurityLevelAbbreviation = &v
	return s
}

func (s *GetTableResponseBodyData) SetSecurityLevelName(v string) *GetTableResponseBodyData {
	s.SecurityLevelName = &v
	return s
}

func (s *GetTableResponseBodyData) SetSimpleNodeInfos(v []*GetTableResponseBodyDataSimpleNodeInfos) *GetTableResponseBodyData {
	s.SimpleNodeInfos = v
	return s
}

func (s *GetTableResponseBodyData) SetStorageType(v string) *GetTableResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *GetTableResponseBodyData) SetStreamTableConfig(v []*GetTableResponseBodyDataStreamTableConfig) *GetTableResponseBodyData {
	s.StreamTableConfig = v
	return s
}

func (s *GetTableResponseBodyData) SetTableSizeInBytes(v int64) *GetTableResponseBodyData {
	s.TableSizeInBytes = &v
	return s
}

func (s *GetTableResponseBodyData) SetVisitCount30d(v int64) *GetTableResponseBodyData {
	s.VisitCount30d = &v
	return s
}

func (s *GetTableResponseBodyData) Validate() error {
	if s.Instructions != nil {
		for _, item := range s.Instructions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SimpleNodeInfos != nil {
		for _, item := range s.SimpleNodeInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type GetTableResponseBodyDataInstructions struct {
	// example:
	//
	// <p>示例内容</p>
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-06-30 00:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 30011211
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 张三
	OwnerNickName *string `json:"OwnerNickName,omitempty" xml:"OwnerNickName,omitempty"`
	// example:
	//
	// 使用指南
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetTableResponseBodyDataInstructions) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyDataInstructions) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyDataInstructions) GetContent() *string {
	return s.Content
}

func (s *GetTableResponseBodyDataInstructions) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetTableResponseBodyDataInstructions) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetTableResponseBodyDataInstructions) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetTableResponseBodyDataInstructions) GetOwnerNickName() *string {
	return s.OwnerNickName
}

func (s *GetTableResponseBodyDataInstructions) GetTitle() *string {
	return s.Title
}

func (s *GetTableResponseBodyDataInstructions) SetContent(v string) *GetTableResponseBodyDataInstructions {
	s.Content = &v
	return s
}

func (s *GetTableResponseBodyDataInstructions) SetGmtCreate(v string) *GetTableResponseBodyDataInstructions {
	s.GmtCreate = &v
	return s
}

func (s *GetTableResponseBodyDataInstructions) SetGmtModified(v string) *GetTableResponseBodyDataInstructions {
	s.GmtModified = &v
	return s
}

func (s *GetTableResponseBodyDataInstructions) SetOwnerId(v string) *GetTableResponseBodyDataInstructions {
	s.OwnerId = &v
	return s
}

func (s *GetTableResponseBodyDataInstructions) SetOwnerNickName(v string) *GetTableResponseBodyDataInstructions {
	s.OwnerNickName = &v
	return s
}

func (s *GetTableResponseBodyDataInstructions) SetTitle(v string) *GetTableResponseBodyDataInstructions {
	s.Title = &v
	return s
}

func (s *GetTableResponseBodyDataInstructions) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyDataSimpleNodeInfos struct {
	BizUnit *GetTableResponseBodyDataSimpleNodeInfosBizUnit `json:"BizUnit,omitempty" xml:"BizUnit,omitempty" type:"Struct"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// n_7443xxxx
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 2345
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// example:
	//
	// NORMAL
	NodeScheduleType *string                                          `json:"NodeScheduleType,omitempty" xml:"NodeScheduleType,omitempty"`
	Owners           []*GetTableResponseBodyDataSimpleNodeInfosOwners `json:"Owners,omitempty" xml:"Owners,omitempty" type:"Repeated"`
	Project          *GetTableResponseBodyDataSimpleNodeInfosProject  `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// example:
	//
	// DLINK
	SubBizType *string `json:"SubBizType,omitempty" xml:"SubBizType,omitempty"`
}

func (s GetTableResponseBodyDataSimpleNodeInfos) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyDataSimpleNodeInfos) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) GetBizUnit() *GetTableResponseBodyDataSimpleNodeInfosBizUnit {
	return s.BizUnit
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) GetEnv() *string {
	return s.Env
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) GetNodeId() *string {
	return s.NodeId
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) GetNodeName() *string {
	return s.NodeName
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) GetNodeScheduleType() *string {
	return s.NodeScheduleType
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) GetOwners() []*GetTableResponseBodyDataSimpleNodeInfosOwners {
	return s.Owners
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) GetProject() *GetTableResponseBodyDataSimpleNodeInfosProject {
	return s.Project
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) GetSubBizType() *string {
	return s.SubBizType
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) SetBizUnit(v *GetTableResponseBodyDataSimpleNodeInfosBizUnit) *GetTableResponseBodyDataSimpleNodeInfos {
	s.BizUnit = v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) SetEnv(v string) *GetTableResponseBodyDataSimpleNodeInfos {
	s.Env = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) SetNodeId(v string) *GetTableResponseBodyDataSimpleNodeInfos {
	s.NodeId = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) SetNodeName(v string) *GetTableResponseBodyDataSimpleNodeInfos {
	s.NodeName = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) SetNodeScheduleType(v string) *GetTableResponseBodyDataSimpleNodeInfos {
	s.NodeScheduleType = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) SetOwners(v []*GetTableResponseBodyDataSimpleNodeInfosOwners) *GetTableResponseBodyDataSimpleNodeInfos {
	s.Owners = v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) SetProject(v *GetTableResponseBodyDataSimpleNodeInfosProject) *GetTableResponseBodyDataSimpleNodeInfos {
	s.Project = v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) SetSubBizType(v string) *GetTableResponseBodyDataSimpleNodeInfos {
	s.SubBizType = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfos) Validate() error {
	if s.BizUnit != nil {
		if err := s.BizUnit.Validate(); err != nil {
			return err
		}
	}
	if s.Owners != nil {
		for _, item := range s.Owners {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTableResponseBodyDataSimpleNodeInfosBizUnit struct {
	// example:
	//
	// 测试板块
	BizUnitDisplayName *string `json:"BizUnitDisplayName,omitempty" xml:"BizUnitDisplayName,omitempty"`
	// example:
	//
	// 2011
	BizUnitId *string `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// LD_test01
	BizUnitName *string `json:"BizUnitName,omitempty" xml:"BizUnitName,omitempty"`
}

func (s GetTableResponseBodyDataSimpleNodeInfosBizUnit) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyDataSimpleNodeInfosBizUnit) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyDataSimpleNodeInfosBizUnit) GetBizUnitDisplayName() *string {
	return s.BizUnitDisplayName
}

func (s *GetTableResponseBodyDataSimpleNodeInfosBizUnit) GetBizUnitId() *string {
	return s.BizUnitId
}

func (s *GetTableResponseBodyDataSimpleNodeInfosBizUnit) GetBizUnitName() *string {
	return s.BizUnitName
}

func (s *GetTableResponseBodyDataSimpleNodeInfosBizUnit) SetBizUnitDisplayName(v string) *GetTableResponseBodyDataSimpleNodeInfosBizUnit {
	s.BizUnitDisplayName = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfosBizUnit) SetBizUnitId(v string) *GetTableResponseBodyDataSimpleNodeInfosBizUnit {
	s.BizUnitId = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfosBizUnit) SetBizUnitName(v string) *GetTableResponseBodyDataSimpleNodeInfosBizUnit {
	s.BizUnitName = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfosBizUnit) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyDataSimpleNodeInfosOwners struct {
	// example:
	//
	// 张三
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetTableResponseBodyDataSimpleNodeInfosOwners) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyDataSimpleNodeInfosOwners) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyDataSimpleNodeInfosOwners) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTableResponseBodyDataSimpleNodeInfosOwners) GetUserId() *string {
	return s.UserId
}

func (s *GetTableResponseBodyDataSimpleNodeInfosOwners) SetDisplayName(v string) *GetTableResponseBodyDataSimpleNodeInfosOwners {
	s.DisplayName = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfosOwners) SetUserId(v string) *GetTableResponseBodyDataSimpleNodeInfosOwners {
	s.UserId = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfosOwners) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyDataSimpleNodeInfosProject struct {
	// example:
	//
	// 测试项目
	ProjectDisplayName *string `json:"ProjectDisplayName,omitempty" xml:"ProjectDisplayName,omitempty"`
	// example:
	//
	// 1011
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// testPrj
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetTableResponseBodyDataSimpleNodeInfosProject) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyDataSimpleNodeInfosProject) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyDataSimpleNodeInfosProject) GetProjectDisplayName() *string {
	return s.ProjectDisplayName
}

func (s *GetTableResponseBodyDataSimpleNodeInfosProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetTableResponseBodyDataSimpleNodeInfosProject) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTableResponseBodyDataSimpleNodeInfosProject) SetProjectDisplayName(v string) *GetTableResponseBodyDataSimpleNodeInfosProject {
	s.ProjectDisplayName = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfosProject) SetProjectId(v string) *GetTableResponseBodyDataSimpleNodeInfosProject {
	s.ProjectId = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfosProject) SetProjectName(v string) *GetTableResponseBodyDataSimpleNodeInfosProject {
	s.ProjectName = &v
	return s
}

func (s *GetTableResponseBodyDataSimpleNodeInfosProject) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyDataStreamTableConfig struct {
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetTableResponseBodyDataStreamTableConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyDataStreamTableConfig) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyDataStreamTableConfig) GetKey() *string {
	return s.Key
}

func (s *GetTableResponseBodyDataStreamTableConfig) GetValue() *string {
	return s.Value
}

func (s *GetTableResponseBodyDataStreamTableConfig) SetKey(v string) *GetTableResponseBodyDataStreamTableConfig {
	s.Key = &v
	return s
}

func (s *GetTableResponseBodyDataStreamTableConfig) SetValue(v string) *GetTableResponseBodyDataStreamTableConfig {
	s.Value = &v
	return s
}

func (s *GetTableResponseBodyDataStreamTableConfig) Validate() error {
	return dara.Validate(s)
}
