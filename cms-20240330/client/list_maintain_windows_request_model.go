// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMaintainWindowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *ListMaintainWindowsRequest
	GetDirection() *string
	SetEnable(v bool) *ListMaintainWindowsRequest
	GetEnable() *bool
	SetMaintainWindowId(v string) *ListMaintainWindowsRequest
	GetMaintainWindowId() *string
	SetMaintainWindowName(v string) *ListMaintainWindowsRequest
	GetMaintainWindowName() *string
	SetMaxResults(v int32) *ListMaintainWindowsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMaintainWindowsRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListMaintainWindowsRequest
	GetOrderBy() *string
	SetWorkspace(v string) *ListMaintainWindowsRequest
	GetWorkspace() *string
}

type ListMaintainWindowsRequest struct {
	// The sort order. Valid values:
	//
	// - **asc**: Ascending order.
	//
	// - **desc**: Descending order. This is the default value.
	//
	// example:
	//
	// desc
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// Filters by enabled status. Valid values:
	//
	// - **true**: Returns only enabled silence policies.
	//
	// - **false**: Returns only paused silence policies.
	//
	// If you do not specify this parameter, the policies are not filtered by enabled status.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The ID of the silence policy. Exact matching is used. If you do not specify this parameter, the policies are not filtered by ID.
	//
	// example:
	//
	// 3ff3fbd0-8a0b-4b31-9b1c-8e3f0a2c5d71
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// Policy Name of the silence policy. Fuzzy matching is used, where a match occurs if Policy Name contains the specified value. If you do not specify this parameter, the policies are not filtered by name.
	//
	// example:
	//
	// silence-for-release
	MaintainWindowName *string `json:"maintainWindowName,omitempty" xml:"maintainWindowName,omitempty"`
	// The maximum number of records to return for this request. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination cursor. You do not need to specify this parameter for the first query. For subsequent queries, you can specify the non-empty nextToken returned in the previous response. This value does not guarantee that the next page contains data.
	//
	// example:
	//
	// Y21zRXZlbnRCYXNlUGFnZT0x
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The field by which to sort. By default, the policies are sorted by creation time. Valid values:
	//
	// - **createTime**: Creation time.
	//
	// - **updateTime**: Update time.
	//
	// - **enable**: Enabled status.
	//
	// If you specify other values, the policies are sorted by creation time.
	//
	// example:
	//
	// createTime
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The name of the workspace. This parameter is required by the backend and is used to isolate silence policy resources across different business workspaces.
	//
	// example:
	//
	// default-cms-xxxx-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListMaintainWindowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMaintainWindowsRequest) GoString() string {
	return s.String()
}

func (s *ListMaintainWindowsRequest) GetDirection() *string {
	return s.Direction
}

func (s *ListMaintainWindowsRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ListMaintainWindowsRequest) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *ListMaintainWindowsRequest) GetMaintainWindowName() *string {
	return s.MaintainWindowName
}

func (s *ListMaintainWindowsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMaintainWindowsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMaintainWindowsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListMaintainWindowsRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListMaintainWindowsRequest) SetDirection(v string) *ListMaintainWindowsRequest {
	s.Direction = &v
	return s
}

func (s *ListMaintainWindowsRequest) SetEnable(v bool) *ListMaintainWindowsRequest {
	s.Enable = &v
	return s
}

func (s *ListMaintainWindowsRequest) SetMaintainWindowId(v string) *ListMaintainWindowsRequest {
	s.MaintainWindowId = &v
	return s
}

func (s *ListMaintainWindowsRequest) SetMaintainWindowName(v string) *ListMaintainWindowsRequest {
	s.MaintainWindowName = &v
	return s
}

func (s *ListMaintainWindowsRequest) SetMaxResults(v int32) *ListMaintainWindowsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMaintainWindowsRequest) SetNextToken(v string) *ListMaintainWindowsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMaintainWindowsRequest) SetOrderBy(v string) *ListMaintainWindowsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListMaintainWindowsRequest) SetWorkspace(v string) *ListMaintainWindowsRequest {
	s.Workspace = &v
	return s
}

func (s *ListMaintainWindowsRequest) Validate() error {
	return dara.Validate(s)
}
