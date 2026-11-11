// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateTransitRouterPolicyTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *DisassociateTransitRouterPolicyTableRequest
	GetAssociationId() *string
	SetClientToken(v string) *DisassociateTransitRouterPolicyTableRequest
	GetClientToken() *string
	SetDryRun(v bool) *DisassociateTransitRouterPolicyTableRequest
	GetDryRun() *bool
}

type DisassociateTransitRouterPolicyTableRequest struct {
	// The ID of the association between the transit router policy table and the network instance connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// pta-123asdfqwe
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run before sending the request. The dry run includes permission and instance status checks. Valid values:
	//
	// - **false*	- (default): sends a normal request. After the request passes the dry run, the network instance connection is disassociated from the transit router policy table.
	//
	// - **true**: performs a dry run without sending the request. The system checks whether required parameters are specified and validates the request format. If the request fails the dry run, an error is returned. If the request passes the dry run, the request ID is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
}

func (s DisassociateTransitRouterPolicyTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociateTransitRouterPolicyTableRequest) GoString() string {
	return s.String()
}

func (s *DisassociateTransitRouterPolicyTableRequest) GetAssociationId() *string {
	return s.AssociationId
}

func (s *DisassociateTransitRouterPolicyTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DisassociateTransitRouterPolicyTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DisassociateTransitRouterPolicyTableRequest) SetAssociationId(v string) *DisassociateTransitRouterPolicyTableRequest {
	s.AssociationId = &v
	return s
}

func (s *DisassociateTransitRouterPolicyTableRequest) SetClientToken(v string) *DisassociateTransitRouterPolicyTableRequest {
	s.ClientToken = &v
	return s
}

func (s *DisassociateTransitRouterPolicyTableRequest) SetDryRun(v bool) *DisassociateTransitRouterPolicyTableRequest {
	s.DryRun = &v
	return s
}

func (s *DisassociateTransitRouterPolicyTableRequest) Validate() error {
	return dara.Validate(s)
}
