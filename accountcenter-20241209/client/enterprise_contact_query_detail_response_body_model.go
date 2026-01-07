// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactQueryDetailResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseContactQueryDetailResponseBody
  GetCode() *string 
  SetData(v *EnterpriseContactQueryDetailResponseBodyData) *EnterpriseContactQueryDetailResponseBody
  GetData() *EnterpriseContactQueryDetailResponseBodyData 
  SetMessage(v string) *EnterpriseContactQueryDetailResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseContactQueryDetailResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseContactQueryDetailResponseBody
  GetSuccess() *bool 
}

type EnterpriseContactQueryDetailResponseBody struct {
  // example:
  // 
  // Success
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data *EnterpriseContactQueryDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // F81F2090-8260-5052-BB93-7DF8996D25EB
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseContactQueryDetailResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactQueryDetailResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseContactQueryDetailResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseContactQueryDetailResponseBody) GetData() *EnterpriseContactQueryDetailResponseBodyData  {
  return s.Data
}

func (s *EnterpriseContactQueryDetailResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseContactQueryDetailResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseContactQueryDetailResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseContactQueryDetailResponseBody) SetCode(v string) *EnterpriseContactQueryDetailResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBody) SetData(v *EnterpriseContactQueryDetailResponseBodyData) *EnterpriseContactQueryDetailResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBody) SetMessage(v string) *EnterpriseContactQueryDetailResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBody) SetRequestId(v string) *EnterpriseContactQueryDetailResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBody) SetSuccess(v bool) *EnterpriseContactQueryDetailResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBody) Validate() error {
  if s.Data != nil {
    if err := s.Data.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseContactQueryDetailResponseBodyData struct {
  // example:
  // 
  // xx@xx.xx
  ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
  // example:
  // 
  // xx
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
  // false
  EmailConfirmed *bool `json:"EmailConfirmed,omitempty" xml:"EmailConfirmed,omitempty"`
  // example:
  // 
  // 3489d3bc-077a-449b-b41e-dd81f7451a42
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
  // false
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

func (s EnterpriseContactQueryDetailResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactQueryDetailResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetContactEmail() *string  {
  return s.ContactEmail
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetContactId() *int64  {
  return s.ContactId
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetContactMobile() *string  {
  return s.ContactMobile
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetContactName() *string  {
  return s.ContactName
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetContactPosition() *string  {
  return s.ContactPosition
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetCustomerId() *string  {
  return s.CustomerId
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetEmailConfirmed() *bool  {
  return s.EmailConfirmed
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetEntityId() *string  {
  return s.EntityId
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetEntityType() *string  {
  return s.EntityType
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetHasSubscription() *bool  {
  return s.HasSubscription
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetMobileConfirmed() *bool  {
  return s.MobileConfirmed
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetSharedContact() *bool  {
  return s.SharedContact
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetUid() *string  {
  return s.Uid
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetUpdateDate() *int64  {
  return s.UpdateDate
}

func (s *EnterpriseContactQueryDetailResponseBodyData) GetUpdateUser() *string  {
  return s.UpdateUser
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetContactEmail(v string) *EnterpriseContactQueryDetailResponseBodyData {
  s.ContactEmail = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetContactId(v int64) *EnterpriseContactQueryDetailResponseBodyData {
  s.ContactId = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetContactMobile(v string) *EnterpriseContactQueryDetailResponseBodyData {
  s.ContactMobile = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetContactName(v string) *EnterpriseContactQueryDetailResponseBodyData {
  s.ContactName = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetContactPosition(v string) *EnterpriseContactQueryDetailResponseBodyData {
  s.ContactPosition = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetCustomerId(v string) *EnterpriseContactQueryDetailResponseBodyData {
  s.CustomerId = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetEmailConfirmed(v bool) *EnterpriseContactQueryDetailResponseBodyData {
  s.EmailConfirmed = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetEntityId(v string) *EnterpriseContactQueryDetailResponseBodyData {
  s.EntityId = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetEntityType(v string) *EnterpriseContactQueryDetailResponseBodyData {
  s.EntityType = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetHasSubscription(v bool) *EnterpriseContactQueryDetailResponseBodyData {
  s.HasSubscription = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetMobileConfirmed(v bool) *EnterpriseContactQueryDetailResponseBodyData {
  s.MobileConfirmed = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetSharedContact(v bool) *EnterpriseContactQueryDetailResponseBodyData {
  s.SharedContact = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetUid(v string) *EnterpriseContactQueryDetailResponseBodyData {
  s.Uid = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetUpdateDate(v int64) *EnterpriseContactQueryDetailResponseBodyData {
  s.UpdateDate = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) SetUpdateUser(v string) *EnterpriseContactQueryDetailResponseBodyData {
  s.UpdateUser = &v
  return s
}

func (s *EnterpriseContactQueryDetailResponseBodyData) Validate() error {
  return dara.Validate(s)
}

