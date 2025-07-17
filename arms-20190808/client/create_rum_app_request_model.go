// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRumAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroup(v string) *CreateRumAppRequest
	GetAppGroup() *string
	SetAppName(v string) *CreateRumAppRequest
	GetAppName() *string
	SetDescription(v string) *CreateRumAppRequest
	GetDescription() *string
	SetLanguage(v string) *CreateRumAppRequest
	GetLanguage() *string
	SetNickName(v string) *CreateRumAppRequest
	GetNickName() *string
	SetPackageName(v string) *CreateRumAppRequest
	GetPackageName() *string
	SetRealRegionId(v string) *CreateRumAppRequest
	GetRealRegionId() *string
	SetRegionId(v string) *CreateRumAppRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateRumAppRequest
	GetResourceGroupId() *string
	SetSiteType(v string) *CreateRumAppRequest
	GetSiteType() *string
	SetSource(v string) *CreateRumAppRequest
	GetSource() *string
	SetTag(v []*CreateRumAppRequestTag) *CreateRumAppRequest
	GetTag() []*CreateRumAppRequestTag
}

type CreateRumAppRequest struct {
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
	Tag []*CreateRumAppRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateRumAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRumAppRequest) GoString() string {
	return s.String()
}

func (s *CreateRumAppRequest) GetAppGroup() *string {
	return s.AppGroup
}

func (s *CreateRumAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateRumAppRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRumAppRequest) GetLanguage() *string {
	return s.Language
}

func (s *CreateRumAppRequest) GetNickName() *string {
	return s.NickName
}

func (s *CreateRumAppRequest) GetPackageName() *string {
	return s.PackageName
}

func (s *CreateRumAppRequest) GetRealRegionId() *string {
	return s.RealRegionId
}

func (s *CreateRumAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRumAppRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateRumAppRequest) GetSiteType() *string {
	return s.SiteType
}

func (s *CreateRumAppRequest) GetSource() *string {
	return s.Source
}

func (s *CreateRumAppRequest) GetTag() []*CreateRumAppRequestTag {
	return s.Tag
}

func (s *CreateRumAppRequest) SetAppGroup(v string) *CreateRumAppRequest {
	s.AppGroup = &v
	return s
}

func (s *CreateRumAppRequest) SetAppName(v string) *CreateRumAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateRumAppRequest) SetDescription(v string) *CreateRumAppRequest {
	s.Description = &v
	return s
}

func (s *CreateRumAppRequest) SetLanguage(v string) *CreateRumAppRequest {
	s.Language = &v
	return s
}

func (s *CreateRumAppRequest) SetNickName(v string) *CreateRumAppRequest {
	s.NickName = &v
	return s
}

func (s *CreateRumAppRequest) SetPackageName(v string) *CreateRumAppRequest {
	s.PackageName = &v
	return s
}

func (s *CreateRumAppRequest) SetRealRegionId(v string) *CreateRumAppRequest {
	s.RealRegionId = &v
	return s
}

func (s *CreateRumAppRequest) SetRegionId(v string) *CreateRumAppRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRumAppRequest) SetResourceGroupId(v string) *CreateRumAppRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateRumAppRequest) SetSiteType(v string) *CreateRumAppRequest {
	s.SiteType = &v
	return s
}

func (s *CreateRumAppRequest) SetSource(v string) *CreateRumAppRequest {
	s.Source = &v
	return s
}

func (s *CreateRumAppRequest) SetTag(v []*CreateRumAppRequestTag) *CreateRumAppRequest {
	s.Tag = v
	return s
}

func (s *CreateRumAppRequest) Validate() error {
	return dara.Validate(s)
}

type CreateRumAppRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// ecs
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateRumAppRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateRumAppRequestTag) GoString() string {
	return s.String()
}

func (s *CreateRumAppRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateRumAppRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateRumAppRequestTag) SetKey(v string) *CreateRumAppRequestTag {
	s.Key = &v
	return s
}

func (s *CreateRumAppRequestTag) SetValue(v string) *CreateRumAppRequestTag {
	s.Value = &v
	return s
}

func (s *CreateRumAppRequestTag) Validate() error {
	return dara.Validate(s)
}
