// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSDKDownloadListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSDKDownloadListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeSDKDownloadListResponseBodyResultObject) *DescribeSDKDownloadListResponseBody
	GetResultObject() []*DescribeSDKDownloadListResponseBodyResultObject
}

type DescribeSDKDownloadListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return object
	ResultObject []*DescribeSDKDownloadListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
}

func (s DescribeSDKDownloadListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDKDownloadListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSDKDownloadListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSDKDownloadListResponseBody) GetResultObject() []*DescribeSDKDownloadListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSDKDownloadListResponseBody) SetRequestId(v string) *DescribeSDKDownloadListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBody) SetResultObject(v []*DescribeSDKDownloadListResponseBodyResultObject) *DescribeSDKDownloadListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSDKDownloadListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSDKDownloadListResponseBodyResultObject struct {
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Developer
	//
	// example:
	//
	// 阿里云安全-风险识别
	Developer *string `json:"developer,omitempty" xml:"developer,omitempty"`
	// Device type.
	//
	// example:
	//
	// ANDROID
	DeviceType *string `json:"deviceType,omitempty" xml:"deviceType,omitempty"`
	// Download URL.
	//
	// example:
	//
	// https://aliyun-xxxx.oss-cn-xxx.xxx.com/sdk/xxx/10056.1/Android-AliyunDeviceEnhance-10056.1-20250611.tgz
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// File MD5.
	//
	// example:
	//
	// E582EEB6B4BC9B5CB168AA5A7DD0EE93
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// Package name
	//
	// example:
	//
	// net.security.device
	PackageName *string `json:"packageName,omitempty" xml:"packageName,omitempty"`
	// Risk recognition SDK privacy policy link
	//
	// example:
	//
	// https://terms.aliyun.com/legal-agreement/terms/suit_bu1_ali_cloud/suit_bu1_ali_cloud202111120818_92724.html
	PrivacyLink *string `json:"privacyLink,omitempty" xml:"privacyLink,omitempty"`
	// Release time
	//
	// example:
	//
	// 1751212800000
	PushTime *string `json:"pushTime,omitempty" xml:"pushTime,omitempty"`
	// SDK version.
	//
	// example:
	//
	// 1
	SdkVersion *string `json:"sdkVersion,omitempty" xml:"sdkVersion,omitempty"`
	// Size
	//
	// example:
	//
	// 4.12 MB
	Size *string `json:"size,omitempty" xml:"size,omitempty"`
}

func (s DescribeSDKDownloadListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSDKDownloadListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetDeveloper() *string {
	return s.Developer
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetDeviceType() *string {
	return s.DeviceType
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetMd5() *string {
	return s.Md5
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetPackageName() *string {
	return s.PackageName
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetPrivacyLink() *string {
	return s.PrivacyLink
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetPushTime() *string {
	return s.PushTime
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) GetSize() *string {
	return s.Size
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetDescription(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetDeveloper(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.Developer = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetDeviceType(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.DeviceType = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetDownloadUrl(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.DownloadUrl = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetMd5(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.Md5 = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetPackageName(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.PackageName = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetPrivacyLink(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.PrivacyLink = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetPushTime(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.PushTime = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetSdkVersion(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.SdkVersion = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) SetSize(v string) *DescribeSDKDownloadListResponseBodyResultObject {
	s.Size = &v
	return s
}

func (s *DescribeSDKDownloadListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
