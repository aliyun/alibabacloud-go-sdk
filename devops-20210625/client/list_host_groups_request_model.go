// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateEndTime(v int64) *ListHostGroupsRequest
	GetCreateEndTime() *int64
	SetCreateStartTime(v int64) *ListHostGroupsRequest
	GetCreateStartTime() *int64
	SetCreatorAccountIds(v string) *ListHostGroupsRequest
	GetCreatorAccountIds() *string
	SetIds(v string) *ListHostGroupsRequest
	GetIds() *string
	SetMaxResults(v int64) *ListHostGroupsRequest
	GetMaxResults() *int64
	SetName(v string) *ListHostGroupsRequest
	GetName() *string
	SetNextToken(v string) *ListHostGroupsRequest
	GetNextToken() *string
	SetPageOrder(v string) *ListHostGroupsRequest
	GetPageOrder() *string
	SetPageSort(v string) *ListHostGroupsRequest
	GetPageSort() *string
}

type ListHostGroupsRequest struct {
	// example:
	//
	// 1586863220000
	CreateEndTime *int64 `json:"createEndTime,omitempty" xml:"createEndTime,omitempty"`
	// example:
	//
	// 1586863220000
	CreateStartTime *int64 `json:"createStartTime,omitempty" xml:"createStartTime,omitempty"`
	// example:
	//
	// 1112122121,3223332
	CreatorAccountIds *string `json:"creatorAccountIds,omitempty" xml:"creatorAccountIds,omitempty"`
	// example:
	//
	// 121,1212121232
	Ids *string `json:"ids,omitempty" xml:"ids,omitempty"`
	// example:
	//
	// 30
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 主机组
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 221212221
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// DESC
	PageOrder *string `json:"pageOrder,omitempty" xml:"pageOrder,omitempty"`
	// example:
	//
	// ID
	PageSort *string `json:"pageSort,omitempty" xml:"pageSort,omitempty"`
}

func (s ListHostGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsRequest) GetCreateEndTime() *int64 {
	return s.CreateEndTime
}

func (s *ListHostGroupsRequest) GetCreateStartTime() *int64 {
	return s.CreateStartTime
}

func (s *ListHostGroupsRequest) GetCreatorAccountIds() *string {
	return s.CreatorAccountIds
}

func (s *ListHostGroupsRequest) GetIds() *string {
	return s.Ids
}

func (s *ListHostGroupsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListHostGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListHostGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHostGroupsRequest) GetPageOrder() *string {
	return s.PageOrder
}

func (s *ListHostGroupsRequest) GetPageSort() *string {
	return s.PageSort
}

func (s *ListHostGroupsRequest) SetCreateEndTime(v int64) *ListHostGroupsRequest {
	s.CreateEndTime = &v
	return s
}

func (s *ListHostGroupsRequest) SetCreateStartTime(v int64) *ListHostGroupsRequest {
	s.CreateStartTime = &v
	return s
}

func (s *ListHostGroupsRequest) SetCreatorAccountIds(v string) *ListHostGroupsRequest {
	s.CreatorAccountIds = &v
	return s
}

func (s *ListHostGroupsRequest) SetIds(v string) *ListHostGroupsRequest {
	s.Ids = &v
	return s
}

func (s *ListHostGroupsRequest) SetMaxResults(v int64) *ListHostGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHostGroupsRequest) SetName(v string) *ListHostGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListHostGroupsRequest) SetNextToken(v string) *ListHostGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListHostGroupsRequest) SetPageOrder(v string) *ListHostGroupsRequest {
	s.PageOrder = &v
	return s
}

func (s *ListHostGroupsRequest) SetPageSort(v string) *ListHostGroupsRequest {
	s.PageSort = &v
	return s
}

func (s *ListHostGroupsRequest) Validate() error {
	return dara.Validate(s)
}
