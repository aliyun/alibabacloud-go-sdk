// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsDataSourceResponseBodyData) *GetMmsDataSourceResponseBody
	GetData() *GetMmsDataSourceResponseBodyData
	SetRequestId(v string) *GetMmsDataSourceResponseBody
	GetRequestId() *string
}

type GetMmsDataSourceResponseBody struct {
	Data *GetMmsDataSourceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 98EC8C47-3D6D-560C-808B-84E494220A32
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceResponseBody) GetData() *GetMmsDataSourceResponseBodyData {
	return s.Data
}

func (s *GetMmsDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsDataSourceResponseBody) SetData(v *GetMmsDataSourceResponseBodyData) *GetMmsDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsDataSourceResponseBody) SetRequestId(v string) *GetMmsDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsDataSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsDataSourceResponseBodyData struct {
	// example:
	//
	// true
	AgentIsOnline *bool                                     `json:"agentIsOnline,omitempty" xml:"agentIsOnline,omitempty"`
	Config        []*GetMmsDataSourceResponseBodyDataConfig `json:"config,omitempty" xml:"config,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-12-17 09:29:58
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 3
	DbNum *int32 `json:"dbNum,omitempty" xml:"dbNum,omitempty"`
	// example:
	//
	// mms_test
	DstProject  *string   `json:"dstProject,omitempty" xml:"dstProject,omitempty"`
	DstProjects []*string `json:"dstProjects,omitempty" xml:"dstProjects,omitempty" type:"Repeated"`
	// example:
	//
	// unexpected exception
	ErrMsg *string `json:"errMsg,omitempty" xml:"errMsg,omitempty"`
	// example:
	//
	// 2000015
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2024-12-17 15:44:17
	LastUpdateTime *string `json:"lastUpdateTime,omitempty" xml:"lastUpdateTime,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// vpc-2zebqp6uojhdla46677tl:cn-shanghai
	Networklink *string `json:"networklink,omitempty" xml:"networklink,omitempty"`
	// example:
	//
	// 10000000
	PartitionNum *int32 `json:"partitionNum,omitempty" xml:"partitionNum,omitempty"`
	// example:
	//
	// 23322
	PartitionsDoingNum *int32 `json:"partitionsDoingNum,omitempty" xml:"partitionsDoingNum,omitempty"`
	// example:
	//
	// 11113
	PartitionsDoneNum *int32 `json:"partitionsDoneNum,omitempty" xml:"partitionsDoneNum,omitempty"`
	// example:
	//
	// 32
	PartitionsFailedNum *int32 `json:"partitionsFailedNum,omitempty" xml:"partitionsFailedNum,omitempty"`
	// example:
	//
	// cn-shanghai
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// example:
	//
	// 1000253
	ScanId *int64 `json:"scanId,omitempty" xml:"scanId,omitempty"`
	// example:
	//
	// STARTED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1000
	TableNum *int32 `json:"tableNum,omitempty" xml:"tableNum,omitempty"`
	// example:
	//
	// 19
	TablesDoingNum *int32 `json:"tablesDoingNum,omitempty" xml:"tablesDoingNum,omitempty"`
	// example:
	//
	// 16
	TablesDoneNum *int32 `json:"tablesDoneNum,omitempty" xml:"tablesDoneNum,omitempty"`
	// example:
	//
	// 2
	TablesFailedNum *int32 `json:"tablesFailedNum,omitempty" xml:"tablesFailedNum,omitempty"`
	// example:
	//
	// 123
	TablesPartDoneNum *int32 `json:"tablesPartDoneNum,omitempty" xml:"tablesPartDoneNum,omitempty"`
	// example:
	//
	// BIGQUERY
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetMmsDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceResponseBodyData) GetAgentIsOnline() *bool {
	return s.AgentIsOnline
}

func (s *GetMmsDataSourceResponseBodyData) GetConfig() []*GetMmsDataSourceResponseBodyDataConfig {
	return s.Config
}

func (s *GetMmsDataSourceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMmsDataSourceResponseBodyData) GetDbNum() *int32 {
	return s.DbNum
}

func (s *GetMmsDataSourceResponseBodyData) GetDstProject() *string {
	return s.DstProject
}

func (s *GetMmsDataSourceResponseBodyData) GetDstProjects() []*string {
	return s.DstProjects
}

func (s *GetMmsDataSourceResponseBodyData) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetMmsDataSourceResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsDataSourceResponseBodyData) GetLastUpdateTime() *string {
	return s.LastUpdateTime
}

func (s *GetMmsDataSourceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMmsDataSourceResponseBodyData) GetNetworklink() *string {
	return s.Networklink
}

func (s *GetMmsDataSourceResponseBodyData) GetPartitionNum() *int32 {
	return s.PartitionNum
}

func (s *GetMmsDataSourceResponseBodyData) GetPartitionsDoingNum() *int32 {
	return s.PartitionsDoingNum
}

func (s *GetMmsDataSourceResponseBodyData) GetPartitionsDoneNum() *int32 {
	return s.PartitionsDoneNum
}

func (s *GetMmsDataSourceResponseBodyData) GetPartitionsFailedNum() *int32 {
	return s.PartitionsFailedNum
}

func (s *GetMmsDataSourceResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetMmsDataSourceResponseBodyData) GetScanId() *int64 {
	return s.ScanId
}

func (s *GetMmsDataSourceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetMmsDataSourceResponseBodyData) GetTableNum() *int32 {
	return s.TableNum
}

func (s *GetMmsDataSourceResponseBodyData) GetTablesDoingNum() *int32 {
	return s.TablesDoingNum
}

func (s *GetMmsDataSourceResponseBodyData) GetTablesDoneNum() *int32 {
	return s.TablesDoneNum
}

func (s *GetMmsDataSourceResponseBodyData) GetTablesFailedNum() *int32 {
	return s.TablesFailedNum
}

func (s *GetMmsDataSourceResponseBodyData) GetTablesPartDoneNum() *int32 {
	return s.TablesPartDoneNum
}

func (s *GetMmsDataSourceResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMmsDataSourceResponseBodyData) SetAgentIsOnline(v bool) *GetMmsDataSourceResponseBodyData {
	s.AgentIsOnline = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetConfig(v []*GetMmsDataSourceResponseBodyDataConfig) *GetMmsDataSourceResponseBodyData {
	s.Config = v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetCreateTime(v string) *GetMmsDataSourceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetDbNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.DbNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetDstProject(v string) *GetMmsDataSourceResponseBodyData {
	s.DstProject = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetDstProjects(v []*string) *GetMmsDataSourceResponseBodyData {
	s.DstProjects = v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetErrMsg(v string) *GetMmsDataSourceResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetId(v int64) *GetMmsDataSourceResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetLastUpdateTime(v string) *GetMmsDataSourceResponseBodyData {
	s.LastUpdateTime = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetName(v string) *GetMmsDataSourceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetNetworklink(v string) *GetMmsDataSourceResponseBodyData {
	s.Networklink = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetPartitionNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.PartitionNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetPartitionsDoingNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.PartitionsDoingNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetPartitionsDoneNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.PartitionsDoneNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetPartitionsFailedNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.PartitionsFailedNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetRegion(v string) *GetMmsDataSourceResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetScanId(v int64) *GetMmsDataSourceResponseBodyData {
	s.ScanId = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetStatus(v string) *GetMmsDataSourceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTableNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TableNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTablesDoingNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TablesDoingNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTablesDoneNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TablesDoneNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTablesFailedNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TablesFailedNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetTablesPartDoneNum(v int32) *GetMmsDataSourceResponseBodyData {
	s.TablesPartDoneNum = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) SetType(v string) *GetMmsDataSourceResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyData) Validate() error {
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

type GetMmsDataSourceResponseBodyDataConfig struct {
	Desc  *string   `json:"desc,omitempty" xml:"desc,omitempty"`
	Enums []*string `json:"enums,omitempty" xml:"enums,omitempty" type:"Repeated"`
	// example:
	//
	// basic_group
	Group *string `json:"group,omitempty" xml:"group,omitempty"`
	// example:
	//
	// bigquery.range.partition.migrate.type
	Key  *string `json:"key,omitempty" xml:"key,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Cluster or Partition
	PlaceHolder *string `json:"placeHolder,omitempty" xml:"placeHolder,omitempty"`
	// example:
	//
	// true
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// example:
	//
	// .keytab
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// Partition
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetMmsDataSourceResponseBodyDataConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMmsDataSourceResponseBodyDataConfig) GoString() string {
	return s.String()
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetDesc() *string {
	return s.Desc
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetEnums() []*string {
	return s.Enums
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetGroup() *string {
	return s.Group
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetKey() *string {
	return s.Key
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetName() *string {
	return s.Name
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetPlaceHolder() *string {
	return s.PlaceHolder
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetRequired() *bool {
	return s.Required
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetSubType() *string {
	return s.SubType
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetType() *string {
	return s.Type
}

func (s *GetMmsDataSourceResponseBodyDataConfig) GetValue() interface{} {
	return s.Value
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetDesc(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Desc = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetEnums(v []*string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Enums = v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetGroup(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Group = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetKey(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Key = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetName(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Name = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetPlaceHolder(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.PlaceHolder = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetRequired(v bool) *GetMmsDataSourceResponseBodyDataConfig {
	s.Required = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetSubType(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.SubType = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetType(v string) *GetMmsDataSourceResponseBodyDataConfig {
	s.Type = &v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) SetValue(v interface{}) *GetMmsDataSourceResponseBodyDataConfig {
	s.Value = v
	return s
}

func (s *GetMmsDataSourceResponseBodyDataConfig) Validate() error {
	return dara.Validate(s)
}
