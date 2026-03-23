// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectApgroupConfigRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApGroupUUId(v string) *EffectApgroupConfigRequest
  GetApGroupUUId() *string 
  SetAppCode(v string) *EffectApgroupConfigRequest
  GetAppCode() *string 
  SetAppName(v string) *EffectApgroupConfigRequest
  GetAppName() *string 
}

type EffectApgroupConfigRequest struct {
  // This parameter is required.
  ApGroupUUId *string `json:"ApGroupUUId,omitempty" xml:"ApGroupUUId,omitempty"`
  // This parameter is required.
  AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
  // This parameter is required.
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s EffectApgroupConfigRequest) String() string {
  return dara.Prettify(s)
}

func (s EffectApgroupConfigRequest) GoString() string {
  return s.String()
}

func (s *EffectApgroupConfigRequest) GetApGroupUUId() *string  {
  return s.ApGroupUUId
}

func (s *EffectApgroupConfigRequest) GetAppCode() *string  {
  return s.AppCode
}

func (s *EffectApgroupConfigRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EffectApgroupConfigRequest) SetApGroupUUId(v string) *EffectApgroupConfigRequest {
  s.ApGroupUUId = &v
  return s
}

func (s *EffectApgroupConfigRequest) SetAppCode(v string) *EffectApgroupConfigRequest {
  s.AppCode = &v
  return s
}

func (s *EffectApgroupConfigRequest) SetAppName(v string) *EffectApgroupConfigRequest {
  s.AppName = &v
  return s
}

func (s *EffectApgroupConfigRequest) Validate() error {
  return dara.Validate(s)
}

