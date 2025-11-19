// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateDesktopImageShrinkRequest
	GetApiKey() *string
	SetAutoCleanUserData(v bool) *CreateDesktopImageShrinkRequest
	GetAutoCleanUserData() *bool
	SetDescription(v string) *CreateDesktopImageShrinkRequest
	GetDescription() *string
	SetDesktopId(v string) *CreateDesktopImageShrinkRequest
	GetDesktopId() *string
	SetDiskType(v string) *CreateDesktopImageShrinkRequest
	GetDiskType() *string
	SetEnableStartUpConfig(v bool) *CreateDesktopImageShrinkRequest
	GetEnableStartUpConfig() *bool
	SetImageName(v string) *CreateDesktopImageShrinkRequest
	GetImageName() *string
	SetSessionId(v string) *CreateDesktopImageShrinkRequest
	GetSessionId() *string
	SetStartUpFilePathListShrink(v string) *CreateDesktopImageShrinkRequest
	GetStartUpFilePathListShrink() *string
}

type CreateDesktopImageShrinkRequest struct {
	// This parameter is required.
	ApiKey            *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	AutoCleanUserData *bool   `json:"AutoCleanUserData,omitempty" xml:"AutoCleanUserData,omitempty"`
	Description       *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// This parameter is required.
	DiskType            *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	EnableStartUpConfig *bool   `json:"EnableStartUpConfig,omitempty" xml:"EnableStartUpConfig,omitempty"`
	// This parameter is required.
	ImageName                 *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	SessionId                 *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StartUpFilePathListShrink *string `json:"StartUpFilePathList,omitempty" xml:"StartUpFilePathList,omitempty"`
}

func (s CreateDesktopImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopImageShrinkRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateDesktopImageShrinkRequest) GetAutoCleanUserData() *bool {
	return s.AutoCleanUserData
}

func (s *CreateDesktopImageShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDesktopImageShrinkRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *CreateDesktopImageShrinkRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateDesktopImageShrinkRequest) GetEnableStartUpConfig() *bool {
	return s.EnableStartUpConfig
}

func (s *CreateDesktopImageShrinkRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateDesktopImageShrinkRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateDesktopImageShrinkRequest) GetStartUpFilePathListShrink() *string {
	return s.StartUpFilePathListShrink
}

func (s *CreateDesktopImageShrinkRequest) SetApiKey(v string) *CreateDesktopImageShrinkRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateDesktopImageShrinkRequest) SetAutoCleanUserData(v bool) *CreateDesktopImageShrinkRequest {
	s.AutoCleanUserData = &v
	return s
}

func (s *CreateDesktopImageShrinkRequest) SetDescription(v string) *CreateDesktopImageShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDesktopImageShrinkRequest) SetDesktopId(v string) *CreateDesktopImageShrinkRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateDesktopImageShrinkRequest) SetDiskType(v string) *CreateDesktopImageShrinkRequest {
	s.DiskType = &v
	return s
}

func (s *CreateDesktopImageShrinkRequest) SetEnableStartUpConfig(v bool) *CreateDesktopImageShrinkRequest {
	s.EnableStartUpConfig = &v
	return s
}

func (s *CreateDesktopImageShrinkRequest) SetImageName(v string) *CreateDesktopImageShrinkRequest {
	s.ImageName = &v
	return s
}

func (s *CreateDesktopImageShrinkRequest) SetSessionId(v string) *CreateDesktopImageShrinkRequest {
	s.SessionId = &v
	return s
}

func (s *CreateDesktopImageShrinkRequest) SetStartUpFilePathListShrink(v string) *CreateDesktopImageShrinkRequest {
	s.StartUpFilePathListShrink = &v
	return s
}

func (s *CreateDesktopImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
