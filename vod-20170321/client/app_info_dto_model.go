// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppInfoDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *AppInfoDTO
	GetAppName() *string
	SetAppType(v int32) *AppInfoDTO
	GetAppType() *int32
	SetCreationTime(v string) *AppInfoDTO
	GetCreationTime() *string
	SetGmtCreate(v string) *AppInfoDTO
	GetGmtCreate() *string
	SetGmtModified(v string) *AppInfoDTO
	GetGmtModified() *string
	SetItemId(v string) *AppInfoDTO
	GetItemId() *string
	SetModificationTime(v string) *AppInfoDTO
	GetModificationTime() *string
	SetPlatforms(v []*AppInfoDTOPlatforms) *AppInfoDTO
	GetPlatforms() []*AppInfoDTOPlatforms
	SetUserId(v int64) *AppInfoDTO
	GetUserId() *int64
}

type AppInfoDTO struct {
	// The application name.
	//
	// example:
	//
	// vod-demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type. Valid values:
	//
	// - 1: Common application.
	//
	// - 2: Embedded SDK.
	//
	// example:
	//
	// 1
	AppType *int32 `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-02T03:04:05Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-01-02 03:04:05
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-01-02 03:04:05
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 3ba9c7d21f4e6a80
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-01-02T03:04:05Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The list of platform information included in the application.
	Platforms []*AppInfoDTOPlatforms `json:"Platforms,omitempty" xml:"Platforms,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 2544469002080159
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s AppInfoDTO) String() string {
	return dara.Prettify(s)
}

func (s AppInfoDTO) GoString() string {
	return s.String()
}

func (s *AppInfoDTO) GetAppName() *string {
	return s.AppName
}

func (s *AppInfoDTO) GetAppType() *int32 {
	return s.AppType
}

func (s *AppInfoDTO) GetCreationTime() *string {
	return s.CreationTime
}

func (s *AppInfoDTO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AppInfoDTO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AppInfoDTO) GetItemId() *string {
	return s.ItemId
}

func (s *AppInfoDTO) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *AppInfoDTO) GetPlatforms() []*AppInfoDTOPlatforms {
	return s.Platforms
}

func (s *AppInfoDTO) GetUserId() *int64 {
	return s.UserId
}

func (s *AppInfoDTO) SetAppName(v string) *AppInfoDTO {
	s.AppName = &v
	return s
}

func (s *AppInfoDTO) SetAppType(v int32) *AppInfoDTO {
	s.AppType = &v
	return s
}

func (s *AppInfoDTO) SetCreationTime(v string) *AppInfoDTO {
	s.CreationTime = &v
	return s
}

func (s *AppInfoDTO) SetGmtCreate(v string) *AppInfoDTO {
	s.GmtCreate = &v
	return s
}

func (s *AppInfoDTO) SetGmtModified(v string) *AppInfoDTO {
	s.GmtModified = &v
	return s
}

func (s *AppInfoDTO) SetItemId(v string) *AppInfoDTO {
	s.ItemId = &v
	return s
}

func (s *AppInfoDTO) SetModificationTime(v string) *AppInfoDTO {
	s.ModificationTime = &v
	return s
}

func (s *AppInfoDTO) SetPlatforms(v []*AppInfoDTOPlatforms) *AppInfoDTO {
	s.Platforms = v
	return s
}

func (s *AppInfoDTO) SetUserId(v int64) *AppInfoDTO {
	s.UserId = &v
	return s
}

func (s *AppInfoDTO) Validate() error {
	if s.Platforms != nil {
		for _, item := range s.Platforms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AppInfoDTOPlatforms struct {
	// The platform ID.
	//
	// example:
	//
	// 3ba9c7d21f4e6a81
	ItemId *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The authorization list.
	LicenseItemIds []*string `json:"LicenseItemIds,omitempty" xml:"LicenseItemIds,omitempty" type:"Repeated"`
	// The package name.
	//
	// example:
	//
	// com.example.vodapp
	PkgName *string `json:"PkgName,omitempty" xml:"PkgName,omitempty"`
	// The package signature.
	//
	// example:
	//
	// 8e610rzjg613h0u58v8473d029h7
	PkgSignature *string `json:"PkgSignature,omitempty" xml:"PkgSignature,omitempty"`
	// The platform type.
	//
	// example:
	//
	// 2
	PlatformType *int64 `json:"PlatformType,omitempty" xml:"PlatformType,omitempty"`
	// The platform type.
	//
	// example:
	//
	// 2
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s AppInfoDTOPlatforms) String() string {
	return dara.Prettify(s)
}

func (s AppInfoDTOPlatforms) GoString() string {
	return s.String()
}

func (s *AppInfoDTOPlatforms) GetItemId() *string {
	return s.ItemId
}

func (s *AppInfoDTOPlatforms) GetLicenseItemIds() []*string {
	return s.LicenseItemIds
}

func (s *AppInfoDTOPlatforms) GetPkgName() *string {
	return s.PkgName
}

func (s *AppInfoDTOPlatforms) GetPkgSignature() *string {
	return s.PkgSignature
}

func (s *AppInfoDTOPlatforms) GetPlatformType() *int64 {
	return s.PlatformType
}

func (s *AppInfoDTOPlatforms) GetType() *int64 {
	return s.Type
}

func (s *AppInfoDTOPlatforms) SetItemId(v string) *AppInfoDTOPlatforms {
	s.ItemId = &v
	return s
}

func (s *AppInfoDTOPlatforms) SetLicenseItemIds(v []*string) *AppInfoDTOPlatforms {
	s.LicenseItemIds = v
	return s
}

func (s *AppInfoDTOPlatforms) SetPkgName(v string) *AppInfoDTOPlatforms {
	s.PkgName = &v
	return s
}

func (s *AppInfoDTOPlatforms) SetPkgSignature(v string) *AppInfoDTOPlatforms {
	s.PkgSignature = &v
	return s
}

func (s *AppInfoDTOPlatforms) SetPlatformType(v int64) *AppInfoDTOPlatforms {
	s.PlatformType = &v
	return s
}

func (s *AppInfoDTOPlatforms) SetType(v int64) *AppInfoDTOPlatforms {
	s.Type = &v
	return s
}

func (s *AppInfoDTOPlatforms) Validate() error {
	return dara.Validate(s)
}
