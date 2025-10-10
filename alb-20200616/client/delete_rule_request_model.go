// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteRuleRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteRuleRequest
	GetDryRun() *bool
	SetRuleId(v string) *DeleteRuleRequest
	GetRuleId() *string
}

type DeleteRuleRequest struct {
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
	// 	- **true**: sends the request without performing the operation. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- **false**: performs a dry run and sends the request. If the request passes the dry run, the `HTTP_2xx` status code is returned and the operation is performed. This is the default value.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the forwarding rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-doc****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteRuleRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteRuleRequest) SetClientToken(v string) *DeleteRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteRuleRequest) SetDryRun(v bool) *DeleteRuleRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteRuleRequest) SetRuleId(v string) *DeleteRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteRuleRequest) Validate() error {
	return dara.Validate(s)
}
