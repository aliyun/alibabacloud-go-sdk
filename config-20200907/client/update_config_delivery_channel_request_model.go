// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateConfigDeliveryChannelRequest
	GetClientToken() *string
	SetCompliantSnapshot(v bool) *UpdateConfigDeliveryChannelRequest
	GetCompliantSnapshot() *bool
	SetConfigurationItemChangeNotification(v bool) *UpdateConfigDeliveryChannelRequest
	GetConfigurationItemChangeNotification() *bool
	SetConfigurationSnapshot(v bool) *UpdateConfigDeliveryChannelRequest
	GetConfigurationSnapshot() *bool
	SetDeliveryChannelCondition(v string) *UpdateConfigDeliveryChannelRequest
	GetDeliveryChannelCondition() *string
	SetDeliveryChannelId(v string) *UpdateConfigDeliveryChannelRequest
	GetDeliveryChannelId() *string
	SetDeliveryChannelName(v string) *UpdateConfigDeliveryChannelRequest
	GetDeliveryChannelName() *string
	SetDeliveryChannelTargetArn(v string) *UpdateConfigDeliveryChannelRequest
	GetDeliveryChannelTargetArn() *string
	SetDeliverySnapshotTime(v string) *UpdateConfigDeliveryChannelRequest
	GetDeliverySnapshotTime() *string
	SetDescription(v string) *UpdateConfigDeliveryChannelRequest
	GetDescription() *string
	SetNonCompliantNotification(v bool) *UpdateConfigDeliveryChannelRequest
	GetNonCompliantNotification() *bool
	SetOversizedDataOSSTargetArn(v string) *UpdateConfigDeliveryChannelRequest
	GetOversizedDataOSSTargetArn() *string
	SetStatus(v int64) *UpdateConfigDeliveryChannelRequest
	GetStatus() *int64
}

type UpdateConfigDeliveryChannelRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must ensure that the token is unique among different requests.
	//
	// The `token` can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to deliver scheduled compliant snapshots. Cloud Config delivers scheduled compliant snapshots at `04:00Z` and `16:00Z` to  Log Service every day. The time is displayed in UTC. Valid values:
	//
	// 	- true: Cloud Config delivers compliant snapshots.
	//
	// 	- false (default): Cloud Config does not deliver scheduled compliant snapshots.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	CompliantSnapshot *bool `json:"CompliantSnapshot,omitempty" xml:"CompliantSnapshot,omitempty"`
	// Specifies whether to deliver resource change logs. If you set this parameter to true, Cloud Config delivers resource change logs to OSS, Log Service, or MNS when the configurations of the resources change. Valid values:
	//
	// 	- true: Cloud Config delivers resource change logs.
	//
	// 	- false (default): Cloud Config does not deliver resource change logs.
	//
	// > This parameter is available for delivery channels of the OSS, SLS, and MNS types.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	ConfigurationItemChangeNotification *bool `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	// Specifies whether to deliver scheduled resource snapshots. Cloud Config delivers scheduled resource snapshots at `04:00Z` and `16:00Z` to OSS, MNS, or Log Service every day. The time is displayed in UTC. Valid values:
	//
	// 	- true: Cloud Config delivers scheduled resource snapshots.
	//
	// 	- false (default): Cloud Config does not deliver scheduled resource snapshots.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	ConfigurationSnapshot *bool `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	// The rule that you want to attach to the delivery channel. This parameter is available when you deliver data of all types to MNS or deliver snapshots to Log Service.
	//
	// 	- If the value of the DeliveryChannelType parameter is MNS, take note of the following settings of the lowest risk level and resource types of the events to which you subscribed:
	//
	//     	- The setting of the lowest risk level for the events to which you want to subscribe is in the following format: `{"filterType":"RuleRiskLevel","value":"1","multiple":false}`.
	//
	//         The `value` field indicates the lowest risk level of the events to which you want to subscribe. Valid values: 1, 2, and 3. The value 1 indicates the high risk level, the value 2 indicates the medium risk level, and the value 3 indicates the low risk level.
	//
	//     	- The setting of the resource types of the events to which you want to subscribe is in the following format: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//         The `values` field indicates the resource types of the events to which you want to subscribe. The value of the field is a JSON array. Examples:
	//
	// `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	//
	// 	- If you set the DeliveryChannelType parameter to SLS, the setting of the resource types of the snapshots to which you want to deliver is in the following format: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//     The `values` field specifies the resource types of the snapshots to which you want to deliver. The value of the field is a JSON array. Examples:
	//
	// `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	//
	// example:
	//
	// [{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]
	DeliveryChannelCondition *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	// The ID of the delivery channel.
	//
	// For more information about how to obtain the ID of a delivery channel, see [DescribeDeliveryChannels](https://help.aliyun.com/document_detail/429841.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cdc-8e45ff4e06a3a8****
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// testoss
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the delivery destination. Valid values:
	//
	// 	- `acs:oss:{RegionId}:{accountId}:{bucketName}` if your delivery destination is an OSS bucket. Example: `acs:oss:cn-shanghai:100931896542****:new-bucket`.
	//
	// 	- `acs:mns:{RegionId}:{accountId}:/topics/{topicName}` if your delivery destination is an MNS topic. Example: `acs:mns:cn-shanghai:100931896542****:/topics/topic1`.
	//
	// 	- `acs:log:{RegionId}:{accountId}:project/{projectName}/logstore/{logstoreName}` if your delivery destination is a Log Service Logstore. Example: `acs:log:cn-shanghai:100931896542****:project/project1/logstore/logstore1`.
	//
	// example:
	//
	// acs:oss:cn-shanghai:100931896542****:new-bucket
	DeliveryChannelTargetArn *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	// The time when you want Cloud Config to deliver scheduled resource snapshots every day.
	//
	// Format: `HH:mmZ`. This time is displayed in UTC.
	//
	// > When you enable the scheduled resource delivery feature, you can configure this parameter to specify a custom delivery time. If you do not configure this parameter, Cloud Config automatically delivers scheduled resource snapshots at `04:00Z` and `16:00Z` every day.
	//
	// example:
	//
	// 09:10Z
	DeliverySnapshotTime *string `json:"DeliverySnapshotTime,omitempty" xml:"DeliverySnapshotTime,omitempty"`
	// The description of the delivery channel.
	//
	// example:
	//
	// My OSS delivery.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to deliver resource non-compliance events. If you set this parameter to true, Cloud Config delivers resource non-compliance events to Log Service or MNS when resources are considered non-compliant. Valid values:
	//
	// 	- true: Cloud Config delivers resource non-compliance events.
	//
	// 	- false (default): Cloud Config does not deliver resource non-compliance events.
	//
	// > This parameter is available only for delivery channels of the SLS or MNS type.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	NonCompliantNotification *bool `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	// The ARN of the OSS bucket to which you want to transfer the delivery data when the size of the data exceeds the specified upper limit of the delivery channel. Format: `acs:oss:{RegionId}:{accountId}:{bucketName}`.
	//
	// If you do not configure this parameter, Cloud Config delivers only summary data.
	//
	// > This parameter is available only for delivery channels of the SLS or MNS type. The maximum storage size of delivery channels of the SLS type is 1 MB, and the maximum storage size of delivery channels of the MNS type is 64 KB.
	//
	// example:
	//
	// acs:oss:cn-shanghai:100931896542****:new-bucket
	OversizedDataOSSTargetArn *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	// Specifies whether to enable the delivery channel. Valid values:
	//
	// 	- 0: Cloud Config disables the delivery channel. Cloud Config retains the most recent delivery configuration and stops resource data delivery.
	//
	// 	- 1 (default): Cloud Config enables the delivery channel.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateConfigDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigDeliveryChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateConfigDeliveryChannelRequest) GetCompliantSnapshot() *bool {
	return s.CompliantSnapshot
}

func (s *UpdateConfigDeliveryChannelRequest) GetConfigurationItemChangeNotification() *bool {
	return s.ConfigurationItemChangeNotification
}

func (s *UpdateConfigDeliveryChannelRequest) GetConfigurationSnapshot() *bool {
	return s.ConfigurationSnapshot
}

func (s *UpdateConfigDeliveryChannelRequest) GetDeliveryChannelCondition() *string {
	return s.DeliveryChannelCondition
}

func (s *UpdateConfigDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *UpdateConfigDeliveryChannelRequest) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *UpdateConfigDeliveryChannelRequest) GetDeliveryChannelTargetArn() *string {
	return s.DeliveryChannelTargetArn
}

func (s *UpdateConfigDeliveryChannelRequest) GetDeliverySnapshotTime() *string {
	return s.DeliverySnapshotTime
}

func (s *UpdateConfigDeliveryChannelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConfigDeliveryChannelRequest) GetNonCompliantNotification() *bool {
	return s.NonCompliantNotification
}

func (s *UpdateConfigDeliveryChannelRequest) GetOversizedDataOSSTargetArn() *string {
	return s.OversizedDataOSSTargetArn
}

func (s *UpdateConfigDeliveryChannelRequest) GetStatus() *int64 {
	return s.Status
}

func (s *UpdateConfigDeliveryChannelRequest) SetClientToken(v string) *UpdateConfigDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetCompliantSnapshot(v bool) *UpdateConfigDeliveryChannelRequest {
	s.CompliantSnapshot = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetConfigurationItemChangeNotification(v bool) *UpdateConfigDeliveryChannelRequest {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetConfigurationSnapshot(v bool) *UpdateConfigDeliveryChannelRequest {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *UpdateConfigDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *UpdateConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDeliveryChannelName(v string) *UpdateConfigDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *UpdateConfigDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDeliverySnapshotTime(v string) *UpdateConfigDeliveryChannelRequest {
	s.DeliverySnapshotTime = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetDescription(v string) *UpdateConfigDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetNonCompliantNotification(v bool) *UpdateConfigDeliveryChannelRequest {
	s.NonCompliantNotification = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetOversizedDataOSSTargetArn(v string) *UpdateConfigDeliveryChannelRequest {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) SetStatus(v int64) *UpdateConfigDeliveryChannelRequest {
	s.Status = &v
	return s
}

func (s *UpdateConfigDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
