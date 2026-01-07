// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseContactQueryPageListResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseContactQueryPageListResponseBody
  GetCode() *string 
  SetData(v []*EnterpriseContactQueryPageListResponseBodyData) *EnterpriseContactQueryPageListResponseBody
  GetData() []*EnterpriseContactQueryPageListResponseBodyData 
  SetMessage(v string) *EnterpriseContactQueryPageListResponseBody
  GetMessage() *string 
  SetPageNo(v int32) *EnterpriseContactQueryPageListResponseBody
  GetPageNo() *int32 
  SetPageSize(v int32) *EnterpriseContactQueryPageListResponseBody
  GetPageSize() *int32 
  SetRequestId(v string) *EnterpriseContactQueryPageListResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseContactQueryPageListResponseBody
  GetSuccess() *bool 
  SetTotalCount(v int32) *EnterpriseContactQueryPageListResponseBody
  GetTotalCount() *int32 
  SetTotalPage(v int32) *EnterpriseContactQueryPageListResponseBody
  GetTotalPage() *int32 
}

type EnterpriseContactQueryPageListResponseBody struct {
  // example:
  // 
  // 200
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data []*EnterpriseContactQueryPageListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
  // msg
  // 
  // example:
  // 
  // successful
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  // example:
  // 
  // 1
  PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
  // example:
  // 
  // 10
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  // example:
  // 
  // C0A6196F-52A0-5EC9-B8D3-263CEF806EC4
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // example:
  // 
  // True
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  // example:
  // 
  // 11
  TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
  // example:
  // 
  // 0
  TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s EnterpriseContactQueryPageListResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactQueryPageListResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseContactQueryPageListResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseContactQueryPageListResponseBody) GetData() []*EnterpriseContactQueryPageListResponseBodyData  {
  return s.Data
}

func (s *EnterpriseContactQueryPageListResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseContactQueryPageListResponseBody) GetPageNo() *int32  {
  return s.PageNo
}

func (s *EnterpriseContactQueryPageListResponseBody) GetPageSize() *int32  {
  return s.PageSize
}

func (s *EnterpriseContactQueryPageListResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseContactQueryPageListResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseContactQueryPageListResponseBody) GetTotalCount() *int32  {
  return s.TotalCount
}

func (s *EnterpriseContactQueryPageListResponseBody) GetTotalPage() *int32  {
  return s.TotalPage
}

func (s *EnterpriseContactQueryPageListResponseBody) SetCode(v string) *EnterpriseContactQueryPageListResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBody) SetData(v []*EnterpriseContactQueryPageListResponseBodyData) *EnterpriseContactQueryPageListResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBody) SetMessage(v string) *EnterpriseContactQueryPageListResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBody) SetPageNo(v int32) *EnterpriseContactQueryPageListResponseBody {
  s.PageNo = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBody) SetPageSize(v int32) *EnterpriseContactQueryPageListResponseBody {
  s.PageSize = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBody) SetRequestId(v string) *EnterpriseContactQueryPageListResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBody) SetSuccess(v bool) *EnterpriseContactQueryPageListResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBody) SetTotalCount(v int32) *EnterpriseContactQueryPageListResponseBody {
  s.TotalCount = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBody) SetTotalPage(v int32) *EnterpriseContactQueryPageListResponseBody {
  s.TotalPage = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBody) Validate() error {
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

type EnterpriseContactQueryPageListResponseBodyData struct {
  // example:
  // 
  // xx@xx.xx
  ContactEmail *string `json:"ContactEmail,omitempty" xml:"ContactEmail,omitempty"`
  // example:
  // 
  // xxx
  ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
  // example:
  // 
  // 1xxxxxxxxxx
  ContactMobile *string `json:"ContactMobile,omitempty" xml:"ContactMobile,omitempty"`
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
  // xx
  UpdateUser *string `json:"UpdateUser,omitempty" xml:"UpdateUser,omitempty"`
}

func (s EnterpriseContactQueryPageListResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseContactQueryPageListResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetContactEmail() *string  {
  return s.ContactEmail
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetContactId() *int64  {
  return s.ContactId
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetContactMobile() *string  {
  return s.ContactMobile
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetContactName() *string  {
  return s.ContactName
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetContactPosition() *string  {
  return s.ContactPosition
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetCustomerId() *string  {
  return s.CustomerId
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetEmailConfirmed() *bool  {
  return s.EmailConfirmed
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetEntityId() *string  {
  return s.EntityId
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetEntityType() *string  {
  return s.EntityType
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetMobileConfirmed() *bool  {
  return s.MobileConfirmed
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetSharedContact() *bool  {
  return s.SharedContact
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetUpdateDate() *int64  {
  return s.UpdateDate
}

func (s *EnterpriseContactQueryPageListResponseBodyData) GetUpdateUser() *string  {
  return s.UpdateUser
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetContactEmail(v string) *EnterpriseContactQueryPageListResponseBodyData {
  s.ContactEmail = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetContactId(v int64) *EnterpriseContactQueryPageListResponseBodyData {
  s.ContactId = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetContactMobile(v string) *EnterpriseContactQueryPageListResponseBodyData {
  s.ContactMobile = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetContactName(v string) *EnterpriseContactQueryPageListResponseBodyData {
  s.ContactName = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetContactPosition(v string) *EnterpriseContactQueryPageListResponseBodyData {
  s.ContactPosition = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetCustomerId(v string) *EnterpriseContactQueryPageListResponseBodyData {
  s.CustomerId = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetEmailConfirmed(v bool) *EnterpriseContactQueryPageListResponseBodyData {
  s.EmailConfirmed = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetEntityId(v string) *EnterpriseContactQueryPageListResponseBodyData {
  s.EntityId = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetEntityType(v string) *EnterpriseContactQueryPageListResponseBodyData {
  s.EntityType = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetMobileConfirmed(v bool) *EnterpriseContactQueryPageListResponseBodyData {
  s.MobileConfirmed = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetSharedContact(v bool) *EnterpriseContactQueryPageListResponseBodyData {
  s.SharedContact = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetUpdateDate(v int64) *EnterpriseContactQueryPageListResponseBodyData {
  s.UpdateDate = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) SetUpdateUser(v string) *EnterpriseContactQueryPageListResponseBodyData {
  s.UpdateUser = &v
  return s
}

func (s *EnterpriseContactQueryPageListResponseBodyData) Validate() error {
  return dara.Validate(s)
}

