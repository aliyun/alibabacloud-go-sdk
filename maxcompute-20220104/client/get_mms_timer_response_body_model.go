// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMmsTimerResponseBodyData) *GetMmsTimerResponseBody
	GetData() *GetMmsTimerResponseBodyData
	SetRequestId(v string) *GetMmsTimerResponseBody
	GetRequestId() *string
}

type GetMmsTimerResponseBody struct {
	Data *GetMmsTimerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 0a06dfe716674588654372173ec0da
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetMmsTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTimerResponseBody) GoString() string {
	return s.String()
}

func (s *GetMmsTimerResponseBody) GetData() *GetMmsTimerResponseBodyData {
	return s.Data
}

func (s *GetMmsTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMmsTimerResponseBody) SetData(v *GetMmsTimerResponseBodyData) *GetMmsTimerResponseBody {
	s.Data = v
	return s
}

func (s *GetMmsTimerResponseBody) SetRequestId(v string) *GetMmsTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMmsTimerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsTimerResponseBodyData struct {
	Config *GetMmsTimerResponseBodyDataConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 1730946421757
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 23
	DbId *int64 `json:"dbId,omitempty" xml:"dbId,omitempty"`
	// example:
	//
	// 2523
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// sale_detail
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Daily
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// example:
	//
	// 2000017
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// mms_test
	SrcDbName *string `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	// example:
	//
	// false
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// example:
	//
	// Daily, Hourly
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 2026-04-01T02:18:01Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// p1=1/p2=abc
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetMmsTimerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTimerResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMmsTimerResponseBodyData) GetConfig() *GetMmsTimerResponseBodyDataConfig {
	return s.Config
}

func (s *GetMmsTimerResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetMmsTimerResponseBodyData) GetDbId() *int64 {
	return s.DbId
}

func (s *GetMmsTimerResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetMmsTimerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetMmsTimerResponseBodyData) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *GetMmsTimerResponseBodyData) GetSourceId() *int64 {
	return s.SourceId
}

func (s *GetMmsTimerResponseBodyData) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *GetMmsTimerResponseBodyData) GetStopped() *bool {
	return s.Stopped
}

func (s *GetMmsTimerResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMmsTimerResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetMmsTimerResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *GetMmsTimerResponseBodyData) SetConfig(v *GetMmsTimerResponseBodyDataConfig) *GetMmsTimerResponseBodyData {
	s.Config = v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetCreateTime(v string) *GetMmsTimerResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetDbId(v int64) *GetMmsTimerResponseBodyData {
	s.DbId = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetId(v int64) *GetMmsTimerResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetName(v string) *GetMmsTimerResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetScheduleType(v string) *GetMmsTimerResponseBodyData {
	s.ScheduleType = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetSourceId(v int64) *GetMmsTimerResponseBodyData {
	s.SourceId = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetSrcDbName(v string) *GetMmsTimerResponseBodyData {
	s.SrcDbName = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetStopped(v bool) *GetMmsTimerResponseBodyData {
	s.Stopped = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetType(v string) *GetMmsTimerResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetUpdateTime(v string) *GetMmsTimerResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) SetValue(v string) *GetMmsTimerResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetMmsTimerResponseBodyData) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMmsTimerResponseBodyDataConfig struct {
	ColumnMapping         map[string]*string `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	EnableDataMigration   *bool              `json:"enableDataMigration,omitempty" xml:"enableDataMigration,omitempty"`
	EnableSchemaMigration *bool              `json:"enableSchemaMigration,omitempty" xml:"enableSchemaMigration,omitempty"`
	// example:
	//
	// true
	EnableVerification *bool                  `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	Others             map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	PartitionFilters   map[string]*string     `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	Partitions         []*int64               `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	TableBlackList     []*string              `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	TableMapping       map[string]*string     `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	TableWhiteList     []*string              `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	Tables             []*string              `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
}

func (s GetMmsTimerResponseBodyDataConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTimerResponseBodyDataConfig) GoString() string {
	return s.String()
}

func (s *GetMmsTimerResponseBodyDataConfig) GetColumnMapping() map[string]*string {
	return s.ColumnMapping
}

func (s *GetMmsTimerResponseBodyDataConfig) GetEnableDataMigration() *bool {
	return s.EnableDataMigration
}

func (s *GetMmsTimerResponseBodyDataConfig) GetEnableSchemaMigration() *bool {
	return s.EnableSchemaMigration
}

func (s *GetMmsTimerResponseBodyDataConfig) GetEnableVerification() *bool {
	return s.EnableVerification
}

func (s *GetMmsTimerResponseBodyDataConfig) GetOthers() map[string]interface{} {
	return s.Others
}

func (s *GetMmsTimerResponseBodyDataConfig) GetPartitionFilters() map[string]*string {
	return s.PartitionFilters
}

func (s *GetMmsTimerResponseBodyDataConfig) GetPartitions() []*int64 {
	return s.Partitions
}

func (s *GetMmsTimerResponseBodyDataConfig) GetTableBlackList() []*string {
	return s.TableBlackList
}

func (s *GetMmsTimerResponseBodyDataConfig) GetTableMapping() map[string]*string {
	return s.TableMapping
}

func (s *GetMmsTimerResponseBodyDataConfig) GetTableWhiteList() []*string {
	return s.TableWhiteList
}

func (s *GetMmsTimerResponseBodyDataConfig) GetTables() []*string {
	return s.Tables
}

func (s *GetMmsTimerResponseBodyDataConfig) SetColumnMapping(v map[string]*string) *GetMmsTimerResponseBodyDataConfig {
	s.ColumnMapping = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetEnableDataMigration(v bool) *GetMmsTimerResponseBodyDataConfig {
	s.EnableDataMigration = &v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetEnableSchemaMigration(v bool) *GetMmsTimerResponseBodyDataConfig {
	s.EnableSchemaMigration = &v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetEnableVerification(v bool) *GetMmsTimerResponseBodyDataConfig {
	s.EnableVerification = &v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetOthers(v map[string]interface{}) *GetMmsTimerResponseBodyDataConfig {
	s.Others = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetPartitionFilters(v map[string]*string) *GetMmsTimerResponseBodyDataConfig {
	s.PartitionFilters = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetPartitions(v []*int64) *GetMmsTimerResponseBodyDataConfig {
	s.Partitions = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetTableBlackList(v []*string) *GetMmsTimerResponseBodyDataConfig {
	s.TableBlackList = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetTableMapping(v map[string]*string) *GetMmsTimerResponseBodyDataConfig {
	s.TableMapping = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetTableWhiteList(v []*string) *GetMmsTimerResponseBodyDataConfig {
	s.TableWhiteList = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) SetTables(v []*string) *GetMmsTimerResponseBodyDataConfig {
	s.Tables = v
	return s
}

func (s *GetMmsTimerResponseBodyDataConfig) Validate() error {
	return dara.Validate(s)
}
