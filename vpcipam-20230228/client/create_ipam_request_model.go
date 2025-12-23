// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateIpamRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateIpamRequest
	GetDryRun() *bool
	SetIpamDescription(v string) *CreateIpamRequest
	GetIpamDescription() *string
	SetIpamName(v string) *CreateIpamRequest
	GetIpamName() *string
	SetOperatingRegionList(v []*string) *CreateIpamRequest
	GetOperatingRegionList() []*string
	SetOwnerAccount(v string) *CreateIpamRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIpamRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIpamRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateIpamRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateIpamRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIpamRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateIpamRequestTag) *CreateIpamRequest
	GetTag() []*CreateIpamRequestTag
}

type CreateIpamRequest struct {
	// The client token used to ensure the idempotence of the request. Use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of the IPAM.
	//
	// It must be 1 to 256 characters in length. Start with a letter but cannot start with `http://` or `https://`. If you do not specify a description, the description is empty by default.
	//
	// example:
	//
	// This is my first Ipam
	IpamDescription *string `json:"IpamDescription,omitempty" xml:"IpamDescription,omitempty"`
	// The name of the IPAM.
	//
	// It must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// abc
	IpamName *string `json:"IpamName,omitempty" xml:"IpamName,omitempty"`
	// The effective regions of the IPAM.
	//
	// This parameter is required.
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the IPAM instance is hosted. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID of the IPAM.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag list.
	Tag []*CreateIpamRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateIpamRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpamRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateIpamRequest) GetIpamDescription() *string {
	return s.IpamDescription
}

func (s *CreateIpamRequest) GetIpamName() *string {
	return s.IpamName
}

func (s *CreateIpamRequest) GetOperatingRegionList() []*string {
	return s.OperatingRegionList
}

func (s *CreateIpamRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIpamRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIpamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpamRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateIpamRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIpamRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIpamRequest) GetTag() []*CreateIpamRequestTag {
	return s.Tag
}

func (s *CreateIpamRequest) SetClientToken(v string) *CreateIpamRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamRequest) SetDryRun(v bool) *CreateIpamRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamRequest) SetIpamDescription(v string) *CreateIpamRequest {
	s.IpamDescription = &v
	return s
}

func (s *CreateIpamRequest) SetIpamName(v string) *CreateIpamRequest {
	s.IpamName = &v
	return s
}

func (s *CreateIpamRequest) SetOperatingRegionList(v []*string) *CreateIpamRequest {
	s.OperatingRegionList = v
	return s
}

func (s *CreateIpamRequest) SetOwnerAccount(v string) *CreateIpamRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpamRequest) SetOwnerId(v int64) *CreateIpamRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpamRequest) SetRegionId(v string) *CreateIpamRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpamRequest) SetResourceGroupId(v string) *CreateIpamRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpamRequest) SetResourceOwnerAccount(v string) *CreateIpamRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpamRequest) SetResourceOwnerId(v int64) *CreateIpamRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpamRequest) SetTag(v []*CreateIpamRequestTag) *CreateIpamRequest {
	s.Tag = v
	return s
}

func (s *CreateIpamRequest) Validate() error {
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

type CreateIpamRequestTag struct {
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

func (s CreateIpamRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpamRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateIpamRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateIpamRequestTag) SetKey(v string) *CreateIpamRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpamRequestTag) SetValue(v string) *CreateIpamRequestTag {
	s.Value = &v
	return s
}

func (s *CreateIpamRequestTag) Validate() error {
	return dara.Validate(s)
}
