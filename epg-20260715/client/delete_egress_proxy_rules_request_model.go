// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteEgressProxyRulesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteEgressProxyRulesRequest
	GetDryRun() *bool
	SetEgressProxyRuleGroupId(v string) *DeleteEgressProxyRulesRequest
	GetEgressProxyRuleGroupId() *string
	SetEgressProxyRuleIds(v []*string) *DeleteEgressProxyRulesRequest
	GetEgressProxyRuleIds() []*string
}

type DeleteEgressProxyRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > Note: If you do not specify this parameter, the system automatically uses the request ID of the API request as the client token. The request ID of each API request is unique.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: performs only a dry run and does not delete the rules. The system checks whether required parameters are specified, the request format is valid, and the service limits are met. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): performs a dry run and sends the actual request. If the request passes the dry run, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress proxy rule group.
	//
	// This parameter is required.
	//
	// example:
	//
	// eprg-uf60966a9633667c****
	EgressProxyRuleGroupId *string `json:"EgressProxyRuleGroupId,omitempty" xml:"EgressProxyRuleGroupId,omitempty"`
	// The IDs of the egress proxy rules.
	//
	// This parameter is required.
	EgressProxyRuleIds []*string `json:"EgressProxyRuleIds,omitempty" xml:"EgressProxyRuleIds,omitempty" type:"Repeated"`
}

func (s DeleteEgressProxyRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteEgressProxyRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteEgressProxyRulesRequest) GetEgressProxyRuleGroupId() *string {
	return s.EgressProxyRuleGroupId
}

func (s *DeleteEgressProxyRulesRequest) GetEgressProxyRuleIds() []*string {
	return s.EgressProxyRuleIds
}

func (s *DeleteEgressProxyRulesRequest) SetClientToken(v string) *DeleteEgressProxyRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteEgressProxyRulesRequest) SetDryRun(v bool) *DeleteEgressProxyRulesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteEgressProxyRulesRequest) SetEgressProxyRuleGroupId(v string) *DeleteEgressProxyRulesRequest {
	s.EgressProxyRuleGroupId = &v
	return s
}

func (s *DeleteEgressProxyRulesRequest) SetEgressProxyRuleIds(v []*string) *DeleteEgressProxyRulesRequest {
	s.EgressProxyRuleIds = v
	return s
}

func (s *DeleteEgressProxyRulesRequest) Validate() error {
	return dara.Validate(s)
}
