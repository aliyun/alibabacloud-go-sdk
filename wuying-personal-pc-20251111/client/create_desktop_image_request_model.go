// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDesktopImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateDesktopImageRequest
	GetApiKey() *string
	SetAutoCleanUserData(v bool) *CreateDesktopImageRequest
	GetAutoCleanUserData() *bool
	SetDescription(v string) *CreateDesktopImageRequest
	GetDescription() *string
	SetDesktopId(v string) *CreateDesktopImageRequest
	GetDesktopId() *string
	SetDiskType(v string) *CreateDesktopImageRequest
	GetDiskType() *string
	SetEnableStartUpConfig(v bool) *CreateDesktopImageRequest
	GetEnableStartUpConfig() *bool
	SetImageName(v string) *CreateDesktopImageRequest
	GetImageName() *string
	SetSessionId(v string) *CreateDesktopImageRequest
	GetSessionId() *string
	SetStartUpFilePathList(v []*string) *CreateDesktopImageRequest
	GetStartUpFilePathList() []*string
}

type CreateDesktopImageRequest struct {
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
	ImageName           *string   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	SessionId           *string   `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StartUpFilePathList []*string `json:"StartUpFilePathList,omitempty" xml:"StartUpFilePathList,omitempty" type:"Repeated"`
}

func (s CreateDesktopImageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDesktopImageRequest) GoString() string {
	return s.String()
}

func (s *CreateDesktopImageRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateDesktopImageRequest) GetAutoCleanUserData() *bool {
	return s.AutoCleanUserData
}

func (s *CreateDesktopImageRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDesktopImageRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *CreateDesktopImageRequest) GetDiskType() *string {
	return s.DiskType
}

func (s *CreateDesktopImageRequest) GetEnableStartUpConfig() *bool {
	return s.EnableStartUpConfig
}

func (s *CreateDesktopImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *CreateDesktopImageRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateDesktopImageRequest) GetStartUpFilePathList() []*string {
	return s.StartUpFilePathList
}

func (s *CreateDesktopImageRequest) SetApiKey(v string) *CreateDesktopImageRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateDesktopImageRequest) SetAutoCleanUserData(v bool) *CreateDesktopImageRequest {
	s.AutoCleanUserData = &v
	return s
}

func (s *CreateDesktopImageRequest) SetDescription(v string) *CreateDesktopImageRequest {
	s.Description = &v
	return s
}

func (s *CreateDesktopImageRequest) SetDesktopId(v string) *CreateDesktopImageRequest {
	s.DesktopId = &v
	return s
}

func (s *CreateDesktopImageRequest) SetDiskType(v string) *CreateDesktopImageRequest {
	s.DiskType = &v
	return s
}

func (s *CreateDesktopImageRequest) SetEnableStartUpConfig(v bool) *CreateDesktopImageRequest {
	s.EnableStartUpConfig = &v
	return s
}

func (s *CreateDesktopImageRequest) SetImageName(v string) *CreateDesktopImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateDesktopImageRequest) SetSessionId(v string) *CreateDesktopImageRequest {
	s.SessionId = &v
	return s
}

func (s *CreateDesktopImageRequest) SetStartUpFilePathList(v []*string) *CreateDesktopImageRequest {
	s.StartUpFilePathList = v
	return s
}

func (s *CreateDesktopImageRequest) Validate() error {
	return dara.Validate(s)
}
