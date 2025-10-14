// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactGroupListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeContactGroupListResponseBody
	GetCode() *string
	SetContactGroupList(v *DescribeContactGroupListResponseBodyContactGroupList) *DescribeContactGroupListResponseBody
	GetContactGroupList() *DescribeContactGroupListResponseBodyContactGroupList
	SetContactGroups(v *DescribeContactGroupListResponseBodyContactGroups) *DescribeContactGroupListResponseBody
	GetContactGroups() *DescribeContactGroupListResponseBodyContactGroups
	SetMessage(v string) *DescribeContactGroupListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeContactGroupListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeContactGroupListResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *DescribeContactGroupListResponseBody
	GetTotal() *int32
}

type DescribeContactGroupListResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about alert groups that were queried.
	ContactGroupList *DescribeContactGroupListResponseBodyContactGroupList `json:"ContactGroupList,omitempty" xml:"ContactGroupList,omitempty" type:"Struct"`
	// The names of alert groups.
	ContactGroups *DescribeContactGroupListResponseBodyContactGroups `json:"ContactGroups,omitempty" xml:"ContactGroups,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 916EE694-03C2-47B6-85EE-5054E3C168D3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of the returned entries.
	//
	// example:
	//
	// 22
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeContactGroupListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeContactGroupListResponseBody) GetContactGroupList() *DescribeContactGroupListResponseBodyContactGroupList {
	return s.ContactGroupList
}

func (s *DescribeContactGroupListResponseBody) GetContactGroups() *DescribeContactGroupListResponseBodyContactGroups {
	return s.ContactGroups
}

func (s *DescribeContactGroupListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeContactGroupListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContactGroupListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeContactGroupListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeContactGroupListResponseBody) SetCode(v string) *DescribeContactGroupListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetContactGroupList(v *DescribeContactGroupListResponseBodyContactGroupList) *DescribeContactGroupListResponseBody {
	s.ContactGroupList = v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetContactGroups(v *DescribeContactGroupListResponseBodyContactGroups) *DescribeContactGroupListResponseBody {
	s.ContactGroups = v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetMessage(v string) *DescribeContactGroupListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetRequestId(v string) *DescribeContactGroupListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetSuccess(v bool) *DescribeContactGroupListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeContactGroupListResponseBody) SetTotal(v int32) *DescribeContactGroupListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeContactGroupListResponseBody) Validate() error {
	if s.ContactGroupList != nil {
		if err := s.ContactGroupList.Validate(); err != nil {
			return err
		}
	}
	if s.ContactGroups != nil {
		if err := s.ContactGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeContactGroupListResponseBodyContactGroupList struct {
	ContactGroup []*DescribeContactGroupListResponseBodyContactGroupListContactGroup `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeContactGroupListResponseBodyContactGroupList) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupListResponseBodyContactGroupList) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBodyContactGroupList) GetContactGroup() []*DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	return s.ContactGroup
}

func (s *DescribeContactGroupListResponseBodyContactGroupList) SetContactGroup(v []*DescribeContactGroupListResponseBodyContactGroupListContactGroup) *DescribeContactGroupListResponseBodyContactGroupList {
	s.ContactGroup = v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupList) Validate() error {
	if s.ContactGroup != nil {
		for _, item := range s.ContactGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeContactGroupListResponseBodyContactGroupListContactGroup struct {
	// The alert contacts in the alert group.
	Contacts *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Struct"`
	// The time when the alert group was created. This value is a UNIX timestamp that represents the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1507070598000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the alert group.
	Describe *string `json:"Describe,omitempty" xml:"Describe,omitempty"`
	// Indicates whether the alert group subscribes to weekly reports. Valid values:
	//
	// 	- true: The alert group subscribes to weekly reports.
	//
	// 	- false: The alert group does not subscribe to weekly reports.
	//
	// example:
	//
	// true
	EnableSubscribed *bool `json:"EnableSubscribed,omitempty" xml:"EnableSubscribed,omitempty"`
	// Indicates whether the alert group can subscribe to weekly reports. Valid values:
	//
	// 	- true: The alert group can subscribe to weekly reports.
	//
	// 	- false: The alert group cannot subscribe to weekly reports.
	//
	// >  The weekly report subscription feature is only available for Alibaba Cloud accounts with more than five Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// true
	EnabledWeeklyReport *bool `json:"EnabledWeeklyReport,omitempty" xml:"EnabledWeeklyReport,omitempty"`
	// The name of the alert group.
	//
	// example:
	//
	// Contact1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The time when the alert group was modified. This value is a UNIX timestamp that represents the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1589447759000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeContactGroupListResponseBodyContactGroupListContactGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupListResponseBodyContactGroupListContactGroup) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) GetContacts() *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts {
	return s.Contacts
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) GetDescribe() *string {
	return s.Describe
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) GetEnableSubscribed() *bool {
	return s.EnableSubscribed
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) GetEnabledWeeklyReport() *bool {
	return s.EnabledWeeklyReport
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) GetName() *string {
	return s.Name
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetContacts(v *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.Contacts = v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetCreateTime(v int64) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetDescribe(v string) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.Describe = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetEnableSubscribed(v bool) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.EnableSubscribed = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetEnabledWeeklyReport(v bool) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.EnabledWeeklyReport = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetName(v string) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.Name = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) SetUpdateTime(v int64) *DescribeContactGroupListResponseBodyContactGroupListContactGroup {
	s.UpdateTime = &v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroup) Validate() error {
	if s.Contacts != nil {
		if err := s.Contacts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts struct {
	Contact []*string `json:"Contact,omitempty" xml:"Contact,omitempty" type:"Repeated"`
}

func (s DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) GetContact() []*string {
	return s.Contact
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) SetContact(v []*string) *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts {
	s.Contact = v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroupListContactGroupContacts) Validate() error {
	return dara.Validate(s)
}

type DescribeContactGroupListResponseBodyContactGroups struct {
	ContactGroup []*string `json:"ContactGroup,omitempty" xml:"ContactGroup,omitempty" type:"Repeated"`
}

func (s DescribeContactGroupListResponseBodyContactGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupListResponseBodyContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupListResponseBodyContactGroups) GetContactGroup() []*string {
	return s.ContactGroup
}

func (s *DescribeContactGroupListResponseBodyContactGroups) SetContactGroup(v []*string) *DescribeContactGroupListResponseBodyContactGroups {
	s.ContactGroup = v
	return s
}

func (s *DescribeContactGroupListResponseBodyContactGroups) Validate() error {
	return dara.Validate(s)
}
