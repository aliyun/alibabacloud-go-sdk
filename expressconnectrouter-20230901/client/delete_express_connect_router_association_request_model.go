// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectRouterAssociationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *DeleteExpressConnectRouterAssociationRequest
	GetAssociationId() *string
	SetClientToken(v string) *DeleteExpressConnectRouterAssociationRequest
	GetClientToken() *string
	SetDeleteAttachment(v bool) *DeleteExpressConnectRouterAssociationRequest
	GetDeleteAttachment() *bool
	SetDryRun(v bool) *DeleteExpressConnectRouterAssociationRequest
	GetDryRun() *bool
	SetEcrId(v string) *DeleteExpressConnectRouterAssociationRequest
	GetEcrId() *string
	SetVersion(v string) *DeleteExpressConnectRouterAssociationRequest
	GetVersion() *string
}

type DeleteExpressConnectRouterAssociationRequest struct {
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
	// You can generate a parameter value from your client to ensure that the parameter value is unique among different requests. The ClientToken parameter supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to delete the association between the Express Connect Router (ECR) and a transit router or VPC. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	DeleteAttachment *bool `json:"DeleteAttachment,omitempty" xml:"DeleteAttachment,omitempty"`
	// Specifies whether to only perform a dry run for this request. Valid values:
	//
	// - **true**: Sends a dry run request and does not modify service resources.
	//
	// - **false*	- (default value): Sends a normal request.
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
	EcrId   *string `json:"EcrId,omitempty" xml:"EcrId,omitempty"`
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DeleteExpressConnectRouterAssociationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectRouterAssociationRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectRouterAssociationRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *DeleteExpressConnectRouterAssociationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteExpressConnectRouterAssociationRequest) GetDeleteAttachment() *bool {
	return s.DeleteAttachment
}

func (s *DeleteExpressConnectRouterAssociationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteExpressConnectRouterAssociationRequest) GetEcrId() *string {
	return s.EcrId
}

func (s *DeleteExpressConnectRouterAssociationRequest) GetVersion() *string {
	return s.Version
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetAssociationId(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.AssociationId = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetClientToken(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetDeleteAttachment(v bool) *DeleteExpressConnectRouterAssociationRequest {
	s.DeleteAttachment = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetDryRun(v bool) *DeleteExpressConnectRouterAssociationRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetEcrId(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.EcrId = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) SetVersion(v string) *DeleteExpressConnectRouterAssociationRequest {
	s.Version = &v
	return s
}

func (s *DeleteExpressConnectRouterAssociationRequest) Validate() error {
	return dara.Validate(s)
}
