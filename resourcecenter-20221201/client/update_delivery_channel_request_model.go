// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelDescription(v string) *UpdateDeliveryChannelRequest
	GetDeliveryChannelDescription() *string
	SetDeliveryChannelFilter(v *UpdateDeliveryChannelRequestDeliveryChannelFilter) *UpdateDeliveryChannelRequest
	GetDeliveryChannelFilter() *UpdateDeliveryChannelRequestDeliveryChannelFilter
	SetDeliveryChannelId(v string) *UpdateDeliveryChannelRequest
	GetDeliveryChannelId() *string
	SetDeliveryChannelName(v string) *UpdateDeliveryChannelRequest
	GetDeliveryChannelName() *string
	SetResourceChangeDelivery(v *UpdateDeliveryChannelRequestResourceChangeDelivery) *UpdateDeliveryChannelRequest
	GetResourceChangeDelivery() *UpdateDeliveryChannelRequestResourceChangeDelivery
	SetResourceSnapshotDelivery(v *UpdateDeliveryChannelRequestResourceSnapshotDelivery) *UpdateDeliveryChannelRequest
	GetResourceSnapshotDelivery() *UpdateDeliveryChannelRequestResourceSnapshotDelivery
}

type UpdateDeliveryChannelRequest struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	DeliveryChannelFilter *UpdateDeliveryChannelRequestDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
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
	// test-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *UpdateDeliveryChannelRequestResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *UpdateDeliveryChannelRequestResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s UpdateDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequest) GetDeliveryChannelDescription() *string {
	return s.DeliveryChannelDescription
}

func (s *UpdateDeliveryChannelRequest) GetDeliveryChannelFilter() *UpdateDeliveryChannelRequestDeliveryChannelFilter {
	return s.DeliveryChannelFilter
}

func (s *UpdateDeliveryChannelRequest) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *UpdateDeliveryChannelRequest) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *UpdateDeliveryChannelRequest) GetResourceChangeDelivery() *UpdateDeliveryChannelRequestResourceChangeDelivery {
	return s.ResourceChangeDelivery
}

func (s *UpdateDeliveryChannelRequest) GetResourceSnapshotDelivery() *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	return s.ResourceSnapshotDelivery
}

func (s *UpdateDeliveryChannelRequest) SetDeliveryChannelDescription(v string) *UpdateDeliveryChannelRequest {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetDeliveryChannelFilter(v *UpdateDeliveryChannelRequestDeliveryChannelFilter) *UpdateDeliveryChannelRequest {
	s.DeliveryChannelFilter = v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetDeliveryChannelId(v string) *UpdateDeliveryChannelRequest {
	s.DeliveryChannelId = &v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetDeliveryChannelName(v string) *UpdateDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetResourceChangeDelivery(v *UpdateDeliveryChannelRequestResourceChangeDelivery) *UpdateDeliveryChannelRequest {
	s.ResourceChangeDelivery = v
	return s
}

func (s *UpdateDeliveryChannelRequest) SetResourceSnapshotDelivery(v *UpdateDeliveryChannelRequestResourceSnapshotDelivery) *UpdateDeliveryChannelRequest {
	s.ResourceSnapshotDelivery = v
	return s
}

func (s *UpdateDeliveryChannelRequest) Validate() error {
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

type UpdateDeliveryChannelRequestDeliveryChannelFilter struct {
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s UpdateDeliveryChannelRequestDeliveryChannelFilter) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryChannelRequestDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestDeliveryChannelFilter) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *UpdateDeliveryChannelRequestDeliveryChannelFilter) SetResourceTypes(v []*string) *UpdateDeliveryChannelRequestDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

func (s *UpdateDeliveryChannelRequestDeliveryChannelFilter) Validate() error {
	return dara.Validate(s)
}

type UpdateDeliveryChannelRequestResourceChangeDelivery struct {
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
	SlsProperties *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
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

func (s UpdateDeliveryChannelRequestResourceChangeDelivery) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryChannelRequestResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) GetEnabled() *string {
	return s.Enabled
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) GetSlsProperties() *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	return s.SlsProperties
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) SetEnabled(v string) *UpdateDeliveryChannelRequestResourceChangeDelivery {
	s.Enabled = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) SetSlsProperties(v *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) *UpdateDeliveryChannelRequestResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) SetTargetArn(v string) *UpdateDeliveryChannelRequestResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) SetTargetType(v string) *UpdateDeliveryChannelRequestResourceChangeDelivery {
	s.TargetType = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceChangeDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceChangeDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}

type UpdateDeliveryChannelRequestResourceSnapshotDelivery struct {
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
	SlsProperties *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
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

func (s UpdateDeliveryChannelRequestResourceSnapshotDelivery) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryChannelRequestResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) GetCustomExpression() *string {
	return s.CustomExpression
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) GetDeliveryTime() *string {
	return s.DeliveryTime
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) GetEnabled() *string {
	return s.Enabled
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) GetSlsProperties() *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	return s.SlsProperties
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetCustomExpression(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetDeliveryTime(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetEnabled(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.Enabled = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetSlsProperties(v *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetTargetArn(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) SetTargetType(v string) *UpdateDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties struct {
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

func (s UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *UpdateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}
