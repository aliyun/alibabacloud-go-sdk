// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVscRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVscRequest
	GetClientToken() *string
	SetDescription(v string) *CreateVscRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateVscRequest
	GetDryRun() *bool
	SetInstanceId(v string) *CreateVscRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *CreateVscRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVscRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateVscRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVscRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateVscRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVscRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateVscRequestTag) *CreateVscRequest
	GetTag() []*CreateVscRequestTag
	SetVscName(v string) *CreateVscRequest
	GetVscName() *string
	SetVscType(v string) *CreateVscRequest
	GetVscType() *string
}

type CreateVscRequest struct {
	// Ensures the idempotence of the request. You can generate a parameter value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the virtual storage channel.
	//
	// example:
	//
	// The description of the device
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run on the request. Valid values:
	//
	// - true: sends a check request and does not query the resource status. The check items include whether the AccessKey pair is valid, whether the Resource Access Management (RAM) user has the required authorization, and whether the required parameters are specified. If the check fails, the corresponding error is returned. If the check succeeds, the DryRunOperation error code is returned.
	//
	// - false: sends a normal request. If the check succeeds, the 2XX HTTP status code is returned and the resource status is queried.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp**c262
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the resource belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
	Tag []*CreateVscRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The name of the virtual storage channel.
	//
	// example:
	//
	// Device 1
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// Specifies the VSC type. Valid values: Primary and Standard. Default value: Primary.
	//
	// example:
	//
	// Primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s CreateVscRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVscRequest) GoString() string {
	return s.String()
}

func (s *CreateVscRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVscRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateVscRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVscRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateVscRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVscRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVscRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVscRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVscRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVscRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVscRequest) GetTag() []*CreateVscRequestTag {
	return s.Tag
}

func (s *CreateVscRequest) GetVscName() *string {
	return s.VscName
}

func (s *CreateVscRequest) GetVscType() *string {
	return s.VscType
}

func (s *CreateVscRequest) SetClientToken(v string) *CreateVscRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVscRequest) SetDescription(v string) *CreateVscRequest {
	s.Description = &v
	return s
}

func (s *CreateVscRequest) SetDryRun(v bool) *CreateVscRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVscRequest) SetInstanceId(v string) *CreateVscRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateVscRequest) SetOwnerAccount(v string) *CreateVscRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVscRequest) SetOwnerId(v int64) *CreateVscRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVscRequest) SetRegionId(v string) *CreateVscRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVscRequest) SetResourceGroupId(v string) *CreateVscRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVscRequest) SetResourceOwnerAccount(v string) *CreateVscRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVscRequest) SetResourceOwnerId(v int64) *CreateVscRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVscRequest) SetTag(v []*CreateVscRequestTag) *CreateVscRequest {
	s.Tag = v
	return s
}

func (s *CreateVscRequest) SetVscName(v string) *CreateVscRequest {
	s.VscName = &v
	return s
}

func (s *CreateVscRequest) SetVscType(v string) *CreateVscRequest {
	s.VscType = &v
	return s
}

func (s *CreateVscRequest) Validate() error {
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

type CreateVscRequestTag struct {
	// The tag key of the instance, cloud disk, and primary ENI created by using this template version. Valid values of N: 1 to 20. If you specify this parameter, it cannot be an empty string. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the command. Valid values of N: 1 to 20. The value can be an empty string.
	//
	// The value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVscRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVscRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVscRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVscRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVscRequestTag) SetKey(v string) *CreateVscRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVscRequestTag) SetValue(v string) *CreateVscRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVscRequestTag) Validate() error {
	return dara.Validate(s)
}
