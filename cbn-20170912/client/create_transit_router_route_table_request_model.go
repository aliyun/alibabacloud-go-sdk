// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterRouteTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateTransitRouterRouteTableRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterRouteTableRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateTransitRouterRouteTableRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterRouteTableRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateTransitRouterRouteTableRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterRouteTableRequest
	GetResourceOwnerId() *int64
	SetRouteTableOptions(v *CreateTransitRouterRouteTableRequestRouteTableOptions) *CreateTransitRouterRouteTableRequest
	GetRouteTableOptions() *CreateTransitRouterRouteTableRequestRouteTableOptions
	SetTag(v []*CreateTransitRouterRouteTableRequestTag) *CreateTransitRouterRouteTableRequest
	GetTag() []*CreateTransitRouterRouteTableRequestTag
	SetTransitRouterId(v string) *CreateTransitRouterRouteTableRequest
	GetTransitRouterId() *string
	SetTransitRouterRouteTableDescription(v string) *CreateTransitRouterRouteTableRequest
	GetTransitRouterRouteTableDescription() *string
	SetTransitRouterRouteTableName(v string) *CreateTransitRouterRouteTableRequest
	GetTransitRouterRouteTableName() *string
}

type CreateTransitRouterRouteTableRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to precheck the request. Check items include permissions and the status of the specified cloud resources. Valid values:
	//
	// 	- **false*	- (default): sends the request. If the request passes the precheck, the custom route table is created.
	//
	// 	- **true**: prechecks the request but does not create the custom route table. If you use this value, the system checks the required parameters and the request syntax. If the request fails to pass the precheck, an error message is returned. If the request passes the check, the system returns the ID of the request.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The features of the route table.
	RouteTableOptions *CreateTransitRouterRouteTableRequestRouteTableOptions `json:"RouteTableOptions,omitempty" xml:"RouteTableOptions,omitempty" type:"Struct"`
	// The tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*CreateTransitRouterRouteTableRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the Enterprise Edition transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The description of the custom route table.
	//
	// The description must be 1 to 256 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testdesc
	TransitRouterRouteTableDescription *string `json:"TransitRouterRouteTableDescription,omitempty" xml:"TransitRouterRouteTableDescription,omitempty"`
	// The name of the custom route table.
	//
	// The name must be 1 to 128 characters in length, and cannot start with http:// or https://. You can also leave this parameter empty.
	//
	// example:
	//
	// testname
	TransitRouterRouteTableName *string `json:"TransitRouterRouteTableName,omitempty" xml:"TransitRouterRouteTableName,omitempty"`
}

func (s CreateTransitRouterRouteTableRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRouteTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterRouteTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterRouteTableRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterRouteTableRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterRouteTableRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterRouteTableRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterRouteTableRequest) GetRouteTableOptions() *CreateTransitRouterRouteTableRequestRouteTableOptions {
	return s.RouteTableOptions
}

func (s *CreateTransitRouterRouteTableRequest) GetTag() []*CreateTransitRouterRouteTableRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterRouteTableRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterRouteTableRequest) GetTransitRouterRouteTableDescription() *string {
	return s.TransitRouterRouteTableDescription
}

func (s *CreateTransitRouterRouteTableRequest) GetTransitRouterRouteTableName() *string {
	return s.TransitRouterRouteTableName
}

func (s *CreateTransitRouterRouteTableRequest) SetClientToken(v string) *CreateTransitRouterRouteTableRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetDryRun(v bool) *CreateTransitRouterRouteTableRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetOwnerAccount(v string) *CreateTransitRouterRouteTableRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetOwnerId(v int64) *CreateTransitRouterRouteTableRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterRouteTableRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetResourceOwnerId(v int64) *CreateTransitRouterRouteTableRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetRouteTableOptions(v *CreateTransitRouterRouteTableRequestRouteTableOptions) *CreateTransitRouterRouteTableRequest {
	s.RouteTableOptions = v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetTag(v []*CreateTransitRouterRouteTableRequestTag) *CreateTransitRouterRouteTableRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetTransitRouterId(v string) *CreateTransitRouterRouteTableRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetTransitRouterRouteTableDescription(v string) *CreateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableDescription = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) SetTransitRouterRouteTableName(v string) *CreateTransitRouterRouteTableRequest {
	s.TransitRouterRouteTableName = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequest) Validate() error {
	if s.RouteTableOptions != nil {
		if err := s.RouteTableOptions.Validate(); err != nil {
			return err
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

type CreateTransitRouterRouteTableRequestRouteTableOptions struct {
	// Specifies whether to enable multi-region equal-cost multi-path (ECMP) routing. Valid values:
	//
	// 	- **disable**(default) If multi-region ECMP routing is disabled, routes that are learned from different regions but have the same prefix and attributes select the transit router with the smallest region ID as the next hop. Region IDs are sorted in alphabetic order. The network latency and bandwidth consumption also vary based on the region. Proceed with caution.
	//
	// 	- **enable*	- If multi-region ECMP routing is enabled, routes that are learned from different regions but have the same prefix and attributes form an ECMP route. The network latency and bandwidth consumption also vary based on the region. Proceed with caution.
	//
	// example:
	//
	// disable
	MultiRegionECMP *string `json:"MultiRegionECMP,omitempty" xml:"MultiRegionECMP,omitempty"`
}

func (s CreateTransitRouterRouteTableRequestRouteTableOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRouteTableRequestRouteTableOptions) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteTableRequestRouteTableOptions) GetMultiRegionECMP() *string {
	return s.MultiRegionECMP
}

func (s *CreateTransitRouterRouteTableRequestRouteTableOptions) SetMultiRegionECMP(v string) *CreateTransitRouterRouteTableRequestRouteTableOptions {
	s.MultiRegionECMP = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequestRouteTableOptions) Validate() error {
	return dara.Validate(s)
}

type CreateTransitRouterRouteTableRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// You can specify at most 20 tag keys.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// The tag value can be 0 to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must have a unique tag value. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTransitRouterRouteTableRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterRouteTableRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterRouteTableRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterRouteTableRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterRouteTableRequestTag) SetKey(v string) *CreateTransitRouterRouteTableRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequestTag) SetValue(v string) *CreateTransitRouterRouteTableRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterRouteTableRequestTag) Validate() error {
	return dara.Validate(s)
}
