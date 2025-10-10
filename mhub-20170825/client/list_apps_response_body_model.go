// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfos(v *ListAppsResponseBodyAppInfos) *ListAppsResponseBody
	GetAppInfos() *ListAppsResponseBodyAppInfos
	SetRequestId(v string) *ListAppsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListAppsResponseBody
	GetTotal() *int32
	SetUbsmsStatus(v string) *ListAppsResponseBody
	GetUbsmsStatus() *string
}

type ListAppsResponseBody struct {
	AppInfos *ListAppsResponseBodyAppInfos `json:"AppInfos,omitempty" xml:"AppInfos,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// enabled
	UbsmsStatus *string `json:"UbsmsStatus,omitempty" xml:"UbsmsStatus,omitempty"`
}

func (s ListAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBody) GetAppInfos() *ListAppsResponseBodyAppInfos {
	return s.AppInfos
}

func (s *ListAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListAppsResponseBody) GetUbsmsStatus() *string {
	return s.UbsmsStatus
}

func (s *ListAppsResponseBody) SetAppInfos(v *ListAppsResponseBodyAppInfos) *ListAppsResponseBody {
	s.AppInfos = v
	return s
}

func (s *ListAppsResponseBody) SetRequestId(v string) *ListAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppsResponseBody) SetTotal(v int32) *ListAppsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAppsResponseBody) SetUbsmsStatus(v string) *ListAppsResponseBody {
	s.UbsmsStatus = &v
	return s
}

func (s *ListAppsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAppsResponseBodyAppInfos struct {
	AppInfo []*ListAppsResponseBodyAppInfosAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Repeated"`
}

func (s ListAppsResponseBodyAppInfos) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponseBodyAppInfos) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyAppInfos) GetAppInfo() []*ListAppsResponseBodyAppInfosAppInfo {
	return s.AppInfo
}

func (s *ListAppsResponseBodyAppInfos) SetAppInfo(v []*ListAppsResponseBodyAppInfosAppInfo) *ListAppsResponseBodyAppInfos {
	s.AppInfo = v
	return s
}

func (s *ListAppsResponseBodyAppInfos) Validate() error {
	return dara.Validate(s)
}

type ListAppsResponseBodyAppInfosAppInfo struct {
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// mobile-live-service
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// com.test.ios
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// example:
	//
	// 图片 base64 字符串
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *int32 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// example:
	//
	// 我的应用
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// example:
	//
	// false
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListAppsResponseBodyAppInfosAppInfo) String() string {
	return dara.Prettify(s)
}

func (s ListAppsResponseBodyAppInfosAppInfo) GoString() string {
	return s.String()
}

func (s *ListAppsResponseBodyAppInfosAppInfo) GetAppKey() *string {
	return s.AppKey
}

func (s *ListAppsResponseBodyAppInfosAppInfo) GetAppName() *string {
	return s.AppName
}

func (s *ListAppsResponseBodyAppInfosAppInfo) GetBundleId() *string {
	return s.BundleId
}

func (s *ListAppsResponseBodyAppInfosAppInfo) GetEncodedIcon() *string {
	return s.EncodedIcon
}

func (s *ListAppsResponseBodyAppInfosAppInfo) GetIndustryId() *int32 {
	return s.IndustryId
}

func (s *ListAppsResponseBodyAppInfosAppInfo) GetName() *string {
	return s.Name
}

func (s *ListAppsResponseBodyAppInfosAppInfo) GetPackageName() *string {
	return s.PackageName
}

func (s *ListAppsResponseBodyAppInfosAppInfo) GetReadonly() *bool {
	return s.Readonly
}

func (s *ListAppsResponseBodyAppInfosAppInfo) GetType() *int32 {
	return s.Type
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetAppKey(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.AppKey = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetAppName(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.AppName = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetBundleId(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.BundleId = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetEncodedIcon(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.EncodedIcon = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetIndustryId(v int32) *ListAppsResponseBodyAppInfosAppInfo {
	s.IndustryId = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetName(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.Name = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetPackageName(v string) *ListAppsResponseBodyAppInfosAppInfo {
	s.PackageName = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetReadonly(v bool) *ListAppsResponseBodyAppInfosAppInfo {
	s.Readonly = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) SetType(v int32) *ListAppsResponseBodyAppInfosAppInfo {
	s.Type = &v
	return s
}

func (s *ListAppsResponseBodyAppInfosAppInfo) Validate() error {
	return dara.Validate(s)
}
