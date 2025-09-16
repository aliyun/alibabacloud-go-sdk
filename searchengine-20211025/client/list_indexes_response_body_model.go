// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListIndexesResponseBody
	GetRequestId() *string
	SetResult(v []*ListIndexesResponseBodyResult) *ListIndexesResponseBody
	GetResult() []*ListIndexesResponseBodyResult
}

type ListIndexesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4FB0325E-8C37-5525-96AC-0333523170A3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of indexes.
	Result []*ListIndexesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListIndexesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIndexesResponseBody) GetResult() []*ListIndexesResponseBodyResult {
	return s.Result
}

func (s *ListIndexesResponseBody) SetRequestId(v string) *ListIndexesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexesResponseBody) SetResult(v []*ListIndexesResponseBodyResult) *ListIndexesResponseBody {
	s.Result = v
	return s
}

func (s *ListIndexesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIndexesResponseBodyResult struct {
	// The index schema, which is a JSON string.
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
	// ha-cn-7mz2kvu2c01_table4
	DataSource *string `json:"dataSource,omitempty" xml:"dataSource,omitempty"`
	// The information about the data source.
	DataSourceInfo *ListIndexesResponseBodyResultDataSourceInfo `json:"dataSourceInfo,omitempty" xml:"dataSourceInfo,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// Description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The deployment name of the index.
	//
	// example:
	//
	// test
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The time when full data in the index was last updated.
	//
	// example:
	//
	// 2023-07-05 10:40:38
	FullUpdateTime *string `json:"fullUpdateTime,omitempty" xml:"fullUpdateTime,omitempty"`
	// The full version of the index.
	//
	// example:
	//
	// 1688523414
	FullVersion *int64 `json:"fullVersion,omitempty" xml:"fullVersion,omitempty"`
	// The time when incremental data in the index was last updated.
	//
	// example:
	//
	// 2023-07-05 10:58:33
	IncUpdateTime *string `json:"incUpdateTime,omitempty" xml:"incUpdateTime,omitempty"`
	// The index size.
	//
	// example:
	//
	// 4689
	IndexSize *int64 `json:"indexSize,omitempty" xml:"indexSize,omitempty"`
	// The index ststus. Valid values: NEW and PUBLISH.
	//
	// example:
	//
	// " "
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
	// The index versions.
	Versions []*ListIndexesResponseBodyResultVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListIndexesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResult) GetContent() *string {
	return s.Content
}

func (s *ListIndexesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIndexesResponseBodyResult) GetDataSource() *string {
	return s.DataSource
}

func (s *ListIndexesResponseBodyResult) GetDataSourceInfo() *ListIndexesResponseBodyResultDataSourceInfo {
	return s.DataSourceInfo
}

func (s *ListIndexesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListIndexesResponseBodyResult) GetDomain() *string {
	return s.Domain
}

func (s *ListIndexesResponseBodyResult) GetFullUpdateTime() *string {
	return s.FullUpdateTime
}

func (s *ListIndexesResponseBodyResult) GetFullVersion() *int64 {
	return s.FullVersion
}

func (s *ListIndexesResponseBodyResult) GetIncUpdateTime() *string {
	return s.IncUpdateTime
}

func (s *ListIndexesResponseBodyResult) GetIndexSize() *int64 {
	return s.IndexSize
}

func (s *ListIndexesResponseBodyResult) GetIndexStatus() *string {
	return s.IndexStatus
}

func (s *ListIndexesResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListIndexesResponseBodyResult) GetPartition() *int32 {
	return s.Partition
}

func (s *ListIndexesResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListIndexesResponseBodyResult) GetVersions() []*ListIndexesResponseBodyResultVersions {
	return s.Versions
}

func (s *ListIndexesResponseBodyResult) SetContent(v string) *ListIndexesResponseBodyResult {
	s.Content = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetCreateTime(v string) *ListIndexesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDataSource(v string) *ListIndexesResponseBodyResult {
	s.DataSource = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDataSourceInfo(v *ListIndexesResponseBodyResultDataSourceInfo) *ListIndexesResponseBodyResult {
	s.DataSourceInfo = v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDescription(v string) *ListIndexesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetDomain(v string) *ListIndexesResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetFullUpdateTime(v string) *ListIndexesResponseBodyResult {
	s.FullUpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetFullVersion(v int64) *ListIndexesResponseBodyResult {
	s.FullVersion = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetIncUpdateTime(v string) *ListIndexesResponseBodyResult {
	s.IncUpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetIndexSize(v int64) *ListIndexesResponseBodyResult {
	s.IndexSize = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetIndexStatus(v string) *ListIndexesResponseBodyResult {
	s.IndexStatus = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetName(v string) *ListIndexesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetPartition(v int32) *ListIndexesResponseBodyResult {
	s.Partition = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetUpdateTime(v string) *ListIndexesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResult) SetVersions(v []*ListIndexesResponseBodyResultVersions) *ListIndexesResponseBodyResult {
	s.Versions = v
	return s
}

func (s *ListIndexesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListIndexesResponseBodyResultDataSourceInfo struct {
	// Indicates whether the automatic full indexing feature is enabled.
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The configuration of MaxCompute data sources.
	Config *ListIndexesResponseBodyResultDataSourceInfoConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// The data center in which the data source is deployed.
	//
	// example:
	//
	// test
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// index1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of resources used for data update.
	//
	// example:
	//
	// 2
	ProcessPartitionCount *int32 `json:"processPartitionCount,omitempty" xml:"processPartitionCount,omitempty"`
	// The configurations of the SARO data source.
	SaroConfig *ListIndexesResponseBodyResultDataSourceInfoSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// The type of the data source. Valid values: odps, swift, saro, oss, and unKnow.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListIndexesResponseBodyResultDataSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBodyResultDataSourceInfo) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) GetConfig() *ListIndexesResponseBodyResultDataSourceInfoConfig {
	return s.Config
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) GetDomain() *string {
	return s.Domain
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) GetName() *string {
	return s.Name
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) GetProcessPartitionCount() *int32 {
	return s.ProcessPartitionCount
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) GetSaroConfig() *ListIndexesResponseBodyResultDataSourceInfoSaroConfig {
	return s.SaroConfig
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) GetType() *string {
	return s.Type
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetAutoBuildIndex(v bool) *ListIndexesResponseBodyResultDataSourceInfo {
	s.AutoBuildIndex = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetConfig(v *ListIndexesResponseBodyResultDataSourceInfoConfig) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Config = v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetDomain(v string) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Domain = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetName(v string) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetProcessPartitionCount(v int32) *ListIndexesResponseBodyResultDataSourceInfo {
	s.ProcessPartitionCount = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetSaroConfig(v *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) *ListIndexesResponseBodyResultDataSourceInfo {
	s.SaroConfig = v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) SetType(v string) *ListIndexesResponseBodyResultDataSourceInfo {
	s.Type = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfo) Validate() error {
	return dara.Validate(s)
}

type ListIndexesResponseBodyResultDataSourceInfoConfig struct {
	// The AccessKey ID of the MaxCompute data source.
	//
	// example:
	//
	// root
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	// The AccessKey secret of the MaxCompute data source.
	//
	// example:
	//
	// root123
	AccessSecret *string `json:"accessSecret,omitempty" xml:"accessSecret,omitempty"`
	// The OSS bucket.
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
	// The shard name.
	//
	// example:
	//
	// ds=12345
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

func (s ListIndexesResponseBodyResultDataSourceInfoConfig) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBodyResultDataSourceInfoConfig) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetBucket() *string {
	return s.Bucket
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetDatabase() *string {
	return s.Database
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetFormat() *string {
	return s.Format
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetOssPath() *string {
	return s.OssPath
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetPartition() *string {
	return s.Partition
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetPath() *string {
	return s.Path
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetProject() *string {
	return s.Project
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetTable() *string {
	return s.Table
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) GetTag() *string {
	return s.Tag
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetAccessKey(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.AccessKey = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetAccessSecret(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.AccessSecret = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetBucket(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Bucket = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetCatalog(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Catalog = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetDatabase(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Database = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetEndpoint(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Endpoint = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetFormat(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Format = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetNamespace(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Namespace = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetOssPath(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.OssPath = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetPartition(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Partition = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetPath(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Path = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetProject(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Project = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetTable(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Table = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) SetTag(v string) *ListIndexesResponseBodyResultDataSourceInfoConfig {
	s.Tag = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoConfig) Validate() error {
	return dara.Validate(s)
}

type ListIndexesResponseBodyResultDataSourceInfoSaroConfig struct {
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
	// dump_odps_demo
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s ListIndexesResponseBodyResultDataSourceInfoSaroConfig) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBodyResultDataSourceInfoSaroConfig) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) GetTableName() *string {
	return s.TableName
}

func (s *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) SetNamespace(v string) *ListIndexesResponseBodyResultDataSourceInfoSaroConfig {
	s.Namespace = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) SetTableName(v string) *ListIndexesResponseBodyResultDataSourceInfoSaroConfig {
	s.TableName = &v
	return s
}

func (s *ListIndexesResponseBodyResultDataSourceInfoSaroConfig) Validate() error {
	return dara.Validate(s)
}

type ListIndexesResponseBodyResultVersions struct {
	// The description of the index version.
	//
	// example:
	//
	// close alarm, by 3.9.2 hotfix workflow
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// The files.
	Files []*ListIndexesResponseBodyResultVersionsFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// The name of the index version.
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
	// The ID of the index version. If the index version is modified, the returned value is null.
	//
	// example:
	//
	// 1
	VersionId *int32 `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s ListIndexesResponseBodyResultVersions) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBodyResultVersions) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultVersions) GetDesc() *string {
	return s.Desc
}

func (s *ListIndexesResponseBodyResultVersions) GetFiles() []*ListIndexesResponseBodyResultVersionsFiles {
	return s.Files
}

func (s *ListIndexesResponseBodyResultVersions) GetName() *string {
	return s.Name
}

func (s *ListIndexesResponseBodyResultVersions) GetStatus() *string {
	return s.Status
}

func (s *ListIndexesResponseBodyResultVersions) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListIndexesResponseBodyResultVersions) GetVersionId() *int32 {
	return s.VersionId
}

func (s *ListIndexesResponseBodyResultVersions) SetDesc(v string) *ListIndexesResponseBodyResultVersions {
	s.Desc = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetFiles(v []*ListIndexesResponseBodyResultVersionsFiles) *ListIndexesResponseBodyResultVersions {
	s.Files = v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetName(v string) *ListIndexesResponseBodyResultVersions {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetStatus(v string) *ListIndexesResponseBodyResultVersions {
	s.Status = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetUpdateTime(v int64) *ListIndexesResponseBodyResultVersions {
	s.UpdateTime = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) SetVersionId(v int32) *ListIndexesResponseBodyResultVersions {
	s.VersionId = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersions) Validate() error {
	return dara.Validate(s)
}

type ListIndexesResponseBodyResultVersionsFiles struct {
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
	// ha-cn-7mz2iv7sq01_qrs
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListIndexesResponseBodyResultVersionsFiles) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBodyResultVersionsFiles) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyResultVersionsFiles) GetFullPathName() *string {
	return s.FullPathName
}

func (s *ListIndexesResponseBodyResultVersionsFiles) GetIsDir() *bool {
	return s.IsDir
}

func (s *ListIndexesResponseBodyResultVersionsFiles) GetIsTemplate() *bool {
	return s.IsTemplate
}

func (s *ListIndexesResponseBodyResultVersionsFiles) GetName() *string {
	return s.Name
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetFullPathName(v string) *ListIndexesResponseBodyResultVersionsFiles {
	s.FullPathName = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetIsDir(v bool) *ListIndexesResponseBodyResultVersionsFiles {
	s.IsDir = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetIsTemplate(v bool) *ListIndexesResponseBodyResultVersionsFiles {
	s.IsTemplate = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) SetName(v string) *ListIndexesResponseBodyResultVersionsFiles {
	s.Name = &v
	return s
}

func (s *ListIndexesResponseBodyResultVersionsFiles) Validate() error {
	return dara.Validate(s)
}
