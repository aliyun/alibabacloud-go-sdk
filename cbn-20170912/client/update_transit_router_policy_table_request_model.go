// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTransitRouterPolicyTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTransitRouterPolicyTableRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateTransitRouterPolicyTableRequest
	GetDescription() *string
	SetDryRun(v bool) *UpdateTransitRouterPolicyTableRequest
	GetDryRun() *bool
	SetName(v string) *UpdateTransitRouterPolicyTableRequest
	GetName() *string
	SetPolicyTableId(v string) *UpdateTransitRouterPolicyTableRequest
	GetPolicyTableId() *string
}

type UpdateTransitRouterPolicyTableRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-4266****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the policy route table.
	//
	// The description can be empty or 1 to 256 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run on the request, including checks on permissions and instance status. Valid values:
	//
	// - **false*	- (default): sends a normal request. If the request passes the check, the policy route table is updated.
	//
	// - **true**: sends a dry run request. Only the check is performed, and the policy route table is not updated. Check items include whether required parameters are specified and the request format. If the check fails, an error is returned. If the check passes, the request ID is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the policy route table.
	//
	// The name can be empty or 1 to 128 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The instance ID of the policy routing table.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptbl-abc123xyz
	PolicyTableId *string `json:"PolicyTableId,omitempty" xml:"PolicyTableId,omitempty"`
}

func (s UpdateTransitRouterPolicyTableRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTransitRouterPolicyTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransitRouterPolicyTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTransitRouterPolicyTableRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateTransitRouterPolicyTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTransitRouterPolicyTableRequest) GetName() *string {
	return s.Name
}

func (s *UpdateTransitRouterPolicyTableRequest) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *UpdateTransitRouterPolicyTableRequest) SetClientToken(v string) *UpdateTransitRouterPolicyTableRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTransitRouterPolicyTableRequest) SetDescription(v string) *UpdateTransitRouterPolicyTableRequest {
	s.Description = &v
	return s
}

func (s *UpdateTransitRouterPolicyTableRequest) SetDryRun(v bool) *UpdateTransitRouterPolicyTableRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTransitRouterPolicyTableRequest) SetName(v string) *UpdateTransitRouterPolicyTableRequest {
	s.Name = &v
	return s
}

func (s *UpdateTransitRouterPolicyTableRequest) SetPolicyTableId(v string) *UpdateTransitRouterPolicyTableRequest {
	s.PolicyTableId = &v
	return s
}

func (s *UpdateTransitRouterPolicyTableRequest) Validate() error {
	return dara.Validate(s)
}
