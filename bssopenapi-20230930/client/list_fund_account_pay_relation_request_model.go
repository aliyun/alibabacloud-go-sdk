// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFundAccountPayRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListFundAccountPayRelationRequest
	GetCurrentPage() *int32
	SetFundAccountId(v string) *ListFundAccountPayRelationRequest
	GetFundAccountId() *string
	SetNbid(v string) *ListFundAccountPayRelationRequest
	GetNbid() *string
	SetPageSize(v int32) *ListFundAccountPayRelationRequest
	GetPageSize() *int32
	SetStatus(v string) *ListFundAccountPayRelationRequest
	GetStatus() *string
}

type ListFundAccountPayRelationRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12323123
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// valid
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFundAccountPayRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFundAccountPayRelationRequest) GoString() string {
	return s.String()
}

func (s *ListFundAccountPayRelationRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListFundAccountPayRelationRequest) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *ListFundAccountPayRelationRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ListFundAccountPayRelationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFundAccountPayRelationRequest) GetStatus() *string {
	return s.Status
}

func (s *ListFundAccountPayRelationRequest) SetCurrentPage(v int32) *ListFundAccountPayRelationRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListFundAccountPayRelationRequest) SetFundAccountId(v string) *ListFundAccountPayRelationRequest {
	s.FundAccountId = &v
	return s
}

func (s *ListFundAccountPayRelationRequest) SetNbid(v string) *ListFundAccountPayRelationRequest {
	s.Nbid = &v
	return s
}

func (s *ListFundAccountPayRelationRequest) SetPageSize(v int32) *ListFundAccountPayRelationRequest {
	s.PageSize = &v
	return s
}

func (s *ListFundAccountPayRelationRequest) SetStatus(v string) *ListFundAccountPayRelationRequest {
	s.Status = &v
	return s
}

func (s *ListFundAccountPayRelationRequest) Validate() error {
	return dara.Validate(s)
}
