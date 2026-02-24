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
	DesAggregatorIds *string `json:"DesAggregatorIds,omitempty" xml:"DesAggregatorIds,omitempty"`
	SrcAggregatorId  *string `json:"SrcAggregatorId,omitempty" xml:"SrcAggregatorId,omitempty"`
	// This parameter is required.
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
