// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateTransitRouterRequest
	GetCenId() *string
	SetClientToken(v string) *CreateTransitRouterRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateTransitRouterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTransitRouterRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterRequest
	GetResourceOwnerId() *int64
	SetSupportMulticast(v bool) *CreateTransitRouterRequest
	GetSupportMulticast() *bool
	SetTag(v []*CreateTransitRouterRequestTag) *CreateTransitRouterRequest
	GetTag() []*CreateTransitRouterRequestTag
	SetTransitRouterCidrList(v []*CreateTransitRouterRequestTransitRouterCidrList) *CreateTransitRouterRequest
	GetTransitRouterCidrList() []*CreateTransitRouterRequestTransitRouterCidrList
	SetTransitRouterDescription(v string) *CreateTransitRouterRequest
	GetTransitRouterDescription() *string
	SetTransitRouterName(v string) *CreateTransitRouterRequest
	GetTransitRouterName() *string
}

type CreateTransitRouterRequest struct {
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system automatically uses **RequestId*	- as **ClientToken**. The value of **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Check items include permissions and the status of the specified cloud resources. Valid values:
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, an Enterprise Edition transit router is created.
	//
	// 	- **true**: checks the request but does not create the Enterprise Edition transit router. If you use this value, the system checks whether the required parameters are set, and whether the request syntax is valid. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Enterprise Edition transit router is deployed.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-zhangjiakou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable the multicast feature for the Enterprise Edition transit router. Valid values:
	//
	// 	- **false*	- (default): no
	//
	// 	- **true**: yes
	//
	// The multicast feature is supported only in specific regions. You can call [ListTransitRouterAvailableResource](https://help.aliyun.com/document_detail/261356.html) to query the regions that support multicast.
	//
	// example:
	//
	// false
	SupportMulticast *bool `json:"SupportMulticast,omitempty" xml:"SupportMulticast,omitempty"`
	// The tags.
	Tag []*CreateTransitRouterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The CIDR blocks to be added to the transit router.
	TransitRouterCidrList []*CreateTransitRouterRequestTransitRouterCidrList `json:"TransitRouterCidrList,omitempty" xml:"TransitRouterCidrList,omitempty" type:"Repeated"`
	// The description of the Enterprise Edition transit router instance.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testdesc
	TransitRouterDescription *string `json:"TransitRouterDescription,omitempty" xml:"TransitRouterDescription,omitempty"`
	// The name of the Enterprise Edition transit router.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testname
	TransitRouterName *string `json:"TransitRouterName,omitempty" xml:"TransitRouterName,omitempty"`
}

func (s CreateTransitRouterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateTransitRouterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterRequest) GetSupportMulticast() *bool {
	return s.SupportMulticast
}

func (s *CreateTransitRouterRequest) GetTag() []*CreateTransitRouterRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterRequest) GetTransitRouterCidrList() []*CreateTransitRouterRequestTransitRouterCidrList {
	return s.TransitRouterCidrList
}

func (s *CreateTransitRouterRequest) GetTransitRouterDescription() *string {
	return s.TransitRouterDescription
}

func (s *CreateTransitRouterRequest) GetTransitRouterName() *string {
	return s.TransitRouterName
}

func (s *CreateTransitRouterRequest) SetCenId(v string) *CreateTransitRouterRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterRequest) SetClientToken(v string) *CreateTransitRouterRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterRequest) SetDryRun(v bool) *CreateTransitRouterRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterRequest) SetOwnerAccount(v string) *CreateTransitRouterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRequest) SetOwnerId(v int64) *CreateTransitRouterRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterRequest) SetRegionId(v string) *CreateTransitRouterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRequest) SetResourceOwnerId(v int64) *CreateTransitRouterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterRequest) SetSupportMulticast(v bool) *CreateTransitRouterRequest {
	s.SupportMulticast = &v
	return s
}

func (s *CreateTransitRouterRequest) SetTag(v []*CreateTransitRouterRequestTag) *CreateTransitRouterRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterRequest) SetTransitRouterCidrList(v []*CreateTransitRouterRequestTransitRouterCidrList) *CreateTransitRouterRequest {
	s.TransitRouterCidrList = v
	return s
}

func (s *CreateTransitRouterRequest) SetTransitRouterDescription(v string) *CreateTransitRouterRequest {
	s.TransitRouterDescription = &v
	return s
}

func (s *CreateTransitRouterRequest) SetTransitRouterName(v string) *CreateTransitRouterRequest {
	s.TransitRouterName = &v
	return s
}

func (s *CreateTransitRouterRequest) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterRequestTag struct {
	// The tag key.
	//
	// The tag keys cannot be an empty string. The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys in each call.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be an empty string or up to 128 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each key-value must be unique. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTransitRouterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterRequestTag) SetKey(v string) *CreateTransitRouterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterRequestTag) SetValue(v string) *CreateTransitRouterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterRequestTransitRouterCidrList struct {
	// The CIDR block of the transit router.
	//
	// example:
	//
	// 192.168.10.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description of the transit router CIDR block.
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the transit router CIDR block.
	//
	// The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to allow the system to automatically add a route that points to the CIDR block to the route table of the transit router.
	//
	// 	- **true*	- (default)
	//
	//     If you set the value to true, after you create a VPN attachment on a private VPN gateway and enable route learning for the VPN attachment, the system automatically adds the following route to the route table of the transit router that is in route learning relationship with the VPN attachment:
	//
	//     A blackhole route whose destination CIDR block is the transit router CIDR block, which refers to the CIDR block from which gateway IP addresses are allocated to the IPsec-VPN connection.
	//
	//     The blackhole route is advertised only to the route tables of virtual border routers (VBRs) connected to the transit router.
	//
	// 	- **false**
	//
	// example:
	//
	// true
	PublishCidrRoute *bool `json:"PublishCidrRoute,omitempty" xml:"PublishCidrRoute,omitempty"`
}

func (s CreateTransitRouterRequestTransitRouterCidrList) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRequestTransitRouterCidrList) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRequestTransitRouterCidrList) GetCidr() *string {
	return s.Cidr
}

func (s *CreateTransitRouterRequestTransitRouterCidrList) GetDescription() *string {
	return s.Description
}

func (s *CreateTransitRouterRequestTransitRouterCidrList) GetName() *string {
	return s.Name
}

func (s *CreateTransitRouterRequestTransitRouterCidrList) GetPublishCidrRoute() *bool {
	return s.PublishCidrRoute
}

func (s *CreateTransitRouterRequestTransitRouterCidrList) SetCidr(v string) *CreateTransitRouterRequestTransitRouterCidrList {
	s.Cidr = &v
	return s
}

func (s *CreateTransitRouterRequestTransitRouterCidrList) SetDescription(v string) *CreateTransitRouterRequestTransitRouterCidrList {
	s.Description = &v
	return s
}

func (s *CreateTransitRouterRequestTransitRouterCidrList) SetName(v string) *CreateTransitRouterRequestTransitRouterCidrList {
	s.Name = &v
	return s
}

func (s *CreateTransitRouterRequestTransitRouterCidrList) SetPublishCidrRoute(v bool) *CreateTransitRouterRequestTransitRouterCidrList {
	s.PublishCidrRoute = &v
	return s
}

func (s *CreateTransitRouterRequestTransitRouterCidrList) Validate() error {
	return dara.Validate(s)
}
