// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateConfigDeliveryChannelRequest
	GetClientToken() *string
	SetCompliantSnapshot(v bool) *CreateConfigDeliveryChannelRequest
	GetCompliantSnapshot() *bool
	SetConfigurationItemChangeNotification(v bool) *CreateConfigDeliveryChannelRequest
	GetConfigurationItemChangeNotification() *bool
	SetConfigurationSnapshot(v bool) *CreateConfigDeliveryChannelRequest
	GetConfigurationSnapshot() *bool
	SetDeliveryChannelCondition(v string) *CreateConfigDeliveryChannelRequest
	GetDeliveryChannelCondition() *string
	SetDeliveryChannelName(v string) *CreateConfigDeliveryChannelRequest
	GetDeliveryChannelName() *string
	SetDeliveryChannelTargetArn(v string) *CreateConfigDeliveryChannelRequest
	GetDeliveryChannelTargetArn() *string
	SetDeliveryChannelType(v string) *CreateConfigDeliveryChannelRequest
	GetDeliveryChannelType() *string
	SetDeliverySnapshotTime(v string) *CreateConfigDeliveryChannelRequest
	GetDeliverySnapshotTime() *string
	SetDescription(v string) *CreateConfigDeliveryChannelRequest
	GetDescription() *string
	SetNonCompliantNotification(v bool) *CreateConfigDeliveryChannelRequest
	GetNonCompliantNotification() *bool
	SetOversizedDataOSSTargetArn(v string) *CreateConfigDeliveryChannelRequest
	GetOversizedDataOSSTargetArn() *string
}

type CreateConfigDeliveryChannelRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must ensure that the token is unique among different requests.
	//
	// The `token` can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 1594295238-f9361358-5843-4294-8d30-b5183fac****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to deliver scheduled compliant snapshots. Cloud Config delivers scheduled compliant snapshots at `04:00Z` and `16:00Z` to Log Service every day. The time is displayed in UTC. Valid values:
	//
	// 	- true: Cloud Config delivers scheduled compliant snapshots.
	//
	// 	- false (default): Cloud Config does not deliver scheduled compliant snapshots.
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
	// example:
	//
	// false
	ConfigurationSnapshot *bool `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	// The rule that you want to attach to the delivery channel, used to specify subscription content filtering conditions.
	//
	// 	- If you specify the minimum risk level and resource types for compliance events, it should be as follows:
	//
	//     	- The setting of the lowest risk level for the events to which you want to subscribe is in the following format: `{"filterType":"RuleRiskLevel","value":"1","multiple":false}`.
	//
	//         The `value` field indicates the lowest risk level of the events to which you want to subscribe. Valid values: 1, 2, and 3. The value 1 indicates the high risk level, the value 2 indicates the medium risk level, and the value 3 indicates the low risk level.
	//
	//         The `multiple` field indicates whether multiple values are supported for this group of filters. Risk levels only support single-value filtering, so the multiple field for compliance event type content can only be set to `false`.
	//
	//     	- The setting of the resource types of the events to which you want to subscribe is in the following format: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//         The `values` field indicates the resource types of the events to which you want to subscribe.
	//
	//         The `multiple` field indicates whether multiple values are supported for this group of filters. Resource types support multi-value filtering; when selecting multiple resource types, the multiple field can be set to true.
	//
	// The value of the field is a JSON array. Examples:
	//
	//         `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	//
	//     	- You can also simultaneously specify both risk levels and resource types, such as:
	//
	//         `[{"filterType":"RuleRiskLevel","value":"2","multiple":false},{"filterType":"ResourceType","values":["ACS::CDN::Domain","ACS::ActionTrail::Trail"],"multiple":true}]`
	//
	// 	- If you specify the resource types for delivering configurations, the resource types are specified by: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//     The `values` field specifies the resource types of the snapshots to which you want to deliver. The value of the field is a JSON array. Examples:
	//
	//       `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
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
	// The Alibaba Cloud Resource Name (ARN) of the delivery destination. Valid values:
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
	// 	- false (default): Cloud Config does not deliver resource non-compliance events.
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
	// > This parameter is available only for delivery channels of the SLS or MNS type. The maximum storage size of delivery channels of the SLS type is 1 MB, and the maximum storage size of delivery channels of the MNS type is 64 KB.
	//
	// example:
	//
	// acs:oss:cn-shanghai:100931896542****:new-bucket
	OversizedDataOSSTargetArn *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
}

func (s CreateConfigDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigDeliveryChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateConfigDeliveryChannelRequest) GetCompliantSnapshot() *bool {
	return s.CompliantSnapshot
}

func (s *CreateConfigDeliveryChannelRequest) GetConfigurationItemChangeNotification() *bool {
	return s.ConfigurationItemChangeNotification
}

func (s *CreateConfigDeliveryChannelRequest) GetConfigurationSnapshot() *bool {
	return s.ConfigurationSnapshot
}

func (s *CreateConfigDeliveryChannelRequest) GetDeliveryChannelCondition() *string {
	return s.DeliveryChannelCondition
}

func (s *CreateConfigDeliveryChannelRequest) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *CreateConfigDeliveryChannelRequest) GetDeliveryChannelTargetArn() *string {
	return s.DeliveryChannelTargetArn
}

func (s *CreateConfigDeliveryChannelRequest) GetDeliveryChannelType() *string {
	return s.DeliveryChannelType
}

func (s *CreateConfigDeliveryChannelRequest) GetDeliverySnapshotTime() *string {
	return s.DeliverySnapshotTime
}

func (s *CreateConfigDeliveryChannelRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConfigDeliveryChannelRequest) GetNonCompliantNotification() *bool {
	return s.NonCompliantNotification
}

func (s *CreateConfigDeliveryChannelRequest) GetOversizedDataOSSTargetArn() *string {
	return s.OversizedDataOSSTargetArn
}

func (s *CreateConfigDeliveryChannelRequest) SetClientToken(v string) *CreateConfigDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetCompliantSnapshot(v bool) *CreateConfigDeliveryChannelRequest {
	s.CompliantSnapshot = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetConfigurationItemChangeNotification(v bool) *CreateConfigDeliveryChannelRequest {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetConfigurationSnapshot(v bool) *CreateConfigDeliveryChannelRequest {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *CreateConfigDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDeliveryChannelName(v string) *CreateConfigDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *CreateConfigDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDeliveryChannelType(v string) *CreateConfigDeliveryChannelRequest {
	s.DeliveryChannelType = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDeliverySnapshotTime(v string) *CreateConfigDeliveryChannelRequest {
	s.DeliverySnapshotTime = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetDescription(v string) *CreateConfigDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetNonCompliantNotification(v bool) *CreateConfigDeliveryChannelRequest {
	s.NonCompliantNotification = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) SetOversizedDataOSSTargetArn(v string) *CreateConfigDeliveryChannelRequest {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *CreateConfigDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
