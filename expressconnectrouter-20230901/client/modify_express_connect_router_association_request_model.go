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
	// This parameter is required.
	//
	// example:
	//
	// ecr-assoc-9p2qxx5phpca2m****
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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
