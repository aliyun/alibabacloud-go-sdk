// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateProjectResponseBodyData) *UpdateProjectResponseBody
	GetData() *UpdateProjectResponseBodyData
	SetErrCode(v string) *UpdateProjectResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *UpdateProjectResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *UpdateProjectResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *UpdateProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateProjectResponseBody
	GetSuccess() *bool
}

type UpdateProjectResponseBody struct {
	// example:
	//
	// []
	Data *UpdateProjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBody) GetData() *UpdateProjectResponseBodyData {
	return s.Data
}

func (s *UpdateProjectResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *UpdateProjectResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *UpdateProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProjectResponseBody) SetData(v *UpdateProjectResponseBodyData) *UpdateProjectResponseBody {
	s.Data = v
	return s
}

func (s *UpdateProjectResponseBody) SetErrCode(v string) *UpdateProjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpdateProjectResponseBody) SetErrMessage(v string) *UpdateProjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpdateProjectResponseBody) SetHttpStatusCode(v int32) *UpdateProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateProjectResponseBody) SetRequestId(v string) *UpdateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectResponseBody) SetSuccess(v bool) *UpdateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProjectResponseBodyData struct {
	// example:
	//
	// 2024-12-10T02:07:16Z
	CreateTime  *string                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectApps []*UpdateProjectResponseBodyDataProjectApps `json:"ProjectApps,omitempty" xml:"ProjectApps,omitempty" type:"Repeated"`
	// example:
	//
	// 56160
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MyProject
	ProjectName *string                                    `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectSDK  []*UpdateProjectResponseBodyDataProjectSDK `json:"ProjectSDK,omitempty" xml:"ProjectSDK,omitempty" type:"Repeated"`
	// example:
	//
	// WebApplication
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s UpdateProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateProjectResponseBodyData) GetProjectApps() []*UpdateProjectResponseBodyDataProjectApps {
	return s.ProjectApps
}

func (s *UpdateProjectResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateProjectResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *UpdateProjectResponseBodyData) GetProjectSDK() []*UpdateProjectResponseBodyDataProjectSDK {
	return s.ProjectSDK
}

func (s *UpdateProjectResponseBodyData) GetProjectType() *string {
	return s.ProjectType
}

func (s *UpdateProjectResponseBodyData) SetCreateTime(v string) *UpdateProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectApps(v []*UpdateProjectResponseBodyDataProjectApps) *UpdateProjectResponseBodyData {
	s.ProjectApps = v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectId(v string) *UpdateProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectName(v string) *UpdateProjectResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectSDK(v []*UpdateProjectResponseBodyDataProjectSDK) *UpdateProjectResponseBodyData {
	s.ProjectSDK = v
	return s
}

func (s *UpdateProjectResponseBodyData) SetProjectType(v string) *UpdateProjectResponseBodyData {
	s.ProjectType = &v
	return s
}

func (s *UpdateProjectResponseBodyData) Validate() error {
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

type UpdateProjectResponseBodyDataProjectApps struct {
	ApplicationAccessIds []*UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds `json:"ApplicationAccessIds,omitempty" xml:"ApplicationAccessIds,omitempty" type:"Repeated"`
	// example:
	//
	// 4498
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1889
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateProjectResponseBodyDataProjectApps) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBodyDataProjectApps) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBodyDataProjectApps) GetApplicationAccessIds() []*UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	return s.ApplicationAccessIds
}

func (s *UpdateProjectResponseBodyDataProjectApps) GetId() *string {
	return s.Id
}

func (s *UpdateProjectResponseBodyDataProjectApps) GetProjectId() *string {
	return s.ProjectId
}

func (s *UpdateProjectResponseBodyDataProjectApps) SetApplicationAccessIds(v []*UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) *UpdateProjectResponseBodyDataProjectApps {
	s.ApplicationAccessIds = v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectApps) SetId(v string) *UpdateProjectResponseBodyDataProjectApps {
	s.Id = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectApps) SetProjectId(v string) *UpdateProjectResponseBodyDataProjectApps {
	s.ProjectId = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectApps) Validate() error {
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

type UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) GetApplicationAccessSecret() *string {
	return s.ApplicationAccessSecret
}

func (s *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessId(v string) *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessId = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessSecret(v string) *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessSecret = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectAppsApplicationAccessIds) Validate() error {
	return dara.Validate(s)
}

type UpdateProjectResponseBodyDataProjectSDK struct {
	// example:
	//
	// 2024-11-01T13:40:53Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// http://demo.com/demo
	DemoUrl    *string `json:"DemoUrl,omitempty" xml:"DemoUrl,omitempty"`
	DeployMode *string `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	// example:
	//
	// PHP
	DevelopLanguage *string `json:"DevelopLanguage,omitempty" xml:"DevelopLanguage,omitempty"`
	// example:
	//
	// http://demo.com/doc
	DocUrl  *string `json:"DocUrl,omitempty" xml:"DocUrl,omitempty"`
	SdkName *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	// example:
	//
	// http://demo.com/sdk.zip
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// 4.13.0
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s UpdateProjectResponseBodyDataProjectSDK) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectResponseBodyDataProjectSDK) GoString() string {
	return s.String()
}

func (s *UpdateProjectResponseBodyDataProjectSDK) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateProjectResponseBodyDataProjectSDK) GetDemoUrl() *string {
	return s.DemoUrl
}

func (s *UpdateProjectResponseBodyDataProjectSDK) GetDeployMode() *string {
	return s.DeployMode
}

func (s *UpdateProjectResponseBodyDataProjectSDK) GetDevelopLanguage() *string {
	return s.DevelopLanguage
}

func (s *UpdateProjectResponseBodyDataProjectSDK) GetDocUrl() *string {
	return s.DocUrl
}

func (s *UpdateProjectResponseBodyDataProjectSDK) GetSdkName() *string {
	return s.SdkName
}

func (s *UpdateProjectResponseBodyDataProjectSDK) GetSdkUrl() *string {
	return s.SdkUrl
}

func (s *UpdateProjectResponseBodyDataProjectSDK) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetCreateTime(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.CreateTime = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetDemoUrl(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.DemoUrl = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetDeployMode(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.DeployMode = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetDevelopLanguage(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.DevelopLanguage = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetDocUrl(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.DocUrl = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetSdkName(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.SdkName = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetSdkUrl(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.SdkUrl = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) SetSdkVersion(v string) *UpdateProjectResponseBodyDataProjectSDK {
	s.SdkVersion = &v
	return s
}

func (s *UpdateProjectResponseBodyDataProjectSDK) Validate() error {
	return dara.Validate(s)
}
