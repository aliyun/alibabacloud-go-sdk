// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEffectApConfigRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApMac(v string) *EffectApConfigRequest
  GetApMac() *string 
  SetAppCode(v string) *EffectApConfigRequest
  GetAppCode() *string 
  SetAppName(v string) *EffectApConfigRequest
  GetAppName() *string 
}

type EffectApConfigRequest struct {
  // This parameter is required.
  ApMac *string `json:"ApMac,omitempty" xml:"ApMac,omitempty"`
  // This parameter is required.
  AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
  // This parameter is required.
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s EffectApConfigRequest) String() string {
  return dara.Prettify(s)
}

func (s EffectApConfigRequest) GoString() string {
  return s.String()
}

func (s *EffectApConfigRequest) GetApMac() *string  {
  return s.ApMac
}

func (s *EffectApConfigRequest) GetAppCode() *string  {
  return s.AppCode
}

func (s *EffectApConfigRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EffectApConfigRequest) SetApMac(v string) *EffectApConfigRequest {
  s.ApMac = &v
  return s
}

func (s *EffectApConfigRequest) SetAppCode(v string) *EffectApConfigRequest {
  s.AppCode = &v
  return s
}

func (s *EffectApConfigRequest) SetAppName(v string) *EffectApConfigRequest {
  s.AppName = &v
  return s
}

func (s *EffectApConfigRequest) Validate() error {
  return dara.Validate(s)
}

