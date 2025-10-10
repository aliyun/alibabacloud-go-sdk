// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteRulesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteRulesRequest
	GetDryRun() *bool
	SetRuleIds(v []*string) *DeleteRulesRequest
	GetRuleIds() []*string
}

type DeleteRulesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF3898
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): performs a dry run and performs the actual request. If the request passes the dry run, a `2xx HTTP` status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The forwarding rules. You can specify at most 100 forwarding rules in each call.
	//
	// This parameter is required.
	RuleIds []*string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s DeleteRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRulesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteRulesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteRulesRequest) GetRuleIds() []*string {
	return s.RuleIds
}

func (s *DeleteRulesRequest) SetClientToken(v string) *DeleteRulesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRulesRequest) SetDryRun(v bool) *DeleteRulesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRulesRequest) SetRuleIds(v []*string) *DeleteRulesRequest {
	s.RuleIds = v
	return s
}

func (s *DeleteRulesRequest) Validate() error {
	return dara.Validate(s)
}
