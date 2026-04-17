// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIdsShrink(v string) *ListContactsShrinkRequest
	GetContactIdsShrink() *string
	SetEmail(v string) *ListContactsShrinkRequest
	GetEmail() *string
	SetGroupId(v string) *ListContactsShrinkRequest
	GetGroupId() *string
	SetName(v string) *ListContactsShrinkRequest
	GetName() *string
	SetPageNumber(v int64) *ListContactsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListContactsShrinkRequest
	GetPageSize() *int64
	SetPhone(v string) *ListContactsShrinkRequest
	GetPhone() *string
	SetQueryUngroupedContacts(v bool) *ListContactsShrinkRequest
	GetQueryUngroupedContacts() *bool
	SetWorkspace(v string) *ListContactsShrinkRequest
	GetWorkspace() *string
}

type ListContactsShrinkRequest struct {
	ContactIdsShrink *string `json:"contactIds,omitempty" xml:"contactIds,omitempty"`
	// example:
	//
	// test@aliyun.com
	Email   *string `json:"email,omitempty" xml:"email,omitempty"`
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
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
	// 100
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 15012345678
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// true
	QueryUngroupedContacts *bool   `json:"queryUngroupedContacts,omitempty" xml:"queryUngroupedContacts,omitempty"`
	Workspace              *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListContactsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContactsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListContactsShrinkRequest) GetContactIdsShrink() *string {
	return s.ContactIdsShrink
}

func (s *ListContactsShrinkRequest) GetEmail() *string {
	return s.Email
}

func (s *ListContactsShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListContactsShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListContactsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListContactsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListContactsShrinkRequest) GetPhone() *string {
	return s.Phone
}

func (s *ListContactsShrinkRequest) GetQueryUngroupedContacts() *bool {
	return s.QueryUngroupedContacts
}

func (s *ListContactsShrinkRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListContactsShrinkRequest) SetContactIdsShrink(v string) *ListContactsShrinkRequest {
	s.ContactIdsShrink = &v
	return s
}

func (s *ListContactsShrinkRequest) SetEmail(v string) *ListContactsShrinkRequest {
	s.Email = &v
	return s
}

func (s *ListContactsShrinkRequest) SetGroupId(v string) *ListContactsShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListContactsShrinkRequest) SetName(v string) *ListContactsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListContactsShrinkRequest) SetPageNumber(v int64) *ListContactsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContactsShrinkRequest) SetPageSize(v int64) *ListContactsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListContactsShrinkRequest) SetPhone(v string) *ListContactsShrinkRequest {
	s.Phone = &v
	return s
}

func (s *ListContactsShrinkRequest) SetQueryUngroupedContacts(v bool) *ListContactsShrinkRequest {
	s.QueryUngroupedContacts = &v
	return s
}

func (s *ListContactsShrinkRequest) SetWorkspace(v string) *ListContactsShrinkRequest {
	s.Workspace = &v
	return s
}

func (s *ListContactsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
