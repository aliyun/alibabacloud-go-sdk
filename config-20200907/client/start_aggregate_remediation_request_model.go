// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAggregateRemediationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *StartAggregateRemediationRequest
	GetAggregatorId() *string
	SetConfigRuleId(v string) *StartAggregateRemediationRequest
	GetConfigRuleId() *string
	SetResourceAccountId(v int64) *StartAggregateRemediationRequest
	GetResourceAccountId() *int64
}

type StartAggregateRemediationRequest struct {
	// The ID of the account group.
	//
	// For information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-6b4a626622af0012****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The rule ID.
	//
	// For more information about how to obtain the ID of a rule, see [ListAggregateConfigRules](https://help.aliyun.com/document_detail/264148.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-6b7c626622af00b4****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The ID of the Alibaba Cloud account to which the resources to be remediated belong. If this parameter is left empty, non-compliant resources of all accounts in the account group are remediated.
	//
	// > You must specify the ID of the current management account or a member account in the account group of the management account.
	//
	// example:
	//
	// 100271897542****
	ResourceAccountId *int64 `json:"ResourceAccountId,omitempty" xml:"ResourceAccountId,omitempty"`
}

func (s StartAggregateRemediationRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAggregateRemediationRequest) GoString() string {
	return s.String()
}

func (s *StartAggregateRemediationRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *StartAggregateRemediationRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *StartAggregateRemediationRequest) GetResourceAccountId() *int64 {
	return s.ResourceAccountId
}

func (s *StartAggregateRemediationRequest) SetAggregatorId(v string) *StartAggregateRemediationRequest {
	s.AggregatorId = &v
	return s
}

func (s *StartAggregateRemediationRequest) SetConfigRuleId(v string) *StartAggregateRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *StartAggregateRemediationRequest) SetResourceAccountId(v int64) *StartAggregateRemediationRequest {
	s.ResourceAccountId = &v
	return s
}

func (s *StartAggregateRemediationRequest) Validate() error {
	return dara.Validate(s)
}
