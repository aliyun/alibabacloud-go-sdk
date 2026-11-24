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
	SetTag(v []*CreateExpressConnectRouterAssociationRequestTag) *CreateExpressConnectRouterAssociationRequest
	GetTag() []*CreateExpressConnectRouterAssociationRequestTag
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
	// The list of allowed route prefixes.
	AllowedPrefixes []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	// The route prefix mode. Valid values:
	//
	// - MatchMode: match mode. Only the configured route CIDR blocks are advertised to the data center, and the previously advertised specific routes are withdrawn.
	//
	// - IncrementalMode: incremental mode. The configured route CIDR blocks are advertised to the data center. The previously advertised specific routes that fall within the configured route CIDR blocks are withdrawn, and the specific routes that do not fall within the configured CIDR blocks remain advertised.
	//
	// example:
	//
	// MatchMode
	AllowedPrefixesMode *string `json:"AllowedPrefixesMode,omitempty" xml:"AllowedPrefixesMode,omitempty"`
	// The region where the associated resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	AssociationRegionId *string `json:"AssociationRegionId,omitempty" xml:"AssociationRegionId,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// > If you specify TransitRouterId, CenId is required. You must first call the CreateCen operation of the Cbn (2017-09-12) service to create a CEN instance and obtain the CEN instance ID.
	//
	// example:
	//
	// cen-of3o1the3i4gwb****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate a parameter value from your client to ensure that the value is unique among different requests. The client token only supports ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to initiate an association when the Express Connect Router is associated with the transit router. Valid values:
	//
	// - **true**: No association is initiated on the transit router side.
	//
	// - **false**: An association is initiated on the transit router side.
	//
	// example:
	//
	// true
	CreateAttachment *bool `json:"CreateAttachment,omitempty" xml:"CreateAttachment,omitempty"`
	// The description of the associated resource. The description must be 0 to 128 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to only perform a dry run for this request. Valid values:
	//
	// - **true**: The system sends a check request without modifying service resources.
	//
	// - **false*	- (default): The system sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	// The list of tags.
	//
	// You can add up to 20 tags.
	Tag []*CreateExpressConnectRouterAssociationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the transit router.
	//
	// > The TransitRouterId and VpcId parameters are mutually exclusive. You must specify one of them. You must first call the CreateTransitRouter operation of the Cbn (2017-09-12) service to create a transit router and obtain the transit router ID.
	//
	// example:
	//
	// tr-2ze4i71c6be454e2l****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
	// The ID of the Alibaba Cloud account to which the transit router instance belongs. The transit router is used to forward routing traffic. Default value: the current logon Alibaba Cloud account ID.
	//
	// > If you want to load a network instance across accounts, this parameter is required.
	//
	// example:
	//
	// 189159362009****
	TransitRouterOwnerId *int64  `json:"TransitRouterOwnerId,omitempty" xml:"TransitRouterOwnerId,omitempty"`
	Version              *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The instance ID of the VPC-connected instance.
	//
	// > The VpcId and TransitRouterId parameters are mutually exclusive. You must specify one of them.
	//
	// > You must first call the CreateVpc operation of the Vpc (2016-04-28) service to create a VPC and obtain the VPC ID.
	//
	// example:
	//
	// vpc-bp1h37fchc6jmfyln****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The Alibaba Cloud account ID to which the VPC-connected instance belongs. Default value: the current logon Alibaba Cloud account ID.
	//
	// > If you want to load a network instance across accounts, this parameter is required.
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

func (s *CreateExpressConnectRouterAssociationRequest) GetTag() []*CreateExpressConnectRouterAssociationRequestTag {
	return s.Tag
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

func (s *CreateExpressConnectRouterAssociationRequest) SetTag(v []*CreateExpressConnectRouterAssociationRequestTag) *CreateExpressConnectRouterAssociationRequest {
	s.Tag = v
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

type CreateExpressConnectRouterAssociationRequestTag struct {
	// The tag key of the resource. If you specify this parameter, it cannot be an empty string. The tag key can be up to 128 characters in length. It cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the resource. The tag value can be an empty string and can be up to 128 characters in length. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateExpressConnectRouterAssociationRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectRouterAssociationRequestTag) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectRouterAssociationRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateExpressConnectRouterAssociationRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateExpressConnectRouterAssociationRequestTag) SetKey(v string) *CreateExpressConnectRouterAssociationRequestTag {
	s.Key = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequestTag) SetValue(v string) *CreateExpressConnectRouterAssociationRequestTag {
	s.Value = &v
	return s
}

func (s *CreateExpressConnectRouterAssociationRequestTag) Validate() error {
	return dara.Validate(s)
}
