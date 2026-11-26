// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKVCacheStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *CreateKVCacheStoreRequest
	GetCapacity() *int64
	SetClientToken(v string) *CreateKVCacheStoreRequest
	GetClientToken() *string
	SetDescription(v string) *CreateKVCacheStoreRequest
	GetDescription() *string
	SetHpnZone(v string) *CreateKVCacheStoreRequest
	GetHpnZone() *string
	SetName(v string) *CreateKVCacheStoreRequest
	GetName() *string
	SetPaymentType(v string) *CreateKVCacheStoreRequest
	GetPaymentType() *string
	SetRegionId(v string) *CreateKVCacheStoreRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateKVCacheStoreRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateKVCacheStoreRequestTag) *CreateKVCacheStoreRequest
	GetTag() []*CreateKVCacheStoreRequestTag
	SetZoneId(v string) *CreateKVCacheStoreRequest
	GetZoneId() *string
}

type CreateKVCacheStoreRequest struct {
	// The storage capacity, in GiB. The minimum capacity is 300 TiB (307,200 GiB), and the capacity scales in increments of 300 TiB.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2395
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The client token that is used to ensure request idempotence. The token can be up to 64 characters in length. Use a UUID.
	//
	// example:
	//
	// YOUR_CLIENT_TOKEN
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the KVCacheStore. The description must be 2 to 256 characters in length and cannot start with http:// or https://. Default value: empty.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The HPN cluster ID, which is used to create an affinity scheduling relationship between the KVCacheStore and the specified HPN cluster. After the KVCacheStore is created, it may have affinity relationships with multiple HPN clusters based on network topology. You can call the GetKVCacheStore operation to query the available HPN clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// B6
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The name of the KVCacheStore. The name must be 2 to 128 characters in length and can contain letters and digits. It can also contain colons (:), underscores (_), periods (.), and hyphens (-). If this parameter is not specified, the default value is the KVCacheStore ID.
	//
	// example:
	//
	// sc-data-warehouse-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing method. Valid value: POSTPAY (pay-as-you-go). Default value: POSTPAY.
	//
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The ID of the region where the KVCacheStore is created. You can call the DescribeRegions operation to query the list of activated regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzafsjd7i4qaq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tag key-value pairs for the resource. A maximum of 20 tags are supported. This parameter overrides the parent TagDTO type and uses the Tag type that is consistent with the Get and List responses.
	Tag []*CreateKVCacheStoreRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID. You can call the DescribeZones operation to query the list of zones in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateKVCacheStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKVCacheStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateKVCacheStoreRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *CreateKVCacheStoreRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKVCacheStoreRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKVCacheStoreRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *CreateKVCacheStoreRequest) GetName() *string {
	return s.Name
}

func (s *CreateKVCacheStoreRequest) GetPaymentType() *string {
	return s.PaymentType
}

func (s *CreateKVCacheStoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateKVCacheStoreRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateKVCacheStoreRequest) GetTag() []*CreateKVCacheStoreRequestTag {
	return s.Tag
}

func (s *CreateKVCacheStoreRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateKVCacheStoreRequest) SetCapacity(v int64) *CreateKVCacheStoreRequest {
	s.Capacity = &v
	return s
}

func (s *CreateKVCacheStoreRequest) SetClientToken(v string) *CreateKVCacheStoreRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKVCacheStoreRequest) SetDescription(v string) *CreateKVCacheStoreRequest {
	s.Description = &v
	return s
}

func (s *CreateKVCacheStoreRequest) SetHpnZone(v string) *CreateKVCacheStoreRequest {
	s.HpnZone = &v
	return s
}

func (s *CreateKVCacheStoreRequest) SetName(v string) *CreateKVCacheStoreRequest {
	s.Name = &v
	return s
}

func (s *CreateKVCacheStoreRequest) SetPaymentType(v string) *CreateKVCacheStoreRequest {
	s.PaymentType = &v
	return s
}

func (s *CreateKVCacheStoreRequest) SetRegionId(v string) *CreateKVCacheStoreRequest {
	s.RegionId = &v
	return s
}

func (s *CreateKVCacheStoreRequest) SetResourceGroupId(v string) *CreateKVCacheStoreRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateKVCacheStoreRequest) SetTag(v []*CreateKVCacheStoreRequestTag) *CreateKVCacheStoreRequest {
	s.Tag = v
	return s
}

func (s *CreateKVCacheStoreRequest) SetZoneId(v string) *CreateKVCacheStoreRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateKVCacheStoreRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateKVCacheStoreRequestTag struct {
	// The tag key of the resource.
	//
	// example:
	//
	// 000098da1005a3df
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource.
	//
	// example:
	//
	// 000088aabb0023f7
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s CreateKVCacheStoreRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateKVCacheStoreRequestTag) GoString() string {
	return s.String()
}

func (s *CreateKVCacheStoreRequestTag) GetTagKey() *string {
	return s.TagKey
}

func (s *CreateKVCacheStoreRequestTag) GetTagValue() *string {
	return s.TagValue
}

func (s *CreateKVCacheStoreRequestTag) SetTagKey(v string) *CreateKVCacheStoreRequestTag {
	s.TagKey = &v
	return s
}

func (s *CreateKVCacheStoreRequestTag) SetTagValue(v string) *CreateKVCacheStoreRequestTag {
	s.TagValue = &v
	return s
}

func (s *CreateKVCacheStoreRequestTag) Validate() error {
	return dara.Validate(s)
}
