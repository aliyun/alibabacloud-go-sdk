// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetPlayConfigRequest
	GetAppId() *string
	SetAppName(v string) *GetPlayConfigRequest
	GetAppName() *string
	SetBrand(v string) *GetPlayConfigRequest
	GetBrand() *string
	SetConfigType(v string) *GetPlayConfigRequest
	GetConfigType() *string
	SetOS(v string) *GetPlayConfigRequest
	GetOS() *string
	SetOSVersion(v string) *GetPlayConfigRequest
	GetOSVersion() *string
	SetSDKVersion(v string) *GetPlayConfigRequest
	GetSDKVersion() *string
}

type GetPlayConfigRequest struct {
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Brand   *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// This parameter is required.
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	OS         *string `json:"OS,omitempty" xml:"OS,omitempty"`
	OSVersion  *string `json:"OSVersion,omitempty" xml:"OSVersion,omitempty"`
	SDKVersion *string `json:"SDKVersion,omitempty" xml:"SDKVersion,omitempty"`
}

func (s GetPlayConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPlayConfigRequest) GoString() string {
	return s.String()
}

func (s *GetPlayConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetPlayConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetPlayConfigRequest) GetBrand() *string {
	return s.Brand
}

func (s *GetPlayConfigRequest) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetPlayConfigRequest) GetOS() *string {
	return s.OS
}

func (s *GetPlayConfigRequest) GetOSVersion() *string {
	return s.OSVersion
}

func (s *GetPlayConfigRequest) GetSDKVersion() *string {
	return s.SDKVersion
}

func (s *GetPlayConfigRequest) SetAppId(v string) *GetPlayConfigRequest {
	s.AppId = &v
	return s
}

func (s *GetPlayConfigRequest) SetAppName(v string) *GetPlayConfigRequest {
	s.AppName = &v
	return s
}

func (s *GetPlayConfigRequest) SetBrand(v string) *GetPlayConfigRequest {
	s.Brand = &v
	return s
}

func (s *GetPlayConfigRequest) SetConfigType(v string) *GetPlayConfigRequest {
	s.ConfigType = &v
	return s
}

func (s *GetPlayConfigRequest) SetOS(v string) *GetPlayConfigRequest {
	s.OS = &v
	return s
}

func (s *GetPlayConfigRequest) SetOSVersion(v string) *GetPlayConfigRequest {
	s.OSVersion = &v
	return s
}

func (s *GetPlayConfigRequest) SetSDKVersion(v string) *GetPlayConfigRequest {
	s.SDKVersion = &v
	return s
}

func (s *GetPlayConfigRequest) Validate() error {
	return dara.Validate(s)
}
