// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFreeLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppItemId(v string) *AddFreeLicenseRequest
	GetAppItemId() *string
	SetAppName(v string) *AddFreeLicenseRequest
	GetAppName() *string
	SetAppPlatforms(v string) *AddFreeLicenseRequest
	GetAppPlatforms() *string
	SetSdkModels(v string) *AddFreeLicenseRequest
	GetSdkModels() *string
}

type AddFreeLicenseRequest struct {
	AppItemId    *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPlatforms *string `json:"AppPlatforms,omitempty" xml:"AppPlatforms,omitempty"`
	SdkModels    *string `json:"SdkModels,omitempty" xml:"SdkModels,omitempty"`
}

func (s AddFreeLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFreeLicenseRequest) GoString() string {
	return s.String()
}

func (s *AddFreeLicenseRequest) GetAppItemId() *string {
	return s.AppItemId
}

func (s *AddFreeLicenseRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddFreeLicenseRequest) GetAppPlatforms() *string {
	return s.AppPlatforms
}

func (s *AddFreeLicenseRequest) GetSdkModels() *string {
	return s.SdkModels
}

func (s *AddFreeLicenseRequest) SetAppItemId(v string) *AddFreeLicenseRequest {
	s.AppItemId = &v
	return s
}

func (s *AddFreeLicenseRequest) SetAppName(v string) *AddFreeLicenseRequest {
	s.AppName = &v
	return s
}

func (s *AddFreeLicenseRequest) SetAppPlatforms(v string) *AddFreeLicenseRequest {
	s.AppPlatforms = &v
	return s
}

func (s *AddFreeLicenseRequest) SetSdkModels(v string) *AddFreeLicenseRequest {
	s.SdkModels = &v
	return s
}

func (s *AddFreeLicenseRequest) Validate() error {
	return dara.Validate(s)
}
