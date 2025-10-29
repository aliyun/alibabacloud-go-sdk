// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCasterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterName(v string) *CreateCasterRequest
	GetCasterName() *string
	SetCasterTemplate(v string) *CreateCasterRequest
	GetCasterTemplate() *string
	SetChargeType(v string) *CreateCasterRequest
	GetChargeType() *string
	SetClientToken(v string) *CreateCasterRequest
	GetClientToken() *string
	SetExpireTime(v string) *CreateCasterRequest
	GetExpireTime() *string
	SetNormType(v int32) *CreateCasterRequest
	GetNormType() *int32
	SetOwnerId(v int64) *CreateCasterRequest
	GetOwnerId() *int64
	SetPurchaseTime(v string) *CreateCasterRequest
	GetPurchaseTime() *string
	SetRegionId(v string) *CreateCasterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCasterRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateCasterRequestTag) *CreateCasterRequest
	GetTag() []*CreateCasterRequestTag
}

type CreateCasterRequest struct {
	// The name of the production studio.
	//
	// example:
	//
	// liveCaster****
	CasterName *string `json:"CasterName,omitempty" xml:"CasterName,omitempty"`
	// The preset resolution of the production studio. If the subscription billing method is used, this parameter supports the following valid values:
	//
	// 	- **lp_ld**: low definition
	//
	// 	- **lp_sd**: standard definition
	//
	// 	- **lp_hd**: high definition
	//
	// 	- **lp_ud**: ultra high definition
	//
	// 	- **lp_ld_v**: low definition (portrait mode)
	//
	// 	- **lp_sd_v**: standard definition (portrait mode)
	//
	// 	- **lp_hd_v**: high definition (portrait mode)
	//
	// 	- **lp_ud_v**: ultra high definition (portrait mode)
	//
	// >  If the pay-as-you-go billing method is used, you must call the [SetCasterConfig](https://help.aliyun.com/document_detail/60271.html) operation to specify the resolution.
	//
	// example:
	//
	// lp_sd
	CasterTemplate *string `json:"CasterTemplate,omitempty" xml:"CasterTemplate,omitempty"`
	// The billing method. Only the pay-as-you-go billing method is supported.***	- Valid values:
	//
	// 	- **PrePaid**: subscription. This billing method is not yet supported.
	//
	// 	- **PostPaid**: pay-as-you-go
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can specify a custom value for this parameter, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The expiration time of the production studio. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  This parameter is valid only if you set the **ChargeType*	- parameter to **PrePaid**.
	//
	// example:
	//
	// 2017-08-22T12:10:10Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The type of the production studio. Valid values:
	//
	// 	- **1**: general mode
	//
	// 	- **6**: playlist mode (for carousel playback)
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	NormType *int32 `json:"NormType,omitempty" xml:"NormType,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The time when the production studio was purchased. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  This parameter is valid only if you set the **ChargeType*	- parameter to **PrePaid**.
	//
	// example:
	//
	// 2017-08-20T12:10:10Z
	PurchaseTime *string `json:"PurchaseTime,omitempty" xml:"PurchaseTime,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. For more information about resource groups, see [Resource groups](https://help.aliyun.com/document_detail/2381067.html).
	//
	// example:
	//
	// rg-aekzw******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags.
	Tag []*CreateCasterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateCasterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCasterRequest) GoString() string {
	return s.String()
}

func (s *CreateCasterRequest) GetCasterName() *string {
	return s.CasterName
}

func (s *CreateCasterRequest) GetCasterTemplate() *string {
	return s.CasterTemplate
}

func (s *CreateCasterRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateCasterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCasterRequest) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateCasterRequest) GetNormType() *int32 {
	return s.NormType
}

func (s *CreateCasterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCasterRequest) GetPurchaseTime() *string {
	return s.PurchaseTime
}

func (s *CreateCasterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCasterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCasterRequest) GetTag() []*CreateCasterRequestTag {
	return s.Tag
}

func (s *CreateCasterRequest) SetCasterName(v string) *CreateCasterRequest {
	s.CasterName = &v
	return s
}

func (s *CreateCasterRequest) SetCasterTemplate(v string) *CreateCasterRequest {
	s.CasterTemplate = &v
	return s
}

func (s *CreateCasterRequest) SetChargeType(v string) *CreateCasterRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateCasterRequest) SetClientToken(v string) *CreateCasterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCasterRequest) SetExpireTime(v string) *CreateCasterRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateCasterRequest) SetNormType(v int32) *CreateCasterRequest {
	s.NormType = &v
	return s
}

func (s *CreateCasterRequest) SetOwnerId(v int64) *CreateCasterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCasterRequest) SetPurchaseTime(v string) *CreateCasterRequest {
	s.PurchaseTime = &v
	return s
}

func (s *CreateCasterRequest) SetRegionId(v string) *CreateCasterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCasterRequest) SetResourceGroupId(v string) *CreateCasterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCasterRequest) SetTag(v []*CreateCasterRequestTag) *CreateCasterRequest {
	s.Tag = v
	return s
}

func (s *CreateCasterRequest) Validate() error {
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

type CreateCasterRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCasterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCasterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCasterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCasterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCasterRequestTag) SetKey(v string) *CreateCasterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCasterRequestTag) SetValue(v string) *CreateCasterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCasterRequestTag) Validate() error {
	return dara.Validate(s)
}
