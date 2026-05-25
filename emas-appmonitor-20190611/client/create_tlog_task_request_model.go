// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTlogTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliYunCurrentPk(v string) *CreateTlogTaskRequest
	GetAliYunCurrentPk() *string
	SetAliYunMainPk(v string) *CreateTlogTaskRequest
	GetAliYunMainPk() *string
	SetAliYunName(v string) *CreateTlogTaskRequest
	GetAliYunName() *string
	SetAppKey(v int64) *CreateTlogTaskRequest
	GetAppKey() *int64
	SetAppVersion(v string) *CreateTlogTaskRequest
	GetAppVersion() *string
	SetBeginDate(v int64) *CreateTlogTaskRequest
	GetBeginDate() *int64
	SetBrand(v string) *CreateTlogTaskRequest
	GetBrand() *string
	SetCity(v string) *CreateTlogTaskRequest
	GetCity() *string
	SetClusterId(v string) *CreateTlogTaskRequest
	GetClusterId() *string
	SetCollectionNums(v int64) *CreateTlogTaskRequest
	GetCollectionNums() *int64
	SetDays(v string) *CreateTlogTaskRequest
	GetDays() *string
	SetDeviceJson(v string) *CreateTlogTaskRequest
	GetDeviceJson() *string
	SetEndDate(v int64) *CreateTlogTaskRequest
	GetEndDate() *int64
	SetErrorType(v string) *CreateTlogTaskRequest
	GetErrorType() *string
	SetModel(v string) *CreateTlogTaskRequest
	GetModel() *string
	SetNotifySettingJson(v string) *CreateTlogTaskRequest
	GetNotifySettingJson() *string
	SetOs(v string) *CreateTlogTaskRequest
	GetOs() *string
	SetOsVersion(v string) *CreateTlogTaskRequest
	GetOsVersion() *string
	SetSourceType(v string) *CreateTlogTaskRequest
	GetSourceType() *string
	SetTaskName(v string) *CreateTlogTaskRequest
	GetTaskName() *string
}

type CreateTlogTaskRequest struct {
	// example:
	//
	// 20000000
	AliYunCurrentPk *string `json:"AliYunCurrentPk,omitempty" xml:"AliYunCurrentPk,omitempty"`
	// example:
	//
	// 20000000
	AliYunMainPk *string `json:"AliYunMainPk,omitempty" xml:"AliYunMainPk,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// emas_user
	AliYunName *string `json:"AliYunName,omitempty" xml:"AliYunName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 24780725
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 1.0.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 1778860800000
	BeginDate *int64 `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// example:
	//
	// Redmi
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	City  *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// temp
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 10
	CollectionNums *int64 `json:"CollectionNums,omitempty" xml:"CollectionNums,omitempty"`
	// example:
	//
	// 1
	Days       *string `json:"Days,omitempty" xml:"Days,omitempty"`
	DeviceJson *string `json:"DeviceJson,omitempty" xml:"DeviceJson,omitempty"`
	// example:
	//
	// 1779465599999
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// temp
	ErrorType *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	// example:
	//
	// iPhone16
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// {"First":true,"Completed":true,"NotifyPhone":true,"NotifyEmail":true,"NotifyWebhook":true,"ContactGroupIds":[1]}
	NotifySettingJson *string `json:"NotifySettingJson,omitempty" xml:"NotifySettingJson,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 1.0.0
	OsVersion *string `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	// example:
	//
	// USER
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	TaskName   *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s CreateTlogTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTlogTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTlogTaskRequest) GetAliYunCurrentPk() *string {
	return s.AliYunCurrentPk
}

func (s *CreateTlogTaskRequest) GetAliYunMainPk() *string {
	return s.AliYunMainPk
}

func (s *CreateTlogTaskRequest) GetAliYunName() *string {
	return s.AliYunName
}

func (s *CreateTlogTaskRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *CreateTlogTaskRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *CreateTlogTaskRequest) GetBeginDate() *int64 {
	return s.BeginDate
}

func (s *CreateTlogTaskRequest) GetBrand() *string {
	return s.Brand
}

func (s *CreateTlogTaskRequest) GetCity() *string {
	return s.City
}

func (s *CreateTlogTaskRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateTlogTaskRequest) GetCollectionNums() *int64 {
	return s.CollectionNums
}

func (s *CreateTlogTaskRequest) GetDays() *string {
	return s.Days
}

func (s *CreateTlogTaskRequest) GetDeviceJson() *string {
	return s.DeviceJson
}

func (s *CreateTlogTaskRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *CreateTlogTaskRequest) GetErrorType() *string {
	return s.ErrorType
}

func (s *CreateTlogTaskRequest) GetModel() *string {
	return s.Model
}

func (s *CreateTlogTaskRequest) GetNotifySettingJson() *string {
	return s.NotifySettingJson
}

func (s *CreateTlogTaskRequest) GetOs() *string {
	return s.Os
}

func (s *CreateTlogTaskRequest) GetOsVersion() *string {
	return s.OsVersion
}

func (s *CreateTlogTaskRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateTlogTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateTlogTaskRequest) SetAliYunCurrentPk(v string) *CreateTlogTaskRequest {
	s.AliYunCurrentPk = &v
	return s
}

func (s *CreateTlogTaskRequest) SetAliYunMainPk(v string) *CreateTlogTaskRequest {
	s.AliYunMainPk = &v
	return s
}

func (s *CreateTlogTaskRequest) SetAliYunName(v string) *CreateTlogTaskRequest {
	s.AliYunName = &v
	return s
}

func (s *CreateTlogTaskRequest) SetAppKey(v int64) *CreateTlogTaskRequest {
	s.AppKey = &v
	return s
}

func (s *CreateTlogTaskRequest) SetAppVersion(v string) *CreateTlogTaskRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateTlogTaskRequest) SetBeginDate(v int64) *CreateTlogTaskRequest {
	s.BeginDate = &v
	return s
}

func (s *CreateTlogTaskRequest) SetBrand(v string) *CreateTlogTaskRequest {
	s.Brand = &v
	return s
}

func (s *CreateTlogTaskRequest) SetCity(v string) *CreateTlogTaskRequest {
	s.City = &v
	return s
}

func (s *CreateTlogTaskRequest) SetClusterId(v string) *CreateTlogTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateTlogTaskRequest) SetCollectionNums(v int64) *CreateTlogTaskRequest {
	s.CollectionNums = &v
	return s
}

func (s *CreateTlogTaskRequest) SetDays(v string) *CreateTlogTaskRequest {
	s.Days = &v
	return s
}

func (s *CreateTlogTaskRequest) SetDeviceJson(v string) *CreateTlogTaskRequest {
	s.DeviceJson = &v
	return s
}

func (s *CreateTlogTaskRequest) SetEndDate(v int64) *CreateTlogTaskRequest {
	s.EndDate = &v
	return s
}

func (s *CreateTlogTaskRequest) SetErrorType(v string) *CreateTlogTaskRequest {
	s.ErrorType = &v
	return s
}

func (s *CreateTlogTaskRequest) SetModel(v string) *CreateTlogTaskRequest {
	s.Model = &v
	return s
}

func (s *CreateTlogTaskRequest) SetNotifySettingJson(v string) *CreateTlogTaskRequest {
	s.NotifySettingJson = &v
	return s
}

func (s *CreateTlogTaskRequest) SetOs(v string) *CreateTlogTaskRequest {
	s.Os = &v
	return s
}

func (s *CreateTlogTaskRequest) SetOsVersion(v string) *CreateTlogTaskRequest {
	s.OsVersion = &v
	return s
}

func (s *CreateTlogTaskRequest) SetSourceType(v string) *CreateTlogTaskRequest {
	s.SourceType = &v
	return s
}

func (s *CreateTlogTaskRequest) SetTaskName(v string) *CreateTlogTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateTlogTaskRequest) Validate() error {
	return dara.Validate(s)
}
