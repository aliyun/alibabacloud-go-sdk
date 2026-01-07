// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountQueryAccountsInfoResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccountInfoDtoList(v []*EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) *EnterpriseAccountQueryAccountsInfoResponseBody
  GetAccountInfoDtoList() []*EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList 
  SetCode(v string) *EnterpriseAccountQueryAccountsInfoResponseBody
  GetCode() *string 
  SetMaxResults(v int32) *EnterpriseAccountQueryAccountsInfoResponseBody
  GetMaxResults() *int32 
  SetMessage(v string) *EnterpriseAccountQueryAccountsInfoResponseBody
  GetMessage() *string 
  SetNextToken(v string) *EnterpriseAccountQueryAccountsInfoResponseBody
  GetNextToken() *string 
  SetRequestId(v string) *EnterpriseAccountQueryAccountsInfoResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountQueryAccountsInfoResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountQueryAccountsInfoResponseBody struct {
  AccountInfoDtoList []*EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList `json:"AccountInfoDtoList,omitempty" xml:"AccountInfoDtoList,omitempty" type:"Repeated"`
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountQueryAccountsInfoResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryAccountsInfoResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) GetAccountInfoDtoList() []*EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList  {
  return s.AccountInfoDtoList
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetAccountInfoDtoList(v []*EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) *EnterpriseAccountQueryAccountsInfoResponseBody {
  s.AccountInfoDtoList = v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetCode(v string) *EnterpriseAccountQueryAccountsInfoResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetMaxResults(v int32) *EnterpriseAccountQueryAccountsInfoResponseBody {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetMessage(v string) *EnterpriseAccountQueryAccountsInfoResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetNextToken(v string) *EnterpriseAccountQueryAccountsInfoResponseBody {
  s.NextToken = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetRequestId(v string) *EnterpriseAccountQueryAccountsInfoResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) SetSuccess(v bool) *EnterpriseAccountQueryAccountsInfoResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBody) Validate() error {
  if s.AccountInfoDtoList != nil {
    for _, item := range s.AccountInfoDtoList {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList struct {
  Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
  BelongToMasterAccount *bool `json:"BelongToMasterAccount,omitempty" xml:"BelongToMasterAccount,omitempty"`
  Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
  EnterpriseEcId *string `json:"EnterpriseEcId,omitempty" xml:"EnterpriseEcId,omitempty"`
  EnterpriseEntityId *string `json:"EnterpriseEntityId,omitempty" xml:"EnterpriseEntityId,omitempty"`
  EnterpriseLicenseNo *string `json:"EnterpriseLicenseNo,omitempty" xml:"EnterpriseLicenseNo,omitempty"`
  EnterpriseName *string `json:"EnterpriseName,omitempty" xml:"EnterpriseName,omitempty"`
  EnterpriseNbId *string `json:"EnterpriseNbId,omitempty" xml:"EnterpriseNbId,omitempty"`
  FreezeLogin *bool `json:"FreezeLogin,omitempty" xml:"FreezeLogin,omitempty"`
  LoginId *string `json:"LoginId,omitempty" xml:"LoginId,omitempty"`
  ManageInviteTimeStamp *string `json:"ManageInviteTimeStamp,omitempty" xml:"ManageInviteTimeStamp,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
  SecurityMobile *string `json:"SecurityMobile,omitempty" xml:"SecurityMobile,omitempty"`
}

func (s EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetAlias() *string  {
  return s.Alias
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetBelongToMasterAccount() *bool  {
  return s.BelongToMasterAccount
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetEmail() *string  {
  return s.Email
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetEnterpriseEcId() *string  {
  return s.EnterpriseEcId
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetEnterpriseEntityId() *string  {
  return s.EnterpriseEntityId
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetEnterpriseLicenseNo() *string  {
  return s.EnterpriseLicenseNo
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetEnterpriseName() *string  {
  return s.EnterpriseName
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetEnterpriseNbId() *string  {
  return s.EnterpriseNbId
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetFreezeLogin() *bool  {
  return s.FreezeLogin
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetLoginId() *string  {
  return s.LoginId
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetManageInviteTimeStamp() *string  {
  return s.ManageInviteTimeStamp
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) GetSecurityMobile() *string  {
  return s.SecurityMobile
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetAlias(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.Alias = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetBelongToMasterAccount(v bool) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.BelongToMasterAccount = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEmail(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.Email = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseEcId(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.EnterpriseEcId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseEntityId(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.EnterpriseEntityId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseLicenseNo(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.EnterpriseLicenseNo = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseName(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.EnterpriseName = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetEnterpriseNbId(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.EnterpriseNbId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetFreezeLogin(v bool) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.FreezeLogin = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetLoginId(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.LoginId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetManageInviteTimeStamp(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.ManageInviteTimeStamp = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetPk(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.Pk = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) SetSecurityMobile(v string) *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList {
  s.SecurityMobile = &v
  return s
}

func (s *EnterpriseAccountQueryAccountsInfoResponseBodyAccountInfoDtoList) Validate() error {
  return dara.Validate(s)
}

