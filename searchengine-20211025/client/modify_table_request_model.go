// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataProcessConfig(v []*ModifyTableRequestDataProcessConfig) *ModifyTableRequest
	GetDataProcessConfig() []*ModifyTableRequestDataProcessConfig
	SetDataSource(v *ModifyTableRequestDataSource) *ModifyTableRequest
	GetDataSource() *ModifyTableRequestDataSource
	SetFieldSchema(v map[string]*string) *ModifyTableRequest
	GetFieldSchema() map[string]*string
	SetPartitionCount(v int32) *ModifyTableRequest
	GetPartitionCount() *int32
	SetPrimaryKey(v string) *ModifyTableRequest
	GetPrimaryKey() *string
	SetRawSchema(v string) *ModifyTableRequest
	GetRawSchema() *string
	SetVectorIndex(v []*ModifyTableRequestVectorIndex) *ModifyTableRequest
	GetVectorIndex() []*ModifyTableRequestVectorIndex
	SetDryRun(v bool) *ModifyTableRequest
	GetDryRun() *bool
}

type ModifyTableRequest struct {
	// The configurations about field processing.
	DataProcessConfig []*ModifyTableRequestDataProcessConfig `json:"dataProcessConfig,omitempty" xml:"dataProcessConfig,omitempty" type:"Repeated"`
	// The configurations of the data source.
	DataSource *ModifyTableRequestDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// The fields.
	FieldSchema map[string]*string `json:"fieldSchema,omitempty" xml:"fieldSchema,omitempty"`
	// The number of data shards.
	//
	// example:
	//
	// 1
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	// The primary key field.
	//
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// The instance schema. If this parameter is specified, the parameters about the index are not required.
	//
	// example:
	//
	// {}
	RawSchema *string `json:"rawSchema,omitempty" xml:"rawSchema,omitempty"`
	// The index schema.
	VectorIndex []*ModifyTableRequestVectorIndex `json:"vectorIndex,omitempty" xml:"vectorIndex,omitempty" type:"Repeated"`
	// Specifies whether to perform only a dry run, without performing the actual request. The system only checks the validity of the data source. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyTableRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableRequest) GoString() string {
	return s.String()
}

func (s *ModifyTableRequest) GetDataProcessConfig() []*ModifyTableRequestDataProcessConfig {
	return s.DataProcessConfig
}

func (s *ModifyTableRequest) GetDataSource() *ModifyTableRequestDataSource {
	return s.DataSource
}

func (s *ModifyTableRequest) GetFieldSchema() map[string]*string {
	return s.FieldSchema
}

func (s *ModifyTableRequest) GetPartitionCount() *int32 {
	return s.PartitionCount
}

func (s *ModifyTableRequest) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *ModifyTableRequest) GetRawSchema() *string {
	return s.RawSchema
}

func (s *ModifyTableRequest) GetVectorIndex() []*ModifyTableRequestVectorIndex {
	return s.VectorIndex
}

func (s *ModifyTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyTableRequest) SetDataProcessConfig(v []*ModifyTableRequestDataProcessConfig) *ModifyTableRequest {
	s.DataProcessConfig = v
	return s
}

func (s *ModifyTableRequest) SetDataSource(v *ModifyTableRequestDataSource) *ModifyTableRequest {
	s.DataSource = v
	return s
}

func (s *ModifyTableRequest) SetFieldSchema(v map[string]*string) *ModifyTableRequest {
	s.FieldSchema = v
	return s
}

func (s *ModifyTableRequest) SetPartitionCount(v int32) *ModifyTableRequest {
	s.PartitionCount = &v
	return s
}

func (s *ModifyTableRequest) SetPrimaryKey(v string) *ModifyTableRequest {
	s.PrimaryKey = &v
	return s
}

func (s *ModifyTableRequest) SetRawSchema(v string) *ModifyTableRequest {
	s.RawSchema = &v
	return s
}

func (s *ModifyTableRequest) SetVectorIndex(v []*ModifyTableRequestVectorIndex) *ModifyTableRequest {
	s.VectorIndex = v
	return s
}

func (s *ModifyTableRequest) SetDryRun(v bool) *ModifyTableRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyTableRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyTableRequestDataProcessConfig struct {
	// The destination field.
	//
	// example:
	//
	// source_image_vector
	DstField *string `json:"dstField,omitempty" xml:"dstField,omitempty"`
	// The method used to process the field. Valid values: copy and vectorize. A value of copy specifies that the value of the source field is copied to the destination field. A value of vectorize specifies that the value of the source field is vectorized by a vectorization model and the output vector is stored in the destination field.
	//
	// example:
	//
	// vectorize
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The information about the model.
	Params *ModifyTableRequestDataProcessConfigParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// The source field.
	//
	// example:
	//
	// source_image
	SrcField *string `json:"srcField,omitempty" xml:"srcField,omitempty"`
}

func (s ModifyTableRequestDataProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableRequestDataProcessConfig) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataProcessConfig) GetDstField() *string {
	return s.DstField
}

func (s *ModifyTableRequestDataProcessConfig) GetOperator() *string {
	return s.Operator
}

func (s *ModifyTableRequestDataProcessConfig) GetParams() *ModifyTableRequestDataProcessConfigParams {
	return s.Params
}

func (s *ModifyTableRequestDataProcessConfig) GetSrcField() *string {
	return s.SrcField
}

func (s *ModifyTableRequestDataProcessConfig) SetDstField(v string) *ModifyTableRequestDataProcessConfig {
	s.DstField = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfig) SetOperator(v string) *ModifyTableRequestDataProcessConfig {
	s.Operator = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfig) SetParams(v *ModifyTableRequestDataProcessConfigParams) *ModifyTableRequestDataProcessConfig {
	s.Params = v
	return s
}

func (s *ModifyTableRequestDataProcessConfig) SetSrcField(v string) *ModifyTableRequestDataProcessConfig {
	s.SrcField = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyTableRequestDataProcessConfigParams struct {
	// The source of the data to be vectorized.
	SrcFieldConfig *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig `json:"srcFieldConfig,omitempty" xml:"srcFieldConfig,omitempty" type:"Struct"`
	// The data type.
	//
	// example:
	//
	// image
	VectorModal *string `json:"vectorModal,omitempty" xml:"vectorModal,omitempty"`
	// The vectorization model.
	//
	// example:
	//
	// clip
	VectorModel *string `json:"vectorModel,omitempty" xml:"vectorModel,omitempty"`
}

func (s ModifyTableRequestDataProcessConfigParams) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableRequestDataProcessConfigParams) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataProcessConfigParams) GetSrcFieldConfig() *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig {
	return s.SrcFieldConfig
}

func (s *ModifyTableRequestDataProcessConfigParams) GetVectorModal() *string {
	return s.VectorModal
}

func (s *ModifyTableRequestDataProcessConfigParams) GetVectorModel() *string {
	return s.VectorModel
}

func (s *ModifyTableRequestDataProcessConfigParams) SetSrcFieldConfig(v *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) *ModifyTableRequestDataProcessConfigParams {
	s.SrcFieldConfig = v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParams) SetVectorModal(v string) *ModifyTableRequestDataProcessConfigParams {
	s.VectorModal = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParams) SetVectorModel(v string) *ModifyTableRequestDataProcessConfigParams {
	s.VectorModel = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParams) Validate() error {
	return dara.Validate(s)
}

type ModifyTableRequestDataProcessConfigParamsSrcFieldConfig struct {
	// The name of the OSS bucket.
	//
	// example:
	//
	// test
	OssBucket *string `json:"ossBucket,omitempty" xml:"ossBucket,omitempty"`
	// The OSS endpoint.
	//
	// example:
	//
	// oss-cn-hangzhou-internal.aliyuncs.com
	OssEndpoint *string `json:"ossEndpoint,omitempty" xml:"ossEndpoint,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// uid
	Uid *string `json:"uid,omitempty" xml:"uid,omitempty"`
}

func (s ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) GetOssBucket() *string {
	return s.OssBucket
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) GetUid() *string {
	return s.Uid
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) SetOssBucket(v string) *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.OssBucket = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) SetOssEndpoint(v string) *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.OssEndpoint = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) SetUid(v string) *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.Uid = &v
	return s
}

func (s *ModifyTableRequestDataProcessConfigParamsSrcFieldConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyTableRequestDataSource struct {
	// Specifies whether to automatically rebuild the index.
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The configurations of the data source.
	Config *ModifyTableRequestDataSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The start timestamp from which incremental data is retrieved.
	//
	// example:
	//
	// 1715160176
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
}

func (s ModifyTableRequestDataSource) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableRequestDataSource) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataSource) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *ModifyTableRequestDataSource) GetConfig() *ModifyTableRequestDataSourceConfig {
	return s.Config
}

func (s *ModifyTableRequestDataSource) GetDataTimeSec() *int32 {
	return s.DataTimeSec
}

func (s *ModifyTableRequestDataSource) SetAutoBuildIndex(v bool) *ModifyTableRequestDataSource {
	s.AutoBuildIndex = &v
	return s
}

func (s *ModifyTableRequestDataSource) SetConfig(v *ModifyTableRequestDataSourceConfig) *ModifyTableRequestDataSource {
	s.Config = v
	return s
}

func (s *ModifyTableRequestDataSource) SetDataTimeSec(v int32) *ModifyTableRequestDataSource {
	s.DataTimeSec = &v
	return s
}

func (s *ModifyTableRequestDataSource) Validate() error {
	return dara.Validate(s)
}

type ModifyTableRequestDataSourceConfig struct {
	// The AccessKey ID of the MaxCompute data source.
	//
	// example:
	//
	// AK
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// The AccessKey secret of the MaxCompute data source.
	//
	// example:
	//
	// AS
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// antsys-shujiang-osstest
	Bucket   *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Catalog  *string `json:"catalog,omitempty" xml:"catalog,omitempty"`
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	// The endpoint of the MaxCompute data source.
	//
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The path of the Object Storage Service (OSS) object.
	//
	// example:
	//
	// oss://opensearch
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The partition in the MaxCompute table.
	//
	// example:
	//
	// ds=20231220
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The name of the MaxCompute project that is used as the data source.
	//
	// example:
	//
	// yw_dw_rpt
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The name of the MaxCompute table that is used as the data source.
	//
	// example:
	//
	// behavior
	Table       *string `json:"table,omitempty" xml:"table,omitempty"`
	TableFormat *string `json:"tableFormat,omitempty" xml:"tableFormat,omitempty"`
	Tag         *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ModifyTableRequestDataSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableRequestDataSourceConfig) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestDataSourceConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ModifyTableRequestDataSourceConfig) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *ModifyTableRequestDataSourceConfig) GetBucket() *string {
	return s.Bucket
}

func (s *ModifyTableRequestDataSourceConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *ModifyTableRequestDataSourceConfig) GetDatabase() *string {
	return s.Database
}

func (s *ModifyTableRequestDataSourceConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModifyTableRequestDataSourceConfig) GetOssPath() *string {
	return s.OssPath
}

func (s *ModifyTableRequestDataSourceConfig) GetPartition() *string {
	return s.Partition
}

func (s *ModifyTableRequestDataSourceConfig) GetProject() *string {
	return s.Project
}

func (s *ModifyTableRequestDataSourceConfig) GetTable() *string {
	return s.Table
}

func (s *ModifyTableRequestDataSourceConfig) GetTableFormat() *string {
	return s.TableFormat
}

func (s *ModifyTableRequestDataSourceConfig) GetTag() *string {
	return s.Tag
}

func (s *ModifyTableRequestDataSourceConfig) SetAccessKey(v string) *ModifyTableRequestDataSourceConfig {
	s.AccessKey = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetAccessSecret(v string) *ModifyTableRequestDataSourceConfig {
	s.AccessSecret = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetBucket(v string) *ModifyTableRequestDataSourceConfig {
	s.Bucket = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetCatalog(v string) *ModifyTableRequestDataSourceConfig {
	s.Catalog = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetDatabase(v string) *ModifyTableRequestDataSourceConfig {
	s.Database = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetEndpoint(v string) *ModifyTableRequestDataSourceConfig {
	s.Endpoint = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetOssPath(v string) *ModifyTableRequestDataSourceConfig {
	s.OssPath = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetPartition(v string) *ModifyTableRequestDataSourceConfig {
	s.Partition = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetProject(v string) *ModifyTableRequestDataSourceConfig {
	s.Project = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetTable(v string) *ModifyTableRequestDataSourceConfig {
	s.Table = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetTableFormat(v string) *ModifyTableRequestDataSourceConfig {
	s.TableFormat = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) SetTag(v string) *ModifyTableRequestDataSourceConfig {
	s.Tag = &v
	return s
}

func (s *ModifyTableRequestDataSourceConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyTableRequestVectorIndex struct {
	// The configurations of the index schema.
	AdvanceParams *ModifyTableRequestVectorIndexAdvanceParams `json:"advanceParams,omitempty" xml:"advanceParams,omitempty" type:"Struct"`
	// The dimension of the vector.
	//
	// example:
	//
	// 128
	Dimension *string `json:"dimension,omitempty" xml:"dimension,omitempty"`
	// The distance type.
	//
	// example:
	//
	// SquaredEuclidean
	DistanceType *string `json:"distanceType,omitempty" xml:"distanceType,omitempty"`
	// The name of the index schema.
	//
	// example:
	//
	// test_api
	IndexName *string `json:"indexName,omitempty" xml:"indexName,omitempty"`
	// The namespace field.
	//
	// example:
	//
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The field that stores the indexes of the elements in sparse vectors.
	//
	// example:
	//
	// sparse_indices
	SparseIndexField *string `json:"sparseIndexField,omitempty" xml:"sparseIndexField,omitempty"`
	// The field that stores the elements in sparse vectors.
	//
	// example:
	//
	// sparse_values
	SparseValueField *string `json:"sparseValueField,omitempty" xml:"sparseValueField,omitempty"`
	// The vector field.
	//
	// example:
	//
	// source_image_vector
	VectorField *string `json:"vectorField,omitempty" xml:"vectorField,omitempty"`
	// The vector retrieval algorithm.
	//
	// example:
	//
	// Qc
	VectorIndexType *string `json:"vectorIndexType,omitempty" xml:"vectorIndexType,omitempty"`
}

func (s ModifyTableRequestVectorIndex) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableRequestVectorIndex) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestVectorIndex) GetAdvanceParams() *ModifyTableRequestVectorIndexAdvanceParams {
	return s.AdvanceParams
}

func (s *ModifyTableRequestVectorIndex) GetDimension() *string {
	return s.Dimension
}

func (s *ModifyTableRequestVectorIndex) GetDistanceType() *string {
	return s.DistanceType
}

func (s *ModifyTableRequestVectorIndex) GetIndexName() *string {
	return s.IndexName
}

func (s *ModifyTableRequestVectorIndex) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyTableRequestVectorIndex) GetSparseIndexField() *string {
	return s.SparseIndexField
}

func (s *ModifyTableRequestVectorIndex) GetSparseValueField() *string {
	return s.SparseValueField
}

func (s *ModifyTableRequestVectorIndex) GetVectorField() *string {
	return s.VectorField
}

func (s *ModifyTableRequestVectorIndex) GetVectorIndexType() *string {
	return s.VectorIndexType
}

func (s *ModifyTableRequestVectorIndex) SetAdvanceParams(v *ModifyTableRequestVectorIndexAdvanceParams) *ModifyTableRequestVectorIndex {
	s.AdvanceParams = v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetDimension(v string) *ModifyTableRequestVectorIndex {
	s.Dimension = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetDistanceType(v string) *ModifyTableRequestVectorIndex {
	s.DistanceType = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetIndexName(v string) *ModifyTableRequestVectorIndex {
	s.IndexName = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetNamespace(v string) *ModifyTableRequestVectorIndex {
	s.Namespace = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetSparseIndexField(v string) *ModifyTableRequestVectorIndex {
	s.SparseIndexField = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetSparseValueField(v string) *ModifyTableRequestVectorIndex {
	s.SparseValueField = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetVectorField(v string) *ModifyTableRequestVectorIndex {
	s.VectorField = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) SetVectorIndexType(v string) *ModifyTableRequestVectorIndex {
	s.VectorIndexType = &v
	return s
}

func (s *ModifyTableRequestVectorIndex) Validate() error {
	return dara.Validate(s)
}

type ModifyTableRequestVectorIndexAdvanceParams struct {
	// The index building parameters.
	//
	// example:
	//
	// {}
	BuildIndexParams *string `json:"buildIndexParams,omitempty" xml:"buildIndexParams,omitempty"`
	// The threshold for linear building.
	//
	// example:
	//
	// 5000
	LinearBuildThreshold *string `json:"linearBuildThreshold,omitempty" xml:"linearBuildThreshold,omitempty"`
	// The minimum number of retrieved candidate sets.
	//
	// example:
	//
	// 20000
	MinScanDocCnt *string `json:"minScanDocCnt,omitempty" xml:"minScanDocCnt,omitempty"`
	// The index retrieval parameters.
	//
	// example:
	//
	// {}
	SearchIndexParams *string `json:"searchIndexParams,omitempty" xml:"searchIndexParams,omitempty"`
}

func (s ModifyTableRequestVectorIndexAdvanceParams) String() string {
	return dara.Prettify(s)
}

func (s ModifyTableRequestVectorIndexAdvanceParams) GoString() string {
	return s.String()
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) GetBuildIndexParams() *string {
	return s.BuildIndexParams
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) GetLinearBuildThreshold() *string {
	return s.LinearBuildThreshold
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) GetMinScanDocCnt() *string {
	return s.MinScanDocCnt
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) GetSearchIndexParams() *string {
	return s.SearchIndexParams
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) SetBuildIndexParams(v string) *ModifyTableRequestVectorIndexAdvanceParams {
	s.BuildIndexParams = &v
	return s
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) SetLinearBuildThreshold(v string) *ModifyTableRequestVectorIndexAdvanceParams {
	s.LinearBuildThreshold = &v
	return s
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) SetMinScanDocCnt(v string) *ModifyTableRequestVectorIndexAdvanceParams {
	s.MinScanDocCnt = &v
	return s
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) SetSearchIndexParams(v string) *ModifyTableRequestVectorIndexAdvanceParams {
	s.SearchIndexParams = &v
	return s
}

func (s *ModifyTableRequestVectorIndexAdvanceParams) Validate() error {
	return dara.Validate(s)
}
