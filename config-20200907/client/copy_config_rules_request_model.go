// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyConfigRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesAggregatorIds(v string) *CopyConfigRulesRequest
	GetDesAggregatorIds() *string
	SetSrcAggregatorId(v string) *CopyConfigRulesRequest
	GetSrcAggregatorId() *string
	SetSrcConfigRuleIds(v string) *CopyConfigRulesRequest
	GetSrcConfigRuleIds() *string
}

type CopyConfigRulesRequest struct {
	// The IDs of the destination account groups into which the rules are replicated. Separate multiple account group IDs with commas (,).
	//
	// > If you leave this parameter empty, the compliance packages are replicated into the same account group.
	//
	// example:
	//
	// ca-4b4e626622af005d****
	DesAggregatorIds *string `json:"DesAggregatorIds,omitempty" xml:"DesAggregatorIds,omitempty"`
	// The ID of the account group to which the rules belong.
	//
	// For more information about how to obtain the ID of an account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// example:
	//
	// ca-24db626622af0060****
	SrcAggregatorId *string `json:"SrcAggregatorId,omitempty" xml:"SrcAggregatorId,omitempty"`
	// The rule IDs. Separate multiple rule IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// cr-4b57626622af0065****,cr-47c1626622af0050****
	SrcConfigRuleIds *string `json:"SrcConfigRuleIds,omitempty" xml:"SrcConfigRuleIds,omitempty"`
}

func (s CopyConfigRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyConfigRulesRequest) GoString() string {
	return s.String()
}

func (s *CopyConfigRulesRequest) GetDesAggregatorIds() *string {
	return s.DesAggregatorIds
}

func (s *CopyConfigRulesRequest) GetSrcAggregatorId() *string {
	return s.SrcAggregatorId
}

func (s *CopyConfigRulesRequest) GetSrcConfigRuleIds() *string {
	return s.SrcConfigRuleIds
}

func (s *CopyConfigRulesRequest) SetDesAggregatorIds(v string) *CopyConfigRulesRequest {
	s.DesAggregatorIds = &v
	return s
}

func (s *CopyConfigRulesRequest) SetSrcAggregatorId(v string) *CopyConfigRulesRequest {
	s.SrcAggregatorId = &v
	return s
}

func (s *CopyConfigRulesRequest) SetSrcConfigRuleIds(v string) *CopyConfigRulesRequest {
	s.SrcConfigRuleIds = &v
	return s
}

func (s *CopyConfigRulesRequest) Validate() error {
	return dara.Validate(s)
}
