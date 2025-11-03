// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHaVipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateHaVipRequest
	GetClientToken() *string
	SetDescription(v string) *CreateHaVipRequest
	GetDescription() *string
	SetIpAddress(v string) *CreateHaVipRequest
	GetIpAddress() *string
	SetName(v string) *CreateHaVipRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateHaVipRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateHaVipRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateHaVipRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateHaVipRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateHaVipRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateHaVipRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateHaVipRequestTag) *CreateHaVipRequest
	GetTag() []*CreateHaVipRequestTag
	SetVSwitchId(v string) *CreateHaVipRequest
	GetVSwitchId() *string
}

type CreateHaVipRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among all requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the HaVip.
	//
	// The description must be 1 to 255 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my HaVip.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP address of the HaVip.
	//
	// The specified IP address must be an idle IP address that falls within the CIDR block of the vSwitch. If this parameter is not set, an idle IP address from the CIDR block of the vSwitch is randomly assigned to the HaVip.
	//
	// example:
	//
	// 192.XX.XX.10
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the HaVip.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the HaVip. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the HaVip belongs.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag of the resource.
	Tag []*CreateHaVipRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which the HaVip belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-asdfjlnaue4g****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateHaVipRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHaVipRequest) GoString() string {
	return s.String()
}

func (s *CreateHaVipRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateHaVipRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHaVipRequest) GetIpAddress() *string {
	return s.IpAddress
}

func (s *CreateHaVipRequest) GetName() *string {
	return s.Name
}

func (s *CreateHaVipRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateHaVipRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateHaVipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHaVipRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateHaVipRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateHaVipRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateHaVipRequest) GetTag() []*CreateHaVipRequestTag {
	return s.Tag
}

func (s *CreateHaVipRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateHaVipRequest) SetClientToken(v string) *CreateHaVipRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateHaVipRequest) SetDescription(v string) *CreateHaVipRequest {
	s.Description = &v
	return s
}

func (s *CreateHaVipRequest) SetIpAddress(v string) *CreateHaVipRequest {
	s.IpAddress = &v
	return s
}

func (s *CreateHaVipRequest) SetName(v string) *CreateHaVipRequest {
	s.Name = &v
	return s
}

func (s *CreateHaVipRequest) SetOwnerAccount(v string) *CreateHaVipRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateHaVipRequest) SetOwnerId(v int64) *CreateHaVipRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateHaVipRequest) SetRegionId(v string) *CreateHaVipRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHaVipRequest) SetResourceGroupId(v string) *CreateHaVipRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateHaVipRequest) SetResourceOwnerAccount(v string) *CreateHaVipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateHaVipRequest) SetResourceOwnerId(v int64) *CreateHaVipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateHaVipRequest) SetTag(v []*CreateHaVipRequestTag) *CreateHaVipRequest {
	s.Tag = v
	return s
}

func (s *CreateHaVipRequest) SetVSwitchId(v string) *CreateHaVipRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateHaVipRequest) Validate() error {
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

type CreateHaVipRequestTag struct {
	// The key of tag N to add to the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length, but cannot contain `http://` or `https://`. The tag value cannot start with `aliyun` or `acs:`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHaVipRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateHaVipRequestTag) GoString() string {
	return s.String()
}

func (s *CreateHaVipRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateHaVipRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateHaVipRequestTag) SetKey(v string) *CreateHaVipRequestTag {
	s.Key = &v
	return s
}

func (s *CreateHaVipRequestTag) SetValue(v string) *CreateHaVipRequestTag {
	s.Value = &v
	return s
}

func (s *CreateHaVipRequestTag) Validate() error {
	return dara.Validate(s)
}
