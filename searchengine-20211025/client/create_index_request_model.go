// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildParallelNum(v int32) *CreateIndexRequest
	GetBuildParallelNum() *int32
	SetContent(v string) *CreateIndexRequest
	GetContent() *string
	SetDataSource(v string) *CreateIndexRequest
	GetDataSource() *string
	SetDataSourceInfo(v *CreateIndexRequestDataSourceInfo) *CreateIndexRequest
	GetDataSourceInfo() *CreateIndexRequestDataSourceInfo
	SetDomain(v string) *CreateIndexRequest
	GetDomain() *string
	SetExtend(v map[string]interface{}) *CreateIndexRequest
	GetExtend() map[string]interface{}
	SetMergeParallelNum(v int32) *CreateIndexRequest
	GetMergeParallelNum() *int32
	SetName(v string) *CreateIndexRequest
	GetName() *string
	SetPartition(v int32) *CreateIndexRequest
	GetPartition() *int32
	SetDryRun(v bool) *CreateIndexRequest
	GetDryRun() *bool
}

type CreateIndexRequest struct {
	// The maximum number of full indexes that can be concurrently built.
	//
	// example:
	//
	// 2
	BuildParallelNum *int32 `json:"buildParallelNum,omitempty" xml:"buildParallelNum,omitempty"`
	// The index schema.
	//
	// example:
	//
	// {\\"summarys\\":{\\"summary_fields\\":[\\"id\\"]},\\"indexs\\":[{\\"index_name\\":\\"index_id\\",\\"index_type\\":\\"PRIMARYKEY64\\",\\"index_fields\\":\\"id\\",\\"has_primary_key_attribute\\":true,\\"is_primary_key_sorted\\":false}],\\"attributes\\":[\\"id\\"],\\"fields\\":[{\\"field_name\\":\\"id\\",\\"field_type\\":\\"UINT16\\"}],\\"table_name\\":\\"index_2\\"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test1
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// The information about the data source. This parameter is required for an OpenSearch Vector Search Edition instance of the new version.
	DataSourceInfo *CreateIndexRequestDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// The data center in which the data source is deployed.
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The extended content of the field configuration. key specifies the vector field and the field that requires embedding.
	//
	// example:
	//
	// {
	//
	//         "vector":
	//
	//         [
	//
	//             "source_image_vector"
	//
	//         ],
	//
	//         "embeding":
	//
	//         [
	//
	//             "source_image"
	//
	//         ],
	//
	//         "description":
	//
	//         []
	//
	//     }
	Extend map[string]interface{} `json:"extend,omitempty" xml:"extend,omitempty"`
	// The maximum number of full indexes that can be concurrently merged.
	//
	// example:
	//
	// 2
	MergeParallelNum *int32 `json:"mergeParallelNum,omitempty" xml:"mergeParallelNum,omitempty"`
	// The index name.
	//
	// example:
	//
	// ha-cn-zvp2qr1sk01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of data shards.
	//
	// example:
	//
	// 20211202
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
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

func (s CreateIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequest) GoString() string {
	return s.String()
}

func (s *CreateIndexRequest) GetBuildParallelNum() *int32 {
	return s.BuildParallelNum
}

func (s *CreateIndexRequest) GetContent() *string {
	return s.Content
}

func (s *CreateIndexRequest) GetDataSource() *string {
	return s.DataSource
}

func (s *CreateIndexRequest) GetDataSourceInfo() *CreateIndexRequestDataSourceInfo {
	return s.DataSourceInfo
}

func (s *CreateIndexRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateIndexRequest) GetExtend() map[string]interface{} {
	return s.Extend
}

func (s *CreateIndexRequest) GetMergeParallelNum() *int32 {
	return s.MergeParallelNum
}

func (s *CreateIndexRequest) GetName() *string {
	return s.Name
}

func (s *CreateIndexRequest) GetPartition() *int32 {
	return s.Partition
}

func (s *CreateIndexRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateIndexRequest) SetBuildParallelNum(v int32) *CreateIndexRequest {
	s.BuildParallelNum = &v
	return s
}

func (s *CreateIndexRequest) SetContent(v string) *CreateIndexRequest {
	s.Content = &v
	return s
}

func (s *CreateIndexRequest) SetDataSource(v string) *CreateIndexRequest {
	s.DataSource = &v
	return s
}

func (s *CreateIndexRequest) SetDataSourceInfo(v *CreateIndexRequestDataSourceInfo) *CreateIndexRequest {
	s.DataSourceInfo = v
	return s
}

func (s *CreateIndexRequest) SetDomain(v string) *CreateIndexRequest {
	s.Domain = &v
	return s
}

func (s *CreateIndexRequest) SetExtend(v map[string]interface{}) *CreateIndexRequest {
	s.Extend = v
	return s
}

func (s *CreateIndexRequest) SetMergeParallelNum(v int32) *CreateIndexRequest {
	s.MergeParallelNum = &v
	return s
}

func (s *CreateIndexRequest) SetName(v string) *CreateIndexRequest {
	s.Name = &v
	return s
}

func (s *CreateIndexRequest) SetPartition(v int32) *CreateIndexRequest {
	s.Partition = &v
	return s
}

func (s *CreateIndexRequest) SetDryRun(v bool) *CreateIndexRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIndexRequest) Validate() error {
	return dara.Validate(s)
}

type CreateIndexRequestDataSourceInfo struct {
	// Specifies whether to enable automatic full indexing.
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The information about the MaxCompute data source.
	Config *CreateIndexRequestDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The start timestamp from which incremental data is retrieved.
	//
	// example:
	//
	// 1709715164
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The data center in which the data source is deployed.
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// ha-cn-35t3n1yuj0d_index_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The maximum number of full indexes that can be concurrently processed.
	//
	// example:
	//
	// 2
	ProcessParallelNum *int32 `json:"processParallelNum,omitempty" xml:"processParallelNum,omitempty"`
	// The number of resources used for data update.
	//
	// example:
	//
	// 4
	ProcessPartitionCount *int32 `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	// The configurations of the SARO data source.
	SaroConfig *CreateIndexRequestDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	Scene      *string                                     `json:"scene,omitempty" xml:"scene,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- odps
	//
	// 	- swift
	//
	// 	- saro
	//
	// 	- oss
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateIndexRequestDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequestDataSourceInfo) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSourceInfo) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *CreateIndexRequestDataSourceInfo) GetConfig() *CreateIndexRequestDataSourceInfoConfig {
	return s.Config
}

func (s *CreateIndexRequestDataSourceInfo) GetDataTimeSec() *int32 {
	return s.DataTimeSec
}

func (s *CreateIndexRequestDataSourceInfo) GetDomain() *string {
	return s.Domain
}

func (s *CreateIndexRequestDataSourceInfo) GetName() *string {
	return s.Name
}

func (s *CreateIndexRequestDataSourceInfo) GetProcessParallelNum() *int32 {
	return s.ProcessParallelNum
}

func (s *CreateIndexRequestDataSourceInfo) GetProcessPartitionCount() *int32 {
	return s.ProcessPartitionCount
}

func (s *CreateIndexRequestDataSourceInfo) GetSaroConfig() *CreateIndexRequestDataSourceInfoSaroConfig {
	return s.SaroConfig
}

func (s *CreateIndexRequestDataSourceInfo) GetScene() *string {
	return s.Scene
}

func (s *CreateIndexRequestDataSourceInfo) GetType() *string {
	return s.Type
}

func (s *CreateIndexRequestDataSourceInfo) SetAutoBuildIndex(v bool) *CreateIndexRequestDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetConfig(v *CreateIndexRequestDataSourceInfoConfig) *CreateIndexRequestDataSourceInfo {
	s.Config = v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetDataTimeSec(v int32) *CreateIndexRequestDataSourceInfo {
	s.DataTimeSec = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetDomain(v string) *CreateIndexRequestDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetName(v string) *CreateIndexRequestDataSourceInfo {
	s.Name = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetProcessParallelNum(v int32) *CreateIndexRequestDataSourceInfo {
	s.ProcessParallelNum = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetProcessPartitionCount(v int32) *CreateIndexRequestDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetSaroConfig(v *CreateIndexRequestDataSourceInfoSaroConfig) *CreateIndexRequestDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetScene(v string) *CreateIndexRequestDataSourceInfo {
	s.Scene = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) SetType(v string) *CreateIndexRequestDataSourceInfo {
	s.Type = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfo) Validate() error {
	return dara.Validate(s)
}

type CreateIndexRequestDataSourceInfoConfig struct {
	// The AccessKey ID of the MaxCompute data source.
	//
	// example:
	//
	// L***p
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// The AccessKey secret of the MaxCompute data source.
	//
	// example:
	//
	// 5**9a6
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// test-bucket
	Bucket   *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Catalog  *string `json:"catalog,omitempty" xml:"catalog,omitempty"`
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	// The endpoint of the MaxCompute or Object Storage Service (OSS) data source.
	//
	// example:
	//
	// https://oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Format   *string `json:"format,omitempty" xml:"format,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// test-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The path of the OSS object.
	//
	// example:
	//
	// /opensearch/oss.json
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The partition in the MaxCompute table. This parameter is required if type is set to odps.
	//
	// example:
	//
	// ds=20230114
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The path of the Apsara File Storage for HDFS data source.
	//
	// example:
	//
	// test-hdfs-path
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The name of the MaxCompute project that is used as the data source.
	//
	// example:
	//
	// bbt_algo_pai
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The table name.
	//
	// example:
	//
	// bbt_rec_swing_u2i2i_score_be_v1
	Table       *string `json:"table,omitempty" xml:"table,omitempty"`
	TableFormat *string `json:"tableFormat,omitempty" xml:"tableFormat,omitempty"`
	Tag         *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s CreateIndexRequestDataSourceInfoConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequestDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetBucket() *string {
	return s.Bucket
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetDatabase() *string {
	return s.Database
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetFormat() *string {
	return s.Format
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetOssPath() *string {
	return s.OssPath
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetPartition() *string {
	return s.Partition
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetPath() *string {
	return s.Path
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetProject() *string {
	return s.Project
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetTable() *string {
	return s.Table
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetTableFormat() *string {
	return s.TableFormat
}

func (s *CreateIndexRequestDataSourceInfoConfig) GetTag() *string {
	return s.Tag
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetAccessKey(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetAccessSecret(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetBucket(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetCatalog(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Catalog = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetDatabase(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Database = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetEndpoint(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetFormat(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Format = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetNamespace(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetOssPath(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetPartition(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetPath(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetProject(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetTable(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Table = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetTableFormat(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.TableFormat = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) SetTag(v string) *CreateIndexRequestDataSourceInfoConfig {
	s.Tag = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoConfig) Validate() error {
	return dara.Validate(s)
}

type CreateIndexRequestDataSourceInfoSaroConfig struct {
	// The namespace of the SARO data source.
	//
	// example:
	//
	// flink-test-fjx-default
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the SARO table.
	//
	// example:
	//
	// device_event_shy_summary_
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s CreateIndexRequestDataSourceInfoSaroConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateIndexRequestDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *CreateIndexRequestDataSourceInfoSaroConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateIndexRequestDataSourceInfoSaroConfig) GetTableName() *string {
	return s.TableName
}

func (s *CreateIndexRequestDataSourceInfoSaroConfig) SetNamespace(v string) *CreateIndexRequestDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoSaroConfig) SetTableName(v string) *CreateIndexRequestDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

func (s *CreateIndexRequestDataSourceInfoSaroConfig) Validate() error {
	return dara.Validate(s)
}
