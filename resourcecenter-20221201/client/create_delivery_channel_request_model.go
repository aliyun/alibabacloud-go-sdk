// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelDescription(v string) *CreateDeliveryChannelRequest
	GetDeliveryChannelDescription() *string
	SetDeliveryChannelFilter(v *CreateDeliveryChannelRequestDeliveryChannelFilter) *CreateDeliveryChannelRequest
	GetDeliveryChannelFilter() *CreateDeliveryChannelRequestDeliveryChannelFilter
	SetDeliveryChannelName(v string) *CreateDeliveryChannelRequest
	GetDeliveryChannelName() *string
	SetResourceChangeDelivery(v *CreateDeliveryChannelRequestResourceChangeDelivery) *CreateDeliveryChannelRequest
	GetResourceChangeDelivery() *CreateDeliveryChannelRequestResourceChangeDelivery
	SetResourceSnapshotDelivery(v *CreateDeliveryChannelRequestResourceSnapshotDelivery) *CreateDeliveryChannelRequest
	GetResourceSnapshotDelivery() *CreateDeliveryChannelRequestResourceSnapshotDelivery
}

type CreateDeliveryChannelRequest struct {
	// The description of the delivery channel.
	//
	// example:
	//
	// 测试投递
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	//
	// This parameter is required.
	DeliveryChannelFilter *CreateDeliveryChannelRequestDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The name of the delivery channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-delivery
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The delivery of resource configuration changes.
	ResourceChangeDelivery *CreateDeliveryChannelRequestResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The scheduled delivery of resource snapshots.
	ResourceSnapshotDelivery *CreateDeliveryChannelRequestResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s CreateDeliveryChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequest) GetDeliveryChannelDescription() *string {
	return s.DeliveryChannelDescription
}

func (s *CreateDeliveryChannelRequest) GetDeliveryChannelFilter() *CreateDeliveryChannelRequestDeliveryChannelFilter {
	return s.DeliveryChannelFilter
}

func (s *CreateDeliveryChannelRequest) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *CreateDeliveryChannelRequest) GetResourceChangeDelivery() *CreateDeliveryChannelRequestResourceChangeDelivery {
	return s.ResourceChangeDelivery
}

func (s *CreateDeliveryChannelRequest) GetResourceSnapshotDelivery() *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	return s.ResourceSnapshotDelivery
}

func (s *CreateDeliveryChannelRequest) SetDeliveryChannelDescription(v string) *CreateDeliveryChannelRequest {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *CreateDeliveryChannelRequest) SetDeliveryChannelFilter(v *CreateDeliveryChannelRequestDeliveryChannelFilter) *CreateDeliveryChannelRequest {
	s.DeliveryChannelFilter = v
	return s
}

func (s *CreateDeliveryChannelRequest) SetDeliveryChannelName(v string) *CreateDeliveryChannelRequest {
	s.DeliveryChannelName = &v
	return s
}

func (s *CreateDeliveryChannelRequest) SetResourceChangeDelivery(v *CreateDeliveryChannelRequestResourceChangeDelivery) *CreateDeliveryChannelRequest {
	s.ResourceChangeDelivery = v
	return s
}

func (s *CreateDeliveryChannelRequest) SetResourceSnapshotDelivery(v *CreateDeliveryChannelRequestResourceSnapshotDelivery) *CreateDeliveryChannelRequest {
	s.ResourceSnapshotDelivery = v
	return s
}

func (s *CreateDeliveryChannelRequest) Validate() error {
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

type CreateDeliveryChannelRequestDeliveryChannelFilter struct {
	// The list of resource types to be delivered.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s CreateDeliveryChannelRequestDeliveryChannelFilter) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryChannelRequestDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestDeliveryChannelFilter) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *CreateDeliveryChannelRequestDeliveryChannelFilter) SetResourceTypes(v []*string) *CreateDeliveryChannelRequestDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

func (s *CreateDeliveryChannelRequestDeliveryChannelFilter) Validate() error {
	return dara.Validate(s)
}

type CreateDeliveryChannelRequestResourceChangeDelivery struct {
	// The SLS configurations.
	SlsProperties *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the destination. Valid values:
	//
	// - If you set `TargetType` to `OSS`, set `TargetArn` to the ARN of an OSS bucket that has the `resourcecenter-` prefix.
	//
	// - If you set `TargetType` to `SLS`, set `TargetArn` to the ARN of an SLS Logstore that has the `resourcecenter-` prefix.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the destination.
	//
	// Valid value: `SLS`.
	//
	// example:
	//
	// SLS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateDeliveryChannelRequestResourceChangeDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryChannelRequestResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) GetSlsProperties() *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	return s.SlsProperties
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) SetSlsProperties(v *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) *CreateDeliveryChannelRequestResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) SetTargetArn(v string) *CreateDeliveryChannelRequestResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) SetTargetType(v string) *CreateDeliveryChannelRequestResourceChangeDelivery {
	s.TargetType = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceChangeDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties struct {
	// The ARN of the destination OSS bucket for oversized files.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. Set this parameter to the ARN of an OSS bucket that has the `resourcecenter-` prefix.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:191142248777****:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}

type CreateDeliveryChannelRequestResourceSnapshotDelivery struct {
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
	// The SLS configurations.
	SlsProperties *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The Alibaba Cloud Resource Name (ARN) of the destination. Valid values:
	//
	// - If you set `TargetType` to `OSS`, set `TargetArn` to the ARN of an Object Storage Service (OSS) bucket that has the `resourcecenter-` prefix. Example: `acs:oss:cn-hangzhou:191142248777****:resourcecenter-oss`.
	//
	// - If you set `TargetType` to `SLS`, set `TargetArn` to the ARN of a Simple Log Service (SLS) Logstore that has the `resourcecenter-` prefix. Example: `acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls`.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the destination. Valid values:
	//
	// - For standard scheduled delivery, set this parameter to `OSS`.
	//
	// - For custom scheduled delivery, set this parameter to `OSS` or `SLS`.
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s CreateDeliveryChannelRequestResourceSnapshotDelivery) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryChannelRequestResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) GetCustomExpression() *string {
	return s.CustomExpression
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) GetDeliveryTime() *string {
	return s.DeliveryTime
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) GetSlsProperties() *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	return s.SlsProperties
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetCustomExpression(v string) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetDeliveryTime(v string) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetSlsProperties(v *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetTargetArn(v string) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) SetTargetType(v string) *CreateDeliveryChannelRequestResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination OSS bucket for oversized files.
	//
	// If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object. Set this parameter to the ARN of an OSS bucket that has the `resourcecenter-` prefix.
	//
	// > This parameter is valid only for custom scheduled delivery. You do not need to specify this parameter for standard scheduled delivery.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:191142248777****:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}
