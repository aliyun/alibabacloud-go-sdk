// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApp interface {
	dara.Model
	String() string
	GoString() string
	SetAppDescription(v string) *App
	GetAppDescription() *string
	SetAppId(v string) *App
	GetAppId() *string
	SetAppKey(v string) *App
	GetAppKey() *string
	SetAppName(v string) *App
	GetAppName() *string
	SetAppRegion(v int64) *App
	GetAppRegion() *int64
	SetAppType(v int64) *App
	GetAppType() *int64
	SetEnglishName(v string) *App
	GetEnglishName() *string
	SetOwnerId(v string) *App
	GetOwnerId() *string
	SetPackageName(v string) *App
	GetPackageName() *string
}

type App struct {
	AppDescription *string `json:"AppDescription,omitempty" xml:"AppDescription,omitempty"`
	AppId          *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppKey         *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName        *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppRegion      *int64  `json:"AppRegion,omitempty" xml:"AppRegion,omitempty"`
	AppType        *int64  `json:"AppType,omitempty" xml:"AppType,omitempty"`
	EnglishName    *string `json:"EnglishName,omitempty" xml:"EnglishName,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PackageName    *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
}

func (s App) String() string {
	return dara.Prettify(s)
}

func (s App) GoString() string {
	return s.String()
}

func (s *App) GetAppDescription() *string {
	return s.AppDescription
}

func (s *App) GetAppId() *string {
	return s.AppId
}

func (s *App) GetAppKey() *string {
	return s.AppKey
}

func (s *App) GetAppName() *string {
	return s.AppName
}

func (s *App) GetAppRegion() *int64 {
	return s.AppRegion
}

func (s *App) GetAppType() *int64 {
	return s.AppType
}

func (s *App) GetEnglishName() *string {
	return s.EnglishName
}

func (s *App) GetOwnerId() *string {
	return s.OwnerId
}

func (s *App) GetPackageName() *string {
	return s.PackageName
}

func (s *App) SetAppDescription(v string) *App {
	s.AppDescription = &v
	return s
}

func (s *App) SetAppId(v string) *App {
	s.AppId = &v
	return s
}

func (s *App) SetAppKey(v string) *App {
	s.AppKey = &v
	return s
}

func (s *App) SetAppName(v string) *App {
	s.AppName = &v
	return s
}

func (s *App) SetAppRegion(v int64) *App {
	s.AppRegion = &v
	return s
}

func (s *App) SetAppType(v int64) *App {
	s.AppType = &v
	return s
}

func (s *App) SetEnglishName(v string) *App {
	s.EnglishName = &v
	return s
}

func (s *App) SetOwnerId(v string) *App {
	s.OwnerId = &v
	return s
}

func (s *App) SetPackageName(v string) *App {
	s.PackageName = &v
	return s
}

func (s *App) Validate() error {
	return dara.Validate(s)
}
