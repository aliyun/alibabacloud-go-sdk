// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectRouterAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedPrefixes(v []*string) *CreateExpressConnectRouterAssociationRequest
	GetAllowedPrefixes() []*string
	SetAllowedPrefixesMode(v string) *CreateExpressConnectRouterAssociationRequest
	GetAllowedPrefixesMode() *string
	SetAssociationRegionId(v string) *CreateExpressConnectRouterAssociationRequest
	GetAssociationRegionId() *string
	SetCenId(v string) *CreateExpressConnectRouterAssociationRequest
	GetCenId() *string
	SetClientToken(v string) *CreateExpressConnectRouterAssociationRequest
	GetClientToken() *string
	SetCreateAttachment(v bool) *CreateExpressConnectRouterAssociationRequest
	GetCreateAttachment() *bool
	SetDescription(v string) *CreateExpressConnectRouterAssociationRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateExpressConnectRouterAssociationRequest
	GetDryRun() *bool
	SetEcrId(v string) *CreateExpressConnectRouterAssociationRequest
	GetEcrId() *string
	SetTransitRouterId(v string) *CreateExpressConnectRouterAssociationRequest
	GetTransitRouterId() *string
	SetTransitRouterOwnerId(v int64) *CreateExpressConnectRouterAssociationRequest
	GetTransitRouterOwnerId() *int64
	SetVersion(v string) *CreateExpressConnectRouterAssociationRequest
	GetVersion() *string
	SetVpcId(v string) *CreateExpressConnectRouterAssociationRequest
	GetVpcId() *string
	SetVpcOwnerId(v int64) *CreateExpressConnectRouterAssociationRequest
	GetVpcOwnerId() *int64
}

type CreateExpressConnectRouterAssociationRequest struct {
	// The allowed route prefixes.
	AllowedPrefixes []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	// The route prefix mode. Valid values:
	//
	// - **MatchMode**: After you distribute new route CIDR blocks to data centers, original specific routes that are distributed are withdrawn.
	//
	// - **IncrementalMode**: After you distribute new route CIDR blocks to data centers, the original specific routes that fall in the CIDR blocks that you configure are withdrawn, and the original specific routes that do not fall in the CIDR blocks are still distributed.
	//
	// example:
	//
	// MatchMode
	AllowedPrefixesMode *string `json:"AllowedPrefixesMode,omitempty" xml:"AllowedPrefixesMode,omitempty"`
	// The region ID of the resource to be associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	AssociationRegionId *string `json:"AssociationRegionId,omitempty" xml:"AssociationRegionId,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// example:
	//
	// cen-of3o1the3i4gwb****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to initiate an association on the TR when the ECR is associated with the TR. Valid values:
	//
	// 	- **true**: You do not need to initiate an association on the TR.
	//
	// 	- **false**: You need to initiate an association on the TR.
	//
	// example:
	//
	// true
	CreateAttachment *bool `json:"CreateAttachment,omitempty" xml:"CreateAttachment,omitempty"`
	// The information about the associated resource. It must be 0 to 128 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ECR ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The TR ID.
	//
	// example:
	//
	// tr-2ze4i71c6be454e2l****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the TR. Default value: ID of the Alibaba Cloud account that logs in.
	//
	// >  If you want to connect to a network instance that belongs to a different account, this parameter is required.
	//
	// example:
	//
	// 189159362009****
	TransitRouterOwnerId *int64  `json:"TransitRouterOwnerId,omitempty" xml:"TransitRouterOwnerId,omitempty"`
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-bp1h37fchc6jmfyln****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the VPC. Default value: ID of the Alibaba Cloud account that logs in.
	//
	// >  If you want to connect to a network instance that belongs to a different account, this parameter is required.
	//
	// example:
	//
	// 132193271328****
	VpcOwnerId *int64 `json:"VpcOwnerId,omitempty" xml:"VpcOwnerId,omitempty"`
}

func (s CreateExpressConnectRouterAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationRequest) GetAllowedPrefixes() []*string {
	return s.AllowedPrefixes
}

func (s *CreateExpressConnectRouterAssociationRequest) GetAllowedPrefixesMode() *string {
	return s.AllowedPrefixesMode
}

func (s *CreateExpressConnectRouterAssociationRequest) GetAssociationRegionId() *string {
	return s.AssociationRegionId
}

func (s *CreateExpressConnectRouterAssociationRequest) GetCenId() *string {
	return s.CenId
}

func (s *CreateExpressConnectRouterAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateExpressConnectRouterAssociationRequest) GetCreateAttachment() *bool {
	return s.CreateAttachment
}

func (s *CreateExpressConnectRouterAssociationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateExpressConnectRouterAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateExpressConnectRouterAssociationRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *CreateExpressConnectRouterAssociationRequest) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *CreateExpressConnectRouterAssociationRequest) GetTransitRouterOwnerId() *int64 {
	return s.TransitRouterOwnerId
}

func (s *CreateExpressConnectRouterAssociationRequest) GetVersion() *string {
	return s.Version
}

func (s *CreateExpressConnectRouterAssociationRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateExpressConnectRouterAssociationRequest) GetVpcOwnerId() *int64 {
	return s.VpcOwnerId
}

func (s *CreateExpressConnectRouterAssociationRequest) SetAllowedPrefixes(v []*string) *CreateExpressConnectRouterAssociationRequest {
	s.AllowedPrefixes = v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetAllowedPrefixesMode(v string) *CreateExpressConnectRouterAssociationRequest {
	s.AllowedPrefixesMode = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetAssociationRegionId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.AssociationRegionId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetCenId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.CenId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetClientToken(v string) *CreateExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetCreateAttachment(v bool) *CreateExpressConnectRouterAssociationRequest {
	s.CreateAttachment = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetDescription(v string) *CreateExpressConnectRouterAssociationRequest {
	s.Description = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetDryRun(v bool) *CreateExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetEcrId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetTransitRouterId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.TransitRouterId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetTransitRouterOwnerId(v int64) *CreateExpressConnectRouterAssociationRequest {
	s.TransitRouterOwnerId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetVersion(v string) *CreateExpressConnectRouterAssociationRequest {
	s.Version = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetVpcId(v string) *CreateExpressConnectRouterAssociationRequest {
	s.VpcId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) SetVpcOwnerId(v int64) *CreateExpressConnectRouterAssociationRequest {
	s.VpcOwnerId = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequest) Validate() error {
	return dara.Validate(s)
}
