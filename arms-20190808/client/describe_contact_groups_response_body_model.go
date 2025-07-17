// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *DescribeContactGroupsResponseBodyPageBean) *DescribeContactGroupsResponseBody
	GetPageBean() *DescribeContactGroupsResponseBodyPageBean
	SetRequestId(v string) *DescribeContactGroupsResponseBody
	GetRequestId() *string
}

type DescribeContactGroupsResponseBody struct {
	// The objects that were returned.
	PageBean *DescribeContactGroupsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 4D6C358A-A58B-4F4B-94CE-F5AAF023****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContactGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponseBody) GetPageBean() *DescribeContactGroupsResponseBodyPageBean {
	return s.PageBean
}

func (s *DescribeContactGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContactGroupsResponseBody) SetPageBean(v *DescribeContactGroupsResponseBodyPageBean) *DescribeContactGroupsResponseBody {
	s.PageBean = v
	return s
}

func (s *DescribeContactGroupsResponseBody) SetRequestId(v string) *DescribeContactGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactGroupsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeContactGroupsResponseBodyPageBean struct {
	// The name of the alert contact group.
	AlertContactGroups []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroups `json:"AlertContactGroups,omitempty" xml:"AlertContactGroups,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of alert contact groups displayed on each page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of alert contact groups.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeContactGroupsResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponseBodyPageBean) GetAlertContactGroups() []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroups {
	return s.AlertContactGroups
}

func (s *DescribeContactGroupsResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *DescribeContactGroupsResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *DescribeContactGroupsResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeContactGroupsResponseBodyPageBean) SetAlertContactGroups(v []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) *DescribeContactGroupsResponseBodyPageBean {
	s.AlertContactGroups = v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBean) SetPage(v int64) *DescribeContactGroupsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBean) SetSize(v int64) *DescribeContactGroupsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBean) SetTotal(v int64) *DescribeContactGroupsResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBean) Validate() error {
	return dara.Validate(s)
}

type DescribeContactGroupsResponseBodyPageBeanAlertContactGroups struct {
	// The ID of the alert contact group.
	//
	// example:
	//
	// 83261
	ArmsContactGroupId *int64 `json:"ArmsContactGroupId,omitempty" xml:"ArmsContactGroupId,omitempty"`
	// The ID of the alert contact group.
	//
	// example:
	//
	// 123
	ContactGroupId *float32 `json:"ContactGroupId,omitempty" xml:"ContactGroupId,omitempty"`
	// The name of the alert contact group.
	//
	// example:
	//
	// TestGroup
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The contact information. If the **IsDetail*	- parameter is set to `false`, no **contact*	- information is displayed.
	Contacts []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
}

func (s DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) GetArmsContactGroupId() *int64 {
	return s.ArmsContactGroupId
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) GetContactGroupId() *float32 {
	return s.ContactGroupId
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) GetContacts() []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	return s.Contacts
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) SetArmsContactGroupId(v int64) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups {
	s.ArmsContactGroupId = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) SetContactGroupId(v float32) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups {
	s.ContactGroupId = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) SetContactGroupName(v string) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups {
	s.ContactGroupName = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) SetContacts(v []*DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups {
	s.Contacts = v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroups) Validate() error {
	return dara.Validate(s)
}

type DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts struct {
	// The ID of the alert contact.
	//
	// example:
	//
	// 100117
	ArmsContactId *int64 `json:"ArmsContactId,omitempty" xml:"ArmsContactId,omitempty"`
	// The ID of the alert contact.
	//
	// example:
	//
	// 456
	ContactId *float32 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// John Doe
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The email address of the alert contact.
	//
	// example:
	//
	// someone@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The mobile number of the alert contact.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) GetArmsContactId() *int64 {
	return s.ArmsContactId
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) GetContactId() *float32 {
	return s.ContactId
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) GetContactName() *string {
	return s.ContactName
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) GetEmail() *string {
	return s.Email
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) GetPhone() *string {
	return s.Phone
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) SetArmsContactId(v int64) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	s.ArmsContactId = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) SetContactId(v float32) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	s.ContactId = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) SetContactName(v string) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	s.ContactName = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) SetEmail(v string) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	s.Email = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) SetPhone(v string) *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts {
	s.Phone = &v
	return s
}

func (s *DescribeContactGroupsResponseBodyPageBeanAlertContactGroupsContacts) Validate() error {
	return dara.Validate(s)
}
