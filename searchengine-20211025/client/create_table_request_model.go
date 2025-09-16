// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataProcessConfig(v []*CreateTableRequestDataProcessConfig) *CreateTableRequest
	GetDataProcessConfig() []*CreateTableRequestDataProcessConfig
	SetDataProcessorCount(v int32) *CreateTableRequest
	GetDataProcessorCount() *int32
	SetDataSource(v *CreateTableRequestDataSource) *CreateTableRequest
	GetDataSource() *CreateTableRequestDataSource
	SetFieldSchema(v map[string]*string) *CreateTableRequest
	GetFieldSchema() map[string]*string
	SetName(v string) *CreateTableRequest
	GetName() *string
	SetPartitionCount(v int32) *CreateTableRequest
	GetPartitionCount() *int32
	SetPrimaryKey(v string) *CreateTableRequest
	GetPrimaryKey() *string
	SetRawSchema(v string) *CreateTableRequest
	GetRawSchema() *string
	SetScene(v string) *CreateTableRequest
	GetScene() *string
	SetVectorIndex(v []*CreateTableRequestVectorIndex) *CreateTableRequest
	GetVectorIndex() []*CreateTableRequestVectorIndex
	SetDryRun(v bool) *CreateTableRequest
	GetDryRun() *bool
}

type CreateTableRequest struct {
	// The configurations about field processing.
	DataProcessConfig []*CreateTableRequestDataProcessConfig `json:"dataProcessConfig,omitempty" xml:"dataProcessConfig,omitempty" type:"Repeated"`
	// The number of resources used for data update.
	//
	// example:
	//
	// 1
	DataProcessorCount *int32 `json:"dataProcessorCount,omitempty" xml:"dataProcessorCount,omitempty"`
	// The configurations of the data source.
	DataSource *CreateTableRequestDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// The fields.
	FieldSchema map[string]*string `json:"fieldSchema,omitempty" xml:"fieldSchema,omitempty"`
	// The index name.
	//
	// example:
	//
	// index_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	Scene     *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The index schema.
	VectorIndex []*CreateTableRequestVectorIndex `json:"vectorIndex,omitempty" xml:"vectorIndex,omitempty" type:"Repeated"`
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

func (s CreateTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTableRequest) GetDataProcessConfig() []*CreateTableRequestDataProcessConfig {
	return s.DataProcessConfig
}

func (s *CreateTableRequest) GetDataProcessorCount() *int32 {
	return s.DataProcessorCount
}

func (s *CreateTableRequest) GetDataSource() *CreateTableRequestDataSource {
	return s.DataSource
}

func (s *CreateTableRequest) GetFieldSchema() map[string]*string {
	return s.FieldSchema
}

func (s *CreateTableRequest) GetName() *string {
	return s.Name
}

func (s *CreateTableRequest) GetPartitionCount() *int32 {
	return s.PartitionCount
}

func (s *CreateTableRequest) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *CreateTableRequest) GetRawSchema() *string {
	return s.RawSchema
}

func (s *CreateTableRequest) GetScene() *string {
	return s.Scene
}

func (s *CreateTableRequest) GetVectorIndex() []*CreateTableRequestVectorIndex {
	return s.VectorIndex
}

func (s *CreateTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTableRequest) SetDataProcessConfig(v []*CreateTableRequestDataProcessConfig) *CreateTableRequest {
	s.DataProcessConfig = v
	return s
}

func (s *CreateTableRequest) SetDataProcessorCount(v int32) *CreateTableRequest {
	s.DataProcessorCount = &v
	return s
}

func (s *CreateTableRequest) SetDataSource(v *CreateTableRequestDataSource) *CreateTableRequest {
	s.DataSource = v
	return s
}

func (s *CreateTableRequest) SetFieldSchema(v map[string]*string) *CreateTableRequest {
	s.FieldSchema = v
	return s
}

func (s *CreateTableRequest) SetName(v string) *CreateTableRequest {
	s.Name = &v
	return s
}

func (s *CreateTableRequest) SetPartitionCount(v int32) *CreateTableRequest {
	s.PartitionCount = &v
	return s
}

func (s *CreateTableRequest) SetPrimaryKey(v string) *CreateTableRequest {
	s.PrimaryKey = &v
	return s
}

func (s *CreateTableRequest) SetRawSchema(v string) *CreateTableRequest {
	s.RawSchema = &v
	return s
}

func (s *CreateTableRequest) SetScene(v string) *CreateTableRequest {
	s.Scene = &v
	return s
}

func (s *CreateTableRequest) SetVectorIndex(v []*CreateTableRequestVectorIndex) *CreateTableRequest {
	s.VectorIndex = v
	return s
}

func (s *CreateTableRequest) SetDryRun(v bool) *CreateTableRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTableRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestDataProcessConfig struct {
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
	Params *CreateTableRequestDataProcessConfigParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// The source field.
	//
	// example:
	//
	// source_image
	SrcField *string `json:"srcField,omitempty" xml:"srcField,omitempty"`
}

func (s CreateTableRequestDataProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestDataProcessConfig) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataProcessConfig) GetDstField() *string {
	return s.DstField
}

func (s *CreateTableRequestDataProcessConfig) GetOperator() *string {
	return s.Operator
}

func (s *CreateTableRequestDataProcessConfig) GetParams() *CreateTableRequestDataProcessConfigParams {
	return s.Params
}

func (s *CreateTableRequestDataProcessConfig) GetSrcField() *string {
	return s.SrcField
}

func (s *CreateTableRequestDataProcessConfig) SetDstField(v string) *CreateTableRequestDataProcessConfig {
	s.DstField = &v
	return s
}

func (s *CreateTableRequestDataProcessConfig) SetOperator(v string) *CreateTableRequestDataProcessConfig {
	s.Operator = &v
	return s
}

func (s *CreateTableRequestDataProcessConfig) SetParams(v *CreateTableRequestDataProcessConfigParams) *CreateTableRequestDataProcessConfig {
	s.Params = v
	return s
}

func (s *CreateTableRequestDataProcessConfig) SetSrcField(v string) *CreateTableRequestDataProcessConfig {
	s.SrcField = &v
	return s
}

func (s *CreateTableRequestDataProcessConfig) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestDataProcessConfigParams struct {
	// The source of the data to be vectorized.
	SrcFieldConfig *CreateTableRequestDataProcessConfigParamsSrcFieldConfig `json:"srcFieldConfig,omitempty" xml:"srcFieldConfig,omitempty" type:"Struct"`
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

func (s CreateTableRequestDataProcessConfigParams) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestDataProcessConfigParams) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataProcessConfigParams) GetSrcFieldConfig() *CreateTableRequestDataProcessConfigParamsSrcFieldConfig {
	return s.SrcFieldConfig
}

func (s *CreateTableRequestDataProcessConfigParams) GetVectorModal() *string {
	return s.VectorModal
}

func (s *CreateTableRequestDataProcessConfigParams) GetVectorModel() *string {
	return s.VectorModel
}

func (s *CreateTableRequestDataProcessConfigParams) SetSrcFieldConfig(v *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) *CreateTableRequestDataProcessConfigParams {
	s.SrcFieldConfig = v
	return s
}

func (s *CreateTableRequestDataProcessConfigParams) SetVectorModal(v string) *CreateTableRequestDataProcessConfigParams {
	s.VectorModal = &v
	return s
}

func (s *CreateTableRequestDataProcessConfigParams) SetVectorModel(v string) *CreateTableRequestDataProcessConfigParams {
	s.VectorModel = &v
	return s
}

func (s *CreateTableRequestDataProcessConfigParams) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestDataProcessConfigParamsSrcFieldConfig struct {
	// The OSS bucket.
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

func (s CreateTableRequestDataProcessConfigParamsSrcFieldConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestDataProcessConfigParamsSrcFieldConfig) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) GetOssBucket() *string {
	return s.OssBucket
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) GetUid() *string {
	return s.Uid
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) SetOssBucket(v string) *CreateTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.OssBucket = &v
	return s
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) SetOssEndpoint(v string) *CreateTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.OssEndpoint = &v
	return s
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) SetUid(v string) *CreateTableRequestDataProcessConfigParamsSrcFieldConfig {
	s.Uid = &v
	return s
}

func (s *CreateTableRequestDataProcessConfigParamsSrcFieldConfig) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestDataSource struct {
	// Specifies whether to automatically rebuild the index.
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The configurations of the data source.
	Config *CreateTableRequestDataSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The start timestamp from which incremental data is retrieved.
	//
	// example:
	//
	// 1715160176
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The data source type. Valid values: odps, swift, and oss.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateTableRequestDataSource) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestDataSource) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataSource) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *CreateTableRequestDataSource) GetConfig() *CreateTableRequestDataSourceConfig {
	return s.Config
}

func (s *CreateTableRequestDataSource) GetDataTimeSec() *int32 {
	return s.DataTimeSec
}

func (s *CreateTableRequestDataSource) GetType() *string {
	return s.Type
}

func (s *CreateTableRequestDataSource) SetAutoBuildIndex(v bool) *CreateTableRequestDataSource {
	s.AutoBuildIndex = &v
	return s
}

func (s *CreateTableRequestDataSource) SetConfig(v *CreateTableRequestDataSourceConfig) *CreateTableRequestDataSource {
	s.Config = v
	return s
}

func (s *CreateTableRequestDataSource) SetDataTimeSec(v int32) *CreateTableRequestDataSource {
	s.DataTimeSec = &v
	return s
}

func (s *CreateTableRequestDataSource) SetType(v string) *CreateTableRequestDataSource {
	s.Type = &v
	return s
}

func (s *CreateTableRequestDataSource) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestDataSourceConfig struct {
	// The AccessKey ID of the MaxCompute data source.
	//
	// example:
	//
	// ak
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// The AccessKey secret of the MaxCompute data source.
	//
	// example:
	//
	// as
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// The OSS bucket.
	//
	// example:
	//
	// antsys-flytest-ci
	Bucket   *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Catalog  *string `json:"catalog,omitempty" xml:"catalog,omitempty"`
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	// The endpoint of the MaxCompute data source.
	//
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The Object Storage Service (OSS) path.
	//
	// example:
	//
	// oss://opensearch
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The partition in the MaxCompute table. This parameter is required if type is set to odps.
	//
	// example:
	//
	// ds=20220713
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The name of the MaxCompute project that is used as the data source.
	//
	// example:
	//
	// project_20210220122847_3218
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The name of the MaxCompute table that is used as the data source.
	//
	// example:
	//
	// test56
	Table       *string `json:"table,omitempty" xml:"table,omitempty"`
	TableFormat *string `json:"tableFormat,omitempty" xml:"tableFormat,omitempty"`
	Tag         *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s CreateTableRequestDataSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestDataSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateTableRequestDataSourceConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateTableRequestDataSourceConfig) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *CreateTableRequestDataSourceConfig) GetBucket() *string {
	return s.Bucket
}

func (s *CreateTableRequestDataSourceConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *CreateTableRequestDataSourceConfig) GetDatabase() *string {
	return s.Database
}

func (s *CreateTableRequestDataSourceConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateTableRequestDataSourceConfig) GetOssPath() *string {
	return s.OssPath
}

func (s *CreateTableRequestDataSourceConfig) GetPartition() *string {
	return s.Partition
}

func (s *CreateTableRequestDataSourceConfig) GetProject() *string {
	return s.Project
}

func (s *CreateTableRequestDataSourceConfig) GetTable() *string {
	return s.Table
}

func (s *CreateTableRequestDataSourceConfig) GetTableFormat() *string {
	return s.TableFormat
}

func (s *CreateTableRequestDataSourceConfig) GetTag() *string {
	return s.Tag
}

func (s *CreateTableRequestDataSourceConfig) SetAccessKey(v string) *CreateTableRequestDataSourceConfig {
	s.AccessKey = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetAccessSecret(v string) *CreateTableRequestDataSourceConfig {
	s.AccessSecret = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetBucket(v string) *CreateTableRequestDataSourceConfig {
	s.Bucket = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetCatalog(v string) *CreateTableRequestDataSourceConfig {
	s.Catalog = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetDatabase(v string) *CreateTableRequestDataSourceConfig {
	s.Database = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetEndpoint(v string) *CreateTableRequestDataSourceConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetOssPath(v string) *CreateTableRequestDataSourceConfig {
	s.OssPath = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetPartition(v string) *CreateTableRequestDataSourceConfig {
	s.Partition = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetProject(v string) *CreateTableRequestDataSourceConfig {
	s.Project = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetTable(v string) *CreateTableRequestDataSourceConfig {
	s.Table = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetTableFormat(v string) *CreateTableRequestDataSourceConfig {
	s.TableFormat = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) SetTag(v string) *CreateTableRequestDataSourceConfig {
	s.Tag = &v
	return s
}

func (s *CreateTableRequestDataSourceConfig) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestVectorIndex struct {
	// The configurations of the index schema.
	AdvanceParams *CreateTableRequestVectorIndexAdvanceParams `json:"advanceParams,omitempty" xml:"advanceParams,omitempty" type:"Struct"`
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
	// case_index
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

func (s CreateTableRequestVectorIndex) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestVectorIndex) GoString() string {
	return s.String()
}

func (s *CreateTableRequestVectorIndex) GetAdvanceParams() *CreateTableRequestVectorIndexAdvanceParams {
	return s.AdvanceParams
}

func (s *CreateTableRequestVectorIndex) GetDimension() *string {
	return s.Dimension
}

func (s *CreateTableRequestVectorIndex) GetDistanceType() *string {
	return s.DistanceType
}

func (s *CreateTableRequestVectorIndex) GetIndexName() *string {
	return s.IndexName
}

func (s *CreateTableRequestVectorIndex) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateTableRequestVectorIndex) GetSparseIndexField() *string {
	return s.SparseIndexField
}

func (s *CreateTableRequestVectorIndex) GetSparseValueField() *string {
	return s.SparseValueField
}

func (s *CreateTableRequestVectorIndex) GetVectorField() *string {
	return s.VectorField
}

func (s *CreateTableRequestVectorIndex) GetVectorIndexType() *string {
	return s.VectorIndexType
}

func (s *CreateTableRequestVectorIndex) SetAdvanceParams(v *CreateTableRequestVectorIndexAdvanceParams) *CreateTableRequestVectorIndex {
	s.AdvanceParams = v
	return s
}

func (s *CreateTableRequestVectorIndex) SetDimension(v string) *CreateTableRequestVectorIndex {
	s.Dimension = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetDistanceType(v string) *CreateTableRequestVectorIndex {
	s.DistanceType = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetIndexName(v string) *CreateTableRequestVectorIndex {
	s.IndexName = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetNamespace(v string) *CreateTableRequestVectorIndex {
	s.Namespace = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetSparseIndexField(v string) *CreateTableRequestVectorIndex {
	s.SparseIndexField = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetSparseValueField(v string) *CreateTableRequestVectorIndex {
	s.SparseValueField = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetVectorField(v string) *CreateTableRequestVectorIndex {
	s.VectorField = &v
	return s
}

func (s *CreateTableRequestVectorIndex) SetVectorIndexType(v string) *CreateTableRequestVectorIndex {
	s.VectorIndexType = &v
	return s
}

func (s *CreateTableRequestVectorIndex) Validate() error {
	return dara.Validate(s)
}

type CreateTableRequestVectorIndexAdvanceParams struct {
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

func (s CreateTableRequestVectorIndexAdvanceParams) String() string {
	return dara.Prettify(s)
}

func (s CreateTableRequestVectorIndexAdvanceParams) GoString() string {
	return s.String()
}

func (s *CreateTableRequestVectorIndexAdvanceParams) GetBuildIndexParams() *string {
	return s.BuildIndexParams
}

func (s *CreateTableRequestVectorIndexAdvanceParams) GetLinearBuildThreshold() *string {
	return s.LinearBuildThreshold
}

func (s *CreateTableRequestVectorIndexAdvanceParams) GetMinScanDocCnt() *string {
	return s.MinScanDocCnt
}

func (s *CreateTableRequestVectorIndexAdvanceParams) GetSearchIndexParams() *string {
	return s.SearchIndexParams
}

func (s *CreateTableRequestVectorIndexAdvanceParams) SetBuildIndexParams(v string) *CreateTableRequestVectorIndexAdvanceParams {
	s.BuildIndexParams = &v
	return s
}

func (s *CreateTableRequestVectorIndexAdvanceParams) SetLinearBuildThreshold(v string) *CreateTableRequestVectorIndexAdvanceParams {
	s.LinearBuildThreshold = &v
	return s
}

func (s *CreateTableRequestVectorIndexAdvanceParams) SetMinScanDocCnt(v string) *CreateTableRequestVectorIndexAdvanceParams {
	s.MinScanDocCnt = &v
	return s
}

func (s *CreateTableRequestVectorIndexAdvanceParams) SetSearchIndexParams(v string) *CreateTableRequestVectorIndexAdvanceParams {
	s.SearchIndexParams = &v
	return s
}

func (s *CreateTableRequestVectorIndexAdvanceParams) Validate() error {
	return dara.Validate(s)
}
