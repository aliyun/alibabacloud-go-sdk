// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterEcrAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *CreateTransitRouterEcrAttachmentRequest
	GetCenId() *string
	SetClientToken(v string) *CreateTransitRouterEcrAttachmentRequest
	GetClientToken() *string
	SetDryRun(v bool) *CreateTransitRouterEcrAttachmentRequest
	GetDryRun() *bool
	SetEcrId(v string) *CreateTransitRouterEcrAttachmentRequest
	GetEcrId() *string
	SetEcrOwnerId(v int64) *CreateTransitRouterEcrAttachmentRequest
	GetEcrOwnerId() *int64
	SetOwnerAccount(v string) *CreateTransitRouterEcrAttachmentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateTransitRouterEcrAttachmentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateTransitRouterEcrAttachmentRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateTransitRouterEcrAttachmentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTransitRouterEcrAttachmentRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateTransitRouterEcrAttachmentRequestTag) *CreateTransitRouterEcrAttachmentRequest
	GetTag() []*CreateTransitRouterEcrAttachmentRequestTag
	SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterEcrAttachmentRequest
	GetTransitRouterAttachmentDescription() *string
	SetTransitRouterAttachmentName(v string) *CreateTransitRouterEcrAttachmentRequest
	GetTransitRouterAttachmentName() *string
	SetTransitRouterId(v string) *CreateTransitRouterEcrAttachmentRequest
	GetTransitRouterId() *string
}

type CreateTransitRouterEcrAttachmentRequest struct {
	// The Cloud Enterprise Network (CEN) instance ID.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without modifying the configurations of the VPN connection. The system checks whether required parameters are specified, the request format is valid, and the service limits are met. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// - **false*	- (default): sends a normal request. After the request passes the dry run, the configurations of the VPN connection are modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The associated Express Connect Router (ECR) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-upyc0viial107r****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The Alibaba Cloud account ID to which the ECR instance belongs. The default value is the current logon Alibaba Cloud account ID.
	//
	// > If you want to load a cross-account network instance, this parameter is required.
	//
	// example:
	//
	// 1250123456123456
	EcrOwnerId   *int64  `json:"EcrOwnerId,omitempty" xml:"EcrOwnerId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the transit router.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query region IDs.
	//
	// example:
	//
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The list of tags.
	//
	// You can specify up to 20 tags at a time.
	Tag []*CreateTransitRouterEcrAttachmentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The description of the ECR connection.
	//
	// The description can be empty or 1 to 256 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The name of the ECR connection.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The transit router instance ID.
	//
	// example:
	//
	// tr-bp1su1ytdxtataupl****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s CreateTransitRouterEcrAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterEcrAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetEcrOwnerId() *int64 {
	return s.EcrOwnerId
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetTag() []*CreateTransitRouterEcrAttachmentRequestTag {
	return s.Tag
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetTransitRouterAttachmentDescription() *string {
	return s.TransitRouterAttachmentDescription
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetTransitRouterAttachmentName() *string {
	return s.TransitRouterAttachmentName
}

func (s *CreateTransitRouterEcrAttachmentRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetCenId(v string) *CreateTransitRouterEcrAttachmentRequest {
	s.CenId = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetClientToken(v string) *CreateTransitRouterEcrAttachmentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetDryRun(v bool) *CreateTransitRouterEcrAttachmentRequest {
	s.DryRun = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetEcrId(v string) *CreateTransitRouterEcrAttachmentRequest {
	s.EcrId = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetEcrOwnerId(v int64) *CreateTransitRouterEcrAttachmentRequest {
	s.EcrOwnerId = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetOwnerAccount(v string) *CreateTransitRouterEcrAttachmentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetOwnerId(v int64) *CreateTransitRouterEcrAttachmentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetRegionId(v string) *CreateTransitRouterEcrAttachmentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetResourceOwnerAccount(v string) *CreateTransitRouterEcrAttachmentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetResourceOwnerId(v int64) *CreateTransitRouterEcrAttachmentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetTag(v []*CreateTransitRouterEcrAttachmentRequestTag) *CreateTransitRouterEcrAttachmentRequest {
	s.Tag = v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetTransitRouterAttachmentDescription(v string) *CreateTransitRouterEcrAttachmentRequest {
	s.TransitRouterAttachmentDescription = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetTransitRouterAttachmentName(v string) *CreateTransitRouterEcrAttachmentRequest {
	s.TransitRouterAttachmentName = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) SetTransitRouterId(v string) *CreateTransitRouterEcrAttachmentRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequest) Validate() error {
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

type CreateTransitRouterEcrAttachmentRequestTag struct {
	// The tag key of the resource.
	//
	// The tag key cannot be an empty string. It can be up to 64 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// You can specify up to 20 tag keys at a time.
	//
	// example:
	//
	// tagtest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource.
	//
	// The tag value cannot be an empty string. It can be up to 128 characters in length, and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	//
	// Each tag key must be paired with a tag value. You can specify up to 20 tag values at a time.
	//
	// example:
	//
	// tagtest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateTransitRouterEcrAttachmentRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterEcrAttachmentRequestTag) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterEcrAttachmentRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateTransitRouterEcrAttachmentRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateTransitRouterEcrAttachmentRequestTag) SetKey(v string) *CreateTransitRouterEcrAttachmentRequestTag {
	s.Key = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequestTag) SetValue(v string) *CreateTransitRouterEcrAttachmentRequestTag {
	s.Value = &v
	return s
}

func (s *CreateTransitRouterEcrAttachmentRequestTag) Validate() error {
	return dara.Validate(s)
}
