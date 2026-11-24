// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectRouterAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *ModifyExpressConnectRouterAssociationRequest
	GetAssociationId() *string
	SetClientToken(v string) *ModifyExpressConnectRouterAssociationRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyExpressConnectRouterAssociationRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyExpressConnectRouterAssociationRequest
	GetDryRun() *bool
	SetEcrId(v string) *ModifyExpressConnectRouterAssociationRequest
	GetEcrId() *string
	SetVersion(v string) *ModifyExpressConnectRouterAssociationRequest
	GetVersion() *string
}

type ModifyExpressConnectRouterAssociationRequest struct {
	// The ID used to associate the Express Connect Router (ECR) with a VPC or transit router.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the parameter value from your client to ensure that the value is unique among different requests. The client token supports only ASCII characters.
	//
	// >
	//
	// > - If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the child instance. The description must be 0 to 128 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this request. Valid values:
	//
	// - true: Sends a dry run request without modifying any service resources.
	//
	// - false (default): Sends a normal request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the Express Connect Router (ECR).
	//
	// This parameter is required.
	//
	// example:
	//
	// ecr-mezk2idmsd0vx2****
	EcrId   *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ModifyExpressConnectRouterAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectRouterAssociationRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *ModifyExpressConnectRouterAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyExpressConnectRouterAssociationRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyExpressConnectRouterAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyExpressConnectRouterAssociationRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *ModifyExpressConnectRouterAssociationRequest) GetVersion() *string {
	return s.Version
}

func (s *ModifyExpressConnectRouterAssociationRequest) SetAssociationId(v string) *ModifyExpressConnectRouterAssociationRequest {
	s.AssociationId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationRequest) SetClientToken(v string) *ModifyExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationRequest) SetDescription(v string) *ModifyExpressConnectRouterAssociationRequest {
	s.Description = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationRequest) SetDryRun(v bool) *ModifyExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationRequest) SetEcrId(v string) *ModifyExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationRequest) SetVersion(v string) *ModifyExpressConnectRouterAssociationRequest {
	s.Version = &v
	return s
}

func (s *ModifyExpressConnectRouterAssociationRequest) Validate() error {
	return dara.Validate(s)
}
