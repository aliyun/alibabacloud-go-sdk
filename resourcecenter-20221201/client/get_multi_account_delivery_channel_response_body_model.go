// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiAccountDeliveryChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryChannelDescription(v string) *GetMultiAccountDeliveryChannelResponseBody
	GetDeliveryChannelDescription() *string
	SetDeliveryChannelFilter(v *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) *GetMultiAccountDeliveryChannelResponseBody
	GetDeliveryChannelFilter() *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter
	SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelResponseBody
	GetDeliveryChannelId() *string
	SetDeliveryChannelName(v string) *GetMultiAccountDeliveryChannelResponseBody
	GetDeliveryChannelName() *string
	SetRequestId(v string) *GetMultiAccountDeliveryChannelResponseBody
	GetRequestId() *string
	SetResourceChangeDelivery(v *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) *GetMultiAccountDeliveryChannelResponseBody
	GetResourceChangeDelivery() *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery
	SetResourceSnapshotDelivery(v *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) *GetMultiAccountDeliveryChannelResponseBody
	GetResourceSnapshotDelivery() *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery
}

type GetMultiAccountDeliveryChannelResponseBody struct {
	// The description of the delivery channel.
	DeliveryChannelDescription *string `json:"DeliveryChannelDescription,omitempty" xml:"DeliveryChannelDescription,omitempty"`
	// The effective scope of the delivery channel.
	DeliveryChannelFilter *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter `json:"DeliveryChannelFilter,omitempty" xml:"DeliveryChannelFilter,omitempty" type:"Struct"`
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
	// FE3EAB47-D3A6-5FEA-8353-31C8C0D36***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The configurations for delivery of resource configuration change events.
	ResourceChangeDelivery *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery `json:"ResourceChangeDelivery,omitempty" xml:"ResourceChangeDelivery,omitempty" type:"Struct"`
	// The configurations for delivery of scheduled resource snapshots.
	ResourceSnapshotDelivery *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery `json:"ResourceSnapshotDelivery,omitempty" xml:"ResourceSnapshotDelivery,omitempty" type:"Struct"`
}

func (s GetMultiAccountDeliveryChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBody) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBody) GetDeliveryChannelDescription() *string {
	return s.DeliveryChannelDescription
}

func (s *GetMultiAccountDeliveryChannelResponseBody) GetDeliveryChannelFilter() *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter {
	return s.DeliveryChannelFilter
}

func (s *GetMultiAccountDeliveryChannelResponseBody) GetDeliveryChannelId() *string {
	return s.DeliveryChannelId
}

func (s *GetMultiAccountDeliveryChannelResponseBody) GetDeliveryChannelName() *string {
	return s.DeliveryChannelName
}

func (s *GetMultiAccountDeliveryChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMultiAccountDeliveryChannelResponseBody) GetResourceChangeDelivery() *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery {
	return s.ResourceChangeDelivery
}

func (s *GetMultiAccountDeliveryChannelResponseBody) GetResourceSnapshotDelivery() *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	return s.ResourceSnapshotDelivery
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelDescription(v string) *GetMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelDescription = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelFilter(v *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) *GetMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelFilter = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelId(v string) *GetMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelId = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetDeliveryChannelName(v string) *GetMultiAccountDeliveryChannelResponseBody {
	s.DeliveryChannelName = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetRequestId(v string) *GetMultiAccountDeliveryChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetResourceChangeDelivery(v *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) *GetMultiAccountDeliveryChannelResponseBody {
	s.ResourceChangeDelivery = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) SetResourceSnapshotDelivery(v *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) *GetMultiAccountDeliveryChannelResponseBody {
	s.ResourceSnapshotDelivery = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBody) Validate() error {
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

type GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter struct {
	// The effective account scopes of the delivery channel.
	AccountScopes []*string `json:"AccountScopes,omitempty" xml:"AccountScopes,omitempty" type:"Repeated"`
	// The effective resource types of the delivery channel.
	ResourceTypes []*string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty" type:"Repeated"`
}

func (s GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) GetAccountScopes() []*string {
	return s.AccountScopes
}

func (s *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) SetAccountScopes(v []*string) *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter {
	s.AccountScopes = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) SetResourceTypes(v []*string) *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter {
	s.ResourceTypes = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyDeliveryChannelFilter) Validate() error {
	return dara.Validate(s)
}

type GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery struct {
	Enabled *string `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The Simple Log Service configurations.
	SlsProperties *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
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

func (s GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) GetEnabled() *string {
	return s.Enabled
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) GetSlsProperties() *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties {
	return s.SlsProperties
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) SetEnabled(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery {
	s.Enabled = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) SetSlsProperties(v *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery {
	s.SlsProperties = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) SetTargetArn(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery {
	s.TargetArn = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) SetTargetType(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery {
	s.TargetType = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties struct {
	// The Alibaba Cloud Resource Name (ARN) of the destination to which large files are delivered.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceChangeDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}

type GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery struct {
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
	SlsProperties *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties `json:"SlsProperties,omitempty" xml:"SlsProperties,omitempty" type:"Struct"`
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

func (s GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) GetCustomExpression() *string {
	return s.CustomExpression
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) GetDeliveryTime() *string {
	return s.DeliveryTime
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) GetEnabled() *string {
	return s.Enabled
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) GetSlsProperties() *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties {
	return s.SlsProperties
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) GetTargetArn() *string {
	return s.TargetArn
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) GetTargetType() *string {
	return s.TargetType
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetCustomExpression(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.CustomExpression = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetDeliveryTime(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.DeliveryTime = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetEnabled(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.Enabled = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetSlsProperties(v *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.SlsProperties = v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetTargetArn(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.TargetArn = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) SetTargetType(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery {
	s.TargetType = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDelivery) Validate() error {
	if s.SlsProperties != nil {
		if err := s.SlsProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties struct {
	// The ARN of the destination to which large files are delivered.
	//
	// example:
	//
	// acs:oss:cn-hangzhou:1911422487776***:resourcecenter-oss
	OversizedDataOssTargetArn *string `json:"OversizedDataOssTargetArn,omitempty" xml:"OversizedDataOssTargetArn,omitempty"`
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) String() string {
	return dara.Prettify(s)
}

func (s GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) GoString() string {
	return s.String()
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) GetOversizedDataOssTargetArn() *string {
	return s.OversizedDataOssTargetArn
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) SetOversizedDataOssTargetArn(v string) *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties {
	s.OversizedDataOssTargetArn = &v
	return s
}

func (s *GetMultiAccountDeliveryChannelResponseBodyResourceSnapshotDeliverySlsProperties) Validate() error {
	return dara.Validate(s)
}
