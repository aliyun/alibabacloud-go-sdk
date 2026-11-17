// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateNetworkAclRequest
	GetClientToken() *string
	SetDescription(v string) *CreateNetworkAclRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateNetworkAclRequest
	GetDryRun() *bool
	SetNetworkAclName(v string) *CreateNetworkAclRequest
	GetNetworkAclName() *string
	SetOwnerAccount(v string) *CreateNetworkAclRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateNetworkAclRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateNetworkAclRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateNetworkAclRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateNetworkAclRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateNetworkAclRequestTag) *CreateNetworkAclRequest
	GetTag() []*CreateNetworkAclRequestTag
	SetVpcId(v string) *CreateNetworkAclRequest
	GetVpcId() *string
}

type CreateNetworkAclRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the network ACL.
	//
	// The description must be 1 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is my NetworkAcl.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to precheck the request. Valid values:
	//
	// 	- **true**: prechecks the request.
	//
	// 	- **false*	- (default): sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the network ACL.
	//
	// The name must be 1 to 128 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// acl-1
	NetworkAclName *string `json:"NetworkAclName,omitempty" xml:"NetworkAclName,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the network ACL.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the resource.
	Tag []*CreateNetworkAclRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC to which the network ACL belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-dsfd34356vdf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateNetworkAclRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclRequest) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateNetworkAclRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateNetworkAclRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateNetworkAclRequest) GetNetworkAclName() *string {
	return s.NetworkAclName
}

func (s *CreateNetworkAclRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateNetworkAclRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateNetworkAclRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNetworkAclRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateNetworkAclRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateNetworkAclRequest) GetTag() []*CreateNetworkAclRequestTag {
	return s.Tag
}

func (s *CreateNetworkAclRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateNetworkAclRequest) SetClientToken(v string) *CreateNetworkAclRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateNetworkAclRequest) SetDescription(v string) *CreateNetworkAclRequest {
	s.Description = &v
	return s
}

func (s *CreateNetworkAclRequest) SetDryRun(v bool) *CreateNetworkAclRequest {
	s.DryRun = &v
	return s
}

func (s *CreateNetworkAclRequest) SetNetworkAclName(v string) *CreateNetworkAclRequest {
	s.NetworkAclName = &v
	return s
}

func (s *CreateNetworkAclRequest) SetOwnerAccount(v string) *CreateNetworkAclRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateNetworkAclRequest) SetOwnerId(v int64) *CreateNetworkAclRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateNetworkAclRequest) SetRegionId(v string) *CreateNetworkAclRequest {
	s.RegionId = &v
	return s
}

func (s *CreateNetworkAclRequest) SetResourceOwnerAccount(v string) *CreateNetworkAclRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateNetworkAclRequest) SetResourceOwnerId(v int64) *CreateNetworkAclRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateNetworkAclRequest) SetTag(v []*CreateNetworkAclRequestTag) *CreateNetworkAclRequest {
	s.Tag = v
	return s
}

func (s *CreateNetworkAclRequest) SetVpcId(v string) *CreateNetworkAclRequest {
	s.VpcId = &v
	return s
}

func (s *CreateNetworkAclRequest) Validate() error {
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

type CreateNetworkAclRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. If you specify this parameter, you cannot specify an empty string.
	//
	// A tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. If you specify this parameter, you can specify an empty string.
	//
	// A tag value can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateNetworkAclRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkAclRequestTag) GoString() string {
	return s.String()
}

func (s *CreateNetworkAclRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateNetworkAclRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateNetworkAclRequestTag) SetKey(v string) *CreateNetworkAclRequestTag {
	s.Key = &v
	return s
}

func (s *CreateNetworkAclRequestTag) SetValue(v string) *CreateNetworkAclRequestTag {
	s.Value = &v
	return s
}

func (s *CreateNetworkAclRequestTag) Validate() error {
	return dara.Validate(s)
}
