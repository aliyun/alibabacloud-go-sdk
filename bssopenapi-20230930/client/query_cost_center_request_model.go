// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryCostCenterRequest
	GetCurrentPage() *int32
	SetEcIdAccountIds(v []*QueryCostCenterRequestEcIdAccountIds) *QueryCostCenterRequest
	GetEcIdAccountIds() []*QueryCostCenterRequestEcIdAccountIds
	SetNbid(v string) *QueryCostCenterRequest
	GetNbid() *string
	SetOwnerAccountId(v int64) *QueryCostCenterRequest
	GetOwnerAccountId() *int64
	SetPageSize(v int32) *QueryCostCenterRequest
	GetPageSize() *int32
	SetParentCostCenterId(v int64) *QueryCostCenterRequest
	GetParentCostCenterId() *int64
}

type QueryCostCenterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage    *int32                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds []*QueryCostCenterRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
}

func (s QueryCostCenterRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterRequest) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryCostCenterRequest) GetEcIdAccountIds() []*QueryCostCenterRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *QueryCostCenterRequest) GetNbid() *string {
	return s.Nbid
}

func (s *QueryCostCenterRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCostCenterRequest) GetParentCostCenterId() *int64 {
	return s.ParentCostCenterId
}

func (s *QueryCostCenterRequest) SetCurrentPage(v int32) *QueryCostCenterRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryCostCenterRequest) SetEcIdAccountIds(v []*QueryCostCenterRequestEcIdAccountIds) *QueryCostCenterRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *QueryCostCenterRequest) SetNbid(v string) *QueryCostCenterRequest {
	s.Nbid = &v
	return s
}

func (s *QueryCostCenterRequest) SetOwnerAccountId(v int64) *QueryCostCenterRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterRequest) SetPageSize(v int32) *QueryCostCenterRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCostCenterRequest) SetParentCostCenterId(v int64) *QueryCostCenterRequest {
	s.ParentCostCenterId = &v
	return s
}

func (s *QueryCostCenterRequest) Validate() error {
	return dara.Validate(s)
}

type QueryCostCenterRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1004064243473974
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s QueryCostCenterRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *QueryCostCenterRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *QueryCostCenterRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *QueryCostCenterRequestEcIdAccountIds) SetAccountIds(v []*int64) *QueryCostCenterRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *QueryCostCenterRequestEcIdAccountIds) SetEcId(v string) *QueryCostCenterRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *QueryCostCenterRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
