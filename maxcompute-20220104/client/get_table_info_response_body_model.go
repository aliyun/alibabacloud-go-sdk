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
	// The response result.
	Data *GetTableInfoResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0a06dd4516687375802853481ec9fd
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	// Indicates whether automatic refresh of the materialized view is enabled. This parameter is returned only when the type is materialized view.
	//
	// example:
	//
	// false
	AutoRefreshEnabled *bool `json:"autoRefreshEnabled,omitempty" xml:"autoRefreshEnabled,omitempty"`
	// The clustering properties. This parameter is returned only when the table type is clustered table.
	ClusterInfo *GetTableInfoResponseBodyDataClusterInfo `json:"clusterInfo,omitempty" xml:"clusterInfo,omitempty" type:"Struct"`
	// The table comment.
	//
	// example:
	//
	// sale_detail
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The DDL statement used to create the table.
	//
	// example:
	//
	// create table if not exists sale_detail( shop_name STRING, customer_id STRING, total_price DOUBLE) partitioned by (sale_date STRING, region STRING);
	CreateTableDDL *string `json:"createTableDDL,omitempty" xml:"createTableDDL,omitempty"`
	// The creation time. The time is in ISO 8601 format and in the UTC time zone.
	//
	// example:
	//
	// 2022-01-17T07:07:47Z
	CreationTime *int64 `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// The display name.
	//
	// example:
	//
	// odps_project.default.sale_detail
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The number of files.
	//
	// example:
	//
	// 200
	FileNum *int64 `json:"fileNum,omitempty" xml:"fileNum,omitempty"`
	// Indicates whether the table is an external table. This parameter is returned only when the type is table.
	//
	// example:
	//
	// false
	IsExternalTable *bool `json:"isExternalTable,omitempty" xml:"isExternalTable,omitempty"`
	// Indicates whether the data in the materialized view is invalid due to changes in the source table data. This parameter is returned only when the type is materialized view.
	//
	// example:
	//
	// false
	IsOutdated *bool `json:"isOutdated,omitempty" xml:"isOutdated,omitempty"`
	// The time when the data of the table or view was last accessed. This parameter is returned only when the table is a non-partitioned table.
	//
	// > This data has been gradually collected since July 2023. If the table has not been accessed after the upgrade, or if it is accessed through ALGO jobs or Hologres direct read, the last_access_time cannot be collected.
	//
	// example:
	//
	// 2023-11-21T02:05:56Z
	LastAccessTime *int64 `json:"lastAccessTime,omitempty" xml:"lastAccessTime,omitempty"`
	// The time when the DDL statement of the table or view was last modified.
	//
	// example:
	//
	// 2023-11-21T02:05:56Z
	LastDDLTime *int64 `json:"lastDDLTime,omitempty" xml:"lastDDLTime,omitempty"`
	// The time when the data of the table or view was last modified. The time is in ISO 8601 format and in the UTC time zone.
	//
	// example:
	//
	// 2023-11-21T02:05:56Z
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The lifecycle in days.
	//
	// example:
	//
	// -1
	Lifecycle *string `json:"lifecycle,omitempty" xml:"lifecycle,omitempty"`
	// The external table location. This parameter is returned only when the type is external table.
	//
	// example:
	//
	// oss://oss-cn-hangzhou-internal.aliyuncs.com/oss-mc-test/Demo1/
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// Indicates whether the table is a <props="china">[materialized view](https://www.alibabacloud.com/help/en/maxcompute/user-guide/materialized-view-operations)
	//
	// <props="intl">[materialized view](https://www.alibabacloud.com/help/maxcompute/user-guide/materialized-view-operations).
	//
	// example:
	//
	// false
	MaterializedView *bool `json:"materializedView,omitempty" xml:"materializedView,omitempty"`
	// The name of the table or view.
	//
	// example:
	//
	// sale_detail
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The column information.
	NativeColumns []*GetTableInfoResponseBodyDataNativeColumns `json:"nativeColumns,omitempty" xml:"nativeColumns,omitempty" type:"Repeated"`
	// The ARN of AliyunODPSDefaultRole in RAM. This parameter is returned only when the type is external table.
	//
	// example:
	//
	// acs:ram::xxxxx:role/aliyunodpsdefaultrole
	OdpsPropertiesRolearn *string `json:"odpsPropertiesRolearn,omitempty" xml:"odpsPropertiesRolearn,omitempty"`
	// Indicates whether to ignore the header. This parameter is returned only when the type is external table.
	//
	// example:
	//
	// true
	OdpsSqlTextOptionFlushHeader *bool `json:"odpsSqlTextOptionFlushHeader,omitempty" xml:"odpsSqlTextOptionFlushHeader,omitempty"`
	// The number of header lines to ignore. This parameter is returned only when the type is external table.
	//
	// example:
	//
	// 1
	OdpsTextOptionHeaderLinesCount *int64 `json:"odpsTextOptionHeaderLinesCount,omitempty" xml:"odpsTextOptionHeaderLinesCount,omitempty"`
	// The account information of the owner of the table or view.
	//
	// example:
	//
	// ALIYUN$odps****@aliyunid.com
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The partition column information. This parameter is returned only for partitioned tables.
	PartitionColumns []*GetTableInfoResponseBodyDataPartitionColumns `json:"partitionColumns,omitempty" xml:"partitionColumns,omitempty" type:"Repeated"`
	// The physical size.
	//
	// example:
	//
	// 2763
	PhysicalSize *int64 `json:"physicalSize,omitempty" xml:"physicalSize,omitempty"`
	// The name of the project to which the table or view belongs.
	//
	// example:
	//
	// odps_project
	ProjectName *string `json:"projectName,omitempty" xml:"projectName,omitempty"`
	// Indicates whether query rewrite through the materialized view is allowed. This parameter is returned only when the type is materialized view.
	//
	// example:
	//
	// false
	RewriteEnabled *bool `json:"rewriteEnabled,omitempty" xml:"rewriteEnabled,omitempty"`
	// The name of the schema to which the table or view belongs.
	//
	// example:
	//
	// default
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The data size of the table. If the table is a non-partitioned table, the value is the data size of the table. If the table is a partitioned table, the system does not calculate the data size of the table, and the value is NULL. The PARTITIONS view contains the data size of each partition in the partitioned table. Unit: bytes.
	//
	// example:
	//
	// 5372
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The external table parser. This parameter is returned only when the type is external table.
	//
	// example:
	//
	// com.aliyun.odps.CsvStorageHandler
	StorageHandler *string `json:"storageHandler,omitempty" xml:"storageHandler,omitempty"`
	// The sensitivity level label of the table. For more information, see <props="china">[Label-based access control](https://www.alibabacloud.com/help/en/maxcompute/user-guide/label-based-access-control)
	//
	// <props="intl">[Label-based access control](https://www.alibabacloud.com/help/maxcompute/user-guide/label-based-access-control).
	//
	// example:
	//
	// 0
	TableLabel *string `json:"tableLabel,omitempty" xml:"tableLabel,omitempty"`
	// The name of the Tablestore table to be accessed. This parameter is returned only when the type is external table.
	//
	// example:
	//
	// ots_tpch_orders
	TablesotreTableName *string `json:"tablesotreTableName,omitempty" xml:"tablesotreTableName,omitempty"`
	// The columns of the Tablestore table to be accessed by MaxCompute, including the primary key columns and attribute columns. This parameter is returned only when the type is external table.
	//
	// example:
	//
	// o_orderkey,o_orderdate,o_custkey,o_orderstatus,o_totalprice
	TablestoreColumnsMapping *string `json:"tablestoreColumnsMapping,omitempty" xml:"tablestoreColumnsMapping,omitempty"`
	// The type of the table or view. Valid values:
	//
	// - **internal**: Internal table.
	//
	// - **external**: External table.
	//
	// - **view**: View.
	//
	// - **materializedView**: <props="china">[Materialized view](https://www.alibabacloud.com/help/en/maxcompute/user-guide/materialized-view-operations)
	//
	// <props="intl">[Materialized view](https://www.alibabacloud.com/help/maxcompute/user-guide/materialized-view-operations).
	//
	// example:
	//
	// internal
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The statement used to create the view. This parameter is returned only when the type is view.
	//
	// example:
	//
	// select shop_name, sum(total_price) from sale_detail group by shop_name
	ViewText *string `json:"viewText,omitempty" xml:"viewText,omitempty"`
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
	// The number of buckets in the clustered table. This is an optional field. A value of 0 indicates that the number of buckets is dynamically determined during job execution.
	//
	// example:
	//
	// 1024
	BucketNum *int64 `json:"bucketNum,omitempty" xml:"bucketNum,omitempty"`
	// The cluster key.
	ClusterCols []*string `json:"clusterCols,omitempty" xml:"clusterCols,omitempty" type:"Repeated"`
	// The clustering type of the table. MaxCompute currently supports <props="china">[Hash Clustering](https://www.alibabacloud.com/help/en/maxcompute/use-cases/hash-clustering)
	//
	// <props="intl">[Hash Clustering](https://www.alibabacloud.com/help/maxcompute/use-cases/hash-clustering) and
	//
	// <props="china"> [Range Clustering](https://www.alibabacloud.com/help/en/maxcompute/use-cases/hash-clustering)
	//
	// <props="intl">[Range Clustering](https://www.alibabacloud.com/help/maxcompute/use-cases/range-clustering).
	//
	// example:
	//
	// Hash
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	// The sort columns.
	SortCols []*GetTableInfoResponseBodyDataClusterInfoSortCols `json:"sortCols,omitempty" xml:"sortCols,omitempty" type:"Repeated"`
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
	// The name of the sort column.
	//
	// example:
	//
	// col_2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The sort order.
	//
	// example:
	//
	// DESC
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
	// The column comment.
	//
	// example:
	//
	// Store name
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The sensitivity level label of the column. For more information, see <props="china">[Label-based access control](https://www.alibabacloud.com/help/en/maxcompute/user-guide/label-based-access-control)
	//
	// <props="intl">[Label-based access control](https://www.alibabacloud.com/help/maxcompute/user-guide/label-based-access-control).
	//
	// example:
	//
	// 0
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The column name.
	//
	// example:
	//
	// shop_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The data type.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	// The partition column comment.
	//
	// example:
	//
	// Sale date
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The sensitivity level label of the column. For more information, see <props="china">[Label-based access control](https://www.alibabacloud.com/help/en/maxcompute/user-guide/label-based-access-control)
	//
	// <props="intl">[Label-based access control](https://www.alibabacloud.com/help/maxcompute/user-guide/label-based-access-control).
	//
	// example:
	//
	// 0
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The partition name.
	//
	// example:
	//
	// ds=\\"20190130\\"
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The data type.
	//
	// example:
	//
	// STRING
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
