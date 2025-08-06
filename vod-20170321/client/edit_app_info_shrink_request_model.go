// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEditAppInfoShrinkRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAppItemId(v string) *EditAppInfoShrinkRequest
  GetAppItemId() *string 
  SetAppName(v string) *EditAppInfoShrinkRequest
  GetAppName() *string 
  SetAppType(v int32) *EditAppInfoShrinkRequest
  GetAppType() *int32 
  SetPlatformsShrink(v string) *EditAppInfoShrinkRequest
  GetPlatformsShrink() *string 
}

type EditAppInfoShrinkRequest struct {
  AppItemId *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
  AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
  // example:
  // 
  // 1-普通应用，2-内嵌SDK.
  AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
  PlatformsShrink *string `json:"Platforms,omitempty" xml:"Platforms,omitempty"`
}

func (s EditAppInfoShrinkRequest) String() string {
  return dara.Prettify(s)
}

func (s EditAppInfoShrinkRequest) GoString() string {
  return s.String()
}

func (s *EditAppInfoShrinkRequest) GetAppItemId() *string  {
  return s.AppItemId
}

func (s *EditAppInfoShrinkRequest) GetAppName() *string  {
  return s.AppName
}

func (s *EditAppInfoShrinkRequest) GetAppType() *int32  {
  return s.AppType
}

func (s *EditAppInfoShrinkRequest) GetPlatformsShrink() *string  {
  return s.PlatformsShrink
}

func (s *EditAppInfoShrinkRequest) SetAppItemId(v string) *EditAppInfoShrinkRequest {
  s.AppItemId = &v
  return s
}

func (s *EditAppInfoShrinkRequest) SetAppName(v string) *EditAppInfoShrinkRequest {
  s.AppName = &v
  return s
}

func (s *EditAppInfoShrinkRequest) SetAppType(v int32) *EditAppInfoShrinkRequest {
  s.AppType = &v
  return s
}

func (s *EditAppInfoShrinkRequest) SetPlatformsShrink(v string) *EditAppInfoShrinkRequest {
  s.PlatformsShrink = &v
  return s
}

func (s *EditAppInfoShrinkRequest) Validate() error {
  return dara.Validate(s)
}

