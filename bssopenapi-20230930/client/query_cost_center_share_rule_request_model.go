// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterShareRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEcIdAccountIds(v []*QueryCostCenterShareRuleRequestEcIdAccountIds) *QueryCostCenterShareRuleRequest
	GetEcIdAccountIds() []*QueryCostCenterShareRuleRequestEcIdAccountIds
	SetMaxResults(v int32) *QueryCostCenterShareRuleRequest
	GetMaxResults() *int32
	SetNbid(v string) *QueryCostCenterShareRuleRequest
	GetNbid() *string
	SetNextToken(v string) *QueryCostCenterShareRuleRequest
	GetNextToken() *string
	SetOwnerAccountId(v int64) *QueryCostCenterShareRuleRequest
	GetOwnerAccountId() *int64
}

type QueryCostCenterShareRuleRequest struct {
	EcIdAccountIds []*QueryCostCenterShareRuleRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nYCisJwqt18pP5E9yb47iu
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1529600453335198
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s QueryCostCenterShareRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterShareRuleRequest) GoString() string {
	return s.String()
}

func (s *QueryCostCenterShareRuleRequest) GetEcIdAccountIds() []*QueryCostCenterShareRuleRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *QueryCostCenterShareRuleRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryCostCenterShareRuleRequest) GetNbid() *string {
	return s.Nbid
}

func (s *QueryCostCenterShareRuleRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCostCenterShareRuleRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterShareRuleRequest) SetEcIdAccountIds(v []*QueryCostCenterShareRuleRequestEcIdAccountIds) *QueryCostCenterShareRuleRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *QueryCostCenterShareRuleRequest) SetMaxResults(v int32) *QueryCostCenterShareRuleRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCostCenterShareRuleRequest) SetNbid(v string) *QueryCostCenterShareRuleRequest {
	s.Nbid = &v
	return s
}

func (s *QueryCostCenterShareRuleRequest) SetNextToken(v string) *QueryCostCenterShareRuleRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCostCenterShareRuleRequest) SetOwnerAccountId(v int64) *QueryCostCenterShareRuleRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterShareRuleRequest) Validate() error {
	return dara.Validate(s)
}

type QueryCostCenterShareRuleRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1004064243473974
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s QueryCostCenterShareRuleRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterShareRuleRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *QueryCostCenterShareRuleRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *QueryCostCenterShareRuleRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *QueryCostCenterShareRuleRequestEcIdAccountIds) SetAccountIds(v []*int64) *QueryCostCenterShareRuleRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *QueryCostCenterShareRuleRequestEcIdAccountIds) SetEcId(v string) *QueryCostCenterShareRuleRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *QueryCostCenterShareRuleRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
