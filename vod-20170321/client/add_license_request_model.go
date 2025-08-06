// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLicenseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppItemId(v string) *AddLicenseRequest
	GetAppItemId() *string
	SetAppName(v string) *AddLicenseRequest
	GetAppName() *string
	SetAppPlatforms(v string) *AddLicenseRequest
	GetAppPlatforms() *string
	SetContractNo(v string) *AddLicenseRequest
	GetContractNo() *string
	SetInstanceId(v string) *AddLicenseRequest
	GetInstanceId() *string
	SetSdkModels(v string) *AddLicenseRequest
	GetSdkModels() *string
}

type AddLicenseRequest struct {
	AppItemId    *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	AppName      *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	AppPlatforms *string `json:"AppPlatforms,omitempty" xml:"AppPlatforms,omitempty"`
	ContractNo   *string `json:"ContractNo,omitempty" xml:"ContractNo,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SdkModels    *string `json:"SdkModels,omitempty" xml:"SdkModels,omitempty"`
}

func (s AddLicenseRequest) String() string {
	return dara.Prettify(s)
}

func (s AddLicenseRequest) GoString() string {
	return s.String()
}

func (s *AddLicenseRequest) GetAppItemId() *string {
	return s.AppItemId
}

func (s *AddLicenseRequest) GetAppName() *string {
	return s.AppName
}

func (s *AddLicenseRequest) GetAppPlatforms() *string {
	return s.AppPlatforms
}

func (s *AddLicenseRequest) GetContractNo() *string {
	return s.ContractNo
}

func (s *AddLicenseRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddLicenseRequest) GetSdkModels() *string {
	return s.SdkModels
}

func (s *AddLicenseRequest) SetAppItemId(v string) *AddLicenseRequest {
	s.AppItemId = &v
	return s
}

func (s *AddLicenseRequest) SetAppName(v string) *AddLicenseRequest {
	s.AppName = &v
	return s
}

func (s *AddLicenseRequest) SetAppPlatforms(v string) *AddLicenseRequest {
	s.AppPlatforms = &v
	return s
}

func (s *AddLicenseRequest) SetContractNo(v string) *AddLicenseRequest {
	s.ContractNo = &v
	return s
}

func (s *AddLicenseRequest) SetInstanceId(v string) *AddLicenseRequest {
	s.InstanceId = &v
	return s
}

func (s *AddLicenseRequest) SetSdkModels(v string) *AddLicenseRequest {
	s.SdkModels = &v
	return s
}

func (s *AddLicenseRequest) Validate() error {
	return dara.Validate(s)
}
