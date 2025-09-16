// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoBuildIndex(v bool) *CreateDataSourceRequest
	GetAutoBuildIndex() *bool
	SetConfig(v *CreateDataSourceRequestConfig) *CreateDataSourceRequest
	GetConfig() *CreateDataSourceRequestConfig
	SetDomain(v string) *CreateDataSourceRequest
	GetDomain() *string
	SetName(v string) *CreateDataSourceRequest
	GetName() *string
	SetSaroConfig(v *CreateDataSourceRequestSaroConfig) *CreateDataSourceRequest
	GetSaroConfig() *CreateDataSourceRequestSaroConfig
	SetType(v string) *CreateDataSourceRequest
	GetType() *string
	SetDryRun(v bool) *CreateDataSourceRequest
	GetDryRun() *bool
}

type CreateDataSourceRequest struct {
	// Specifies whether to automatically rebuild the index.
	//
	// example:
	//
	// true
	AutoBuildIndex *bool `json:"autoBuildIndex,omitempty" xml:"autoBuildIndex,omitempty"`
	// The configuration information.
	Config *CreateDataSourceRequestConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
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
	// The configurations of the SARO data source.
	SaroConfig *CreateDataSourceRequestSaroConfig `json:"saroConfig,omitempty" xml:"saroConfig,omitempty" type:"Struct"`
	// The type of the data source. Valid values: odps, oss, and swift.
	//
	// example:
	//
	// odps
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) GetAutoBuildIndex() *bool {
	return s.AutoBuildIndex
}

func (s *CreateDataSourceRequest) GetConfig() *CreateDataSourceRequestConfig {
	return s.Config
}

func (s *CreateDataSourceRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateDataSourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataSourceRequest) GetSaroConfig() *CreateDataSourceRequestSaroConfig {
	return s.SaroConfig
}

func (s *CreateDataSourceRequest) GetType() *string {
	return s.Type
}

func (s *CreateDataSourceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDataSourceRequest) SetAutoBuildIndex(v bool) *CreateDataSourceRequest {
	s.AutoBuildIndex = &v
	return s
}

func (s *CreateDataSourceRequest) SetConfig(v *CreateDataSourceRequestConfig) *CreateDataSourceRequest {
	s.Config = v
	return s
}

func (s *CreateDataSourceRequest) SetDomain(v string) *CreateDataSourceRequest {
	s.Domain = &v
	return s
}

func (s *CreateDataSourceRequest) SetName(v string) *CreateDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDataSourceRequest) SetSaroConfig(v *CreateDataSourceRequestSaroConfig) *CreateDataSourceRequest {
	s.SaroConfig = v
	return s
}

func (s *CreateDataSourceRequest) SetType(v string) *CreateDataSourceRequest {
	s.Type = &v
	return s
}

func (s *CreateDataSourceRequest) SetDryRun(v bool) *CreateDataSourceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDataSourceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDataSourceRequestConfig struct {
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
	// opensearch
	Bucket   *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Catalog  *string `json:"catalog,omitempty" xml:"catalog,omitempty"`
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	// The endpoint of the MaxCompute or Object Storage Service (OSS) data source.
	//
	// example:
	//
	// http://service.cn-hangzhou.maxcompute.aliyun-inc.com/api
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The namespace.
	//
	// example:
	//
	// aegis-ops
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The path of the OSS object.
	//
	// example:
	//
	// /opensearch/search
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
	// test-hdfs-path
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
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	Tag   *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s CreateDataSourceRequestConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestConfig) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateDataSourceRequestConfig) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *CreateDataSourceRequestConfig) GetBucket() *string {
	return s.Bucket
}

func (s *CreateDataSourceRequestConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *CreateDataSourceRequestConfig) GetDatabase() *string {
	return s.Database
}

func (s *CreateDataSourceRequestConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *CreateDataSourceRequestConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateDataSourceRequestConfig) GetOssPath() *string {
	return s.OssPath
}

func (s *CreateDataSourceRequestConfig) GetPartition() *string {
	return s.Partition
}

func (s *CreateDataSourceRequestConfig) GetPath() *string {
	return s.Path
}

func (s *CreateDataSourceRequestConfig) GetProject() *string {
	return s.Project
}

func (s *CreateDataSourceRequestConfig) GetTable() *string {
	return s.Table
}

func (s *CreateDataSourceRequestConfig) GetTag() *string {
	return s.Tag
}

func (s *CreateDataSourceRequestConfig) SetAccessKey(v string) *CreateDataSourceRequestConfig {
	s.AccessKey = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetAccessSecret(v string) *CreateDataSourceRequestConfig {
	s.AccessSecret = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetBucket(v string) *CreateDataSourceRequestConfig {
	s.Bucket = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetCatalog(v string) *CreateDataSourceRequestConfig {
	s.Catalog = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetDatabase(v string) *CreateDataSourceRequestConfig {
	s.Database = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetEndpoint(v string) *CreateDataSourceRequestConfig {
	s.Endpoint = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetNamespace(v string) *CreateDataSourceRequestConfig {
	s.Namespace = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetOssPath(v string) *CreateDataSourceRequestConfig {
	s.OssPath = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetPartition(v string) *CreateDataSourceRequestConfig {
	s.Partition = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetPath(v string) *CreateDataSourceRequestConfig {
	s.Path = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetProject(v string) *CreateDataSourceRequestConfig {
	s.Project = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetTable(v string) *CreateDataSourceRequestConfig {
	s.Table = &v
	return s
}

func (s *CreateDataSourceRequestConfig) SetTag(v string) *CreateDataSourceRequestConfig {
	s.Tag = &v
	return s
}

func (s *CreateDataSourceRequestConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDataSourceRequestSaroConfig struct {
	// The namespace of the SARO data source.
	//
	// example:
	//
	// igraph-cn-x0r3e3abe02
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the SARO table.
	//
	// example:
	//
	// index_hdfs
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
}

func (s CreateDataSourceRequestSaroConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceRequestSaroConfig) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequestSaroConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateDataSourceRequestSaroConfig) GetTableName() *string {
	return s.TableName
}

func (s *CreateDataSourceRequestSaroConfig) SetNamespace(v string) *CreateDataSourceRequestSaroConfig {
	s.Namespace = &v
	return s
}

func (s *CreateDataSourceRequestSaroConfig) SetTableName(v string) *CreateDataSourceRequestSaroConfig {
	s.TableName = &v
	return s
}

func (s *CreateDataSourceRequestSaroConfig) Validate() error {
	return dara.Validate(s)
}
