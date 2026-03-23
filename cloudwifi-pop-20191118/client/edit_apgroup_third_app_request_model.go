// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditApgroupThirdAppRequest interface {
  dara.Model
  String() string
  GoString() string
  SetApgroupId(v int64) *EditApgroupThirdAppRequest
  GetApgroupId() *int64 
  SetAppCode(v string) *EditApgroupThirdAppRequest
  GetAppCode() *string 
  SetAppData(v string) *EditApgroupThirdAppRequest
  GetAppData() *string 
  SetAppName(v string) *EditApgroupThirdAppRequest
  GetAppName() *string 
  SetApplyToSubGroup(v int32) *EditApgroupThirdAppRequest
  GetApplyToSubGroup() *int32 
  SetCategory(v int32) *EditApgroupThirdAppRequest
  GetCategory() *int32 
  SetConfigType(v string) *EditApgroupThirdAppRequest
  GetConfigType() *string 
  SetId(v int64) *EditApgroupThirdAppRequest
  GetId() *int64 
  SetInheritParentGroup(v int32) *EditApgroupThirdAppRequest
  GetInheritParentGroup() *int32 
  SetThirdAppName(v string) *EditApgroupThirdAppRequest
  GetThirdAppName() *string 
}

type EditApgroupThirdAppRequest struct {
  // This parameter is required.
  ApgroupId *int64 `json:"ApgroupId,omitempty" xml:"ApgroupId,omitempty"`
  // This parameter is required.
  AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
  AppData *string `json:"AppData,omitempty" xml:"AppData,omitempty"`
  // This parameter is required.
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  ApplyToSubGroup *int32 `json:"ApplyToSubGroup,omitempty" xml:"ApplyToSubGroup,omitempty"`
  Category *int32 `json:"Category,omitempty" xml:"Category,omitempty"`
  ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
  Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
  InheritParentGroup *int32 `json:"InheritParentGroup,omitempty" xml:"InheritParentGroup,omitempty"`
  ThirdAppName *string `json:"ThirdAppName,omitempty" xml:"ThirdAppName,omitempty"`
}

func (s EditApgroupThirdAppRequest) String() string {
  return dara.Prettify(s)
}

func (s EditApgroupThirdAppRequest) GoString() string {
  return s.String()
}

func (s *EditApgroupThirdAppRequest) GetApgroupId() *int64  {
  return s.ApgroupId
}

func (s *EditApgroupThirdAppRequest) GetAppCode() *string  {
  return s.AppCode
}

func (s *EditApgroupThirdAppRequest) GetAppData() *string  {
  return s.AppData
}

func (s *EditApgroupThirdAppRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EditApgroupThirdAppRequest) GetApplyToSubGroup() *int32  {
  return s.ApplyToSubGroup
}

func (s *EditApgroupThirdAppRequest) GetCategory() *int32  {
  return s.Category
}

func (s *EditApgroupThirdAppRequest) GetConfigType() *string  {
  return s.ConfigType
}

func (s *EditApgroupThirdAppRequest) GetId() *int64  {
  return s.Id
}

func (s *EditApgroupThirdAppRequest) GetInheritParentGroup() *int32  {
  return s.InheritParentGroup
}

func (s *EditApgroupThirdAppRequest) GetThirdAppName() *string  {
  return s.ThirdAppName
}

func (s *EditApgroupThirdAppRequest) SetApgroupId(v int64) *EditApgroupThirdAppRequest {
  s.ApgroupId = &v
  return s
}

func (s *EditApgroupThirdAppRequest) SetAppCode(v string) *EditApgroupThirdAppRequest {
  s.AppCode = &v
  return s
}

func (s *EditApgroupThirdAppRequest) SetAppData(v string) *EditApgroupThirdAppRequest {
  s.AppData = &v
  return s
}

func (s *EditApgroupThirdAppRequest) SetAppName(v string) *EditApgroupThirdAppRequest {
  s.AppName = &v
  return s
}

func (s *EditApgroupThirdAppRequest) SetApplyToSubGroup(v int32) *EditApgroupThirdAppRequest {
  s.ApplyToSubGroup = &v
  return s
}

func (s *EditApgroupThirdAppRequest) SetCategory(v int32) *EditApgroupThirdAppRequest {
  s.Category = &v
  return s
}

func (s *EditApgroupThirdAppRequest) SetConfigType(v string) *EditApgroupThirdAppRequest {
  s.ConfigType = &v
  return s
}

func (s *EditApgroupThirdAppRequest) SetId(v int64) *EditApgroupThirdAppRequest {
  s.Id = &v
  return s
}

func (s *EditApgroupThirdAppRequest) SetInheritParentGroup(v int32) *EditApgroupThirdAppRequest {
  s.InheritParentGroup = &v
  return s
}

func (s *EditApgroupThirdAppRequest) SetThirdAppName(v string) *EditApgroupThirdAppRequest {
  s.ThirdAppName = &v
  return s
}

func (s *EditApgroupThirdAppRequest) Validate() error {
  return dara.Validate(s)
}

