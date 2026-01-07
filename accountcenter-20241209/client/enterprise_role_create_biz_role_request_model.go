// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleCreateBizRoleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizPermissionCodeListJson(v string) *EnterpriseRoleCreateBizRoleRequest
  GetBizPermissionCodeListJson() *string 
  SetBizRoleDesc(v string) *EnterpriseRoleCreateBizRoleRequest
  GetBizRoleDesc() *string 
  SetBizRoleName(v string) *EnterpriseRoleCreateBizRoleRequest
  GetBizRoleName() *string 
  SetEncryptTicket(v string) *EnterpriseRoleCreateBizRoleRequest
  GetEncryptTicket() *string 
  SetOrientedEcId(v string) *EnterpriseRoleCreateBizRoleRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseRoleCreateBizRoleRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseRoleCreateBizRoleRequest
  GetOrientedNbId() *string 
  SetResourceType(v string) *EnterpriseRoleCreateBizRoleRequest
  GetResourceType() *string 
}

type EnterpriseRoleCreateBizRoleRequest struct {
  BizPermissionCodeListJson *string `json:"BizPermissionCodeListJson,omitempty" xml:"BizPermissionCodeListJson,omitempty"`
  BizRoleDesc *string `json:"BizRoleDesc,omitempty" xml:"BizRoleDesc,omitempty"`
  BizRoleName *string `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
  ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s EnterpriseRoleCreateBizRoleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleCreateBizRoleRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleCreateBizRoleRequest) GetBizPermissionCodeListJson() *string  {
  return s.BizPermissionCodeListJson
}

func (s *EnterpriseRoleCreateBizRoleRequest) GetBizRoleDesc() *string  {
  return s.BizRoleDesc
}

func (s *EnterpriseRoleCreateBizRoleRequest) GetBizRoleName() *string  {
  return s.BizRoleName
}

func (s *EnterpriseRoleCreateBizRoleRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseRoleCreateBizRoleRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseRoleCreateBizRoleRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseRoleCreateBizRoleRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseRoleCreateBizRoleRequest) GetResourceType() *string  {
  return s.ResourceType
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetBizPermissionCodeListJson(v string) *EnterpriseRoleCreateBizRoleRequest {
  s.BizPermissionCodeListJson = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetBizRoleDesc(v string) *EnterpriseRoleCreateBizRoleRequest {
  s.BizRoleDesc = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetBizRoleName(v string) *EnterpriseRoleCreateBizRoleRequest {
  s.BizRoleName = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetEncryptTicket(v string) *EnterpriseRoleCreateBizRoleRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetOrientedEcId(v string) *EnterpriseRoleCreateBizRoleRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetOrientedLeId(v string) *EnterpriseRoleCreateBizRoleRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetOrientedNbId(v string) *EnterpriseRoleCreateBizRoleRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) SetResourceType(v string) *EnterpriseRoleCreateBizRoleRequest {
  s.ResourceType = &v
  return s
}

func (s *EnterpriseRoleCreateBizRoleRequest) Validate() error {
  return dara.Validate(s)
}

