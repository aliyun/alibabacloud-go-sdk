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
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-7qthudw0ll6jmc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the ECR.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-upyc0viial107r****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The ID of the Alibaba Cloud account to which the ECR belongs. By default, the ID of the current Alibaba Cloud account is specified.
	//
	// >  If you want to connect to a network instance that belongs to a different account, this parameter is required.
	//
	// example:
	//
	// 1250123456123456
	EcrOwnerId   *int64  `json:"EcrOwnerId,omitempty" xml:"EcrOwnerId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the transit router.
	//
	// You can call the [DescribeChildInstanceRegions](https://help.aliyun.com/document_detail/132080.html) operation to query the most recent region list.
	//
	// example:
	//
	// eu-central-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	//
	// You can specify at most 20 tags in each call.
	Tag []*CreateTransitRouterEcrAttachmentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The description of the ECR connection.
	//
	// This parameter is optional. If you enter a description, it must be 1 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// testdesc
	TransitRouterAttachmentDescription *string `json:"TransitRouterAttachmentDescription,omitempty" xml:"TransitRouterAttachmentDescription,omitempty"`
	// The name of the ECR connection.
	//
	// The name can be empty or 1 to 128 characters in length, and cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	TransitRouterAttachmentName *string `json:"TransitRouterAttachmentName,omitempty" xml:"TransitRouterAttachmentName,omitempty"`
	// The ID of the transit router.
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
	return dara.Validate(s)
}

type CreateTransitRouterEcrAttachmentRequestTag struct {
	// The tag key.
	//
	// The tag key cannot be an empty string. The tag key can be up to 64 characters in length. It cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
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
	// Each key-value pair must be unique. You can specify values for at most 20 tag keys in each call.
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
