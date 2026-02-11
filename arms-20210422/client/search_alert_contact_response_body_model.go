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
	PageBean  *SearchAlertContactResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchAlertContactResponseBodyPageBean struct {
	Contacts   []*SearchAlertContactResponseBodyPageBeanContacts `json:"Contacts,omitempty" xml:"Contacts,omitempty" type:"Repeated"`
	PageNumber *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount *int32                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

type SearchAlertContactResponseBodyPageBeanContacts struct {
	ContactId   *int64  `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DingRobot   *string `json:"DingRobot,omitempty" xml:"DingRobot,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	Phone       *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	SystemNoc   *bool   `json:"SystemNoc,omitempty" xml:"SystemNoc,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId      *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Webhook     *string `json:"Webhook,omitempty" xml:"Webhook,omitempty"`
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
