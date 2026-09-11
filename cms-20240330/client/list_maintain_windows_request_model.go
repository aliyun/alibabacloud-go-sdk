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
	// example:
	//
	// asc
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// 123-12-312-31-23123
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	// example:
	//
	// test
	MaintainWindowName *string `json:"maintainWindowName,omitempty" xml:"maintainWindowName,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// 123123***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// enable
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// example:
	//
	// workspace-test
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
