// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannel(v *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) *GetAggregateConfigDeliveryChannelResponseBody
	GetDeliveryChannel() *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel
	SetRequestId(v string) *GetAggregateConfigDeliveryChannelResponseBody
	GetRequestId() *string
}

type GetAggregateConfigDeliveryChannelResponseBody struct {
	// The information about a delivery channel.
	DeliveryChannel *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel `json:"DeliveryChannel,omitempty" xml:"DeliveryChannel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DC300244-FCE3-5061-8214-C27ECB668487
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAggregateConfigDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigDeliveryChannelResponseBody) GetDeliveryChannel() *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	return s.DeliveryChannel
}

func (s *GetAggregateConfigDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAggregateConfigDeliveryChannelResponseBody) SetDeliveryChannel(v *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) *GetAggregateConfigDeliveryChannelResponseBody {
	s.DeliveryChannel = v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBody) SetRequestId(v string) *GetAggregateConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBody) Validate() error {
	if s.DeliveryChannel != nil {
		if err := s.DeliveryChannel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel struct {
	// The ID of the member in the account group.
	//
	// example:
	//
	// 120886317861****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The ID of the account group.
	//
	// example:
	//
	// ca-a4e5626622af0079****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// Indicates whether the specified destination receives scheduled compliant snapshots. Cloud Config delivers scheduled compliant snapshots at `04:00Z` and `16:00Z` to Log Service every day. The time is displayed in UTC. Valid values:
	//
	// 	- true: The specified destination receives scheduled compliant snapshots.
	//
	// 	- false: The specified destination does not receive scheduled compliant snapshots.
	//
	// example:
	//
	// false
	CompliantSnapshot *bool `json:"CompliantSnapshot,omitempty" xml:"CompliantSnapshot,omitempty"`
	// Indicates whether the specified destination receives resource change logs. If the value of this parameter is true, Cloud Config delivers the resource change logs to OSS, Log Service, or MNS when the configurations of the resources change. Valid values:
	//
	// 	- true: The specified destination receives resource change logs.
	//
	// 	- false: The specified destination does not receive resource change logs.
	//
	// example:
	//
	// true
	ConfigurationItemChangeNotification *bool `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	// Indicates whether the specified destination receives scheduled resource snapshots. Cloud Config delivers scheduled resource snapshots at `04:00Z` and `16:00Z` to OSS, MNS, or Log Service every day. The time is displayed in UTC. Valid values:
	//
	// 	- true: The specified destination receives scheduled resource snapshots.
	//
	// 	- false: The specified destination does not receive scheduled resource snapshots.
	//
	// example:
	//
	// false
	ConfigurationSnapshot *bool `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	// The ARN of the role that is assigned to the delivery channel.
	//
	// example:
	//
	// acs:ram::120886317861****:role/aliyunserviceroleforconfig
	DeliveryChannelAssumeRoleArn *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	// The rule that is attached to the delivery channel. This parameter is available when you deliver data of all types to MNS or deliver snapshots to Log Service.
	//
	// 	- If the value of the DeliveryChannelType parameter is MNS, take note of the following settings of the lowest risk level and resource types of the events to which you subscribed:
	//
	//     	- The lowest risk level of the events to which you want to subscribe is in the following format: `{"filterType":"RuleRiskLevel","value":"1","multiple":false}`.
	//
	//         The `value` field indicates the risk level of the events to which you want to subscribe. Valid values: 1, 2, and 3. The value 1 indicates the high risk level, the value 2 indicates the medium risk level, and the value 3 indicates the low risk level.
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
	// example:
	//
	// cdc-d9106457e0d900b1****
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// myDeliveryChannel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The ARN of the delivery destination.
	//
	// 	- If the value of the DeliveryChannelType parameter is OSS, the value of this parameter is the ARN of the destination OSS bucket.
	//
	// 	- If the value of the DeliveryChannelType parameter is MNS, the value of this parameter is the ARN of the destination MNS topic.
	//
	// 	- If the value of the DeliveryChannelType parameter is SLS, the value of this parameter is the ARN of the destination Log Service Logstore.
	//
	// example:
	//
	// acs:oss:cn-shanghai:120886317861****:new-bucket
	DeliveryChannelTargetArn *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	// The type of the delivery channel. Valid values:
	//
	// 	- OSS: Object Storage Service (OSS)
	//
	// 	- MNS: Message Service (MNS)
	//
	// 	- SLS: Log Service
	//
	// example:
	//
	// OSS
	DeliveryChannelType *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	// The time when Cloud Config delivers scheduled resources snapshots every day.
	//
	// Format: `HH:mmZ`. This time is displayed in UTC.
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
	// Indicates whether the specified destination receives resource non-compliance events. If the value of this parameter is true, Cloud Config delivers resource non-compliance events to Log Service or MNS when resources are evaluated as non-compliant. Valid values:
	//
	// 	- true: The specified destination receives resource non-compliance events.
	//
	// 	- false: The specified destination does not receive resource non-compliance events.
	//
	// example:
	//
	// false
	NonCompliantNotification *bool `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	// The ARN of the OSS bucket to which the delivered data is transferred when the size of the data exceeds the specified upper limit of the delivery channel.
	//
	// example:
	//
	// acs:oss:cn-shanghai:100931896542****:new-bucket
	OversizedDataOSSTargetArn *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	// The status of the delivery channel. Valid values:
	//
	// 	- 0: The delivery channel is disabled.
	//
	// 	- 1: The delivery channel is enabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetAccountId() *string {
	return s.AccountId
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetCompliantSnapshot() *bool {
	return s.CompliantSnapshot
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetConfigurationItemChangeNotification() *bool {
	return s.ConfigurationItemChangeNotification
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetConfigurationSnapshot() *bool {
	return s.ConfigurationSnapshot
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelAssumeRoleArn() *string {
	return s.DeliveryChannelAssumeRoleArn
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelCondition() *string {
	return s.DeliveryChannelCondition
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelTargetArn() *string {
	return s.DeliveryChannelTargetArn
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelType() *string {
	return s.DeliveryChannelType
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliverySnapshotTime() *string {
	return s.DeliverySnapshotTime
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetDescription() *string {
	return s.Description
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetNonCompliantNotification() *bool {
	return s.NonCompliantNotification
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetOversizedDataOSSTargetArn() *string {
	return s.OversizedDataOSSTargetArn
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) GetStatus() *int32 {
	return s.Status
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetAccountId(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.AccountId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetAggregatorId(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetCompliantSnapshot(v bool) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.CompliantSnapshot = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetConfigurationItemChangeNotification(v bool) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetConfigurationSnapshot(v bool) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelAssumeRoleArn(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelCondition(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelId(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelName(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelTargetArn(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelType(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelType = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliverySnapshotTime(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliverySnapshotTime = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetDescription(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.Description = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetNonCompliantNotification(v bool) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.NonCompliantNotification = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetOversizedDataOSSTargetArn(v string) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) SetStatus(v int32) *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.Status = &v
	return s
}

func (s *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel) Validate() error {
	return dara.Validate(s)
}
