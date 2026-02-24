// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAggregateConfigDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *UpdateAggregateConfigDeliveryChannelRequest
	GetAggregatorId() *string
	SetClientToken(v string) *UpdateAggregateConfigDeliveryChannelRequest
	GetClientToken() *string
	SetCompliantSnapshot(v bool) *UpdateAggregateConfigDeliveryChannelRequest
	GetCompliantSnapshot() *bool
	SetConfigurationItemChangeNotification(v bool) *UpdateAggregateConfigDeliveryChannelRequest
	GetConfigurationItemChangeNotification() *bool
	SetConfigurationSnapshot(v bool) *UpdateAggregateConfigDeliveryChannelRequest
	GetConfigurationSnapshot() *bool
	SetDeliveryChannelCondition(v string) *UpdateAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelCondition() *string
	SetDeliveryChannelId(v string) *UpdateAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelId() *string
	SetDeliveryChannelName(v string) *UpdateAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelName() *string
	SetDeliveryChannelTargetArn(v string) *UpdateAggregateConfigDeliveryChannelRequest
	GetDeliveryChannelTargetArn() *string
	SetDeliverySnapshotTime(v string) *UpdateAggregateConfigDeliveryChannelRequest
	GetDeliverySnapshotTime() *string
	SetDescription(v string) *UpdateAggregateConfigDeliveryChannelRequest
	GetDescription() *string
	SetNonCompliantNotification(v bool) *UpdateAggregateConfigDeliveryChannelRequest
	GetNonCompliantNotification() *bool
	SetOversizedDataOSSTargetArn(v string) *UpdateAggregateConfigDeliveryChannelRequest
	GetOversizedDataOSSTargetArn() *string
	SetStatus(v int64) *UpdateAggregateConfigDeliveryChannelRequest
	GetStatus() *int64
}

type UpdateAggregateConfigDeliveryChannelRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// if can be null:
	// true
	CompliantSnapshot *bool `json:"CompliantSnapshot,omitempty" xml:"CompliantSnapshot,omitempty"`
	// if can be null:
	// true
	ConfigurationItemChangeNotification *bool `json:"ConfigurationItemChangeNotification,omitempty" xml:"ConfigurationItemChangeNotification,omitempty"`
	// if can be null:
	// true
	ConfigurationSnapshot    *bool   `json:"ConfigurationSnapshot,omitempty" xml:"ConfigurationSnapshot,omitempty"`
	DeliveryChannelCondition *string `json:"DeliveryChannelCondition,omitempty" xml:"DeliveryChannelCondition,omitempty"`
	// This parameter is required.
	DeliveryChannelId        *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	DeliveryChannelName      *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	DeliveryChannelTargetArn *string `json:"DeliveryChannelTargetArn,omitempty" xml:"DeliveryChannelTargetArn,omitempty"`
	DeliverySnapshotTime     *string `json:"DeliverySnapshotTime,omitempty" xml:"DeliverySnapshotTime,omitempty"`
	Description              *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// if can be null:
	// true
	NonCompliantNotification  *bool   `json:"NonCompliantNotification,omitempty" xml:"NonCompliantNotification,omitempty"`
	OversizedDataOSSTargetArn *string `json:"OversizedDataOSSTargetArn,omitempty" xml:"OversizedDataOSSTargetArn,omitempty"`
	Status                    *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateAggregateConfigDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAggregateConfigDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetCompliantSnapshot() *bool {
	return s.CompliantSnapshot
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetConfigurationItemChangeNotification() *bool {
	return s.ConfigurationItemChangeNotification
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetConfigurationSnapshot() *bool {
	return s.ConfigurationSnapshot
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetDeliveryChannelCondition() *string {
	return s.DeliveryChannelCondition
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetDeliveryChannelTargetArn() *string {
	return s.DeliveryChannelTargetArn
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetDeliverySnapshotTime() *string {
	return s.DeliverySnapshotTime
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetNonCompliantNotification() *bool {
	return s.NonCompliantNotification
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetOversizedDataOSSTargetArn() *string {
	return s.OversizedDataOSSTargetArn
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) GetStatus() *int64 {
	return s.Status
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetAggregatorId(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.AggregatorId = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetClientToken(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetCompliantSnapshot(v bool) *UpdateAggregateConfigDeliveryChannelRequest {
	s.CompliantSnapshot = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetConfigurationItemChangeNotification(v bool) *UpdateAggregateConfigDeliveryChannelRequest {
	s.ConfigurationItemChangeNotification = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetConfigurationSnapshot(v bool) *UpdateAggregateConfigDeliveryChannelRequest {
	s.ConfigurationSnapshot = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelCondition(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelCondition = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelId(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelName(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDeliveryChannelTargetArn(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.DeliveryChannelTargetArn = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDeliverySnapshotTime(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.DeliverySnapshotTime = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetDescription(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.Description = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetNonCompliantNotification(v bool) *UpdateAggregateConfigDeliveryChannelRequest {
	s.NonCompliantNotification = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetOversizedDataOSSTargetArn(v string) *UpdateAggregateConfigDeliveryChannelRequest {
	s.OversizedDataOSSTargetArn = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) SetStatus(v int64) *UpdateAggregateConfigDeliveryChannelRequest {
	s.Status = &v
	return s
}

func (s *UpdateAggregateConfigDeliveryChannelRequest) Validate() error {
	return dara.Validate(s)
}
