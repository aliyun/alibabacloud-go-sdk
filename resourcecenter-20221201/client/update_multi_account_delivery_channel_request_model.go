// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMultiAccountDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelDescription(v string) *UpdateMultiAccountDeliveryChannelRequest
	GetDeliveryChannelDescription() *string
	SetDeliveryChannelFilter(v *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) *UpdateMultiAccountDeliveryChannelRequest
	GetDeliveryChannelFilter() *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter
	SetDeliveryChannelId(v string) *UpdateMultiAccountDeliveryChannelRequest
	GetDeliveryChannelId() *string
	SetDeliveryChannelName(v string) *UpdateMultiAccountDeliveryChannelRequest
	GetDeliveryChannelName() *string
	SetResourceChangeDelivery(v *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) *UpdateMultiAccountDeliveryChannelRequest
	GetResourceChangeDelivery() *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery
	SetResourceSnapshotDelivery(v *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) *UpdateMultiAccountDeliveryChannelRequest
	GetResourceSnapshotDelivery() *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery
}

type UpdateMultiAccountDeliveryChannelRequest struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	DeliveryChannelFilter *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The ID of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-4m6ffpkgu***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-multi-account-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s UpdateMultiAccountDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequest) GetDeliveryChannelDescription() *string {
	return s.DeliveryChannelDescription
}

func (s *UpdateMultiAccountDeliveryChannelRequest) GetDeliveryChannelFilter() *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	return s.DeliveryChannelFilter
}

func (s *UpdateMultiAccountDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *UpdateMultiAccountDeliveryChannelRequest) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *UpdateMultiAccountDeliveryChannelRequest) GetResourceChangeDelivery() *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	return s.ResourceChangeDelivery
}

func (s *UpdateMultiAccountDeliveryChannelRequest) GetResourceSnapshotDelivery() *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	return s.ResourceSnapshotDelivery
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetDeliveryChannelDescription(v string) *UpdateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetDeliveryChannelFilter(v *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) *UpdateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelFilter = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetDeliveryChannelId(v string) *UpdateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetDeliveryChannelName(v string) *UpdateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetResourceChangeDelivery(v *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) *UpdateMultiAccountDeliveryChannelRequest {
	s.ResourceChangeDelivery = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) SetResourceSnapshotDelivery(v *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) *UpdateMultiAccountDeliveryChannelRequest {
	s.ResourceSnapshotDelivery = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequest) Validate() error {
	if s.DeliveryChannelFilter != nil {
		if err := s.DeliveryChannelFilter.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceChangeDelivery != nil {
		if err := s.ResourceChangeDelivery.Validate(); err != nil {
			return err
		}
	}
	if s.ResourceSnapshotDelivery != nil {
		if err := s.ResourceSnapshotDelivery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter struct {
	// An array of effective account scopes for the delivery channel.
	AccountScopes []*string `json:"AccountScopes,omitempty" xml:"AccountScopes,omitempty" type:"Repeated"`
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) GetAccountScopes() []*string {
	return s.AccountScopes
}

func (s *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) SetAccountScopes(v []*string) *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	s.AccountScopes = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) SetResourceTypes(v []*string) *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) Validate() error {
	return dara.Validate(s)
}

type UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery struct {
	// Specifies whether to enable delivery of resource configuration change events. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination. Valid values:
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 1911422487776***:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination.
	//
	// Set the value to `SLS`.
	//
	// example:
	//
	// SLS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GetEnabled() *string {
	return s.Enabled
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GetSlsProperties() *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	return s.SlsProperties
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetEnabled(v string) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.Enabled = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetSlsProperties(v *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetTargetArn(v string) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetTargetType(v string) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.TargetType = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}

type UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery struct {
	// The custom expression.
	//
	// example:
	//
	// select 	- from resources limit 100;
	CustomExpression *string `json:"CustomExpression,omitempty" xml:"CustomExpression,omitempty"`
	// The delivery time.
	//
	// example:
	//
	// 09:00Z
	DeliveryTime *string `json:"DeliveryTime,omitempty" xml:"DeliveryTime,omitempty"`
	// Specifies whether to enable delivery of scheduled resource snapshots. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The Alibaba Cloud Resource Name (ARN) of the delivery destination. Valid values:
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with `resourcecenter-`. Example: `acs:oss:cn-hangzhou:191142248777****:resourcecenter-oss`.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with `resourcecenter-`. Example: `acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination. Valid values:
	//
	// 	- `OSS` for standard delivery
	//
	// 	- `OSS` or `SLS` for custom delivery
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetCustomExpression() *string {
	return s.CustomExpression
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetDeliveryTime() *string {
	return s.DeliveryTime
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetEnabled() *string {
	return s.Enabled
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetSlsProperties() *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	return s.SlsProperties
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetCustomExpression(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetDeliveryTime(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetEnabled(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.Enabled = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetSlsProperties(v *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetTargetArn(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetTargetType(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// >  This parameter takes effect only if you use custom delivery for scheduled resource snapshots. You do not need to configure this parameter if you use standard delivery for scheduled resource snapshots.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *UpdateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}
