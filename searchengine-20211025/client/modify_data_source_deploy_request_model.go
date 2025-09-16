// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataSourceDeployRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBuildIndex(v bool) *ModifyDataSourceDeployRequest
	GetAutoBuildIndex() *bool
	SetExtend(v *ModifyDataSourceDeployRequestExtend) *ModifyDataSourceDeployRequest
	GetExtend() *ModifyDataSourceDeployRequestExtend
	SetProcessor(v *ModifyDataSourceDeployRequestProcessor) *ModifyDataSourceDeployRequest
	GetProcessor() *ModifyDataSourceDeployRequestProcessor
	SetStorage(v *ModifyDataSourceDeployRequestStorage) *ModifyDataSourceDeployRequest
	GetStorage() *ModifyDataSourceDeployRequestStorage
	SetSwift(v *ModifyDataSourceDeployRequestSwift) *ModifyDataSourceDeployRequest
	GetSwift() *ModifyDataSourceDeployRequestSwift
	SetDryRun(v bool) *ModifyDataSourceDeployRequest
	GetDryRun() *bool
	SetGenerationId(v int64) *ModifyDataSourceDeployRequest
	GetGenerationId() *int64
}

type ModifyDataSourceDeployRequest struct {
	// Specifies whether to enable the automatic full indexing feature.
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The extended information.
	Extend *ModifyDataSourceDeployRequestExtend `json:"extend,omitempty" xml:"extend,omitempty" type:"Struct"`
	// The parameters of the process.
	Processor *ModifyDataSourceDeployRequestProcessor `json:"processor,omitempty" xml:"processor,omitempty" type:"Struct"`
	// The information about the data source.
	Storage *ModifyDataSourceDeployRequestStorage `json:"storage,omitempty" xml:"storage,omitempty" type:"Struct"`
	// The information about the incremental data source Swift.
	Swift *ModifyDataSourceDeployRequestSwift `json:"swift,omitempty" xml:"swift,omitempty" type:"Struct"`
	// Specifies whether to perform only a dry run, without performing the actual request. The system only checks the validity of the data source. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
	// The ID of the full index version.
	//
	// example:
	//
	// 1708674867
	GenerationId *int64 `json:"generationId,omitempty" xml:"generationId,omitempty"`
}

func (s ModifyDataSourceDeployRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequest) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *ModifyDataSourceDeployRequest) GetExtend() *ModifyDataSourceDeployRequestExtend {
	return s.Extend
}

func (s *ModifyDataSourceDeployRequest) GetProcessor() *ModifyDataSourceDeployRequestProcessor {
	return s.Processor
}

func (s *ModifyDataSourceDeployRequest) GetStorage() *ModifyDataSourceDeployRequestStorage {
	return s.Storage
}

func (s *ModifyDataSourceDeployRequest) GetSwift() *ModifyDataSourceDeployRequestSwift {
	return s.Swift
}

func (s *ModifyDataSourceDeployRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDataSourceDeployRequest) GetGenerationId() *int64 {
	return s.GenerationId
}

func (s *ModifyDataSourceDeployRequest) SetAutoBuildIndex(v bool) *ModifyDataSourceDeployRequest {
	s.AutoBuildIndex = &v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetExtend(v *ModifyDataSourceDeployRequestExtend) *ModifyDataSourceDeployRequest {
	s.Extend = v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetProcessor(v *ModifyDataSourceDeployRequestProcessor) *ModifyDataSourceDeployRequest {
	s.Processor = v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetStorage(v *ModifyDataSourceDeployRequestStorage) *ModifyDataSourceDeployRequest {
	s.Storage = v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetSwift(v *ModifyDataSourceDeployRequestSwift) *ModifyDataSourceDeployRequest {
	s.Swift = v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetDryRun(v bool) *ModifyDataSourceDeployRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDataSourceDeployRequest) SetGenerationId(v int64) *ModifyDataSourceDeployRequest {
	s.GenerationId = &v
	return s
}

func (s *ModifyDataSourceDeployRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyDataSourceDeployRequestExtend struct {
	// The information about the Apsara File Storage for HDFS data source.
	Hdfs *ModifyDataSourceDeployRequestExtendHdfs `json:"hdfs,omitempty" xml:"hdfs,omitempty" type:"Struct"`
	// The information about the MaxCompute data source.
	Odps *ModifyDataSourceDeployRequestExtendOdps `json:"odps,omitempty" xml:"odps,omitempty" type:"Struct"`
	// The information about the OSS data source.
	Oss *ModifyDataSourceDeployRequestExtendOss `json:"oss,omitempty" xml:"oss,omitempty" type:"Struct"`
	// The information about the SARO data source. This parameter is applicable to the SARO data source used in the intranet of Alibaba Group.
	Saro *ModifyDataSourceDeployRequestExtendSaro `json:"saro,omitempty" xml:"saro,omitempty" type:"Struct"`
}

func (s ModifyDataSourceDeployRequestExtend) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtend) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtend) GetHdfs() *ModifyDataSourceDeployRequestExtendHdfs {
	return s.Hdfs
}

func (s *ModifyDataSourceDeployRequestExtend) GetOdps() *ModifyDataSourceDeployRequestExtendOdps {
	return s.Odps
}

func (s *ModifyDataSourceDeployRequestExtend) GetOss() *ModifyDataSourceDeployRequestExtendOss {
	return s.Oss
}

func (s *ModifyDataSourceDeployRequestExtend) GetSaro() *ModifyDataSourceDeployRequestExtendSaro {
	return s.Saro
}

func (s *ModifyDataSourceDeployRequestExtend) SetHdfs(v *ModifyDataSourceDeployRequestExtendHdfs) *ModifyDataSourceDeployRequestExtend {
	s.Hdfs = v
	return s
}

func (s *ModifyDataSourceDeployRequestExtend) SetOdps(v *ModifyDataSourceDeployRequestExtendOdps) *ModifyDataSourceDeployRequestExtend {
	s.Odps = v
	return s
}

func (s *ModifyDataSourceDeployRequestExtend) SetOss(v *ModifyDataSourceDeployRequestExtendOss) *ModifyDataSourceDeployRequestExtend {
	s.Oss = v
	return s
}

func (s *ModifyDataSourceDeployRequestExtend) SetSaro(v *ModifyDataSourceDeployRequestExtendSaro) *ModifyDataSourceDeployRequestExtend {
	s.Saro = v
	return s
}

func (s *ModifyDataSourceDeployRequestExtend) Validate() error {
	return dara.Validate(s)
}

type ModifyDataSourceDeployRequestExtendHdfs struct {
	// The path of the Apsara File Storage for HDFS data source.
	//
	// example:
	//
	// ymsh-service/src/main/java/cn/ymsh/util/jd
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ModifyDataSourceDeployRequestExtendHdfs) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtendHdfs) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtendHdfs) GetPath() *string {
	return s.Path
}

func (s *ModifyDataSourceDeployRequestExtendHdfs) SetPath(v string) *ModifyDataSourceDeployRequestExtendHdfs {
	s.Path = &v
	return s
}

func (s *ModifyDataSourceDeployRequestExtendHdfs) Validate() error {
	return dara.Validate(s)
}

type ModifyDataSourceDeployRequestExtendOdps struct {
	// The partitions in the MaxCompute table.
	Partitions map[string]*string `json:"partitions,omitempty" xml:"partitions,omitempty"`
}

func (s ModifyDataSourceDeployRequestExtendOdps) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtendOdps) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtendOdps) GetPartitions() map[string]*string {
	return s.Partitions
}

func (s *ModifyDataSourceDeployRequestExtendOdps) SetPartitions(v map[string]*string) *ModifyDataSourceDeployRequestExtendOdps {
	s.Partitions = v
	return s
}

func (s *ModifyDataSourceDeployRequestExtendOdps) Validate() error {
	return dara.Validate(s)
}

type ModifyDataSourceDeployRequestExtendOss struct {
	// The path of the OSS data source.
	//
	// example:
	//
	// oss://test
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ModifyDataSourceDeployRequestExtendOss) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtendOss) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtendOss) GetPath() *string {
	return s.Path
}

func (s *ModifyDataSourceDeployRequestExtendOss) SetPath(v string) *ModifyDataSourceDeployRequestExtendOss {
	s.Path = &v
	return s
}

func (s *ModifyDataSourceDeployRequestExtendOss) Validate() error {
	return dara.Validate(s)
}

type ModifyDataSourceDeployRequestExtendSaro struct {
	// The path of the SARO data source.
	//
	// example:
	//
	// /
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The version number of the SARO data source.
	//
	// example:
	//
	// 1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ModifyDataSourceDeployRequestExtendSaro) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployRequestExtendSaro) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestExtendSaro) GetPath() *string {
	return s.Path
}

func (s *ModifyDataSourceDeployRequestExtendSaro) GetVersion() *string {
	return s.Version
}

func (s *ModifyDataSourceDeployRequestExtendSaro) SetPath(v string) *ModifyDataSourceDeployRequestExtendSaro {
	s.Path = &v
	return s
}

func (s *ModifyDataSourceDeployRequestExtendSaro) SetVersion(v string) *ModifyDataSourceDeployRequestExtendSaro {
	s.Version = &v
	return s
}

func (s *ModifyDataSourceDeployRequestExtendSaro) Validate() error {
	return dara.Validate(s)
}

type ModifyDataSourceDeployRequestProcessor struct {
	// The startup parameters of the process.
	//
	// example:
	//
	// {}
	Args *string `json:"args,omitempty" xml:"args,omitempty"`
	// The resource information.
	//
	// example:
	//
	// {}
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s ModifyDataSourceDeployRequestProcessor) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployRequestProcessor) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestProcessor) GetArgs() *string {
	return s.Args
}

func (s *ModifyDataSourceDeployRequestProcessor) GetResource() *string {
	return s.Resource
}

func (s *ModifyDataSourceDeployRequestProcessor) SetArgs(v string) *ModifyDataSourceDeployRequestProcessor {
	s.Args = &v
	return s
}

func (s *ModifyDataSourceDeployRequestProcessor) SetResource(v string) *ModifyDataSourceDeployRequestProcessor {
	s.Resource = &v
	return s
}

func (s *ModifyDataSourceDeployRequestProcessor) Validate() error {
	return dara.Validate(s)
}

type ModifyDataSourceDeployRequestStorage struct {
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
	// The namespace. This parameter is applicable to the SARO data source used in the intranet of Alibaba Group.
	//
	// example:
	//
	// dp-dev
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The Object Storage Service (OSS) path.
	//
	// example:
	//
	// /opensearch
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The partition in the MaxCompute table.
	//
	// example:
	//
	// ds=20220713
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// The file path in the Apsara File Storage for HDFS file system.
	//
	// example:
	//
	// /ude_jobs/iflow_offline_data_access
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The name of the MaxCompute project that is used as the data source.
	//
	// example:
	//
	// kubenest
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

func (s ModifyDataSourceDeployRequestStorage) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployRequestStorage) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestStorage) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ModifyDataSourceDeployRequestStorage) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *ModifyDataSourceDeployRequestStorage) GetBucket() *string {
	return s.Bucket
}

func (s *ModifyDataSourceDeployRequestStorage) GetCatalog() *string {
	return s.Catalog
}

func (s *ModifyDataSourceDeployRequestStorage) GetDatabase() *string {
	return s.Database
}

func (s *ModifyDataSourceDeployRequestStorage) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ModifyDataSourceDeployRequestStorage) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyDataSourceDeployRequestStorage) GetOssPath() *string {
	return s.OssPath
}

func (s *ModifyDataSourceDeployRequestStorage) GetPartition() *string {
	return s.Partition
}

func (s *ModifyDataSourceDeployRequestStorage) GetPath() *string {
	return s.Path
}

func (s *ModifyDataSourceDeployRequestStorage) GetProject() *string {
	return s.Project
}

func (s *ModifyDataSourceDeployRequestStorage) GetTable() *string {
	return s.Table
}

func (s *ModifyDataSourceDeployRequestStorage) GetTableFormat() *string {
	return s.TableFormat
}

func (s *ModifyDataSourceDeployRequestStorage) GetTag() *string {
	return s.Tag
}

func (s *ModifyDataSourceDeployRequestStorage) SetAccessKey(v string) *ModifyDataSourceDeployRequestStorage {
	s.AccessKey = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetAccessSecret(v string) *ModifyDataSourceDeployRequestStorage {
	s.AccessSecret = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetBucket(v string) *ModifyDataSourceDeployRequestStorage {
	s.Bucket = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetCatalog(v string) *ModifyDataSourceDeployRequestStorage {
	s.Catalog = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetDatabase(v string) *ModifyDataSourceDeployRequestStorage {
	s.Database = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetEndpoint(v string) *ModifyDataSourceDeployRequestStorage {
	s.Endpoint = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetNamespace(v string) *ModifyDataSourceDeployRequestStorage {
	s.Namespace = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetOssPath(v string) *ModifyDataSourceDeployRequestStorage {
	s.OssPath = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetPartition(v string) *ModifyDataSourceDeployRequestStorage {
	s.Partition = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetPath(v string) *ModifyDataSourceDeployRequestStorage {
	s.Path = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetProject(v string) *ModifyDataSourceDeployRequestStorage {
	s.Project = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetTable(v string) *ModifyDataSourceDeployRequestStorage {
	s.Table = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetTableFormat(v string) *ModifyDataSourceDeployRequestStorage {
	s.TableFormat = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) SetTag(v string) *ModifyDataSourceDeployRequestStorage {
	s.Tag = &v
	return s
}

func (s *ModifyDataSourceDeployRequestStorage) Validate() error {
	return dara.Validate(s)
}

type ModifyDataSourceDeployRequestSwift struct {
	// The topic.
	//
	// example:
	//
	// ha-cn-0ju2rps6c08_api
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// zk
	//
	// example:
	//
	// zk
	Zk *string `json:"zk,omitempty" xml:"zk,omitempty"`
}

func (s ModifyDataSourceDeployRequestSwift) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployRequestSwift) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployRequestSwift) GetTopic() *string {
	return s.Topic
}

func (s *ModifyDataSourceDeployRequestSwift) GetZk() *string {
	return s.Zk
}

func (s *ModifyDataSourceDeployRequestSwift) SetTopic(v string) *ModifyDataSourceDeployRequestSwift {
	s.Topic = &v
	return s
}

func (s *ModifyDataSourceDeployRequestSwift) SetZk(v string) *ModifyDataSourceDeployRequestSwift {
	s.Zk = &v
	return s
}

func (s *ModifyDataSourceDeployRequestSwift) Validate() error {
	return dara.Validate(s)
}
