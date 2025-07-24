// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliverySettings(v *GetUserInformationResponseBodyDeliverySettings) *GetUserInformationResponseBody
	GetDeliverySettings() *GetUserInformationResponseBodyDeliverySettings
	SetRequestId(v string) *GetUserInformationResponseBody
	GetRequestId() *string
}

type GetUserInformationResponseBody struct {
	// The delivery settings.
	DeliverySettings *GetUserInformationResponseBodyDeliverySettings `json:"DeliverySettings,omitempty" xml:"DeliverySettings,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 52EBAF16-22F6-53DB-AE1E-44764FC62AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserInformationResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInformationResponseBody) GetDeliverySettings() *GetUserInformationResponseBodyDeliverySettings {
	return s.DeliverySettings
}

func (s *GetUserInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserInformationResponseBody) SetDeliverySettings(v *GetUserInformationResponseBodyDeliverySettings) *GetUserInformationResponseBody {
	s.DeliverySettings = v
	return s
}

func (s *GetUserInformationResponseBody) SetRequestId(v string) *GetUserInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserInformationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserInformationResponseBodyDeliverySettings struct {
	// Indicates whether screencast delivery to OSS is enabled. Valid values:
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
	// 0101data
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// Indicates whether screencast delivery to Object Storage Service (OSS) is enabled. Valid values:
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
	// /test
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
}

func (s GetUserInformationResponseBodyDeliverySettings) String() string {
	return dara.Prettify(s)
}

func (s GetUserInformationResponseBodyDeliverySettings) GoString() string {
	return s.String()
}

func (s *GetUserInformationResponseBodyDeliverySettings) GetActiontrailDeliveryToOssEnabled() *bool {
	return s.ActiontrailDeliveryToOssEnabled
}

func (s *GetUserInformationResponseBodyDeliverySettings) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *GetUserInformationResponseBodyDeliverySettings) GetOssEnabled() *bool {
	return s.OssEnabled
}

func (s *GetUserInformationResponseBodyDeliverySettings) GetOssExpirationDays() *int64 {
	return s.OssExpirationDays
}

func (s *GetUserInformationResponseBodyDeliverySettings) GetOssPath() *string {
	return s.OssPath
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetActiontrailDeliveryToOssEnabled(v bool) *GetUserInformationResponseBodyDeliverySettings {
	s.ActiontrailDeliveryToOssEnabled = &v
	return s
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetOssBucketName(v string) *GetUserInformationResponseBodyDeliverySettings {
	s.OssBucketName = &v
	return s
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetOssEnabled(v bool) *GetUserInformationResponseBodyDeliverySettings {
	s.OssEnabled = &v
	return s
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetOssExpirationDays(v int64) *GetUserInformationResponseBodyDeliverySettings {
	s.OssExpirationDays = &v
	return s
}

func (s *GetUserInformationResponseBodyDeliverySettings) SetOssPath(v string) *GetUserInformationResponseBodyDeliverySettings {
	s.OssPath = &v
	return s
}

func (s *GetUserInformationResponseBodyDeliverySettings) Validate() error {
	return dara.Validate(s)
}
