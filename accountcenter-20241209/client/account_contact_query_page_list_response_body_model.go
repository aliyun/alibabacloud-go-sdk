// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactQueryPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AccountContactQueryPageListResponseBody
	GetCode() *string
	SetData(v []*AccountContactQueryPageListResponseBodyData) *AccountContactQueryPageListResponseBody
	GetData() []*AccountContactQueryPageListResponseBodyData
	SetMessage(v string) *AccountContactQueryPageListResponseBody
	GetMessage() *string
	SetPageNo(v int32) *AccountContactQueryPageListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *AccountContactQueryPageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *AccountContactQueryPageListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AccountContactQueryPageListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *AccountContactQueryPageListResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *AccountContactQueryPageListResponseBody
	GetTotalPage() *int32
}

type AccountContactQueryPageListResponseBody struct {
	// example:
	//
	// 200
	Code *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*AccountContactQueryPageListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 19
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 27A90A55-120A-59DC-812E-62448D440E95
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s AccountContactQueryPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AccountContactQueryPageListResponseBody) GoString() string {
	return s.String()
}

func (s *AccountContactQueryPageListResponseBody) GetCode() *string {
	return s.Code
}

func (s *AccountContactQueryPageListResponseBody) GetData() []*AccountContactQueryPageListResponseBodyData {
	return s.Data
}

func (s *AccountContactQueryPageListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AccountContactQueryPageListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *AccountContactQueryPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *AccountContactQueryPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AccountContactQueryPageListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AccountContactQueryPageListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AccountContactQueryPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *AccountContactQueryPageListResponseBody) SetCode(v string) *AccountContactQueryPageListResponseBody {
	s.Code = &v
	return s
}

func (s *AccountContactQueryPageListResponseBody) SetData(v []*AccountContactQueryPageListResponseBodyData) *AccountContactQueryPageListResponseBody {
	s.Data = v
	return s
}

func (s *AccountContactQueryPageListResponseBody) SetMessage(v string) *AccountContactQueryPageListResponseBody {
	s.Message = &v
	return s
}

func (s *AccountContactQueryPageListResponseBody) SetPageNo(v int32) *AccountContactQueryPageListResponseBody {
	s.PageNo = &v
	return s
}

func (s *AccountContactQueryPageListResponseBody) SetPageSize(v int32) *AccountContactQueryPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *AccountContactQueryPageListResponseBody) SetRequestId(v string) *AccountContactQueryPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccountContactQueryPageListResponseBody) SetSuccess(v bool) *AccountContactQueryPageListResponseBody {
	s.Success = &v
	return s
}

func (s *AccountContactQueryPageListResponseBody) SetTotalCount(v int32) *AccountContactQueryPageListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *AccountContactQueryPageListResponseBody) SetTotalPage(v int32) *AccountContactQueryPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *AccountContactQueryPageListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AccountContactQueryPageListResponseBodyData struct {
	// example:
	//
	// xxx@xxx.xx
	ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
	// example:
	//
	// xxx
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// 1xxxxxxxxxx
	ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
	// example:
	//
	// xxx
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// 1
	ContactPosition *string `json:"ContactPosition,omitempty" xml:"ContactPosition,omitempty"`
	// example:
	//
	// xxx
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// example:
	//
	// true
	EmailConfirmed *bool `json:"EmailConfirmed,omitempty" xml:"EmailConfirmed,omitempty"`
	// example:
	//
	// xxx
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// leId/customerId
	//
	// example:
	//
	// customerId
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// true
	MobileConfirmed *bool `json:"MobileConfirmed,omitempty" xml:"MobileConfirmed,omitempty"`
	// example:
	//
	// false
	SharedContact *bool `json:"SharedContact,omitempty" xml:"SharedContact,omitempty"`
	// example:
	//
	// xxx
	UpdateDate *int64 `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// example:
	//
	// xxx
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s AccountContactQueryPageListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AccountContactQueryPageListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccountContactQueryPageListResponseBodyData) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *AccountContactQueryPageListResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *AccountContactQueryPageListResponseBodyData) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *AccountContactQueryPageListResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *AccountContactQueryPageListResponseBodyData) GetContactPosition() *string {
	return s.ContactPosition
}

func (s *AccountContactQueryPageListResponseBodyData) GetCustomerId() *string {
	return s.CustomerId
}

func (s *AccountContactQueryPageListResponseBodyData) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *AccountContactQueryPageListResponseBodyData) GetEntityId() *string {
	return s.EntityId
}

func (s *AccountContactQueryPageListResponseBodyData) GetEntityType() *string {
	return s.EntityType
}

func (s *AccountContactQueryPageListResponseBodyData) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *AccountContactQueryPageListResponseBodyData) GetSharedContact() *bool {
	return s.SharedContact
}

func (s *AccountContactQueryPageListResponseBodyData) GetUpdateDate() *int64 {
	return s.UpdateDate
}

func (s *AccountContactQueryPageListResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *AccountContactQueryPageListResponseBodyData) SetContactEmail(v string) *AccountContactQueryPageListResponseBodyData {
	s.ContactEmail = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetContactId(v int64) *AccountContactQueryPageListResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetContactMobile(v string) *AccountContactQueryPageListResponseBodyData {
	s.ContactMobile = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetContactName(v string) *AccountContactQueryPageListResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetContactPosition(v string) *AccountContactQueryPageListResponseBodyData {
	s.ContactPosition = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetCustomerId(v string) *AccountContactQueryPageListResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetEmailConfirmed(v bool) *AccountContactQueryPageListResponseBodyData {
	s.EmailConfirmed = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetEntityId(v string) *AccountContactQueryPageListResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetEntityType(v string) *AccountContactQueryPageListResponseBodyData {
	s.EntityType = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetMobileConfirmed(v bool) *AccountContactQueryPageListResponseBodyData {
	s.MobileConfirmed = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetSharedContact(v bool) *AccountContactQueryPageListResponseBodyData {
	s.SharedContact = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetUpdateDate(v int64) *AccountContactQueryPageListResponseBodyData {
	s.UpdateDate = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) SetUpdateUser(v string) *AccountContactQueryPageListResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *AccountContactQueryPageListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
