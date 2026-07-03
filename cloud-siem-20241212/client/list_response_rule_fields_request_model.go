// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResponseRuleFieldsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListResponseRuleFieldsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListResponseRuleFieldsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResponseRuleFieldsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListResponseRuleFieldsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListResponseRuleFieldsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListResponseRuleFieldsRequest
	GetRoleType() *int32
}

type ListResponseRuleFieldsRequest struct {
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
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that marks the position from which to start reading. Leave this parameter empty to read from the beginning.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Chinese mainland.
	//
	// - ap-southeast-1: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type. Valid values:
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts in the enterprise.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ListResponseRuleFieldsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResponseRuleFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListResponseRuleFieldsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListResponseRuleFieldsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResponseRuleFieldsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResponseRuleFieldsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListResponseRuleFieldsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListResponseRuleFieldsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListResponseRuleFieldsRequest) SetLang(v string) *ListResponseRuleFieldsRequest {
	s.Lang = &v
	return s
}

func (s *ListResponseRuleFieldsRequest) SetMaxResults(v int32) *ListResponseRuleFieldsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResponseRuleFieldsRequest) SetNextToken(v string) *ListResponseRuleFieldsRequest {
	s.NextToken = &v
	return s
}

func (s *ListResponseRuleFieldsRequest) SetRegionId(v string) *ListResponseRuleFieldsRequest {
	s.RegionId = &v
	return s
}

func (s *ListResponseRuleFieldsRequest) SetRoleFor(v int64) *ListResponseRuleFieldsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListResponseRuleFieldsRequest) SetRoleType(v int32) *ListResponseRuleFieldsRequest {
	s.RoleType = &v
	return s
}

func (s *ListResponseRuleFieldsRequest) Validate() error {
	return dara.Validate(s)
}
