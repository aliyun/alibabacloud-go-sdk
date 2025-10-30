// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServicePublishedApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataServicePublishedApisResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListDataServicePublishedApisResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataServicePublishedApisResponseBody
	GetMessage() *string
	SetPageResult(v *ListDataServicePublishedApisResponseBodyPageResult) *ListDataServicePublishedApisResponseBody
	GetPageResult() *ListDataServicePublishedApisResponseBodyPageResult
	SetRequestId(v string) *ListDataServicePublishedApisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataServicePublishedApisResponseBody
	GetSuccess() *bool
}

type ListDataServicePublishedApisResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message    *string                                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListDataServicePublishedApisResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataServicePublishedApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataServicePublishedApisResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataServicePublishedApisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataServicePublishedApisResponseBody) GetPageResult() *ListDataServicePublishedApisResponseBodyPageResult {
	return s.PageResult
}

func (s *ListDataServicePublishedApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataServicePublishedApisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataServicePublishedApisResponseBody) SetCode(v string) *ListDataServicePublishedApisResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetHttpStatusCode(v int32) *ListDataServicePublishedApisResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetMessage(v string) *ListDataServicePublishedApisResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetPageResult(v *ListDataServicePublishedApisResponseBodyPageResult) *ListDataServicePublishedApisResponseBody {
	s.PageResult = v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetRequestId(v string) *ListDataServicePublishedApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) SetSuccess(v bool) *ListDataServicePublishedApisResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataServicePublishedApisResponseBodyPageResult struct {
	ApiList []*ListDataServicePublishedApisResponseBodyPageResultApiList `json:"ApiList,omitempty" xml:"ApiList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataServicePublishedApisResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyPageResult) GetApiList() []*ListDataServicePublishedApisResponseBodyPageResultApiList {
	return s.ApiList
}

func (s *ListDataServicePublishedApisResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataServicePublishedApisResponseBodyPageResult) SetApiList(v []*ListDataServicePublishedApisResponseBodyPageResultApiList) *ListDataServicePublishedApisResponseBodyPageResult {
	s.ApiList = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResult) SetTotalCount(v int32) *ListDataServicePublishedApisResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResult) Validate() error {
	if s.ApiList != nil {
		for _, item := range s.ApiList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServicePublishedApisResponseBodyPageResultApiList struct {
	// example:
	//
	// 1022
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// test
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// example:
	//
	// 1
	AppCount    *int32                                                                  `json:"AppCount,omitempty" xml:"AppCount,omitempty"`
	AppInfoList []*ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList `json:"AppInfoList,omitempty" xml:"AppInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	ApplyStatus *int32 `json:"ApplyStatus,omitempty" xml:"ApplyStatus,omitempty"`
	// example:
	//
	// 21
	CallCount *int32 `json:"CallCount,omitempty" xml:"CallCount,omitempty"`
	// example:
	//
	// 1
	CreateType *int32 `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	// example:
	//
	// 0 0 0/1 	- 	- *
	CustomUpdateRate *string `json:"CustomUpdateRate,omitempty" xml:"CustomUpdateRate,omitempty"`
	// example:
	//
	// 2025-06-30 08:00:00
	DeployTime *string `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// example:
	//
	// test xx
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	ExecuteMode *int32 `json:"ExecuteMode,omitempty" xml:"ExecuteMode,omitempty"`
	// example:
	//
	// 102131
	GroupId *int32 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 1022
	LogicUnitNo *int64 `json:"LogicUnitNo,omitempty" xml:"LogicUnitNo,omitempty"`
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// 30012011
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 张三
	OwnerUserName *string `json:"OwnerUserName,omitempty" xml:"OwnerUserName,omitempty"`
	// example:
	//
	// 102101
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1
	UpdateRate *int32 `json:"UpdateRate,omitempty" xml:"UpdateRate,omitempty"`
	// example:
	//
	// 2025-06-30 08:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListDataServicePublishedApisResponseBodyPageResultApiList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyPageResultApiList) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetApiName() *string {
	return s.ApiName
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetAppCount() *int32 {
	return s.AppCount
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetAppInfoList() []*ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList {
	return s.AppInfoList
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetApplyStatus() *int32 {
	return s.ApplyStatus
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetCallCount() *int32 {
	return s.CallCount
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetCreateType() *int32 {
	return s.CreateType
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetCustomUpdateRate() *string {
	return s.CustomUpdateRate
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetDeployTime() *string {
	return s.DeployTime
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetDescription() *string {
	return s.Description
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetExecuteMode() *int32 {
	return s.ExecuteMode
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetGroupId() *int32 {
	return s.GroupId
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetGroupName() *string {
	return s.GroupName
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetLogicUnitNo() *int64 {
	return s.LogicUnitNo
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetMode() *int32 {
	return s.Mode
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetOwner() *string {
	return s.Owner
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetOwnerUserName() *string {
	return s.OwnerUserName
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetUpdateRate() *int32 {
	return s.UpdateRate
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) GetVersion() *string {
	return s.Version
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetApiId(v int64) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.ApiId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetApiName(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.ApiName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetAppCount(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.AppCount = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetAppInfoList(v []*ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.AppInfoList = v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetApplyStatus(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.ApplyStatus = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetCallCount(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.CallCount = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetCreateType(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.CreateType = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetCustomUpdateRate(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.CustomUpdateRate = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetDeployTime(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.DeployTime = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetDescription(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.Description = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetExecuteMode(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.ExecuteMode = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetGroupId(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.GroupId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetGroupName(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.GroupName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetLogicUnitNo(v int64) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.LogicUnitNo = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetMode(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.Mode = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetOwner(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.Owner = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetOwnerUserName(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.OwnerUserName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetProjectId(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.ProjectId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetProjectName(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.ProjectName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetUpdateRate(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.UpdateRate = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetUpdateTime(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.UpdateTime = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) SetVersion(v string) *ListDataServicePublishedApisResponseBodyPageResultApiList {
	s.Version = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiList) Validate() error {
	if s.AppInfoList != nil {
		for _, item := range s.AppInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList struct {
	// example:
	//
	// 10211
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// appKey
	//
	// example:
	//
	// 1201
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// test
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) String() string {
	return dara.Prettify(s)
}

func (s ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) GoString() string {
	return s.String()
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) GetAppId() *int32 {
	return s.AppId
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) GetAppKey() *int64 {
	return s.AppKey
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) GetAppName() *string {
	return s.AppName
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) SetAppId(v int32) *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList {
	s.AppId = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) SetAppKey(v int64) *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList {
	s.AppKey = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) SetAppName(v string) *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList {
	s.AppName = &v
	return s
}

func (s *ListDataServicePublishedApisResponseBodyPageResultApiListAppInfoList) Validate() error {
	return dara.Validate(s)
}
