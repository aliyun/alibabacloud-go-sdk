// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertEventInfo(v *GetAlertEventResponseBodyAlertEventInfo) *GetAlertEventResponseBody
	GetAlertEventInfo() *GetAlertEventResponseBodyAlertEventInfo
	SetCode(v string) *GetAlertEventResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetAlertEventResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAlertEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAlertEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAlertEventResponseBody
	GetSuccess() *bool
}

type GetAlertEventResponseBody struct {
	AlertEventInfo *GetAlertEventResponseBodyAlertEventInfo `json:"AlertEventInfo,omitempty" xml:"AlertEventInfo,omitempty" type:"Struct"`
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
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAlertEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBody) GetAlertEventInfo() *GetAlertEventResponseBodyAlertEventInfo {
	return s.AlertEventInfo
}

func (s *GetAlertEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAlertEventResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAlertEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAlertEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlertEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAlertEventResponseBody) SetAlertEventInfo(v *GetAlertEventResponseBodyAlertEventInfo) *GetAlertEventResponseBody {
	s.AlertEventInfo = v
	return s
}

func (s *GetAlertEventResponseBody) SetCode(v string) *GetAlertEventResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlertEventResponseBody) SetHttpStatusCode(v int32) *GetAlertEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAlertEventResponseBody) SetMessage(v string) *GetAlertEventResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlertEventResponseBody) SetRequestId(v string) *GetAlertEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlertEventResponseBody) SetSuccess(v bool) *GetAlertEventResponseBody {
	s.Success = &v
	return s
}

func (s *GetAlertEventResponseBody) Validate() error {
	if s.AlertEventInfo != nil {
		if err := s.AlertEventInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertEventResponseBodyAlertEventInfo struct {
	// example:
	//
	// ONCE
	AlertFrequency    *string                                                     `json:"AlertFrequency,omitempty" xml:"AlertFrequency,omitempty"`
	AlertObject       *GetAlertEventResponseBodyAlertEventInfoAlertObject         `json:"AlertObject,omitempty" xml:"AlertObject,omitempty" type:"Struct"`
	AlertReason       *GetAlertEventResponseBodyAlertEventInfoAlertReason         `json:"AlertReason,omitempty" xml:"AlertReason,omitempty" type:"Struct"`
	AlertReceiverList []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverList `json:"AlertReceiverList,omitempty" xml:"AlertReceiverList,omitempty" type:"Repeated"`
	BelongProject     *GetAlertEventResponseBodyAlertEventInfoBelongProject       `json:"BelongProject,omitempty" xml:"BelongProject,omitempty" type:"Struct"`
	// example:
	//
	// 2024-11-05 00:00:00
	DoNotDisturbEndTime *string `json:"DoNotDisturbEndTime,omitempty" xml:"DoNotDisturbEndTime,omitempty"`
	// example:
	//
	// 2024-11-05 16:19:33
	FirstAlertTime *string `json:"FirstAlertTime,omitempty" xml:"FirstAlertTime,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 2024-11-05 16:19:33
	LatestAlertTime *string `json:"LatestAlertTime,omitempty" xml:"LatestAlertTime,omitempty"`
	// example:
	//
	// FINISH
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	TotalAlertTimes *int64                                            `json:"TotalAlertTimes,omitempty" xml:"TotalAlertTimes,omitempty"`
	UrlConfig       *GetAlertEventResponseBodyAlertEventInfoUrlConfig `json:"UrlConfig,omitempty" xml:"UrlConfig,omitempty" type:"Struct"`
}

func (s GetAlertEventResponseBodyAlertEventInfo) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfo) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetAlertFrequency() *string {
	return s.AlertFrequency
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetAlertObject() *GetAlertEventResponseBodyAlertEventInfoAlertObject {
	return s.AlertObject
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetAlertReason() *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	return s.AlertReason
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetAlertReceiverList() []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	return s.AlertReceiverList
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetBelongProject() *GetAlertEventResponseBodyAlertEventInfoBelongProject {
	return s.BelongProject
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetDoNotDisturbEndTime() *string {
	return s.DoNotDisturbEndTime
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetFirstAlertTime() *string {
	return s.FirstAlertTime
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetId() *int64 {
	return s.Id
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetLatestAlertTime() *string {
	return s.LatestAlertTime
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetStatus() *string {
	return s.Status
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetTotalAlertTimes() *int64 {
	return s.TotalAlertTimes
}

func (s *GetAlertEventResponseBodyAlertEventInfo) GetUrlConfig() *GetAlertEventResponseBodyAlertEventInfoUrlConfig {
	return s.UrlConfig
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetAlertFrequency(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.AlertFrequency = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetAlertObject(v *GetAlertEventResponseBodyAlertEventInfoAlertObject) *GetAlertEventResponseBodyAlertEventInfo {
	s.AlertObject = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetAlertReason(v *GetAlertEventResponseBodyAlertEventInfoAlertReason) *GetAlertEventResponseBodyAlertEventInfo {
	s.AlertReason = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetAlertReceiverList(v []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) *GetAlertEventResponseBodyAlertEventInfo {
	s.AlertReceiverList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetBelongProject(v *GetAlertEventResponseBodyAlertEventInfoBelongProject) *GetAlertEventResponseBodyAlertEventInfo {
	s.BelongProject = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetDoNotDisturbEndTime(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.DoNotDisturbEndTime = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetFirstAlertTime(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.FirstAlertTime = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetId(v int64) *GetAlertEventResponseBodyAlertEventInfo {
	s.Id = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetLatestAlertTime(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.LatestAlertTime = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetStatus(v string) *GetAlertEventResponseBodyAlertEventInfo {
	s.Status = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetTotalAlertTimes(v int64) *GetAlertEventResponseBodyAlertEventInfo {
	s.TotalAlertTimes = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) SetUrlConfig(v *GetAlertEventResponseBodyAlertEventInfoUrlConfig) *GetAlertEventResponseBodyAlertEventInfo {
	s.UrlConfig = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfo) Validate() error {
	if s.AlertObject != nil {
		if err := s.AlertObject.Validate(); err != nil {
			return err
		}
	}
	if s.AlertReason != nil {
		if err := s.AlertReason.Validate(); err != nil {
			return err
		}
	}
	if s.AlertReceiverList != nil {
		for _, item := range s.AlertReceiverList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BelongProject != nil {
		if err := s.BelongProject.Validate(); err != nil {
			return err
		}
	}
	if s.UrlConfig != nil {
		if err := s.UrlConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAlertEventResponseBodyAlertEventInfoAlertObject struct {
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ALL
	SourceSystemType *string `json:"SourceSystemType,omitempty" xml:"SourceSystemType,omitempty"`
	// example:
	//
	// VDM_BATCH_PYTHON37
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertObject) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertObject) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) GetName() *string {
	return s.Name
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) GetSourceSystemType() *string {
	return s.SourceSystemType
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) GetType() *string {
	return s.Type
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) SetName(v string) *GetAlertEventResponseBodyAlertEventInfoAlertObject {
	s.Name = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) SetSourceSystemType(v string) *GetAlertEventResponseBodyAlertEventInfoAlertObject {
	s.SourceSystemType = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) SetType(v string) *GetAlertEventResponseBodyAlertEventInfoAlertObject {
	s.Type = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertObject) Validate() error {
	return dara.Validate(s)
}

type GetAlertEventResponseBodyAlertEventInfoAlertReason struct {
	AlertReasonParamList []*GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList `json:"AlertReasonParamList,omitempty" xml:"AlertReasonParamList,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-11-05 16:19:32
	BizDate *string `json:"BizDate,omitempty" xml:"BizDate,omitempty"`
	// example:
	//
	// VDM_BATCH_FINISH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// t_6340131422711644160_20241104_6340142
	UniqueKey *string `json:"UniqueKey,omitempty" xml:"UniqueKey,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReason) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReason) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) GetAlertReasonParamList() []*GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList {
	return s.AlertReasonParamList
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) GetBizDate() *string {
	return s.BizDate
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) GetType() *string {
	return s.Type
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) GetUniqueKey() *string {
	return s.UniqueKey
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) SetAlertReasonParamList(v []*GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	s.AlertReasonParamList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) SetBizDate(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	s.BizDate = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) SetType(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	s.Type = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) SetUniqueKey(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReason {
	s.UniqueKey = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReason) Validate() error {
	if s.AlertReasonParamList != nil {
		for _, item := range s.AlertReasonParamList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList struct {
	// example:
	//
	// biz_date
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 2024-11-04 00:00:00
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) GetKey() *string {
	return s.Key
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) GetValue() *string {
	return s.Value
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) SetKey(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList {
	s.Key = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) SetValue(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList {
	s.Value = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReasonAlertReasonParamList) Validate() error {
	return dara.Validate(s)
}

type GetAlertEventResponseBodyAlertEventInfoAlertReceiverList struct {
	AlertChannelTypeList     []*string `json:"AlertChannelTypeList,omitempty" xml:"AlertChannelTypeList,omitempty" type:"Repeated"`
	CustomAlertChannelIdList []*string `json:"CustomAlertChannelIdList,omitempty" xml:"CustomAlertChannelIdList,omitempty" type:"Repeated"`
	// example:
	//
	// test
	OnCallTableName *string `json:"OnCallTableName,omitempty" xml:"OnCallTableName,omitempty"`
	// example:
	//
	// OWNER
	Type     *string                                                             `json:"Type,omitempty" xml:"Type,omitempty"`
	UserList []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList `json:"UserList,omitempty" xml:"UserList,omitempty" type:"Repeated"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetAlertChannelTypeList() []*string {
	return s.AlertChannelTypeList
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetCustomAlertChannelIdList() []*string {
	return s.CustomAlertChannelIdList
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetOnCallTableName() *string {
	return s.OnCallTableName
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetType() *string {
	return s.Type
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) GetUserList() []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList {
	return s.UserList
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetAlertChannelTypeList(v []*string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.AlertChannelTypeList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetCustomAlertChannelIdList(v []*string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.CustomAlertChannelIdList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetOnCallTableName(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.OnCallTableName = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetType(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.Type = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) SetUserList(v []*GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList {
	s.UserList = v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverList) Validate() error {
	if s.UserList != nil {
		for _, item := range s.UserList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList struct {
	// example:
	//
	// Admin
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) GetName() *string {
	return s.Name
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) SetName(v string) *GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList {
	s.Name = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoAlertReceiverListUserList) Validate() error {
	return dara.Validate(s)
}

type GetAlertEventResponseBodyAlertEventInfoBelongProject struct {
	// example:
	//
	// biz_1
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	// example:
	//
	// project_1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoBelongProject) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoBelongProject) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) GetBizName() *string {
	return s.BizName
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) SetBizName(v string) *GetAlertEventResponseBodyAlertEventInfoBelongProject {
	s.BizName = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) SetProjectName(v string) *GetAlertEventResponseBodyAlertEventInfoBelongProject {
	s.ProjectName = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoBelongProject) Validate() error {
	return dara.Validate(s)
}

type GetAlertEventResponseBodyAlertEventInfoUrlConfig struct {
	// example:
	//
	// https://dataphin.com/ops/test3
	AlertConfigUrl *string `json:"AlertConfigUrl,omitempty" xml:"AlertConfigUrl,omitempty"`
	// example:
	//
	// https://dataphin.com/ops/test2
	LogUrl *string `json:"LogUrl,omitempty" xml:"LogUrl,omitempty"`
	// example:
	//
	// https://dataphin.com/ops/test1
	ObjectUrl *string `json:"ObjectUrl,omitempty" xml:"ObjectUrl,omitempty"`
}

func (s GetAlertEventResponseBodyAlertEventInfoUrlConfig) String() string {
	return dara.Prettify(s)
}

func (s GetAlertEventResponseBodyAlertEventInfoUrlConfig) GoString() string {
	return s.String()
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) GetAlertConfigUrl() *string {
	return s.AlertConfigUrl
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) GetLogUrl() *string {
	return s.LogUrl
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) GetObjectUrl() *string {
	return s.ObjectUrl
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) SetAlertConfigUrl(v string) *GetAlertEventResponseBodyAlertEventInfoUrlConfig {
	s.AlertConfigUrl = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) SetLogUrl(v string) *GetAlertEventResponseBodyAlertEventInfoUrlConfig {
	s.LogUrl = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) SetObjectUrl(v string) *GetAlertEventResponseBodyAlertEventInfoUrlConfig {
	s.ObjectUrl = &v
	return s
}

func (s *GetAlertEventResponseBodyAlertEventInfoUrlConfig) Validate() error {
	return dara.Validate(s)
}
