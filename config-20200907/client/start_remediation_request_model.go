// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRemediationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *StartRemediationRequest
	GetConfigRuleId() *string
}

type StartRemediationRequest struct {
	// The rule ID.
	//
	// For more information about how to obtain the ID of a rule, see [ListConfigRules](https://help.aliyun.com/document_detail/169607.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-8a973ac2e2be00a2****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
}

func (s StartRemediationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartRemediationRequest) GoString() string {
	return s.String()
}

func (s *StartRemediationRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *StartRemediationRequest) SetConfigRuleId(v string) *StartRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *StartRemediationRequest) Validate() error {
	return dara.Validate(s)
}
