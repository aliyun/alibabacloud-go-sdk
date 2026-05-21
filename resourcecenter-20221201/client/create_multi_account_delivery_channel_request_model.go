// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultiAccountDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelDescription(v string) *CreateMultiAccountDeliveryChannelRequest
	GetDeliveryChannelDescription() *string
	SetDeliveryChannelFilter(v *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) *CreateMultiAccountDeliveryChannelRequest
	GetDeliveryChannelFilter() *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter
	SetDeliveryChannelName(v string) *CreateMultiAccountDeliveryChannelRequest
	GetDeliveryChannelName() *string
	SetResourceChangeDelivery(v *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) *CreateMultiAccountDeliveryChannelRequest
	GetResourceChangeDelivery() *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery
	SetResourceSnapshotDelivery(v *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) *CreateMultiAccountDeliveryChannelRequest
	GetResourceSnapshotDelivery() *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery
}

type CreateMultiAccountDeliveryChannelRequest struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	//
	// This parameter is required.
	DeliveryChannelFilter *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The name of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-multi-account-delivery
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s CreateMultiAccountDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequest) GetDeliveryChannelDescription() *string {
	return s.DeliveryChannelDescription
}

func (s *CreateMultiAccountDeliveryChannelRequest) GetDeliveryChannelFilter() *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	return s.DeliveryChannelFilter
}

func (s *CreateMultiAccountDeliveryChannelRequest) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *CreateMultiAccountDeliveryChannelRequest) GetResourceChangeDelivery() *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	return s.ResourceChangeDelivery
}

func (s *CreateMultiAccountDeliveryChannelRequest) GetResourceSnapshotDelivery() *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	return s.ResourceSnapshotDelivery
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetDeliveryChannelDescription(v string) *CreateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetDeliveryChannelFilter(v *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) *CreateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelFilter = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetDeliveryChannelName(v string) *CreateMultiAccountDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetResourceChangeDelivery(v *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) *CreateMultiAccountDeliveryChannelRequest {
	s.ResourceChangeDelivery = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequest) SetResourceSnapshotDelivery(v *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) *CreateMultiAccountDeliveryChannelRequest {
	s.ResourceSnapshotDelivery = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequest) Validate() error {
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

type CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter struct {
	// An array of effective account scopes for the delivery channel.
	//
	// This parameter is required.
	AccountScopes []*string `json:"AccountScopes,omitempty" xml:"AccountScopes,omitempty" type:"Repeated"`
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) GetAccountScopes() []*string {
	return s.AccountScopes
}

func (s *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) SetAccountScopes(v []*string) *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	s.AccountScopes = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) SetResourceTypes(v []*string) *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestDeliveryChannelFilter) Validate() error {
	return dara.Validate(s)
}

type CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery struct {
	// The Simple Log Service configurations.
	SlsProperties *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
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

func (s CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GetSlsProperties() *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	return s.SlsProperties
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetSlsProperties(v *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetTargetArn(v string) *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) SetTargetType(v string) *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery {
	s.TargetType = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. You need to set this parameter to the ARN of a bucket whose name is prefixed with `resourcecenter-`.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceChangeDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}

type CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery struct {
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
	// The Simple Log Service configurations.
	SlsProperties *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
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

func (s CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetCustomExpression() *string {
	return s.CustomExpression
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetDeliveryTime() *string {
	return s.DeliveryTime
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetSlsProperties() *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	return s.SlsProperties
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetCustomExpression(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetDeliveryTime(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetSlsProperties(v *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetTargetArn(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) SetTargetType(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties struct {
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

func (s CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *CreateMultiAccountDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}
