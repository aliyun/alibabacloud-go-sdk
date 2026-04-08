// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsTimerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnMapping(v map[string]*string) *CreateMmsTimerRequest
	GetColumnMapping() map[string]*string
	SetEnableDataMigration(v bool) *CreateMmsTimerRequest
	GetEnableDataMigration() *bool
	SetEnableSchemaMigration(v bool) *CreateMmsTimerRequest
	GetEnableSchemaMigration() *bool
	SetEnableVerification(v bool) *CreateMmsTimerRequest
	GetEnableVerification() *bool
	SetName(v string) *CreateMmsTimerRequest
	GetName() *string
	SetOthers(v map[string]interface{}) *CreateMmsTimerRequest
	GetOthers() map[string]interface{}
	SetPartitionFilters(v map[string]*string) *CreateMmsTimerRequest
	GetPartitionFilters() map[string]*string
	SetPartitions(v []*int64) *CreateMmsTimerRequest
	GetPartitions() []*int64
	SetScheduleType(v string) *CreateMmsTimerRequest
	GetScheduleType() *string
	SetSourceId(v int64) *CreateMmsTimerRequest
	GetSourceId() *int64
	SetSrcDbName(v string) *CreateMmsTimerRequest
	GetSrcDbName() *string
	SetTableBlackList(v []*string) *CreateMmsTimerRequest
	GetTableBlackList() []*string
	SetTableMapping(v map[string]*string) *CreateMmsTimerRequest
	GetTableMapping() map[string]*string
	SetTableWhiteList(v []*string) *CreateMmsTimerRequest
	GetTableWhiteList() []*string
	SetTables(v []*string) *CreateMmsTimerRequest
	GetTables() []*string
	SetValue(v string) *CreateMmsTimerRequest
	GetValue() *string
}

type CreateMmsTimerRequest struct {
	ColumnMapping map[string]*string `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	// example:
	//
	// true
	EnableDataMigration *bool `json:"enableDataMigration,omitempty" xml:"enableDataMigration,omitempty"`
	// example:
	//
	// true
	EnableSchemaMigration *bool `json:"enableSchemaMigration,omitempty" xml:"enableSchemaMigration,omitempty"`
	// example:
	//
	// false
	EnableVerification *bool `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	// example:
	//
	// planA
	Name             *string                `json:"name,omitempty" xml:"name,omitempty"`
	Others           map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	PartitionFilters map[string]*string     `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	Partitions       []*int64               `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	// example:
	//
	// Daily
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// example:
	//
	// 2000014
	SourceId *int64 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// example:
	//
	// src_db
	SrcDbName      *string            `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	TableBlackList []*string          `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	TableMapping   map[string]*string `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	TableWhiteList []*string          `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	Tables         []*string          `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	// example:
	//
	// 12:00
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateMmsTimerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsTimerRequest) GoString() string {
	return s.String()
}

func (s *CreateMmsTimerRequest) GetColumnMapping() map[string]*string {
	return s.ColumnMapping
}

func (s *CreateMmsTimerRequest) GetEnableDataMigration() *bool {
	return s.EnableDataMigration
}

func (s *CreateMmsTimerRequest) GetEnableSchemaMigration() *bool {
	return s.EnableSchemaMigration
}

func (s *CreateMmsTimerRequest) GetEnableVerification() *bool {
	return s.EnableVerification
}

func (s *CreateMmsTimerRequest) GetName() *string {
	return s.Name
}

func (s *CreateMmsTimerRequest) GetOthers() map[string]interface{} {
	return s.Others
}

func (s *CreateMmsTimerRequest) GetPartitionFilters() map[string]*string {
	return s.PartitionFilters
}

func (s *CreateMmsTimerRequest) GetPartitions() []*int64 {
	return s.Partitions
}

func (s *CreateMmsTimerRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateMmsTimerRequest) GetSourceId() *int64 {
	return s.SourceId
}

func (s *CreateMmsTimerRequest) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *CreateMmsTimerRequest) GetTableBlackList() []*string {
	return s.TableBlackList
}

func (s *CreateMmsTimerRequest) GetTableMapping() map[string]*string {
	return s.TableMapping
}

func (s *CreateMmsTimerRequest) GetTableWhiteList() []*string {
	return s.TableWhiteList
}

func (s *CreateMmsTimerRequest) GetTables() []*string {
	return s.Tables
}

func (s *CreateMmsTimerRequest) GetValue() *string {
	return s.Value
}

func (s *CreateMmsTimerRequest) SetColumnMapping(v map[string]*string) *CreateMmsTimerRequest {
	s.ColumnMapping = v
	return s
}

func (s *CreateMmsTimerRequest) SetEnableDataMigration(v bool) *CreateMmsTimerRequest {
	s.EnableDataMigration = &v
	return s
}

func (s *CreateMmsTimerRequest) SetEnableSchemaMigration(v bool) *CreateMmsTimerRequest {
	s.EnableSchemaMigration = &v
	return s
}

func (s *CreateMmsTimerRequest) SetEnableVerification(v bool) *CreateMmsTimerRequest {
	s.EnableVerification = &v
	return s
}

func (s *CreateMmsTimerRequest) SetName(v string) *CreateMmsTimerRequest {
	s.Name = &v
	return s
}

func (s *CreateMmsTimerRequest) SetOthers(v map[string]interface{}) *CreateMmsTimerRequest {
	s.Others = v
	return s
}

func (s *CreateMmsTimerRequest) SetPartitionFilters(v map[string]*string) *CreateMmsTimerRequest {
	s.PartitionFilters = v
	return s
}

func (s *CreateMmsTimerRequest) SetPartitions(v []*int64) *CreateMmsTimerRequest {
	s.Partitions = v
	return s
}

func (s *CreateMmsTimerRequest) SetScheduleType(v string) *CreateMmsTimerRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateMmsTimerRequest) SetSourceId(v int64) *CreateMmsTimerRequest {
	s.SourceId = &v
	return s
}

func (s *CreateMmsTimerRequest) SetSrcDbName(v string) *CreateMmsTimerRequest {
	s.SrcDbName = &v
	return s
}

func (s *CreateMmsTimerRequest) SetTableBlackList(v []*string) *CreateMmsTimerRequest {
	s.TableBlackList = v
	return s
}

func (s *CreateMmsTimerRequest) SetTableMapping(v map[string]*string) *CreateMmsTimerRequest {
	s.TableMapping = v
	return s
}

func (s *CreateMmsTimerRequest) SetTableWhiteList(v []*string) *CreateMmsTimerRequest {
	s.TableWhiteList = v
	return s
}

func (s *CreateMmsTimerRequest) SetTables(v []*string) *CreateMmsTimerRequest {
	s.Tables = v
	return s
}

func (s *CreateMmsTimerRequest) SetValue(v string) *CreateMmsTimerRequest {
	s.Value = &v
	return s
}

func (s *CreateMmsTimerRequest) Validate() error {
	return dara.Validate(s)
}
