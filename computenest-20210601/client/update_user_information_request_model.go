// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliverySettings(v *UpdateUserInformationRequestDeliverySettings) *UpdateUserInformationRequest
	GetDeliverySettings() *UpdateUserInformationRequestDeliverySettings
	SetRegionId(v string) *UpdateUserInformationRequest
	GetRegionId() *string
}

type UpdateUserInformationRequest struct {
	// The modified delivery settings.
	DeliverySettings *UpdateUserInformationRequestDeliverySettings `json:"DeliverySettings,omitempty" xml:"DeliverySettings,omitempty" type:"Struct"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateUserInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInformationRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserInformationRequest) GetDeliverySettings() *UpdateUserInformationRequestDeliverySettings {
	return s.DeliverySettings
}

func (s *UpdateUserInformationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateUserInformationRequest) SetDeliverySettings(v *UpdateUserInformationRequestDeliverySettings) *UpdateUserInformationRequest {
	s.DeliverySettings = v
	return s
}

func (s *UpdateUserInformationRequest) SetRegionId(v string) *UpdateUserInformationRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateUserInformationRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateUserInformationRequestDeliverySettings struct {
	// Specifies whether to enable screencast delivery to OSS. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	ActiontrailDeliveryToOssEnabled *bool `json:"ActiontrailDeliveryToOssEnabled,omitempty" xml:"ActiontrailDeliveryToOssEnabled,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// "mybucket"
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Specifies whether to enable screencast delivery to Object Storage Service (OSS). Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	OssEnabled *bool `json:"OssEnabled,omitempty" xml:"OssEnabled,omitempty"`
	// The number of days for which the screencasts are saved.
	//
	// example:
	//
	// 7
	OssExpirationDays *int64 `json:"OssExpirationDays,omitempty" xml:"OssExpirationDays,omitempty"`
	// The OSS path.
	//
	// example:
	//
	// "path1/path2/"
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s UpdateUserInformationRequestDeliverySettings) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInformationRequestDeliverySettings) GoString() string {
	return s.String()
}

func (s *UpdateUserInformationRequestDeliverySettings) GetActiontrailDeliveryToOssEnabled() *bool {
	return s.ActiontrailDeliveryToOssEnabled
}

func (s *UpdateUserInformationRequestDeliverySettings) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *UpdateUserInformationRequestDeliverySettings) GetOssEnabled() *bool {
	return s.OssEnabled
}

func (s *UpdateUserInformationRequestDeliverySettings) GetOssExpirationDays() *int64 {
	return s.OssExpirationDays
}

func (s *UpdateUserInformationRequestDeliverySettings) GetOssPath() *string {
	return s.OssPath
}

func (s *UpdateUserInformationRequestDeliverySettings) SetActiontrailDeliveryToOssEnabled(v bool) *UpdateUserInformationRequestDeliverySettings {
	s.ActiontrailDeliveryToOssEnabled = &v
	return s
}

func (s *UpdateUserInformationRequestDeliverySettings) SetOssBucketName(v string) *UpdateUserInformationRequestDeliverySettings {
	s.OssBucketName = &v
	return s
}

func (s *UpdateUserInformationRequestDeliverySettings) SetOssEnabled(v bool) *UpdateUserInformationRequestDeliverySettings {
	s.OssEnabled = &v
	return s
}

func (s *UpdateUserInformationRequestDeliverySettings) SetOssExpirationDays(v int64) *UpdateUserInformationRequestDeliverySettings {
	s.OssExpirationDays = &v
	return s
}

func (s *UpdateUserInformationRequestDeliverySettings) SetOssPath(v string) *UpdateUserInformationRequestDeliverySettings {
	s.OssPath = &v
	return s
}

func (s *UpdateUserInformationRequestDeliverySettings) Validate() error {
	return dara.Validate(s)
}
