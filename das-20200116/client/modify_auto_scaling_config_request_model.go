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
	// The configuration items for the bandwidth elastic scaling feature.
	Bandwidth *ModifyAutoScalingConfigRequestBandwidth `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty" type:"Struct"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The configuration items for the resource elastic scaling feature.
	Resource *ModifyAutoScalingConfigRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The configuration items for the shard elastic scaling feature.
	Shard *ModifyAutoScalingConfigRequestShard `json:"Shard,omitempty" xml:"Shard,omitempty" type:"Struct"`
	// The configuration items for the specification elastic scaling feature.
	Spec *ModifyAutoScalingConfigRequestSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	// The configuration items for the automatic storage expansion feature.
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
	if s.Bandwidth != nil {
		if err := s.Bandwidth.Validate(); err != nil {
			return err
		}
	}
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	if s.Shard != nil {
		if err := s.Shard.Validate(); err != nil {
			return err
		}
	}
	if s.Spec != nil {
		if err := s.Spec.Validate(); err != nil {
			return err
		}
	}
	if s.Storage != nil {
		if err := s.Storage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAutoScalingConfigRequestBandwidth struct {
	// Specifies whether to apply the configurations under the **Bandwidth*	- configuration item for the bandwidth elastic scaling feature:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// The threshold of average bandwidth usage that triggers automatic bandwidth downgrade. Unit: %. Valid values:
	//
	// - **10**
	//
	// - **20**
	//
	// - **30**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 30
	BandwidthUsageLowerThreshold *int32 `json:"BandwidthUsageLowerThreshold,omitempty" xml:"BandwidthUsageLowerThreshold,omitempty"`
	// The threshold of average bandwidth usage that triggers automatic bandwidth upgrade. Unit: %. Valid values:
	//
	// - **50**
	//
	// - **60**
	//
	// - **70**
	//
	// - **80**
	//
	// - **90**
	//
	// - **95**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	BandwidthUsageUpperThreshold *int32 `json:"BandwidthUsageUpperThreshold,omitempty" xml:"BandwidthUsageUpperThreshold,omitempty"`
	// Specifies whether to enable automatic bandwidth downgrade:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	Downgrade *bool `json:"Downgrade,omitempty" xml:"Downgrade,omitempty"`
	// The observation window for automatic bandwidth upgrade. The parameter value consists of a number and a time unit suffix. The time suffix **m*	- indicates minutes. Valid values:
	//
	// - **1m**
	//
	// - **5m**
	//
	// - **10m**
	//
	// - **15m**
	//
	// - **30m**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	ObservationWindowSize *string `json:"ObservationWindowSize,omitempty" xml:"ObservationWindowSize,omitempty"`
	// Specifies whether to enable automatic bandwidth upgrade:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
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
	// Specifies whether to apply the configurations under the **Resource*	- configuration item for the resource elastic scaling feature:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// The threshold of average CPU utilization that triggers automatic resource scale-up. Unit: %. Valid values:
	//
	// - **70**
	//
	// - **80**
	//
	// - **90**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	CpuUsageUpperThreshold *int32 `json:"CpuUsageUpperThreshold,omitempty" xml:"CpuUsageUpperThreshold,omitempty"`
	// The observation window for scale-down. The parameter value consists of a number and a time unit suffix. The time suffix **m*	- indicates minutes. Valid values:
	//
	// - **1m**
	//
	// - **3m**
	//
	// - **5m**
	//
	// - **10m**
	//
	// - **20m**
	//
	// - **30m**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	DowngradeObservationWindowSize *string `json:"DowngradeObservationWindowSize,omitempty" xml:"DowngradeObservationWindowSize,omitempty"`
	// Specifies whether to enable resource elastic scaling:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The observation window for scale-up. The parameter value consists of a number and a time unit suffix. The time suffix **m*	- indicates minutes. Valid values:
	//
	// - **1m**
	//
	// - **3m**
	//
	// - **5m**
	//
	// - **10m**
	//
	// - **20m**
	//
	// - **30m**
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
	// Specifies whether to apply the configurations under the **Shard*	- configuration item for the shard elastic scaling feature:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// 	Notice: The shard elastic scaling feature is supported only when the target database instance is Tair (Redis® OSS-Compatible) cloud-native (formerly cloud disk) edition in cluster architecture on the China site.
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// Specifies whether to enable automatic shard reduction:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// > The automatic shard reduction feature is currently in canary release.
	//
	// example:
	//
	// true
	Downgrade *bool `json:"Downgrade,omitempty" xml:"Downgrade,omitempty"`
	// The observation window for automatic shard reduction. The parameter value consists of a number and a time unit suffix. The time suffix **h*	- indicates hours, and **d*	- indicates days. Valid values:
	//
	// - **1h**
	//
	// - **2h**
	//
	// - **3h**
	//
	// - **1d**
	//
	// - **7d**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1h
	DowngradeObservationWindowSize *string `json:"DowngradeObservationWindowSize,omitempty" xml:"DowngradeObservationWindowSize,omitempty"`
	// The maximum total number of shards for the instance. The value must be a positive integer. Valid values: 4 to 32.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 16
	MaxShards *int32 `json:"MaxShards,omitempty" xml:"MaxShards,omitempty"`
	// The threshold of average memory utilization that triggers automatic shard reduction. Unit: %. Valid values:
	//
	// - **10**
	//
	// - **20**
	//
	// - **30**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 30
	MemUsageLowerThreshold *int32 `json:"MemUsageLowerThreshold,omitempty" xml:"MemUsageLowerThreshold,omitempty"`
	// The threshold of average memory utilization that triggers automatic shard addition. Unit: %. Valid values:
	//
	// - **50**
	//
	// - **60**
	//
	// - **70**
	//
	// - **80**
	//
	// - **90**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	MemUsageUpperThreshold *int32 `json:"MemUsageUpperThreshold,omitempty" xml:"MemUsageUpperThreshold,omitempty"`
	// The minimum total number of shards for the instance. The value must be a positive integer. Valid values: 4 to 32.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 4
	MinShards *int32 `json:"MinShards,omitempty" xml:"MinShards,omitempty"`
	// Specifies whether to enable automatic shard addition:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	Upgrade *bool `json:"Upgrade,omitempty" xml:"Upgrade,omitempty"`
	// The observation window for automatic shard addition. The parameter value consists of a number and a time unit suffix. The time suffix **m*	- indicates minutes. Valid values:
	//
	// - **5m**
	//
	// - **10m**
	//
	// - **15m**
	//
	// - **30m**
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
	// Specifies whether to apply the configurations under the **Spec*	- configuration item for the specification elastic scaling feature:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// The cool-down period. The parameter value consists of a number and a time unit suffix. The supported database instances and valid values are as follows. The time suffix **m*	- indicates minutes, **h*	- indicates hours, and **d*	- indicates days:
	//
	// - PolarDB for MySQL Cluster Edition: **5m**, **10m**, **30m**, **1h**, **2h**, **3h**, **1d**, and **7d**.
	//
	// - ApsaraDB RDS for MySQL high availability series cloud disk edition: **5m**, **10m**, **30m**, **1h**, **2h**, **3h**, **1d**, and **7d**.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	CoolDownTime *string `json:"CoolDownTime,omitempty" xml:"CoolDownTime,omitempty"`
	// The threshold of average CPU utilization that triggers automatic specification scale-up. Unit: %. Valid values:
	//
	// - **50**
	//
	// - **60**
	//
	// - **70**
	//
	// - **80**
	//
	// - **90**
	//
	// 	Notice: You must set this parameter when the target database instance is PolarDB for MySQL Cluster Edition or ApsaraDB RDS for MySQL high availability series cloud disk edition.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	CpuUsageUpperThreshold *int32 `json:"CpuUsageUpperThreshold,omitempty" xml:"CpuUsageUpperThreshold,omitempty"`
	// Specifies whether automatic specification scale-down is enabled:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// 	Notice: You must set this parameter when the target database instance is PolarDB for MySQL Cluster Edition or ApsaraDB RDS for MySQL high availability series cloud disk edition.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// true
	Downgrade *bool `json:"Downgrade,omitempty" xml:"Downgrade,omitempty"`
	// The maximum number of read-only nodes for the instance.
	//
	// 	Notice: You must set this parameter when the target database instance is PolarDB for MySQL Cluster Edition.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 10
	MaxReadOnlyNodes *int32 `json:"MaxReadOnlyNodes,omitempty" xml:"MaxReadOnlyNodes,omitempty"`
	// The upper limit of the specification for automatic scale-up. You can only scale up to a higher specification within the same series of the database instance. For more details, see the specification documentation for each database instance:
	//
	// <props="china">
	//
	// - For PolarDB for MySQL Cluster Edition, see [Compute node specifications](https://help.aliyun.com/document_detail/102542.html).
	//
	// - For ApsaraDB RDS for MySQL high availability series cloud disk edition, see [Product specifications](https://help.aliyun.com/document_detail/276974.html).
	//
	// - For Tair (Redis® OSS-Compatible) cloud-native (formerly cloud disk) open-source edition, see [Instance specifications](https://help.aliyun.com/document_detail/144986.html).
	//
	//
	//
	// <props="intl">
	//
	// - For PolarDB for MySQL Cluster Edition, see [Compute node specifications](https://help.aliyun.com/document_detail/102542.html).
	//
	// - For ApsaraDB RDS for MySQL high availability series cloud disk edition, see [Product specifications](https://help.aliyun.com/document_detail/276974.html).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// polar.mysql.x8.12xlarge
	MaxSpec *string `json:"MaxSpec,omitempty" xml:"MaxSpec,omitempty"`
	// The threshold of average memory utilization that triggers automatic specification scale-up. Unit: %. Valid values:
	//
	// - **50**
	//
	// - **60**
	//
	// - **70**
	//
	// - **80**
	//
	// - **90**
	//
	// 	Notice: You must set this parameter when the target database instance is Tair (Redis® OSS-Compatible) cloud-native (formerly cloud disk) open-source edition on the China site.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	MemUsageUpperThreshold *int32 `json:"MemUsageUpperThreshold,omitempty" xml:"MemUsageUpperThreshold,omitempty"`
	// The observation window. The parameter value consists of a number and a time unit suffix. The supported database instances and valid values are as follows. The time suffix **m*	- indicates minutes, and **h*	- indicates hours:
	//
	// - PolarDB for MySQL Cluster Edition: **5m**, **10m**, **15m**, and **30m**.
	//
	// - ApsaraDB RDS for MySQL high availability series cloud disk edition: **5m**, **20m**, **30m**, **40m**, and **1h**.
	//
	// - Tair (Redis® OSS-Compatible) cloud-native (formerly cloud disk) open-source edition: **5m**, **10m**, **15m**, and **30m**.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 5m
	ObservationWindowSize *string `json:"ObservationWindowSize,omitempty" xml:"ObservationWindowSize,omitempty"`
	// Specifies whether to enable automatic specification scale-up:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
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
	// Specifies whether to apply the configurations under the **Storage*	- configuration item for the automatic storage expansion feature:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	Apply *bool `json:"Apply,omitempty" xml:"Apply,omitempty"`
	// The threshold of average storage space utilization that triggers automatic storage expansion. Unit: %. Valid values:
	//
	// - **50**
	//
	// - **60**
	//
	// - **70**
	//
	// - **80**
	//
	// - **90**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 70
	DiskUsageUpperThreshold *int32 `json:"DiskUsageUpperThreshold,omitempty" xml:"DiskUsageUpperThreshold,omitempty"`
	// The upper limit of the storage space. The value must be greater than or equal to the current total storage space of the instance. Valid values:
	//
	// - ESSD cloud disk upper limit: 32,000 GB.
	//
	// - Standard SSD upper limit: 6,000 GB.
	//
	// > Standard SSDs are retired. [Upgrade to ESSD cloud disks](https://help.aliyun.com/document_detail/314678.html).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 32000
	MaxStorage *int32 `json:"MaxStorage,omitempty" xml:"MaxStorage,omitempty"`
	// Specifies whether to enable automatic storage expansion:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
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
