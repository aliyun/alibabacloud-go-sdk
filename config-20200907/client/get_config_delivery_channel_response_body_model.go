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
	// The ID of the current Alibaba Cloud account.
	//
	// example:
	//
	// 120886317861****
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Indicates whether to deliver compliance snapshots. Cloud Config delivers compliance and non-compliance information of resources to SLS. Valid values:
	//
	// - true: Deliver compliance snapshots.
	//
	// - false: Do not deliver compliance snapshots.
	//
	// example:
	//
	// false
	CompliantSnapshot *bool `json:"CompliantSnapshot,omitempty" xml:"CompliantSnapshot,omitempty"`
	// Indicates whether to deliver resource configuration changes. When the configuration of a resource changes, Cloud Config delivers the resource configuration changes to OSS, SLS, or MNS. Valid values:
	//
	// - true: Deliver resource configuration changes.
	//
	// - false: Do not deliver resource configuration changes.
	//
	// example:
	//
	// true
	ConfigurationItemChangeNotification *bool `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	// Indicates whether to deliver scheduled resource snapshots. Cloud Config delivers scheduled resource snapshots to OSS or SLS at `04:00Z` and `16:00Z` (UTC) every day. Valid values:
	//
	// - true: Deliver scheduled resource snapshots.
	//
	// - false: Do not deliver scheduled resource snapshots.
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
	// The rule that is attached to the delivery channel. This parameter is available only for delivery channels of the MNS type and for snapshot deliveries to delivery channels of the SLS type.
	//
	// - If you deliver data to an MNS topic, you can specify the lowest risk level of the events to subscribe to and the resource types to subscribe to.
	//
	//   - The lowest risk level of the events to subscribe to: `{"filterType":"RuleRiskLevel","value":"1","multiple":false}`.
	//
	//     The \\`value\\` parameter indicates the risk level. Valid values: 1 (high), 2 (medium), and 3 (low).
	//
	//   - The resource types to subscribe to: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//     The \\`values\\` parameter indicates the resource types. The value is a JSON array.
	//
	//     Example:
	//
	//     `[{"filterType":"ResourceType","values":["ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage","ACS::CDN::Domain","ACS::CEN::CenBandwidthPackage","ACS::CEN::CenInstance","ACS::CEN::Flowlog","ACS::DdosCoo::Instance"],"multiple":true}]`
	//
	// - If you deliver snapshots to an SLS Logstore, you can specify the resource types to deliver: `{"filterType":"ResourceType","values":["ACS::ACK::Cluster","ACS::ActionTrail::Trail","ACS::CBWP::CommonBandwidthPackage"],"multiple":true}`.
	//
	//   The \\`values\\` parameter indicates the resource types. The value is a JSON array.
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
	// - If the DeliveryChannelType parameter is set to OSS, this parameter is the ARN of the destination OSS bucket.
	//
	// - If the DeliveryChannelType parameter is set to MNS, this parameter is the ARN of the destination MNS topic.
	//
	// - If the DeliveryChannelType parameter is set to SLS, this parameter is the ARN of the destination Simple Log Service Logstore.
	//
	// example:
	//
	// acs:oss:cn-shanghai:120886317861****:new-bucket
	DeliveryChannelTargetArn *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	// The type of the delivery channel. Valid values:
	//
	// - OSS: Object Storage Service.
	//
	// - MNS: Simple Message Queue (formerly MNS).
	//
	// - SLS: Simple Log Service.
	//
	// example:
	//
	// OSS
	DeliveryChannelType *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	// The time when Cloud Config starts to deliver scheduled resource snapshots every day.
	//
	// The time is in the `HH:mmZ` format (UTC).
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
	// Indicates whether to deliver resource non-compliance events. When a resource is evaluated as non-compliant, Cloud Config delivers the non-compliance events to SLS or MNS. Valid values:
	//
	// - true: Deliver resource non-compliance events.
	//
	// - false: Do not deliver resource non-compliance events.
	//
	// example:
	//
	// false
	NonCompliantNotification *bool `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	// The ARN of the OSS bucket to which the delivered data is transferred when the size of the data exceeds the limit of the delivery channel.
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
