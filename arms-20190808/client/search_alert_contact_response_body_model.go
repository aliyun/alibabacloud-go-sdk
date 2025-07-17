// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchAlertContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *SearchAlertContactResponseBodyPageBean) *SearchAlertContactResponseBody
	GetPageBean() *SearchAlertContactResponseBodyPageBean
	SetRequestId(v string) *SearchAlertContactResponseBody
	GetRequestId() *string
}

type SearchAlertContactResponseBody struct {
	// The returned struct.
	PageBean *SearchAlertContactResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 21E85B16-75A6-429A-9F65-8AAC9A54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SearchAlertContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBody) GetPageBean() *SearchAlertContactResponseBodyPageBean {
	return s.PageBean
}

func (s *SearchAlertContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchAlertContactResponseBody) SetPageBean(v *SearchAlertContactResponseBodyPageBean) *SearchAlertContactResponseBody {
	s.PageBean = v
	return s
}

func (s *SearchAlertContactResponseBody) SetRequestId(v string) *SearchAlertContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAlertContactResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchAlertContactResponseBodyPageBean struct {
	// The information about the alert contacts.
	Contacts []*SearchAlertContactResponseBodyPageBeanContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 23
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAlertContactResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBodyPageBean) GetContacts() []*SearchAlertContactResponseBodyPageBeanContacts {
	return s.Contacts
}

func (s *SearchAlertContactResponseBodyPageBean) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchAlertContactResponseBodyPageBean) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchAlertContactResponseBodyPageBean) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SearchAlertContactResponseBodyPageBean) SetContacts(v []*SearchAlertContactResponseBodyPageBeanContacts) *SearchAlertContactResponseBodyPageBean {
	s.Contacts = v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) SetPageNumber(v int32) *SearchAlertContactResponseBodyPageBean {
	s.PageNumber = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) SetPageSize(v int32) *SearchAlertContactResponseBodyPageBean {
	s.PageSize = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) SetTotalCount(v int32) *SearchAlertContactResponseBodyPageBean {
	s.TotalCount = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBean) Validate() error {
	return dara.Validate(s)
}

type SearchAlertContactResponseBodyPageBeanContacts struct {
	// The ID of the alert contact.
	//
	// example:
	//
	// 123
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// The name of the alert contact.
	//
	// example:
	//
	// John Doe
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// The contact group to which the contact belongs. If your contacts are added to multiple contact groups, the contact groups are separated by vertical bars (|).
	//
	// example:
	//
	// Default Group | SRE Group
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The timestamp generated when the alert contact was created.
	//
	// example:
	//
	// 1572349025000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The webhook URL of the DingTalk chatbot.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=91f2f6****
	DingRobot *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
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
	// 1381111*****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Indicates whether the alert contact receives system notifications. Valid values:
	//
	// 	- `true`: The alert contact receives system notifications.
	//
	// 	- `false`: The alert contact does not receive system notifications.
	//
	// example:
	//
	// false
	SystemNoc *bool `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	// The timestamp generated when the alert contact was updated.
	//
	// example:
	//
	// 1580258717000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 113197164949****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The information about the webhook.
	//
	// example:
	//
	// {\\"body\\":\\"{   \\\\\\"msg_type\\\\\\": \\\\\\"text\\\\\\",   \\\\\\"content\\\\\\": {     \\\\\\"text\\\\\\": \\\\\\"$content\\\\\\"   } }\\",\\"header\\":{\\"Arms-Content-Type\\":\\"json\\"},\\"method\\":\\"post\\",\\"params\\":{},\\"url\\":\\"https://***",\\"userId\\":\\"1131971649***\\"}",
	Webhook *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
}

func (s SearchAlertContactResponseBodyPageBeanContacts) String() string {
	return dara.Prettify(s)
}

func (s SearchAlertContactResponseBodyPageBeanContacts) GoString() string {
	return s.String()
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetContactId() *int64 {
	return s.ContactId
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetContactName() *string {
	return s.ContactName
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetContent() *string {
	return s.Content
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetDingRobot() *string {
	return s.DingRobot
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetEmail() *string {
	return s.Email
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetPhone() *string {
	return s.Phone
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetSystemNoc() *bool {
	return s.SystemNoc
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetUserId() *string {
	return s.UserId
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) GetWebhook() *string {
	return s.Webhook
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContactId(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.ContactId = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContactName(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.ContactName = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetContent(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Content = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetCreateTime(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.CreateTime = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetDingRobot(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.DingRobot = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetEmail(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Email = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetPhone(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Phone = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetResourceGroupId(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.ResourceGroupId = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetSystemNoc(v bool) *SearchAlertContactResponseBodyPageBeanContacts {
	s.SystemNoc = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetUpdateTime(v int64) *SearchAlertContactResponseBodyPageBeanContacts {
	s.UpdateTime = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetUserId(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.UserId = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) SetWebhook(v string) *SearchAlertContactResponseBodyPageBeanContacts {
	s.Webhook = &v
	return s
}

func (s *SearchAlertContactResponseBodyPageBeanContacts) Validate() error {
	return dara.Validate(s)
}
