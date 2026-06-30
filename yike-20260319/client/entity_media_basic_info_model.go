// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEntityMediaBasicInfo interface {
  dara.Model
  String() string
  GoString() string
  SetAppId(v string) *EntityMediaBasicInfo
  GetAppId() *string 
  SetBiz(v string) *EntityMediaBasicInfo
  GetBiz() *string 
  SetCreateTime(v string) *EntityMediaBasicInfo
  GetCreateTime() *string 
  SetEntityId(v string) *EntityMediaBasicInfo
  GetEntityId() *string 
  SetEntityMediaId(v string) *EntityMediaBasicInfo
  GetEntityMediaId() *string 
  SetModifiedTime(v string) *EntityMediaBasicInfo
  GetModifiedTime() *string 
  SetStatus(v string) *EntityMediaBasicInfo
  GetStatus() *string 
  SetUserData(v string) *EntityMediaBasicInfo
  GetUserData() *string 
}

type EntityMediaBasicInfo struct {
  AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
  Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
  CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
  EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
  EntityMediaId *string `json:"EntityMediaId,omitempty" xml:"EntityMediaId,omitempty"`
  ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
  Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
  UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s EntityMediaBasicInfo) String() string {
  return dara.Prettify(s)
}

func (s EntityMediaBasicInfo) GoString() string {
  return s.String()
}

func (s *EntityMediaBasicInfo) GetAppId() *string  {
  return s.AppId
}

func (s *EntityMediaBasicInfo) GetBiz() *string  {
  return s.Biz
}

func (s *EntityMediaBasicInfo) GetCreateTime() *string  {
  return s.CreateTime
}

func (s *EntityMediaBasicInfo) GetEntityId() *string  {
  return s.EntityId
}

func (s *EntityMediaBasicInfo) GetEntityMediaId() *string  {
  return s.EntityMediaId
}

func (s *EntityMediaBasicInfo) GetModifiedTime() *string  {
  return s.ModifiedTime
}

func (s *EntityMediaBasicInfo) GetStatus() *string  {
  return s.Status
}

func (s *EntityMediaBasicInfo) GetUserData() *string  {
  return s.UserData
}

func (s *EntityMediaBasicInfo) SetAppId(v string) *EntityMediaBasicInfo {
  s.AppId = &v
  return s
}

func (s *EntityMediaBasicInfo) SetBiz(v string) *EntityMediaBasicInfo {
  s.Biz = &v
  return s
}

func (s *EntityMediaBasicInfo) SetCreateTime(v string) *EntityMediaBasicInfo {
  s.CreateTime = &v
  return s
}

func (s *EntityMediaBasicInfo) SetEntityId(v string) *EntityMediaBasicInfo {
  s.EntityId = &v
  return s
}

func (s *EntityMediaBasicInfo) SetEntityMediaId(v string) *EntityMediaBasicInfo {
  s.EntityMediaId = &v
  return s
}

func (s *EntityMediaBasicInfo) SetModifiedTime(v string) *EntityMediaBasicInfo {
  s.ModifiedTime = &v
  return s
}

func (s *EntityMediaBasicInfo) SetStatus(v string) *EntityMediaBasicInfo {
  s.Status = &v
  return s
}

func (s *EntityMediaBasicInfo) SetUserData(v string) *EntityMediaBasicInfo {
  s.UserData = &v
  return s
}

func (s *EntityMediaBasicInfo) Validate() error {
  return dara.Validate(s)
}

