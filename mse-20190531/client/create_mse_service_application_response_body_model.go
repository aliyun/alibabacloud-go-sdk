// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMseServiceApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateMseServiceApplicationResponseBodyData) *CreateMseServiceApplicationResponseBody
	GetData() *CreateMseServiceApplicationResponseBodyData
	SetMessage(v string) *CreateMseServiceApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMseServiceApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateMseServiceApplicationResponseBody
	GetSuccess() *string
}

type CreateMseServiceApplicationResponseBody struct {
	// The data structure.
	Data *CreateMseServiceApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 904F6C04-1284-****-8ED2-FFC57E507A72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMseServiceApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMseServiceApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMseServiceApplicationResponseBody) GetData() *CreateMseServiceApplicationResponseBodyData {
	return s.Data
}

func (s *CreateMseServiceApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMseServiceApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMseServiceApplicationResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateMseServiceApplicationResponseBody) SetData(v *CreateMseServiceApplicationResponseBodyData) *CreateMseServiceApplicationResponseBody {
	s.Data = v
	return s
}

func (s *CreateMseServiceApplicationResponseBody) SetMessage(v string) *CreateMseServiceApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBody) SetRequestId(v string) *CreateMseServiceApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBody) SetSuccess(v string) *CreateMseServiceApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateMseServiceApplicationResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// xxxxxxxx@xxxxxxxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1631001140913
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The additional information.
	//
	// example:
	//
	// {}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The programming language of the application.
	//
	// example:
	//
	// JAVA
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The license key in use.
	//
	// example:
	//
	// xxxxxxxx@xxxxxxxxxx
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The source type.
	//
	// example:
	//
	// edasmsc
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status. Valid values: 1: available; 2: deleted.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 1632979237663
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 1234567890
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The version information.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateMseServiceApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMseServiceApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMseServiceApplicationResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateMseServiceApplicationResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *CreateMseServiceApplicationResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateMseServiceApplicationResponseBodyData) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *CreateMseServiceApplicationResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *CreateMseServiceApplicationResponseBodyData) GetLicenseKey() *string {
	return s.LicenseKey
}

func (s *CreateMseServiceApplicationResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateMseServiceApplicationResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *CreateMseServiceApplicationResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *CreateMseServiceApplicationResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CreateMseServiceApplicationResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *CreateMseServiceApplicationResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *CreateMseServiceApplicationResponseBodyData) SetAppId(v string) *CreateMseServiceApplicationResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetAppName(v string) *CreateMseServiceApplicationResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetCreateTime(v int64) *CreateMseServiceApplicationResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetExtraInfo(v string) *CreateMseServiceApplicationResponseBodyData {
	s.ExtraInfo = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetLanguage(v string) *CreateMseServiceApplicationResponseBodyData {
	s.Language = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetLicenseKey(v string) *CreateMseServiceApplicationResponseBodyData {
	s.LicenseKey = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetRegionId(v string) *CreateMseServiceApplicationResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetSource(v string) *CreateMseServiceApplicationResponseBodyData {
	s.Source = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetStatus(v int32) *CreateMseServiceApplicationResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetUpdateTime(v int64) *CreateMseServiceApplicationResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetUserId(v string) *CreateMseServiceApplicationResponseBodyData {
	s.UserId = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) SetVersion(v string) *CreateMseServiceApplicationResponseBodyData {
	s.Version = &v
	return s
}

func (s *CreateMseServiceApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
