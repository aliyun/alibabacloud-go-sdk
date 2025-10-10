// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppInfo(v *CreateAppResponseBodyAppInfo) *CreateAppResponseBody
	GetAppInfo() *CreateAppResponseBodyAppInfo
	SetRequestId(v string) *CreateAppResponseBody
	GetRequestId() *string
}

type CreateAppResponseBody struct {
	AppInfo *CreateAppResponseBodyAppInfo `json:"AppInfo,omitempty" xml:"AppInfo,omitempty" type:"Struct"`
	// example:
	//
	// 126D4DDD-05A5-49B1-B18C-39C4A929****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) GetAppInfo() *CreateAppResponseBodyAppInfo {
	return s.AppInfo
}

func (s *CreateAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAppResponseBody) SetAppInfo(v *CreateAppResponseBodyAppInfo) *CreateAppResponseBody {
	s.AppInfo = v
	return s
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAppResponseBodyAppInfo struct {
	// example:
	//
	// 123456
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// com.test.ios
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2020-12-16T06:25:52Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// com.test.android
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// example:
	//
	// 123456
	ProductId *int32 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateAppResponseBodyAppInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAppResponseBodyAppInfo) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyAppInfo) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateAppResponseBodyAppInfo) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateAppResponseBodyAppInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateAppResponseBodyAppInfo) GetDescription() *string {
	return s.Description
}

func (s *CreateAppResponseBodyAppInfo) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *CreateAppResponseBodyAppInfo) GetName() *string {
	return s.Name
}

func (s *CreateAppResponseBodyAppInfo) GetPackageName() *string {
	return s.PackageName
}

func (s *CreateAppResponseBodyAppInfo) GetProductId() *int32 {
	return s.ProductId
}

func (s *CreateAppResponseBodyAppInfo) GetType() *int32 {
	return s.Type
}

func (s *CreateAppResponseBodyAppInfo) SetAppKey(v string) *CreateAppResponseBodyAppInfo {
	s.AppKey = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetBundleId(v string) *CreateAppResponseBodyAppInfo {
	s.BundleId = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetCreateTime(v string) *CreateAppResponseBodyAppInfo {
	s.CreateTime = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetDescription(v string) *CreateAppResponseBodyAppInfo {
	s.Description = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetModifyTime(v string) *CreateAppResponseBodyAppInfo {
	s.ModifyTime = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetName(v string) *CreateAppResponseBodyAppInfo {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetPackageName(v string) *CreateAppResponseBodyAppInfo {
	s.PackageName = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetProductId(v int32) *CreateAppResponseBodyAppInfo {
	s.ProductId = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) SetType(v int32) *CreateAppResponseBodyAppInfo {
	s.Type = &v
	return s
}

func (s *CreateAppResponseBodyAppInfo) Validate() error {
	return dara.Validate(s)
}
