// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *QueryProjectResponseBodyData) *QueryProjectResponseBody
	GetData() *QueryProjectResponseBodyData
	SetErrCode(v string) *QueryProjectResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryProjectResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryProjectResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryProjectResponseBody
	GetSuccess() *bool
}

type QueryProjectResponseBody struct {
	// example:
	//
	// []
	Data *QueryProjectResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s QueryProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBody) GetData() *QueryProjectResponseBodyData {
	return s.Data
}

func (s *QueryProjectResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryProjectResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryProjectResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryProjectResponseBody) SetData(v *QueryProjectResponseBodyData) *QueryProjectResponseBody {
	s.Data = v
	return s
}

func (s *QueryProjectResponseBody) SetErrCode(v string) *QueryProjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryProjectResponseBody) SetErrMessage(v string) *QueryProjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryProjectResponseBody) SetHttpStatusCode(v int32) *QueryProjectResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryProjectResponseBody) SetRequestId(v string) *QueryProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProjectResponseBody) SetSuccess(v bool) *QueryProjectResponseBody {
	s.Success = &v
	return s
}

func (s *QueryProjectResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryProjectResponseBodyData struct {
	// example:
	//
	// 2024-11-01T13:40:53Z
	CreateTime  *string                                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectApps []*QueryProjectResponseBodyDataProjectApps `json:"ProjectApps,omitempty" xml:"ProjectApps,omitempty" type:"Repeated"`
	// example:
	//
	// 67055
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MyProject
	ProjectName *string                                   `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectSDK  []*QueryProjectResponseBodyDataProjectSDK `json:"ProjectSDK,omitempty" xml:"ProjectSDK,omitempty" type:"Repeated"`
	// example:
	//
	// WebApplication
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s QueryProjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryProjectResponseBodyData) GetProjectApps() []*QueryProjectResponseBodyDataProjectApps {
	return s.ProjectApps
}

func (s *QueryProjectResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryProjectResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryProjectResponseBodyData) GetProjectSDK() []*QueryProjectResponseBodyDataProjectSDK {
	return s.ProjectSDK
}

func (s *QueryProjectResponseBodyData) GetProjectType() *string {
	return s.ProjectType
}

func (s *QueryProjectResponseBodyData) SetCreateTime(v string) *QueryProjectResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectApps(v []*QueryProjectResponseBodyDataProjectApps) *QueryProjectResponseBodyData {
	s.ProjectApps = v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectId(v string) *QueryProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectName(v string) *QueryProjectResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectSDK(v []*QueryProjectResponseBodyDataProjectSDK) *QueryProjectResponseBodyData {
	s.ProjectSDK = v
	return s
}

func (s *QueryProjectResponseBodyData) SetProjectType(v string) *QueryProjectResponseBodyData {
	s.ProjectType = &v
	return s
}

func (s *QueryProjectResponseBodyData) Validate() error {
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

type QueryProjectResponseBodyDataProjectApps struct {
	ApplicationAccessIds []*QueryProjectResponseBodyDataProjectAppsApplicationAccessIds `json:"ApplicationAccessIds,omitempty" xml:"ApplicationAccessIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2144
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 159
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryProjectResponseBodyDataProjectApps) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectResponseBodyDataProjectApps) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyDataProjectApps) GetApplicationAccessIds() []*QueryProjectResponseBodyDataProjectAppsApplicationAccessIds {
	return s.ApplicationAccessIds
}

func (s *QueryProjectResponseBodyDataProjectApps) GetId() *string {
	return s.Id
}

func (s *QueryProjectResponseBodyDataProjectApps) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryProjectResponseBodyDataProjectApps) SetApplicationAccessIds(v []*QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) *QueryProjectResponseBodyDataProjectApps {
	s.ApplicationAccessIds = v
	return s
}

func (s *QueryProjectResponseBodyDataProjectApps) SetId(v string) *QueryProjectResponseBodyDataProjectApps {
	s.Id = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectApps) SetProjectId(v string) *QueryProjectResponseBodyDataProjectApps {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectApps) Validate() error {
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

type QueryProjectResponseBodyDataProjectAppsApplicationAccessIds struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) GetApplicationAccessSecret() *string {
	return s.ApplicationAccessSecret
}

func (s *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessId(v string) *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessId = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessSecret(v string) *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessSecret = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectAppsApplicationAccessIds) Validate() error {
	return dara.Validate(s)
}

type QueryProjectResponseBodyDataProjectSDK struct {
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
	// JAVA
	DevelopLanguage *string `json:"DevelopLanguage,omitempty" xml:"DevelopLanguage,omitempty"`
	// example:
	//
	// http://demo.com/doc
	DocUrl *string `json:"DocUrl,omitempty" xml:"DocUrl,omitempty"`
	// example:
	//
	// JSSDK
	SdkName *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	// example:
	//
	// http://demo.com/sdk.zip
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// 5.1.0
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s QueryProjectResponseBodyDataProjectSDK) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectResponseBodyDataProjectSDK) GoString() string {
	return s.String()
}

func (s *QueryProjectResponseBodyDataProjectSDK) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryProjectResponseBodyDataProjectSDK) GetDemoUrl() *string {
	return s.DemoUrl
}

func (s *QueryProjectResponseBodyDataProjectSDK) GetDeployMode() *string {
	return s.DeployMode
}

func (s *QueryProjectResponseBodyDataProjectSDK) GetDevelopLanguage() *string {
	return s.DevelopLanguage
}

func (s *QueryProjectResponseBodyDataProjectSDK) GetDocUrl() *string {
	return s.DocUrl
}

func (s *QueryProjectResponseBodyDataProjectSDK) GetSdkName() *string {
	return s.SdkName
}

func (s *QueryProjectResponseBodyDataProjectSDK) GetSdkUrl() *string {
	return s.SdkUrl
}

func (s *QueryProjectResponseBodyDataProjectSDK) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetCreateTime(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetDemoUrl(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.DemoUrl = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetDeployMode(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.DeployMode = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetDevelopLanguage(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.DevelopLanguage = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetDocUrl(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.DocUrl = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetSdkName(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.SdkName = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetSdkUrl(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.SdkUrl = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) SetSdkVersion(v string) *QueryProjectResponseBodyDataProjectSDK {
	s.SdkVersion = &v
	return s
}

func (s *QueryProjectResponseBodyDataProjectSDK) Validate() error {
	return dara.Validate(s)
}
