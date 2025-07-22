// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoScalingConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v *ModifyAutoScalingConfigRequestBandwidth) *ModifyAutoScalingConfigRequest
	GetBandwidth() *ModifyAutoScalingConfigRequestBandwidth
	SetInstanceId(v string) *ModifyAutoScalingConfigRequest
	GetInstanceId() *string
	SetResource(v *ModifyAutoScalingConfigRequestResource) *ModifyAutoScalingConfigRequest
	GetResource() *ModifyAutoScalingConfigRequestResource
	SetShard(v *ModifyAutoScalingConfigRequestShard) *ModifyAutoScalingConfigRequest
	GetShard() *ModifyAutoScalingConfigRequestShard
	SetSpec(v *ModifyAutoScalingConfigRequestSpec) *ModifyAutoScalingConfigRequest
	GetSpec() *ModifyAutoScalingConfigRequestSpec
	SetStorage(v *ModifyAutoScalingConfigRequestStorage) *ModifyAutoScalingConfigRequest
	GetStorage() *ModifyAutoScalingConfigRequestStorage
}

type ModifyAutoScalingConfigRequest struct {
	// The configuration item of the bandwidth auto scaling feature.
	Bandwidth *ModifyAutoScalingConfigRequestBandwidth `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty" type:"Struct"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configuration item of the resource auto scaling feature.
	Resource *ModifyAutoScalingConfigRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The configuration item of the shard auto scaling feature.
	Shard *ModifyAutoScalingConfigRequestShard `json:"Shard,omitempty" xml:"Shard,omitempty" type:"Struct"`
	// The configuration item of the specification auto scaling feature.
	Spec *ModifyAutoScalingConfigRequestSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	// The configuration item of the automatic storage expansion feature.
	Storage *ModifyAutoScalingConfigRequestStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
}

func (s ModifyAutoScalingConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigRequest) GetBandwidth() *ModifyAutoScalingConfigRequestBandwidth {
	return s.Bandwidth
}

func (s *ModifyAutoScalingConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyAutoScalingConfigRequest) GetResource() *ModifyAutoScalingConfigRequestResource {
	return s.Resource
}

func (s *ModifyAutoScalingConfigRequest) GetShard() *ModifyAutoScalingConfigRequestShard {
	return s.Shard
}

func (s *ModifyAutoScalingConfigRequest) GetSpec() *ModifyAutoScalingConfigRequestSpec {
	return s.Spec
}

func (s *ModifyAutoScalingConfigRequest) GetStorage() *ModifyAutoScalingConfigRequestStorage {
	return s.Storage
}

func (s *ModifyAutoScalingConfigRequest) SetBandwidth(v *ModifyAutoScalingConfigRequestBandwidth) *ModifyAutoScalingConfigRequest {
	s.Bandwidth = v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetInstanceId(v string) *ModifyAutoScalingConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetResource(v *ModifyAutoScalingConfigRequestResource) *ModifyAutoScalingConfigRequest {
	s.Resource = v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetShard(v *ModifyAutoScalingConfigRequestShard) *ModifyAutoScalingConfigRequest {
	s.Shard = v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetSpec(v *ModifyAutoScalingConfigRequestSpec) *ModifyAutoScalingConfigRequest {
	s.Spec = v
	return s
}

func (s *ModifyAutoScalingConfigRequest) SetStorage(v *ModifyAutoScalingConfigRequestStorage) *ModifyAutoScalingConfigRequest {
	s.Storage = v
	return s
}

func (s *ModifyAutoScalingConfigRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyAutoScalingConfigRequestBandwidth struct {
	// Specifies whether to apply the **Bandwidth*	- configuration of the bandwidth auto scaling feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// The average bandwidth usage threshold that triggers automatic bandwidth downgrade. Unit: %. Valid values:
	//
	// 	- **10**
	//
	// 	- **20**
	//
	// 	- **30**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 30
	BandwidthUsageLowerThreshold *int32 `json:"BandwidthUsageLowerThreshold,omitempty" xml:"BandwidthUsageLowerThreshold,omitempty"`
	// The average bandwidth usage threshold that triggers automatic bandwidth upgrade. Unit: %. Valid values:
	//
	// 	- **50**
	//
	// 	- **60**
	//
	// 	- **70**
	//
	// 	- **80**
	//
	// 	- **90**
	//
	// 	- **95**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	BandwidthUsageUpperThreshold *int32 `json:"BandwidthUsageUpperThreshold,omitempty" xml:"BandwidthUsageUpperThreshold,omitempty"`
	// Specifies whether to enable automatic bandwidth downgrade. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Downgrade *bool `json:"Downgrade,omitempty" xml:"Downgrade,omitempty"`
	// The observation window of the bandwidth auto scaling feature. The value of this parameter consists of a numeric value and a time unit suffix. The **m*	- time unit suffix specifies the minute. Valid values:
	//
	// 	- **1m**
	//
	// 	- **5m**
	//
	// 	- **10m**
	//
	// 	- **15m**
	//
	// 	- **30m**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	ObservationWindowSize *string `json:"ObservationWindowSize,omitempty" xml:"ObservationWindowSize,omitempty"`
	// Specifies whether to enable automatic bandwidth upgrade. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Upgrade *bool `json:"Upgrade,omitempty" xml:"Upgrade,omitempty"`
}

func (s ModifyAutoScalingConfigRequestBandwidth) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigRequestBandwidth) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigRequestBandwidth) GetApply() *bool {
	return s.Apply
}

func (s *ModifyAutoScalingConfigRequestBandwidth) GetBandwidthUsageLowerThreshold() *int32 {
	return s.BandwidthUsageLowerThreshold
}

func (s *ModifyAutoScalingConfigRequestBandwidth) GetBandwidthUsageUpperThreshold() *int32 {
	return s.BandwidthUsageUpperThreshold
}

func (s *ModifyAutoScalingConfigRequestBandwidth) GetDowngrade() *bool {
	return s.Downgrade
}

func (s *ModifyAutoScalingConfigRequestBandwidth) GetObservationWindowSize() *string {
	return s.ObservationWindowSize
}

func (s *ModifyAutoScalingConfigRequestBandwidth) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *ModifyAutoScalingConfigRequestBandwidth) SetApply(v bool) *ModifyAutoScalingConfigRequestBandwidth {
	s.Apply = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestBandwidth) SetBandwidthUsageLowerThreshold(v int32) *ModifyAutoScalingConfigRequestBandwidth {
	s.BandwidthUsageLowerThreshold = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestBandwidth) SetBandwidthUsageUpperThreshold(v int32) *ModifyAutoScalingConfigRequestBandwidth {
	s.BandwidthUsageUpperThreshold = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestBandwidth) SetDowngrade(v bool) *ModifyAutoScalingConfigRequestBandwidth {
	s.Downgrade = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestBandwidth) SetObservationWindowSize(v string) *ModifyAutoScalingConfigRequestBandwidth {
	s.ObservationWindowSize = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestBandwidth) SetUpgrade(v bool) *ModifyAutoScalingConfigRequestBandwidth {
	s.Upgrade = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestBandwidth) Validate() error {
	return dara.Validate(s)
}

type ModifyAutoScalingConfigRequestResource struct {
	// Specifies whether to apply the **Resource*	- configuration of the resource auto scaling feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// The average CPU utilization threshold that triggers automatic resource scale-out. Unit: %. Valid values:
	//
	// 	- **70**
	//
	// 	- **80**
	//
	// 	- **90**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	CpuUsageUpperThreshold *int32 `json:"CpuUsageUpperThreshold,omitempty" xml:"CpuUsageUpperThreshold,omitempty"`
	// The observation window of the automatic resource scale-in feature. The value of this parameter consists of a numeric value and a time unit suffix. The **m*	- time unit suffix specifies the minute. Valid values:
	//
	// 	- **1m**
	//
	// 	- **3m**
	//
	// 	- **5m**
	//
	// 	- **10m**
	//
	// 	- **20m**
	//
	// 	- **30m**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	DowngradeObservationWindowSize *string `json:"DowngradeObservationWindowSize,omitempty" xml:"DowngradeObservationWindowSize,omitempty"`
	// Specifies whether to enable resource auto scaling. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The observation window of the automatic resource scale-out feature. The value of this parameter consists of a numeric value and a time unit suffix. The **m*	- time unit suffix specifies the minute. Valid values:
	//
	// 	- **1m**
	//
	// 	- **3m**
	//
	// 	- **5m**
	//
	// 	- **10m**
	//
	// 	- **20m**
	//
	// 	- **30m**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	UpgradeObservationWindowSize *string `json:"UpgradeObservationWindowSize,omitempty" xml:"UpgradeObservationWindowSize,omitempty"`
}

func (s ModifyAutoScalingConfigRequestResource) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigRequestResource) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigRequestResource) GetApply() *bool {
	return s.Apply
}

func (s *ModifyAutoScalingConfigRequestResource) GetCpuUsageUpperThreshold() *int32 {
	return s.CpuUsageUpperThreshold
}

func (s *ModifyAutoScalingConfigRequestResource) GetDowngradeObservationWindowSize() *string {
	return s.DowngradeObservationWindowSize
}

func (s *ModifyAutoScalingConfigRequestResource) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyAutoScalingConfigRequestResource) GetUpgradeObservationWindowSize() *string {
	return s.UpgradeObservationWindowSize
}

func (s *ModifyAutoScalingConfigRequestResource) SetApply(v bool) *ModifyAutoScalingConfigRequestResource {
	s.Apply = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestResource) SetCpuUsageUpperThreshold(v int32) *ModifyAutoScalingConfigRequestResource {
	s.CpuUsageUpperThreshold = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestResource) SetDowngradeObservationWindowSize(v string) *ModifyAutoScalingConfigRequestResource {
	s.DowngradeObservationWindowSize = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestResource) SetEnable(v bool) *ModifyAutoScalingConfigRequestResource {
	s.Enable = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestResource) SetUpgradeObservationWindowSize(v string) *ModifyAutoScalingConfigRequestResource {
	s.UpgradeObservationWindowSize = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestResource) Validate() error {
	return dara.Validate(s)
}

type ModifyAutoScalingConfigRequestShard struct {
	// Specifies whether to apply the **Shard*	- configuration of the shard auto scaling feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  The shard auto scaling feature is available only for Tair (Redis OSS-compatible) cloud-native cluster instances on the China site (aliyun.com).
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// Specifies whether to enable automatic shard removal. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  The automatic shard removal feature is in a canary release.
	//
	// example:
	//
	// true
	Downgrade *bool `json:"Downgrade,omitempty" xml:"Downgrade,omitempty"`
	// The observation window of the automatic shard removal feature. The value of this parameter consists of a numeric value and a time unit suffix. The **h*	- time unit suffix specifies the hour. The **d*	- time unit suffix specifies the day. Valid values:
	//
	// 	- **1h**
	//
	// 	- **2h**
	//
	// 	- **3h**
	//
	// 	- **1d**
	//
	// 	- **7d**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1h
	DowngradeObservationWindowSize *string `json:"DowngradeObservationWindowSize,omitempty" xml:"DowngradeObservationWindowSize,omitempty"`
	// The maximum number of shards in the instance. The value must be a positive integer. Valid values: 4 to 32.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 16
	MaxShards *int32 `json:"MaxShards,omitempty" xml:"MaxShards,omitempty"`
	// The average memory usage threshold that triggers automatic shard removal. Unit: %. Valid values:
	//
	// 	- **10**
	//
	// 	- **20**
	//
	// 	- **30**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 30
	MemUsageLowerThreshold *int32 `json:"MemUsageLowerThreshold,omitempty" xml:"MemUsageLowerThreshold,omitempty"`
	// The average memory usage threshold that triggers automatic shard addition. Unit: %. Valid values:
	//
	// 	- **50**
	//
	// 	- **60**
	//
	// 	- **70**
	//
	// 	- **80**
	//
	// 	- **90**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	MemUsageUpperThreshold *int32 `json:"MemUsageUpperThreshold,omitempty" xml:"MemUsageUpperThreshold,omitempty"`
	// The minimum number of shards in the instance. The value must be a positive integer. Valid values: 4 to 32.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 4
	MinShards *int32 `json:"MinShards,omitempty" xml:"MinShards,omitempty"`
	// Specifies whether to enable automatic shard addition. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Upgrade *bool `json:"Upgrade,omitempty" xml:"Upgrade,omitempty"`
	// The observation window of the automatic shard addition feature. The value of this parameter consists of a numeric value and a time unit suffix. The **m*	- time unit suffix specifies the minute. Valid values:
	//
	// 	- **5m**
	//
	// 	- **10m**
	//
	// 	- **15m**
	//
	// 	- **30m**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	UpgradeObservationWindowSize *string `json:"UpgradeObservationWindowSize,omitempty" xml:"UpgradeObservationWindowSize,omitempty"`
}

func (s ModifyAutoScalingConfigRequestShard) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigRequestShard) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigRequestShard) GetApply() *bool {
	return s.Apply
}

func (s *ModifyAutoScalingConfigRequestShard) GetDowngrade() *bool {
	return s.Downgrade
}

func (s *ModifyAutoScalingConfigRequestShard) GetDowngradeObservationWindowSize() *string {
	return s.DowngradeObservationWindowSize
}

func (s *ModifyAutoScalingConfigRequestShard) GetMaxShards() *int32 {
	return s.MaxShards
}

func (s *ModifyAutoScalingConfigRequestShard) GetMemUsageLowerThreshold() *int32 {
	return s.MemUsageLowerThreshold
}

func (s *ModifyAutoScalingConfigRequestShard) GetMemUsageUpperThreshold() *int32 {
	return s.MemUsageUpperThreshold
}

func (s *ModifyAutoScalingConfigRequestShard) GetMinShards() *int32 {
	return s.MinShards
}

func (s *ModifyAutoScalingConfigRequestShard) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *ModifyAutoScalingConfigRequestShard) GetUpgradeObservationWindowSize() *string {
	return s.UpgradeObservationWindowSize
}

func (s *ModifyAutoScalingConfigRequestShard) SetApply(v bool) *ModifyAutoScalingConfigRequestShard {
	s.Apply = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestShard) SetDowngrade(v bool) *ModifyAutoScalingConfigRequestShard {
	s.Downgrade = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestShard) SetDowngradeObservationWindowSize(v string) *ModifyAutoScalingConfigRequestShard {
	s.DowngradeObservationWindowSize = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestShard) SetMaxShards(v int32) *ModifyAutoScalingConfigRequestShard {
	s.MaxShards = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestShard) SetMemUsageLowerThreshold(v int32) *ModifyAutoScalingConfigRequestShard {
	s.MemUsageLowerThreshold = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestShard) SetMemUsageUpperThreshold(v int32) *ModifyAutoScalingConfigRequestShard {
	s.MemUsageUpperThreshold = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestShard) SetMinShards(v int32) *ModifyAutoScalingConfigRequestShard {
	s.MinShards = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestShard) SetUpgrade(v bool) *ModifyAutoScalingConfigRequestShard {
	s.Upgrade = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestShard) SetUpgradeObservationWindowSize(v string) *ModifyAutoScalingConfigRequestShard {
	s.UpgradeObservationWindowSize = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestShard) Validate() error {
	return dara.Validate(s)
}

type ModifyAutoScalingConfigRequestSpec struct {
	// Specifies whether to apply the **Spec*	- configuration of the specification auto scaling feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// The quiescent period. The value of this parameter consists of a numeric value and a time unit suffix. The **m*	- time unit suffix specifies the minute, the **h*	- time unit suffix specifies the hour, and the **d*	- time unit suffix specifies the day.
	//
	// 	- Valid values for PolarDB for MySQL Cluster Edition instances: **5m**, **10m**, **30m**, **1h**, **2h**, **3h**, **1d**, and **7d**.
	//
	// 	- Valid values for ApsaraDB RDS for MySQL High-availability Edition instances that use standard SSDs or Enterprise SSDs (ESSDs): **5m**, **10m**, **30m**, **1h**, **2h**, **3h**, **1d**, and **7d**.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	CoolDownTime *string `json:"CoolDownTime,omitempty" xml:"CoolDownTime,omitempty"`
	// The average CPU utilization threshold that triggers automatic specification scale-up. Unit: %. Valid values:
	//
	// 	- **50**
	//
	// 	- **60**
	//
	// 	- **70**
	//
	// 	- **80**
	//
	// 	- **90**
	//
	// >  This parameter must be specified if the database instance is a PolarDB for MySQL Cluster Edition instance or an ApsaraDB RDS for MySQL High-availability Edition instance that uses standard SSDs or ESSDs.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	CpuUsageUpperThreshold *int32 `json:"CpuUsageUpperThreshold,omitempty" xml:"CpuUsageUpperThreshold,omitempty"`
	// Specifies whether to enable automatic specification scale-down. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  This parameter must be specified if the database instance is a PolarDB for MySQL Cluster Edition instance or an ApsaraDB RDS for MySQL High-availability Edition instance that uses standard SSDs or ESSDs.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Downgrade *bool `json:"Downgrade,omitempty" xml:"Downgrade,omitempty"`
	// The maximum number of read-only nodes of the instance.
	//
	// >  This parameter must be specified if the database instance is a PolarDB for MySQL Cluster Edition instance.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 10
	MaxReadOnlyNodes *int32 `json:"MaxReadOnlyNodes,omitempty" xml:"MaxReadOnlyNodes,omitempty"`
	// The maximum specifications to which the database instance can be scaled up. The database instance can be upgraded only to a database instance of the same edition with higher specifications. For information about the specifications of different database instances, see the following topics:
	//
	// 	- PolarDB for MySQL Cluster Edition instances: [Specifications of compute nodes](https://help.aliyun.com/document_detail/102542.html)
	//
	// 	- ApsaraDB RDS for MySQL High-availability Edition instances that use standard SSDs or ESSDs: [Specifications](https://help.aliyun.com/document_detail/276974.html)
	//
	// if can be null:
	// true
	//
	// example:
	//
	// polar.mysql.x8.12xlarge
	MaxSpec *string `json:"MaxSpec,omitempty" xml:"MaxSpec,omitempty"`
	// The average memory usage threshold that triggers automatic specification scale-up. Unit: %. Valid values:
	//
	// 	- **50**
	//
	// 	- **60**
	//
	// 	- **70**
	//
	// 	- **80**
	//
	// 	- **90**
	//
	// >  This parameter must be specified if the database instance is a Tair (Redis OSS-compatible) Community Edition cloud-native instance on the China site (aliyun.com).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	MemUsageUpperThreshold *int32 `json:"MemUsageUpperThreshold,omitempty" xml:"MemUsageUpperThreshold,omitempty"`
	// The observation window. The value of this parameter consists of a numeric value and a time unit suffix. The **m*	- time unit suffix specifies the minute and the **h*	- time unit suffix specifies the hour.
	//
	// 	- Valid values for PolarDB for MySQL Cluster Edition instances: **5m**, **10m**, **15m**, and **30m**.
	//
	// 	- Valid values for ApsaraDB RDS for MySQL High-availability Edition instances that use standard SSDs or ESSDs: **5m**, **20m**, **30m**, **40m**, and **1h**.
	//
	// 	- Valid values for Tair (Redis OSS-compatible) Community Edition cloud-native instances: **5m**, **10m**, **15m**, and **30m**.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	ObservationWindowSize *string `json:"ObservationWindowSize,omitempty" xml:"ObservationWindowSize,omitempty"`
	// Specifies whether to enable automatic specification scale-up. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Upgrade *bool `json:"Upgrade,omitempty" xml:"Upgrade,omitempty"`
}

func (s ModifyAutoScalingConfigRequestSpec) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigRequestSpec) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigRequestSpec) GetApply() *bool {
	return s.Apply
}

func (s *ModifyAutoScalingConfigRequestSpec) GetCoolDownTime() *string {
	return s.CoolDownTime
}

func (s *ModifyAutoScalingConfigRequestSpec) GetCpuUsageUpperThreshold() *int32 {
	return s.CpuUsageUpperThreshold
}

func (s *ModifyAutoScalingConfigRequestSpec) GetDowngrade() *bool {
	return s.Downgrade
}

func (s *ModifyAutoScalingConfigRequestSpec) GetMaxReadOnlyNodes() *int32 {
	return s.MaxReadOnlyNodes
}

func (s *ModifyAutoScalingConfigRequestSpec) GetMaxSpec() *string {
	return s.MaxSpec
}

func (s *ModifyAutoScalingConfigRequestSpec) GetMemUsageUpperThreshold() *int32 {
	return s.MemUsageUpperThreshold
}

func (s *ModifyAutoScalingConfigRequestSpec) GetObservationWindowSize() *string {
	return s.ObservationWindowSize
}

func (s *ModifyAutoScalingConfigRequestSpec) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *ModifyAutoScalingConfigRequestSpec) SetApply(v bool) *ModifyAutoScalingConfigRequestSpec {
	s.Apply = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestSpec) SetCoolDownTime(v string) *ModifyAutoScalingConfigRequestSpec {
	s.CoolDownTime = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestSpec) SetCpuUsageUpperThreshold(v int32) *ModifyAutoScalingConfigRequestSpec {
	s.CpuUsageUpperThreshold = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestSpec) SetDowngrade(v bool) *ModifyAutoScalingConfigRequestSpec {
	s.Downgrade = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestSpec) SetMaxReadOnlyNodes(v int32) *ModifyAutoScalingConfigRequestSpec {
	s.MaxReadOnlyNodes = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestSpec) SetMaxSpec(v string) *ModifyAutoScalingConfigRequestSpec {
	s.MaxSpec = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestSpec) SetMemUsageUpperThreshold(v int32) *ModifyAutoScalingConfigRequestSpec {
	s.MemUsageUpperThreshold = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestSpec) SetObservationWindowSize(v string) *ModifyAutoScalingConfigRequestSpec {
	s.ObservationWindowSize = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestSpec) SetUpgrade(v bool) *ModifyAutoScalingConfigRequestSpec {
	s.Upgrade = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestSpec) Validate() error {
	return dara.Validate(s)
}

type ModifyAutoScalingConfigRequestStorage struct {
	// Specifies whether to apply the **Storage*	- configuration of the automatic storage expansion feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// The average storage usage threshold that triggers automatic storage expansion. Unit: %. Valid values:
	//
	// 	- **50**
	//
	// 	- **60**
	//
	// 	- **70**
	//
	// 	- **80**
	//
	// 	- **90**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	DiskUsageUpperThreshold *int32 `json:"DiskUsageUpperThreshold,omitempty" xml:"DiskUsageUpperThreshold,omitempty"`
	// The maximum storage size of the database instance. Unit: GB. The value must be greater than or equal to the total storage size of the instance.
	//
	// 	- If the instance uses ESSDs, the maximum value of this parameter can be 32000.
	//
	// 	- If the instance uses standard SSDs, the maximum value of this parameter can be 6000.
	//
	// >  The standard SSD storage type is phased out. We recommend that you [upgrade the storage type of your instance from standard SSDs to ESSDs](https://help.aliyun.com/document_detail/314678.html).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 32000
	MaxStorage *int32 `json:"MaxStorage,omitempty" xml:"MaxStorage,omitempty"`
	// Specifies whether to enable automatic storage expansion. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Upgrade *bool `json:"Upgrade,omitempty" xml:"Upgrade,omitempty"`
}

func (s ModifyAutoScalingConfigRequestStorage) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoScalingConfigRequestStorage) GoString() string {
	return s.String()
}

func (s *ModifyAutoScalingConfigRequestStorage) GetApply() *bool {
	return s.Apply
}

func (s *ModifyAutoScalingConfigRequestStorage) GetDiskUsageUpperThreshold() *int32 {
	return s.DiskUsageUpperThreshold
}

func (s *ModifyAutoScalingConfigRequestStorage) GetMaxStorage() *int32 {
	return s.MaxStorage
}

func (s *ModifyAutoScalingConfigRequestStorage) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *ModifyAutoScalingConfigRequestStorage) SetApply(v bool) *ModifyAutoScalingConfigRequestStorage {
	s.Apply = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestStorage) SetDiskUsageUpperThreshold(v int32) *ModifyAutoScalingConfigRequestStorage {
	s.DiskUsageUpperThreshold = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestStorage) SetMaxStorage(v int32) *ModifyAutoScalingConfigRequestStorage {
	s.MaxStorage = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestStorage) SetUpgrade(v bool) *ModifyAutoScalingConfigRequestStorage {
	s.Upgrade = &v
	return s
}

func (s *ModifyAutoScalingConfigRequestStorage) Validate() error {
	return dara.Validate(s)
}
