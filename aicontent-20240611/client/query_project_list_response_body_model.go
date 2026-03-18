// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryProjectListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*QueryProjectListResponseBodyData) *QueryProjectListResponseBody
	GetData() []*QueryProjectListResponseBodyData
	SetErrCode(v string) *QueryProjectListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *QueryProjectListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *QueryProjectListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *QueryProjectListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryProjectListResponseBody
	GetSuccess() *bool
}

type QueryProjectListResponseBody struct {
	// example:
	//
	// []
	Data []*QueryProjectListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s QueryProjectListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBody) GetData() []*QueryProjectListResponseBodyData {
	return s.Data
}

func (s *QueryProjectListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *QueryProjectListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *QueryProjectListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryProjectListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryProjectListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryProjectListResponseBody) SetData(v []*QueryProjectListResponseBodyData) *QueryProjectListResponseBody {
	s.Data = v
	return s
}

func (s *QueryProjectListResponseBody) SetErrCode(v string) *QueryProjectListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *QueryProjectListResponseBody) SetErrMessage(v string) *QueryProjectListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *QueryProjectListResponseBody) SetHttpStatusCode(v int32) *QueryProjectListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryProjectListResponseBody) SetRequestId(v string) *QueryProjectListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryProjectListResponseBody) SetSuccess(v bool) *QueryProjectListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryProjectListResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryProjectListResponseBodyData struct {
	// example:
	//
	// 2025-02-18 12:10:22
	CreateTime  *string                                        `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ProjectApps []*QueryProjectListResponseBodyDataProjectApps `json:"ProjectApps,omitempty" xml:"ProjectApps,omitempty" type:"Repeated"`
	// example:
	//
	// 4910
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// MyProject
	ProjectName *string                                       `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectSDK  []*QueryProjectListResponseBodyDataProjectSDK `json:"ProjectSDK,omitempty" xml:"ProjectSDK,omitempty" type:"Repeated"`
	// example:
	//
	// WebApplication
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
}

func (s QueryProjectListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryProjectListResponseBodyData) GetProjectApps() []*QueryProjectListResponseBodyDataProjectApps {
	return s.ProjectApps
}

func (s *QueryProjectListResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryProjectListResponseBodyData) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryProjectListResponseBodyData) GetProjectSDK() []*QueryProjectListResponseBodyDataProjectSDK {
	return s.ProjectSDK
}

func (s *QueryProjectListResponseBodyData) GetProjectType() *string {
	return s.ProjectType
}

func (s *QueryProjectListResponseBodyData) SetCreateTime(v string) *QueryProjectListResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectApps(v []*QueryProjectListResponseBodyDataProjectApps) *QueryProjectListResponseBodyData {
	s.ProjectApps = v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectId(v string) *QueryProjectListResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectName(v string) *QueryProjectListResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectSDK(v []*QueryProjectListResponseBodyDataProjectSDK) *QueryProjectListResponseBodyData {
	s.ProjectSDK = v
	return s
}

func (s *QueryProjectListResponseBodyData) SetProjectType(v string) *QueryProjectListResponseBodyData {
	s.ProjectType = &v
	return s
}

func (s *QueryProjectListResponseBodyData) Validate() error {
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

type QueryProjectListResponseBodyDataProjectApps struct {
	ApplicationAccessIds []*QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds `json:"ApplicationAccessIds,omitempty" xml:"ApplicationAccessIds,omitempty" type:"Repeated"`
	// example:
	//
	// 4700
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 4747
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s QueryProjectListResponseBodyDataProjectApps) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectListResponseBodyDataProjectApps) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBodyDataProjectApps) GetApplicationAccessIds() []*QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds {
	return s.ApplicationAccessIds
}

func (s *QueryProjectListResponseBodyDataProjectApps) GetId() *string {
	return s.Id
}

func (s *QueryProjectListResponseBodyDataProjectApps) GetProjectId() *string {
	return s.ProjectId
}

func (s *QueryProjectListResponseBodyDataProjectApps) SetApplicationAccessIds(v []*QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) *QueryProjectListResponseBodyDataProjectApps {
	s.ApplicationAccessIds = v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectApps) SetId(v string) *QueryProjectListResponseBodyDataProjectApps {
	s.Id = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectApps) SetProjectId(v string) *QueryProjectListResponseBodyDataProjectApps {
	s.ProjectId = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectApps) Validate() error {
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

type QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds struct {
	// example:
	//
	// 1234567890
	ApplicationAccessId *string `json:"applicationAccessId,omitempty" xml:"applicationAccessId,omitempty"`
	// example:
	//
	// MyAppSecret
	ApplicationAccessSecret *string `json:"applicationAccessSecret,omitempty" xml:"applicationAccessSecret,omitempty"`
}

func (s QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) GetApplicationAccessId() *string {
	return s.ApplicationAccessId
}

func (s *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) GetApplicationAccessSecret() *string {
	return s.ApplicationAccessSecret
}

func (s *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessId(v string) *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessId = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) SetApplicationAccessSecret(v string) *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds {
	s.ApplicationAccessSecret = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectAppsApplicationAccessIds) Validate() error {
	return dara.Validate(s)
}

type QueryProjectListResponseBodyDataProjectSDK struct {
	// example:
	//
	// 2024-07-16T08:23:19Z
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
	// GO AUTH
	SdkName *string `json:"SdkName,omitempty" xml:"SdkName,omitempty"`
	// example:
	//
	// http://demo.com/sdk.zip
	SdkUrl *string `json:"SdkUrl,omitempty" xml:"SdkUrl,omitempty"`
	// example:
	//
	// .3.52
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
}

func (s QueryProjectListResponseBodyDataProjectSDK) String() string {
	return dara.Prettify(s)
}

func (s QueryProjectListResponseBodyDataProjectSDK) GoString() string {
	return s.String()
}

func (s *QueryProjectListResponseBodyDataProjectSDK) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryProjectListResponseBodyDataProjectSDK) GetDemoUrl() *string {
	return s.DemoUrl
}

func (s *QueryProjectListResponseBodyDataProjectSDK) GetDeployMode() *string {
	return s.DeployMode
}

func (s *QueryProjectListResponseBodyDataProjectSDK) GetDevelopLanguage() *string {
	return s.DevelopLanguage
}

func (s *QueryProjectListResponseBodyDataProjectSDK) GetDocUrl() *string {
	return s.DocUrl
}

func (s *QueryProjectListResponseBodyDataProjectSDK) GetSdkName() *string {
	return s.SdkName
}

func (s *QueryProjectListResponseBodyDataProjectSDK) GetSdkUrl() *string {
	return s.SdkUrl
}

func (s *QueryProjectListResponseBodyDataProjectSDK) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetCreateTime(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.CreateTime = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetDemoUrl(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.DemoUrl = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetDeployMode(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.DeployMode = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetDevelopLanguage(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.DevelopLanguage = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetDocUrl(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.DocUrl = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetSdkName(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.SdkName = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetSdkUrl(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.SdkUrl = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) SetSdkVersion(v string) *QueryProjectListResponseBodyDataProjectSDK {
	s.SdkVersion = &v
	return s
}

func (s *QueryProjectListResponseBodyDataProjectSDK) Validate() error {
	return dara.Validate(s)
}
