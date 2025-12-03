// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListApplicationMembersResponseBody
	GetCurrent() *int64
	SetPageSize(v int64) *ListApplicationMembersResponseBody
	GetPageSize() *int64
	SetPages(v int64) *ListApplicationMembersResponseBody
	GetPages() *int64
	SetRecords(v []*ListApplicationMembersResponseBodyRecords) *ListApplicationMembersResponseBody
	GetRecords() []*ListApplicationMembersResponseBodyRecords
	SetRequestId(v string) *ListApplicationMembersResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListApplicationMembersResponseBody
	GetTotal() *int64
}

type ListApplicationMembersResponseBody struct {
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 10
	Pages   *int64                                       `json:"pages,omitempty" xml:"pages,omitempty"`
	Records []*ListApplicationMembersResponseBodyRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
	// example:
	//
	// FC93CE1A-8D7A-13A9-8306-7465DE2E5C0F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListApplicationMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationMembersResponseBody) GetCurrent() *int64 {
	return s.Current
}

func (s *ListApplicationMembersResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationMembersResponseBody) GetPages() *int64 {
	return s.Pages
}

func (s *ListApplicationMembersResponseBody) GetRecords() []*ListApplicationMembersResponseBodyRecords {
	return s.Records
}

func (s *ListApplicationMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationMembersResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListApplicationMembersResponseBody) SetCurrent(v int64) *ListApplicationMembersResponseBody {
	s.Current = &v
	return s
}

func (s *ListApplicationMembersResponseBody) SetPageSize(v int64) *ListApplicationMembersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListApplicationMembersResponseBody) SetPages(v int64) *ListApplicationMembersResponseBody {
	s.Pages = &v
	return s
}

func (s *ListApplicationMembersResponseBody) SetRecords(v []*ListApplicationMembersResponseBodyRecords) *ListApplicationMembersResponseBody {
	s.Records = v
	return s
}

func (s *ListApplicationMembersResponseBody) SetRequestId(v string) *ListApplicationMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationMembersResponseBody) SetTotal(v int64) *ListApplicationMembersResponseBody {
	s.Total = &v
	return s
}

func (s *ListApplicationMembersResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationMembersResponseBodyRecords struct {
	// example:
	//
	// http://
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// 成语描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 成员1
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// 1332695887xxxxxx
	Id       *string                                              `json:"id,omitempty" xml:"id,omitempty"`
	RoleList []*ListApplicationMembersResponseBodyRecordsRoleList `json:"roleList,omitempty" xml:"roleList,omitempty" type:"Repeated"`
	// example:
	//
	// User
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListApplicationMembersResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMembersResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *ListApplicationMembersResponseBodyRecords) GetAvatar() *string {
	return s.Avatar
}

func (s *ListApplicationMembersResponseBodyRecords) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationMembersResponseBodyRecords) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListApplicationMembersResponseBodyRecords) GetId() *string {
	return s.Id
}

func (s *ListApplicationMembersResponseBodyRecords) GetRoleList() []*ListApplicationMembersResponseBodyRecordsRoleList {
	return s.RoleList
}

func (s *ListApplicationMembersResponseBodyRecords) GetType() *string {
	return s.Type
}

func (s *ListApplicationMembersResponseBodyRecords) SetAvatar(v string) *ListApplicationMembersResponseBodyRecords {
	s.Avatar = &v
	return s
}

func (s *ListApplicationMembersResponseBodyRecords) SetDescription(v string) *ListApplicationMembersResponseBodyRecords {
	s.Description = &v
	return s
}

func (s *ListApplicationMembersResponseBodyRecords) SetDisplayName(v string) *ListApplicationMembersResponseBodyRecords {
	s.DisplayName = &v
	return s
}

func (s *ListApplicationMembersResponseBodyRecords) SetId(v string) *ListApplicationMembersResponseBodyRecords {
	s.Id = &v
	return s
}

func (s *ListApplicationMembersResponseBodyRecords) SetRoleList(v []*ListApplicationMembersResponseBodyRecordsRoleList) *ListApplicationMembersResponseBodyRecords {
	s.RoleList = v
	return s
}

func (s *ListApplicationMembersResponseBodyRecords) SetType(v string) *ListApplicationMembersResponseBodyRecords {
	s.Type = &v
	return s
}

func (s *ListApplicationMembersResponseBodyRecords) Validate() error {
	if s.RoleList != nil {
		for _, item := range s.RoleList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationMembersResponseBodyRecordsRoleList struct {
	// example:
	//
	// 开发者
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// developer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListApplicationMembersResponseBodyRecordsRoleList) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationMembersResponseBodyRecordsRoleList) GoString() string {
	return s.String()
}

func (s *ListApplicationMembersResponseBodyRecordsRoleList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListApplicationMembersResponseBodyRecordsRoleList) GetName() *string {
	return s.Name
}

func (s *ListApplicationMembersResponseBodyRecordsRoleList) SetDisplayName(v string) *ListApplicationMembersResponseBodyRecordsRoleList {
	s.DisplayName = &v
	return s
}

func (s *ListApplicationMembersResponseBodyRecordsRoleList) SetName(v string) *ListApplicationMembersResponseBodyRecordsRoleList {
	s.Name = &v
	return s
}

func (s *ListApplicationMembersResponseBodyRecordsRoleList) Validate() error {
	return dara.Validate(s)
}
