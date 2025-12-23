// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamResourceDiscoveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateIpamResourceDiscoveryRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateIpamResourceDiscoveryRequest
	GetDryRun() *bool
	SetIpamResourceDiscoveryDescription(v string) *CreateIpamResourceDiscoveryRequest
	GetIpamResourceDiscoveryDescription() *string
	SetIpamResourceDiscoveryName(v string) *CreateIpamResourceDiscoveryRequest
	GetIpamResourceDiscoveryName() *string
	SetOperatingRegionList(v []*string) *CreateIpamResourceDiscoveryRequest
	GetOperatingRegionList() []*string
	SetOwnerAccount(v string) *CreateIpamResourceDiscoveryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateIpamResourceDiscoveryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateIpamResourceDiscoveryRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateIpamResourceDiscoveryRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateIpamResourceDiscoveryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateIpamResourceDiscoveryRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateIpamResourceDiscoveryRequestTag) *CreateIpamResourceDiscoveryRequest
	GetTag() []*CreateIpamResourceDiscoveryRequestTag
}

type CreateIpamResourceDiscoveryRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without sending the actual request. Valid value:
	//
	// 	- **true**: Performs the dry run without creating a custom resource discovery instance. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- **false*	- (default): Performs a dry run and the actual request. If the request passes the dry run, an HTTP 2xx status code is returned and a custom resource discovery instance is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The description of resource discovery.
	//
	// example:
	//
	// description
	IpamResourceDiscoveryDescription *string `json:"IpamResourceDiscoveryDescription,omitempty" xml:"IpamResourceDiscoveryDescription,omitempty"`
	// The name of the resource discovery.
	//
	// example:
	//
	// name
	IpamResourceDiscoveryName *string `json:"IpamResourceDiscoveryName,omitempty" xml:"IpamResourceDiscoveryName,omitempty"`
	// The list of effective regions.
	//
	// This parameter is required.
	OperatingRegionList []*string `json:"OperatingRegionList,omitempty" xml:"OperatingRegionList,omitempty" type:"Repeated"`
	OwnerAccount        *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The request region.
	//
	// >  The request region is the hosted region of the resource discovery instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2sermdd6****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information.
	Tag []*CreateIpamResourceDiscoveryRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateIpamResourceDiscoveryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamResourceDiscoveryRequest) GoString() string {
	return s.String()
}

func (s *CreateIpamResourceDiscoveryRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateIpamResourceDiscoveryRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateIpamResourceDiscoveryRequest) GetIpamResourceDiscoveryDescription() *string {
	return s.IpamResourceDiscoveryDescription
}

func (s *CreateIpamResourceDiscoveryRequest) GetIpamResourceDiscoveryName() *string {
	return s.IpamResourceDiscoveryName
}

func (s *CreateIpamResourceDiscoveryRequest) GetOperatingRegionList() []*string {
	return s.OperatingRegionList
}

func (s *CreateIpamResourceDiscoveryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateIpamResourceDiscoveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateIpamResourceDiscoveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateIpamResourceDiscoveryRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateIpamResourceDiscoveryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateIpamResourceDiscoveryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateIpamResourceDiscoveryRequest) GetTag() []*CreateIpamResourceDiscoveryRequestTag {
	return s.Tag
}

func (s *CreateIpamResourceDiscoveryRequest) SetClientToken(v string) *CreateIpamResourceDiscoveryRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetDryRun(v bool) *CreateIpamResourceDiscoveryRequest {
	s.DryRun = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryDescription(v string) *CreateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryDescription = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetIpamResourceDiscoveryName(v string) *CreateIpamResourceDiscoveryRequest {
	s.IpamResourceDiscoveryName = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetOperatingRegionList(v []*string) *CreateIpamResourceDiscoveryRequest {
	s.OperatingRegionList = v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetOwnerAccount(v string) *CreateIpamResourceDiscoveryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetOwnerId(v int64) *CreateIpamResourceDiscoveryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetRegionId(v string) *CreateIpamResourceDiscoveryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetResourceGroupId(v string) *CreateIpamResourceDiscoveryRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetResourceOwnerAccount(v string) *CreateIpamResourceDiscoveryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetResourceOwnerId(v int64) *CreateIpamResourceDiscoveryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) SetTag(v []*CreateIpamResourceDiscoveryRequestTag) *CreateIpamResourceDiscoveryRequest {
	s.Tag = v
	return s
}

func (s *CreateIpamResourceDiscoveryRequest) Validate() error {
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

type CreateIpamResourceDiscoveryRequestTag struct {
	// The tag keys. You can specify at most 20 tag keys. It cannot be an empty string.
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

func (s CreateIpamResourceDiscoveryRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamResourceDiscoveryRequestTag) GoString() string {
	return s.String()
}

func (s *CreateIpamResourceDiscoveryRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateIpamResourceDiscoveryRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateIpamResourceDiscoveryRequestTag) SetKey(v string) *CreateIpamResourceDiscoveryRequestTag {
	s.Key = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequestTag) SetValue(v string) *CreateIpamResourceDiscoveryRequestTag {
	s.Value = &v
	return s
}

func (s *CreateIpamResourceDiscoveryRequestTag) Validate() error {
	return dara.Validate(s)
}
