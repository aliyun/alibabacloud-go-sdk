// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupQueryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DsgUserGroupQueryListRequest
	GetName() *string
	SetOwner(v string) *DsgUserGroupQueryListRequest
	GetOwner() *string
	SetPageNumber(v int32) *DsgUserGroupQueryListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DsgUserGroupQueryListRequest
	GetPageSize() *int32
	SetProjectName(v string) *DsgUserGroupQueryListRequest
	GetProjectName() *string
}

type DsgUserGroupQueryListRequest struct {
	// The keyword of the user group name. A fuzzy match is performed based on the keyword to search for the user group.
	//
	// example:
	//
	// yun_group
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the user group.
	//
	// example:
	//
	// user1
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the compute engine. If you want to query the information about a MaxCompute user group, you need to configure this parameter.
	//
	// example:
	//
	// dev_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s DsgUserGroupQueryListRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupQueryListRequest) GoString() string {
	return s.String()
}

func (s *DsgUserGroupQueryListRequest) GetName() *string {
	return s.Name
}

func (s *DsgUserGroupQueryListRequest) GetOwner() *string {
	return s.Owner
}

func (s *DsgUserGroupQueryListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgUserGroupQueryListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgUserGroupQueryListRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DsgUserGroupQueryListRequest) SetName(v string) *DsgUserGroupQueryListRequest {
	s.Name = &v
	return s
}

func (s *DsgUserGroupQueryListRequest) SetOwner(v string) *DsgUserGroupQueryListRequest {
	s.Owner = &v
	return s
}

func (s *DsgUserGroupQueryListRequest) SetPageNumber(v int32) *DsgUserGroupQueryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DsgUserGroupQueryListRequest) SetPageSize(v int32) *DsgUserGroupQueryListRequest {
	s.PageSize = &v
	return s
}

func (s *DsgUserGroupQueryListRequest) SetProjectName(v string) *DsgUserGroupQueryListRequest {
	s.ProjectName = &v
	return s
}

func (s *DsgUserGroupQueryListRequest) Validate() error {
	return dara.Validate(s)
}
