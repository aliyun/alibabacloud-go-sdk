// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *DescribeContactsResponseBodyPageBean) *DescribeContactsResponseBody
	GetPageBean() *DescribeContactsResponseBodyPageBean
	SetRequestId(v string) *DescribeContactsResponseBody
	GetRequestId() *string
}

type DescribeContactsResponseBody struct {
	// The objects that were returned.
	PageBean *DescribeContactsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 21E85B16-75A6-429A-9F65-8AAC9A54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeContactsResponseBody) GetPageBean() *DescribeContactsResponseBodyPageBean {
	return s.PageBean
}

func (s *DescribeContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeContactsResponseBody) SetPageBean(v *DescribeContactsResponseBodyPageBean) *DescribeContactsResponseBody {
	s.PageBean = v
	return s
}

func (s *DescribeContactsResponseBody) SetRequestId(v string) *DescribeContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeContactsResponseBody) Validate() error {
	if s.PageBean != nil {
		if err := s.PageBean.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeContactsResponseBodyPageBean struct {
	// The alert contacts.
	AlertContacts []*DescribeContactsResponseBodyPageBeanAlertContacts `json:"AlertContacts,omitempty" xml:"AlertContacts,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of alert contacts returned per page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of alert contacts.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeContactsResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *DescribeContactsResponseBodyPageBean) GetAlertContacts() []*DescribeContactsResponseBodyPageBeanAlertContacts {
	return s.AlertContacts
}

func (s *DescribeContactsResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *DescribeContactsResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *DescribeContactsResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeContactsResponseBodyPageBean) SetAlertContacts(v []*DescribeContactsResponseBodyPageBeanAlertContacts) *DescribeContactsResponseBodyPageBean {
	s.AlertContacts = v
	return s
}

func (s *DescribeContactsResponseBodyPageBean) SetPage(v int64) *DescribeContactsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBean) SetSize(v int64) *DescribeContactsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBean) SetTotal(v int64) *DescribeContactsResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBean) Validate() error {
	if s.AlertContacts != nil {
		for _, item := range s.AlertContacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeContactsResponseBodyPageBeanAlertContacts struct {
	// The ID of the alert contact.
	//
	// example:
	//
	// 100147
	ArmsContactId *int64 `json:"ArmsContactId,omitempty" xml:"ArmsContactId,omitempty"`
	// The ID of the alert contact.
	//
	// example:
	//
	// 123
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
	// Indicates whether the email address was verified.
	//
	// example:
	//
	// true
	IsEmailVerify *bool `json:"IsEmailVerify,omitempty" xml:"IsEmailVerify,omitempty"`
	// Indicates whether the mobile number was verified. Valid values:
	//
	// 	- `false`: no
	//
	// 	- `true`: yes
	//
	// example:
	//
	// false
	IsVerify *bool `json:"IsVerify,omitempty" xml:"IsVerify,omitempty"`
	// The mobile number of the alert contact.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The operation that you want to perform if phone calls fail to be answered. Valid values:
	//
	// 	- 0: No operation is performed.
	//
	// 	- 1: A phone call is made again.
	//
	// 	- 2: A text message is sent.
	//
	// 	- 3 (default value): The global default value is used.
	//
	// example:
	//
	// 3
	ReissueSendNotice *int64 `json:"ReissueSendNotice,omitempty" xml:"ReissueSendNotice,omitempty"`
}

func (s DescribeContactsResponseBodyPageBeanAlertContacts) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactsResponseBodyPageBeanAlertContacts) GoString() string {
	return s.String()
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) GetArmsContactId() *int64 {
	return s.ArmsContactId
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) GetContactId() *float32 {
	return s.ContactId
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) GetContactName() *string {
	return s.ContactName
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) GetEmail() *string {
	return s.Email
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) GetIsEmailVerify() *bool {
	return s.IsEmailVerify
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) GetIsVerify() *bool {
	return s.IsVerify
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) GetPhone() *string {
	return s.Phone
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) GetReissueSendNotice() *int64 {
	return s.ReissueSendNotice
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetArmsContactId(v int64) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.ArmsContactId = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetContactId(v float32) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.ContactId = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetContactName(v string) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.ContactName = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetEmail(v string) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.Email = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetIsEmailVerify(v bool) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.IsEmailVerify = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetIsVerify(v bool) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.IsVerify = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetPhone(v string) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.Phone = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) SetReissueSendNotice(v int64) *DescribeContactsResponseBodyPageBeanAlertContacts {
	s.ReissueSendNotice = &v
	return s
}

func (s *DescribeContactsResponseBodyPageBeanAlertContacts) Validate() error {
	return dara.Validate(s)
}
