// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetIndexResponseBody
	GetRequestId() *string
	SetResult(v *GetIndexResponseBodyResult) *GetIndexResponseBody
	GetResult() *GetIndexResponseBodyResult
}

type GetIndexResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4FB0325E-8C37-5525-96AC-0333523170A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The index information.
	Result *GetIndexResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIndexResponseBody) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIndexResponseBody) GetResult() *GetIndexResponseBodyResult {
	return s.Result
}

func (s *GetIndexResponseBody) SetRequestId(v string) *GetIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIndexResponseBody) SetResult(v *GetIndexResponseBodyResult) *GetIndexResponseBody {
	s.Result = v
	return s
}

func (s *GetIndexResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetIndexResponseBodyResult struct {
	// The cluster information.
	Cluster map[string]*ResultClusterValue `json:"cluster,omitempty" xml:"cluster,omitempty"`
	// The configuration information.
	Config map[string]map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The configuration that takes effect next time.
	ConfigWhenBuild map[string]map[string]interface{} `json:"configWhenBuild,omitempty" xml:"configWhenBuild,omitempty"`
	// The file content.
	//
	// example:
	//
	// {"summarys":{"parameter":{"file_compressor":"zstd"},"summary_fields":["id"]},"file_compress":[{"name":"file_compressor","type":"zstd"},{"name":"no_compressor","type":""}],"indexs":[{"index_fields":"name","index_name":"ids","index_type":"STRING"},{"has_primary_key_attribute":true,"index_fields":"id","is_primary_key_sorted":false,"index_name":"id","index_type":"PRIMARYKEY64"}],"attributes":[{"file_compress":"no_compressor","field_name":"id"}],"fields":[{"user_defined_param":{},"compress_type":"uniq","field_type":"STRING","field_name":"id"},{"compress_type":"uniq","field_type":"STRING","field_name":"name"}],"table_name":"api"}
	Content    *string `json:"content,omitempty" xml:"content,omitempty"`
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// ha-cn-tl32nd2nq01_00
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// The information about the data source.
	DataSourceInfo *GetIndexResponseBodyResultDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// The description of the index version.
	//
	// example:
	//
	// test index
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The deployment name of the index.
	//
	// example:
	//
	// sz_vpc_domain_1
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// Extended information
	Extend map[string][]*string `json:"extend,omitempty" xml:"extend,omitempty"`
	// The time when full data in the index was last updated.
	//
	// example:
	//
	// 2024-06-20 08:52:54
	FullUpdateTime *string `json:"fullUpdateTime,omitempty" xml:"fullUpdateTime,omitempty"`
	// The data version.
	//
	// example:
	//
	// 1688523414
	FullVersion *int64 `json:"fullVersion,omitempty" xml:"fullVersion,omitempty"`
	// The time when incremental data in the index was last updated.
	//
	// example:
	//
	// 2024-06-20 08:52:54
	IncUpdateTime *string `json:"incUpdateTime,omitempty" xml:"incUpdateTime,omitempty"`
	// The index size.
	//
	// example:
	//
	// 4689
	IndexSize *int64 `json:"indexSize,omitempty" xml:"indexSize,omitempty"`
	// The status of the index version. Valid values:
	//
	// 	- NEW: The index version is created.
	//
	// 	- PUBLISH: The index version is normal.
	//
	// 	- IN_USE: The index version is in use.
	//
	// 	- NOT_USE: The index version is not used.
	//
	// 	- STOP_USE: The index version is being stopped.
	//
	// 	- RESTORE_USE: The index version is being restored.
	//
	// 	- FAIL: The index version failed to be created.
	//
	// example:
	//
	// IN_USE
	IndexStatus *string `json:"indexStatus,omitempty" xml:"indexStatus,omitempty"`
	// The index name.
	//
	// example:
	//
	// general
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of shards.
	//
	// example:
	//
	// 2
	Partition  *int32  `json:"partition,omitempty" xml:"partition,omitempty"`
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The information about the versions.
	Versions []*GetIndexResponseBodyResultVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetIndexResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetIndexResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResult) GetCluster() map[string]*ResultClusterValue {
	return s.Cluster
}

func (s *GetIndexResponseBodyResult) GetConfig() map[string]map[string]interface{} {
	return s.Config
}

func (s *GetIndexResponseBodyResult) GetConfigWhenBuild() map[string]map[string]interface{} {
	return s.ConfigWhenBuild
}

func (s *GetIndexResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *GetIndexResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetIndexResponseBodyResult) GetDataSource() *string {
	return s.DataSource
}

func (s *GetIndexResponseBodyResult) GetDataSourceInfo() *GetIndexResponseBodyResultDataSourceInfo {
	return s.DataSourceInfo
}

func (s *GetIndexResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *GetIndexResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *GetIndexResponseBodyResult) GetExtend() map[string][]*string {
	return s.Extend
}

func (s *GetIndexResponseBodyResult) GetFullUpdateTime() *string {
	return s.FullUpdateTime
}

func (s *GetIndexResponseBodyResult) GetFullVersion() *int64 {
	return s.FullVersion
}

func (s *GetIndexResponseBodyResult) GetIncUpdateTime() *string {
	return s.IncUpdateTime
}

func (s *GetIndexResponseBodyResult) GetIndexSize() *int64 {
	return s.IndexSize
}

func (s *GetIndexResponseBodyResult) GetIndexStatus() *string {
	return s.IndexStatus
}

func (s *GetIndexResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetIndexResponseBodyResult) GetPartition() *int32 {
	return s.Partition
}

func (s *GetIndexResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetIndexResponseBodyResult) GetVersions() []*GetIndexResponseBodyResultVersions {
	return s.Versions
}

func (s *GetIndexResponseBodyResult) SetCluster(v map[string]*ResultClusterValue) *GetIndexResponseBodyResult {
	s.Cluster = v
	return s
}

func (s *GetIndexResponseBodyResult) SetConfig(v map[string]map[string]interface{}) *GetIndexResponseBodyResult {
	s.Config = v
	return s
}

func (s *GetIndexResponseBodyResult) SetConfigWhenBuild(v map[string]map[string]interface{}) *GetIndexResponseBodyResult {
	s.ConfigWhenBuild = v
	return s
}

func (s *GetIndexResponseBodyResult) SetContent(v string) *GetIndexResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetCreateTime(v string) *GetIndexResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetDataSource(v string) *GetIndexResponseBodyResult {
	s.DataSource = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetDataSourceInfo(v *GetIndexResponseBodyResultDataSourceInfo) *GetIndexResponseBodyResult {
	s.DataSourceInfo = v
	return s
}

func (s *GetIndexResponseBodyResult) SetDescription(v string) *GetIndexResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetDomain(v string) *GetIndexResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetExtend(v map[string][]*string) *GetIndexResponseBodyResult {
	s.Extend = v
	return s
}

func (s *GetIndexResponseBodyResult) SetFullUpdateTime(v string) *GetIndexResponseBodyResult {
	s.FullUpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetFullVersion(v int64) *GetIndexResponseBodyResult {
	s.FullVersion = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetIncUpdateTime(v string) *GetIndexResponseBodyResult {
	s.IncUpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetIndexSize(v int64) *GetIndexResponseBodyResult {
	s.IndexSize = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetIndexStatus(v string) *GetIndexResponseBodyResult {
	s.IndexStatus = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetName(v string) *GetIndexResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetPartition(v int32) *GetIndexResponseBodyResult {
	s.Partition = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetUpdateTime(v string) *GetIndexResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResult) SetVersions(v []*GetIndexResponseBodyResultVersions) *GetIndexResponseBodyResult {
	s.Versions = v
	return s
}

func (s *GetIndexResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetIndexResponseBodyResultDataSourceInfo struct {
	// Indicates whether the automatic full indexing feature is enabled.
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The configuration of MaxCompute data sources.
	Config *GetIndexResponseBodyResultDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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
	// ha-cn-pl32rf0****_test_api
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
	// 2
	ProcessPartitionCount *int32 `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	// The configurations of the SARO data source.
	SaroConfig *GetIndexResponseBodyResultDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// The type of the data source. Valid values: odps, swift, saro, oss, and unKnow.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetIndexResponseBodyResultDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetIndexResponseBodyResultDataSourceInfo) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultDataSourceInfo) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *GetIndexResponseBodyResultDataSourceInfo) GetConfig() *GetIndexResponseBodyResultDataSourceInfoConfig {
	return s.Config
}

func (s *GetIndexResponseBodyResultDataSourceInfo) GetDomain() *string {
	return s.Domain
}

func (s *GetIndexResponseBodyResultDataSourceInfo) GetName() *string {
	return s.Name
}

func (s *GetIndexResponseBodyResultDataSourceInfo) GetProcessParallelNum() *int32 {
	return s.ProcessParallelNum
}

func (s *GetIndexResponseBodyResultDataSourceInfo) GetProcessPartitionCount() *int32 {
	return s.ProcessPartitionCount
}

func (s *GetIndexResponseBodyResultDataSourceInfo) GetSaroConfig() *GetIndexResponseBodyResultDataSourceInfoSaroConfig {
	return s.SaroConfig
}

func (s *GetIndexResponseBodyResultDataSourceInfo) GetType() *string {
	return s.Type
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetAutoBuildIndex(v bool) *GetIndexResponseBodyResultDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetConfig(v *GetIndexResponseBodyResultDataSourceInfoConfig) *GetIndexResponseBodyResultDataSourceInfo {
	s.Config = v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetDomain(v string) *GetIndexResponseBodyResultDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetName(v string) *GetIndexResponseBodyResultDataSourceInfo {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetProcessParallelNum(v int32) *GetIndexResponseBodyResultDataSourceInfo {
	s.ProcessParallelNum = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetProcessPartitionCount(v int32) *GetIndexResponseBodyResultDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetSaroConfig(v *GetIndexResponseBodyResultDataSourceInfoSaroConfig) *GetIndexResponseBodyResultDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) SetType(v string) *GetIndexResponseBodyResultDataSourceInfo {
	s.Type = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfo) Validate() error {
	return dara.Validate(s)
}

type GetIndexResponseBodyResultDataSourceInfoConfig struct {
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
	// ha3test-oss
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
	// TEST_dump_demo_sj_na61hunbu2_share_holo
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The Object Storage Service (OSS) path.
	//
	// example:
	//
	// /test_opensearch/sift_oss_test.data
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The partition in the MaxCompute table. Example: ds=20180102.
	//
	// example:
	//
	// ds=20220713
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The file path in the Apsara File Storage for HDFS file system.
	//
	// example:
	//
	// http://test_opensearch/sift_oss_test.data
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The name of the MaxCompute project that is used as the data source.
	//
	// example:
	//
	// tisplus_dev
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// The name of the MaxCompute table that is used as the data source.
	//
	// example:
	//
	// dump_odps_demo
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	Tag   *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s GetIndexResponseBodyResultDataSourceInfoConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIndexResponseBodyResultDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetBucket() *string {
	return s.Bucket
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetDatabase() *string {
	return s.Database
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetFormat() *string {
	return s.Format
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetOssPath() *string {
	return s.OssPath
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetPartition() *string {
	return s.Partition
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetPath() *string {
	return s.Path
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetProject() *string {
	return s.Project
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetTable() *string {
	return s.Table
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) GetTag() *string {
	return s.Tag
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetAccessKey(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetAccessSecret(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetBucket(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetCatalog(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Catalog = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetDatabase(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Database = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetEndpoint(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetFormat(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Format = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetNamespace(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetOssPath(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetPartition(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetPath(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetProject(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetTable(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Table = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) SetTag(v string) *GetIndexResponseBodyResultDataSourceInfoConfig {
	s.Tag = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoConfig) Validate() error {
	return dara.Validate(s)
}

type GetIndexResponseBodyResultDataSourceInfoSaroConfig struct {
	// The namespace of the SARO data source.
	//
	// example:
	//
	// TEST_dump_demo_sj_na61hunbu2_share_holo
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the SARO table.
	//
	// example:
	//
	// llm
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s GetIndexResponseBodyResultDataSourceInfoSaroConfig) String() string {
	return dara.Prettify(s)
}

func (s GetIndexResponseBodyResultDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultDataSourceInfoSaroConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *GetIndexResponseBodyResultDataSourceInfoSaroConfig) GetTableName() *string {
	return s.TableName
}

func (s *GetIndexResponseBodyResultDataSourceInfoSaroConfig) SetNamespace(v string) *GetIndexResponseBodyResultDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoSaroConfig) SetTableName(v string) *GetIndexResponseBodyResultDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

func (s *GetIndexResponseBodyResultDataSourceInfoSaroConfig) Validate() error {
	return dara.Validate(s)
}

type GetIndexResponseBodyResultVersions struct {
	// The description of the version.
	//
	// example:
	//
	// close alarm, by 3.9.2 hotfix workflow
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The information about the files.
	Files []*GetIndexResponseBodyResultVersionsFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The version name.
	//
	// example:
	//
	// ha-cn-7pp2ngv4s02_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the index version. Valid values:
	//
	// 	- NEW: The index version is created.
	//
	// 	- PUBLISH: The index version is normal.
	//
	// 	- IN_USE: The index version is in use.
	//
	// 	- NOT_USE: The index version is not used.
	//
	// 	- STOP_USE: The index version is being stopped.
	//
	// 	- RESTORE_USE: The index version is being restored.
	//
	// 	- FAIL: The index version failed to be created.
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the index version was updated.
	//
	// example:
	//
	// " "
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The version ID.
	//
	// example:
	//
	// 1
	VersionId *int32 `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetIndexResponseBodyResultVersions) String() string {
	return dara.Prettify(s)
}

func (s GetIndexResponseBodyResultVersions) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultVersions) GetDesc() *string {
	return s.Desc
}

func (s *GetIndexResponseBodyResultVersions) GetFiles() []*GetIndexResponseBodyResultVersionsFiles {
	return s.Files
}

func (s *GetIndexResponseBodyResultVersions) GetName() *string {
	return s.Name
}

func (s *GetIndexResponseBodyResultVersions) GetStatus() *string {
	return s.Status
}

func (s *GetIndexResponseBodyResultVersions) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetIndexResponseBodyResultVersions) GetVersionId() *int32 {
	return s.VersionId
}

func (s *GetIndexResponseBodyResultVersions) SetDesc(v string) *GetIndexResponseBodyResultVersions {
	s.Desc = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetFiles(v []*GetIndexResponseBodyResultVersionsFiles) *GetIndexResponseBodyResultVersions {
	s.Files = v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetName(v string) *GetIndexResponseBodyResultVersions {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetStatus(v string) *GetIndexResponseBodyResultVersions {
	s.Status = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetUpdateTime(v int64) *GetIndexResponseBodyResultVersions {
	s.UpdateTime = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) SetVersionId(v int32) *GetIndexResponseBodyResultVersions {
	s.VersionId = &v
	return s
}

func (s *GetIndexResponseBodyResultVersions) Validate() error {
	return dara.Validate(s)
}

type GetIndexResponseBodyResultVersionsFiles struct {
	// The full path of the file.
	//
	// example:
	//
	// " "
	FullPathName *string `json:"fullPathName,omitempty" xml:"fullPathName,omitempty"`
	// Indicates whether the file is a directory.
	//
	// example:
	//
	// True
	IsDir *bool `json:"isDir,omitempty" xml:"isDir,omitempty"`
	// Indicates whether the file is a template.
	//
	// example:
	//
	// True
	IsTemplate *bool `json:"isTemplate,omitempty" xml:"isTemplate,omitempty"`
	// The file name.
	//
	// example:
	//
	// qrs.json
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetIndexResponseBodyResultVersionsFiles) String() string {
	return dara.Prettify(s)
}

func (s GetIndexResponseBodyResultVersionsFiles) GoString() string {
	return s.String()
}

func (s *GetIndexResponseBodyResultVersionsFiles) GetFullPathName() *string {
	return s.FullPathName
}

func (s *GetIndexResponseBodyResultVersionsFiles) GetIsDir() *bool {
	return s.IsDir
}

func (s *GetIndexResponseBodyResultVersionsFiles) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *GetIndexResponseBodyResultVersionsFiles) GetName() *string {
	return s.Name
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetFullPathName(v string) *GetIndexResponseBodyResultVersionsFiles {
	s.FullPathName = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetIsDir(v bool) *GetIndexResponseBodyResultVersionsFiles {
	s.IsDir = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetIsTemplate(v bool) *GetIndexResponseBodyResultVersionsFiles {
	s.IsTemplate = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) SetName(v string) *GetIndexResponseBodyResultVersionsFiles {
	s.Name = &v
	return s
}

func (s *GetIndexResponseBodyResultVersionsFiles) Validate() error {
	return dara.Validate(s)
}
