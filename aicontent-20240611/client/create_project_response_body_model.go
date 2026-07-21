// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateProjectResponseBodyData) *CreateProjectResponseBody
	GetData() *CreateProjectResponseBodyData
	SetErrCode(v string) *CreateProjectResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateProjectResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CreateProjectResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CreateProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProjectResponseBody
	GetSuccess() *bool
}

type CreateProjectResponseBody struct {
	// The returned data object.
	//
	// example:
	//
	// []
	Data *CreateProjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) GetData() *CreateProjectResponseBodyData {
	return s.Data
}

func (s *CreateProjectResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateProjectResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProjectResponseBody) SetData(v *CreateProjectResponseBodyData) *CreateProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateProjectResponseBody) SetErrCode(v string) *CreateProjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateProjectResponseBody) SetErrMessage(v string) *CreateProjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateProjectResponseBody) SetHttpStatusCode(v int32) *CreateProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectResponseBodyData struct {
	// The project\\"s creation time.
	//
	// example:
	//
	// 2023-02-15T09:17:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The applications in the project.
	ProjectApps []*CreateProjectResponseBodyDataProjectApps `json:"ProjectApps,omitempty" xml:"ProjectApps,omitempty" type:"Repeated"`
	// The project ID.
	//
	// example:
	//
	// 124187
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The project name.
	//
	// example:
	//
	// MyProject
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The SDKs for the project.
	ProjectSDK []*CreateProjectResponseBodyDataProjectSDK `json:"ProjectSDK,omitempty" xml:"ProjectSDK,omitempty" type:"Repeated"`
	// The project type.
	//
	// example:
	//
	// WebApplication
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s CreateProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateProjectResponseBodyData) GetProjectApps() []*CreateProjectResponseBodyDataProjectApps {
	return s.ProjectApps
}

func (s *CreateProjectResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateProjectResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateProjectResponseBodyData) GetProjectSDK() []*CreateProjectResponseBodyDataProjectSDK {
	return s.ProjectSDK
}

func (s *CreateProjectResponseBodyData) GetProjectType() *string {
	return s.ProjectType
}

func (s *CreateProjectResponseBodyData) SetCreateTime(v string) *CreateProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectApps(v []*CreateProjectResponseBodyDataProjectApps) *CreateProjectResponseBodyData {
	s.ProjectApps = v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectId(v string) *CreateProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectName(v string) *CreateProjectResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectSDK(v []*CreateProjectResponseBodyDataProjectSDK) *CreateProjectResponseBodyData {
	s.ProjectSDK = v
	return s
}

func (s *CreateProjectResponseBodyData) SetProjectType(v string) *CreateProjectResponseBodyData {
	s.ProjectType = &v
	return s
}

func (s *CreateProjectResponseBodyData) Validate() error {
	if s.ProjectApps != nil {
		for _, item := range s.ProjectApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProjectSDK != nil {
		for _, item := range s.ProjectSDK {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateProjectResponseBodyDataProjectApps struct {
	// The access credentials for the application.
	ApplicationAccessIds []*CreateProjectResponseBodyDataProjectAppsApplicationAccessIds `json:"ApplicationAccessIds,omitempty" xml:"ApplicationAccessIds,omitempty" type:"Repeated"`
	// The internal ID of the application.
	//
	// example:
	//
	// 4867
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 4910
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateProjectResponseBodyDataProjectApps) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyDataProjectApps) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataProjectApps) GetApplicationAccessIds() []*CreateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	return s.ApplicationAccessIds
}

func (s *CreateProjectResponseBodyDataProjectApps) GetId() *string {
	return s.Id
}

func (s *CreateProjectResponseBodyDataProjectApps) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateProjectResponseBodyDataProjectApps) SetApplicationAccessIds(v []*CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) *CreateProjectResponseBodyDataProjectApps {
	s.ApplicationAccessIds = v
	return s
}

func (s *CreateProjectResponseBodyDataProjectApps) SetId(v string) *CreateProjectResponseBodyDataProjectApps {
	s.Id = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectApps) SetProjectId(v string) *CreateProjectResponseBodyDataProjectApps {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectApps) Validate() error {
	if s.ApplicationAccessIds != nil {
		for _, item := range s.ApplicationAccessIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateProjectResponseBodyDataProjectAppsApplicationAccessIds struct {
	// The app key.
	//
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// The app secret. This parameter is returned only for requests sent from the console.
	//
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) GetApplicationAccessSecret() *string {
	return s.ApplicationAccessSecret
}

func (s *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessId(v string) *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessId = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessSecret(v string) *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessSecret = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectAppsApplicationAccessIds) Validate() error {
	return dara.Validate(s)
}

type CreateProjectResponseBodyDataProjectSDK struct {
	// The SDK\\"s creation time.
	//
	// example:
	//
	// 2023-02-15T09:17:39Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The demo\\"s URL.
	//
	// example:
	//
	// http://demo.com/demo
	DemoUrl *string `json:"DemoUrl,omitempty" xml:"DemoUrl,omitempty"`
	// The deployment mode. Valid values: client or server.
	//
	// example:
	//
	// 客户端
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// The development language.
	//
	// example:
	//
	// C++
	DevelopLanguage *string `json:"DevelopLanguage,omitempty" xml:"DevelopLanguage,omitempty"`
	// The SDK\\"s documentation URL.
	//
	// example:
	//
	// http://demo.com/doc
	DocUrl *string `json:"DocUrl,omitempty" xml:"DocUrl,omitempty"`
	// The name of the SDK.
	//
	// example:
	//
	// C SDK
	SdkName *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	// The SDK\\"s download URL.
	//
	// example:
	//
	// http://demo.com/sdk.zip
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// The SDK version.
	//
	// example:
	//
	// 4.12.8
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s CreateProjectResponseBodyDataProjectSDK) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyDataProjectSDK) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyDataProjectSDK) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateProjectResponseBodyDataProjectSDK) GetDemoUrl() *string {
	return s.DemoUrl
}

func (s *CreateProjectResponseBodyDataProjectSDK) GetDeployMode() *string {
	return s.DeployMode
}

func (s *CreateProjectResponseBodyDataProjectSDK) GetDevelopLanguage() *string {
	return s.DevelopLanguage
}

func (s *CreateProjectResponseBodyDataProjectSDK) GetDocUrl() *string {
	return s.DocUrl
}

func (s *CreateProjectResponseBodyDataProjectSDK) GetSdkName() *string {
	return s.SdkName
}

func (s *CreateProjectResponseBodyDataProjectSDK) GetSdkUrl() *string {
	return s.SdkUrl
}

func (s *CreateProjectResponseBodyDataProjectSDK) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetCreateTime(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.CreateTime = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetDemoUrl(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.DemoUrl = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetDeployMode(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.DeployMode = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetDevelopLanguage(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.DevelopLanguage = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetDocUrl(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.DocUrl = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetSdkName(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.SdkName = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetSdkUrl(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.SdkUrl = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) SetSdkVersion(v string) *CreateProjectResponseBodyDataProjectSDK {
	s.SdkVersion = &v
	return s
}

func (s *CreateProjectResponseBodyDataProjectSDK) Validate() error {
	return dara.Validate(s)
}
