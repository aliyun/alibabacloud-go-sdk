// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleQueryAccountForRoleGrantByPageResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetAccounts(v []*EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetAccounts() []*EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts 
  SetCode(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetCode() *string 
  SetMaxResults(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetMaxResults() *int32 
  SetMessage(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetMessage() *string 
  SetNextToken(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetNextToken() *string 
  SetPageNo(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetPageNo() *int32 
  SetPageSize(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetPageSize() *int32 
  SetRequestId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetSuccess() *bool 
  SetTotalCount(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetTotalCount() *int32 
  SetTotalPage(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody
  GetTotalPage() *int32 
}

type EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody struct {
  Accounts []*EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Repeated"`
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
  PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
  PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
  TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
  TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetAccounts() []*EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts  {
  return s.Accounts
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetPageNo() *int32  {
  return s.PageNo
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetPageSize() *int32  {
  return s.PageSize
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetTotalCount() *int32  {
  return s.TotalCount
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) GetTotalPage() *int32  {
  return s.TotalPage
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetAccounts(v []*EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.Accounts = v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetCode(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetMaxResults(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetMessage(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetNextToken(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.NextToken = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetPageNo(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.PageNo = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetPageSize(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.PageSize = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetRequestId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetSuccess(v bool) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetTotalCount(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.TotalCount = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) SetTotalPage(v int32) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody {
  s.TotalPage = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBody) Validate() error {
  if s.Accounts != nil {
    for _, item := range s.Accounts {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts struct {
  Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
  AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
  Granted *bool `json:"Granted,omitempty" xml:"Granted,omitempty"`
  Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) GetAlias() *string  {
  return s.Alias
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) GetAliyunId() *string  {
  return s.AliyunId
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) GetGranted() *bool  {
  return s.Granted
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) GetPk() *string  {
  return s.Pk
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) SetAlias(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts {
  s.Alias = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) SetAliyunId(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts {
  s.AliyunId = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) SetGranted(v bool) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts {
  s.Granted = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) SetPk(v string) *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts {
  s.Pk = &v
  return s
}

func (s *EnterpriseRoleQueryAccountForRoleGrantByPageResponseBodyAccounts) Validate() error {
  return dara.Validate(s)
}

