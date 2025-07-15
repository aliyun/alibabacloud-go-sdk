// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCastersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterList(v *DescribeCastersResponseBodyCasterList) *DescribeCastersResponseBody
	GetCasterList() *DescribeCastersResponseBodyCasterList
	SetRequestId(v string) *DescribeCastersResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeCastersResponseBody
	GetTotal() *int32
}

type DescribeCastersResponseBody struct {
	// The production studios.
	CasterList *DescribeCastersResponseBodyCasterList `json:"CasterList,omitempty" xml:"CasterList,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of production studios.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeCastersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCastersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCastersResponseBody) GetCasterList() *DescribeCastersResponseBodyCasterList {
	return s.CasterList
}

func (s *DescribeCastersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCastersResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeCastersResponseBody) SetCasterList(v *DescribeCastersResponseBodyCasterList) *DescribeCastersResponseBody {
	s.CasterList = v
	return s
}

func (s *DescribeCastersResponseBody) SetRequestId(v string) *DescribeCastersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCastersResponseBody) SetTotal(v int32) *DescribeCastersResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeCastersResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCastersResponseBodyCasterList struct {
	Caster []*DescribeCastersResponseBodyCasterListCaster `json:"Caster,omitempty" xml:"Caster,omitempty" type:"Repeated"`
}

func (s DescribeCastersResponseBodyCasterList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCastersResponseBodyCasterList) GoString() string {
	return s.String()
}

func (s *DescribeCastersResponseBodyCasterList) GetCaster() []*DescribeCastersResponseBodyCasterListCaster {
	return s.Caster
}

func (s *DescribeCastersResponseBodyCasterList) SetCaster(v []*DescribeCastersResponseBodyCasterListCaster) *DescribeCastersResponseBodyCasterList {
	s.Caster = v
	return s
}

func (s *DescribeCastersResponseBodyCasterList) Validate() error {
	return dara.Validate(s)
}

type DescribeCastersResponseBodyCasterListCaster struct {
	// The ID of the production studio. You can specify the ID in a request to query the streaming URLs of the production studio, start the production studio, add a video resource, a layout, a component, or a playlist to the production studio, or query layouts of the production studio.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The name of the production studio.
	//
	// example:
	//
	// liveCaster****
	CasterName *string `json:"CasterName,omitempty" xml:"CasterName,omitempty"`
	// The resolution in which the production studio plays videos. This parameter is returned if the subscription billing method is used. Valid values:
	//
	// 	- lp_ld: low definition
	//
	// 	- lp_sd: standard definition
	//
	// 	- lp_hd: high definition
	//
	// 	- lp_ud: ultra high definition.
	//
	// 	- lp_ld_v: low definition (portrait mode)
	//
	// 	- lp_sd_v: standard definition (portrait mode)
	//
	// 	- lp_hd_v: high definition (portrait mode)
	//
	// 	- lp_ud_v: ultra high definition (portrait mode)
	//
	// example:
	//
	// lp_ld
	CasterTemplate *string `json:"CasterTemplate,omitempty" xml:"CasterTemplate,omitempty"`
	// Indicates whether the channel is enabled for the production studio.
	//
	// 	- 0: The channel is disabled.
	//
	// 	- 1: The channel is enabled.
	//
	// example:
	//
	// 1
	ChannelEnable *int32 `json:"ChannelEnable,omitempty" xml:"ChannelEnable,omitempty"`
	// The billing method. Valid values:
	//
	// 	- PrePaid: the subscription billing method
	//
	// 	- PostPaid: the pay-as-you-go billing method
	//
	// example:
	//
	// PrePaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	ClientTokenId *string `json:"ClientTokenId,omitempty" xml:"ClientTokenId,omitempty"`
	// The time when the production studio was created.
	//
	// example:
	//
	// 2017-08-30 12:02:57.0
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The streaming duration. Format: hh:mm:ss.
	//
	// example:
	//
	// 1:02:33
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The time when the production studio expires.
	//
	// example:
	//
	// 2018-08-30 12:02:57.0
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the production studio was last modified. For example, the time when the production studio was last started, stopped, or modified is returned.
	//
	// example:
	//
	// 2021-12-01T05:08:45Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The type of the production studio. Valid values:
	//
	// 	- 0: playlist mode
	//
	// 	- 1: general mode
	//
	// example:
	//
	// 1
	NormType *int32 `json:"NormType,omitempty" xml:"NormType,omitempty"`
	// The time when the production studio was purchased.
	//
	// example:
	//
	// 2017-08-30 12:02:57.0
	PurchaseTime *string `json:"PurchaseTime,omitempty" xml:"PurchaseTime,omitempty"`
	// The resource group ID. For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/2381067.html).
	//
	// example:
	//
	// rg-aekzw******
	ResourceGroupId *int32 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The time when the production studio was started. This parameter is returned if the production studio is in the streaming status.
	//
	// example:
	//
	// 2017-08-30 18:02:57.0
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the production studio. Valid values:
	//
	// 	- 0: idle
	//
	// 	- 1: streaming
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags.
	Tags *DescribeCastersResponseBodyCasterListCasterTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeCastersResponseBodyCasterListCaster) String() string {
	return dara.Prettify(s)
}

func (s DescribeCastersResponseBodyCasterListCaster) GoString() string {
	return s.String()
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetCasterName() *string {
	return s.CasterName
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetCasterTemplate() *string {
	return s.CasterTemplate
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetChannelEnable() *int32 {
	return s.ChannelEnable
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetClientTokenId() *string {
	return s.ClientTokenId
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetDuration() *string {
	return s.Duration
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetLastModified() *string {
	return s.LastModified
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetNormType() *int32 {
	return s.NormType
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetPurchaseTime() *string {
	return s.PurchaseTime
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetResourceGroupId() *int32 {
	return s.ResourceGroupId
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCastersResponseBodyCasterListCaster) GetTags() *DescribeCastersResponseBodyCasterListCasterTags {
	return s.Tags
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetCasterId(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.CasterId = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetCasterName(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.CasterName = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetCasterTemplate(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.CasterTemplate = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetChannelEnable(v int32) *DescribeCastersResponseBodyCasterListCaster {
	s.ChannelEnable = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetChargeType(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.ChargeType = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetClientTokenId(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.ClientTokenId = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetCreateTime(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.CreateTime = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetDuration(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.Duration = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetExpireTime(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.ExpireTime = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetLastModified(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.LastModified = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetNormType(v int32) *DescribeCastersResponseBodyCasterListCaster {
	s.NormType = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetPurchaseTime(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.PurchaseTime = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetResourceGroupId(v int32) *DescribeCastersResponseBodyCasterListCaster {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetStartTime(v string) *DescribeCastersResponseBodyCasterListCaster {
	s.StartTime = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetStatus(v int32) *DescribeCastersResponseBodyCasterListCaster {
	s.Status = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) SetTags(v *DescribeCastersResponseBodyCasterListCasterTags) *DescribeCastersResponseBodyCasterListCaster {
	s.Tags = v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCaster) Validate() error {
	return dara.Validate(s)
}

type DescribeCastersResponseBodyCasterListCasterTags struct {
	Tag []*DescribeCastersResponseBodyCasterListCasterTagsTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s DescribeCastersResponseBodyCasterListCasterTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCastersResponseBodyCasterListCasterTags) GoString() string {
	return s.String()
}

func (s *DescribeCastersResponseBodyCasterListCasterTags) GetTag() []*DescribeCastersResponseBodyCasterListCasterTagsTag {
	return s.Tag
}

func (s *DescribeCastersResponseBodyCasterListCasterTags) SetTag(v []*DescribeCastersResponseBodyCasterListCasterTagsTag) *DescribeCastersResponseBodyCasterListCasterTags {
	s.Tag = v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCasterTags) Validate() error {
	return dara.Validate(s)
}

type DescribeCastersResponseBodyCasterListCasterTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeCastersResponseBodyCasterListCasterTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCastersResponseBodyCasterListCasterTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCastersResponseBodyCasterListCasterTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeCastersResponseBodyCasterListCasterTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeCastersResponseBodyCasterListCasterTagsTag) SetTagKey(v string) *DescribeCastersResponseBodyCasterListCasterTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCasterTagsTag) SetTagValue(v string) *DescribeCastersResponseBodyCasterListCasterTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeCastersResponseBodyCasterListCasterTagsTag) Validate() error {
	return dara.Validate(s)
}
