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
	// The Cloud Enterprise Network (CEN) instance ID.
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
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to execute a dry run, without performing the actual request. The dry run includes permission verification, instance status verification, and forwarding and routing checks. Valid values:
	//
	// - **false*	- (default): sends a normal request and creates the Enterprise Edition transit router instance after the request passes the check.
	//
	// - **true**: sends a check request, without creating the Enterprise Edition transit router instance. The system checks the required parameters, request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the error code `DryRunOperation` is returned.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Enterprise Edition transit router instance.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the region ID.
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
	// - **false*	- (default): disables the multicast feature.
	//
	// - **true**: enables the multicast feature.
	//
	// Only Enterprise Edition transit routers in some regions support the multicast feature. You can call the [ListTransitRouterAvailableResource](https://help.aliyun.com/document_detail/261356.html) operation to query the regions that support the multicast feature.
	//
	// example:
	//
	// false
	SupportMulticast *bool `json:"SupportMulticast,omitempty" xml:"SupportMulticast,omitempty"`
	// The tag information.
	Tag []*CreateTransitRouterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The list of transit router CIDR blocks.
	TransitRouterCidrList []*CreateTransitRouterRequestTransitRouterCidrList `json:"TransitRouterCidrList,omitempty" xml:"TransitRouterCidrList,omitempty" type:"Repeated"`
	// The description of the Enterprise Edition transit router instance.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterDescription *string `json:"TransitRouterDescription,omitempty" xml:"TransitRouterDescription,omitempty"`
	// The name of the Enterprise Edition transit router instance.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TransitRouterCidrList != nil {
		for _, item := range s.TransitRouterCidrList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateTransitRouterRequestTag struct {
	// The tag key of the resource.
	//
	// Once specified, the tag key cannot be an empty string. The tag key can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// Once specified, the tag value cannot be empty. The tag value can be up to 128 characters in length, and cannot start with aliyun or acs:, or contain http:// or https://.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
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
	// The transit router CIDR block.
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
	// Specifies whether to allow the system to automatically add a route for the transit router CIDR block to the transit router route table.
	//
	// - **true*	- (default): allows the system.
	//
	//
	//
	//      If you select true, after you create a VPN connection of the private gateway type and create a route learning relationship for the VPN connection, the system automatically adds the following route entry to the transit router route table that has a route learning relationship with the VPN connection:
	//
	//     A blackhole route whose destination CIDR block is the transit router CIDR block from which gateway IP addresses are allocated for the IPsec connection.
	//
	//      The blackhole route is propagated only to the route tables of virtual border router (VBR) instances connected to the transit router.
	//
	// - **false**: does not allow the system.
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
