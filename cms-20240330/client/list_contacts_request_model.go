// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactIds(v []*string) *ListContactsRequest
	GetContactIds() []*string
	SetEmail(v string) *ListContactsRequest
	GetEmail() *string
	SetGroupId(v string) *ListContactsRequest
	GetGroupId() *string
	SetName(v string) *ListContactsRequest
	GetName() *string
	SetPageNumber(v int64) *ListContactsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListContactsRequest
	GetPageSize() *int64
	SetPhone(v string) *ListContactsRequest
	GetPhone() *string
	SetQueryUngroupedContacts(v bool) *ListContactsRequest
	GetQueryUngroupedContacts() *bool
	SetWorkspace(v string) *ListContactsRequest
	GetWorkspace() *string
}

type ListContactsRequest struct {
	ContactIds []*string `json:"contactIds,omitempty" xml:"contactIds,omitempty" type:"Repeated"`
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

func (s ListContactsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListContactsRequest) GoString() string {
	return s.String()
}

func (s *ListContactsRequest) GetContactIds() []*string {
	return s.ContactIds
}

func (s *ListContactsRequest) GetEmail() *string {
	return s.Email
}

func (s *ListContactsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListContactsRequest) GetName() *string {
	return s.Name
}

func (s *ListContactsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListContactsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListContactsRequest) GetPhone() *string {
	return s.Phone
}

func (s *ListContactsRequest) GetQueryUngroupedContacts() *bool {
	return s.QueryUngroupedContacts
}

func (s *ListContactsRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListContactsRequest) SetContactIds(v []*string) *ListContactsRequest {
	s.ContactIds = v
	return s
}

func (s *ListContactsRequest) SetEmail(v string) *ListContactsRequest {
	s.Email = &v
	return s
}

func (s *ListContactsRequest) SetGroupId(v string) *ListContactsRequest {
	s.GroupId = &v
	return s
}

func (s *ListContactsRequest) SetName(v string) *ListContactsRequest {
	s.Name = &v
	return s
}

func (s *ListContactsRequest) SetPageNumber(v int64) *ListContactsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListContactsRequest) SetPageSize(v int64) *ListContactsRequest {
	s.PageSize = &v
	return s
}

func (s *ListContactsRequest) SetPhone(v string) *ListContactsRequest {
	s.Phone = &v
	return s
}

func (s *ListContactsRequest) SetQueryUngroupedContacts(v bool) *ListContactsRequest {
	s.QueryUngroupedContacts = &v
	return s
}

func (s *ListContactsRequest) SetWorkspace(v string) *ListContactsRequest {
	s.Workspace = &v
	return s
}

func (s *ListContactsRequest) Validate() error {
	return dara.Validate(s)
}
