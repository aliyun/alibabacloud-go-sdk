// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregateConfigDeliveryChannelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannels(v []*ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) *ListAggregateConfigDeliveryChannelsResponseBody
	GetDeliveryChannels() []*ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels
	SetRequestId(v string) *ListAggregateConfigDeliveryChannelsResponseBody
	GetRequestId() *string
}

type ListAggregateConfigDeliveryChannelsResponseBody struct {
	// The information about the delivery channels.
	DeliveryChannels []*ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DC300244-FCE3-5061-8214-C27ECB668487
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAggregateConfigDeliveryChannelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigDeliveryChannelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigDeliveryChannelsResponseBody) GetDeliveryChannels() []*ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	return s.DeliveryChannels
}

func (s *ListAggregateConfigDeliveryChannelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAggregateConfigDeliveryChannelsResponseBody) SetDeliveryChannels(v []*ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) *ListAggregateConfigDeliveryChannelsResponseBody {
	s.DeliveryChannels = v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBody) SetRequestId(v string) *ListAggregateConfigDeliveryChannelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBody) Validate() error {
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

type ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels struct {
	// The ID of the member in the account group.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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

func (s ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) String() string {
	return dara.Prettify(s)
}

func (s ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GoString() string {
	return s.String()
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetCompliantSnapshot() *bool {
	return s.CompliantSnapshot
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetConfigurationItemChangeNotification() *bool {
	return s.ConfigurationItemChangeNotification
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetConfigurationSnapshot() *bool {
	return s.ConfigurationSnapshot
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelAssumeRoleArn() *string {
	return s.DeliveryChannelAssumeRoleArn
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelCondition() *string {
	return s.DeliveryChannelCondition
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelTargetArn() *string {
	return s.DeliveryChannelTargetArn
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliveryChannelType() *string {
	return s.DeliveryChannelType
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDeliverySnapshotTime() *string {
	return s.DeliverySnapshotTime
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetDescription() *string {
	return s.Description
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetNonCompliantNotification() *bool {
	return s.NonCompliantNotification
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetOversizedDataOSSTargetArn() *string {
	return s.OversizedDataOSSTargetArn
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) GetStatus() *int32 {
	return s.Status
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetAccountId(v int64) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.AccountId = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetAggregatorId(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.AggregatorId = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetCompliantSnapshot(v bool) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.CompliantSnapshot = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationItemChangeNotification(v bool) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetConfigurationSnapshot(v bool) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelAssumeRoleArn(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelAssumeRoleArn = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelCondition(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelId(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelId = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelName(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelName = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelTargetArn(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliveryChannelType(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliveryChannelType = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDeliverySnapshotTime(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.DeliverySnapshotTime = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetDescription(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.Description = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetNonCompliantNotification(v bool) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.NonCompliantNotification = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetOversizedDataOSSTargetArn(v string) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) SetStatus(v int32) *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels {
	s.Status = &v
	return s
}

func (s *ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels) Validate() error {
	return dara.Validate(s)
}
