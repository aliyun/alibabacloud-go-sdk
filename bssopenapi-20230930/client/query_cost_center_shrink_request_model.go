// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCostCenterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *QueryCostCenterShrinkRequest
	GetCurrentPage() *int32
	SetEcIdAccountIdsShrink(v string) *QueryCostCenterShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetNbid(v string) *QueryCostCenterShrinkRequest
	GetNbid() *string
	SetOwnerAccountId(v int64) *QueryCostCenterShrinkRequest
	GetOwnerAccountId() *int64
	SetPageSize(v int32) *QueryCostCenterShrinkRequest
	GetPageSize() *int32
	SetParentCostCenterId(v int64) *QueryCostCenterShrinkRequest
	GetParentCostCenterId() *int64
}

type QueryCostCenterShrinkRequest struct {
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of enterprises and accounts. If this parameter is left empty, the current account is queried.
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// The ID of the primary sales channel. If this parameter is left empty, the sales channel ID of the current user is used by default.
	//
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// The ID of the user who owns the financial unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the parent financial unit. A value of -1 indicates the root financial unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
}

func (s QueryCostCenterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCostCenterShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCostCenterShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryCostCenterShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *QueryCostCenterShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *QueryCostCenterShrinkRequest) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *QueryCostCenterShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryCostCenterShrinkRequest) GetParentCostCenterId() *int64 {
	return s.ParentCostCenterId
}

func (s *QueryCostCenterShrinkRequest) SetCurrentPage(v int32) *QueryCostCenterShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetEcIdAccountIdsShrink(v string) *QueryCostCenterShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetNbid(v string) *QueryCostCenterShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetOwnerAccountId(v int64) *QueryCostCenterShrinkRequest {
	s.OwnerAccountId = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetPageSize(v int32) *QueryCostCenterShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) SetParentCostCenterId(v int64) *QueryCostCenterShrinkRequest {
	s.ParentCostCenterId = &v
	return s
}

func (s *QueryCostCenterShrinkRequest) Validate() error {
	return dara.Validate(s)
}
