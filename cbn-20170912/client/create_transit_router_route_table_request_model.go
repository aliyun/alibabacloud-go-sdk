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
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, including permission and instance status verification. Valid values:
	//
	// - **false*	- (default): sends a normal request and creates the custom route table after the request passes the verification.
	//
	// - **true**: sends a check request without creating the custom route table. The system checks the required parameters, request format, and other conditions. If the check fails, the corresponding error is returned. If the check passes, the corresponding request ID is returned.
	//
	// example:
	//
	// false
	DryRun               *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The route table options.
	RouteTableOptions *CreateTransitRouterRouteTableRequestRouteTableOptions `json:"RouteTableOptions,omitempty" xml:"RouteTableOptions,omitempty" type:"Struct"`
	// The tag information.
	//
	// You can specify up to 20 tags at a time.
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
	// The description can be empty or 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterRouteTableDescription *string `json:"TransitRouterRouteTableDescription,omitempty" xml:"TransitRouterRouteTableDescription,omitempty"`
	// The name of the custom route table.
	//
	// The name can be empty or 1 to 128 characters in length and cannot start with http:// or https://.
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
	// The multi-region equal-cost multi-path (ECMP) routing setting. Valid values:
	//
	// - **disable*	- (default): disables multi-region ECMP routing. After multi-region ECMP routing is disabled, routes with the same prefix learned from different regions prefer the transit router whose Region ID is the smallest (sorted alphabetically) as the next hop when other route attributes are the same. This may change traffic latency and bandwidth consumption between regions. Make sure that you fully evaluate the impact before disabling this feature.
	//
	// - **enable**: enables multi-region ECMP routing. After multi-region ECMP routing is enabled, routes with the same prefix learned from different regions form ECMP routes when other route attributes are the same. This may change traffic latency and bandwidth consumption between regions. Make sure that you fully evaluate the impact before enabling this feature.
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
	// Once specified, the tag key cannot be an empty string. The tag key can be up to 64 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// test
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// Once specified, the tag value cannot be empty. The tag value can be up to 128 characters in length and cannot start with aliyun or acs:. It cannot contain http:// or https://.
	//
	// Each tag key corresponds to one tag value. You can specify up to 20 tag values at a time.
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
