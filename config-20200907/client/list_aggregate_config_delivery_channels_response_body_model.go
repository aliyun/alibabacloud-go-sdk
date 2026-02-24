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
	DeliveryChannels []*ListAggregateConfigDeliveryChannelsResponseBodyDeliveryChannels `json:"DeliveryChannels,omitempty" xml:"DeliveryChannels,omitempty" type:"Repeated"`
	RequestId        *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	AccountId                           *int64  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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
