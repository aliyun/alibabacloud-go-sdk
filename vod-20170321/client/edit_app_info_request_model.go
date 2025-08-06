// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditAppInfoRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppItemId(v string) *EditAppInfoRequest
  GetAppItemId() *string 
  SetAppName(v string) *EditAppInfoRequest
  GetAppName() *string 
  SetAppType(v int32) *EditAppInfoRequest
  GetAppType() *int32 
  SetPlatforms(v []*EditAppInfoRequestPlatforms) *EditAppInfoRequest
  GetPlatforms() []*EditAppInfoRequestPlatforms 
}

type EditAppInfoRequest struct {
  AppItemId *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // example:
  // 
  // 1-普通应用，2-内嵌SDK.
  AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
  Platforms []*EditAppInfoRequestPlatforms `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
}

func (s EditAppInfoRequest) String() string {
  return dara.Prettify(s)
}

func (s EditAppInfoRequest) GoString() string {
  return s.String()
}

func (s *EditAppInfoRequest) GetAppItemId() *string  {
  return s.AppItemId
}

func (s *EditAppInfoRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EditAppInfoRequest) GetAppType() *int32  {
  return s.AppType
}

func (s *EditAppInfoRequest) GetPlatforms() []*EditAppInfoRequestPlatforms  {
  return s.Platforms
}

func (s *EditAppInfoRequest) SetAppItemId(v string) *EditAppInfoRequest {
  s.AppItemId = &v
  return s
}

func (s *EditAppInfoRequest) SetAppName(v string) *EditAppInfoRequest {
  s.AppName = &v
  return s
}

func (s *EditAppInfoRequest) SetAppType(v int32) *EditAppInfoRequest {
  s.AppType = &v
  return s
}

func (s *EditAppInfoRequest) SetPlatforms(v []*EditAppInfoRequestPlatforms) *EditAppInfoRequest {
  s.Platforms = v
  return s
}

func (s *EditAppInfoRequest) Validate() error {
  return dara.Validate(s)
}

type EditAppInfoRequestPlatforms struct {
  PkgName *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
  PkgSignature *string `json:"PkgSignature,omitempty" xml:"PkgSignature,omitempty"`
  // example:
  // 
  // 1
  PlatformType *int64 `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
}

func (s EditAppInfoRequestPlatforms) String() string {
  return dara.Prettify(s)
}

func (s EditAppInfoRequestPlatforms) GoString() string {
  return s.String()
}

func (s *EditAppInfoRequestPlatforms) GetPkgName() *string  {
  return s.PkgName
}

func (s *EditAppInfoRequestPlatforms) GetPkgSignature() *string  {
  return s.PkgSignature
}

func (s *EditAppInfoRequestPlatforms) GetPlatformType() *int64  {
  return s.PlatformType
}

func (s *EditAppInfoRequestPlatforms) SetPkgName(v string) *EditAppInfoRequestPlatforms {
  s.PkgName = &v
  return s
}

func (s *EditAppInfoRequestPlatforms) SetPkgSignature(v string) *EditAppInfoRequestPlatforms {
  s.PkgSignature = &v
  return s
}

func (s *EditAppInfoRequestPlatforms) SetPlatformType(v int64) *EditAppInfoRequestPlatforms {
  s.PlatformType = &v
  return s
}

func (s *EditAppInfoRequestPlatforms) Validate() error {
  return dara.Validate(s)
}

