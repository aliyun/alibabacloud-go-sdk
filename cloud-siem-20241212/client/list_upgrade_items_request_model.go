// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUpgradeItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListUpgradeItemsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListUpgradeItemsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListUpgradeItemsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListUpgradeItemsRequest
	GetRegionId() *string
	SetRoleFor(v string) *ListUpgradeItemsRequest
	GetRoleFor() *string
}

type ListUpgradeItemsRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *string `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListUpgradeItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUpgradeItemsRequest) GoString() string {
	return s.String()
}

func (s *ListUpgradeItemsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListUpgradeItemsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListUpgradeItemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUpgradeItemsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListUpgradeItemsRequest) GetRoleFor() *string {
	return s.RoleFor
}

func (s *ListUpgradeItemsRequest) SetLang(v string) *ListUpgradeItemsRequest {
	s.Lang = &v
	return s
}

func (s *ListUpgradeItemsRequest) SetMaxResults(v int32) *ListUpgradeItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListUpgradeItemsRequest) SetNextToken(v string) *ListUpgradeItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListUpgradeItemsRequest) SetRegionId(v string) *ListUpgradeItemsRequest {
	s.RegionId = &v
	return s
}

func (s *ListUpgradeItemsRequest) SetRoleFor(v string) *ListUpgradeItemsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListUpgradeItemsRequest) Validate() error {
	return dara.Validate(s)
}
