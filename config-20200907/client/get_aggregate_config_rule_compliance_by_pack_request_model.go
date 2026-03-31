// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateConfigRuleComplianceByPackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *GetAggregateConfigRuleComplianceByPackRequest
	GetAggregatorId() *string
	SetCompliancePackId(v string) *GetAggregateConfigRuleComplianceByPackRequest
	GetCompliancePackId() *string
}

type GetAggregateConfigRuleComplianceByPackRequest struct {
	// The ID of the account group.
	//
	// For information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-04b3fd170e340007****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the compliance package.
	//
	// For information about how to obtain the ID of a compliance package, see [ListAggregateCompliancePacks](https://help.aliyun.com/document_detail/262059.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-541e626622af0087****
	CompliancePackId *string `json:"CompliancePackId,omitempty" xml:"CompliancePackId,omitempty"`
}

func (s GetAggregateConfigRuleComplianceByPackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateConfigRuleComplianceByPackRequest) GoString() string {
	return s.String()
}

func (s *GetAggregateConfigRuleComplianceByPackRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *GetAggregateConfigRuleComplianceByPackRequest) GetCompliancePackId() *string {
	return s.CompliancePackId
}

func (s *GetAggregateConfigRuleComplianceByPackRequest) SetAggregatorId(v string) *GetAggregateConfigRuleComplianceByPackRequest {
	s.AggregatorId = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackRequest) SetCompliancePackId(v string) *GetAggregateConfigRuleComplianceByPackRequest {
	s.CompliancePackId = &v
	return s
}

func (s *GetAggregateConfigRuleComplianceByPackRequest) Validate() error {
	return dara.Validate(s)
}
