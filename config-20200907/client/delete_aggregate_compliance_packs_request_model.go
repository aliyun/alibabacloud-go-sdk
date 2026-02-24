// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateCompliancePacksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *DeleteAggregateCompliancePacksRequest
	GetAggregatorId() *string
	SetClientToken(v string) *DeleteAggregateCompliancePacksRequest
	GetClientToken() *string
	SetCompliancePackIds(v string) *DeleteAggregateCompliancePacksRequest
	GetCompliancePackIds() *string
	SetDeleteRule(v bool) *DeleteAggregateCompliancePacksRequest
	GetDeleteRule() *bool
}

type DeleteAggregateCompliancePacksRequest struct {
	// This parameter is required.
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	CompliancePackIds *string `json:"CompliancePackIds,omitempty" xml:"CompliancePackIds,omitempty"`
	DeleteRule        *bool   `json:"DeleteRule,omitempty" xml:"DeleteRule,omitempty"`
}

func (s DeleteAggregateCompliancePacksRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateCompliancePacksRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregateCompliancePacksRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *DeleteAggregateCompliancePacksRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAggregateCompliancePacksRequest) GetCompliancePackIds() *string {
	return s.CompliancePackIds
}

func (s *DeleteAggregateCompliancePacksRequest) GetDeleteRule() *bool {
	return s.DeleteRule
}

func (s *DeleteAggregateCompliancePacksRequest) SetAggregatorId(v string) *DeleteAggregateCompliancePacksRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) SetClientToken(v string) *DeleteAggregateCompliancePacksRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) SetCompliancePackIds(v string) *DeleteAggregateCompliancePacksRequest {
	s.CompliancePackIds = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) SetDeleteRule(v bool) *DeleteAggregateCompliancePacksRequest {
	s.DeleteRule = &v
	return s
}

func (s *DeleteAggregateCompliancePacksRequest) Validate() error {
	return dara.Validate(s)
}
