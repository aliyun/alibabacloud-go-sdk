// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateTransitRouterShrinkRequest
	GetCenId() *string
	SetClientToken(v string) *CreateTransitRouterShrinkRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterShrinkRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateTransitRouterShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTransitRouterShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterShrinkRequest
	GetResourceOwnerId() *int64
	SetSupportMulticast(v bool) *CreateTransitRouterShrinkRequest
	GetSupportMulticast() *bool
	SetTag(v []*CreateTransitRouterShrinkRequestTag) *CreateTransitRouterShrinkRequest
	GetTag() []*CreateTransitRouterShrinkRequestTag
	SetTransitRouterCidrListShrink(v string) *CreateTransitRouterShrinkRequest
	GetTransitRouterCidrListShrink() *string
	SetTransitRouterDescription(v string) *CreateTransitRouterShrinkRequest
	GetTransitRouterDescription() *string
	SetTransitRouterName(v string) *CreateTransitRouterShrinkRequest
	GetTransitRouterName() *string
}

type CreateTransitRouterShrinkRequest struct {
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
	Tag []*CreateTransitRouterShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The list of transit router CIDR blocks.
	TransitRouterCidrListShrink *string `json:"TransitRouterCidrList,omitempty" xml:"TransitRouterCidrList,omitempty"`
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

func (s CreateTransitRouterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterShrinkRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateTransitRouterShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterShrinkRequest) GetSupportMulticast() *bool {
	return s.SupportMulticast
}

func (s *CreateTransitRouterShrinkRequest) GetTag() []*CreateTransitRouterShrinkRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterShrinkRequest) GetTransitRouterCidrListShrink() *string {
	return s.TransitRouterCidrListShrink
}

func (s *CreateTransitRouterShrinkRequest) GetTransitRouterDescription() *string {
	return s.TransitRouterDescription
}

func (s *CreateTransitRouterShrinkRequest) GetTransitRouterName() *string {
	return s.TransitRouterName
}

func (s *CreateTransitRouterShrinkRequest) SetCenId(v string) *CreateTransitRouterShrinkRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetClientToken(v string) *CreateTransitRouterShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetDryRun(v bool) *CreateTransitRouterShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetOwnerAccount(v string) *CreateTransitRouterShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetOwnerId(v int64) *CreateTransitRouterShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetRegionId(v string) *CreateTransitRouterShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetResourceOwnerId(v int64) *CreateTransitRouterShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetSupportMulticast(v bool) *CreateTransitRouterShrinkRequest {
	s.SupportMulticast = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetTag(v []*CreateTransitRouterShrinkRequestTag) *CreateTransitRouterShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetTransitRouterCidrListShrink(v string) *CreateTransitRouterShrinkRequest {
	s.TransitRouterCidrListShrink = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetTransitRouterDescription(v string) *CreateTransitRouterShrinkRequest {
	s.TransitRouterDescription = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) SetTransitRouterName(v string) *CreateTransitRouterShrinkRequest {
	s.TransitRouterName = &v
	return s
}

func (s *CreateTransitRouterShrinkRequest) Validate() error {
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

type CreateTransitRouterShrinkRequestTag struct {
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

func (s CreateTransitRouterShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterShrinkRequestTag) SetKey(v string) *CreateTransitRouterShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterShrinkRequestTag) SetValue(v string) *CreateTransitRouterShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
