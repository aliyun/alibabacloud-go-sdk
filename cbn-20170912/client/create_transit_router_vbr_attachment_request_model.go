// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterVbrAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterVbrAttachmentRequest
	GetAutoPublishRouteEnabled() *bool
	SetCenId(v string) *CreateTransitRouterVbrAttachmentRequest
	GetCenId() *string
	SetClientToken(v string) *CreateTransitRouterVbrAttachmentRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterVbrAttachmentRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *CreateTransitRouterVbrAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterVbrAttachmentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTransitRouterVbrAttachmentRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterVbrAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterVbrAttachmentRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateTransitRouterVbrAttachmentRequestTag) *CreateTransitRouterVbrAttachmentRequest
	GetTag() []*CreateTransitRouterVbrAttachmentRequestTag
	SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVbrAttachmentRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentName(v string) *CreateTransitRouterVbrAttachmentRequest
	GetTransitRouterAttachmentName() *string
	SetTransitRouterId(v string) *CreateTransitRouterVbrAttachmentRequest
	GetTransitRouterId() *string
	SetVbrId(v string) *CreateTransitRouterVbrAttachmentRequest
	GetVbrId() *string
	SetVbrOwnerId(v int64) *CreateTransitRouterVbrAttachmentRequest
	GetVbrOwnerId() *int64
}

type CreateTransitRouterVbrAttachmentRequest struct {
	// Specifies whether to enable the Enterprise Edition transit router instance that can forward traffic to automatically publish route entry and routing information to the VBR instance.
	//
	// Valid values:
	//
	// - **false*	- (default): no.
	//
	// - **true**: yes.
	//
	// example:
	//
	// false
	AutoPublishRouteEnabled *bool `json:"AutoPublishRouteEnabled,omitempty" xml:"AutoPublishRouteEnabled,omitempty"`
	// The Cloud Enterprise Network (CEN) instance ID. This parameter is required when TransitRouterId is not specified. The system performs automatic creation of an Enterprise Edition transit router instance that can forward and perform routing in the region specified by RegionId.
	//
	// example:
	//
	// cen-j3jzhw1zpau2km****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request, including permission and instance status checks. Valid values:
	//
	// - **false*	- (default): sends a normal request and creates the VBR connection after the check passes.
	//
	// - **true**: sends a check request and only performs the check without creating the VBR connection. The check items include whether required parameters are specified and the request format. If the check fails, an error is returned. If the check passes, the request ID is returned.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the VBR. This parameter is required and must be used in combination with VbrId and CenId or TransitRouterId.
	//
	// You can call [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) to query region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tag information.
	//
	// You can specify up to 20 tags at a time.
	Tag []*CreateTransitRouterVbrAttachmentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The description of the VBR connection.
	//
	// The description can be empty or 1 to 256 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The name of the VBR connection.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// testname
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The Enterprise Edition transit router instance ID. This parameter is required when CenId is not specified. It is used to create a VBR connection on an existing Enterprise Edition transit router instance that can forward and perform routing.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The VBR instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vbr-bp1svadp4lq38janc****
	VbrId *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	// The Alibaba Cloud account ID to which the VBR instance belongs. The default value is the Alibaba Cloud account ID of the current logon.
	//
	// > This parameter is required if you want to load a cross-account network instance.
	//
	// example:
	//
	// 1250123456123456
	VbrOwnerId *int64 `json:"VbrOwnerId,omitempty" xml:"VbrOwnerId,omitempty"`
}

func (s CreateTransitRouterVbrAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVbrAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetAutoPublishRouteEnabled() *bool {
	return s.AutoPublishRouteEnabled
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetTag() []*CreateTransitRouterVbrAttachmentRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetVbrId() *string {
	return s.VbrId
}

func (s *CreateTransitRouterVbrAttachmentRequest) GetVbrOwnerId() *int64 {
	return s.VbrOwnerId
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetAutoPublishRouteEnabled(v bool) *CreateTransitRouterVbrAttachmentRequest {
	s.AutoPublishRouteEnabled = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetCenId(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetClientToken(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetDryRun(v bool) *CreateTransitRouterVbrAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetOwnerAccount(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetOwnerId(v int64) *CreateTransitRouterVbrAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetRegionId(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetResourceOwnerId(v int64) *CreateTransitRouterVbrAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetTag(v []*CreateTransitRouterVbrAttachmentRequestTag) *CreateTransitRouterVbrAttachmentRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetTransitRouterAttachmentName(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetTransitRouterId(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetVbrId(v string) *CreateTransitRouterVbrAttachmentRequest {
	s.VbrId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) SetVbrOwnerId(v int64) *CreateTransitRouterVbrAttachmentRequest {
	s.VbrOwnerId = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequest) Validate() error {
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

type CreateTransitRouterVbrAttachmentRequestTag struct {
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. It can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`, or contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value cannot be empty. It can be up to 128 characters in length, and cannot start with aliyun or acs:, or contain http:// or https://.
	//
	// Each tag key corresponds to a tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTransitRouterVbrAttachmentRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterVbrAttachmentRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterVbrAttachmentRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterVbrAttachmentRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterVbrAttachmentRequestTag) SetKey(v string) *CreateTransitRouterVbrAttachmentRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequestTag) SetValue(v string) *CreateTransitRouterVbrAttachmentRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterVbrAttachmentRequestTag) Validate() error {
	return dara.Validate(s)
}
