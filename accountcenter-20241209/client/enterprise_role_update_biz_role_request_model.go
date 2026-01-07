// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseRoleUpdateBizRoleRequest interface {
  dara.Model
  String() string
  GoString() string
  SetBizPermissionCodeListJson(v string) *EnterpriseRoleUpdateBizRoleRequest
  GetBizPermissionCodeListJson() *string 
  SetBizRoleCode(v string) *EnterpriseRoleUpdateBizRoleRequest
  GetBizRoleCode() *string 
  SetBizRoleDesc(v string) *EnterpriseRoleUpdateBizRoleRequest
  GetBizRoleDesc() *string 
  SetBizRoleName(v string) *EnterpriseRoleUpdateBizRoleRequest
  GetBizRoleName() *string 
  SetEncryptTicket(v string) *EnterpriseRoleUpdateBizRoleRequest
  GetEncryptTicket() *string 
  SetOrientedEcId(v string) *EnterpriseRoleUpdateBizRoleRequest
  GetOrientedEcId() *string 
  SetOrientedLeId(v string) *EnterpriseRoleUpdateBizRoleRequest
  GetOrientedLeId() *string 
  SetOrientedNbId(v string) *EnterpriseRoleUpdateBizRoleRequest
  GetOrientedNbId() *string 
}

type EnterpriseRoleUpdateBizRoleRequest struct {
  BizPermissionCodeListJson *string `json:"BizPermissionCodeListJson,omitempty" xml:"BizPermissionCodeListJson,omitempty"`
  BizRoleCode *string `json:"BizRoleCode,omitempty" xml:"BizRoleCode,omitempty"`
  BizRoleDesc *string `json:"BizRoleDesc,omitempty" xml:"BizRoleDesc,omitempty"`
  BizRoleName *string `json:"BizRoleName,omitempty" xml:"BizRoleName,omitempty"`
  EncryptTicket *string `json:"EncryptTicket,omitempty" xml:"EncryptTicket,omitempty"`
  OrientedEcId *string `json:"OrientedEcId,omitempty" xml:"OrientedEcId,omitempty"`
  OrientedLeId *string `json:"OrientedLeId,omitempty" xml:"OrientedLeId,omitempty"`
  OrientedNbId *string `json:"OrientedNbId,omitempty" xml:"OrientedNbId,omitempty"`
}

func (s EnterpriseRoleUpdateBizRoleRequest) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseRoleUpdateBizRoleRequest) GoString() string {
  return s.String()
}

func (s *EnterpriseRoleUpdateBizRoleRequest) GetBizPermissionCodeListJson() *string  {
  return s.BizPermissionCodeListJson
}

func (s *EnterpriseRoleUpdateBizRoleRequest) GetBizRoleCode() *string  {
  return s.BizRoleCode
}

func (s *EnterpriseRoleUpdateBizRoleRequest) GetBizRoleDesc() *string  {
  return s.BizRoleDesc
}

func (s *EnterpriseRoleUpdateBizRoleRequest) GetBizRoleName() *string  {
  return s.BizRoleName
}

func (s *EnterpriseRoleUpdateBizRoleRequest) GetEncryptTicket() *string  {
  return s.EncryptTicket
}

func (s *EnterpriseRoleUpdateBizRoleRequest) GetOrientedEcId() *string  {
  return s.OrientedEcId
}

func (s *EnterpriseRoleUpdateBizRoleRequest) GetOrientedLeId() *string  {
  return s.OrientedLeId
}

func (s *EnterpriseRoleUpdateBizRoleRequest) GetOrientedNbId() *string  {
  return s.OrientedNbId
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetBizPermissionCodeListJson(v string) *EnterpriseRoleUpdateBizRoleRequest {
  s.BizPermissionCodeListJson = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetBizRoleCode(v string) *EnterpriseRoleUpdateBizRoleRequest {
  s.BizRoleCode = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetBizRoleDesc(v string) *EnterpriseRoleUpdateBizRoleRequest {
  s.BizRoleDesc = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetBizRoleName(v string) *EnterpriseRoleUpdateBizRoleRequest {
  s.BizRoleName = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetEncryptTicket(v string) *EnterpriseRoleUpdateBizRoleRequest {
  s.EncryptTicket = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetOrientedEcId(v string) *EnterpriseRoleUpdateBizRoleRequest {
  s.OrientedEcId = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetOrientedLeId(v string) *EnterpriseRoleUpdateBizRoleRequest {
  s.OrientedLeId = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) SetOrientedNbId(v string) *EnterpriseRoleUpdateBizRoleRequest {
  s.OrientedNbId = &v
  return s
}

func (s *EnterpriseRoleUpdateBizRoleRequest) Validate() error {
  return dara.Validate(s)
}

