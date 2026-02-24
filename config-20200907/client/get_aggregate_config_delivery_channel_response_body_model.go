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
	DeliveryChannel *GetAggregateConfigDeliveryChannelResponseBodyDeliveryChannel `json:"DeliveryChannel,omitempty" xml:"DeliveryChannel,omitempty" type:"Struct"`
	RequestId       *string                                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId                           *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	AggregatorId                        *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	CompliantSnapshot                   *bool   `json:"CompliantSnapshot,omitempty" xml:"CompliantSnapshot,omitempty"`
	ConfigurationItemChangeNotification *bool   `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	ConfigurationSnapshot               *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelAssumeRoleArn        *string `json:"DeliveryChannelAssumeRoleArn,omitempty" xml:"DeliveryChannelAssumeRoleArn,omitempty"`
	DeliveryChannelCondition            *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	DeliveryChannelId                   *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	DeliveryChannelName                 *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn            *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	DeliveryChannelType                 *string `json:"DeliveryChannelType,omitempty" xml:"DeliveryChannelType,omitempty"`
	DeliverySnapshotTime                *string `json:"DeliverySnapshotTime,omitempty" xml:"DeliverySnapshotTime,omitempty"`
	Description                         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	NonCompliantNotification            *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn           *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	Status                              *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
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
