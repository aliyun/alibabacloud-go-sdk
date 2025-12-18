// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConfigDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannel(v *GetConfigDeliveryChannelResponseBodyDeliveryChannel) *GetConfigDeliveryChannelResponseBody
	GetDeliveryChannel() *GetConfigDeliveryChannelResponseBodyDeliveryChannel
	SetRequestId(v string) *GetConfigDeliveryChannelResponseBody
	GetRequestId() *string
}

type GetConfigDeliveryChannelResponseBody struct {
	// The information about the delivery channel.
	DeliveryChannel *GetConfigDeliveryChannelResponseBodyDeliveryChannel `json:"DeliveryChannel,omitempty" xml:"DeliveryChannel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DC300244-FCE3-5061-8214-C27ECB668487
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetConfigDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConfigDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetConfigDeliveryChannelResponseBody) GetDeliveryChannel() *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	return s.DeliveryChannel
}

func (s *GetConfigDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConfigDeliveryChannelResponseBody) SetDeliveryChannel(v *GetConfigDeliveryChannelResponseBodyDeliveryChannel) *GetConfigDeliveryChannelResponseBody {
	s.DeliveryChannel = v
	return s
}

func (s *GetConfigDeliveryChannelResponseBody) SetRequestId(v string) *GetConfigDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBody) Validate() error {
	if s.DeliveryChannel != nil {
		if err := s.DeliveryChannel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetConfigDeliveryChannelResponseBodyDeliveryChannel struct {
	// The ID of your Alibaba Cloud account.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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
	// The Alibaba Cloud Resource Name (ARN) of the role assumed by the delivery channel.
	//
	// example:
	//
	// acs:ram::120886317861****:role/aliyunserviceroleforconfig
	DeliveryChannelAssumeRoleArn *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	// The rule that is attached to the delivery channel. This parameter is available when you deliver data of all types to MNS or deliver snapshots to Log Service.
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
	// The Alibaba Cloud Resource Name (ARN) of the delivery destination.
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
	// The ARN of the OSS bucket to which you want to transfer the delivery data when the size of the data exceeds the specified upper limit of the delivery channel.
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

func (s GetConfigDeliveryChannelResponseBodyDeliveryChannel) String() string {
	return dara.Prettify(s)
}

func (s GetConfigDeliveryChannelResponseBodyDeliveryChannel) GoString() string {
	return s.String()
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetAccountId() *int64 {
	return s.AccountId
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetCompliantSnapshot() *bool {
	return s.CompliantSnapshot
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetConfigurationItemChangeNotification() *bool {
	return s.ConfigurationItemChangeNotification
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetConfigurationSnapshot() *bool {
	return s.ConfigurationSnapshot
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelAssumeRoleArn() *string {
	return s.DeliveryChannelAssumeRoleArn
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelCondition() *string {
	return s.DeliveryChannelCondition
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelTargetArn() *string {
	return s.DeliveryChannelTargetArn
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliveryChannelType() *string {
	return s.DeliveryChannelType
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetDeliverySnapshotTime() *string {
	return s.DeliverySnapshotTime
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetDescription() *string {
	return s.Description
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetNonCompliantNotification() *bool {
	return s.NonCompliantNotification
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetOversizedDataOSSTargetArn() *string {
	return s.OversizedDataOSSTargetArn
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) GetStatus() *int32 {
	return s.Status
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetAccountId(v int64) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.AccountId = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetCompliantSnapshot(v bool) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.CompliantSnapshot = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetConfigurationItemChangeNotification(v bool) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetConfigurationSnapshot(v bool) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelAssumeRoleArn(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelCondition(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelId(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelName(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelTargetArn(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliveryChannelType(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliveryChannelType = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDeliverySnapshotTime(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.DeliverySnapshotTime = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetDescription(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.Description = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetNonCompliantNotification(v bool) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.NonCompliantNotification = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetOversizedDataOSSTargetArn(v string) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) SetStatus(v int32) *GetConfigDeliveryChannelResponseBodyDeliveryChannel {
	s.Status = &v
	return s
}

func (s *GetConfigDeliveryChannelResponseBodyDeliveryChannel) Validate() error {
	return dara.Validate(s)
}
