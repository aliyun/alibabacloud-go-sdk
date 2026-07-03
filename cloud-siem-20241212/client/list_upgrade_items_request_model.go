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
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to retrieve the next page of results. Leave this parameter empty for the first request. To retrieve the next page, set this parameter to the NextToken value from the previous response.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the Data Management center is located. Select a region based on where your assets are. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member. An administrator can use this parameter to switch to the member\\"s view.
	//
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
