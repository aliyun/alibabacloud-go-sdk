// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBuildParallelNum(v int32) *ModifyIndexRequest
	GetBuildParallelNum() *int32
	SetCluster(v map[string]map[string]interface{}) *ModifyIndexRequest
	GetCluster() map[string]map[string]interface{}
	SetClusterConfigName(v string) *ModifyIndexRequest
	GetClusterConfigName() *string
	SetConfig(v map[string]*ConfigValue) *ModifyIndexRequest
	GetConfig() map[string]*ConfigValue
	SetContent(v string) *ModifyIndexRequest
	GetContent() *string
	SetDataSource(v string) *ModifyIndexRequest
	GetDataSource() *string
	SetDataSourceInfo(v *ModifyIndexRequestDataSourceInfo) *ModifyIndexRequest
	GetDataSourceInfo() *ModifyIndexRequestDataSourceInfo
	SetDescription(v string) *ModifyIndexRequest
	GetDescription() *string
	SetDomain(v string) *ModifyIndexRequest
	GetDomain() *string
	SetExtend(v map[string]interface{}) *ModifyIndexRequest
	GetExtend() map[string]interface{}
	SetMergeParallelNum(v int32) *ModifyIndexRequest
	GetMergeParallelNum() *int32
	SetPartition(v int32) *ModifyIndexRequest
	GetPartition() *int32
	SetPushMode(v string) *ModifyIndexRequest
	GetPushMode() *string
	SetDryRun(v bool) *ModifyIndexRequest
	GetDryRun() *bool
}

type ModifyIndexRequest struct {
	// The maximum number of full indexes that can be concurrently built.
	//
	// example:
	//
	// 2
	BuildParallelNum *int32 `json:"buildParallelNum,omitempty" xml:"buildParallelNum,omitempty"`
	// The cluster information.
	Cluster map[string]map[string]interface{} `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The name of the configuration file.
	//
	// example:
	//
	// ha-cn-35t3r02iq03@ha-cn-35t3r02iq03_test_api@hz_pre_vpc_domain_1@test_api@index_config_v1
	ClusterConfigName *string `json:"clusterConfigName,omitempty" xml:"clusterConfigName,omitempty"`
	// The information about the offline configuration.
	Config map[string]*ConfigValue `json:"config,omitempty" xml:"config,omitempty"`
	// The file content.
	//
	// example:
	//
	// {\\"summarys\\":{\\"summary_fields\\":[\\"id\\"]},\\"indexs\\":[{\\"index_name\\":\\"index_id\\",\\"index_type\\":\\"PRIMARYKEY64\\",\\"index_fields\\":\\"id\\",\\"has_primary_key_attribute\\":true,\\"is_primary_key_sorted\\":false}],\\"attributes\\":[\\"id\\"],\\"fields\\":[{\\"field_name\\":\\"id\\",\\"field_type\\":\\"UINT16\\"}],\\"table_name\\":\\"index_2\\"}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// ha-cn-35t3n1yuj0d_index_1
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// The information about the data source, which is required for the new version of OpenSearch Vector Search Edition.
	DataSourceInfo *ModifyIndexRequestDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// The description of the data source.
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the data center in which the data source is deployed.
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string                `json:"domain,omitempty" xml:"domain,omitempty"`
	Extend map[string]interface{} `json:"extend,omitempty" xml:"extend,omitempty"`
	// The maximum number of full indexes that can be concurrently merged.
	//
	// example:
	//
	// 2
	MergeParallelNum *int32 `json:"mergeParallelNum,omitempty" xml:"mergeParallelNum,omitempty"`
	// The number of shards.
	//
	// example:
	//
	// 2
	Partition *int32 `json:"partition,omitempty" xml:"partition,omitempty"`
	// The push mode of the configuration. By default, only the configuration is pushed.
	//
	// example:
	//
	// PUSH_ONLY
	PushMode *string `json:"pushMode,omitempty" xml:"pushMode,omitempty"`
	// Specifies whether to check the validity of input parameters. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**: checks only the validity of input parameters.
	//
	// 	- **false**: checks the validity of input parameters and creates an attribution configuration.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexRequest) GoString() string {
	return s.String()
}

func (s *ModifyIndexRequest) GetBuildParallelNum() *int32 {
	return s.BuildParallelNum
}

func (s *ModifyIndexRequest) GetCluster() map[string]map[string]interface{} {
	return s.Cluster
}

func (s *ModifyIndexRequest) GetClusterConfigName() *string {
	return s.ClusterConfigName
}

func (s *ModifyIndexRequest) GetConfig() map[string]*ConfigValue {
	return s.Config
}

func (s *ModifyIndexRequest) GetContent() *string {
	return s.Content
}

func (s *ModifyIndexRequest) GetDataSource() *string {
	return s.DataSource
}

func (s *ModifyIndexRequest) GetDataSourceInfo() *ModifyIndexRequestDataSourceInfo {
	return s.DataSourceInfo
}

func (s *ModifyIndexRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyIndexRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyIndexRequest) GetExtend() map[string]interface{} {
	return s.Extend
}

func (s *ModifyIndexRequest) GetMergeParallelNum() *int32 {
	return s.MergeParallelNum
}

func (s *ModifyIndexRequest) GetPartition() *int32 {
	return s.Partition
}

func (s *ModifyIndexRequest) GetPushMode() *string {
	return s.PushMode
}

func (s *ModifyIndexRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyIndexRequest) SetBuildParallelNum(v int32) *ModifyIndexRequest {
	s.BuildParallelNum = &v
	return s
}

func (s *ModifyIndexRequest) SetCluster(v map[string]map[string]interface{}) *ModifyIndexRequest {
	s.Cluster = v
	return s
}

func (s *ModifyIndexRequest) SetClusterConfigName(v string) *ModifyIndexRequest {
	s.ClusterConfigName = &v
	return s
}

func (s *ModifyIndexRequest) SetConfig(v map[string]*ConfigValue) *ModifyIndexRequest {
	s.Config = v
	return s
}

func (s *ModifyIndexRequest) SetContent(v string) *ModifyIndexRequest {
	s.Content = &v
	return s
}

func (s *ModifyIndexRequest) SetDataSource(v string) *ModifyIndexRequest {
	s.DataSource = &v
	return s
}

func (s *ModifyIndexRequest) SetDataSourceInfo(v *ModifyIndexRequestDataSourceInfo) *ModifyIndexRequest {
	s.DataSourceInfo = v
	return s
}

func (s *ModifyIndexRequest) SetDescription(v string) *ModifyIndexRequest {
	s.Description = &v
	return s
}

func (s *ModifyIndexRequest) SetDomain(v string) *ModifyIndexRequest {
	s.Domain = &v
	return s
}

func (s *ModifyIndexRequest) SetExtend(v map[string]interface{}) *ModifyIndexRequest {
	s.Extend = v
	return s
}

func (s *ModifyIndexRequest) SetMergeParallelNum(v int32) *ModifyIndexRequest {
	s.MergeParallelNum = &v
	return s
}

func (s *ModifyIndexRequest) SetPartition(v int32) *ModifyIndexRequest {
	s.Partition = &v
	return s
}

func (s *ModifyIndexRequest) SetPushMode(v string) *ModifyIndexRequest {
	s.PushMode = &v
	return s
}

func (s *ModifyIndexRequest) SetDryRun(v bool) *ModifyIndexRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyIndexRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyIndexRequestDataSourceInfo struct {
	// Specifies whether to enable the automatic full indexing feature.
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The reindexing method. Valid values: api: API data source. indexRecover: data recovery by using indexing.
	//
	// example:
	//
	// api
	BuildMode *string `json:"buildMode,omitempty" xml:"buildMode,omitempty"`
	// The configurations of the MaxCompute data source.
	Config *ModifyIndexRequestDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The start timestamp from which incremental data is retrieved.
	//
	// example:
	//
	// 1709715164
	DataTimeSec *int32 `json:"dataTimeSec,omitempty" xml:"dataTimeSec,omitempty"`
	// The offline deployment name of the data source.
	//
	// example:
	//
	// vpc_hz_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The ID of the index version from which data is restored.
	//
	// example:
	//
	// 4
	Generation *int64 `json:"generation,omitempty" xml:"generation,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// ha-cn-35t3n1yuj0d_index_1
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	OssDataPath *string `json:"ossDataPath,omitempty" xml:"ossDataPath,omitempty"`
	Partition   *string `json:"partition,omitempty" xml:"partition,omitempty"`
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
	SaroConfig *ModifyIndexRequestDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// The type of the data source. Valid values: odps, swift, saro, oss, and unKnow.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModifyIndexRequestDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexRequestDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ModifyIndexRequestDataSourceInfo) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *ModifyIndexRequestDataSourceInfo) GetBuildMode() *string {
	return s.BuildMode
}

func (s *ModifyIndexRequestDataSourceInfo) GetConfig() *ModifyIndexRequestDataSourceInfoConfig {
	return s.Config
}

func (s *ModifyIndexRequestDataSourceInfo) GetDataTimeSec() *int32 {
	return s.DataTimeSec
}

func (s *ModifyIndexRequestDataSourceInfo) GetDomain() *string {
	return s.Domain
}

func (s *ModifyIndexRequestDataSourceInfo) GetGeneration() *int64 {
	return s.Generation
}

func (s *ModifyIndexRequestDataSourceInfo) GetName() *string {
	return s.Name
}

func (s *ModifyIndexRequestDataSourceInfo) GetOssDataPath() *string {
	return s.OssDataPath
}

func (s *ModifyIndexRequestDataSourceInfo) GetPartition() *string {
	return s.Partition
}

func (s *ModifyIndexRequestDataSourceInfo) GetProcessParallelNum() *int32 {
	return s.ProcessParallelNum
}

func (s *ModifyIndexRequestDataSourceInfo) GetProcessPartitionCount() *int32 {
	return s.ProcessPartitionCount
}

func (s *ModifyIndexRequestDataSourceInfo) GetSaroConfig() *ModifyIndexRequestDataSourceInfoSaroConfig {
	return s.SaroConfig
}

func (s *ModifyIndexRequestDataSourceInfo) GetType() *string {
	return s.Type
}

func (s *ModifyIndexRequestDataSourceInfo) SetAutoBuildIndex(v bool) *ModifyIndexRequestDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetBuildMode(v string) *ModifyIndexRequestDataSourceInfo {
	s.BuildMode = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetConfig(v *ModifyIndexRequestDataSourceInfoConfig) *ModifyIndexRequestDataSourceInfo {
	s.Config = v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetDataTimeSec(v int32) *ModifyIndexRequestDataSourceInfo {
	s.DataTimeSec = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetDomain(v string) *ModifyIndexRequestDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetGeneration(v int64) *ModifyIndexRequestDataSourceInfo {
	s.Generation = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetName(v string) *ModifyIndexRequestDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetOssDataPath(v string) *ModifyIndexRequestDataSourceInfo {
	s.OssDataPath = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetPartition(v string) *ModifyIndexRequestDataSourceInfo {
	s.Partition = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetProcessParallelNum(v int32) *ModifyIndexRequestDataSourceInfo {
	s.ProcessParallelNum = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetProcessPartitionCount(v int32) *ModifyIndexRequestDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetSaroConfig(v *ModifyIndexRequestDataSourceInfoSaroConfig) *ModifyIndexRequestDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) SetType(v string) *ModifyIndexRequestDataSourceInfo {
	s.Type = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfo) Validate() error {
	return dara.Validate(s)
}

type ModifyIndexRequestDataSourceInfoConfig struct {
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
	// The endpoint of the MaxCompute data source.
	//
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Format   *string `json:"format,omitempty" xml:"format,omitempty"`
	// The namespace. This parameter is applicable to the SARO data source used in the intranet of Alibaba Group.
	//
	// example:
	//
	// test-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The Object Storage Service (OSS) path.
	//
	// example:
	//
	// /opensearch/oss.json
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The partition in the MaxCompute table. Example: ds=20180102.
	//
	// example:
	//
	// ds=20230114
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The file path in the Apsara File Storage for HDFS file system.
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
	// The name of the MaxCompute table that is used as the data source.
	//
	// example:
	//
	// item
	Table       *string `json:"table,omitempty" xml:"table,omitempty"`
	TableFormat *string `json:"tableFormat,omitempty" xml:"tableFormat,omitempty"`
	Tag         *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ModifyIndexRequestDataSourceInfoConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexRequestDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetBucket() *string {
	return s.Bucket
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetDatabase() *string {
	return s.Database
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetFormat() *string {
	return s.Format
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetOssPath() *string {
	return s.OssPath
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetPartition() *string {
	return s.Partition
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetPath() *string {
	return s.Path
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetProject() *string {
	return s.Project
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetTable() *string {
	return s.Table
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetTableFormat() *string {
	return s.TableFormat
}

func (s *ModifyIndexRequestDataSourceInfoConfig) GetTag() *string {
	return s.Tag
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetAccessKey(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetAccessSecret(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetBucket(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetCatalog(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Catalog = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetDatabase(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Database = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetEndpoint(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetFormat(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Format = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetNamespace(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetOssPath(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetPartition(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetPath(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetProject(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetTable(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Table = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetTableFormat(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.TableFormat = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) SetTag(v string) *ModifyIndexRequestDataSourceInfoConfig {
	s.Tag = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyIndexRequestDataSourceInfoSaroConfig struct {
	// The namespace to which the SARO data source belongs.
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

func (s ModifyIndexRequestDataSourceInfoSaroConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyIndexRequestDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *ModifyIndexRequestDataSourceInfoSaroConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyIndexRequestDataSourceInfoSaroConfig) GetTableName() *string {
	return s.TableName
}

func (s *ModifyIndexRequestDataSourceInfoSaroConfig) SetNamespace(v string) *ModifyIndexRequestDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoSaroConfig) SetTableName(v string) *ModifyIndexRequestDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

func (s *ModifyIndexRequestDataSourceInfoSaroConfig) Validate() error {
	return dara.Validate(s)
}
