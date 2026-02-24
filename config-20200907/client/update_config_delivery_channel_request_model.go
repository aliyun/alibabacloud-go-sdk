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
	// A client token used to ensure the idempotence of the request. Use a client to generate the token, and make sure that the token is unique among different requests.
	//
	// The `ClientToken` parameter can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to deliver resource compliance snapshots. Cloud Config delivers resource compliance and non-compliance information to SLS. Valid values:
	//
	// - true: The resource compliance snapshots are delivered.
	//
	// - false: The resource compliance snapshots are not delivered.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	CompliantSnapshot *bool `json:"CompliantSnapshot,omitempty" xml:"CompliantSnapshot,omitempty"`
	// Specifies whether to deliver the resource configuration history. Cloud Config delivers the resource configuration history to OSS, SLS, or MNS when the configuration of a resource changes. Valid values:
	//
	// - true: The resource configuration history is delivered.
	//
	// - false (default): The resource configuration history is not delivered.
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
	// Specifies whether to deliver scheduled resource snapshots. Cloud Config delivers scheduled resource snapshots to OSS or SLS at `04:00Z` and `16:00Z` (UTC) every day. Valid values:
	//
	// - true: The scheduled resource snapshots are delivered.
	//
	// - false (default): The scheduled resource snapshots are not delivered.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	ConfigurationSnapshot *bool `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	// The rule that is attached to the delivery channel. This parameter is applicable to all deliveries to MNS and snapshot deliveries to SLS.
	//
	// - If you specify the minimum risk level of events and the resource types for an MNS subscription, use the following formats:
	//
	//   - The minimum risk level of the subscribed events: `{"filterType":"RuleRiskLevel","value":"1","multiple":false}`.
	//
	//     `value` specifies the risk level. Valid values: 1 (high risk), 2 (medium risk), and 3 (low risk).
	//
	//   - The resource types of the subscribed events: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//     `values` specifies the resource types of the subscribed events. The value is a JSON array.
	//
	//     Example:
	//
	//     `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	//
	// - If you specify the resource types of snapshots delivered to SLS, use the following format: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//   `values` specifies the resource types of the snapshots to deliver. The value is a JSON array.
	//
	//   Example:
	//
	//   `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	//
	// example:
	//
	// [{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]
	DeliveryChannelCondition *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	// The ID of the delivery channel.
	//
	// For more information about how to obtain the delivery channel ID, see [ListConfigDeliveryChannels](https://help.aliyun.com/document_detail/429841.html).
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
	// - If the delivery channel is Object Storage Service (OSS), the value is in the format of `acs:oss:{RegionId}:{accountId}:{bucketName}`. Example: `acs:oss:cn-shanghai:100931896542****:new-bucket`.
	//
	// - If the delivery channel is MNS, the value is in the format of `acs:mns:{RegionId}:{accountId}:/topics/{topicName}`. Example: `acs:mns:cn-shanghai:100931896542****:/topics/topic1`.
	//
	// - If the delivery channel is Simple Log Service (SLS), the value is in the format of `acs:log:{RegionId}:{accountId}:project/{projectName}/logstore/{logstoreName}`. Example: `acs:log:cn-shanghai:100931896542****:project/project1/logstore/logstore1`.
	//
	// example:
	//
	// acs:oss:cn-shanghai:100931896542****:new-bucket
	DeliveryChannelTargetArn *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	// The time of day when the scheduled resource snapshot is delivered.
	//
	// The value is in the `HH:mmZ` format. The time is in UTC.
	//
	// > If you enable scheduled delivery of resource snapshots, use this parameter to specify a delivery time. If you do not specify this parameter, Cloud Config delivers the scheduled resource snapshots at `04:00Z` and `16:00Z` by default.
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
	// Specifies whether to deliver resource non-compliance events. Cloud Config delivers resource non-compliance events to SLS or MNS when a resource is evaluated as non-compliant. Valid values:
	//
	// - true: The resource non-compliance events are delivered.
	//
	// - false (default): The resource non-compliance events are not delivered.
	//
	// > This parameter is available only for delivery channels of the SLS and MNS types.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	NonCompliantNotification *bool `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	// The ARN of the OSS bucket where data is delivered when the data size exceeds the limit of the delivery channel. The value is in the format of `acs:oss:{RegionId}:{accountId}:{bucketName}`.
	//
	// If you do not specify this parameter, Cloud Config delivers only the summary of the data.
	//
	// > This parameter is available only for delivery channels of the SLS and MNS types. The maximum size of data that can be delivered to SLS is 1 MB. The maximum size of data that can be delivered to MNS is 64 KB.
	//
	// example:
	//
	// acs:oss:cn-shanghai:100931896542****:new-bucket
	OversizedDataOSSTargetArn *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	// The status of the delivery channel. Valid values:
	//
	// - 0: The delivery channel is disabled. Cloud Config retains the most recent delivery configuration and stops delivering resource data.
	//
	// - 1 (default): The delivery channel is enabled.
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
