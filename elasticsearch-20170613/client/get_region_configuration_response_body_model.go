// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRegionConfigurationResponseBody
	GetRequestId() *string
	SetResult(v *GetRegionConfigurationResponseBodyResult) *GetRegionConfigurationResponseBody
	GetResult() *GetRegionConfigurationResponseBodyResult
}

type GetRegionConfigurationResponseBody struct {
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6F******
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetRegionConfigurationResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s GetRegionConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegionConfigurationResponseBody) GetResult() *GetRegionConfigurationResponseBodyResult {
	return s.Result
}

func (s *GetRegionConfigurationResponseBody) SetRequestId(v string) *GetRegionConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegionConfigurationResponseBody) SetResult(v *GetRegionConfigurationResponseBodyResult) *GetRegionConfigurationResponseBody {
	s.Result = v
	return s
}

func (s *GetRegionConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResult struct {
	ClientNodeAmountRange *GetRegionConfigurationResponseBodyResultClientNodeAmountRange `json:"clientNodeAmountRange,omitempty" xml:"clientNodeAmountRange,omitempty" type:"Struct"`
	ClientNodeDiskList    []*GetRegionConfigurationResponseBodyResultClientNodeDiskList  `json:"clientNodeDiskList,omitempty" xml:"clientNodeDiskList,omitempty" type:"Repeated"`
	ClientNodeSpec        []*string                                                      `json:"clientNodeSpec,omitempty" xml:"clientNodeSpec,omitempty" type:"Repeated"`
	// example:
	//
	// https://common-buy.aliyun.com/?commodityCode=elasticsearch&orderType=BUY#/buy
	CreateUrl             *string                                                        `json:"createUrl,omitempty" xml:"createUrl,omitempty"`
	DataDiskList          []*GetRegionConfigurationResponseBodyResultDataDiskList        `json:"dataDiskList,omitempty" xml:"dataDiskList,omitempty" type:"Repeated"`
	ElasticNodeProperties *GetRegionConfigurationResponseBodyResultElasticNodeProperties `json:"elasticNodeProperties,omitempty" xml:"elasticNodeProperties,omitempty" type:"Struct"`
	// example:
	//
	// production
	Env                  *string                                                         `json:"env,omitempty" xml:"env,omitempty"`
	EsVersions           []*string                                                       `json:"esVersions,omitempty" xml:"esVersions,omitempty" type:"Repeated"`
	EsVersionsLatestList []*GetRegionConfigurationResponseBodyResultEsVersionsLatestList `json:"esVersionsLatestList,omitempty" xml:"esVersionsLatestList,omitempty" type:"Repeated"`
	InstanceSupportNodes []*string                                                       `json:"instanceSupportNodes,omitempty" xml:"instanceSupportNodes,omitempty" type:"Repeated"`
	JvmConfine           *GetRegionConfigurationResponseBodyResultJvmConfine             `json:"jvmConfine,omitempty" xml:"jvmConfine,omitempty" type:"Struct"`
	KibanaNodeProperties *GetRegionConfigurationResponseBodyResultKibanaNodeProperties   `json:"kibanaNodeProperties,omitempty" xml:"kibanaNodeProperties,omitempty" type:"Struct"`
	LogstashZones        []*string                                                       `json:"logstashZones,omitempty" xml:"logstashZones,omitempty" type:"Repeated"`
	MasterDiskList       []*GetRegionConfigurationResponseBodyResultMasterDiskList       `json:"masterDiskList,omitempty" xml:"masterDiskList,omitempty" type:"Repeated"`
	MasterSpec           []*string                                                       `json:"masterSpec,omitempty" xml:"masterSpec,omitempty" type:"Repeated"`
	Node                 *GetRegionConfigurationResponseBodyResultNode                   `json:"node,omitempty" xml:"node,omitempty" type:"Struct"`
	NodeSpecList         []*GetRegionConfigurationResponseBodyResultNodeSpecList         `json:"nodeSpecList,omitempty" xml:"nodeSpecList,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou
	RegionId           *string                                                     `json:"regionId,omitempty" xml:"regionId,omitempty"`
	SupportVersions    []*GetRegionConfigurationResponseBodyResultSupportVersions  `json:"supportVersions,omitempty" xml:"supportVersions,omitempty" type:"Repeated"`
	WarmNodeProperties *GetRegionConfigurationResponseBodyResultWarmNodeProperties `json:"warmNodeProperties,omitempty" xml:"warmNodeProperties,omitempty" type:"Struct"`
	Zones              []*string                                                   `json:"zones,omitempty" xml:"zones,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResult) GetClientNodeAmountRange() *GetRegionConfigurationResponseBodyResultClientNodeAmountRange {
	return s.ClientNodeAmountRange
}

func (s *GetRegionConfigurationResponseBodyResult) GetClientNodeDiskList() []*GetRegionConfigurationResponseBodyResultClientNodeDiskList {
	return s.ClientNodeDiskList
}

func (s *GetRegionConfigurationResponseBodyResult) GetClientNodeSpec() []*string {
	return s.ClientNodeSpec
}

func (s *GetRegionConfigurationResponseBodyResult) GetCreateUrl() *string {
	return s.CreateUrl
}

func (s *GetRegionConfigurationResponseBodyResult) GetDataDiskList() []*GetRegionConfigurationResponseBodyResultDataDiskList {
	return s.DataDiskList
}

func (s *GetRegionConfigurationResponseBodyResult) GetElasticNodeProperties() *GetRegionConfigurationResponseBodyResultElasticNodeProperties {
	return s.ElasticNodeProperties
}

func (s *GetRegionConfigurationResponseBodyResult) GetEnv() *string {
	return s.Env
}

func (s *GetRegionConfigurationResponseBodyResult) GetEsVersions() []*string {
	return s.EsVersions
}

func (s *GetRegionConfigurationResponseBodyResult) GetEsVersionsLatestList() []*GetRegionConfigurationResponseBodyResultEsVersionsLatestList {
	return s.EsVersionsLatestList
}

func (s *GetRegionConfigurationResponseBodyResult) GetInstanceSupportNodes() []*string {
	return s.InstanceSupportNodes
}

func (s *GetRegionConfigurationResponseBodyResult) GetJvmConfine() *GetRegionConfigurationResponseBodyResultJvmConfine {
	return s.JvmConfine
}

func (s *GetRegionConfigurationResponseBodyResult) GetKibanaNodeProperties() *GetRegionConfigurationResponseBodyResultKibanaNodeProperties {
	return s.KibanaNodeProperties
}

func (s *GetRegionConfigurationResponseBodyResult) GetLogstashZones() []*string {
	return s.LogstashZones
}

func (s *GetRegionConfigurationResponseBodyResult) GetMasterDiskList() []*GetRegionConfigurationResponseBodyResultMasterDiskList {
	return s.MasterDiskList
}

func (s *GetRegionConfigurationResponseBodyResult) GetMasterSpec() []*string {
	return s.MasterSpec
}

func (s *GetRegionConfigurationResponseBodyResult) GetNode() *GetRegionConfigurationResponseBodyResultNode {
	return s.Node
}

func (s *GetRegionConfigurationResponseBodyResult) GetNodeSpecList() []*GetRegionConfigurationResponseBodyResultNodeSpecList {
	return s.NodeSpecList
}

func (s *GetRegionConfigurationResponseBodyResult) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRegionConfigurationResponseBodyResult) GetSupportVersions() []*GetRegionConfigurationResponseBodyResultSupportVersions {
	return s.SupportVersions
}

func (s *GetRegionConfigurationResponseBodyResult) GetWarmNodeProperties() *GetRegionConfigurationResponseBodyResultWarmNodeProperties {
	return s.WarmNodeProperties
}

func (s *GetRegionConfigurationResponseBodyResult) GetZones() []*string {
	return s.Zones
}

func (s *GetRegionConfigurationResponseBodyResult) SetClientNodeAmountRange(v *GetRegionConfigurationResponseBodyResultClientNodeAmountRange) *GetRegionConfigurationResponseBodyResult {
	s.ClientNodeAmountRange = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetClientNodeDiskList(v []*GetRegionConfigurationResponseBodyResultClientNodeDiskList) *GetRegionConfigurationResponseBodyResult {
	s.ClientNodeDiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetClientNodeSpec(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.ClientNodeSpec = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetCreateUrl(v string) *GetRegionConfigurationResponseBodyResult {
	s.CreateUrl = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetDataDiskList(v []*GetRegionConfigurationResponseBodyResultDataDiskList) *GetRegionConfigurationResponseBodyResult {
	s.DataDiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetElasticNodeProperties(v *GetRegionConfigurationResponseBodyResultElasticNodeProperties) *GetRegionConfigurationResponseBodyResult {
	s.ElasticNodeProperties = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetEnv(v string) *GetRegionConfigurationResponseBodyResult {
	s.Env = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetEsVersions(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.EsVersions = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetEsVersionsLatestList(v []*GetRegionConfigurationResponseBodyResultEsVersionsLatestList) *GetRegionConfigurationResponseBodyResult {
	s.EsVersionsLatestList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetInstanceSupportNodes(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.InstanceSupportNodes = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetJvmConfine(v *GetRegionConfigurationResponseBodyResultJvmConfine) *GetRegionConfigurationResponseBodyResult {
	s.JvmConfine = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetKibanaNodeProperties(v *GetRegionConfigurationResponseBodyResultKibanaNodeProperties) *GetRegionConfigurationResponseBodyResult {
	s.KibanaNodeProperties = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetLogstashZones(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.LogstashZones = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetMasterDiskList(v []*GetRegionConfigurationResponseBodyResultMasterDiskList) *GetRegionConfigurationResponseBodyResult {
	s.MasterDiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetMasterSpec(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.MasterSpec = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetNode(v *GetRegionConfigurationResponseBodyResultNode) *GetRegionConfigurationResponseBodyResult {
	s.Node = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetNodeSpecList(v []*GetRegionConfigurationResponseBodyResultNodeSpecList) *GetRegionConfigurationResponseBodyResult {
	s.NodeSpecList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetRegionId(v string) *GetRegionConfigurationResponseBodyResult {
	s.RegionId = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetSupportVersions(v []*GetRegionConfigurationResponseBodyResultSupportVersions) *GetRegionConfigurationResponseBodyResult {
	s.SupportVersions = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetWarmNodeProperties(v *GetRegionConfigurationResponseBodyResultWarmNodeProperties) *GetRegionConfigurationResponseBodyResult {
	s.WarmNodeProperties = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) SetZones(v []*string) *GetRegionConfigurationResponseBodyResult {
	s.Zones = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultClientNodeAmountRange struct {
	// example:
	//
	// 25
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	// example:
	//
	// 2
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultClientNodeAmountRange) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultClientNodeAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeAmountRange) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeAmountRange) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeAmountRange) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultClientNodeAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeAmountRange) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultClientNodeAmountRange {
	s.MinAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeAmountRange) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultClientNodeDiskList struct {
	// example:
	//
	// cloud_efficiency
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// 20
	MaxSize *int32 `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	// example:
	//
	// 20
	MinSize *int32 `json:"minSize,omitempty" xml:"minSize,omitempty"`
	// example:
	//
	// 18
	ScaleLimit *int32 `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultClientNodeDiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultClientNodeDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) GetScaleLimit() *int32 {
	return s.ScaleLimit
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultClientNodeDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultClientNodeDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultClientNodeDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultClientNodeDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultClientNodeDiskList) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultDataDiskList struct {
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// 5120
	MaxSize *int32 `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	// example:
	//
	// 20
	MinSize *int32 `json:"minSize,omitempty" xml:"minSize,omitempty"`
	// example:
	//
	// 2048
	ScaleLimit    *int32    `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
	ValueLimitSet []*string `json:"valueLimitSet,omitempty" xml:"valueLimitSet,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultDataDiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultDataDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) GetScaleLimit() *int32 {
	return s.ScaleLimit
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) GetValueLimitSet() []*string {
	return s.ValueLimitSet
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) SetValueLimitSet(v []*string) *GetRegionConfigurationResponseBodyResultDataDiskList {
	s.ValueLimitSet = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultDataDiskList) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultElasticNodeProperties struct {
	AmountRange *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange `json:"amountRange,omitempty" xml:"amountRange,omitempty" type:"Struct"`
	DiskList    []*GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList  `json:"diskList,omitempty" xml:"diskList,omitempty" type:"Repeated"`
	Spec        []*string                                                                 `json:"spec,omitempty" xml:"spec,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultElasticNodeProperties) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultElasticNodeProperties) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) GetAmountRange() *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange {
	return s.AmountRange
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) GetDiskList() []*GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	return s.DiskList
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) GetSpec() []*string {
	return s.Spec
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) SetAmountRange(v *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) *GetRegionConfigurationResponseBodyResultElasticNodeProperties {
	s.AmountRange = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) SetDiskList(v []*GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) *GetRegionConfigurationResponseBodyResultElasticNodeProperties {
	s.DiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) SetSpec(v []*string) *GetRegionConfigurationResponseBodyResultElasticNodeProperties {
	s.Spec = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodeProperties) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange struct {
	// example:
	//
	// 25
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	// example:
	//
	// 2
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange {
	s.MinAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesAmountRange) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList struct {
	// example:
	//
	// true
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// 5120
	MaxSize *int32 `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	// example:
	//
	// 500
	MinSize *int32 `json:"minSize,omitempty" xml:"minSize,omitempty"`
	// example:
	//
	// 2048
	ScaleLimit    *int32    `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
	ValueLimitSet []*string `json:"valueLimitSet,omitempty" xml:"valueLimitSet,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) GetScaleLimit() *int32 {
	return s.ScaleLimit
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) GetValueLimitSet() []*string {
	return s.ValueLimitSet
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetDiskEncryption(v bool) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.DiskEncryption = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) SetValueLimitSet(v []*string) *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList {
	s.ValueLimitSet = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultElasticNodePropertiesDiskList) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultEsVersionsLatestList struct {
	// example:
	//
	// 5.5_with_X-Pack
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 5.5.3_with_X-Pack
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultEsVersionsLatestList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultEsVersionsLatestList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultEsVersionsLatestList) GetKey() *string {
	return s.Key
}

func (s *GetRegionConfigurationResponseBodyResultEsVersionsLatestList) GetValue() *string {
	return s.Value
}

func (s *GetRegionConfigurationResponseBodyResultEsVersionsLatestList) SetKey(v string) *GetRegionConfigurationResponseBodyResultEsVersionsLatestList {
	s.Key = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultEsVersionsLatestList) SetValue(v string) *GetRegionConfigurationResponseBodyResultEsVersionsLatestList {
	s.Value = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultEsVersionsLatestList) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultJvmConfine struct {
	// example:
	//
	// 32
	Memory            *int32    `json:"memory,omitempty" xml:"memory,omitempty"`
	SupportEsVersions []*string `json:"supportEsVersions,omitempty" xml:"supportEsVersions,omitempty" type:"Repeated"`
	SupportGcs        []*string `json:"supportGcs,omitempty" xml:"supportGcs,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultJvmConfine) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultJvmConfine) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) GetMemory() *int32 {
	return s.Memory
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) GetSupportEsVersions() []*string {
	return s.SupportEsVersions
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) GetSupportGcs() []*string {
	return s.SupportGcs
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) SetMemory(v int32) *GetRegionConfigurationResponseBodyResultJvmConfine {
	s.Memory = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) SetSupportEsVersions(v []*string) *GetRegionConfigurationResponseBodyResultJvmConfine {
	s.SupportEsVersions = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) SetSupportGcs(v []*string) *GetRegionConfigurationResponseBodyResultJvmConfine {
	s.SupportGcs = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultJvmConfine) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultKibanaNodeProperties struct {
	AmountRange *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange `json:"amountRange,omitempty" xml:"amountRange,omitempty" type:"Struct"`
	Spec        []*string                                                                `json:"spec,omitempty" xml:"spec,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultKibanaNodeProperties) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultKibanaNodeProperties) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodeProperties) GetAmountRange() *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange {
	return s.AmountRange
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodeProperties) GetSpec() []*string {
	return s.Spec
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodeProperties) SetAmountRange(v *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) *GetRegionConfigurationResponseBodyResultKibanaNodeProperties {
	s.AmountRange = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodeProperties) SetSpec(v []*string) *GetRegionConfigurationResponseBodyResultKibanaNodeProperties {
	s.Spec = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodeProperties) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange struct {
	// example:
	//
	// 20
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	// example:
	//
	// 1
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange {
	s.MinAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultKibanaNodePropertiesAmountRange) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultMasterDiskList struct {
	// example:
	//
	// cloud_ssd
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// 20
	MaxSize *int32 `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	// example:
	//
	// 20
	MinSize *int32 `json:"minSize,omitempty" xml:"minSize,omitempty"`
	// example:
	//
	// 20
	ScaleLimit *int32 `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultMasterDiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultMasterDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) GetScaleLimit() *int32 {
	return s.ScaleLimit
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultMasterDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultMasterDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultMasterDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultMasterDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultMasterDiskList) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultNode struct {
	// example:
	//
	// 50
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	// example:
	//
	// 2
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultNode) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultNode) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultNode) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetRegionConfigurationResponseBodyResultNode) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *GetRegionConfigurationResponseBodyResultNode) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultNode {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNode) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultNode {
	s.MinAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNode) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultNodeSpecList struct {
	// example:
	//
	// 16
	CpuCount *int32 `json:"cpuCount,omitempty" xml:"cpuCount,omitempty"`
	// example:
	//
	// 44000
	Disk *int32 `json:"disk,omitempty" xml:"disk,omitempty"`
	// example:
	//
	// local_efficiency
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 64
	MemorySize *int32 `json:"memorySize,omitempty" xml:"memorySize,omitempty"`
	// example:
	//
	// elasticsearch.sn2ne.large
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
	// example:
	//
	// local_efficiency
	SpecGroupType *string `json:"specGroupType,omitempty" xml:"specGroupType,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultNodeSpecList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultNodeSpecList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) GetCpuCount() *int32 {
	return s.CpuCount
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) GetDisk() *int32 {
	return s.Disk
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) GetEnable() *bool {
	return s.Enable
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) GetSpec() *string {
	return s.Spec
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) GetSpecGroupType() *string {
	return s.SpecGroupType
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetCpuCount(v int32) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.CpuCount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetDisk(v int32) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.Disk = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetEnable(v bool) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.Enable = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetMemorySize(v int32) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.MemorySize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetSpec(v string) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.Spec = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) SetSpecGroupType(v string) *GetRegionConfigurationResponseBodyResultNodeSpecList {
	s.SpecGroupType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultNodeSpecList) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultSupportVersions struct {
	// example:
	//
	// x-pack
	InstanceCategory   *string                                                                      `json:"instanceCategory,omitempty" xml:"instanceCategory,omitempty"`
	SupportVersionList []*GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList `json:"supportVersionList,omitempty" xml:"supportVersionList,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultSupportVersions) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultSupportVersions) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersions) GetInstanceCategory() *string {
	return s.InstanceCategory
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersions) GetSupportVersionList() []*GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList {
	return s.SupportVersionList
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersions) SetInstanceCategory(v string) *GetRegionConfigurationResponseBodyResultSupportVersions {
	s.InstanceCategory = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersions) SetSupportVersionList(v []*GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) *GetRegionConfigurationResponseBodyResultSupportVersions {
	s.SupportVersionList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersions) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList struct {
	// example:
	//
	// 5.5
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// 5.5.3
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) GetKey() *string {
	return s.Key
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) GetValue() *string {
	return s.Value
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) SetKey(v string) *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList {
	s.Key = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) SetValue(v string) *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList {
	s.Value = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultSupportVersionsSupportVersionList) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultWarmNodeProperties struct {
	AmountRange *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange `json:"amountRange,omitempty" xml:"amountRange,omitempty" type:"Struct"`
	DiskList    []*GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList  `json:"diskList,omitempty" xml:"diskList,omitempty" type:"Repeated"`
	Spec        []*string                                                              `json:"spec,omitempty" xml:"spec,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultWarmNodeProperties) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultWarmNodeProperties) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) GetAmountRange() *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange {
	return s.AmountRange
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) GetDiskList() []*GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	return s.DiskList
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) GetSpec() []*string {
	return s.Spec
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) SetAmountRange(v *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) *GetRegionConfigurationResponseBodyResultWarmNodeProperties {
	s.AmountRange = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) SetDiskList(v []*GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) *GetRegionConfigurationResponseBodyResultWarmNodeProperties {
	s.DiskList = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) SetSpec(v []*string) *GetRegionConfigurationResponseBodyResultWarmNodeProperties {
	s.Spec = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodeProperties) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange struct {
	// example:
	//
	// 50
	MaxAmount *int32 `json:"maxAmount,omitempty" xml:"maxAmount,omitempty"`
	// example:
	//
	// 2
	MinAmount *int32 `json:"minAmount,omitempty" xml:"minAmount,omitempty"`
}

func (s GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) GetMaxAmount() *int32 {
	return s.MaxAmount
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) SetMaxAmount(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange {
	s.MaxAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) SetMinAmount(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange {
	s.MinAmount = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesAmountRange) Validate() error {
	return dara.Validate(s)
}

type GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList struct {
	// example:
	//
	// true
	DiskEncryption *bool `json:"diskEncryption,omitempty" xml:"diskEncryption,omitempty"`
	// example:
	//
	// cloud_efficiency
	DiskType *string `json:"diskType,omitempty" xml:"diskType,omitempty"`
	// example:
	//
	// 5120
	MaxSize *int32 `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	// example:
	//
	// 500
	MinSize *int32 `json:"minSize,omitempty" xml:"minSize,omitempty"`
	// example:
	//
	// 2048
	ScaleLimit    *int32    `json:"scaleLimit,omitempty" xml:"scaleLimit,omitempty"`
	ValueLimitSet []*string `json:"valueLimitSet,omitempty" xml:"valueLimitSet,omitempty" type:"Repeated"`
}

func (s GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) String() string {
	return dara.Prettify(s)
}

func (s GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) GoString() string {
	return s.String()
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) GetDiskEncryption() *bool {
	return s.DiskEncryption
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) GetDiskType() *string {
	return s.DiskType
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) GetMaxSize() *int32 {
	return s.MaxSize
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) GetMinSize() *int32 {
	return s.MinSize
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) GetScaleLimit() *int32 {
	return s.ScaleLimit
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) GetValueLimitSet() []*string {
	return s.ValueLimitSet
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetDiskEncryption(v bool) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.DiskEncryption = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetDiskType(v string) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.DiskType = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetMaxSize(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.MaxSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetMinSize(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.MinSize = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetScaleLimit(v int32) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.ScaleLimit = &v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) SetValueLimitSet(v []*string) *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList {
	s.ValueLimitSet = v
	return s
}

func (s *GetRegionConfigurationResponseBodyResultWarmNodePropertiesDiskList) Validate() error {
	return dara.Validate(s)
}
