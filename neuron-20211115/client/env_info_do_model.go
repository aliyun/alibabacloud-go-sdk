// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnvInfoDO interface {
  dara.Model
  String() string
  GoString() string
  SetEnv(v string) *EnvInfoDO
  GetEnv() *string 
  SetEnvType(v string) *EnvInfoDO
  GetEnvType() *string 
  SetOrgType(v string) *EnvInfoDO
  GetOrgType() *string 
  SetProductId(v int64) *EnvInfoDO
  GetProductId() *int64 
  SetProductName(v string) *EnvInfoDO
  GetProductName() *string 
  SetRegion(v string) *EnvInfoDO
  GetRegion() *string 
}

type EnvInfoDO struct {
  // example:
  // 
  // TEST
  Env *string `json:"env,omitempty" xml:"env,omitempty"`
  // example:
  // 
  // TEST
  EnvType *string `json:"envType,omitempty" xml:"envType,omitempty"`
  // example:
  // 
  // INNER
  OrgType *string `json:"orgType,omitempty" xml:"orgType,omitempty"`
  // example:
  // 
  // 231
  ProductId *int64 `json:"productId,omitempty" xml:"productId,omitempty"`
  // example:
  // 
  // yunmall
  ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
  // example:
  // 
  // cn-zhangjiakou
  Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s EnvInfoDO) String() string {
  return dara.Prettify(s)
}

func (s EnvInfoDO) GoString() string {
  return s.String()
}

func (s *EnvInfoDO) GetEnv() *string  {
  return s.Env
}

func (s *EnvInfoDO) GetEnvType() *string  {
  return s.EnvType
}

func (s *EnvInfoDO) GetOrgType() *string  {
  return s.OrgType
}

func (s *EnvInfoDO) GetProductId() *int64  {
  return s.ProductId
}

func (s *EnvInfoDO) GetProductName() *string  {
  return s.ProductName
}

func (s *EnvInfoDO) GetRegion() *string  {
  return s.Region
}

func (s *EnvInfoDO) SetEnv(v string) *EnvInfoDO {
  s.Env = &v
  return s
}

func (s *EnvInfoDO) SetEnvType(v string) *EnvInfoDO {
  s.EnvType = &v
  return s
}

func (s *EnvInfoDO) SetOrgType(v string) *EnvInfoDO {
  s.OrgType = &v
  return s
}

func (s *EnvInfoDO) SetProductId(v int64) *EnvInfoDO {
  s.ProductId = &v
  return s
}

func (s *EnvInfoDO) SetProductName(v string) *EnvInfoDO {
  s.ProductName = &v
  return s
}

func (s *EnvInfoDO) SetRegion(v string) *EnvInfoDO {
  s.Region = &v
  return s
}

func (s *EnvInfoDO) Validate() error {
  return dara.Validate(s)
}

