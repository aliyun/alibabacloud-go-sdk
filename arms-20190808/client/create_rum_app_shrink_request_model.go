// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRumAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroup(v string) *CreateRumAppShrinkRequest
	GetAppGroup() *string
	SetAppName(v string) *CreateRumAppShrinkRequest
	GetAppName() *string
	SetDescription(v string) *CreateRumAppShrinkRequest
	GetDescription() *string
	SetLanguage(v string) *CreateRumAppShrinkRequest
	GetLanguage() *string
	SetNickName(v string) *CreateRumAppShrinkRequest
	GetNickName() *string
	SetPackageName(v string) *CreateRumAppShrinkRequest
	GetPackageName() *string
	SetRealRegionId(v string) *CreateRumAppShrinkRequest
	GetRealRegionId() *string
	SetRegionId(v string) *CreateRumAppShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateRumAppShrinkRequest
	GetResourceGroupId() *string
	SetSiteType(v string) *CreateRumAppShrinkRequest
	GetSiteType() *string
	SetSource(v string) *CreateRumAppShrinkRequest
	GetSource() *string
	SetTagShrink(v string) *CreateRumAppShrinkRequest
	GetTagShrink() *string
}

type CreateRumAppShrinkRequest struct {
	// The name of the application group.
	//
	// example:
	//
	// default
	AppGroup *string `json:"AppGroup,omitempty" xml:"AppGroup,omitempty"`
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Monitoring description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The language used by the client.
	//
	// example:
	//
	// java
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The nickname of the application.
	//
	// example:
	//
	// test-app
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The name of the Android application package. This parameter is required if you create an Android application.
	//
	// example:
	//
	// com.xxxx.xxxxxx
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The region where the application resides. You can leave this parameter empty or set it to China East 2 Finance.
	//
	// example:
	//
	// cn-shanghai-finance-1
	RealRegionId *string `json:"RealRegionId,omitempty" xml:"RealRegionId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmxyexli2****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The application type. Valid values: web, miniapp, ios, and android.
	//
	// This parameter is required.
	//
	// example:
	//
	// web
	SiteType *string `json:"SiteType,omitempty" xml:"SiteType,omitempty"`
	// The source. This is a reserved parameter.
	//
	// example:
	//
	// arms
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The list of tags. You can specify a maximum of 20 tags.
	TagShrink *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s CreateRumAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRumAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRumAppShrinkRequest) GetAppGroup() *string {
	return s.AppGroup
}

func (s *CreateRumAppShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateRumAppShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRumAppShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateRumAppShrinkRequest) GetNickName() *string {
	return s.NickName
}

func (s *CreateRumAppShrinkRequest) GetPackageName() *string {
	return s.PackageName
}

func (s *CreateRumAppShrinkRequest) GetRealRegionId() *string {
	return s.RealRegionId
}

func (s *CreateRumAppShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRumAppShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRumAppShrinkRequest) GetSiteType() *string {
	return s.SiteType
}

func (s *CreateRumAppShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *CreateRumAppShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *CreateRumAppShrinkRequest) SetAppGroup(v string) *CreateRumAppShrinkRequest {
	s.AppGroup = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetAppName(v string) *CreateRumAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetDescription(v string) *CreateRumAppShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetLanguage(v string) *CreateRumAppShrinkRequest {
	s.Language = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetNickName(v string) *CreateRumAppShrinkRequest {
	s.NickName = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetPackageName(v string) *CreateRumAppShrinkRequest {
	s.PackageName = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetRealRegionId(v string) *CreateRumAppShrinkRequest {
	s.RealRegionId = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetRegionId(v string) *CreateRumAppShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetResourceGroupId(v string) *CreateRumAppShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetSiteType(v string) *CreateRumAppShrinkRequest {
	s.SiteType = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetSource(v string) *CreateRumAppShrinkRequest {
	s.Source = &v
	return s
}

func (s *CreateRumAppShrinkRequest) SetTagShrink(v string) *CreateRumAppShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *CreateRumAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
