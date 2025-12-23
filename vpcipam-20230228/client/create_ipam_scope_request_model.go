// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamScopeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateIpamScopeRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateIpamScopeRequest
	GetDryRun() *bool
	SetIpamId(v string) *CreateIpamScopeRequest
	GetIpamId() *string
	SetIpamScopeDescription(v string) *CreateIpamScopeRequest
	GetIpamScopeDescription() *string
	SetIpamScopeName(v string) *CreateIpamScopeRequest
	GetIpamScopeName() *string
	SetIpamScopeType(v string) *CreateIpamScopeRequest
	GetIpamScopeType() *string
	SetOwnerAccount(v string) *CreateIpamScopeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIpamScopeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIpamScopeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateIpamScopeRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateIpamScopeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIpamScopeRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateIpamScopeRequestTag) *CreateIpamScopeRequest
	GetTag() []*CreateIpamScopeRequestTag
}

type CreateIpamScopeRequest struct {
	// The client token used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized RAM users, and missing parameter values. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and the actual request. After the request passes the dry run, a 2xx HTTP status code is returned and the IPAM scope is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the IPAM.
	//
	// This parameter is required.
	//
	// example:
	//
	// ipam-ccxbnsbhew0d6t****
	IpamId *string `json:"IpamId,omitempty" xml:"IpamId,omitempty"`
	// The description of the IPAM scope.
	//
	// It must be 1 to 256 characters in length. It must start with a letter but cannot start with `http://` or `https://`. This parameter is empty by default.
	//
	// example:
	//
	// test description
	IpamScopeDescription *string `json:"IpamScopeDescription,omitempty" xml:"IpamScopeDescription,omitempty"`
	// The name of the IPAM scope.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	IpamScopeName *string `json:"IpamScopeName,omitempty" xml:"IpamScopeName,omitempty"`
	// The type of IPAM scope: **private**
	//
	//
	// **Usage notes*	- You can create only private IPAM scopes.
	//
	// example:
	//
	// private
	IpamScopeType *string `json:"IpamScopeType,omitempty" xml:"IpamScopeType,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the IPAM scope.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tag []*CreateIpamScopeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateIpamScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamScopeRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamScopeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpamScopeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateIpamScopeRequest) GetIpamId() *string {
	return s.IpamId
}

func (s *CreateIpamScopeRequest) GetIpamScopeDescription() *string {
	return s.IpamScopeDescription
}

func (s *CreateIpamScopeRequest) GetIpamScopeName() *string {
	return s.IpamScopeName
}

func (s *CreateIpamScopeRequest) GetIpamScopeType() *string {
	return s.IpamScopeType
}

func (s *CreateIpamScopeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIpamScopeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIpamScopeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpamScopeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateIpamScopeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIpamScopeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIpamScopeRequest) GetTag() []*CreateIpamScopeRequestTag {
	return s.Tag
}

func (s *CreateIpamScopeRequest) SetClientToken(v string) *CreateIpamScopeRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamScopeRequest) SetDryRun(v bool) *CreateIpamScopeRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamScopeRequest) SetIpamId(v string) *CreateIpamScopeRequest {
	s.IpamId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetIpamScopeDescription(v string) *CreateIpamScopeRequest {
	s.IpamScopeDescription = &v
	return s
}

func (s *CreateIpamScopeRequest) SetIpamScopeName(v string) *CreateIpamScopeRequest {
	s.IpamScopeName = &v
	return s
}

func (s *CreateIpamScopeRequest) SetIpamScopeType(v string) *CreateIpamScopeRequest {
	s.IpamScopeType = &v
	return s
}

func (s *CreateIpamScopeRequest) SetOwnerAccount(v string) *CreateIpamScopeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpamScopeRequest) SetOwnerId(v int64) *CreateIpamScopeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetRegionId(v string) *CreateIpamScopeRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetResourceGroupId(v string) *CreateIpamScopeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetResourceOwnerAccount(v string) *CreateIpamScopeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpamScopeRequest) SetResourceOwnerId(v int64) *CreateIpamScopeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpamScopeRequest) SetTag(v []*CreateIpamScopeRequestTag) *CreateIpamScopeRequest {
	s.Tag = v
	return s
}

func (s *CreateIpamScopeRequest) Validate() error {
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

type CreateIpamScopeRequestTag struct {
	// The tag key of the resource. You can specify at most 20 tag keys. It cannot be an empty string.
	//
	// The tag key can be up to 64 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). It must start with a letter but cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. You can specify empty strings as tag values.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateIpamScopeRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamScopeRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpamScopeRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateIpamScopeRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateIpamScopeRequestTag) SetKey(v string) *CreateIpamScopeRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpamScopeRequestTag) SetValue(v string) *CreateIpamScopeRequestTag {
	s.Value = &v
	return s
}

func (s *CreateIpamScopeRequestTag) Validate() error {
	return dara.Validate(s)
}
