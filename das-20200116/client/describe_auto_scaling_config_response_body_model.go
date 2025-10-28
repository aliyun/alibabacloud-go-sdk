// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoScalingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAutoScalingConfigResponseBody
	GetCode() *string
	SetData(v *DescribeAutoScalingConfigResponseBodyData) *DescribeAutoScalingConfigResponseBody
	GetData() *DescribeAutoScalingConfigResponseBodyData
	SetMessage(v string) *DescribeAutoScalingConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAutoScalingConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeAutoScalingConfigResponseBody
	GetSuccess() *string
}

type DescribeAutoScalingConfigResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configurations of the auto scaling feature for instances.
	Data *DescribeAutoScalingConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutoScalingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAutoScalingConfigResponseBody) GetData() *DescribeAutoScalingConfigResponseBodyData {
	return s.Data
}

func (s *DescribeAutoScalingConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAutoScalingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoScalingConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeAutoScalingConfigResponseBody) SetCode(v string) *DescribeAutoScalingConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBody) SetData(v *DescribeAutoScalingConfigResponseBodyData) *DescribeAutoScalingConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutoScalingConfigResponseBody) SetMessage(v string) *DescribeAutoScalingConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBody) SetRequestId(v string) *DescribeAutoScalingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBody) SetSuccess(v string) *DescribeAutoScalingConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoScalingConfigResponseBodyData struct {
	// The configurations of the automatic bandwidth adjustment feature.
	Bandwidth *DescribeAutoScalingConfigResponseBodyDataBandwidth `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty" type:"Struct"`
	// The configurations of the auto scaling feature for local resources.
	Resource *DescribeAutoScalingConfigResponseBodyDataResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The configurations of the auto scaling feature for shards.
	Shard *DescribeAutoScalingConfigResponseBodyDataShard `json:"Shard,omitempty" xml:"Shard,omitempty" type:"Struct"`
	// The configurations of the auto scaling feature for specifications.
	Spec *DescribeAutoScalingConfigResponseBodyDataSpec `json:"Spec,omitempty" xml:"Spec,omitempty" type:"Struct"`
	// The configurations of the automatic storage expansion feature.
	Storage *DescribeAutoScalingConfigResponseBodyDataStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
}

func (s DescribeAutoScalingConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingConfigResponseBodyData) GetBandwidth() *DescribeAutoScalingConfigResponseBodyDataBandwidth {
	return s.Bandwidth
}

func (s *DescribeAutoScalingConfigResponseBodyData) GetResource() *DescribeAutoScalingConfigResponseBodyDataResource {
	return s.Resource
}

func (s *DescribeAutoScalingConfigResponseBodyData) GetShard() *DescribeAutoScalingConfigResponseBodyDataShard {
	return s.Shard
}

func (s *DescribeAutoScalingConfigResponseBodyData) GetSpec() *DescribeAutoScalingConfigResponseBodyDataSpec {
	return s.Spec
}

func (s *DescribeAutoScalingConfigResponseBodyData) GetStorage() *DescribeAutoScalingConfigResponseBodyDataStorage {
	return s.Storage
}

func (s *DescribeAutoScalingConfigResponseBodyData) SetBandwidth(v *DescribeAutoScalingConfigResponseBodyDataBandwidth) *DescribeAutoScalingConfigResponseBodyData {
	s.Bandwidth = v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyData) SetResource(v *DescribeAutoScalingConfigResponseBodyDataResource) *DescribeAutoScalingConfigResponseBodyData {
	s.Resource = v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyData) SetShard(v *DescribeAutoScalingConfigResponseBodyDataShard) *DescribeAutoScalingConfigResponseBodyData {
	s.Shard = v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyData) SetSpec(v *DescribeAutoScalingConfigResponseBodyDataSpec) *DescribeAutoScalingConfigResponseBodyData {
	s.Spec = v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyData) SetStorage(v *DescribeAutoScalingConfigResponseBodyDataStorage) *DescribeAutoScalingConfigResponseBodyData {
	s.Storage = v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyData) Validate() error {
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

type DescribeAutoScalingConfigResponseBodyDataBandwidth struct {
	// The average bandwidth usage threshold that triggers automatic bandwidth downgrade. Unit: %.
	//
	// example:
	//
	// 30
	BandwidthUsageLowerThreshold *int32 `json:"BandwidthUsageLowerThreshold,omitempty" xml:"BandwidthUsageLowerThreshold,omitempty"`
	// The average bandwidth usage threshold that triggers automatic bandwidth adjustment. Unit: %.
	//
	// example:
	//
	// 70
	BandwidthUsageUpperThreshold *int32 `json:"BandwidthUsageUpperThreshold,omitempty" xml:"BandwidthUsageUpperThreshold,omitempty"`
	// Indicates whether the automatic bandwidth downgrade feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Downgrade *bool `json:"Downgrade,omitempty" xml:"Downgrade,omitempty"`
	// The observation window of the automatic bandwidth adjustment feature. The return value consists of a numeric value and a time unit suffix. Valid values of the time unit suffix:
	//
	// 	- **s**: seconds.
	//
	// 	- **m**: minutes.
	//
	// 	- **h**: hours.
	//
	// 	- **d**: days.
	//
	// >  A value of **5m*	- indicates 5 minutes.
	//
	// example:
	//
	// 5m
	ObservationWindowSize *string `json:"ObservationWindowSize,omitempty" xml:"ObservationWindowSize,omitempty"`
	// Indicates whether the automatic bandwidth adjustment feature is enabled. Valid values:
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

func (s DescribeAutoScalingConfigResponseBodyDataBandwidth) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingConfigResponseBodyDataBandwidth) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) GetBandwidthUsageLowerThreshold() *int32 {
	return s.BandwidthUsageLowerThreshold
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) GetBandwidthUsageUpperThreshold() *int32 {
	return s.BandwidthUsageUpperThreshold
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) GetDowngrade() *bool {
	return s.Downgrade
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) GetObservationWindowSize() *string {
	return s.ObservationWindowSize
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) SetBandwidthUsageLowerThreshold(v int32) *DescribeAutoScalingConfigResponseBodyDataBandwidth {
	s.BandwidthUsageLowerThreshold = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) SetBandwidthUsageUpperThreshold(v int32) *DescribeAutoScalingConfigResponseBodyDataBandwidth {
	s.BandwidthUsageUpperThreshold = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) SetDowngrade(v bool) *DescribeAutoScalingConfigResponseBodyDataBandwidth {
	s.Downgrade = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) SetObservationWindowSize(v string) *DescribeAutoScalingConfigResponseBodyDataBandwidth {
	s.ObservationWindowSize = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) SetUpgrade(v bool) *DescribeAutoScalingConfigResponseBodyDataBandwidth {
	s.Upgrade = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataBandwidth) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoScalingConfigResponseBodyDataResource struct {
	// The scale-out step size of CPU.
	//
	// example:
	//
	// 2
	CpuStep *int32 `json:"CpuStep,omitempty" xml:"CpuStep,omitempty"`
	// The average CPU utilization threshold that triggers automatic scale-out of local resources. Unit: %.
	//
	// example:
	//
	// 70
	CpuUsageUpperThreshold *int32 `json:"CpuUsageUpperThreshold,omitempty" xml:"CpuUsageUpperThreshold,omitempty"`
	// The observation window of the automatic scale-in feature for local resources. The return value consists of a numeric value and a time unit suffix. Valid values of the time unit suffix:
	//
	// 	- **s**: seconds.
	//
	// 	- **m**: minutes.
	//
	// 	- **h**: hours.
	//
	// 	- **d**: days.
	//
	// >  A value of **5m*	- indicates 5 minutes.
	//
	// example:
	//
	// 5m
	DowngradeObservationWindowSize *string `json:"DowngradeObservationWindowSize,omitempty" xml:"DowngradeObservationWindowSize,omitempty"`
	// Indicates whether the auto scaling feature is enabled for local resources. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The observation window of the automatic scale-out feature for local resources. The return value consists of a numeric value and a time unit suffix. Valid values of the time unit suffix:
	//
	// 	- **s**: seconds.
	//
	// 	- **m**: minutes.
	//
	// 	- **h**: hours.
	//
	// 	- **d**: days.
	//
	// >  A value of **5m*	- indicates 5 minutes.
	//
	// example:
	//
	// 5m
	UpgradeObservationWindowSize *string `json:"UpgradeObservationWindowSize,omitempty" xml:"UpgradeObservationWindowSize,omitempty"`
}

func (s DescribeAutoScalingConfigResponseBodyDataResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingConfigResponseBodyDataResource) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) GetCpuStep() *int32 {
	return s.CpuStep
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) GetCpuUsageUpperThreshold() *int32 {
	return s.CpuUsageUpperThreshold
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) GetDowngradeObservationWindowSize() *string {
	return s.DowngradeObservationWindowSize
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) GetEnable() *bool {
	return s.Enable
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) GetUpgradeObservationWindowSize() *string {
	return s.UpgradeObservationWindowSize
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) SetCpuStep(v int32) *DescribeAutoScalingConfigResponseBodyDataResource {
	s.CpuStep = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) SetCpuUsageUpperThreshold(v int32) *DescribeAutoScalingConfigResponseBodyDataResource {
	s.CpuUsageUpperThreshold = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) SetDowngradeObservationWindowSize(v string) *DescribeAutoScalingConfigResponseBodyDataResource {
	s.DowngradeObservationWindowSize = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) SetEnable(v bool) *DescribeAutoScalingConfigResponseBodyDataResource {
	s.Enable = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) SetUpgradeObservationWindowSize(v string) *DescribeAutoScalingConfigResponseBodyDataResource {
	s.UpgradeObservationWindowSize = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataResource) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoScalingConfigResponseBodyDataShard struct {
	// Indicates whether the feature of automatically removing shards is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Downgrade *bool `json:"Downgrade,omitempty" xml:"Downgrade,omitempty"`
	// The observation window of the feature of automatically removing shards. The return value consists of a numeric value and a time unit suffix. Valid values of the time unit suffix:
	//
	// 	- **s**: seconds.
	//
	// 	- **m**: minutes.
	//
	// 	- **h**: hours.
	//
	// 	- **d**: days.
	//
	// >  A value of **1d*	- indicates one day.
	//
	// example:
	//
	// 1d
	DowngradeObservationWindowSize *string `json:"DowngradeObservationWindowSize,omitempty" xml:"DowngradeObservationWindowSize,omitempty"`
	// The maximum number of shards in the instance.
	//
	// example:
	//
	// 16
	MaxShards *int32 `json:"MaxShards,omitempty" xml:"MaxShards,omitempty"`
	// The average memory usage threshold that triggers automatic removal of shards. Unit: %.
	//
	// example:
	//
	// 30
	MemUsageLowerThreshold *int32 `json:"MemUsageLowerThreshold,omitempty" xml:"MemUsageLowerThreshold,omitempty"`
	// The average memory usage threshold that triggers automatic adding of shards. Unit: %.
	//
	// example:
	//
	// 70
	MemUsageUpperThreshold *int32 `json:"MemUsageUpperThreshold,omitempty" xml:"MemUsageUpperThreshold,omitempty"`
	// The minimum number of shards in the instance.
	//
	// example:
	//
	// 4
	MinShards *int32 `json:"MinShards,omitempty" xml:"MinShards,omitempty"`
	// Indicates whether the feature of automatically adding shards is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Upgrade *bool `json:"Upgrade,omitempty" xml:"Upgrade,omitempty"`
	// The observation window of the feature of automatically adding shards. The return value consists of a numeric value and a time unit suffix. Valid values of the time unit suffix:
	//
	// 	- **s**: seconds.
	//
	// 	- **m**: minutes.
	//
	// 	- **h**: hours.
	//
	// 	- **d**: days.
	//
	// >  A value of **5m*	- indicates 5 minutes.
	//
	// example:
	//
	// 5m
	UpgradeObservationWindowSize *string `json:"UpgradeObservationWindowSize,omitempty" xml:"UpgradeObservationWindowSize,omitempty"`
}

func (s DescribeAutoScalingConfigResponseBodyDataShard) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingConfigResponseBodyDataShard) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) GetDowngrade() *bool {
	return s.Downgrade
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) GetDowngradeObservationWindowSize() *string {
	return s.DowngradeObservationWindowSize
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) GetMaxShards() *int32 {
	return s.MaxShards
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) GetMemUsageLowerThreshold() *int32 {
	return s.MemUsageLowerThreshold
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) GetMemUsageUpperThreshold() *int32 {
	return s.MemUsageUpperThreshold
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) GetMinShards() *int32 {
	return s.MinShards
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) GetUpgradeObservationWindowSize() *string {
	return s.UpgradeObservationWindowSize
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) SetDowngrade(v bool) *DescribeAutoScalingConfigResponseBodyDataShard {
	s.Downgrade = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) SetDowngradeObservationWindowSize(v string) *DescribeAutoScalingConfigResponseBodyDataShard {
	s.DowngradeObservationWindowSize = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) SetMaxShards(v int32) *DescribeAutoScalingConfigResponseBodyDataShard {
	s.MaxShards = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) SetMemUsageLowerThreshold(v int32) *DescribeAutoScalingConfigResponseBodyDataShard {
	s.MemUsageLowerThreshold = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) SetMemUsageUpperThreshold(v int32) *DescribeAutoScalingConfigResponseBodyDataShard {
	s.MemUsageUpperThreshold = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) SetMinShards(v int32) *DescribeAutoScalingConfigResponseBodyDataShard {
	s.MinShards = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) SetUpgrade(v bool) *DescribeAutoScalingConfigResponseBodyDataShard {
	s.Upgrade = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) SetUpgradeObservationWindowSize(v string) *DescribeAutoScalingConfigResponseBodyDataShard {
	s.UpgradeObservationWindowSize = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataShard) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoScalingConfigResponseBodyDataSpec struct {
	// The quiescent period. The return value consists of a numeric value and a time unit suffix. Valid values of the time unit suffix:
	//
	// 	- **s**: seconds.
	//
	// 	- **m**: minutes.
	//
	// 	- **h**: hours.
	//
	// 	- **d**: days.
	//
	// >  A value of **5m*	- indicates 5 minutes.
	//
	// example:
	//
	// 5m
	CoolDownTime *string `json:"CoolDownTime,omitempty" xml:"CoolDownTime,omitempty"`
	// The average CPU utilization threshold that triggers automatic specification scale-up. Unit: %.
	//
	// example:
	//
	// 70
	CpuUsageUpperThreshold *int32 `json:"CpuUsageUpperThreshold,omitempty" xml:"CpuUsageUpperThreshold,omitempty"`
	// Indicates whether the automatic specification scale-down feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Downgrade *bool `json:"Downgrade,omitempty" xml:"Downgrade,omitempty"`
	// The maximum number of read-only nodes of the instance.
	//
	// example:
	//
	// 10
	MaxReadOnlyNodes *int32 `json:"MaxReadOnlyNodes,omitempty" xml:"MaxReadOnlyNodes,omitempty"`
	// The maximum specifications to which the cluster can be scaled up. For more information about the specifications of each type of supported database instances, see the following topics:
	//
	// 	- PolarDB for MySQL Cluster Edition instances: [Compute node specifications of PolarDB for MySQL Enterprise Edition](https://help.aliyun.com/document_detail/102542.html)
	//
	// 	- ApsaraDB RDS for MySQL High-availability Edition instances that use standard SSDs or ESSDs: [Specifications](https://help.aliyun.com/document_detail/276974.html)
	//
	// example:
	//
	// polar.mysql.x8.12xlarge
	MaxSpec *string `json:"MaxSpec,omitempty" xml:"MaxSpec,omitempty"`
	// The average memory usage threshold that triggers automatic specification scale-up. Unit: %.
	//
	// example:
	//
	// 70
	MemUsageUpperThreshold *int32 `json:"MemUsageUpperThreshold,omitempty" xml:"MemUsageUpperThreshold,omitempty"`
	// The observation window. The return value consists of a numeric value and a time unit suffix. Valid values of the time unit suffix:
	//
	// 	- **s**: seconds.
	//
	// 	- **m**: minutes.
	//
	// 	- **h**: hours.
	//
	// 	- **d**: days.
	//
	// >  A value of **5m*	- indicates 5 minutes.
	//
	// example:
	//
	// 5m
	ObservationWindowSize *string `json:"ObservationWindowSize,omitempty" xml:"ObservationWindowSize,omitempty"`
	// Indicates whether the automatic specification scale-up feature is enabled. Valid values:
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

func (s DescribeAutoScalingConfigResponseBodyDataSpec) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingConfigResponseBodyDataSpec) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) GetCoolDownTime() *string {
	return s.CoolDownTime
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) GetCpuUsageUpperThreshold() *int32 {
	return s.CpuUsageUpperThreshold
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) GetDowngrade() *bool {
	return s.Downgrade
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) GetMaxReadOnlyNodes() *int32 {
	return s.MaxReadOnlyNodes
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) GetMaxSpec() *string {
	return s.MaxSpec
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) GetMemUsageUpperThreshold() *int32 {
	return s.MemUsageUpperThreshold
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) GetObservationWindowSize() *string {
	return s.ObservationWindowSize
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) SetCoolDownTime(v string) *DescribeAutoScalingConfigResponseBodyDataSpec {
	s.CoolDownTime = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) SetCpuUsageUpperThreshold(v int32) *DescribeAutoScalingConfigResponseBodyDataSpec {
	s.CpuUsageUpperThreshold = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) SetDowngrade(v bool) *DescribeAutoScalingConfigResponseBodyDataSpec {
	s.Downgrade = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) SetMaxReadOnlyNodes(v int32) *DescribeAutoScalingConfigResponseBodyDataSpec {
	s.MaxReadOnlyNodes = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) SetMaxSpec(v string) *DescribeAutoScalingConfigResponseBodyDataSpec {
	s.MaxSpec = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) SetMemUsageUpperThreshold(v int32) *DescribeAutoScalingConfigResponseBodyDataSpec {
	s.MemUsageUpperThreshold = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) SetObservationWindowSize(v string) *DescribeAutoScalingConfigResponseBodyDataSpec {
	s.ObservationWindowSize = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) SetUpgrade(v bool) *DescribeAutoScalingConfigResponseBodyDataSpec {
	s.Upgrade = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataSpec) Validate() error {
	return dara.Validate(s)
}

type DescribeAutoScalingConfigResponseBodyDataStorage struct {
	// The average storage usage threshold that triggers automatic storage expansion. Unit: %.
	//
	// example:
	//
	// 70
	DiskUsageUpperThreshold *int32 `json:"DiskUsageUpperThreshold,omitempty" xml:"DiskUsageUpperThreshold,omitempty"`
	// The maximum storage size. Unit: GB.
	//
	// example:
	//
	// 32000
	MaxStorage *int32 `json:"MaxStorage,omitempty" xml:"MaxStorage,omitempty"`
	// Indicates whether the automatic storage expansion feature is enabled. Valid values:
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

func (s DescribeAutoScalingConfigResponseBodyDataStorage) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoScalingConfigResponseBodyDataStorage) GoString() string {
	return s.String()
}

func (s *DescribeAutoScalingConfigResponseBodyDataStorage) GetDiskUsageUpperThreshold() *int32 {
	return s.DiskUsageUpperThreshold
}

func (s *DescribeAutoScalingConfigResponseBodyDataStorage) GetMaxStorage() *int32 {
	return s.MaxStorage
}

func (s *DescribeAutoScalingConfigResponseBodyDataStorage) GetUpgrade() *bool {
	return s.Upgrade
}

func (s *DescribeAutoScalingConfigResponseBodyDataStorage) SetDiskUsageUpperThreshold(v int32) *DescribeAutoScalingConfigResponseBodyDataStorage {
	s.DiskUsageUpperThreshold = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataStorage) SetMaxStorage(v int32) *DescribeAutoScalingConfigResponseBodyDataStorage {
	s.MaxStorage = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataStorage) SetUpgrade(v bool) *DescribeAutoScalingConfigResponseBodyDataStorage {
	s.Upgrade = &v
	return s
}

func (s *DescribeAutoScalingConfigResponseBodyDataStorage) Validate() error {
	return dara.Validate(s)
}
