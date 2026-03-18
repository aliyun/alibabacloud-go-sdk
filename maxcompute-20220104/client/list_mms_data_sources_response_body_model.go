// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMmsDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMmsDataSourcesResponseBodyData) *ListMmsDataSourcesResponseBody
	GetData() *ListMmsDataSourcesResponseBodyData
	SetRequestId(v string) *ListMmsDataSourcesResponseBody
	GetRequestId() *string
}

type ListMmsDataSourcesResponseBody struct {
	// The returned data.
	Data *ListMmsDataSourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C1F7715F-D316-5AB6-BD02-5241083F4003
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMmsDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponseBody) GetData() *ListMmsDataSourcesResponseBodyData {
	return s.Data
}

func (s *ListMmsDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMmsDataSourcesResponseBody) SetData(v *ListMmsDataSourcesResponseBodyData) *ListMmsDataSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListMmsDataSourcesResponseBody) SetRequestId(v string) *ListMmsDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMmsDataSourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMmsDataSourcesResponseBodyData struct {
	// The list of data sources.
	ObjectList []*ListMmsDataSourcesResponseBodyDataObjectList `json:"objectList,omitempty" xml:"objectList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 9
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMmsDataSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponseBodyData) GetObjectList() []*ListMmsDataSourcesResponseBodyDataObjectList {
	return s.ObjectList
}

func (s *ListMmsDataSourcesResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMmsDataSourcesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMmsDataSourcesResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListMmsDataSourcesResponseBodyData) SetObjectList(v []*ListMmsDataSourcesResponseBodyDataObjectList) *ListMmsDataSourcesResponseBodyData {
	s.ObjectList = v
	return s
}

func (s *ListMmsDataSourcesResponseBodyData) SetPageNum(v int32) *ListMmsDataSourcesResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyData) SetPageSize(v int32) *ListMmsDataSourcesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyData) SetTotal(v int32) *ListMmsDataSourcesResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyData) Validate() error {
	if s.ObjectList != nil {
		for _, item := range s.ObjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMmsDataSourcesResponseBodyDataObjectList struct {
	// Indicates whether the data source instance or its associated agent is started.
	//
	// example:
	//
	// true
	AgentIsOnline *bool `json:"agentIsOnline,omitempty" xml:"agentIsOnline,omitempty"`
	// The configurations of the data source.
	Config []*ListMmsDataSourcesResponseBodyDataObjectListConfig `json:"config,omitempty" xml:"config,omitempty" type:"Repeated"`
	// The time when the data source was created.
	//
	// example:
	//
	// 2024-12-17 09:29:58
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The number of databases in the data source.
	//
	// example:
	//
	// 3
	DbNum *int32 `json:"dbNum,omitempty" xml:"dbNum,omitempty"`
	// The default destination MaxCompute project.
	//
	// example:
	//
	// mms_test
	DstProject *string `json:"dstProject,omitempty" xml:"dstProject,omitempty"`
	// The list of destination MaxCompute projects.
	DstProjects []*string `json:"dstProjects,omitempty" xml:"dstProjects,omitempty" type:"Repeated"`
	// The reason why the data source instance failed to start or shut down. This parameter is returned only when the value of \\`status\\` is \\`START_FAILED\\` or \\`STOP_FAILED\\`.
	//
	// example:
	//
	// unexpected exception
	ErrMsg *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 2000015
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// The last time when the metadata was synchronized.
	//
	// example:
	//
	// 2024-12-17 15:44:17
	LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The MaxCompute network connectivity ID is the region ID.
	//
	// example:
	//
	// vpc-2zebqp6uojhdla46677tl:cn-beijing
	Networklink *string `json:"networklink,omitempty" xml:"networklink,omitempty"`
	// The number of partitions in the data source.
	//
	// example:
	//
	// 10000000
	PartitionNum *int32 `json:"partitionNum,omitempty" xml:"partitionNum,omitempty"`
	// The number of partitions that are being migrated.
	//
	// example:
	//
	// 2332
	PartitionsDoingNum *int32 `json:"partitionsDoingNum,omitempty" xml:"partitionsDoingNum,omitempty"`
	// The number of partitions that have been migrated.
	//
	// example:
	//
	// 23
	PartitionsDoneNum *int32 `json:"partitionsDoneNum,omitempty" xml:"partitionsDoneNum,omitempty"`
	// The number of partitions that failed to be migrated.
	//
	// example:
	//
	// 2323
	PartitionsFailedNum *int32 `json:"partitionsFailedNum,omitempty" xml:"partitionsFailedNum,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The ID of the metadata synchronization task.
	//
	// example:
	//
	// 1000253
	ScanId *int64 `json:"scanId,omitempty" xml:"scanId,omitempty"`
	// The status of the data source.
	//
	// example:
	//
	// STARTED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The number of tables in the data source.
	//
	// example:
	//
	// 1000
	TableNum *int32 `json:"tableNum,omitempty" xml:"tableNum,omitempty"`
	// The number of tables that are being migrated.
	//
	// example:
	//
	// 18
	TablesDoingNum *int32 `json:"tablesDoingNum,omitempty" xml:"tablesDoingNum,omitempty"`
	// The number of tables that have been migrated.
	//
	// example:
	//
	// 2323
	TablesDoneNum *int32 `json:"tablesDoneNum,omitempty" xml:"tablesDoneNum,omitempty"`
	// The number of tables that failed to be migrated.
	//
	// example:
	//
	// 2
	TablesFailedNum *int32 `json:"tablesFailedNum,omitempty" xml:"tablesFailedNum,omitempty"`
	// The number of tables that are partially migrated.
	//
	// example:
	//
	// 22
	TablesPartDoneNum *int32 `json:"tablesPartDoneNum,omitempty" xml:"tablesPartDoneNum,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListMmsDataSourcesResponseBodyDataObjectList) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourcesResponseBodyDataObjectList) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetAgentIsOnline() *bool {
	return s.AgentIsOnline
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetConfig() []*ListMmsDataSourcesResponseBodyDataObjectListConfig {
	return s.Config
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetDbNum() *int32 {
	return s.DbNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetDstProject() *string {
	return s.DstProject
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetDstProjects() []*string {
	return s.DstProjects
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetId() *int64 {
	return s.Id
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetName() *string {
	return s.Name
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetNetworklink() *string {
	return s.Networklink
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetPartitionsDoingNum() *int32 {
	return s.PartitionsDoingNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetPartitionsDoneNum() *int32 {
	return s.PartitionsDoneNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetPartitionsFailedNum() *int32 {
	return s.PartitionsFailedNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetRegion() *string {
	return s.Region
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetScanId() *int64 {
	return s.ScanId
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetStatus() *string {
	return s.Status
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetTableNum() *int32 {
	return s.TableNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetTablesDoingNum() *int32 {
	return s.TablesDoingNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetTablesDoneNum() *int32 {
	return s.TablesDoneNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetTablesFailedNum() *int32 {
	return s.TablesFailedNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetTablesPartDoneNum() *int32 {
	return s.TablesPartDoneNum
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) GetType() *string {
	return s.Type
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetAgentIsOnline(v bool) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.AgentIsOnline = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetConfig(v []*ListMmsDataSourcesResponseBodyDataObjectListConfig) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Config = v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetCreateTime(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.CreateTime = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetDbNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.DbNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetDstProject(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.DstProject = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetDstProjects(v []*string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.DstProjects = v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetErrMsg(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.ErrMsg = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetId(v int64) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Id = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetLastUpdateTime(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.LastUpdateTime = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetName(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Name = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetNetworklink(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Networklink = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetPartitionNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.PartitionNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetPartitionsDoingNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.PartitionsDoingNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetPartitionsDoneNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.PartitionsDoneNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetPartitionsFailedNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.PartitionsFailedNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetRegion(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Region = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetScanId(v int64) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.ScanId = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetStatus(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Status = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTableNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TableNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTablesDoingNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TablesDoingNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTablesDoneNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TablesDoneNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTablesFailedNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TablesFailedNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetTablesPartDoneNum(v int32) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.TablesPartDoneNum = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) SetType(v string) *ListMmsDataSourcesResponseBodyDataObjectList {
	s.Type = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectList) Validate() error {
	if s.Config != nil {
		for _, item := range s.Config {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMmsDataSourcesResponseBodyDataObjectListConfig struct {
	// The description of the configuration.
	//
	// example:
	//
	// 范围分区表迁移方式
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The enumerated values for the configuration.
	Enums []*string `json:"enums,omitempty" xml:"enums,omitempty" type:"Repeated"`
	// The configuration group.
	//
	// example:
	//
	// basic_group
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// The English identifier for the configuration.
	//
	// example:
	//
	// bigquery.range.partition.migrate.type
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The name of the configuration.
	//
	// example:
	//
	// 范围分区表迁移方式
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// An example value for the configuration.
	//
	// example:
	//
	// Cluster or Partition
	PlaceHolder *string `json:"placeHolder,omitempty" xml:"placeHolder,omitempty"`
	// Indicates whether the configuration is required.
	//
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// If \\`type\\` is set to \\`file\\`, \\`subType\\` specifies the file type, such as \\`.keytab\\`.
	//
	// example:
	//
	// .keytab
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// The type of the configuration. Valid values: \\`boolean\\`, \\`int\\`, \\`map\\`, \\`string\\`, \\`password\\`, and \\`file\\`.
	//
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The value of the configuration.
	//
	// example:
	//
	// Partition
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListMmsDataSourcesResponseBodyDataObjectListConfig) String() string {
	return dara.Prettify(s)
}

func (s ListMmsDataSourcesResponseBodyDataObjectListConfig) GoString() string {
	return s.String()
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetDesc() *string {
	return s.Desc
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetEnums() []*string {
	return s.Enums
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetGroup() *string {
	return s.Group
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetKey() *string {
	return s.Key
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetName() *string {
	return s.Name
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetPlaceHolder() *string {
	return s.PlaceHolder
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetRequired() *bool {
	return s.Required
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetSubType() *string {
	return s.SubType
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetType() *string {
	return s.Type
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) GetValue() interface{} {
	return s.Value
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetDesc(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Desc = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetEnums(v []*string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Enums = v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetGroup(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Group = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetKey(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Key = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetName(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Name = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetPlaceHolder(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.PlaceHolder = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetRequired(v bool) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Required = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetSubType(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.SubType = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetType(v string) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Type = &v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) SetValue(v interface{}) *ListMmsDataSourcesResponseBodyDataObjectListConfig {
	s.Value = v
	return s
}

func (s *ListMmsDataSourcesResponseBodyDataObjectListConfig) Validate() error {
	return dara.Validate(s)
}
