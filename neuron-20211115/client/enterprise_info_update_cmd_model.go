// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnterpriseInfoUpdateCmd interface {
  dara.Model
  String() string
  GoString() string
  SetDescription(v string) *EnterpriseInfoUpdateCmd
  GetDescription() *string 
  SetId(v int64) *EnterpriseInfoUpdateCmd
  GetId() *int64 
  SetName(v string) *EnterpriseInfoUpdateCmd
  GetName() *string 
}

type EnterpriseInfoUpdateCmd struct {
  // example:
  // 
  // 互联网企业
  Description *string `json:"description,omitempty" xml:"description,omitempty"`
  // This parameter is required.
  // 
  // example:
  // 
  // 1
  Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
  // example:
  // 
  // 阿里
  Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s EnterpriseInfoUpdateCmd) String() string {
  return dara.Prettify(s)
}

func (s EnterpriseInfoUpdateCmd) GoString() string {
  return s.String()
}

func (s *EnterpriseInfoUpdateCmd) GetDescription() *string  {
  return s.Description
}

func (s *EnterpriseInfoUpdateCmd) GetId() *int64  {
  return s.Id
}

func (s *EnterpriseInfoUpdateCmd) GetName() *string  {
  return s.Name
}

func (s *EnterpriseInfoUpdateCmd) SetDescription(v string) *EnterpriseInfoUpdateCmd {
  s.Description = &v
  return s
}

func (s *EnterpriseInfoUpdateCmd) SetId(v int64) *EnterpriseInfoUpdateCmd {
  s.Id = &v
  return s
}

func (s *EnterpriseInfoUpdateCmd) SetName(v string) *EnterpriseInfoUpdateCmd {
  s.Name = &v
  return s
}

func (s *EnterpriseInfoUpdateCmd) Validate() error {
  return dara.Validate(s)
}

