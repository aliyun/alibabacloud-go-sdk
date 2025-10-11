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
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *CreateDeliveryChannelRequestResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
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
	return dara.Validate(s)
}

type CreateDeliveryChannelRequestDeliveryChannelFilter struct {
	// An array of effective resource types for the delivery channel.
	//
	// 	- Example: ["ACS::VPC::VPC", "ACS::ECS::Instance"].
	//
	// 	- If you want to deliver items of all resource types supported by Resource Center, set this parameter to ["ALL"].
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
	// The Simple Log Service configurations.
	SlsProperties *CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination.
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with resourcecenter-.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with resourcecenter-.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination.
	//
	// Valid values:
	//
	// 	- `SLS`
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
	return dara.Validate(s)
}

type CreateDeliveryChannelRequestResourceChangeDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// 	- If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object.
	//
	// 	- You need to set this parameter to the ARN of a bucket whose name is prefixed with resourcecenter-.
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
	// The Simple Log Service configurations.
	SlsProperties *CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The Alibaba Cloud Resource Name (ARN) of the delivery destination.
	//
	// 	- If you set `TargetType` to `OSS`, you must set `TargetArn` to the ARN of a bucket whose name is prefixed with resourcecenter-.
	//
	// 	- If you set `TargetType` to `SLS`, you must set `TargetArn` to the ARN of a Logstore whose name is prefixed with resourcecenter-.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 191142248777****:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the delivery destination.
	//
	// Valid values:
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
	return dara.Validate(s)
}

type CreateDeliveryChannelRequestResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// 	- If the size of a resource configuration change event exceeds 1 MB, the event is delivered as an OSS object.
	//
	// 	- You need to set this parameter to the ARN of a bucket whose name is prefixed with resourcecenter-.
	//
	// 	- This parameter takes effect only if you use custom delivery for scheduled resource snapshots. You do not need to configure this parameter if you use standard delivery for scheduled resource snapshots.
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
