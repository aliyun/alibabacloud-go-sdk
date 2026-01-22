// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterId(v int64) *QueryCostCenterResourceRequest
	GetCostCenterId() *int64
	SetEcIdAccountIds(v []*QueryCostCenterResourceRequestEcIdAccountIds) *QueryCostCenterResourceRequest
	GetEcIdAccountIds() []*QueryCostCenterResourceRequestEcIdAccountIds
	SetMaxResults(v int32) *QueryCostCenterResourceRequest
	GetMaxResults() *int32
	SetNbid(v string) *QueryCostCenterResourceRequest
	GetNbid() *string
	SetNextToken(v string) *QueryCostCenterResourceRequest
	GetNextToken() *string
	SetOwnerAccountId(v int64) *QueryCostCenterResourceRequest
	GetOwnerAccountId() *int64
}

type QueryCostCenterResourceRequest struct {
	// example:
	//
	// 123456
	CostCenterId   *int64                                          `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	EcIdAccountIds []*QueryCostCenterResourceRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// CAESEgoQCg4KCmd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1234567812345678
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s QueryCostCenterResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterResourceRequest) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceRequest) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *QueryCostCenterResourceRequest) GetEcIdAccountIds() []*QueryCostCenterResourceRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *QueryCostCenterResourceRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *QueryCostCenterResourceRequest) GetNbid() *string {
	return s.Nbid
}

func (s *QueryCostCenterResourceRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCostCenterResourceRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterResourceRequest) SetCostCenterId(v int64) *QueryCostCenterResourceRequest {
	s.CostCenterId = &v
	return s
}

func (s *QueryCostCenterResourceRequest) SetEcIdAccountIds(v []*QueryCostCenterResourceRequestEcIdAccountIds) *QueryCostCenterResourceRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *QueryCostCenterResourceRequest) SetMaxResults(v int32) *QueryCostCenterResourceRequest {
	s.MaxResults = &v
	return s
}

func (s *QueryCostCenterResourceRequest) SetNbid(v string) *QueryCostCenterResourceRequest {
	s.Nbid = &v
	return s
}

func (s *QueryCostCenterResourceRequest) SetNextToken(v string) *QueryCostCenterResourceRequest {
	s.NextToken = &v
	return s
}

func (s *QueryCostCenterResourceRequest) SetOwnerAccountId(v int64) *QueryCostCenterResourceRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterResourceRequest) Validate() error {
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryCostCenterResourceRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1501603440974415
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s QueryCostCenterResourceRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterResourceRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *QueryCostCenterResourceRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *QueryCostCenterResourceRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *QueryCostCenterResourceRequestEcIdAccountIds) SetAccountIds(v []*int64) *QueryCostCenterResourceRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *QueryCostCenterResourceRequestEcIdAccountIds) SetEcId(v string) *QueryCostCenterResourceRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *QueryCostCenterResourceRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
