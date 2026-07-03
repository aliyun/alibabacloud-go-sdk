// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListTagsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListTagsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListTagsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListTagsRequest
	GetRoleType() *int32
	SetTargetRelation(v string) *ListTagsRequest
	GetTargetRelation() *string
	SetTargetType(v string) *ListTagsRequest
	GetTargetType() *string
	SetTargetUuid(v string) *ListTagsRequest
	GetTargetUuid() *string
}

type ListTagsRequest struct {
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
	// The maximum number of results to return when using NextToken-based pagination. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or if no more results exist. If a next query is available, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAASLVeIxed4466E0LVmGkzwS6hJKd9DGVGMDRM6Lu****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the threat analysis data management center is located. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: the asset is in the Chinese mainland.
	//
	// - ap-southeast-1: the asset is outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator switches to when viewing as another member.
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
	// The target relationship.
	//
	// example:
	//
	// {"incident_uuid":"3fce2d3b65fc2bd05cac6d******"}
	TargetRelation *string `json:"TargetRelation,omitempty" xml:"TargetRelation,omitempty"`
	// The scan object type. Valid values:
	//
	// - **1**: snapshot
	//
	// - **2**: image
	//
	// example:
	//
	// AliUId
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The UUID of the target asset.
	//
	// example:
	//
	// a1b2c3-****
	TargetUuid *string `json:"TargetUuid,omitempty" xml:"TargetUuid,omitempty"`
}

func (s ListTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListTagsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListTagsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListTagsRequest) GetTargetRelation() *string {
	return s.TargetRelation
}

func (s *ListTagsRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ListTagsRequest) GetTargetUuid() *string {
	return s.TargetUuid
}

func (s *ListTagsRequest) SetLang(v string) *ListTagsRequest {
	s.Lang = &v
	return s
}

func (s *ListTagsRequest) SetMaxResults(v int32) *ListTagsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagsRequest) SetNextToken(v string) *ListTagsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagsRequest) SetRegionId(v string) *ListTagsRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagsRequest) SetRoleFor(v int64) *ListTagsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListTagsRequest) SetRoleType(v int32) *ListTagsRequest {
	s.RoleType = &v
	return s
}

func (s *ListTagsRequest) SetTargetRelation(v string) *ListTagsRequest {
	s.TargetRelation = &v
	return s
}

func (s *ListTagsRequest) SetTargetType(v string) *ListTagsRequest {
	s.TargetType = &v
	return s
}

func (s *ListTagsRequest) SetTargetUuid(v string) *ListTagsRequest {
	s.TargetUuid = &v
	return s
}

func (s *ListTagsRequest) Validate() error {
	return dara.Validate(s)
}
