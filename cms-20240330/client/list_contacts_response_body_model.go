// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContacts(v []*ListContactsResponseBodyContacts) *ListContactsResponseBody
	GetContacts() []*ListContactsResponseBodyContacts
	SetPageNumber(v int64) *ListContactsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListContactsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListContactsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListContactsResponseBody
	GetTotal() *int64
}

type ListContactsResponseBody struct {
	Contacts []*ListContactsResponseBodyContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 56
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactsResponseBody) GetContacts() []*ListContactsResponseBodyContacts {
	return s.Contacts
}

func (s *ListContactsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListContactsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContactsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListContactsResponseBody) SetContacts(v []*ListContactsResponseBodyContacts) *ListContactsResponseBody {
	s.Contacts = v
	return s
}

func (s *ListContactsResponseBody) SetPageNumber(v int64) *ListContactsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListContactsResponseBody) SetPageSize(v int64) *ListContactsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListContactsResponseBody) SetRequestId(v string) *ListContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactsResponseBody) SetTotal(v int64) *ListContactsResponseBody {
	s.Total = &v
	return s
}

func (s *ListContactsResponseBody) Validate() error {
	if s.Contacts != nil {
		for _, item := range s.Contacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListContactsResponseBodyContacts struct {
	// example:
	//
	// test
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	// example:
	//
	// test@aliyun.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// true
	EmailVerify *bool              `json:"emailVerify,omitempty" xml:"emailVerify,omitempty"`
	GroupList   []*string          `json:"groupList,omitempty" xml:"groupList,omitempty" type:"Repeated"`
	ImUserIds   map[string]*string `json:"imUserIds,omitempty" xml:"imUserIds,omitempty"`
	// example:
	//
	// zh_CN
	Lang *string `json:"lang,omitempty" xml:"lang,omitempty"`
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 130123456789
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// example:
	//
	// true
	PhoneVerify *bool `json:"phoneVerify,omitempty" xml:"phoneVerify,omitempty"`
	// example:
	//
	// 2024-10-22 02:21:51
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	Workspace  *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListContactsResponseBodyContacts) String() string {
	return dara.Prettify(s)
}

func (s ListContactsResponseBodyContacts) GoString() string {
	return s.String()
}

func (s *ListContactsResponseBodyContacts) GetContactId() *string {
	return s.ContactId
}

func (s *ListContactsResponseBodyContacts) GetEmail() *string {
	return s.Email
}

func (s *ListContactsResponseBodyContacts) GetEmailVerify() *bool {
	return s.EmailVerify
}

func (s *ListContactsResponseBodyContacts) GetGroupList() []*string {
	return s.GroupList
}

func (s *ListContactsResponseBodyContacts) GetImUserIds() map[string]*string {
	return s.ImUserIds
}

func (s *ListContactsResponseBodyContacts) GetLang() *string {
	return s.Lang
}

func (s *ListContactsResponseBodyContacts) GetName() *string {
	return s.Name
}

func (s *ListContactsResponseBodyContacts) GetPhone() *string {
	return s.Phone
}

func (s *ListContactsResponseBodyContacts) GetPhoneVerify() *bool {
	return s.PhoneVerify
}

func (s *ListContactsResponseBodyContacts) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListContactsResponseBodyContacts) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListContactsResponseBodyContacts) SetContactId(v string) *ListContactsResponseBodyContacts {
	s.ContactId = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetEmail(v string) *ListContactsResponseBodyContacts {
	s.Email = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetEmailVerify(v bool) *ListContactsResponseBodyContacts {
	s.EmailVerify = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetGroupList(v []*string) *ListContactsResponseBodyContacts {
	s.GroupList = v
	return s
}

func (s *ListContactsResponseBodyContacts) SetImUserIds(v map[string]*string) *ListContactsResponseBodyContacts {
	s.ImUserIds = v
	return s
}

func (s *ListContactsResponseBodyContacts) SetLang(v string) *ListContactsResponseBodyContacts {
	s.Lang = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetName(v string) *ListContactsResponseBodyContacts {
	s.Name = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetPhone(v string) *ListContactsResponseBodyContacts {
	s.Phone = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetPhoneVerify(v bool) *ListContactsResponseBodyContacts {
	s.PhoneVerify = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetUpdateTime(v string) *ListContactsResponseBodyContacts {
	s.UpdateTime = &v
	return s
}

func (s *ListContactsResponseBodyContacts) SetWorkspace(v string) *ListContactsResponseBodyContacts {
	s.Workspace = &v
	return s
}

func (s *ListContactsResponseBodyContacts) Validate() error {
	return dara.Validate(s)
}
