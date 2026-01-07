// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleQueryBizRoleDetailResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetBizRoleDetail(v *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) *EnterpriseRoleQueryBizRoleDetailResponseBody
  GetBizRoleDetail() *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail 
  SetCode(v string) *EnterpriseRoleQueryBizRoleDetailResponseBody
  GetCode() *string 
  SetMessage(v string) *EnterpriseRoleQueryBizRoleDetailResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseRoleQueryBizRoleDetailResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseRoleQueryBizRoleDetailResponseBody
  GetSuccess() *bool 
}

type EnterpriseRoleQueryBizRoleDetailResponseBody struct {
  BizRoleDetail *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail `json:"BizRoleDetail,omitempty" xml:"BizRoleDetail,omitempty" type:"Struct"`
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) GetBizRoleDetail() *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail  {
  return s.BizRoleDetail
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetBizRoleDetail(v *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) *EnterpriseRoleQueryBizRoleDetailResponseBody {
  s.BizRoleDetail = v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetCode(v string) *EnterpriseRoleQueryBizRoleDetailResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetMessage(v string) *EnterpriseRoleQueryBizRoleDetailResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetRequestId(v string) *EnterpriseRoleQueryBizRoleDetailResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) SetSuccess(v bool) *EnterpriseRoleQueryBizRoleDetailResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBody) Validate() error {
  if s.BizRoleDetail != nil {
    if err := s.BizRoleDetail.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail struct {
  BizPermissions []*EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions `json:"BizPermissions,omitempty" xml:"BizPermissions,omitempty" type:"Repeated"`
  BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
  BizRoleDesc *string `json:"BizRoleDesc,omitempty" xml:"BizRoleDesc,omitempty"`
  BizRoleName *string `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
  SrcType *string `json:"SrcType,omitempty" xml:"SrcType,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) GetBizPermissions() []*EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions  {
  return s.BizPermissions
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) GetBizRoleCode() *string  {
  return s.BizRoleCode
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) GetBizRoleDesc() *string  {
  return s.BizRoleDesc
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) GetBizRoleName() *string  {
  return s.BizRoleName
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) GetSrcType() *string  {
  return s.SrcType
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetBizPermissions(v []*EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
  s.BizPermissions = v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetBizRoleCode(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
  s.BizRoleCode = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetBizRoleDesc(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
  s.BizRoleDesc = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetBizRoleName(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
  s.BizRoleName = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetResourceType(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
  s.ResourceType = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) SetSrcType(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail {
  s.SrcType = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetail) Validate() error {
  if s.BizPermissions != nil {
    for _, item := range s.BizPermissions {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions struct {
  BizPermissionCode *string `json:"BizPermissionCode,omitempty" xml:"BizPermissionCode,omitempty"`
  BizPermissionDesc *string `json:"BizPermissionDesc,omitempty" xml:"BizPermissionDesc,omitempty"`
  BizPermissionName *string `json:"BizPermissionName,omitempty" xml:"BizPermissionName,omitempty"`
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) GetBizPermissionCode() *string  {
  return s.BizPermissionCode
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) GetBizPermissionDesc() *string  {
  return s.BizPermissionDesc
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) GetBizPermissionName() *string  {
  return s.BizPermissionName
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) SetBizPermissionCode(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions {
  s.BizPermissionCode = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) SetBizPermissionDesc(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions {
  s.BizPermissionDesc = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) SetBizPermissionName(v string) *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions {
  s.BizPermissionName = &v
  return s
}

func (s *EnterpriseRoleQueryBizRoleDetailResponseBodyBizRoleDetailBizPermissions) Validate() error {
  return dara.Validate(s)
}

