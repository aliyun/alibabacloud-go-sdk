// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigDeliveryChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannels(v []*ListConfigDeliveryChannelsResponseBodyDeliveryChannels) *ListConfigDeliveryChannelsResponseBody
	GetDeliveryChannels() []*ListConfigDeliveryChannelsResponseBodyDeliveryChannels
	SetRequestId(v string) *ListConfigDeliveryChannelsResponseBody
	GetRequestId() *string
}

type ListConfigDeliveryChannelsResponseBody struct {
	// The information about the delivery channels.
	DeliveryChannels []*ListConfigDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DC300244-FCE3-5061-8214-C27ECB668487
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigDeliveryChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigDeliveryChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigDeliveryChannelsResponseBody) GetDeliveryChannels() []*ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	return s.DeliveryChannels
}

func (s *ListConfigDeliveryChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigDeliveryChannelsResponseBody) SetDeliveryChannels(v []*ListConfigDeliveryChannelsResponseBodyDeliveryChannels) *ListConfigDeliveryChannelsResponseBody {
	s.DeliveryChannels = v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBody) SetRequestId(v string) *ListConfigDeliveryChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBody) Validate() error {
	if s.DeliveryChannels != nil {
		for _, item := range s.DeliveryChannels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConfigDeliveryChannelsResponseBodyDeliveryChannels struct {
	// The ID of your Alibaba Cloud account.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Indicates whether the specified destination receives scheduled compliant snapshots. Cloud Config delivers scheduled compliant snapshots at `04:00Z` and `16:00Z` to  Log Service every day. The time is displayed in UTC. Valid values:
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
	// The ARN of the role assumed by the delivery channel.
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

func (s ListConfigDeliveryChannelsResponseBodyDeliveryChannels) String() string {
	return dara.Prettify(s)
}

func (s ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GoString() string {
	return s.String()
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetCompliantSnapshot() *bool {
	return s.CompliantSnapshot
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetConfigurationItemChangeNotification() *bool {
	return s.ConfigurationItemChangeNotification
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetConfigurationSnapshot() *bool {
	return s.ConfigurationSnapshot
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelAssumeRoleArn() *string {
	return s.DeliveryChannelAssumeRoleArn
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelCondition() *string {
	return s.DeliveryChannelCondition
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelTargetArn() *string {
	return s.DeliveryChannelTargetArn
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelType() *string {
	return s.DeliveryChannelType
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliverySnapshotTime() *string {
	return s.DeliverySnapshotTime
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDescription() *string {
	return s.Description
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetNonCompliantNotification() *bool {
	return s.NonCompliantNotification
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetOversizedDataOSSTargetArn() *string {
	return s.OversizedDataOSSTargetArn
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) GetStatus() *int32 {
	return s.Status
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetAccountId(v int64) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.AccountId = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetCompliantSnapshot(v bool) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.CompliantSnapshot = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationItemChangeNotification(v bool) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationSnapshot(v bool) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelAssumeRoleArn(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelCondition(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelId(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelName(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelTargetArn(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelType(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelType = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliverySnapshotTime(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliverySnapshotTime = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDescription(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.Description = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetNonCompliantNotification(v bool) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.NonCompliantNotification = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetOversizedDataOSSTargetArn(v string) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) SetStatus(v int32) *ListConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.Status = &v
	return s
}

func (s *ListConfigDeliveryChannelsResponseBodyDeliveryChannels) Validate() error {
	return dara.Validate(s)
}
