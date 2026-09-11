// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIds(v []*string) *ListContactGroupsRequest
	GetContactGroupIds() []*string
	SetName(v string) *ListContactGroupsRequest
	GetName() *string
	SetPageNumber(v int64) *ListContactGroupsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListContactGroupsRequest
	GetPageSize() *int64
	SetWorkspace(v string) *ListContactGroupsRequest
	GetWorkspace() *string
}

type ListContactGroupsRequest struct {
	// The contact group IDs.
	ContactGroupIds []*string `json:"contactGroupIds,omitempty" xml:"contactGroupIds,omitempty" type:"Repeated"`
	// The contact name.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 100.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// default-cms-xxxxxx-cn-beijing
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListContactGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContactGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListContactGroupsRequest) GetContactGroupIds() []*string {
	return s.ContactGroupIds
}

func (s *ListContactGroupsRequest) GetName() *string {
	return s.Name
}

func (s *ListContactGroupsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListContactGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListContactGroupsRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListContactGroupsRequest) SetContactGroupIds(v []*string) *ListContactGroupsRequest {
	s.ContactGroupIds = v
	return s
}

func (s *ListContactGroupsRequest) SetName(v string) *ListContactGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListContactGroupsRequest) SetPageNumber(v int64) *ListContactGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContactGroupsRequest) SetPageSize(v int64) *ListContactGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContactGroupsRequest) SetWorkspace(v string) *ListContactGroupsRequest {
	s.Workspace = &v
	return s
}

func (s *ListContactGroupsRequest) Validate() error {
	return dara.Validate(s)
}
