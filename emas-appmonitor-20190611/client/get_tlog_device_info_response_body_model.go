// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogDeviceInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v int64) *GetTlogDeviceInfoResponseBody
	GetErrorCode() *int64
	SetMessage(v string) *GetTlogDeviceInfoResponseBody
	GetMessage() *string
	SetModel(v *GetTlogDeviceInfoResponseBodyModel) *GetTlogDeviceInfoResponseBody
	GetModel() *GetTlogDeviceInfoResponseBodyModel
	SetRequestId(v string) *GetTlogDeviceInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTlogDeviceInfoResponseBody
	GetSuccess() *bool
}

type GetTlogDeviceInfoResponseBody struct {
	// example:
	//
	// 500
	ErrorCode *int64 `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// successful
	Message *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *GetTlogDeviceInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// AB8AB5EC-9636-421D-AE7C-BB227DFC95B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTlogDeviceInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTlogDeviceInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetTlogDeviceInfoResponseBody) GetErrorCode() *int64 {
	return s.ErrorCode
}

func (s *GetTlogDeviceInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTlogDeviceInfoResponseBody) GetModel() *GetTlogDeviceInfoResponseBodyModel {
	return s.Model
}

func (s *GetTlogDeviceInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTlogDeviceInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTlogDeviceInfoResponseBody) SetErrorCode(v int64) *GetTlogDeviceInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBody) SetMessage(v string) *GetTlogDeviceInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBody) SetModel(v *GetTlogDeviceInfoResponseBodyModel) *GetTlogDeviceInfoResponseBody {
	s.Model = v
	return s
}

func (s *GetTlogDeviceInfoResponseBody) SetRequestId(v string) *GetTlogDeviceInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBody) SetSuccess(v bool) *GetTlogDeviceInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTlogDeviceInfoResponseBodyModel struct {
	// example:
	//
	// 1
	AppBuild *string `json:"AppBuild,omitempty" xml:"AppBuild,omitempty"`
	// example:
	//
	// 1001@iphoneos
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// appKey
	//
	// example:
	//
	// 1001
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// V20250224102631
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// Hinova
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// example:
	//
	// 1739808000000
	ClientTime *int64 `json:"ClientTime,omitempty" xml:"ClientTime,omitempty"`
	// example:
	//
	// ad-0001t9c1b6y48kqcd44s-105
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	// example:
	//
	// ASUS_X00GD
	DeviceModel *string `json:"DeviceModel,omitempty" xml:"DeviceModel,omitempty"`
	Geo         *string `json:"Geo,omitempty" xml:"Geo,omitempty"`
	// example:
	//
	// 10
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 460049842500442
	Imsi *string `json:"Imsi,omitempty" xml:"Imsi,omitempty"`
	// example:
	//
	// 0:0:0:0:0:ffff:8ccd:be5
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 1
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// example:
	//
	// 1739808000000
	ServerTime *int64 `json:"ServerTime,omitempty" xml:"ServerTime,omitempty"`
	// example:
	//
	// 1739808000000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// userId
	//
	// example:
	//
	// userId
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// userNick
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetTlogDeviceInfoResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s GetTlogDeviceInfoResponseBodyModel) GoString() string {
	return s.String()
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetAppBuild() *string {
	return s.AppBuild
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetAppId() *string {
	return s.AppId
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetAppKey() *string {
	return s.AppKey
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetBrand() *string {
	return s.Brand
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetClientTime() *int64 {
	return s.ClientTime
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetDeviceModel() *string {
	return s.DeviceModel
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetGeo() *string {
	return s.Geo
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetId() *string {
	return s.Id
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetImsi() *string {
	return s.Imsi
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetIp() *string {
	return s.Ip
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetOs() *string {
	return s.Os
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetOsVersion() *string {
	return s.OsVersion
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetServerTime() *int64 {
	return s.ServerTime
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetUserId() *string {
	return s.UserId
}

func (s *GetTlogDeviceInfoResponseBodyModel) GetUserName() *string {
	return s.UserName
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetAppBuild(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.AppBuild = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetAppId(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.AppId = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetAppKey(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.AppKey = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetAppVersion(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.AppVersion = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetBrand(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.Brand = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetClientTime(v int64) *GetTlogDeviceInfoResponseBodyModel {
	s.ClientTime = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetDeviceId(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.DeviceId = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetDeviceModel(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.DeviceModel = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetGeo(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.Geo = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetId(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.Id = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetImsi(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.Imsi = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetIp(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.Ip = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetOs(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.Os = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetOsVersion(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.OsVersion = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetServerTime(v int64) *GetTlogDeviceInfoResponseBodyModel {
	s.ServerTime = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetUpdateTime(v int64) *GetTlogDeviceInfoResponseBodyModel {
	s.UpdateTime = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetUserId(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.UserId = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) SetUserName(v string) *GetTlogDeviceInfoResponseBodyModel {
	s.UserName = &v
	return s
}

func (s *GetTlogDeviceInfoResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
