// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppGroup(v string) *GetRumAppsRequest
	GetAppGroup() *string
	SetAppId(v string) *GetRumAppsRequest
	GetAppId() *string
	SetAppName(v string) *GetRumAppsRequest
	GetAppName() *string
	SetAppType(v string) *GetRumAppsRequest
	GetAppType() *string
	SetRegionId(v string) *GetRumAppsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *GetRumAppsRequest
	GetResourceGroupId() *string
	SetTags(v []*GetRumAppsRequestTags) *GetRumAppsRequest
	GetTags() []*GetRumAppsRequestTags
}

type GetRumAppsRequest struct {
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
	Tags []*GetRumAppsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetRumAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppsRequest) GoString() string {
	return s.String()
}

func (s *GetRumAppsRequest) GetAppGroup() *string {
	return s.AppGroup
}

func (s *GetRumAppsRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetRumAppsRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetRumAppsRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetRumAppsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumAppsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetRumAppsRequest) GetTags() []*GetRumAppsRequestTags {
	return s.Tags
}

func (s *GetRumAppsRequest) SetAppGroup(v string) *GetRumAppsRequest {
	s.AppGroup = &v
	return s
}

func (s *GetRumAppsRequest) SetAppId(v string) *GetRumAppsRequest {
	s.AppId = &v
	return s
}

func (s *GetRumAppsRequest) SetAppName(v string) *GetRumAppsRequest {
	s.AppName = &v
	return s
}

func (s *GetRumAppsRequest) SetAppType(v string) *GetRumAppsRequest {
	s.AppType = &v
	return s
}

func (s *GetRumAppsRequest) SetRegionId(v string) *GetRumAppsRequest {
	s.RegionId = &v
	return s
}

func (s *GetRumAppsRequest) SetResourceGroupId(v string) *GetRumAppsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetRumAppsRequest) SetTags(v []*GetRumAppsRequestTags) *GetRumAppsRequest {
	s.Tags = v
	return s
}

func (s *GetRumAppsRequest) Validate() error {
	return dara.Validate(s)
}

type GetRumAppsRequestTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetRumAppsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s GetRumAppsRequestTags) GoString() string {
	return s.String()
}

func (s *GetRumAppsRequestTags) GetKey() *string {
	return s.Key
}

func (s *GetRumAppsRequestTags) GetValue() *string {
	return s.Value
}

func (s *GetRumAppsRequestTags) SetKey(v string) *GetRumAppsRequestTags {
	s.Key = &v
	return s
}

func (s *GetRumAppsRequestTags) SetValue(v string) *GetRumAppsRequestTags {
	s.Value = &v
	return s
}

func (s *GetRumAppsRequestTags) Validate() error {
	return dara.Validate(s)
}
