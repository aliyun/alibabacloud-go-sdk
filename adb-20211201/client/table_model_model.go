// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTableModel interface {
	dara.Model
	String() string
	GoString() string
	SetArchiveType(v string) *TableModel
	GetArchiveType() *string
	SetBlockSize(v int64) *TableModel
	GetBlockSize() *int64
	SetBucket(v int64) *TableModel
	GetBucket() *int64
	SetBucketCount(v int64) *TableModel
	GetBucketCount() *int64
	SetCols(v []*FieldSchemaModel) *TableModel
	GetCols() []*FieldSchemaModel
	SetComment(v string) *TableModel
	GetComment() *string
	SetCompression(v string) *TableModel
	GetCompression() *string
	SetCreateTime(v string) *TableModel
	GetCreateTime() *string
	SetCreatedBySource(v string) *TableModel
	GetCreatedBySource() *string
	SetCreatedByUser(v string) *TableModel
	GetCreatedByUser() *string
	SetCurrentVersion(v int64) *TableModel
	GetCurrentVersion() *int64
	SetDbName(v string) *TableModel
	GetDbName() *string
	SetDictEncode(v bool) *TableModel
	GetDictEncode() *bool
	SetDistributeColumns(v []*FieldSchemaModel) *TableModel
	GetDistributeColumns() []*FieldSchemaModel
	SetDistributeType(v string) *TableModel
	GetDistributeType() *string
	SetEnableDfs(v bool) *TableModel
	GetEnableDfs() *bool
	SetHotPartitionCount(v int64) *TableModel
	GetHotPartitionCount() *int64
	SetIndexes(v []*CstoreIndexModel) *TableModel
	GetIndexes() []*CstoreIndexModel
	SetIsAllIndex(v bool) *TableModel
	GetIsAllIndex() *bool
	SetIsFulltextDict(v bool) *TableModel
	GetIsFulltextDict() *bool
	SetMaxColumnId(v int64) *TableModel
	GetMaxColumnId() *int64
	SetParameters(v map[string]*string) *TableModel
	GetParameters() map[string]*string
	SetPartitionColumn(v string) *TableModel
	GetPartitionColumn() *string
	SetPartitionCount(v int64) *TableModel
	GetPartitionCount() *int64
	SetPartitionKeys(v []*FieldSchemaModel) *TableModel
	GetPartitionKeys() []*FieldSchemaModel
	SetPartitionType(v string) *TableModel
	GetPartitionType() *string
	SetPhysicalDatabaseName(v string) *TableModel
	GetPhysicalDatabaseName() *string
	SetPhysicalTableName(v string) *TableModel
	GetPhysicalTableName() *string
	SetPreviousVersion(v int64) *TableModel
	GetPreviousVersion() *int64
	SetRawTableName(v string) *TableModel
	GetRawTableName() *string
	SetRouteColumns(v []*FieldSchemaModel) *TableModel
	GetRouteColumns() []*FieldSchemaModel
	SetRouteEffectiveColumn(v *FieldSchemaModel) *TableModel
	GetRouteEffectiveColumn() *FieldSchemaModel
	SetRouteType(v string) *TableModel
	GetRouteType() *string
	SetRtEngineType(v string) *TableModel
	GetRtEngineType() *string
	SetRtIndexAll(v bool) *TableModel
	GetRtIndexAll() *bool
	SetRtModeType(v string) *TableModel
	GetRtModeType() *string
	SetSd(v *StorageDescriptorModel) *TableModel
	GetSd() *StorageDescriptorModel
	SetStoragePolicy(v string) *TableModel
	GetStoragePolicy() *string
	SetSubpartitionColumn(v string) *TableModel
	GetSubpartitionColumn() *string
	SetSubpartitionCount(v int64) *TableModel
	GetSubpartitionCount() *int64
	SetSubpartitionType(v string) *TableModel
	GetSubpartitionType() *string
	SetTableEngineName(v string) *TableModel
	GetTableEngineName() *string
	SetTableName(v string) *TableModel
	GetTableName() *string
	SetTableType(v string) *TableModel
	GetTableType() *string
	SetTblId(v int64) *TableModel
	GetTblId() *int64
	SetTemporary(v bool) *TableModel
	GetTemporary() *bool
	SetUpdateTime(v string) *TableModel
	GetUpdateTime() *string
	SetViewExpandedText(v string) *TableModel
	GetViewExpandedText() *string
	SetViewOriginalText(v string) *TableModel
	GetViewOriginalText() *string
	SetViewSecurityMode(v string) *TableModel
	GetViewSecurityMode() *string
}

type TableModel struct {
	// The archive type.
	//
	// example:
	//
	// ArchiveType
	ArchiveType *string `json:"ArchiveType,omitempty" xml:"ArchiveType,omitempty"`
	// The block size.
	//
	// example:
	//
	// 64
	BlockSize *int64 `json:"BlockSize,omitempty" xml:"BlockSize,omitempty"`
	// The bucket ID.
	//
	// example:
	//
	// 16
	Bucket *int64 `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The number of buckets.
	//
	// example:
	//
	// 16
	BucketCount *int64 `json:"BucketCount,omitempty" xml:"BucketCount,omitempty"`
	// The column information.
	Cols []*FieldSchemaModel `json:"Cols,omitempty" xml:"Cols,omitempty" type:"Repeated"`
	// The description.
	//
	// example:
	//
	// description
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The compression method.
	//
	// example:
	//
	// Compression
	Compression *string `json:"Compression,omitempty" xml:"Compression,omitempty"`
	// The time when the table was created.
	//
	// example:
	//
	// 2023-01-05 13:17:55
	CreateTime      *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBySource *string `json:"CreatedBySource,omitempty" xml:"CreatedBySource,omitempty"`
	CreatedByUser   *string `json:"CreatedByUser,omitempty" xml:"CreatedByUser,omitempty"`
	// The current version.
	//
	// example:
	//
	// 2
	CurrentVersion *int64 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The name of the logical database.
	//
	// example:
	//
	// example
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates whether the dictionary is encrypted.
	//
	// example:
	//
	// false
	DictEncode *bool `json:"DictEncode,omitempty" xml:"DictEncode,omitempty"`
	// The distribution columns.
	DistributeColumns []*FieldSchemaModel `json:"DistributeColumns,omitempty" xml:"DistributeColumns,omitempty" type:"Repeated"`
	// The distribution type.
	//
	// example:
	//
	// DistributeType
	DistributeType *string `json:"DistributeType,omitempty" xml:"DistributeType,omitempty"`
	// Indicates whether DFS is allowed.
	//
	// example:
	//
	// false
	EnableDfs *bool `json:"EnableDfs,omitempty" xml:"EnableDfs,omitempty"`
	// The number of hot partitions.
	//
	// example:
	//
	// 32
	HotPartitionCount *int64 `json:"HotPartitionCount,omitempty" xml:"HotPartitionCount,omitempty"`
	// The indexes.
	Indexes []*CstoreIndexModel `json:"Indexes,omitempty" xml:"Indexes,omitempty" type:"Repeated"`
	// Indicates whether the index is a full index.
	//
	// example:
	//
	// true
	IsAllIndex *bool `json:"IsAllIndex,omitempty" xml:"IsAllIndex,omitempty"`
	// Indicates whether the table is a full-text index dictionary.
	//
	// example:
	//
	// false
	IsFulltextDict *bool `json:"IsFulltextDict,omitempty" xml:"IsFulltextDict,omitempty"`
	// The maximum column ID.
	//
	// example:
	//
	// MaxColumnId
	MaxColumnId *int64 `json:"MaxColumnId,omitempty" xml:"MaxColumnId,omitempty"`
	// The parameters.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The information about the partition key column.
	//
	// example:
	//
	// colName
	PartitionColumn *string `json:"PartitionColumn,omitempty" xml:"PartitionColumn,omitempty"`
	// The number of partitions.
	//
	// example:
	//
	// 16
	PartitionCount *int64 `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	// The partition keys.
	PartitionKeys []*FieldSchemaModel `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	// The partition type.
	//
	// example:
	//
	// PartitionType
	PartitionType *string `json:"PartitionType,omitempty" xml:"PartitionType,omitempty"`
	// The name of the physical database.
	//
	// example:
	//
	// physicalDatabaseName
	PhysicalDatabaseName *string `json:"PhysicalDatabaseName,omitempty" xml:"PhysicalDatabaseName,omitempty"`
	// The name of the physical table.
	//
	// example:
	//
	// physicalTableName
	PhysicalTableName *string `json:"PhysicalTableName,omitempty" xml:"PhysicalTableName,omitempty"`
	// The previous version.
	//
	// example:
	//
	// 1
	PreviousVersion *int64 `json:"PreviousVersion,omitempty" xml:"PreviousVersion,omitempty"`
	// The raw table name.
	//
	// example:
	//
	// RawTableName
	RawTableName *string `json:"RawTableName,omitempty" xml:"RawTableName,omitempty"`
	// The routing columns.
	RouteColumns []*FieldSchemaModel `json:"RouteColumns,omitempty" xml:"RouteColumns,omitempty" type:"Repeated"`
	// The effective routing column.
	RouteEffectiveColumn *FieldSchemaModel `json:"RouteEffectiveColumn,omitempty" xml:"RouteEffectiveColumn,omitempty"`
	// The routing type.
	//
	// example:
	//
	// routeType
	RouteType *string `json:"RouteType,omitempty" xml:"RouteType,omitempty"`
	// The routing engine type.
	//
	// example:
	//
	// RtEngineType
	RtEngineType *string `json:"RtEngineType,omitempty" xml:"RtEngineType,omitempty"`
	// Indicates whether to route all indexes.
	//
	// example:
	//
	// false
	RtIndexAll *bool `json:"RtIndexAll,omitempty" xml:"RtIndexAll,omitempty"`
	// The routing mode type.
	//
	// example:
	//
	// RtModeType
	RtModeType *string `json:"RtModeType,omitempty" xml:"RtModeType,omitempty"`
	// The description of the storage.
	Sd *StorageDescriptorModel `json:"Sd,omitempty" xml:"Sd,omitempty"`
	// The storage policy.
	//
	// example:
	//
	// StoragePolicy
	StoragePolicy *string `json:"StoragePolicy,omitempty" xml:"StoragePolicy,omitempty"`
	// The information about the subpartition column.
	//
	// example:
	//
	// SubpartitionColumn
	SubpartitionColumn *string `json:"SubpartitionColumn,omitempty" xml:"SubpartitionColumn,omitempty"`
	// The number of subpartitions.
	//
	// example:
	//
	// 64
	SubpartitionCount *int64 `json:"SubpartitionCount,omitempty" xml:"SubpartitionCount,omitempty"`
	// The subpartition type.
	//
	// example:
	//
	// SubpartitionColumn
	SubpartitionType *string `json:"SubpartitionType,omitempty" xml:"SubpartitionType,omitempty"`
	// The name of the table engine.
	//
	// example:
	//
	// hive
	TableEngineName *string `json:"TableEngineName,omitempty" xml:"TableEngineName,omitempty"`
	// The name of the logical table.
	//
	// example:
	//
	// tableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The table type.
	//
	// example:
	//
	// external_table
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	// The table ID.
	//
	// example:
	//
	// 123
	TblId *int64 `json:"TblId,omitempty" xml:"TblId,omitempty"`
	// Indicates whether the table is a temporary table.
	//
	// example:
	//
	// false
	Temporary *bool `json:"Temporary,omitempty" xml:"Temporary,omitempty"`
	// The time when the table was last updated.
	//
	// example:
	//
	// 2023-01-05 13:17:55
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The normalized SQL statement that is used to create the view.
	//
	// example:
	//
	// ViewExpandedText
	ViewExpandedText *string `json:"ViewExpandedText,omitempty" xml:"ViewExpandedText,omitempty"`
	// The SQL statement used to create the view.
	//
	// example:
	//
	// ViewOriginalText
	ViewOriginalText *string `json:"ViewOriginalText,omitempty" xml:"ViewOriginalText,omitempty"`
	// The security mode of the view.
	//
	// example:
	//
	// ViewSecurityMode
	ViewSecurityMode *string `json:"ViewSecurityMode,omitempty" xml:"ViewSecurityMode,omitempty"`
}

func (s TableModel) String() string {
	return dara.Prettify(s)
}

func (s TableModel) GoString() string {
	return s.String()
}

func (s *TableModel) GetArchiveType() *string {
	return s.ArchiveType
}

func (s *TableModel) GetBlockSize() *int64 {
	return s.BlockSize
}

func (s *TableModel) GetBucket() *int64 {
	return s.Bucket
}

func (s *TableModel) GetBucketCount() *int64 {
	return s.BucketCount
}

func (s *TableModel) GetCols() []*FieldSchemaModel {
	return s.Cols
}

func (s *TableModel) GetComment() *string {
	return s.Comment
}

func (s *TableModel) GetCompression() *string {
	return s.Compression
}

func (s *TableModel) GetCreateTime() *string {
	return s.CreateTime
}

func (s *TableModel) GetCreatedBySource() *string {
	return s.CreatedBySource
}

func (s *TableModel) GetCreatedByUser() *string {
	return s.CreatedByUser
}

func (s *TableModel) GetCurrentVersion() *int64 {
	return s.CurrentVersion
}

func (s *TableModel) GetDbName() *string {
	return s.DbName
}

func (s *TableModel) GetDictEncode() *bool {
	return s.DictEncode
}

func (s *TableModel) GetDistributeColumns() []*FieldSchemaModel {
	return s.DistributeColumns
}

func (s *TableModel) GetDistributeType() *string {
	return s.DistributeType
}

func (s *TableModel) GetEnableDfs() *bool {
	return s.EnableDfs
}

func (s *TableModel) GetHotPartitionCount() *int64 {
	return s.HotPartitionCount
}

func (s *TableModel) GetIndexes() []*CstoreIndexModel {
	return s.Indexes
}

func (s *TableModel) GetIsAllIndex() *bool {
	return s.IsAllIndex
}

func (s *TableModel) GetIsFulltextDict() *bool {
	return s.IsFulltextDict
}

func (s *TableModel) GetMaxColumnId() *int64 {
	return s.MaxColumnId
}

func (s *TableModel) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *TableModel) GetPartitionColumn() *string {
	return s.PartitionColumn
}

func (s *TableModel) GetPartitionCount() *int64 {
	return s.PartitionCount
}

func (s *TableModel) GetPartitionKeys() []*FieldSchemaModel {
	return s.PartitionKeys
}

func (s *TableModel) GetPartitionType() *string {
	return s.PartitionType
}

func (s *TableModel) GetPhysicalDatabaseName() *string {
	return s.PhysicalDatabaseName
}

func (s *TableModel) GetPhysicalTableName() *string {
	return s.PhysicalTableName
}

func (s *TableModel) GetPreviousVersion() *int64 {
	return s.PreviousVersion
}

func (s *TableModel) GetRawTableName() *string {
	return s.RawTableName
}

func (s *TableModel) GetRouteColumns() []*FieldSchemaModel {
	return s.RouteColumns
}

func (s *TableModel) GetRouteEffectiveColumn() *FieldSchemaModel {
	return s.RouteEffectiveColumn
}

func (s *TableModel) GetRouteType() *string {
	return s.RouteType
}

func (s *TableModel) GetRtEngineType() *string {
	return s.RtEngineType
}

func (s *TableModel) GetRtIndexAll() *bool {
	return s.RtIndexAll
}

func (s *TableModel) GetRtModeType() *string {
	return s.RtModeType
}

func (s *TableModel) GetSd() *StorageDescriptorModel {
	return s.Sd
}

func (s *TableModel) GetStoragePolicy() *string {
	return s.StoragePolicy
}

func (s *TableModel) GetSubpartitionColumn() *string {
	return s.SubpartitionColumn
}

func (s *TableModel) GetSubpartitionCount() *int64 {
	return s.SubpartitionCount
}

func (s *TableModel) GetSubpartitionType() *string {
	return s.SubpartitionType
}

func (s *TableModel) GetTableEngineName() *string {
	return s.TableEngineName
}

func (s *TableModel) GetTableName() *string {
	return s.TableName
}

func (s *TableModel) GetTableType() *string {
	return s.TableType
}

func (s *TableModel) GetTblId() *int64 {
	return s.TblId
}

func (s *TableModel) GetTemporary() *bool {
	return s.Temporary
}

func (s *TableModel) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *TableModel) GetViewExpandedText() *string {
	return s.ViewExpandedText
}

func (s *TableModel) GetViewOriginalText() *string {
	return s.ViewOriginalText
}

func (s *TableModel) GetViewSecurityMode() *string {
	return s.ViewSecurityMode
}

func (s *TableModel) SetArchiveType(v string) *TableModel {
	s.ArchiveType = &v
	return s
}

func (s *TableModel) SetBlockSize(v int64) *TableModel {
	s.BlockSize = &v
	return s
}

func (s *TableModel) SetBucket(v int64) *TableModel {
	s.Bucket = &v
	return s
}

func (s *TableModel) SetBucketCount(v int64) *TableModel {
	s.BucketCount = &v
	return s
}

func (s *TableModel) SetCols(v []*FieldSchemaModel) *TableModel {
	s.Cols = v
	return s
}

func (s *TableModel) SetComment(v string) *TableModel {
	s.Comment = &v
	return s
}

func (s *TableModel) SetCompression(v string) *TableModel {
	s.Compression = &v
	return s
}

func (s *TableModel) SetCreateTime(v string) *TableModel {
	s.CreateTime = &v
	return s
}

func (s *TableModel) SetCreatedBySource(v string) *TableModel {
	s.CreatedBySource = &v
	return s
}

func (s *TableModel) SetCreatedByUser(v string) *TableModel {
	s.CreatedByUser = &v
	return s
}

func (s *TableModel) SetCurrentVersion(v int64) *TableModel {
	s.CurrentVersion = &v
	return s
}

func (s *TableModel) SetDbName(v string) *TableModel {
	s.DbName = &v
	return s
}

func (s *TableModel) SetDictEncode(v bool) *TableModel {
	s.DictEncode = &v
	return s
}

func (s *TableModel) SetDistributeColumns(v []*FieldSchemaModel) *TableModel {
	s.DistributeColumns = v
	return s
}

func (s *TableModel) SetDistributeType(v string) *TableModel {
	s.DistributeType = &v
	return s
}

func (s *TableModel) SetEnableDfs(v bool) *TableModel {
	s.EnableDfs = &v
	return s
}

func (s *TableModel) SetHotPartitionCount(v int64) *TableModel {
	s.HotPartitionCount = &v
	return s
}

func (s *TableModel) SetIndexes(v []*CstoreIndexModel) *TableModel {
	s.Indexes = v
	return s
}

func (s *TableModel) SetIsAllIndex(v bool) *TableModel {
	s.IsAllIndex = &v
	return s
}

func (s *TableModel) SetIsFulltextDict(v bool) *TableModel {
	s.IsFulltextDict = &v
	return s
}

func (s *TableModel) SetMaxColumnId(v int64) *TableModel {
	s.MaxColumnId = &v
	return s
}

func (s *TableModel) SetParameters(v map[string]*string) *TableModel {
	s.Parameters = v
	return s
}

func (s *TableModel) SetPartitionColumn(v string) *TableModel {
	s.PartitionColumn = &v
	return s
}

func (s *TableModel) SetPartitionCount(v int64) *TableModel {
	s.PartitionCount = &v
	return s
}

func (s *TableModel) SetPartitionKeys(v []*FieldSchemaModel) *TableModel {
	s.PartitionKeys = v
	return s
}

func (s *TableModel) SetPartitionType(v string) *TableModel {
	s.PartitionType = &v
	return s
}

func (s *TableModel) SetPhysicalDatabaseName(v string) *TableModel {
	s.PhysicalDatabaseName = &v
	return s
}

func (s *TableModel) SetPhysicalTableName(v string) *TableModel {
	s.PhysicalTableName = &v
	return s
}

func (s *TableModel) SetPreviousVersion(v int64) *TableModel {
	s.PreviousVersion = &v
	return s
}

func (s *TableModel) SetRawTableName(v string) *TableModel {
	s.RawTableName = &v
	return s
}

func (s *TableModel) SetRouteColumns(v []*FieldSchemaModel) *TableModel {
	s.RouteColumns = v
	return s
}

func (s *TableModel) SetRouteEffectiveColumn(v *FieldSchemaModel) *TableModel {
	s.RouteEffectiveColumn = v
	return s
}

func (s *TableModel) SetRouteType(v string) *TableModel {
	s.RouteType = &v
	return s
}

func (s *TableModel) SetRtEngineType(v string) *TableModel {
	s.RtEngineType = &v
	return s
}

func (s *TableModel) SetRtIndexAll(v bool) *TableModel {
	s.RtIndexAll = &v
	return s
}

func (s *TableModel) SetRtModeType(v string) *TableModel {
	s.RtModeType = &v
	return s
}

func (s *TableModel) SetSd(v *StorageDescriptorModel) *TableModel {
	s.Sd = v
	return s
}

func (s *TableModel) SetStoragePolicy(v string) *TableModel {
	s.StoragePolicy = &v
	return s
}

func (s *TableModel) SetSubpartitionColumn(v string) *TableModel {
	s.SubpartitionColumn = &v
	return s
}

func (s *TableModel) SetSubpartitionCount(v int64) *TableModel {
	s.SubpartitionCount = &v
	return s
}

func (s *TableModel) SetSubpartitionType(v string) *TableModel {
	s.SubpartitionType = &v
	return s
}

func (s *TableModel) SetTableEngineName(v string) *TableModel {
	s.TableEngineName = &v
	return s
}

func (s *TableModel) SetTableName(v string) *TableModel {
	s.TableName = &v
	return s
}

func (s *TableModel) SetTableType(v string) *TableModel {
	s.TableType = &v
	return s
}

func (s *TableModel) SetTblId(v int64) *TableModel {
	s.TblId = &v
	return s
}

func (s *TableModel) SetTemporary(v bool) *TableModel {
	s.Temporary = &v
	return s
}

func (s *TableModel) SetUpdateTime(v string) *TableModel {
	s.UpdateTime = &v
	return s
}

func (s *TableModel) SetViewExpandedText(v string) *TableModel {
	s.ViewExpandedText = &v
	return s
}

func (s *TableModel) SetViewOriginalText(v string) *TableModel {
	s.ViewOriginalText = &v
	return s
}

func (s *TableModel) SetViewSecurityMode(v string) *TableModel {
	s.ViewSecurityMode = &v
	return s
}

func (s *TableModel) Validate() error {
	if s.Cols != nil {
		for _, item := range s.Cols {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DistributeColumns != nil {
		for _, item := range s.DistributeColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Indexes != nil {
		for _, item := range s.Indexes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PartitionKeys != nil {
		for _, item := range s.PartitionKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RouteColumns != nil {
		for _, item := range s.RouteColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RouteEffectiveColumn != nil {
		if err := s.RouteEffectiveColumn.Validate(); err != nil {
			return err
		}
	}
	if s.Sd != nil {
		if err := s.Sd.Validate(); err != nil {
			return err
		}
	}
	return nil
}
