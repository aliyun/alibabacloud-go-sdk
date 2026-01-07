// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAccountContactQueryDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AccountContactQueryDetailResponseBody
	GetCode() *string
	SetData(v *AccountContactQueryDetailResponseBodyData) *AccountContactQueryDetailResponseBody
	GetData() *AccountContactQueryDetailResponseBodyData
	SetMessage(v string) *AccountContactQueryDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *AccountContactQueryDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AccountContactQueryDetailResponseBody
	GetSuccess() *bool
}

type AccountContactQueryDetailResponseBody struct {
	// example:
	//
	// 200
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *AccountContactQueryDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DA632E90-32DB-52DE-823B-4A182169D954
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AccountContactQueryDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AccountContactQueryDetailResponseBody) GoString() string {
	return s.String()
}

func (s *AccountContactQueryDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *AccountContactQueryDetailResponseBody) GetData() *AccountContactQueryDetailResponseBodyData {
	return s.Data
}

func (s *AccountContactQueryDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AccountContactQueryDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AccountContactQueryDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AccountContactQueryDetailResponseBody) SetCode(v string) *AccountContactQueryDetailResponseBody {
	s.Code = &v
	return s
}

func (s *AccountContactQueryDetailResponseBody) SetData(v *AccountContactQueryDetailResponseBodyData) *AccountContactQueryDetailResponseBody {
	s.Data = v
	return s
}

func (s *AccountContactQueryDetailResponseBody) SetMessage(v string) *AccountContactQueryDetailResponseBody {
	s.Message = &v
	return s
}

func (s *AccountContactQueryDetailResponseBody) SetRequestId(v string) *AccountContactQueryDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *AccountContactQueryDetailResponseBody) SetSuccess(v bool) *AccountContactQueryDetailResponseBody {
	s.Success = &v
	return s
}

func (s *AccountContactQueryDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AccountContactQueryDetailResponseBodyData struct {
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
	// 4
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
	// xxx
	HasSubscription *bool `json:"HasSubscription,omitempty" xml:"HasSubscription,omitempty"`
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
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
	// example:
	//
	// xxx
	UpdateDate *int64 `json:"UpdateDate,omitempty" xml:"UpdateDate,omitempty"`
	// example:
	//
	// xxx
	UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s AccountContactQueryDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AccountContactQueryDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *AccountContactQueryDetailResponseBodyData) GetContactEmail() *string {
	return s.ContactEmail
}

func (s *AccountContactQueryDetailResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *AccountContactQueryDetailResponseBodyData) GetContactMobile() *string {
	return s.ContactMobile
}

func (s *AccountContactQueryDetailResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *AccountContactQueryDetailResponseBodyData) GetContactPosition() *string {
	return s.ContactPosition
}

func (s *AccountContactQueryDetailResponseBodyData) GetCustomerId() *string {
	return s.CustomerId
}

func (s *AccountContactQueryDetailResponseBodyData) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *AccountContactQueryDetailResponseBodyData) GetEntityId() *string {
	return s.EntityId
}

func (s *AccountContactQueryDetailResponseBodyData) GetEntityType() *string {
	return s.EntityType
}

func (s *AccountContactQueryDetailResponseBodyData) GetHasSubscription() *bool {
	return s.HasSubscription
}

func (s *AccountContactQueryDetailResponseBodyData) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *AccountContactQueryDetailResponseBodyData) GetSharedContact() *bool {
	return s.SharedContact
}

func (s *AccountContactQueryDetailResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *AccountContactQueryDetailResponseBodyData) GetUpdateDate() *int64 {
	return s.UpdateDate
}

func (s *AccountContactQueryDetailResponseBodyData) GetUpdateUser() *string {
	return s.UpdateUser
}

func (s *AccountContactQueryDetailResponseBodyData) SetContactEmail(v string) *AccountContactQueryDetailResponseBodyData {
	s.ContactEmail = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetContactId(v int64) *AccountContactQueryDetailResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetContactMobile(v string) *AccountContactQueryDetailResponseBodyData {
	s.ContactMobile = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetContactName(v string) *AccountContactQueryDetailResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetContactPosition(v string) *AccountContactQueryDetailResponseBodyData {
	s.ContactPosition = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetCustomerId(v string) *AccountContactQueryDetailResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetEmailConfirmed(v bool) *AccountContactQueryDetailResponseBodyData {
	s.EmailConfirmed = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetEntityId(v string) *AccountContactQueryDetailResponseBodyData {
	s.EntityId = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetEntityType(v string) *AccountContactQueryDetailResponseBodyData {
	s.EntityType = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetHasSubscription(v bool) *AccountContactQueryDetailResponseBodyData {
	s.HasSubscription = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetMobileConfirmed(v bool) *AccountContactQueryDetailResponseBodyData {
	s.MobileConfirmed = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetSharedContact(v bool) *AccountContactQueryDetailResponseBodyData {
	s.SharedContact = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetUid(v string) *AccountContactQueryDetailResponseBodyData {
	s.Uid = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetUpdateDate(v int64) *AccountContactQueryDetailResponseBodyData {
	s.UpdateDate = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) SetUpdateUser(v string) *AccountContactQueryDetailResponseBodyData {
	s.UpdateUser = &v
	return s
}

func (s *AccountContactQueryDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
