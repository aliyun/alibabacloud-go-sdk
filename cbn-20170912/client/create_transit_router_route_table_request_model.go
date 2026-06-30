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
	// Generate a token from your client to ensure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the request as the **ClientToken**. The **RequestId*	- of each request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run. The system checks the request for potential issues, including required parameters, request format, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the request. If the request passes the dry run, a custom route table is created.
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
	// The tag.
	//
	// You can specify up to 20 tags in each call.
	Tag []*CreateTransitRouterRouteTableRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the Enterprise Edition transit router instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The description of the custom route table.
	//
	// The description can be empty or 1 to 256 characters in length. It cannot start with \\`http\\://\\` or \\`https\\://\\`.
	//
	// example:
	//
	// testdesc
	TransitRouterRouteTableDescription *string `json:"TransitRouterRouteTableDescription,omitempty" xml:"TransitRouterRouteTableDescription,omitempty"`
	// The name of the custom route table.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with \\`http\\://\\` or \\`https\\://\\`.
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
	// The multi-region equal-cost multi-path (ECMP) routing feature. Valid values:
	//
	// - **disable*	- (default): disables the multi-region ECMP routing feature. If you disable the multi-region ECMP routing feature, routes that are learned from different regions but have the same prefix and attributes select the transit router with the smallest region ID as the next hop. The region ID is sorted in alphabetical order. In this case, the latency and bandwidth consumption of the traffic may change. Make sure that you are aware of the impact before you disable the feature.
	//
	// - **enable**: enables the multi-region ECMP routing feature. If you enable the multi-region ECMP routing feature, routes that are learned from different regions but have the same prefix and attributes are considered ECMP routes. In this case, the latency and bandwidth consumption of the traffic may change. Make sure that you are aware of the impact before you enable the feature.
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
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with \\`aliyun\\` or \\`acs:\\`. It cannot contain \\`http\\://\\` or \\`https\\://\\`.
	//
	// You can specify up to 20 tag keys.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value can be an empty string or a string of up to 128 characters. It cannot start with \\`aliyun\\` or \\`acs:\\` and cannot contain \\`http\\://\\` or \\`https\\://\\`.
	//
	// Each tag key must have a unique tag value. You can specify up to 20 tag values.
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
