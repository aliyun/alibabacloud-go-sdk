// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnMapping(v map[string]*string) *CreateMmsJobRequest
	GetColumnMapping() map[string]*string
	SetDstDbName(v string) *CreateMmsJobRequest
	GetDstDbName() *string
	SetDstSchemaName(v string) *CreateMmsJobRequest
	GetDstSchemaName() *string
	SetEnableVerification(v bool) *CreateMmsJobRequest
	GetEnableVerification() *bool
	SetEta(v string) *CreateMmsJobRequest
	GetEta() *string
	SetIncrement(v bool) *CreateMmsJobRequest
	GetIncrement() *bool
	SetName(v string) *CreateMmsJobRequest
	GetName() *string
	SetOthers(v map[string]interface{}) *CreateMmsJobRequest
	GetOthers() map[string]interface{}
	SetPartitionFilters(v map[string]*string) *CreateMmsJobRequest
	GetPartitionFilters() map[string]*string
	SetPartitions(v []*int64) *CreateMmsJobRequest
	GetPartitions() []*int64
	SetSchemaOnly(v bool) *CreateMmsJobRequest
	GetSchemaOnly() *bool
	SetSourceId(v int64) *CreateMmsJobRequest
	GetSourceId() *int64
	SetSourceName(v string) *CreateMmsJobRequest
	GetSourceName() *string
	SetSrcDbName(v string) *CreateMmsJobRequest
	GetSrcDbName() *string
	SetSrcSchemaName(v string) *CreateMmsJobRequest
	GetSrcSchemaName() *string
	SetTableBlackList(v []*string) *CreateMmsJobRequest
	GetTableBlackList() []*string
	SetTableMapping(v map[string]*string) *CreateMmsJobRequest
	GetTableMapping() map[string]*string
	SetTableWhiteList(v []*string) *CreateMmsJobRequest
	GetTableWhiteList() []*string
	SetTables(v []*string) *CreateMmsJobRequest
	GetTables() []*string
	SetTaskType(v string) *CreateMmsJobRequest
	GetTaskType() *string
}

type CreateMmsJobRequest struct {
	ColumnMapping      map[string]*string     `json:"columnMapping,omitempty" xml:"columnMapping,omitempty"`
	DstDbName          *string                `json:"dstDbName,omitempty" xml:"dstDbName,omitempty"`
	DstSchemaName      *string                `json:"dstSchemaName,omitempty" xml:"dstSchemaName,omitempty"`
	EnableVerification *bool                  `json:"enableVerification,omitempty" xml:"enableVerification,omitempty"`
	Eta                *string                `json:"eta,omitempty" xml:"eta,omitempty"`
	Increment          *bool                  `json:"increment,omitempty" xml:"increment,omitempty"`
	Name               *string                `json:"name,omitempty" xml:"name,omitempty"`
	Others             map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
	PartitionFilters   map[string]*string     `json:"partitionFilters,omitempty" xml:"partitionFilters,omitempty"`
	Partitions         []*int64               `json:"partitions,omitempty" xml:"partitions,omitempty" type:"Repeated"`
	SchemaOnly         *bool                  `json:"schemaOnly,omitempty" xml:"schemaOnly,omitempty"`
	SourceId           *int64                 `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	SourceName         *string                `json:"sourceName,omitempty" xml:"sourceName,omitempty"`
	SrcDbName          *string                `json:"srcDbName,omitempty" xml:"srcDbName,omitempty"`
	SrcSchemaName      *string                `json:"srcSchemaName,omitempty" xml:"srcSchemaName,omitempty"`
	TableBlackList     []*string              `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	TableMapping       map[string]*string     `json:"tableMapping,omitempty" xml:"tableMapping,omitempty"`
	TableWhiteList     []*string              `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	Tables             []*string              `json:"tables,omitempty" xml:"tables,omitempty" type:"Repeated"`
	// MOCK, HIVE: hive udtf task, HIVE_DATAX: hive datax task, COPY_TASK: odps Copy Task, ODPS_INSERT_OVERWRITE: odps simple insert overwrite task, MC2MC_VERIFY, OSS, HIVE_OSS
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s CreateMmsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMmsJobRequest) GetColumnMapping() map[string]*string {
	return s.ColumnMapping
}

func (s *CreateMmsJobRequest) GetDstDbName() *string {
	return s.DstDbName
}

func (s *CreateMmsJobRequest) GetDstSchemaName() *string {
	return s.DstSchemaName
}

func (s *CreateMmsJobRequest) GetEnableVerification() *bool {
	return s.EnableVerification
}

func (s *CreateMmsJobRequest) GetEta() *string {
	return s.Eta
}

func (s *CreateMmsJobRequest) GetIncrement() *bool {
	return s.Increment
}

func (s *CreateMmsJobRequest) GetName() *string {
	return s.Name
}

func (s *CreateMmsJobRequest) GetOthers() map[string]interface{} {
	return s.Others
}

func (s *CreateMmsJobRequest) GetPartitionFilters() map[string]*string {
	return s.PartitionFilters
}

func (s *CreateMmsJobRequest) GetPartitions() []*int64 {
	return s.Partitions
}

func (s *CreateMmsJobRequest) GetSchemaOnly() *bool {
	return s.SchemaOnly
}

func (s *CreateMmsJobRequest) GetSourceId() *int64 {
	return s.SourceId
}

func (s *CreateMmsJobRequest) GetSourceName() *string {
	return s.SourceName
}

func (s *CreateMmsJobRequest) GetSrcDbName() *string {
	return s.SrcDbName
}

func (s *CreateMmsJobRequest) GetSrcSchemaName() *string {
	return s.SrcSchemaName
}

func (s *CreateMmsJobRequest) GetTableBlackList() []*string {
	return s.TableBlackList
}

func (s *CreateMmsJobRequest) GetTableMapping() map[string]*string {
	return s.TableMapping
}

func (s *CreateMmsJobRequest) GetTableWhiteList() []*string {
	return s.TableWhiteList
}

func (s *CreateMmsJobRequest) GetTables() []*string {
	return s.Tables
}

func (s *CreateMmsJobRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateMmsJobRequest) SetColumnMapping(v map[string]*string) *CreateMmsJobRequest {
	s.ColumnMapping = v
	return s
}

func (s *CreateMmsJobRequest) SetDstDbName(v string) *CreateMmsJobRequest {
	s.DstDbName = &v
	return s
}

func (s *CreateMmsJobRequest) SetDstSchemaName(v string) *CreateMmsJobRequest {
	s.DstSchemaName = &v
	return s
}

func (s *CreateMmsJobRequest) SetEnableVerification(v bool) *CreateMmsJobRequest {
	s.EnableVerification = &v
	return s
}

func (s *CreateMmsJobRequest) SetEta(v string) *CreateMmsJobRequest {
	s.Eta = &v
	return s
}

func (s *CreateMmsJobRequest) SetIncrement(v bool) *CreateMmsJobRequest {
	s.Increment = &v
	return s
}

func (s *CreateMmsJobRequest) SetName(v string) *CreateMmsJobRequest {
	s.Name = &v
	return s
}

func (s *CreateMmsJobRequest) SetOthers(v map[string]interface{}) *CreateMmsJobRequest {
	s.Others = v
	return s
}

func (s *CreateMmsJobRequest) SetPartitionFilters(v map[string]*string) *CreateMmsJobRequest {
	s.PartitionFilters = v
	return s
}

func (s *CreateMmsJobRequest) SetPartitions(v []*int64) *CreateMmsJobRequest {
	s.Partitions = v
	return s
}

func (s *CreateMmsJobRequest) SetSchemaOnly(v bool) *CreateMmsJobRequest {
	s.SchemaOnly = &v
	return s
}

func (s *CreateMmsJobRequest) SetSourceId(v int64) *CreateMmsJobRequest {
	s.SourceId = &v
	return s
}

func (s *CreateMmsJobRequest) SetSourceName(v string) *CreateMmsJobRequest {
	s.SourceName = &v
	return s
}

func (s *CreateMmsJobRequest) SetSrcDbName(v string) *CreateMmsJobRequest {
	s.SrcDbName = &v
	return s
}

func (s *CreateMmsJobRequest) SetSrcSchemaName(v string) *CreateMmsJobRequest {
	s.SrcSchemaName = &v
	return s
}

func (s *CreateMmsJobRequest) SetTableBlackList(v []*string) *CreateMmsJobRequest {
	s.TableBlackList = v
	return s
}

func (s *CreateMmsJobRequest) SetTableMapping(v map[string]*string) *CreateMmsJobRequest {
	s.TableMapping = v
	return s
}

func (s *CreateMmsJobRequest) SetTableWhiteList(v []*string) *CreateMmsJobRequest {
	s.TableWhiteList = v
	return s
}

func (s *CreateMmsJobRequest) SetTables(v []*string) *CreateMmsJobRequest {
	s.Tables = v
	return s
}

func (s *CreateMmsJobRequest) SetTaskType(v string) *CreateMmsJobRequest {
	s.TaskType = &v
	return s
}

func (s *CreateMmsJobRequest) Validate() error {
	return dara.Validate(s)
}
