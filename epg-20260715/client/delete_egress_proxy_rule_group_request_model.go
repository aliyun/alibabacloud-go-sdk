// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyRuleGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEgressProxyRuleGroupRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteEgressProxyRuleGroupRequest
	GetDryRun() *bool
	SetEgressProxyRuleGroupId(v string) *DeleteEgressProxyRuleGroupRequest
	GetEgressProxyRuleGroupId() *string
}

type DeleteEgressProxyRuleGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value of this parameter on the client side. Make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > **Note**: If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request is unique.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: sends a check request and does not delete the rule group. The system checks whether the required parameters are specified, the request format is valid, and the service limits are met. If the check fails, the corresponding error is returned. If the check succeeds, the error code DryRunOperation is returned.
	//
	// - **false*	- (default): sends a regular request. If the check succeeds, the HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress proxy rule group that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// eprg-uf60966a9633667c****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
}

func (s DeleteEgressProxyRuleGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyRuleGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyRuleGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEgressProxyRuleGroupRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteEgressProxyRuleGroupRequest) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *DeleteEgressProxyRuleGroupRequest) SetClientToken(v string) *DeleteEgressProxyRuleGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEgressProxyRuleGroupRequest) SetDryRun(v bool) *DeleteEgressProxyRuleGroupRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteEgressProxyRuleGroupRequest) SetEgressProxyRuleGroupId(v string) *DeleteEgressProxyRuleGroupRequest {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *DeleteEgressProxyRuleGroupRequest) Validate() error {
	return dara.Validate(s)
}
