// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTransitRouterPolicyTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachmentId(v string) *AssociateTransitRouterPolicyTableRequest
	GetAttachmentId() *string
	SetClientToken(v string) *AssociateTransitRouterPolicyTableRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateTransitRouterPolicyTableRequest
	GetDryRun() *bool
	SetPolicyTableId(v string) *AssociateTransitRouterPolicyTableRequest
	GetPolicyTableId() *string
}

type AssociateTransitRouterPolicyTableRequest struct {
	// The ID of the network instance connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// tr-attach-2nalp6yksc805w****
	AttachmentId *string `json:"AttachmentId,omitempty" xml:"AttachmentId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client and make sure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. The dry run includes checks on permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a normal request, performs the check, and associates the policy route table if the check passes.
	//
	// - **true**: sends a dry run request, performs only the check, and does not associate the policy route table. The check items include whether required parameters are specified and the request format. If the check fails, an error is returned. If the check passes, the request ID is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The instance ID of the policy routing table.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptbl-abc123xyz
	PolicyTableId *string `json:"PolicyTableId,omitempty" xml:"PolicyTableId,omitempty"`
}

func (s AssociateTransitRouterPolicyTableRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateTransitRouterPolicyTableRequest) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterPolicyTableRequest) GetAttachmentId() *string {
	return s.AttachmentId
}

func (s *AssociateTransitRouterPolicyTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateTransitRouterPolicyTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateTransitRouterPolicyTableRequest) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *AssociateTransitRouterPolicyTableRequest) SetAttachmentId(v string) *AssociateTransitRouterPolicyTableRequest {
	s.AttachmentId = &v
	return s
}

func (s *AssociateTransitRouterPolicyTableRequest) SetClientToken(v string) *AssociateTransitRouterPolicyTableRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateTransitRouterPolicyTableRequest) SetDryRun(v bool) *AssociateTransitRouterPolicyTableRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateTransitRouterPolicyTableRequest) SetPolicyTableId(v string) *AssociateTransitRouterPolicyTableRequest {
	s.PolicyTableId = &v
	return s
}

func (s *AssociateTransitRouterPolicyTableRequest) Validate() error {
	return dara.Validate(s)
}
