// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAggregateConfigDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *CreateAggregateConfigDeliveryChannelRequest
	GetAggregatorId() *string
	SetClientToken(v string) *CreateAggregateConfigDeliveryChannelRequest
	GetClientToken() *string
	SetCompliantSnapshot(v bool) *CreateAggregateConfigDeliveryChannelRequest
	GetCompliantSnapshot() *bool
	SetConfigurationItemChangeNotification(v bool) *CreateAggregateConfigDeliveryChannelRequest
	GetConfigurationItemChangeNotification() *bool
	SetConfigurationSnapshot(v bool) *CreateAggregateConfigDeliveryChannelRequest
	GetConfigurationSnapshot() *bool
	SetDeliveryChannelCondition(v string) *CreateAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelCondition() *string
	SetDeliveryChannelName(v string) *CreateAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelName() *string
	SetDeliveryChannelTargetArn(v string) *CreateAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelTargetArn() *string
	SetDeliveryChannelType(v string) *CreateAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelType() *string
	SetDeliverySnapshotTime(v string) *CreateAggregateConfigDeliveryChannelRequest
	GetDeliverySnapshotTime() *string
	SetDescription(v string) *CreateAggregateConfigDeliveryChannelRequest
	GetDescription() *string
	SetNonCompliantNotification(v bool) *CreateAggregateConfigDeliveryChannelRequest
	GetNonCompliantNotification() *bool
	SetOversizedDataOSSTargetArn(v string) *CreateAggregateConfigDeliveryChannelRequest
	GetOversizedDataOSSTargetArn() *string
}

type CreateAggregateConfigDeliveryChannelRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The `token` can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html)
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to deliver scheduled compliant snapshots. Cloud Config delivers scheduled compliant snapshots to Log Service at `04:00Z` and `16:00Z` every day. The time is displayed in UTC. Valid values:
	//
	// 	- true: Cloud Config delivers scheduled compliant snapshots.
	//
	// 	- false: Cloud Config does not deliver scheduled compliant snapshots. This is the default value.
	//
	// example:
	//
	// false
	CompliantSnapshot *bool `json:"CompliantSnapshot,omitempty" xml:"CompliantSnapshot,omitempty"`
	// Specifies whether to deliver resource change logs. If you set this parameter to true, Cloud Config delivers resource change logs to OSS, Log Service, or MNS when the configurations of the resources change. Valid values:
	//
	// 	- true: Cloud Config delivers resource change logs.
	//
	// 	- false: Cloud Config does not deliver resource change logs. This is the default value.
	//
	// > This parameter is available for delivery channels of the OSS, SLS, and MNS types.
	//
	// example:
	//
	// false
	ConfigurationItemChangeNotification *bool `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	// Specifies whether to deliver scheduled resource snapshots. Cloud Config delivers scheduled resource snapshots to OSS, MNS, or Log Service at `04:00Z` and `16:00Z` every day. The time is displayed in UTC. Valid values:
	//
	// 	- true: Cloud Config delivers scheduled resource snapshots.
	//
	// 	- false: Cloud Config does not deliver scheduled resource snapshots. This is the default value.
	//
	// example:
	//
	// false
	ConfigurationSnapshot *bool `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	// The rule that you want to attach to the delivery channel. This parameter is available when you deliver data of all types to MNS or deliver snapshots to Log Service.
	//
	// 	- If you set the DeliveryChannelType parameter to MNS, take note of the following settings of the lowest risk level and the resource types of the events to which you want to subscribe:
	//
	//     	- The lowest risk level of the events to which you want to subscribe is in the following format: `{"filterType":"RuleRiskLevel","value":"1","multiple":false}`.
	//
	//         The `value` field indicates the risk level of the events to which you want to subscribe. Valid values: 1, 2, and 3. The value 1 indicates the high risk level, the value 2 indicates the medium risk level, and the value 3 indicates the low risk level.
	//
	//     	- The setting of the resource types of the events to which you want to subscribe is in the following format: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//         The `values` field indicates the resource types of the events to which you want to subscribe. The value of the field is a JSON array. Example:
	//
	// `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	//
	// 	- If you set the DeliveryChannelType parameter to SLS, the setting of the resource types of the snapshots to which you want to deliver is in the following format: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//     The `values` field specifies the resource types of the snapshots to which you want to deliver. The value of the field is a JSON array. Example:
	//
	// `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	//
	// example:
	//
	// [{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]
	DeliveryChannelCondition *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	// The name of the delivery channel.
	//
	// > If you do not configure this parameter, this parameter is left empty.
	//
	// example:
	//
	// testoss
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The ARN of the delivery destination. Valid values:
	//
	// 	- `acs:oss:{RegionId}:{accountId}:{bucketName}` if your delivery destination is an OSS bucket. Example: `acs:oss:cn-shanghai:100931896542****:new-bucket`.
	//
	// 	- `acs:mns:{RegionId}:{accountId}:/topics/{topicName}` if your delivery destination is an MNS topic. Example: `acs:mns:cn-shanghai:100931896542****:/topics/topic1`.
	//
	// 	- `acs:log:{RegionId}:{accountId}:project/{projectName}/logstore/{logstoreName}` if your delivery destination is a Log Service Logstore. Example: `acs:log:cn-shanghai:100931896542****:project/project1/logstore/logstore1`.
	//
	// This parameter is required.
	//
	// example:
	//
	// acs:oss:cn-shanghai:100931896542****:new-bucket
	DeliveryChannelTargetArn *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	// The type of the delivery channel. Valid values:
	//
	// 	- OSS: Object Storage Service (OSS)
	//
	// 	- MNS: Message Service (MNS)
	//
	// 	- SLS: Log Service
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS
	DeliveryChannelType *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
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
	// 	- false: Cloud Config does not deliver resource non-compliance events. This is the default value.
	//
	// > This parameter is available only for delivery channels of the SLS or MNS type.
	//
	// example:
	//
	// false
	NonCompliantNotification *bool `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	// The ARN of the OSS bucket to which you want to transfer the delivery data when the size of the data exceeds the specified upper limit of the delivery channel. Format: `acs:oss:{RegionId}:{accountId}:{bucketName}`.
	//
	// If you do not configure this parameter, Cloud Config delivers only summary data.
	//
	// > This parameter is available only for delivery channels of the SLS or MNS type. The upper limit on the storage size of delivery channels of the SLS type is 1 MB, and the maximum storage size of delivery channels of the MNS type is 64 KB.
	//
	// example:
	//
	// acs:oss:cn-shanghai:100931896542****:new-bucket
	OversizedDataOSSTargetArn *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
}

func (s CreateAggregateConfigDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAggregateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetCompliantSnapshot() *bool {
	return s.CompliantSnapshot
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetConfigurationItemChangeNotification() *bool {
	return s.ConfigurationItemChangeNotification
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetConfigurationSnapshot() *bool {
	return s.ConfigurationSnapshot
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetDeliveryChannelCondition() *string {
	return s.DeliveryChannelCondition
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetDeliveryChannelTargetArn() *string {
	return s.DeliveryChannelTargetArn
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetDeliveryChannelType() *string {
	return s.DeliveryChannelType
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetDeliverySnapshotTime() *string {
	return s.DeliverySnapshotTime
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetNonCompliantNotification() *bool {
	return s.NonCompliantNotification
}

func (s *CreateAggregateConfigDeliveryChannelRequest) GetOversizedDataOSSTargetArn() *string {
	return s.OversizedDataOSSTargetArn
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetAggregatorId(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.AggregatorId = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetClientToken(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetCompliantSnapshot(v bool) *CreateAggregateConfigDeliveryChannelRequest {
	s.CompliantSnapshot = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetConfigurationItemChangeNotification(v bool) *CreateAggregateConfigDeliveryChannelRequest {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetConfigurationSnapshot(v bool) *CreateAggregateConfigDeliveryChannelRequest {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelName(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelType(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelType = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDeliverySnapshotTime(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.DeliverySnapshotTime = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetDescription(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetNonCompliantNotification(v bool) *CreateAggregateConfigDeliveryChannelRequest {
	s.NonCompliantNotification = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) SetOversizedDataOSSTargetArn(v string) *CreateAggregateConfigDeliveryChannelRequest {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *CreateAggregateConfigDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
