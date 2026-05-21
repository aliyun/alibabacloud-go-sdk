// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableSupportAccountRequest interface {
  dara.Model
  String() string
  GoString() string
  SetRegionId(v string) *EnableSupportAccountRequest
  GetRegionId() *string 
  SetEnabled(v bool) *EnableSupportAccountRequest
  GetEnabled() *bool 
  SetExpireTime(v string) *EnableSupportAccountRequest
  GetExpireTime() *string 
  SetPassword(v string) *EnableSupportAccountRequest
  GetPassword() *string 
}

type EnableSupportAccountRequest struct {
  // example:
  // 
  // cn-hangzhou
  RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
  // example:
  // 
  // true
  Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
  // example:
  // 
  // 2023-03-24 11:19:05
  ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
  // example:
  // 
  // xxxx
  Password *string `json:"password,omitempty" xml:"password,omitempty"`
}

func (s EnableSupportAccountRequest) String() string {
  return dara.Prettify(s)
}

func (s EnableSupportAccountRequest) GoString() string {
  return s.String()
}

func (s *EnableSupportAccountRequest) GetRegionId() *string  {
  return s.RegionId
}

func (s *EnableSupportAccountRequest) GetEnabled() *bool  {
  return s.Enabled
}

func (s *EnableSupportAccountRequest) GetExpireTime() *string  {
  return s.ExpireTime
}

func (s *EnableSupportAccountRequest) GetPassword() *string  {
  return s.Password
}

func (s *EnableSupportAccountRequest) SetRegionId(v string) *EnableSupportAccountRequest {
  s.RegionId = &v
  return s
}

func (s *EnableSupportAccountRequest) SetEnabled(v bool) *EnableSupportAccountRequest {
  s.Enabled = &v
  return s
}

func (s *EnableSupportAccountRequest) SetExpireTime(v string) *EnableSupportAccountRequest {
  s.ExpireTime = &v
  return s
}

func (s *EnableSupportAccountRequest) SetPassword(v string) *EnableSupportAccountRequest {
  s.Password = &v
  return s
}

func (s *EnableSupportAccountRequest) Validate() error {
  return dara.Validate(s)
}

