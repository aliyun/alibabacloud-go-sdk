// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetTableInfoResponseBodyData) *GetTableInfoResponseBody
	GetData() *GetTableInfoResponseBodyData
	SetRequestId(v string) *GetTableInfoResponseBody
	GetRequestId() *string
}

type GetTableInfoResponseBody struct {
	Data      *GetTableInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	RequestId *string                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetTableInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBody) GetData() *GetTableInfoResponseBodyData {
	return s.Data
}

func (s *GetTableInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableInfoResponseBody) SetData(v *GetTableInfoResponseBodyData) *GetTableInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetTableInfoResponseBody) SetRequestId(v string) *GetTableInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTableInfoResponseBodyData struct {
	AutoRefreshEnabled             *bool                                           `json:"autoRefreshEnabled,omitempty" xml:"autoRefreshEnabled,omitempty"`
	ClusterInfo                    *GetTableInfoResponseBodyDataClusterInfo        `json:"clusterInfo,omitempty" xml:"clusterInfo,omitempty" type:"Struct"`
	Comment                        *string                                         `json:"comment,omitempty" xml:"comment,omitempty"`
	CreateTableDDL                 *string                                         `json:"createTableDDL,omitempty" xml:"createTableDDL,omitempty"`
	CreationTime                   *int64                                          `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	DisplayName                    *string                                         `json:"displayName,omitempty" xml:"displayName,omitempty"`
	FileNum                        *int64                                          `json:"fileNum,omitempty" xml:"fileNum,omitempty"`
	IsExternalTable                *bool                                           `json:"isExternalTable,omitempty" xml:"isExternalTable,omitempty"`
	IsOutdated                     *bool                                           `json:"isOutdated,omitempty" xml:"isOutdated,omitempty"`
	LastAccessTime                 *int64                                          `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	LastDDLTime                    *int64                                          `json:"lastDDLTime,omitempty" xml:"lastDDLTime,omitempty"`
	LastModifiedTime               *int64                                          `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	Lifecycle                      *string                                         `json:"lifecycle,omitempty" xml:"lifecycle,omitempty"`
	Location                       *string                                         `json:"location,omitempty" xml:"location,omitempty"`
	MaterializedView               *bool                                           `json:"materializedView,omitempty" xml:"materializedView,omitempty"`
	Name                           *string                                         `json:"name,omitempty" xml:"name,omitempty"`
	NativeColumns                  []*GetTableInfoResponseBodyDataNativeColumns    `json:"nativeColumns,omitempty" xml:"nativeColumns,omitempty" type:"Repeated"`
	OdpsPropertiesRolearn          *string                                         `json:"odpsPropertiesRolearn,omitempty" xml:"odpsPropertiesRolearn,omitempty"`
	OdpsSqlTextOptionFlushHeader   *bool                                           `json:"odpsSqlTextOptionFlushHeader,omitempty" xml:"odpsSqlTextOptionFlushHeader,omitempty"`
	OdpsTextOptionHeaderLinesCount *int64                                          `json:"odpsTextOptionHeaderLinesCount,omitempty" xml:"odpsTextOptionHeaderLinesCount,omitempty"`
	Owner                          *string                                         `json:"owner,omitempty" xml:"owner,omitempty"`
	PartitionColumns               []*GetTableInfoResponseBodyDataPartitionColumns `json:"partitionColumns,omitempty" xml:"partitionColumns,omitempty" type:"Repeated"`
	PhysicalSize                   *int64                                          `json:"physicalSize,omitempty" xml:"physicalSize,omitempty"`
	ProjectName                    *string                                         `json:"projectName,omitempty" xml:"projectName,omitempty"`
	RewriteEnabled                 *bool                                           `json:"rewriteEnabled,omitempty" xml:"rewriteEnabled,omitempty"`
	Schema                         *string                                         `json:"schema,omitempty" xml:"schema,omitempty"`
	Size                           *int64                                          `json:"size,omitempty" xml:"size,omitempty"`
	StorageHandler                 *string                                         `json:"storageHandler,omitempty" xml:"storageHandler,omitempty"`
	TableLabel                     *string                                         `json:"tableLabel,omitempty" xml:"tableLabel,omitempty"`
	TablesotreTableName            *string                                         `json:"tablesotreTableName,omitempty" xml:"tablesotreTableName,omitempty"`
	TablestoreColumnsMapping       *string                                         `json:"tablestoreColumnsMapping,omitempty" xml:"tablestoreColumnsMapping,omitempty"`
	Type                           *string                                         `json:"type,omitempty" xml:"type,omitempty"`
	ViewText                       *string                                         `json:"viewText,omitempty" xml:"viewText,omitempty"`
}

func (s GetTableInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTableInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyData) GetAutoRefreshEnabled() *bool {
	return s.AutoRefreshEnabled
}

func (s *GetTableInfoResponseBodyData) GetClusterInfo() *GetTableInfoResponseBodyDataClusterInfo {
	return s.ClusterInfo
}

func (s *GetTableInfoResponseBodyData) GetComment() *string {
	return s.Comment
}

func (s *GetTableInfoResponseBodyData) GetCreateTableDDL() *string {
	return s.CreateTableDDL
}

func (s *GetTableInfoResponseBodyData) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *GetTableInfoResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetTableInfoResponseBodyData) GetFileNum() *int64 {
	return s.FileNum
}

func (s *GetTableInfoResponseBodyData) GetIsExternalTable() *bool {
	return s.IsExternalTable
}

func (s *GetTableInfoResponseBodyData) GetIsOutdated() *bool {
	return s.IsOutdated
}

func (s *GetTableInfoResponseBodyData) GetLastAccessTime() *int64 {
	return s.LastAccessTime
}

func (s *GetTableInfoResponseBodyData) GetLastDDLTime() *int64 {
	return s.LastDDLTime
}

func (s *GetTableInfoResponseBodyData) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *GetTableInfoResponseBodyData) GetLifecycle() *string {
	return s.Lifecycle
}

func (s *GetTableInfoResponseBodyData) GetLocation() *string {
	return s.Location
}

func (s *GetTableInfoResponseBodyData) GetMaterializedView() *bool {
	return s.MaterializedView
}

func (s *GetTableInfoResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetTableInfoResponseBodyData) GetNativeColumns() []*GetTableInfoResponseBodyDataNativeColumns {
	return s.NativeColumns
}

func (s *GetTableInfoResponseBodyData) GetOdpsPropertiesRolearn() *string {
	return s.OdpsPropertiesRolearn
}

func (s *GetTableInfoResponseBodyData) GetOdpsSqlTextOptionFlushHeader() *bool {
	return s.OdpsSqlTextOptionFlushHeader
}

func (s *GetTableInfoResponseBodyData) GetOdpsTextOptionHeaderLinesCount() *int64 {
	return s.OdpsTextOptionHeaderLinesCount
}

func (s *GetTableInfoResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetTableInfoResponseBodyData) GetPartitionColumns() []*GetTableInfoResponseBodyDataPartitionColumns {
	return s.PartitionColumns
}

func (s *GetTableInfoResponseBodyData) GetPhysicalSize() *int64 {
	return s.PhysicalSize
}

func (s *GetTableInfoResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTableInfoResponseBodyData) GetRewriteEnabled() *bool {
	return s.RewriteEnabled
}

func (s *GetTableInfoResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *GetTableInfoResponseBodyData) GetSize() *int64 {
	return s.Size
}

func (s *GetTableInfoResponseBodyData) GetStorageHandler() *string {
	return s.StorageHandler
}

func (s *GetTableInfoResponseBodyData) GetTableLabel() *string {
	return s.TableLabel
}

func (s *GetTableInfoResponseBodyData) GetTablesotreTableName() *string {
	return s.TablesotreTableName
}

func (s *GetTableInfoResponseBodyData) GetTablestoreColumnsMapping() *string {
	return s.TablestoreColumnsMapping
}

func (s *GetTableInfoResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetTableInfoResponseBodyData) GetViewText() *string {
	return s.ViewText
}

func (s *GetTableInfoResponseBodyData) SetAutoRefreshEnabled(v bool) *GetTableInfoResponseBodyData {
	s.AutoRefreshEnabled = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetClusterInfo(v *GetTableInfoResponseBodyDataClusterInfo) *GetTableInfoResponseBodyData {
	s.ClusterInfo = v
	return s
}

func (s *GetTableInfoResponseBodyData) SetComment(v string) *GetTableInfoResponseBodyData {
	s.Comment = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetCreateTableDDL(v string) *GetTableInfoResponseBodyData {
	s.CreateTableDDL = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetCreationTime(v int64) *GetTableInfoResponseBodyData {
	s.CreationTime = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetDisplayName(v string) *GetTableInfoResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetFileNum(v int64) *GetTableInfoResponseBodyData {
	s.FileNum = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetIsExternalTable(v bool) *GetTableInfoResponseBodyData {
	s.IsExternalTable = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetIsOutdated(v bool) *GetTableInfoResponseBodyData {
	s.IsOutdated = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLastAccessTime(v int64) *GetTableInfoResponseBodyData {
	s.LastAccessTime = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLastDDLTime(v int64) *GetTableInfoResponseBodyData {
	s.LastDDLTime = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLastModifiedTime(v int64) *GetTableInfoResponseBodyData {
	s.LastModifiedTime = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLifecycle(v string) *GetTableInfoResponseBodyData {
	s.Lifecycle = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetLocation(v string) *GetTableInfoResponseBodyData {
	s.Location = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetMaterializedView(v bool) *GetTableInfoResponseBodyData {
	s.MaterializedView = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetName(v string) *GetTableInfoResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetNativeColumns(v []*GetTableInfoResponseBodyDataNativeColumns) *GetTableInfoResponseBodyData {
	s.NativeColumns = v
	return s
}

func (s *GetTableInfoResponseBodyData) SetOdpsPropertiesRolearn(v string) *GetTableInfoResponseBodyData {
	s.OdpsPropertiesRolearn = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetOdpsSqlTextOptionFlushHeader(v bool) *GetTableInfoResponseBodyData {
	s.OdpsSqlTextOptionFlushHeader = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetOdpsTextOptionHeaderLinesCount(v int64) *GetTableInfoResponseBodyData {
	s.OdpsTextOptionHeaderLinesCount = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetOwner(v string) *GetTableInfoResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetPartitionColumns(v []*GetTableInfoResponseBodyDataPartitionColumns) *GetTableInfoResponseBodyData {
	s.PartitionColumns = v
	return s
}

func (s *GetTableInfoResponseBodyData) SetPhysicalSize(v int64) *GetTableInfoResponseBodyData {
	s.PhysicalSize = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetProjectName(v string) *GetTableInfoResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetRewriteEnabled(v bool) *GetTableInfoResponseBodyData {
	s.RewriteEnabled = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetSchema(v string) *GetTableInfoResponseBodyData {
	s.Schema = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetSize(v int64) *GetTableInfoResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetStorageHandler(v string) *GetTableInfoResponseBodyData {
	s.StorageHandler = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetTableLabel(v string) *GetTableInfoResponseBodyData {
	s.TableLabel = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetTablesotreTableName(v string) *GetTableInfoResponseBodyData {
	s.TablesotreTableName = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetTablestoreColumnsMapping(v string) *GetTableInfoResponseBodyData {
	s.TablestoreColumnsMapping = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetType(v string) *GetTableInfoResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetTableInfoResponseBodyData) SetViewText(v string) *GetTableInfoResponseBodyData {
	s.ViewText = &v
	return s
}

func (s *GetTableInfoResponseBodyData) Validate() error {
	if s.ClusterInfo != nil {
		if err := s.ClusterInfo.Validate(); err != nil {
			return err
		}
	}
	if s.NativeColumns != nil {
		for _, item := range s.NativeColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PartitionColumns != nil {
		for _, item := range s.PartitionColumns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableInfoResponseBodyDataClusterInfo struct {
	BucketNum   *int64                                             `json:"bucketNum,omitempty" xml:"bucketNum,omitempty"`
	ClusterCols []*string                                          `json:"clusterCols,omitempty" xml:"clusterCols,omitempty" type:"Repeated"`
	ClusterType *string                                            `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	SortCols    []*GetTableInfoResponseBodyDataClusterInfoSortCols `json:"sortCols,omitempty" xml:"sortCols,omitempty" type:"Repeated"`
}

func (s GetTableInfoResponseBodyDataClusterInfo) String() string {
	return dara.Prettify(s)
}

func (s GetTableInfoResponseBodyDataClusterInfo) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyDataClusterInfo) GetBucketNum() *int64 {
	return s.BucketNum
}

func (s *GetTableInfoResponseBodyDataClusterInfo) GetClusterCols() []*string {
	return s.ClusterCols
}

func (s *GetTableInfoResponseBodyDataClusterInfo) GetClusterType() *string {
	return s.ClusterType
}

func (s *GetTableInfoResponseBodyDataClusterInfo) GetSortCols() []*GetTableInfoResponseBodyDataClusterInfoSortCols {
	return s.SortCols
}

func (s *GetTableInfoResponseBodyDataClusterInfo) SetBucketNum(v int64) *GetTableInfoResponseBodyDataClusterInfo {
	s.BucketNum = &v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfo) SetClusterCols(v []*string) *GetTableInfoResponseBodyDataClusterInfo {
	s.ClusterCols = v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfo) SetClusterType(v string) *GetTableInfoResponseBodyDataClusterInfo {
	s.ClusterType = &v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfo) SetSortCols(v []*GetTableInfoResponseBodyDataClusterInfoSortCols) *GetTableInfoResponseBodyDataClusterInfo {
	s.SortCols = v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfo) Validate() error {
	if s.SortCols != nil {
		for _, item := range s.SortCols {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTableInfoResponseBodyDataClusterInfoSortCols struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
}

func (s GetTableInfoResponseBodyDataClusterInfoSortCols) String() string {
	return dara.Prettify(s)
}

func (s GetTableInfoResponseBodyDataClusterInfoSortCols) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyDataClusterInfoSortCols) GetName() *string {
	return s.Name
}

func (s *GetTableInfoResponseBodyDataClusterInfoSortCols) GetOrder() *string {
	return s.Order
}

func (s *GetTableInfoResponseBodyDataClusterInfoSortCols) SetName(v string) *GetTableInfoResponseBodyDataClusterInfoSortCols {
	s.Name = &v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfoSortCols) SetOrder(v string) *GetTableInfoResponseBodyDataClusterInfoSortCols {
	s.Order = &v
	return s
}

func (s *GetTableInfoResponseBodyDataClusterInfoSortCols) Validate() error {
	return dara.Validate(s)
}

type GetTableInfoResponseBodyDataNativeColumns struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTableInfoResponseBodyDataNativeColumns) String() string {
	return dara.Prettify(s)
}

func (s GetTableInfoResponseBodyDataNativeColumns) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyDataNativeColumns) GetComment() *string {
	return s.Comment
}

func (s *GetTableInfoResponseBodyDataNativeColumns) GetLabel() *string {
	return s.Label
}

func (s *GetTableInfoResponseBodyDataNativeColumns) GetName() *string {
	return s.Name
}

func (s *GetTableInfoResponseBodyDataNativeColumns) GetType() *string {
	return s.Type
}

func (s *GetTableInfoResponseBodyDataNativeColumns) SetComment(v string) *GetTableInfoResponseBodyDataNativeColumns {
	s.Comment = &v
	return s
}

func (s *GetTableInfoResponseBodyDataNativeColumns) SetLabel(v string) *GetTableInfoResponseBodyDataNativeColumns {
	s.Label = &v
	return s
}

func (s *GetTableInfoResponseBodyDataNativeColumns) SetName(v string) *GetTableInfoResponseBodyDataNativeColumns {
	s.Name = &v
	return s
}

func (s *GetTableInfoResponseBodyDataNativeColumns) SetType(v string) *GetTableInfoResponseBodyDataNativeColumns {
	s.Type = &v
	return s
}

func (s *GetTableInfoResponseBodyDataNativeColumns) Validate() error {
	return dara.Validate(s)
}

type GetTableInfoResponseBodyDataPartitionColumns struct {
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	Label   *string `json:"label,omitempty" xml:"label,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTableInfoResponseBodyDataPartitionColumns) String() string {
	return dara.Prettify(s)
}

func (s GetTableInfoResponseBodyDataPartitionColumns) GoString() string {
	return s.String()
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) GetComment() *string {
	return s.Comment
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) GetLabel() *string {
	return s.Label
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) GetName() *string {
	return s.Name
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) GetType() *string {
	return s.Type
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) SetComment(v string) *GetTableInfoResponseBodyDataPartitionColumns {
	s.Comment = &v
	return s
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) SetLabel(v string) *GetTableInfoResponseBodyDataPartitionColumns {
	s.Label = &v
	return s
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) SetName(v string) *GetTableInfoResponseBodyDataPartitionColumns {
	s.Name = &v
	return s
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) SetType(v string) *GetTableInfoResponseBodyDataPartitionColumns {
	s.Type = &v
	return s
}

func (s *GetTableInfoResponseBodyDataPartitionColumns) Validate() error {
	return dara.Validate(s)
}
