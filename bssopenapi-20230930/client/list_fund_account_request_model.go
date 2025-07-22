// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFundAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNbid(v string) *ListFundAccountRequest
	GetNbid() *string
	SetQueryOnlyInUse(v bool) *ListFundAccountRequest
	GetQueryOnlyInUse() *bool
	SetQueryOnlyManage(v bool) *ListFundAccountRequest
	GetQueryOnlyManage() *bool
}

type ListFundAccountRequest struct {
	// example:
	//
	// 2084210001
	Nbid            *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	QueryOnlyInUse  *bool   `json:"QueryOnlyInUse,omitempty" xml:"QueryOnlyInUse,omitempty"`
	QueryOnlyManage *bool   `json:"QueryOnlyManage,omitempty" xml:"QueryOnlyManage,omitempty"`
}

func (s ListFundAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFundAccountRequest) GoString() string {
	return s.String()
}

func (s *ListFundAccountRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ListFundAccountRequest) GetQueryOnlyInUse() *bool {
	return s.QueryOnlyInUse
}

func (s *ListFundAccountRequest) GetQueryOnlyManage() *bool {
	return s.QueryOnlyManage
}

func (s *ListFundAccountRequest) SetNbid(v string) *ListFundAccountRequest {
	s.Nbid = &v
	return s
}

func (s *ListFundAccountRequest) SetQueryOnlyInUse(v bool) *ListFundAccountRequest {
	s.QueryOnlyInUse = &v
	return s
}

func (s *ListFundAccountRequest) SetQueryOnlyManage(v bool) *ListFundAccountRequest {
	s.QueryOnlyManage = &v
	return s
}

func (s *ListFundAccountRequest) Validate() error {
	return dara.Validate(s)
}
