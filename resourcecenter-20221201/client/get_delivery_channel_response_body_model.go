// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelDescription(v string) *GetDeliveryChannelResponseBody
	GetDeliveryChannelDescription() *string
	SetDeliveryChannelFilter(v *GetDeliveryChannelResponseBodyDeliveryChannelFilter) *GetDeliveryChannelResponseBody
	GetDeliveryChannelFilter() *GetDeliveryChannelResponseBodyDeliveryChannelFilter
	SetDeliveryChannelId(v string) *GetDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetDeliveryChannelName(v string) *GetDeliveryChannelResponseBody
	GetDeliveryChannelName() *string
	SetRequestId(v string) *GetDeliveryChannelResponseBody
	GetRequestId() *string
	SetResourceChangeDelivery(v *GetDeliveryChannelResponseBodyResourceChangeDelivery) *GetDeliveryChannelResponseBody
	GetResourceChangeDelivery() *GetDeliveryChannelResponseBodyResourceChangeDelivery
	SetResourceSnapshotDelivery(v *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) *GetDeliveryChannelResponseBody
	GetResourceSnapshotDelivery() *GetDeliveryChannelResponseBodyResourceSnapshotDelivery
}

type GetDeliveryChannelResponseBody struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	DeliveryChannelFilter *GetDeliveryChannelResponseBodyDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
	// The ID of the delivery channel.
	//
	// example:
	//
	// dc-6q79dm4o9***
	DeliveryChannelId *string `json:"DeliveryChannelId,omitempty" xml:"DeliveryChannelId,omitempty"`
	// The name of the delivery channel.
	//
	// example:
	//
	// test-delivery-channel
	DeliveryChannelName *string `json:"DeliveryChannelName,omitempty" xml:"DeliveryChannelName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 17502A1B-7026-5551-8E1C-CCABD0CBC***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *GetDeliveryChannelResponseBodyResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *GetDeliveryChannelResponseBodyResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s GetDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBody) GetDeliveryChannelDescription() *string {
	return s.DeliveryChannelDescription
}

func (s *GetDeliveryChannelResponseBody) GetDeliveryChannelFilter() *GetDeliveryChannelResponseBodyDeliveryChannelFilter {
	return s.DeliveryChannelFilter
}

func (s *GetDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetDeliveryChannelResponseBody) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *GetDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDeliveryChannelResponseBody) GetResourceChangeDelivery() *GetDeliveryChannelResponseBodyResourceChangeDelivery {
	return s.ResourceChangeDelivery
}

func (s *GetDeliveryChannelResponseBody) GetResourceSnapshotDelivery() *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	return s.ResourceSnapshotDelivery
}

func (s *GetDeliveryChannelResponseBody) SetDeliveryChannelDescription(v string) *GetDeliveryChannelResponseBody {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetDeliveryChannelFilter(v *GetDeliveryChannelResponseBodyDeliveryChannelFilter) *GetDeliveryChannelResponseBody {
	s.DeliveryChannelFilter = v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *GetDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetDeliveryChannelName(v string) *GetDeliveryChannelResponseBody {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetRequestId(v string) *GetDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetResourceChangeDelivery(v *GetDeliveryChannelResponseBodyResourceChangeDelivery) *GetDeliveryChannelResponseBody {
	s.ResourceChangeDelivery = v
	return s
}

func (s *GetDeliveryChannelResponseBody) SetResourceSnapshotDelivery(v *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) *GetDeliveryChannelResponseBody {
	s.ResourceSnapshotDelivery = v
	return s
}

func (s *GetDeliveryChannelResponseBody) Validate() error {
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

type GetDeliveryChannelResponseBodyDeliveryChannelFilter struct {
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s GetDeliveryChannelResponseBodyDeliveryChannelFilter) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyDeliveryChannelFilter) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *GetDeliveryChannelResponseBodyDeliveryChannelFilter) SetResourceTypes(v []*string) *GetDeliveryChannelResponseBodyDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

func (s *GetDeliveryChannelResponseBodyDeliveryChannelFilter) Validate() error {
	return dara.Validate(s)
}

type GetDeliveryChannelResponseBodyResourceChangeDelivery struct {
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination.
	//
	// example:
	//
	// acs:log:cn-hangzhou: 1911422487776***:project/delivery/logstore/resourcecenter-sls
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the destination.
	//
	// example:
	//
	// SLS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetDeliveryChannelResponseBodyResourceChangeDelivery) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) GetEnabled() *string {
	return s.Enabled
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) GetSlsProperties() *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties {
	return s.SlsProperties
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) SetEnabled(v string) *GetDeliveryChannelResponseBodyResourceChangeDelivery {
	s.Enabled = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) SetSlsProperties(v *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) *GetDeliveryChannelResponseBodyResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) SetTargetArn(v string) *GetDeliveryChannelResponseBodyResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) SetTargetType(v string) *GetDeliveryChannelResponseBodyResourceChangeDelivery {
	s.TargetType = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties struct {
	// The Alibaba Cloud Resource Name (ARN) of the destination to which large files are delivered.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}

type GetDeliveryChannelResponseBodyResourceSnapshotDelivery struct {
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
	Enabled      *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
	// The ARN of the delivery destination.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	TargetArn *string `json:"TargetArn,omitempty" xml:"TargetArn,omitempty"`
	// The type of the destination.
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s GetDeliveryChannelResponseBodyResourceSnapshotDelivery) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) GetCustomExpression() *string {
	return s.CustomExpression
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) GetDeliveryTime() *string {
	return s.DeliveryTime
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) GetEnabled() *string {
	return s.Enabled
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) GetSlsProperties() *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties {
	return s.SlsProperties
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetCustomExpression(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetDeliveryTime(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetEnabled(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.Enabled = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetSlsProperties(v *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetTargetArn(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) SetTargetType(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *GetDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}
