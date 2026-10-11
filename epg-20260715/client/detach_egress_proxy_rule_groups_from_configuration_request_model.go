// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachEgressProxyRuleGroupsFromConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetachEgressProxyRuleGroupsFromConfigurationRequest
	GetClientToken() *string
	SetDryRun(v bool) *DetachEgressProxyRuleGroupsFromConfigurationRequest
	GetDryRun() *bool
	SetEgressProxyConfigurationId(v string) *DetachEgressProxyRuleGroupsFromConfigurationRequest
	GetEgressProxyConfigurationId() *string
	SetEgressProxyRuleGroupIds(v []*string) *DetachEgressProxyRuleGroupsFromConfigurationRequest
	GetEgressProxyRuleGroupIds() []*string
}

type DetachEgressProxyRuleGroupsFromConfigurationRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > Note If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId of each API request is unique.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run on the request. Valid values:
	//
	// - **true**: Sends a check request and does not disassociate the egress proxy rule groups. The check items include whether the required parameters are specified, the request format, and service limits. If the check fails, the corresponding error is returned. If the check succeeds, the DryRunOperation error code is returned.
	//
	// - **false*	- (default): Sends a normal request. After the check succeeds, an HTTP 2xx status code is returned and the operation is directly executed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the egress proxy configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// epc-hn302f6a9a284c4****
	EgressProxyConfigurationId *string `json:"EgressProxyConfigurationId,omitempty" xml:"EgressProxyConfigurationId,omitempty"`
	// The IDs of the egress proxy rule groups to be disassociated from the egress proxy configuration.
	//
	// This parameter is required.
	EgressProxyRuleGroupIds []*string `json:"EgressProxyRuleGroupIds,omitempty" xml:"EgressProxyRuleGroupIds,omitempty" type:"Repeated"`
}

func (s DetachEgressProxyRuleGroupsFromConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachEgressProxyRuleGroupsFromConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationRequest) GetEgressProxyConfigurationId() *string {
	return s.EgressProxyConfigurationId
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationRequest) GetEgressProxyRuleGroupIds() []*string {
	return s.EgressProxyRuleGroupIds
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationRequest) SetClientToken(v string) *DetachEgressProxyRuleGroupsFromConfigurationRequest {
	s.ClientToken = &v
	return s
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationRequest) SetDryRun(v bool) *DetachEgressProxyRuleGroupsFromConfigurationRequest {
	s.DryRun = &v
	return s
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationRequest) SetEgressProxyConfigurationId(v string) *DetachEgressProxyRuleGroupsFromConfigurationRequest {
	s.EgressProxyConfigurationId = &v
	return s
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationRequest) SetEgressProxyRuleGroupIds(v []*string) *DetachEgressProxyRuleGroupsFromConfigurationRequest {
	s.EgressProxyRuleGroupIds = v
	return s
}

func (s *DetachEgressProxyRuleGroupsFromConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
