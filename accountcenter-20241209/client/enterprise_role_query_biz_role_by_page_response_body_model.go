// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleQueryBizRoleByPageResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetBizRoles(v []*EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetBizRoles() []*EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles 
  SetCode(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetCode() *string 
  SetMaxResults(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetMaxResults() *int32 
  SetMessage(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetMessage() *string 
  SetNextToken(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetNextToken() *string 
  SetPageNo(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetPageNo() *int32 
  SetPageSize(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetPageSize() *int32 
  SetRequestId(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetSuccess() *bool 
  SetTotalCount(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetTotalCount() *int32 
  SetTotalPage(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody
  GetTotalPage() *int32 
}

type EnterpriseRoleQueryBizRoleByPageResponseBody struct {
  BizRoles []*EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles `json:"BizRoles,omitempty" xml:"BizRoles,omitempty" type:"Repeated"`
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

func (s EnterpriseRoleQueryBizRoleByPageResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleByPageResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetBizRoles() []*EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles  {
  return s.BizRoles
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetMaxResults() *int32  {
  return s.MaxResults
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetNextToken() *string  {
  return s.NextToken
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetPageNo() *int32  {
  return s.PageNo
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetPageSize() *int32  {
  return s.PageSize
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetTotalCount() *int32  {
  return s.TotalCount
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) GetTotalPage() *int32  {
  return s.TotalPage
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetBizRoles(v []*EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.BizRoles = v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetCode(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetMaxResults(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.MaxResults = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetMessage(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetNextToken(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.NextToken = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetPageNo(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.PageNo = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetPageSize(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.PageSize = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetRequestId(v string) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetSuccess(v bool) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetTotalCount(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.TotalCount = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) SetTotalPage(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBody {
  s.TotalPage = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBody) Validate() error {
  if s.BizRoles != nil {
    for _, item := range s.BizRoles {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles struct {
  BizPermissionCount *int32 `json:"BizPermissionCount,omitempty" xml:"BizPermissionCount,omitempty"`
  BizPermissionNameList []*string `json:"BizPermissionNameList,omitempty" xml:"BizPermissionNameList,omitempty" type:"Repeated"`
  BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
  BizRoleDesc *string `json:"BizRoleDesc,omitempty" xml:"BizRoleDesc,omitempty"`
  BizRoleName *string `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
  GrantedPkCount *int32 `json:"GrantedPkCount,omitempty" xml:"GrantedPkCount,omitempty"`
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GetBizPermissionCount() *int32  {
  return s.BizPermissionCount
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GetBizPermissionNameList() []*string  {
  return s.BizPermissionNameList
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GetBizRoleCode() *string  {
  return s.BizRoleCode
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GetBizRoleDesc() *string  {
  return s.BizRoleDesc
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GetBizRoleName() *string  {
  return s.BizRoleName
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GetGrantedPkCount() *int32  {
  return s.GrantedPkCount
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) GetSrcType() *string  {
  return s.SrcType
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizPermissionCount(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
  s.BizPermissionCount = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizPermissionNameList(v []*string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
  s.BizPermissionNameList = v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizRoleCode(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
  s.BizRoleCode = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizRoleDesc(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
  s.BizRoleDesc = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetBizRoleName(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
  s.BizRoleName = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetGrantedPkCount(v int32) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
  s.GrantedPkCount = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetResourceType(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
  s.ResourceType = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) SetSrcType(v string) *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles {
  s.SrcType = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleByPageResponseBodyBizRoles) Validate() error {
  return dara.Validate(s)
}

