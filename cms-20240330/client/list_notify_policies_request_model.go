// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNotifyPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListNotifyPoliciesRequest
	GetMaxResults() *int32
	SetName(v string) *ListNotifyPoliciesRequest
	GetName() *string
	SetNextToken(v string) *ListNotifyPoliciesRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListNotifyPoliciesRequest
	GetOrderBy() *string
	SetOrderDesc(v string) *ListNotifyPoliciesRequest
	GetOrderDesc() *string
	SetWorkspace(v string) *ListNotifyPoliciesRequest
	GetWorkspace() *string
}

type ListNotifyPoliciesRequest struct {
	// The maximum number of entries to return in this request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The policy name used for fuzzy filtering.
	//
	// example:
	//
	// prod-alert
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token. Leave this parameter empty for the first page. For subsequent pages, set this parameter to the nextToken value returned in the previous response.
	//
	// example:
	//
	// eyJjdXJzb3IiOjEwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The field used for sorting. Valid values: createTime, updateTime, and name.
	//
	// example:
	//
	// createTime
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// Specifies whether to sort results in descending order. Valid values:
	//
	// - true: descending order.
	//
	// - false: ascending order.
	//
	// example:
	//
	// true
	OrderDesc *string `json:"orderDesc,omitempty" xml:"orderDesc,omitempty"`
	// The workspace ID. This parameter is used to isolate notify policy resources across different business spaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-xxxx-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListNotifyPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNotifyPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListNotifyPoliciesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNotifyPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListNotifyPoliciesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNotifyPoliciesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListNotifyPoliciesRequest) GetOrderDesc() *string {
	return s.OrderDesc
}

func (s *ListNotifyPoliciesRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListNotifyPoliciesRequest) SetMaxResults(v int32) *ListNotifyPoliciesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNotifyPoliciesRequest) SetName(v string) *ListNotifyPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListNotifyPoliciesRequest) SetNextToken(v string) *ListNotifyPoliciesRequest {
	s.NextToken = &v
	return s
}

func (s *ListNotifyPoliciesRequest) SetOrderBy(v string) *ListNotifyPoliciesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListNotifyPoliciesRequest) SetOrderDesc(v string) *ListNotifyPoliciesRequest {
	s.OrderDesc = &v
	return s
}

func (s *ListNotifyPoliciesRequest) SetWorkspace(v string) *ListNotifyPoliciesRequest {
	s.Workspace = &v
	return s
}

func (s *ListNotifyPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
