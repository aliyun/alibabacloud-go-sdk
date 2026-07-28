// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVpcPrefixListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVpcPrefixListRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateVpcPrefixListRequest
	GetDryRun() *bool
	SetIpVersion(v string) *CreateVpcPrefixListRequest
	GetIpVersion() *string
	SetMaxEntries(v int32) *CreateVpcPrefixListRequest
	GetMaxEntries() *int32
	SetOwnerAccount(v string) *CreateVpcPrefixListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateVpcPrefixListRequest
	GetOwnerId() *int64
	SetPrefixListDescription(v string) *CreateVpcPrefixListRequest
	GetPrefixListDescription() *string
	SetPrefixListEntries(v []*CreateVpcPrefixListRequestPrefixListEntries) *CreateVpcPrefixListRequest
	GetPrefixListEntries() []*CreateVpcPrefixListRequestPrefixListEntries
	SetPrefixListName(v string) *CreateVpcPrefixListRequest
	GetPrefixListName() *string
	SetRegionId(v string) *CreateVpcPrefixListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateVpcPrefixListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateVpcPrefixListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVpcPrefixListRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateVpcPrefixListRequestTag) *CreateVpcPrefixListRequest
	GetTag() []*CreateVpcPrefixListRequestTag
}

type CreateVpcPrefixListRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The ClientToken value can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without creating the prefix list. The system checks the required parameters, request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the check succeeds, an HTTP 2xx status code is returned and the prefix list is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IP version. Valid values:
	//
	// - **IPv4*	- (default): IPv4.
	//
	// - **IPv6**: IPv6.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The maximum number of entries for Classless Inter-Domain Routing (CIDR) blocks in the prefix list. Default value: 50.
	//
	// example:
	//
	// 50
	MaxEntries   *int32  `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the prefix list.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// description
	PrefixListDescription *string `json:"PrefixListDescription,omitempty" xml:"PrefixListDescription,omitempty"`
	// The Classless Inter-Domain Routing (CIDR) block information of the prefix list.
	PrefixListEntries []*CreateVpcPrefixListRequestPrefixListEntries `json:"PrefixListEntries,omitempty" xml:"PrefixListEntries,omitempty" type:"Repeated"`
	// The name of the prefix list.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// name
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	// The region ID of the prefix list that you want to create.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the prefix list belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*CreateVpcPrefixListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateVpcPrefixListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcPrefixListRequest) GoString() string {
	return s.String()
}

func (s *CreateVpcPrefixListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVpcPrefixListRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateVpcPrefixListRequest) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateVpcPrefixListRequest) GetMaxEntries() *int32 {
	return s.MaxEntries
}

func (s *CreateVpcPrefixListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateVpcPrefixListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVpcPrefixListRequest) GetPrefixListDescription() *string {
	return s.PrefixListDescription
}

func (s *CreateVpcPrefixListRequest) GetPrefixListEntries() []*CreateVpcPrefixListRequestPrefixListEntries {
	return s.PrefixListEntries
}

func (s *CreateVpcPrefixListRequest) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *CreateVpcPrefixListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVpcPrefixListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVpcPrefixListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVpcPrefixListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVpcPrefixListRequest) GetTag() []*CreateVpcPrefixListRequestTag {
	return s.Tag
}

func (s *CreateVpcPrefixListRequest) SetClientToken(v string) *CreateVpcPrefixListRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetDryRun(v bool) *CreateVpcPrefixListRequest {
	s.DryRun = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetIpVersion(v string) *CreateVpcPrefixListRequest {
	s.IpVersion = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetMaxEntries(v int32) *CreateVpcPrefixListRequest {
	s.MaxEntries = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetOwnerAccount(v string) *CreateVpcPrefixListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetOwnerId(v int64) *CreateVpcPrefixListRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetPrefixListDescription(v string) *CreateVpcPrefixListRequest {
	s.PrefixListDescription = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetPrefixListEntries(v []*CreateVpcPrefixListRequestPrefixListEntries) *CreateVpcPrefixListRequest {
	s.PrefixListEntries = v
	return s
}

func (s *CreateVpcPrefixListRequest) SetPrefixListName(v string) *CreateVpcPrefixListRequest {
	s.PrefixListName = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetRegionId(v string) *CreateVpcPrefixListRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetResourceGroupId(v string) *CreateVpcPrefixListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetResourceOwnerAccount(v string) *CreateVpcPrefixListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetResourceOwnerId(v int64) *CreateVpcPrefixListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVpcPrefixListRequest) SetTag(v []*CreateVpcPrefixListRequestTag) *CreateVpcPrefixListRequest {
	s.Tag = v
	return s
}

func (s *CreateVpcPrefixListRequest) Validate() error {
	if s.PrefixListEntries != nil {
		for _, item := range s.PrefixListEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateVpcPrefixListRequestPrefixListEntries struct {
	// The Classless Inter-Domain Routing (CIDR) block of the prefix list.
	//
	// example:
	//
	// 192.168.0.0/16
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description of the Classless Inter-Domain Routing (CIDR) block in the prefix list.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// CIDR
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s CreateVpcPrefixListRequestPrefixListEntries) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcPrefixListRequestPrefixListEntries) GoString() string {
	return s.String()
}

func (s *CreateVpcPrefixListRequestPrefixListEntries) GetCidr() *string {
	return s.Cidr
}

func (s *CreateVpcPrefixListRequestPrefixListEntries) GetDescription() *string {
	return s.Description
}

func (s *CreateVpcPrefixListRequestPrefixListEntries) SetCidr(v string) *CreateVpcPrefixListRequestPrefixListEntries {
	s.Cidr = &v
	return s
}

func (s *CreateVpcPrefixListRequestPrefixListEntries) SetDescription(v string) *CreateVpcPrefixListRequestPrefixListEntries {
	s.Description = &v
	return s
}

func (s *CreateVpcPrefixListRequestPrefixListEntries) Validate() error {
	return dara.Validate(s)
}

type CreateVpcPrefixListRequestTag struct {
	// The tag key of the resource. You can specify up to 20 tag keys. The tag key cannot be an empty string.
	//
	// The tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. You can specify up to 20 tag values. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVpcPrefixListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVpcPrefixListRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVpcPrefixListRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVpcPrefixListRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVpcPrefixListRequestTag) SetKey(v string) *CreateVpcPrefixListRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVpcPrefixListRequestTag) SetValue(v string) *CreateVpcPrefixListRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVpcPrefixListRequestTag) Validate() error {
	return dara.Validate(s)
}
