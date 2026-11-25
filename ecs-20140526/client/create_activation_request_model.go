// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActivationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateActivationRequest
	GetClientToken() *string
	SetDescription(v string) *CreateActivationRequest
	GetDescription() *string
	SetInstanceCount(v int32) *CreateActivationRequest
	GetInstanceCount() *int32
	SetInstanceName(v string) *CreateActivationRequest
	GetInstanceName() *string
	SetIpAddressRange(v string) *CreateActivationRequest
	GetIpAddressRange() *string
	SetOwnerAccount(v string) *CreateActivationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateActivationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateActivationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateActivationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateActivationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateActivationRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateActivationRequestTag) *CreateActivationRequest
	GetTag() []*CreateActivationRequestTag
	SetTimeToLiveInHours(v int64) *CreateActivationRequest
	GetTimeToLiveInHours() *int64
}

type CreateActivationRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The ClientToken parameter can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the activation code. The description must be 1 to 100 characters in length.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of times the activation code can be used to register managed instances. Valid values: 1 to 1000.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The default instance name prefix. The prefix must be 2 to 50 characters in length. It must start with a letter and cannot start with a special character or a digit. It can contain only periods (.), underscores (_), hyphens (-), and colons (:). It cannot start with `http://` or `https://`.
	//
	// Instances registered by using the activation code created by this operation will use this name as a prefix to generate sequential instance names. You can also specify a new instance name to override this default value when you register the managed instance.
	//
	// When you register a managed instance, if the InstanceName parameter is specified, the instance name is generated in the <InstanceName>-001 format, where the number of digits in 001 depends on the number of digits in the InstanceCount value. If the InstanceName parameter is not specified, the hostname of the host is used as the instance name.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP addresses of the hosts that are allowed to use the activation code. The value can be an IPv4 address, an IPv6 address, or a CIDR block.
	//
	// example:
	//
	// 0.0.0.0/0
	IpAddressRange *string `json:"IpAddressRange,omitempty" xml:"IpAddressRange,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Supported regions: China (Qingdao), China (Beijing), China (Zhangjiakou), China (Hohhot), China (Ulanqab), China (Hangzhou), China (Shanghai), China (Shenzhen), China (Heyuan), China (Guangzhou), China (Chengdu), Hong Kong (China), Singapore, Japan (Tokyo), US (Silicon Valley), and US (Virginia).
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the activation code belongs.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
	Tag []*CreateActivationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The validity period of the activation code. After the activation code expires, it cannot be used to register new instances. Unit: hours. Valid values: 1 to 4.
	//
	// Default value: 4.
	//
	// example:
	//
	// 4
	TimeToLiveInHours *int64 `json:"TimeToLiveInHours,omitempty" xml:"TimeToLiveInHours,omitempty"`
}

func (s CreateActivationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateActivationRequest) GoString() string {
	return s.String()
}

func (s *CreateActivationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateActivationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateActivationRequest) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *CreateActivationRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateActivationRequest) GetIpAddressRange() *string {
	return s.IpAddressRange
}

func (s *CreateActivationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateActivationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateActivationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateActivationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateActivationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateActivationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateActivationRequest) GetTag() []*CreateActivationRequestTag {
	return s.Tag
}

func (s *CreateActivationRequest) GetTimeToLiveInHours() *int64 {
	return s.TimeToLiveInHours
}

func (s *CreateActivationRequest) SetClientToken(v string) *CreateActivationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateActivationRequest) SetDescription(v string) *CreateActivationRequest {
	s.Description = &v
	return s
}

func (s *CreateActivationRequest) SetInstanceCount(v int32) *CreateActivationRequest {
	s.InstanceCount = &v
	return s
}

func (s *CreateActivationRequest) SetInstanceName(v string) *CreateActivationRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateActivationRequest) SetIpAddressRange(v string) *CreateActivationRequest {
	s.IpAddressRange = &v
	return s
}

func (s *CreateActivationRequest) SetOwnerAccount(v string) *CreateActivationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateActivationRequest) SetOwnerId(v int64) *CreateActivationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateActivationRequest) SetRegionId(v string) *CreateActivationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateActivationRequest) SetResourceGroupId(v string) *CreateActivationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateActivationRequest) SetResourceOwnerAccount(v string) *CreateActivationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateActivationRequest) SetResourceOwnerId(v int64) *CreateActivationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateActivationRequest) SetTag(v []*CreateActivationRequestTag) *CreateActivationRequest {
	s.Tag = v
	return s
}

func (s *CreateActivationRequest) SetTimeToLiveInHours(v int64) *CreateActivationRequest {
	s.TimeToLiveInHours = &v
	return s
}

func (s *CreateActivationRequest) Validate() error {
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

type CreateActivationRequestTag struct {
	// The tag key of the managed instance activation code (dynamic code). N ranges from 1 to 20. If you specify this parameter, it cannot be an empty string.
	//
	// If you use a single tag to filter resources, the resource count returned cannot exceed 1,000. If you use multiple tags to filter resources, the resource count of resources to which you attach multiple tags cannot exceed 1,000. If the resource count exceeds 1,000, you must call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query the resources.
	//
	// The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the managed instance activation code. N ranges from 1 to 20. This parameter can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateActivationRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateActivationRequestTag) GoString() string {
	return s.String()
}

func (s *CreateActivationRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateActivationRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateActivationRequestTag) SetKey(v string) *CreateActivationRequestTag {
	s.Key = &v
	return s
}

func (s *CreateActivationRequestTag) SetValue(v string) *CreateActivationRequestTag {
	s.Value = &v
	return s
}

func (s *CreateActivationRequestTag) Validate() error {
	return dara.Validate(s)
}
