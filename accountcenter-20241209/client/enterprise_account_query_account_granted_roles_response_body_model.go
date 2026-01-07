// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseAccountQueryAccountGrantedRolesResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetCode(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBody
  GetCode() *string 
  SetData(v []*EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) *EnterpriseAccountQueryAccountGrantedRolesResponseBody
  GetData() []*EnterpriseAccountQueryAccountGrantedRolesResponseBodyData 
  SetMessage(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBody
  GetMessage() *string 
  SetRequestId(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBody
  GetRequestId() *string 
  SetSuccess(v bool) *EnterpriseAccountQueryAccountGrantedRolesResponseBody
  GetSuccess() *bool 
}

type EnterpriseAccountQueryAccountGrantedRolesResponseBody struct {
  Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
  Data []*EnterpriseAccountQueryAccountGrantedRolesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
  Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponseBody) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) GetCode() *string  {
  return s.Code
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) GetData() []*EnterpriseAccountQueryAccountGrantedRolesResponseBodyData  {
  return s.Data
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) GetMessage() *string  {
  return s.Message
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) GetSuccess() *bool  {
  return s.Success
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetCode(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
  s.Code = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetData(v []*EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
  s.Data = v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetMessage(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
  s.Message = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetRequestId(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) SetSuccess(v bool) *EnterpriseAccountQueryAccountGrantedRolesResponseBody {
  s.Success = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBody) Validate() error {
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

type EnterpriseAccountQueryAccountGrantedRolesResponseBodyData struct {
  BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
  BizRoleName *string `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) GoString() string {
  return s.String()
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) GetBizRoleCode() *string  {
  return s.BizRoleCode
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) GetBizRoleName() *string  {
  return s.BizRoleName
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) SetBizRoleCode(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData {
  s.BizRoleCode = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) SetBizRoleName(v string) *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData {
  s.BizRoleName = &v
  return s
}

func (s *EnterpriseAccountQueryAccountGrantedRolesResponseBodyData) Validate() error {
  return dara.Validate(s)
}

