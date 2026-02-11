// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertContactGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroups(v []*SearchAlertContactGroupResponseBodyContactGroups) *SearchAlertContactGroupResponseBody
	GetContactGroups() []*SearchAlertContactGroupResponseBodyContactGroups
	SetRequestId(v string) *SearchAlertContactGroupResponseBody
	GetRequestId() *string
}

type SearchAlertContactGroupResponseBody struct {
	ContactGroups []*SearchAlertContactGroupResponseBodyContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Repeated"`
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertContactGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactGroupResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBody) GetContactGroups() []*SearchAlertContactGroupResponseBodyContactGroups {
	return s.ContactGroups
}

func (s *SearchAlertContactGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchAlertContactGroupResponseBody) SetContactGroups(v []*SearchAlertContactGroupResponseBodyContactGroups) *SearchAlertContactGroupResponseBody {
	s.ContactGroups = v
	return s
}

func (s *SearchAlertContactGroupResponseBody) SetRequestId(v string) *SearchAlertContactGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBody) Validate() error {
	if s.ContactGroups != nil {
		for _, item := range s.ContactGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchAlertContactGroupResponseBodyContactGroups struct {
	ContactGroupId   *int64                                                      `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	ContactGroupName *string                                                     `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	Contacts         []*SearchAlertContactGroupResponseBodyContactGroupsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	CreateTime       *int64                                                      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	UpdateTime       *int64                                                      `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId           *string                                                     `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertContactGroupResponseBodyContactGroups) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactGroupResponseBodyContactGroups) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) GetContactGroupId() *int64 {
	return s.ContactGroupId
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) GetContacts() []*SearchAlertContactGroupResponseBodyContactGroupsContacts {
	return s.Contacts
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) GetUserId() *string {
	return s.UserId
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContactGroupId(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.ContactGroupId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContactGroupName(v string) *SearchAlertContactGroupResponseBodyContactGroups {
	s.ContactGroupName = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetContacts(v []*SearchAlertContactGroupResponseBodyContactGroupsContacts) *SearchAlertContactGroupResponseBodyContactGroups {
	s.Contacts = v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetCreateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetUpdateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroups {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) SetUserId(v string) *SearchAlertContactGroupResponseBodyContactGroups {
	s.UserId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroups) Validate() error {
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

type SearchAlertContactGroupResponseBodyContactGroupsContacts struct {
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DingRobot   *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	SystemNoc   *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SearchAlertContactGroupResponseBodyContactGroupsContacts) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactGroupResponseBodyContactGroupsContacts) GoString() string {
	return s.String()
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) GetContactId() *int64 {
	return s.ContactId
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) GetContactName() *string {
	return s.ContactName
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) GetDingRobot() *string {
	return s.DingRobot
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) GetEmail() *string {
	return s.Email
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) GetPhone() *string {
	return s.Phone
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) GetSystemNoc() *bool {
	return s.SystemNoc
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) GetUserId() *string {
	return s.UserId
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetContactId(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetContactName(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetCreateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetDingRobot(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.DingRobot = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetEmail(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.Email = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetPhone(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetSystemNoc(v bool) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.SystemNoc = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetUpdateTime(v int64) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) SetUserId(v string) *SearchAlertContactGroupResponseBodyContactGroupsContacts {
	s.UserId = &v
	return s
}

func (s *SearchAlertContactGroupResponseBodyContactGroupsContacts) Validate() error {
	return dara.Validate(s)
}
