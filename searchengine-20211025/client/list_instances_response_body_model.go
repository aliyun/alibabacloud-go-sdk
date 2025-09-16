// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetResult(v []*ListInstancesResponseBodyResult) *ListInstancesResponseBody
	GetResult() []*ListInstancesResponseBodyResult
	SetTotalCount(v int32) *ListInstancesResponseBody
	GetTotalCount() *int32
}

type ListInstancesResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// 89B968E6-1E41-58DF-BB25-5F98ECC759CE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
	Result []*ListInstancesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned
	//
	// example:
	//
	// 5
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetResult() []*ListInstancesResponseBodyResult {
	return s.Result
}

func (s *ListInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetResult(v []*ListInstancesResponseBodyResult) *ListInstancesResponseBody {
	s.Result = v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyResult struct {
	// The billing method.
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The commodity code of the instance.
	//
	// example:
	//
	// ""
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2022-06-04T02:03:21Z
	CreateTime        *string                                             `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DataSourceDetails []*ListInstancesResponseBodyResultDataSourceDetails `json:"dataSourceDetails,omitempty" xml:"dataSourceDetails,omitempty" type:"Repeated"`
	// The description of the instance.
	//
	// example:
	//
	// Emergency test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Edition     *string `json:"edition,omitempty" xml:"edition,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 1634885083
	ExpiredTime *string `json:"expiredTime,omitempty" xml:"expiredTime,omitempty"`
	// Indicates whether an overdue payment is involved.
	//
	// example:
	//
	// false
	InDebt *bool `json:"inDebt,omitempty" xml:"inDebt,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// ha-cn-2r42n8oh001
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock state of the instance.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The network information of the instance.
	Network *ListInstancesResponseBodyResultNetwork `json:"network,omitempty" xml:"network,omitempty" type:"Struct"`
	NoQrs   *bool                                   `json:"noQrs,omitempty" xml:"noQrs,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzgpiswzbksdi
	ResourceGroupId *string                              `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Spec            *ListInstancesResponseBodyResultSpec `json:"spec,omitempty" xml:"spec,omitempty" type:"Struct"`
	// The instance status.
	//
	// example:
	//
	// 2
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The tags of the instance.
	Tags []*ListInstancesResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The time when the instance was updated.
	//
	// example:
	//
	// 2018-12-06T11:17:49.0
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	UserName   *string `json:"userName,omitempty" xml:"userName,omitempty"`
	Version    *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListInstancesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListInstancesResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListInstancesResponseBodyResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInstancesResponseBodyResult) GetDataSourceDetails() []*ListInstancesResponseBodyResultDataSourceDetails {
	return s.DataSourceDetails
}

func (s *ListInstancesResponseBodyResult) GetDescription() *string {
	return s.Description
}

func (s *ListInstancesResponseBodyResult) GetEdition() *string {
	return s.Edition
}

func (s *ListInstancesResponseBodyResult) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListInstancesResponseBodyResult) GetInDebt() *bool {
	return s.InDebt
}

func (s *ListInstancesResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyResult) GetLockMode() *string {
	return s.LockMode
}

func (s *ListInstancesResponseBodyResult) GetNetwork() *ListInstancesResponseBodyResultNetwork {
	return s.Network
}

func (s *ListInstancesResponseBodyResult) GetNoQrs() *bool {
	return s.NoQrs
}

func (s *ListInstancesResponseBodyResult) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyResult) GetSpec() *ListInstancesResponseBodyResultSpec {
	return s.Spec
}

func (s *ListInstancesResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyResult) GetTags() []*ListInstancesResponseBodyResultTags {
	return s.Tags
}

func (s *ListInstancesResponseBodyResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListInstancesResponseBodyResult) GetUserName() *string {
	return s.UserName
}

func (s *ListInstancesResponseBodyResult) GetVersion() *string {
	return s.Version
}

func (s *ListInstancesResponseBodyResult) SetChargeType(v string) *ListInstancesResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetCommodityCode(v string) *ListInstancesResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetCreateTime(v string) *ListInstancesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetDataSourceDetails(v []*ListInstancesResponseBodyResultDataSourceDetails) *ListInstancesResponseBodyResult {
	s.DataSourceDetails = v
	return s
}

func (s *ListInstancesResponseBodyResult) SetDescription(v string) *ListInstancesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetEdition(v string) *ListInstancesResponseBodyResult {
	s.Edition = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetExpiredTime(v string) *ListInstancesResponseBodyResult {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetInDebt(v bool) *ListInstancesResponseBodyResult {
	s.InDebt = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetInstanceId(v string) *ListInstancesResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetLockMode(v string) *ListInstancesResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetNetwork(v *ListInstancesResponseBodyResultNetwork) *ListInstancesResponseBodyResult {
	s.Network = v
	return s
}

func (s *ListInstancesResponseBodyResult) SetNoQrs(v bool) *ListInstancesResponseBodyResult {
	s.NoQrs = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetResourceGroupId(v string) *ListInstancesResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetSpec(v *ListInstancesResponseBodyResultSpec) *ListInstancesResponseBodyResult {
	s.Spec = v
	return s
}

func (s *ListInstancesResponseBodyResult) SetStatus(v string) *ListInstancesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetTags(v []*ListInstancesResponseBodyResultTags) *ListInstancesResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyResult) SetUpdateTime(v string) *ListInstancesResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetUserName(v string) *ListInstancesResponseBodyResult {
	s.UserName = &v
	return s
}

func (s *ListInstancesResponseBodyResult) SetVersion(v string) *ListInstancesResponseBodyResult {
	s.Version = &v
	return s
}

func (s *ListInstancesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyResultDataSourceDetails struct {
	Config    *ListInstancesResponseBodyResultDataSourceDetailsConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	IndexName *string                                                 `json:"indexName,omitempty" xml:"indexName,omitempty"`
	Type      *string                                                 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListInstancesResponseBodyResultDataSourceDetails) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyResultDataSourceDetails) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultDataSourceDetails) GetConfig() *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	return s.Config
}

func (s *ListInstancesResponseBodyResultDataSourceDetails) GetIndexName() *string {
	return s.IndexName
}

func (s *ListInstancesResponseBodyResultDataSourceDetails) GetType() *string {
	return s.Type
}

func (s *ListInstancesResponseBodyResultDataSourceDetails) SetConfig(v *ListInstancesResponseBodyResultDataSourceDetailsConfig) *ListInstancesResponseBodyResultDataSourceDetails {
	s.Config = v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetails) SetIndexName(v string) *ListInstancesResponseBodyResultDataSourceDetails {
	s.IndexName = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetails) SetType(v string) *ListInstancesResponseBodyResultDataSourceDetails {
	s.Type = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetails) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyResultDataSourceDetailsConfig struct {
	AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty"`
	Bucket    *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	Catalog   *string `json:"catalog,omitempty" xml:"catalog,omitempty"`
	Database  *string `json:"database,omitempty" xml:"database,omitempty"`
	Endpoint  *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	OssPath   *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	Partition *string `json:"partition,omitempty" xml:"partition,omitempty"`
	Path      *string `json:"path,omitempty" xml:"path,omitempty"`
	Project   *string `json:"project,omitempty" xml:"project,omitempty"`
	Table     *string `json:"table,omitempty" xml:"table,omitempty"`
	Tag       *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListInstancesResponseBodyResultDataSourceDetailsConfig) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyResultDataSourceDetailsConfig) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetAccessKey() *string {
	return s.AccessKey
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetBucket() *string {
	return s.Bucket
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetCatalog() *string {
	return s.Catalog
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetDatabase() *string {
	return s.Database
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetNamespace() *string {
	return s.Namespace
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetOssPath() *string {
	return s.OssPath
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetPartition() *string {
	return s.Partition
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetPath() *string {
	return s.Path
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetProject() *string {
	return s.Project
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetTable() *string {
	return s.Table
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) GetTag() *string {
	return s.Tag
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetAccessKey(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.AccessKey = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetBucket(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Bucket = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetCatalog(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Catalog = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetDatabase(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Database = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetEndpoint(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetNamespace(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Namespace = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetOssPath(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.OssPath = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetPartition(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Partition = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetPath(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Path = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetProject(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Project = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetTable(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Table = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) SetTag(v string) *ListInstancesResponseBodyResultDataSourceDetailsConfig {
	s.Tag = &v
	return s
}

func (s *ListInstancesResponseBodyResultDataSourceDetailsConfig) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyResultNetwork struct {
	Allow *string `json:"allow,omitempty" xml:"allow,omitempty"`
	// The instance endpoint.
	//
	// example:
	//
	// ""
	Endpoint       *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	PublicEndpoint *string `json:"publicEndpoint,omitempty" xml:"publicEndpoint,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp11ldcf59q2nbwkqgj6z
	VSwitchId *string `json:"vSwitchId,omitempty" xml:"vSwitchId,omitempty"`
	// The ID of the virtual private cloud (VPC) in which the instance is deployed.
	//
	// example:
	//
	// vpc-wz9axk41d9vffoc79x0oe
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s ListInstancesResponseBodyResultNetwork) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyResultNetwork) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultNetwork) GetAllow() *string {
	return s.Allow
}

func (s *ListInstancesResponseBodyResultNetwork) GetEndpoint() *string {
	return s.Endpoint
}

func (s *ListInstancesResponseBodyResultNetwork) GetPublicEndpoint() *string {
	return s.PublicEndpoint
}

func (s *ListInstancesResponseBodyResultNetwork) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListInstancesResponseBodyResultNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *ListInstancesResponseBodyResultNetwork) SetAllow(v string) *ListInstancesResponseBodyResultNetwork {
	s.Allow = &v
	return s
}

func (s *ListInstancesResponseBodyResultNetwork) SetEndpoint(v string) *ListInstancesResponseBodyResultNetwork {
	s.Endpoint = &v
	return s
}

func (s *ListInstancesResponseBodyResultNetwork) SetPublicEndpoint(v string) *ListInstancesResponseBodyResultNetwork {
	s.PublicEndpoint = &v
	return s
}

func (s *ListInstancesResponseBodyResultNetwork) SetVSwitchId(v string) *ListInstancesResponseBodyResultNetwork {
	s.VSwitchId = &v
	return s
}

func (s *ListInstancesResponseBodyResultNetwork) SetVpcId(v string) *ListInstancesResponseBodyResultNetwork {
	s.VpcId = &v
	return s
}

func (s *ListInstancesResponseBodyResultNetwork) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyResultSpec struct {
	QrsResource    *ListInstancesResponseBodyResultSpecQrsResource    `json:"qrsResource,omitempty" xml:"qrsResource,omitempty" type:"Struct"`
	SearchResource *ListInstancesResponseBodyResultSpecSearchResource `json:"searchResource,omitempty" xml:"searchResource,omitempty" type:"Struct"`
}

func (s ListInstancesResponseBodyResultSpec) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyResultSpec) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultSpec) GetQrsResource() *ListInstancesResponseBodyResultSpecQrsResource {
	return s.QrsResource
}

func (s *ListInstancesResponseBodyResultSpec) GetSearchResource() *ListInstancesResponseBodyResultSpecSearchResource {
	return s.SearchResource
}

func (s *ListInstancesResponseBodyResultSpec) SetQrsResource(v *ListInstancesResponseBodyResultSpecQrsResource) *ListInstancesResponseBodyResultSpec {
	s.QrsResource = v
	return s
}

func (s *ListInstancesResponseBodyResultSpec) SetSearchResource(v *ListInstancesResponseBodyResultSpecSearchResource) *ListInstancesResponseBodyResultSpec {
	s.SearchResource = v
	return s
}

func (s *ListInstancesResponseBodyResultSpec) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyResultSpecQrsResource struct {
	Category  *string `json:"category,omitempty" xml:"category,omitempty"`
	Cpu       *int32  `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Disk      *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	Mem       *int32  `json:"mem,omitempty" xml:"mem,omitempty"`
	NodeCount *int32  `json:"nodeCount,omitempty" xml:"nodeCount,omitempty"`
}

func (s ListInstancesResponseBodyResultSpecQrsResource) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyResultSpecQrsResource) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) GetCategory() *string {
	return s.Category
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) GetDisk() *int32 {
	return s.Disk
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) GetMem() *int32 {
	return s.Mem
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) SetCategory(v string) *ListInstancesResponseBodyResultSpecQrsResource {
	s.Category = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) SetCpu(v int32) *ListInstancesResponseBodyResultSpecQrsResource {
	s.Cpu = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) SetDisk(v int32) *ListInstancesResponseBodyResultSpecQrsResource {
	s.Disk = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) SetMem(v int32) *ListInstancesResponseBodyResultSpecQrsResource {
	s.Mem = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) SetNodeCount(v int32) *ListInstancesResponseBodyResultSpecQrsResource {
	s.NodeCount = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecQrsResource) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyResultSpecSearchResource struct {
	Category  *string `json:"category,omitempty" xml:"category,omitempty"`
	Cpu       *int32  `json:"cpu,omitempty" xml:"cpu,omitempty"`
	Disk      *int32  `json:"disk,omitempty" xml:"disk,omitempty"`
	Mem       *int32  `json:"mem,omitempty" xml:"mem,omitempty"`
	NodeCount *int32  `json:"nodeCount,omitempty" xml:"nodeCount,omitempty"`
}

func (s ListInstancesResponseBodyResultSpecSearchResource) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyResultSpecSearchResource) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) GetCategory() *string {
	return s.Category
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) GetDisk() *int32 {
	return s.Disk
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) GetMem() *int32 {
	return s.Mem
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) SetCategory(v string) *ListInstancesResponseBodyResultSpecSearchResource {
	s.Category = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) SetCpu(v int32) *ListInstancesResponseBodyResultSpecSearchResource {
	s.Cpu = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) SetDisk(v int32) *ListInstancesResponseBodyResultSpecSearchResource {
	s.Disk = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) SetMem(v int32) *ListInstancesResponseBodyResultSpecSearchResource {
	s.Mem = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) SetNodeCount(v int32) *ListInstancesResponseBodyResultSpecSearchResource {
	s.NodeCount = &v
	return s
}

func (s *ListInstancesResponseBodyResultSpecSearchResource) Validate() error {
	return dara.Validate(s)
}

type ListInstancesResponseBodyResultTags struct {
	// The tag key.
	//
	// example:
	//
	// env
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// oboms-disk
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListInstancesResponseBodyResultTags) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyResultTags) GetKey() *string {
	return s.Key
}

func (s *ListInstancesResponseBodyResultTags) GetValue() *string {
	return s.Value
}

func (s *ListInstancesResponseBodyResultTags) SetKey(v string) *ListInstancesResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyResultTags) SetValue(v string) *ListInstancesResponseBodyResultTags {
	s.Value = &v
	return s
}

func (s *ListInstancesResponseBodyResultTags) Validate() error {
	return dara.Validate(s)
}
