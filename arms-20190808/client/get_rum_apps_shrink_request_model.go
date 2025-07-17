// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumAppsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroup(v string) *GetRumAppsShrinkRequest
	GetAppGroup() *string
	SetAppId(v string) *GetRumAppsShrinkRequest
	GetAppId() *string
	SetAppName(v string) *GetRumAppsShrinkRequest
	GetAppName() *string
	SetAppType(v string) *GetRumAppsShrinkRequest
	GetAppType() *string
	SetRegionId(v string) *GetRumAppsShrinkRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *GetRumAppsShrinkRequest
	GetResourceGroupId() *string
	SetTagsShrink(v string) *GetRumAppsShrinkRequest
	GetTagsShrink() *string
}

type GetRumAppsShrinkRequest struct {
	// The group to which the application belongs.
	//
	// example:
	//
	// default
	AppGroup *string `json:"AppGroup,omitempty" xml:"AppGroup,omitempty"`
	// The application ID.
	//
	// example:
	//
	// b590lhguqs@28f515462******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application. You can specify only one application name in each request.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application type. Currently, supported application types include: web, mini program, Android, iOS, Windows, macOS, and HarmonyOS.
	//
	// example:
	//
	// web
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
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
	// The tags.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s GetRumAppsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetRumAppsShrinkRequest) GetAppGroup() *string {
	return s.AppGroup
}

func (s *GetRumAppsShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetRumAppsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetRumAppsShrinkRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetRumAppsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumAppsShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetRumAppsShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *GetRumAppsShrinkRequest) SetAppGroup(v string) *GetRumAppsShrinkRequest {
	s.AppGroup = &v
	return s
}

func (s *GetRumAppsShrinkRequest) SetAppId(v string) *GetRumAppsShrinkRequest {
	s.AppId = &v
	return s
}

func (s *GetRumAppsShrinkRequest) SetAppName(v string) *GetRumAppsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *GetRumAppsShrinkRequest) SetAppType(v string) *GetRumAppsShrinkRequest {
	s.AppType = &v
	return s
}

func (s *GetRumAppsShrinkRequest) SetRegionId(v string) *GetRumAppsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetRumAppsShrinkRequest) SetResourceGroupId(v string) *GetRumAppsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRumAppsShrinkRequest) SetTagsShrink(v string) *GetRumAppsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *GetRumAppsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
