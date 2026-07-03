// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryViewsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListQueryViewsRequest
	GetLang() *string
	SetMaxResults(v int32) *ListQueryViewsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListQueryViewsRequest
	GetNextToken() *string
	SetQueryViewScene(v string) *ListQueryViewsRequest
	GetQueryViewScene() *string
	SetQueryViewType(v string) *ListQueryViewsRequest
	GetQueryViewType() *string
	SetRegionId(v string) *ListQueryViewsRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListQueryViewsRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *ListQueryViewsRequest
	GetRoleType() *int32
}

type ListQueryViewsRequest struct {
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
	// The maximum number of results to return per request when using the NextToken-based pagination. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. You do not need to specify this parameter for the first query or if no more results exist. If more results exist, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The scene to which the query view belongs.
	//
	// example:
	//
	// Alert
	QueryViewScene *string `json:"QueryViewScene,omitempty" xml:"QueryViewScene,omitempty"`
	// The view type. If this parameter is left empty, all views are returned.
	//
	// example:
	//
	// preset
	QueryViewType *string `json:"QueryViewType,omitempty" xml:"QueryViewType,omitempty"`
	// The region where the threat analysis data management center is located. Specify the management center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets belong to the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Your assets belong to regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
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

func (s ListQueryViewsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListQueryViewsRequest) GoString() string {
	return s.String()
}

func (s *ListQueryViewsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListQueryViewsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListQueryViewsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListQueryViewsRequest) GetQueryViewScene() *string {
	return s.QueryViewScene
}

func (s *ListQueryViewsRequest) GetQueryViewType() *string {
	return s.QueryViewType
}

func (s *ListQueryViewsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListQueryViewsRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListQueryViewsRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *ListQueryViewsRequest) SetLang(v string) *ListQueryViewsRequest {
	s.Lang = &v
	return s
}

func (s *ListQueryViewsRequest) SetMaxResults(v int32) *ListQueryViewsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListQueryViewsRequest) SetNextToken(v string) *ListQueryViewsRequest {
	s.NextToken = &v
	return s
}

func (s *ListQueryViewsRequest) SetQueryViewScene(v string) *ListQueryViewsRequest {
	s.QueryViewScene = &v
	return s
}

func (s *ListQueryViewsRequest) SetQueryViewType(v string) *ListQueryViewsRequest {
	s.QueryViewType = &v
	return s
}

func (s *ListQueryViewsRequest) SetRegionId(v string) *ListQueryViewsRequest {
	s.RegionId = &v
	return s
}

func (s *ListQueryViewsRequest) SetRoleFor(v int64) *ListQueryViewsRequest {
	s.RoleFor = &v
	return s
}

func (s *ListQueryViewsRequest) SetRoleType(v int32) *ListQueryViewsRequest {
	s.RoleType = &v
	return s
}

func (s *ListQueryViewsRequest) Validate() error {
	return dara.Validate(s)
}
