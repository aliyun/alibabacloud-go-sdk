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
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
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
	// The maximum number of times that the activation code can be used to register managed instances. Valid values: 1 to 1000.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The default prefix of instance names. The prefix must be 2 to 50 characters in length. It must start with a letter and cannot start with a special character or digit. It can contain only periods (.), underscores (_), hyphens (-), and colons (:) as special characters. It cannot start with `http://` or `https://`.
	//
	// Instances registered with the activation code created by calling this operation use this name as a prefix to generate sequential instance names. You can also specify a new instance name to override this default value when you register a managed instance.
	//
	// When you register a managed instance, if the InstanceName value is specified, instance names in the format of `<InstanceName>-001` are generated. The number of digits in the sequential number `001` depends on the number of digits in the InstanceCount value. If the InstanceName value is not specified, the hostname of the host is used as the instance name.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP addresses of hosts that are allowed to use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
	//
	// example:
	//
	// 0.0.0.0/0
	IpAddressRange *string `json:"IpAddressRange,omitempty" xml:"IpAddressRange,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Currently supported regions: China North 1 (Qingdao), China North 2 (Beijing), China North 3 (Zhangjiakou), China North 5 (Hohhot), China North 6 (Ulanqab), China East 1 (Hangzhou), China East 2 (Shanghai), China South 1 (Shenzhen), China South 2 (Heyuan), China South 3 (Guangzhou), China Southwest 1 (Chengdu), China (Hong Kong), Singapore, Japan (Tokyo), US (Silicon Valley), US (Virginia).
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to view region IDs and other information.
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
	// The tags.
	Tag []*CreateActivationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The validity period of the activation code. After the validity period expires, the activation code cannot be used to register new instances. Unit: hours. Valid values: 1 to 4.
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
	// The key of the tag for the managed instance activation code. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If you use a single tag to filter resources, the number of resources that are returned under that tag cannot exceed 1,000. If you use multiple tags to filter resources, the number of resources that are bound with all specified tags cannot exceed 1,000. If the number of resources exceeds 1,000, you must call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation to query the resources.
	//
	// The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag for the managed instance activation code. Valid values of N: 1 to 20. The tag value can be an empty string.
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
