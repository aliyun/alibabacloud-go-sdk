// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPolicyTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterPolicyTableRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterPolicyTableRequest
	GetDryRun() *bool
	SetPolicyTableId(v string) *DeleteTransitRouterPolicyTableRequest
	GetPolicyTableId() *string
}

type DeleteTransitRouterPolicyTableRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the request. Checks include permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, the policy route table is deleted.
	//
	// - **true**: performs a dry run. The policy route table is not deleted. The system checks whether required parameters are specified and whether the request format is valid. If the request fails the dry run, an error is returned. If the request passes the dry run, the request ID is returned.
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

func (s DeleteTransitRouterPolicyTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPolicyTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPolicyTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterPolicyTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterPolicyTableRequest) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *DeleteTransitRouterPolicyTableRequest) SetClientToken(v string) *DeleteTransitRouterPolicyTableRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterPolicyTableRequest) SetDryRun(v bool) *DeleteTransitRouterPolicyTableRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterPolicyTableRequest) SetPolicyTableId(v string) *DeleteTransitRouterPolicyTableRequest {
	s.PolicyTableId = &v
	return s
}

func (s *DeleteTransitRouterPolicyTableRequest) Validate() error {
	return dara.Validate(s)
}
