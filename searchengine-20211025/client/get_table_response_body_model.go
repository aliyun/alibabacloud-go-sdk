// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetTableResponseBody
	GetRequestId() *string
	SetResult(v *GetTableResponseBodyResult) *GetTableResponseBody
	GetResult() *GetTableResponseBodyResult
}

type GetTableResponseBody struct {
	// requestId
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
	Result *GetTableResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTableResponseBody) GetResult() *GetTableResponseBodyResult {
	return s.Result
}

func (s *GetTableResponseBody) SetRequestId(v string) *GetTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableResponseBody) SetResult(v *GetTableResponseBodyResult) *GetTableResponseBody {
	s.Result = v
	return s
}

func (s *GetTableResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyResult struct {
	// The configurations about field processing.
	DataProcessConfig []*GetTableResponseBodyResultDataProcessConfig `json:"dataProcessConfig,omitempty" xml:"dataProcessConfig,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DataProcessorCount *int32                                `json:"dataProcessorCount,omitempty" xml:"dataProcessorCount,omitempty"`
	DataSource         *GetTableResponseBodyResultDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// The field. The value is a key-value pair in which the key indicates the field name and value indicates the field type.
	FieldSchema map[string]*string `json:"fieldSchema,omitempty" xml:"fieldSchema,omitempty"`
	// example:
	//
	// test_oss
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PartitionCount *int32 `json:"partitionCount,omitempty" xml:"partitionCount,omitempty"`
	// example:
	//
	// id
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	// example:
	//
	// {}
	RawSchema *string `json:"rawSchema,omitempty" xml:"rawSchema,omitempty"`
	// The state of the index table. Valid values: NEW, PUBLISH, IN_USE, NOT_USE, STOP_USE, RESTORE_USE, and FAIL. After an index is created in an OpenSearch Retrieval Engine Edition instance, the index enters the IN_USE state. If the first full index fails to be created in an OpenSearch Vector Search Edition instance of the new version, the index is in the FAIL state.
	//
	// example:
	//
	// IN_USE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The index schema.
	VectorIndex []*GetTableResponseBodyResultVectorIndex `json:"vectorIndex,omitempty" xml:"vectorIndex,omitempty" type:"Repeated"`
}

func (s GetTableResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResult) GetDataProcessConfig() []*GetTableResponseBodyResultDataProcessConfig {
	return s.DataProcessConfig
}

func (s *GetTableResponseBodyResult) GetDataProcessorCount() *int32 {
	return s.DataProcessorCount
}

func (s *GetTableResponseBodyResult) GetDataSource() *GetTableResponseBodyResultDataSource {
	return s.DataSource
}

func (s *GetTableResponseBodyResult) GetFieldSchema() map[string]*string {
	return s.FieldSchema
}

func (s *GetTableResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetTableResponseBodyResult) GetPartitionCount() *int32 {
	return s.PartitionCount
}

func (s *GetTableResponseBodyResult) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GetTableResponseBodyResult) GetRawSchema() *string {
	return s.RawSchema
}

func (s *GetTableResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *GetTableResponseBodyResult) GetVectorIndex() []*GetTableResponseBodyResultVectorIndex {
	return s.VectorIndex
}

func (s *GetTableResponseBodyResult) SetDataProcessConfig(v []*GetTableResponseBodyResultDataProcessConfig) *GetTableResponseBodyResult {
	s.DataProcessConfig = v
	return s
}

func (s *GetTableResponseBodyResult) SetDataProcessorCount(v int32) *GetTableResponseBodyResult {
	s.DataProcessorCount = &v
	return s
}

func (s *GetTableResponseBodyResult) SetDataSource(v *GetTableResponseBodyResultDataSource) *GetTableResponseBodyResult {
	s.DataSource = v
	return s
}

func (s *GetTableResponseBodyResult) SetFieldSchema(v map[string]*string) *GetTableResponseBodyResult {
	s.FieldSchema = v
	return s
}

func (s *GetTableResponseBodyResult) SetName(v string) *GetTableResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetTableResponseBodyResult) SetPartitionCount(v int32) *GetTableResponseBodyResult {
	s.PartitionCount = &v
	return s
}

func (s *GetTableResponseBodyResult) SetPrimaryKey(v string) *GetTableResponseBodyResult {
	s.PrimaryKey = &v
	return s
}

func (s *GetTableResponseBodyResult) SetRawSchema(v string) *GetTableResponseBodyResult {
	s.RawSchema = &v
	return s
}

func (s *GetTableResponseBodyResult) SetStatus(v string) *GetTableResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetTableResponseBodyResult) SetVectorIndex(v []*GetTableResponseBodyResultVectorIndex) *GetTableResponseBodyResult {
	s.VectorIndex = v
	return s
}

func (s *GetTableResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyResultDataProcessConfig struct {
	// The destination field.
	//
	// example:
	//
	// source_image_vector
	DstField *string `json:"dstField,omitempty" xml:"dstField,omitempty"`
	// The method used to process the field. Valid values: copy and vectorize. A value of copy indicates that the value of the source field is copied to the destination field. A value of vectorize indicates that the value of the source field is vectorized by a vectorization model and the output vector is stored in the destination field.
	//
	// example:
	//
	// vectorize
	Operator *string `json:"operator,omitempty" xml:"operator,omitempty"`
	// The information about the model.
	Params *GetTableResponseBodyResultDataProcessConfigParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// The source field.
	//
	// example:
	//
	// source_image
	SrcField *string `json:"srcField,omitempty" xml:"srcField,omitempty"`
}

func (s GetTableResponseBodyResultDataProcessConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyResultDataProcessConfig) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataProcessConfig) GetDstField() *string {
	return s.DstField
}

func (s *GetTableResponseBodyResultDataProcessConfig) GetOperator() *string {
	return s.Operator
}

func (s *GetTableResponseBodyResultDataProcessConfig) GetParams() *GetTableResponseBodyResultDataProcessConfigParams {
	return s.Params
}

func (s *GetTableResponseBodyResultDataProcessConfig) GetSrcField() *string {
	return s.SrcField
}

func (s *GetTableResponseBodyResultDataProcessConfig) SetDstField(v string) *GetTableResponseBodyResultDataProcessConfig {
	s.DstField = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfig) SetOperator(v string) *GetTableResponseBodyResultDataProcessConfig {
	s.Operator = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfig) SetParams(v *GetTableResponseBodyResultDataProcessConfigParams) *GetTableResponseBodyResultDataProcessConfig {
	s.Params = v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfig) SetSrcField(v string) *GetTableResponseBodyResultDataProcessConfig {
	s.SrcField = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfig) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyResultDataProcessConfigParams struct {
	// The source of the data to be vectorized.
	SrcFieldConfig *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig `json:"srcFieldConfig,omitempty" xml:"srcFieldConfig,omitempty" type:"Struct"`
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

func (s GetTableResponseBodyResultDataProcessConfigParams) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyResultDataProcessConfigParams) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) GetSrcFieldConfig() *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig {
	return s.SrcFieldConfig
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) GetVectorModal() *string {
	return s.VectorModal
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) GetVectorModel() *string {
	return s.VectorModel
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) SetSrcFieldConfig(v *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) *GetTableResponseBodyResultDataProcessConfigParams {
	s.SrcFieldConfig = v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) SetVectorModal(v string) *GetTableResponseBodyResultDataProcessConfigParams {
	s.VectorModal = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) SetVectorModel(v string) *GetTableResponseBodyResultDataProcessConfigParams {
	s.VectorModel = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParams) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig struct {
	// OSS Bucket
	//
	// example:
	//
	// test
	OssBucket *string `json:"ossBucket,omitempty" xml:"ossBucket,omitempty"`
	// The Object Storage Service (OSS) endpoint.
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

func (s GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) GetOssBucket() *string {
	return s.OssBucket
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) GetOssEndpoint() *string {
	return s.OssEndpoint
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) GetUid() *string {
	return s.Uid
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) SetOssBucket(v string) *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig {
	s.OssBucket = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) SetOssEndpoint(v string) *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig {
	s.OssEndpoint = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) SetUid(v string) *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig {
	s.Uid = &v
	return s
}

func (s *GetTableResponseBodyResultDataProcessConfigParamsSrcFieldConfig) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyResultDataSource struct {
	// example:
	//
	// true
	AutoBuildIndex *bool                                       `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Config         *GetTableResponseBodyResultDataSourceConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// 1715160176
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetTableResponseBodyResultDataSource) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyResultDataSource) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataSource) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *GetTableResponseBodyResultDataSource) GetConfig() *GetTableResponseBodyResultDataSourceConfig {
	return s.Config
}

func (s *GetTableResponseBodyResultDataSource) GetDataTimeSec() *int32 {
	return s.DataTimeSec
}

func (s *GetTableResponseBodyResultDataSource) GetType() *string {
	return s.Type
}

func (s *GetTableResponseBodyResultDataSource) SetAutoBuildIndex(v bool) *GetTableResponseBodyResultDataSource {
	s.AutoBuildIndex = &v
	return s
}

func (s *GetTableResponseBodyResultDataSource) SetConfig(v *GetTableResponseBodyResultDataSourceConfig) *GetTableResponseBodyResultDataSource {
	s.Config = v
	return s
}

func (s *GetTableResponseBodyResultDataSource) SetDataTimeSec(v int32) *GetTableResponseBodyResultDataSource {
	s.DataTimeSec = &v
	return s
}

func (s *GetTableResponseBodyResultDataSource) SetType(v string) *GetTableResponseBodyResultDataSource {
	s.Type = &v
	return s
}

func (s *GetTableResponseBodyResultDataSource) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyResultDataSourceConfig struct {
	// AK
	//
	// example:
	//
	// ak
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// AS
	//
	// example:
	//
	// as
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// example:
	//
	// heytea-ops-oss
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// /opensearch_index_data/sift_oss_test.data
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// example:
	//
	// ds=20220808
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// vendor/sebastian/comparator/src/exceptions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// dp_pdm_marketing_prod
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// test_add
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
}

func (s GetTableResponseBodyResultDataSourceConfig) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyResultDataSourceConfig) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetBucket() *string {
	return s.Bucket
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetOssPath() *string {
	return s.OssPath
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetPartition() *string {
	return s.Partition
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetPath() *string {
	return s.Path
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetProject() *string {
	return s.Project
}

func (s *GetTableResponseBodyResultDataSourceConfig) GetTable() *string {
	return s.Table
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetAccessKey(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.AccessKey = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetAccessSecret(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.AccessSecret = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetBucket(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Bucket = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetEndpoint(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Endpoint = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetNamespace(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Namespace = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetOssPath(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.OssPath = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetPartition(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Partition = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetPath(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Path = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetProject(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Project = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) SetTable(v string) *GetTableResponseBodyResultDataSourceConfig {
	s.Table = &v
	return s
}

func (s *GetTableResponseBodyResultDataSourceConfig) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyResultVectorIndex struct {
	// The configurations of the index schema.
	AdvanceParams *GetTableResponseBodyResultVectorIndexAdvanceParams `json:"advanceParams,omitempty" xml:"advanceParams,omitempty" type:"Struct"`
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
	// test_odps
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

func (s GetTableResponseBodyResultVectorIndex) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyResultVectorIndex) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultVectorIndex) GetAdvanceParams() *GetTableResponseBodyResultVectorIndexAdvanceParams {
	return s.AdvanceParams
}

func (s *GetTableResponseBodyResultVectorIndex) GetDimension() *string {
	return s.Dimension
}

func (s *GetTableResponseBodyResultVectorIndex) GetDistanceType() *string {
	return s.DistanceType
}

func (s *GetTableResponseBodyResultVectorIndex) GetIndexName() *string {
	return s.IndexName
}

func (s *GetTableResponseBodyResultVectorIndex) GetNamespace() *string {
	return s.Namespace
}

func (s *GetTableResponseBodyResultVectorIndex) GetSparseIndexField() *string {
	return s.SparseIndexField
}

func (s *GetTableResponseBodyResultVectorIndex) GetSparseValueField() *string {
	return s.SparseValueField
}

func (s *GetTableResponseBodyResultVectorIndex) GetVectorField() *string {
	return s.VectorField
}

func (s *GetTableResponseBodyResultVectorIndex) GetVectorIndexType() *string {
	return s.VectorIndexType
}

func (s *GetTableResponseBodyResultVectorIndex) SetAdvanceParams(v *GetTableResponseBodyResultVectorIndexAdvanceParams) *GetTableResponseBodyResultVectorIndex {
	s.AdvanceParams = v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetDimension(v string) *GetTableResponseBodyResultVectorIndex {
	s.Dimension = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetDistanceType(v string) *GetTableResponseBodyResultVectorIndex {
	s.DistanceType = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetIndexName(v string) *GetTableResponseBodyResultVectorIndex {
	s.IndexName = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetNamespace(v string) *GetTableResponseBodyResultVectorIndex {
	s.Namespace = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetSparseIndexField(v string) *GetTableResponseBodyResultVectorIndex {
	s.SparseIndexField = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetSparseValueField(v string) *GetTableResponseBodyResultVectorIndex {
	s.SparseValueField = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetVectorField(v string) *GetTableResponseBodyResultVectorIndex {
	s.VectorField = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) SetVectorIndexType(v string) *GetTableResponseBodyResultVectorIndex {
	s.VectorIndexType = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndex) Validate() error {
	return dara.Validate(s)
}

type GetTableResponseBodyResultVectorIndexAdvanceParams struct {
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

func (s GetTableResponseBodyResultVectorIndexAdvanceParams) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponseBodyResultVectorIndexAdvanceParams) GoString() string {
	return s.String()
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) GetBuildIndexParams() *string {
	return s.BuildIndexParams
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) GetLinearBuildThreshold() *string {
	return s.LinearBuildThreshold
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) GetMinScanDocCnt() *string {
	return s.MinScanDocCnt
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) GetSearchIndexParams() *string {
	return s.SearchIndexParams
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) SetBuildIndexParams(v string) *GetTableResponseBodyResultVectorIndexAdvanceParams {
	s.BuildIndexParams = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) SetLinearBuildThreshold(v string) *GetTableResponseBodyResultVectorIndexAdvanceParams {
	s.LinearBuildThreshold = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) SetMinScanDocCnt(v string) *GetTableResponseBodyResultVectorIndexAdvanceParams {
	s.MinScanDocCnt = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) SetSearchIndexParams(v string) *GetTableResponseBodyResultVectorIndexAdvanceParams {
	s.SearchIndexParams = &v
	return s
}

func (s *GetTableResponseBodyResultVectorIndexAdvanceParams) Validate() error {
	return dara.Validate(s)
}
