// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateApplicationResponseBody
	GetCode() *int32
	SetData(v *CreateApplicationResponseBodyData) *CreateApplicationResponseBody
	GetData() *CreateApplicationResponseBodyData
	SetHttpStatusCode(v int32) *CreateApplicationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateApplicationResponseBody
	GetSuccess() *string
}

type CreateApplicationResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	//
	// example:
	//
	// data
	Data *CreateApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 68D91223-CCE9-5F9C-B538-20F617DA48B1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The request result. Valid values:
	//
	// - `true`: The request is successful.
	//
	// - `false`: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateApplicationResponseBody) GetData() *CreateApplicationResponseBodyData {
	return s.Data
}

func (s *CreateApplicationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateApplicationResponseBody) SetCode(v int32) *CreateApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApplicationResponseBody) SetData(v *CreateApplicationResponseBodyData) *CreateApplicationResponseBody {
	s.Data = v
	return s
}

func (s *CreateApplicationResponseBody) SetHttpStatusCode(v int32) *CreateApplicationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApplicationResponseBody) SetMessage(v string) *CreateApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApplicationResponseBody) SetRequestId(v string) *CreateApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationResponseBody) SetSuccess(v string) *CreateApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApplicationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApplicationResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// xxxxxxxx@xxxxxxxxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// spring-cloud-a
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1631001140913
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The extra information.
	//
	// example:
	//
	// {\\"rpcTypes\\":[\\"dubbo\\",\\"springCloud\\"]}
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The programming language of the application.
	//
	// example:
	//
	// JAVA
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The license key corresponding to the user.
	//
	// example:
	//
	// xxxxxxxx@xxxxxxxxxx
	LicenseKey *string `json:"LicenseKey,omitempty" xml:"LicenseKey,omitempty"`
	// The name of the MSE namespace.
	//
	// example:
	//
	// prod
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The deployment method of the application. Valid values:
	//
	// - ACK: Container Service for Kubernetes (ACK).
	//
	// - Normal: Other access methods.
	//
	// example:
	//
	// ACK
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The application status. A value of 1 indicates normal.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1632979237663
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The Alibaba Cloud user ID.
	//
	// example:
	//
	// 1888888888
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s CreateApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateApplicationResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *CreateApplicationResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *CreateApplicationResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *CreateApplicationResponseBodyData) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *CreateApplicationResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *CreateApplicationResponseBodyData) GetLicenseKey() *string {
	return s.LicenseKey
}

func (s *CreateApplicationResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateApplicationResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *CreateApplicationResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *CreateApplicationResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *CreateApplicationResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *CreateApplicationResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *CreateApplicationResponseBodyData) SetAppId(v string) *CreateApplicationResponseBodyData {
	s.AppId = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetAppName(v string) *CreateApplicationResponseBodyData {
	s.AppName = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetCreateTime(v int64) *CreateApplicationResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetExtraInfo(v string) *CreateApplicationResponseBodyData {
	s.ExtraInfo = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetLanguage(v string) *CreateApplicationResponseBodyData {
	s.Language = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetLicenseKey(v string) *CreateApplicationResponseBodyData {
	s.LicenseKey = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetNamespace(v string) *CreateApplicationResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetRegionId(v string) *CreateApplicationResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetSource(v string) *CreateApplicationResponseBodyData {
	s.Source = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetStatus(v int32) *CreateApplicationResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetUpdateTime(v int64) *CreateApplicationResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetUserId(v string) *CreateApplicationResponseBodyData {
	s.UserId = &v
	return s
}

func (s *CreateApplicationResponseBodyData) SetVersion(v string) *CreateApplicationResponseBodyData {
	s.Version = &v
	return s
}

func (s *CreateApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
