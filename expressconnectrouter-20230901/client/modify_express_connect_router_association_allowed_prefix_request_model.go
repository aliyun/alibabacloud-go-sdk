// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterAssociationAllowedPrefixRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedPrefixes(v []*string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest
	GetAllowedPrefixes() []*string
	SetAllowedPrefixesMode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest
	GetAllowedPrefixesMode() *string
	SetAssociationId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest
	GetAssociationId() *string
	SetClientToken(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest
	GetDryRun() *bool
	SetEcrId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest
	GetEcrId() *string
	SetOwnerAccount(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest
	GetOwnerAccount() *string
	SetVersion(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest
	GetVersion() *string
}

type ModifyExpressConnectRouterAssociationAllowedPrefixRequest struct {
	// The list of allowed route prefixes.
	AllowedPrefixes []*string `json:"AllowedPrefixes,omitempty" xml:"AllowedPrefixes,omitempty" type:"Repeated"`
	// The prefix routing pattern. Valid values:
	//
	// - MatchMode: Publishes only the configured route prefixes to your data center. Existing specific routes are withdrawn.
	//
	// - IncrementalMode: Publishes the configured route prefixes to your data center. Withdraw previously published more specific routes that are covered by the configured prefixes. More specific routes that are not covered by the configured prefixes remain published.
	//
	// example:
	//
	// MatchMode
	AllowedPrefixesMode *string `json:"AllowedPrefixesMode,omitempty" xml:"AllowedPrefixesMode,omitempty"`
	// The ID of the association between the ECR and the VPC or TR.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// A client token used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to make sure that the value is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: Performs a dry run.
	//
	// - **false*	- (default): Sends a normal request.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of ECR.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId        *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	Version      *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GetAllowedPrefixes() []*string {
	return s.AllowedPrefixes
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GetAllowedPrefixesMode() *string {
	return s.AllowedPrefixesMode
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) GetVersion() *string {
	return s.Version
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetAllowedPrefixes(v []*string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.AllowedPrefixes = v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetAllowedPrefixesMode(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.AllowedPrefixesMode = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetAssociationId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.AssociationId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetClientToken(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetDryRun(v bool) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetEcrId(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetOwnerAccount(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) SetVersion(v string) *ModifyExpressConnectRouterAssociationAllowedPrefixRequest {
	s.Version = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationAllowedPrefixRequest) Validate() error {
	return dara.Validate(s)
}
