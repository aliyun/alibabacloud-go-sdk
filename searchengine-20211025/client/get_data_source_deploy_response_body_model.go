// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceDeployResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetDataSourceDeployResponseBody
	GetRequestId() *string
	SetResult(v *GetDataSourceDeployResponseBodyResult) *GetDataSourceDeployResponseBody
	GetResult() *GetDataSourceDeployResponseBodyResult
}

type GetDataSourceDeployResponseBody struct {
	// requestId
	//
	// example:
	//
	// 2AE63638-5420-56DC-BF59-37D8174039A0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result.
	Result *GetDataSourceDeployResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetDataSourceDeployResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataSourceDeployResponseBody) GetResult() *GetDataSourceDeployResponseBodyResult {
	return s.Result
}

func (s *GetDataSourceDeployResponseBody) SetRequestId(v string) *GetDataSourceDeployResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceDeployResponseBody) SetResult(v *GetDataSourceDeployResponseBodyResult) *GetDataSourceDeployResponseBody {
	s.Result = v
	return s
}

func (s *GetDataSourceDeployResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDeployResponseBodyResult struct {
	// example:
	//
	// true
	AutoBuildIndex *bool                                        `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	Extend         *GetDataSourceDeployResponseBodyResultExtend `json:"extend,omitempty" xml:"extend,omitempty" type:"Struct"`
	// The parameters of the process.
	Processor *GetDataSourceDeployResponseBodyResultProcessor `json:"processor,omitempty" xml:"processor,omitempty" type:"Struct"`
	// The information about the data source.
	Storage *GetDataSourceDeployResponseBodyResultStorage `json:"storage,omitempty" xml:"storage,omitempty" type:"Struct"`
	// The information about the incremental data source Swift.
	Swift *GetDataSourceDeployResponseBodyResultSwift `json:"swift,omitempty" xml:"swift,omitempty" type:"Struct"`
}

func (s GetDataSourceDeployResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResult) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *GetDataSourceDeployResponseBodyResult) GetExtend() *GetDataSourceDeployResponseBodyResultExtend {
	return s.Extend
}

func (s *GetDataSourceDeployResponseBodyResult) GetProcessor() *GetDataSourceDeployResponseBodyResultProcessor {
	return s.Processor
}

func (s *GetDataSourceDeployResponseBodyResult) GetStorage() *GetDataSourceDeployResponseBodyResultStorage {
	return s.Storage
}

func (s *GetDataSourceDeployResponseBodyResult) GetSwift() *GetDataSourceDeployResponseBodyResultSwift {
	return s.Swift
}

func (s *GetDataSourceDeployResponseBodyResult) SetAutoBuildIndex(v bool) *GetDataSourceDeployResponseBodyResult {
	s.AutoBuildIndex = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetExtend(v *GetDataSourceDeployResponseBodyResultExtend) *GetDataSourceDeployResponseBodyResult {
	s.Extend = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetProcessor(v *GetDataSourceDeployResponseBodyResultProcessor) *GetDataSourceDeployResponseBodyResult {
	s.Processor = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetStorage(v *GetDataSourceDeployResponseBodyResultStorage) *GetDataSourceDeployResponseBodyResult {
	s.Storage = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) SetSwift(v *GetDataSourceDeployResponseBodyResultSwift) *GetDataSourceDeployResponseBodyResult {
	s.Swift = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDeployResponseBodyResultExtend struct {
	Hdfs *GetDataSourceDeployResponseBodyResultExtendHdfs `json:"hdfs,omitempty" xml:"hdfs,omitempty" type:"Struct"`
	Odps *GetDataSourceDeployResponseBodyResultExtendOdps `json:"odps,omitempty" xml:"odps,omitempty" type:"Struct"`
	Oss  *GetDataSourceDeployResponseBodyResultExtendOss  `json:"oss,omitempty" xml:"oss,omitempty" type:"Struct"`
	Saro *GetDataSourceDeployResponseBodyResultExtendSaro `json:"saro,omitempty" xml:"saro,omitempty" type:"Struct"`
}

func (s GetDataSourceDeployResponseBodyResultExtend) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtend) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtend) GetHdfs() *GetDataSourceDeployResponseBodyResultExtendHdfs {
	return s.Hdfs
}

func (s *GetDataSourceDeployResponseBodyResultExtend) GetOdps() *GetDataSourceDeployResponseBodyResultExtendOdps {
	return s.Odps
}

func (s *GetDataSourceDeployResponseBodyResultExtend) GetOss() *GetDataSourceDeployResponseBodyResultExtendOss {
	return s.Oss
}

func (s *GetDataSourceDeployResponseBodyResultExtend) GetSaro() *GetDataSourceDeployResponseBodyResultExtendSaro {
	return s.Saro
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetHdfs(v *GetDataSourceDeployResponseBodyResultExtendHdfs) *GetDataSourceDeployResponseBodyResultExtend {
	s.Hdfs = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetOdps(v *GetDataSourceDeployResponseBodyResultExtendOdps) *GetDataSourceDeployResponseBodyResultExtend {
	s.Odps = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetOss(v *GetDataSourceDeployResponseBodyResultExtendOss) *GetDataSourceDeployResponseBodyResultExtend {
	s.Oss = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) SetSaro(v *GetDataSourceDeployResponseBodyResultExtendSaro) *GetDataSourceDeployResponseBodyResultExtend {
	s.Saro = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtend) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDeployResponseBodyResultExtendHdfs struct {
	// example:
	//
	// dist-dmj-job/src/main/java
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendHdfs) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendHdfs) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendHdfs) GetPath() *string {
	return s.Path
}

func (s *GetDataSourceDeployResponseBodyResultExtendHdfs) SetPath(v string) *GetDataSourceDeployResponseBodyResultExtendHdfs {
	s.Path = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtendHdfs) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDeployResponseBodyResultExtendOdps struct {
	Partitions map[string]*string `json:"partitions,omitempty" xml:"partitions,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendOdps) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendOdps) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendOdps) GetPartitions() map[string]*string {
	return s.Partitions
}

func (s *GetDataSourceDeployResponseBodyResultExtendOdps) SetPartitions(v map[string]*string) *GetDataSourceDeployResponseBodyResultExtendOdps {
	s.Partitions = v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtendOdps) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDeployResponseBodyResultExtendOss struct {
	// example:
	//
	// oss://opensearch
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendOss) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendOss) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendOss) GetPath() *string {
	return s.Path
}

func (s *GetDataSourceDeployResponseBodyResultExtendOss) SetPath(v string) *GetDataSourceDeployResponseBodyResultExtendOss {
	s.Path = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtendOss) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDeployResponseBodyResultExtendSaro struct {
	// example:
	//
	// dist-dmj-job/src/main/java
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 0.6.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultExtendSaro) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultExtendSaro) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultExtendSaro) GetPath() *string {
	return s.Path
}

func (s *GetDataSourceDeployResponseBodyResultExtendSaro) GetVersion() *string {
	return s.Version
}

func (s *GetDataSourceDeployResponseBodyResultExtendSaro) SetPath(v string) *GetDataSourceDeployResponseBodyResultExtendSaro {
	s.Path = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtendSaro) SetVersion(v string) *GetDataSourceDeployResponseBodyResultExtendSaro {
	s.Version = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultExtendSaro) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDeployResponseBodyResultProcessor struct {
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

func (s GetDataSourceDeployResponseBodyResultProcessor) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultProcessor) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultProcessor) GetArgs() *string {
	return s.Args
}

func (s *GetDataSourceDeployResponseBodyResultProcessor) GetResource() *string {
	return s.Resource
}

func (s *GetDataSourceDeployResponseBodyResultProcessor) SetArgs(v string) *GetDataSourceDeployResponseBodyResultProcessor {
	s.Args = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultProcessor) SetResource(v string) *GetDataSourceDeployResponseBodyResultProcessor {
	s.Resource = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultProcessor) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDeployResponseBodyResultStorage struct {
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
	// antsys-miniapp-chongwen-static
	Bucket   *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Catalog  *string `json:"catalog,omitempty" xml:"catalog,omitempty"`
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	// The endpoint of the MaxCompute data source.
	//
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// example:
	//
	// lazada-campaign-flink
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The Object Storage Service (OSS) path.
	//
	// example:
	//
	// oss://opensearch
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The partition in the MaxCompute table. Example: ds=20180102.
	//
	// example:
	//
	// ds=20220926
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	// example:
	//
	// /beiming_xobject/dwd_xobjectsandbox__list_create_action_by_new/
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// wireless_1688_personal_rec
	Project *string `json:"project,omitempty" xml:"project,omitempty"`
	// example:
	//
	// behavior
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	Tag   *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultStorage) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultStorage) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetBucket() *string {
	return s.Bucket
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetCatalog() *string {
	return s.Catalog
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetDatabase() *string {
	return s.Database
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetNamespace() *string {
	return s.Namespace
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetOssPath() *string {
	return s.OssPath
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetPartition() *string {
	return s.Partition
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetPath() *string {
	return s.Path
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetProject() *string {
	return s.Project
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetTable() *string {
	return s.Table
}

func (s *GetDataSourceDeployResponseBodyResultStorage) GetTag() *string {
	return s.Tag
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetAccessKey(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.AccessKey = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetAccessSecret(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.AccessSecret = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetBucket(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Bucket = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetCatalog(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Catalog = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetDatabase(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Database = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetEndpoint(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Endpoint = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetNamespace(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Namespace = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetOssPath(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.OssPath = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetPartition(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Partition = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetPath(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Path = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetProject(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Project = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetTable(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Table = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) SetTag(v string) *GetDataSourceDeployResponseBodyResultStorage {
	s.Tag = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultStorage) Validate() error {
	return dara.Validate(s)
}

type GetDataSourceDeployResponseBodyResultSwift struct {
	// The topic.
	//
	// example:
	//
	// topic
	Topic *string `json:"topic,omitempty" xml:"topic,omitempty"`
	// zk
	//
	// example:
	//
	// zk
	Zk *string `json:"zk,omitempty" xml:"zk,omitempty"`
}

func (s GetDataSourceDeployResponseBodyResultSwift) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponseBodyResultSwift) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponseBodyResultSwift) GetTopic() *string {
	return s.Topic
}

func (s *GetDataSourceDeployResponseBodyResultSwift) GetZk() *string {
	return s.Zk
}

func (s *GetDataSourceDeployResponseBodyResultSwift) SetTopic(v string) *GetDataSourceDeployResponseBodyResultSwift {
	s.Topic = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultSwift) SetZk(v string) *GetDataSourceDeployResponseBodyResultSwift {
	s.Zk = &v
	return s
}

func (s *GetDataSourceDeployResponseBodyResultSwift) Validate() error {
	return dara.Validate(s)
}
