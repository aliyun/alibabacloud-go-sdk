// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAppInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfo(v *QueryAppInfoResponseBodyAppInfo) *QueryAppInfoResponseBody
	GetAppInfo() *QueryAppInfoResponseBodyAppInfo
	SetRequestId(v string) *QueryAppInfoResponseBody
	GetRequestId() *string
}

type QueryAppInfoResponseBody struct {
	AppInfo *QueryAppInfoResponseBodyAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929BFB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryAppInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAppInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAppInfoResponseBody) GetAppInfo() *QueryAppInfoResponseBodyAppInfo {
	return s.AppInfo
}

func (s *QueryAppInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAppInfoResponseBody) SetAppInfo(v *QueryAppInfoResponseBodyAppInfo) *QueryAppInfoResponseBody {
	s.AppInfo = v
	return s
}

func (s *QueryAppInfoResponseBody) SetRequestId(v string) *QueryAppInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAppInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryAppInfoResponseBodyAppInfo struct {
	// example:
	//
	// 123456
	AppKey  *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// com.test.ios
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// example:
	//
	// false
	CertDevelopAvail *bool `json:"CertDevelopAvail,omitempty" xml:"CertDevelopAvail,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	CertDevelopExpiration *string `json:"CertDevelopExpiration,omitempty" xml:"CertDevelopExpiration,omitempty"`
	// example:
	//
	// false
	CertProductAvail *bool `json:"CertProductAvail,omitempty" xml:"CertProductAvail,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	CertProductExpiration *string `json:"CertProductExpiration,omitempty" xml:"CertProductExpiration,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EncodedIcon *string `json:"EncodedIcon,omitempty" xml:"EncodedIcon,omitempty"`
	// example:
	//
	// 1
	IndustryId *int32 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// example:
	//
	// 1234
	ProductId *int64 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// false
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryAppInfoResponseBodyAppInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryAppInfoResponseBodyAppInfo) GoString() string {
	return s.String()
}

func (s *QueryAppInfoResponseBodyAppInfo) GetAppKey() *string {
	return s.AppKey
}

func (s *QueryAppInfoResponseBodyAppInfo) GetAppName() *string {
	return s.AppName
}

func (s *QueryAppInfoResponseBodyAppInfo) GetBundleId() *string {
	return s.BundleId
}

func (s *QueryAppInfoResponseBodyAppInfo) GetCertDevelopAvail() *bool {
	return s.CertDevelopAvail
}

func (s *QueryAppInfoResponseBodyAppInfo) GetCertDevelopExpiration() *string {
	return s.CertDevelopExpiration
}

func (s *QueryAppInfoResponseBodyAppInfo) GetCertProductAvail() *bool {
	return s.CertProductAvail
}

func (s *QueryAppInfoResponseBodyAppInfo) GetCertProductExpiration() *string {
	return s.CertProductExpiration
}

func (s *QueryAppInfoResponseBodyAppInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryAppInfoResponseBodyAppInfo) GetEncodedIcon() *string {
	return s.EncodedIcon
}

func (s *QueryAppInfoResponseBodyAppInfo) GetIndustryId() *int32 {
	return s.IndustryId
}

func (s *QueryAppInfoResponseBodyAppInfo) GetPackageName() *string {
	return s.PackageName
}

func (s *QueryAppInfoResponseBodyAppInfo) GetProductId() *int64 {
	return s.ProductId
}

func (s *QueryAppInfoResponseBodyAppInfo) GetReadonly() *bool {
	return s.Readonly
}

func (s *QueryAppInfoResponseBodyAppInfo) GetStatus() *int32 {
	return s.Status
}

func (s *QueryAppInfoResponseBodyAppInfo) GetType() *int32 {
	return s.Type
}

func (s *QueryAppInfoResponseBodyAppInfo) SetAppKey(v string) *QueryAppInfoResponseBodyAppInfo {
	s.AppKey = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetAppName(v string) *QueryAppInfoResponseBodyAppInfo {
	s.AppName = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetBundleId(v string) *QueryAppInfoResponseBodyAppInfo {
	s.BundleId = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCertDevelopAvail(v bool) *QueryAppInfoResponseBodyAppInfo {
	s.CertDevelopAvail = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCertDevelopExpiration(v string) *QueryAppInfoResponseBodyAppInfo {
	s.CertDevelopExpiration = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCertProductAvail(v bool) *QueryAppInfoResponseBodyAppInfo {
	s.CertProductAvail = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCertProductExpiration(v string) *QueryAppInfoResponseBodyAppInfo {
	s.CertProductExpiration = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetCreateTime(v string) *QueryAppInfoResponseBodyAppInfo {
	s.CreateTime = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetEncodedIcon(v string) *QueryAppInfoResponseBodyAppInfo {
	s.EncodedIcon = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetIndustryId(v int32) *QueryAppInfoResponseBodyAppInfo {
	s.IndustryId = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetPackageName(v string) *QueryAppInfoResponseBodyAppInfo {
	s.PackageName = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetProductId(v int64) *QueryAppInfoResponseBodyAppInfo {
	s.ProductId = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetReadonly(v bool) *QueryAppInfoResponseBodyAppInfo {
	s.Readonly = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetStatus(v int32) *QueryAppInfoResponseBodyAppInfo {
	s.Status = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) SetType(v int32) *QueryAppInfoResponseBodyAppInfo {
	s.Type = &v
	return s
}

func (s *QueryAppInfoResponseBodyAppInfo) Validate() error {
	return dara.Validate(s)
}
