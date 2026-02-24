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
	// A list of delivery channels.
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
	// The ID of the current Alibaba Cloud account.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Specifies whether to receive compliance snapshots. Cloud Config delivers resource compliance and non-compliance information to SLS. Valid values:
	//
	// - true: Receive compliance snapshots.
	//
	// - false: Do not receive compliance snapshots.
	//
	// example:
	//
	// false
	CompliantSnapshot *bool `json:"CompliantSnapshot,omitempty" xml:"CompliantSnapshot,omitempty"`
	// Specifies whether to receive the resource configuration history. When a resource configuration changes, Cloud Config delivers the resource configuration history to OSS, SLS, or MNS. Valid values:
	//
	// - true: Receive the resource configuration history.
	//
	// - false: Do not receive the resource configuration history.
	//
	// example:
	//
	// true
	ConfigurationItemChangeNotification *bool `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	// Specifies whether to receive scheduled resource snapshots. Cloud Config delivers scheduled resource snapshots to OSS or SLS at `04:00Z` and `16:00Z` (UTC) every day. Valid values:
	//
	// - true: Receive scheduled resource snapshots.
	//
	// - false: Do not receive scheduled resource snapshots.
	//
	// example:
	//
	// false
	ConfigurationSnapshot *bool `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	// The ARN of the role that is assumed by the delivery channel.
	//
	// example:
	//
	// acs:ram::120886317861****:role/aliyunserviceroleforconfig
	DeliveryChannelAssumeRoleArn *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	// The rule attached to the delivery channel. This parameter is supported for MNS channels and for snapshot delivery to SLS channels.
	//
	// - To subscribe to MNS events, specify the minimum risk level and resource types:
	//
	//   - Minimum risk level of subscribed events: `{"filterType":"RuleRiskLevel","value":"1","multiple":false}`.
	//
	//     \\`value\\` specifies the risk level. Valid values: 1 (high risk), 2 (medium risk), and 3 (low risk).
	//
	//   - Resource types of subscribed events: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//     \\`values\\` is a JSON array of the resource types.
	//
	//     Example:
	//
	//     `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	//
	// - To deliver snapshots to SLS, specify the resource types: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//   \\`values\\` is a JSON array of the resource types.
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
	// - If \\`DeliveryChannelType\\` is \\`OSS\\`, this parameter is the ARN of an OSS bucket.
	//
	// - If \\`DeliveryChannelType\\` is \\`MNS\\`, this parameter is the ARN of an MNS topic.
	//
	// - If \\`DeliveryChannelType\\` is \\`SLS\\`, this parameter is the ARN of an SLS Logstore.
	//
	// example:
	//
	// acs:oss:cn-shanghai:120886317861****:new-bucket
	DeliveryChannelTargetArn *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	// The type of the delivery channel. Valid values:
	//
	// - OSS: Object Storage Service (OSS).
	//
	// - MNS: Simple Message Queue (formerly MNS).
	//
	// - SLS: Simple Log Service (SLS).
	//
	// example:
	//
	// OSS
	DeliveryChannelType *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	// The time when scheduled resource snapshots start to be delivered every day.
	//
	// The format is `HH:mmZ` (UTC).
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
	// Specifies whether to receive resource non-compliance events. When a resource becomes non-compliant, Cloud Config delivers a resource non-compliance event to SLS or MNS. Valid values:
	//
	// - true: Receive resource non-compliance events.
	//
	// - false: Do not receive resource non-compliance events.
	//
	// example:
	//
	// false
	NonCompliantNotification *bool `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	// The destination OSS bucket to which the data is transferred when the size of the data exceeds the limit of the delivery channel.
	//
	// example:
	//
	// acs:oss:cn-shanghai:100931896542****:new-bucket
	OversizedDataOSSTargetArn *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	// The status of the delivery channel. Valid values:
	//
	// - 0: The delivery channel is disabled.
	//
	// - 1: The delivery channel is enabled.
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
