// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *QueryCostCenterRuleRequest
	GetCostCenterId() *int64
	SetEcIdAccountIds(v []*QueryCostCenterRuleRequestEcIdAccountIds) *QueryCostCenterRuleRequest
	GetEcIdAccountIds() []*QueryCostCenterRuleRequestEcIdAccountIds
	SetNbid(v string) *QueryCostCenterRuleRequest
	GetNbid() *string
}

type QueryCostCenterRuleRequest struct {
	// example:
	//
	// 597745
	CostCenterId   *int64                                      `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	EcIdAccountIds []*QueryCostCenterRuleRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
}

func (s QueryCostCenterRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterRuleRequest) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRuleRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *QueryCostCenterRuleRequest) GetEcIdAccountIds() []*QueryCostCenterRuleRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *QueryCostCenterRuleRequest) GetNbid() *string {
	return s.Nbid
}

func (s *QueryCostCenterRuleRequest) SetCostCenterId(v int64) *QueryCostCenterRuleRequest {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterRuleRequest) SetEcIdAccountIds(v []*QueryCostCenterRuleRequestEcIdAccountIds) *QueryCostCenterRuleRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *QueryCostCenterRuleRequest) SetNbid(v string) *QueryCostCenterRuleRequest {
	s.Nbid = &v
	return s
}

func (s *QueryCostCenterRuleRequest) Validate() error {
	return dara.Validate(s)
}

type QueryCostCenterRuleRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1004064243473974
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s QueryCostCenterRuleRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterRuleRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRuleRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *QueryCostCenterRuleRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *QueryCostCenterRuleRequestEcIdAccountIds) SetAccountIds(v []*int64) *QueryCostCenterRuleRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *QueryCostCenterRuleRequestEcIdAccountIds) SetEcId(v string) *QueryCostCenterRuleRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *QueryCostCenterRuleRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
