// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupIdsShrink(v string) *ListContactGroupsShrinkRequest
	GetContactGroupIdsShrink() *string
	SetName(v string) *ListContactGroupsShrinkRequest
	GetName() *string
	SetPageNumber(v int64) *ListContactGroupsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListContactGroupsShrinkRequest
	GetPageSize() *int64
	SetWorkspace(v string) *ListContactGroupsShrinkRequest
	GetWorkspace() *string
}

type ListContactGroupsShrinkRequest struct {
	ContactGroupIdsShrink *string `json:"contactGroupIds,omitempty" xml:"contactGroupIds,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize  *int64  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListContactGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContactGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListContactGroupsShrinkRequest) GetContactGroupIdsShrink() *string {
	return s.ContactGroupIdsShrink
}

func (s *ListContactGroupsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListContactGroupsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListContactGroupsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListContactGroupsShrinkRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListContactGroupsShrinkRequest) SetContactGroupIdsShrink(v string) *ListContactGroupsShrinkRequest {
	s.ContactGroupIdsShrink = &v
	return s
}

func (s *ListContactGroupsShrinkRequest) SetName(v string) *ListContactGroupsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListContactGroupsShrinkRequest) SetPageNumber(v int64) *ListContactGroupsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContactGroupsShrinkRequest) SetPageSize(v int64) *ListContactGroupsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListContactGroupsShrinkRequest) SetWorkspace(v string) *ListContactGroupsShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *ListContactGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
