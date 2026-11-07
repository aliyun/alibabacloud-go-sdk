// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterHealthCheckTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListClusterHealthCheckTaskResponseBody
	GetCode() *int32
	SetData(v *ListClusterHealthCheckTaskResponseBodyData) *ListClusterHealthCheckTaskResponseBody
	GetData() *ListClusterHealthCheckTaskResponseBodyData
	SetDynamicCode(v string) *ListClusterHealthCheckTaskResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListClusterHealthCheckTaskResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *ListClusterHealthCheckTaskResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListClusterHealthCheckTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListClusterHealthCheckTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListClusterHealthCheckTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClusterHealthCheckTaskResponseBody
	GetSuccess() *bool
}

type ListClusterHealthCheckTaskResponseBody struct {
	// Code is only used to synchronize with success
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListClusterHealthCheckTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic part in the error message.
	//
	// example:
	//
	// code
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message used to replace \\*\\*%s\\*\\	- in the **ErrMessage*	- parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid**, and **DynamicMessage*	- returns **DtsJobId**, it indicates that the **DtsJobId*	- request parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code. Take note of the following items:
	//
	// - Successful request: The **ErrorCode*	- field is not returned.
	//
	// - Failed request: The **ErrorCode*	- field is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// The request is processed successfully.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4E9FDCFE-0738-493B-B801-82BDFBCB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClusterHealthCheckTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHealthCheckTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterHealthCheckTaskResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListClusterHealthCheckTaskResponseBody) GetData() *ListClusterHealthCheckTaskResponseBodyData {
	return s.Data
}

func (s *ListClusterHealthCheckTaskResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListClusterHealthCheckTaskResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListClusterHealthCheckTaskResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListClusterHealthCheckTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListClusterHealthCheckTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClusterHealthCheckTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterHealthCheckTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClusterHealthCheckTaskResponseBody) SetCode(v int32) *ListClusterHealthCheckTaskResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBody) SetData(v *ListClusterHealthCheckTaskResponseBodyData) *ListClusterHealthCheckTaskResponseBody {
	s.Data = v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBody) SetDynamicCode(v string) *ListClusterHealthCheckTaskResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBody) SetDynamicMessage(v string) *ListClusterHealthCheckTaskResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBody) SetErrorCode(v string) *ListClusterHealthCheckTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBody) SetHttpStatusCode(v int32) *ListClusterHealthCheckTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBody) SetMessage(v string) *ListClusterHealthCheckTaskResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBody) SetRequestId(v string) *ListClusterHealthCheckTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBody) SetSuccess(v bool) *ListClusterHealthCheckTaskResponseBody {
	s.Success = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClusterHealthCheckTaskResponseBodyData struct {
	// The page number of the list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of health check tasks.
	Result []*ListClusterHealthCheckTaskResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of queried rules.
	//
	// example:
	//
	// 9
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListClusterHealthCheckTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHealthCheckTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClusterHealthCheckTaskResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClusterHealthCheckTaskResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterHealthCheckTaskResponseBodyData) GetResult() []*ListClusterHealthCheckTaskResponseBodyDataResult {
	return s.Result
}

func (s *ListClusterHealthCheckTaskResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListClusterHealthCheckTaskResponseBodyData) SetPageNumber(v int32) *ListClusterHealthCheckTaskResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyData) SetPageSize(v int32) *ListClusterHealthCheckTaskResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyData) SetResult(v []*ListClusterHealthCheckTaskResponseBodyDataResult) *ListClusterHealthCheckTaskResponseBodyData {
	s.Result = v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyData) SetTotalSize(v int32) *ListClusterHealthCheckTaskResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterHealthCheckTaskResponseBodyDataResult struct {
	// The complete version number.
	//
	// example:
	//
	// 1.2.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The billing method.
	//
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// Nacos-Ans
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the task was created.
	//
	// example:
	//
	// 2022-06-20T06:51:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Redundant field.
	//
	// example:
	//
	// null
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// mse_ingresspost-cn-0jbvrcex****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the user to which the cluster belongs.
	//
	// example:
	//
	// 123456
	PrimaryUser *string `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	// The number of nodes in the cluster.
	//
	// example:
	//
	// 3
	Replica *string `json:"Replica,omitempty" xml:"Replica,omitempty"`
	// The list of risk items.
	RiskList []*ListClusterHealthCheckTaskResponseBodyDataResultRiskList `json:"RiskList,omitempty" xml:"RiskList,omitempty" type:"Repeated"`
	// The total score.
	//
	// example:
	//
	// 60
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The specifications.
	//
	// example:
	//
	// MSE_SC_2_4_200_c
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The job states.
	//
	// example:
	//
	// FINISH
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of inspection items.
	//
	// example:
	//
	// 10
	TotalItem *int32 `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	// The total number of risk items.
	//
	// example:
	//
	// 3
	TotalRisk *int32 `json:"TotalRisk,omitempty" xml:"TotalRisk,omitempty"`
	// Redundant field.
	//
	// example:
	//
	// null
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the instance was last updated.
	//
	// example:
	//
	// 2022-11-12 15:07:55
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version number.
	//
	// example:
	//
	// NACOS_ANS_1_2_1_3
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s ListClusterHealthCheckTaskResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHealthCheckTaskResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetId() *int32 {
	return s.Id
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetPrimaryUser() *string {
	return s.PrimaryUser
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetReplica() *string {
	return s.Replica
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetRiskList() []*ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	return s.RiskList
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetScore() *int32 {
	return s.Score
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetSpec() *string {
	return s.Spec
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetStatus() *string {
	return s.Status
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetTotalRisk() *int32 {
	return s.TotalRisk
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetType() *string {
	return s.Type
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) GetVersionCode() *string {
	return s.VersionCode
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetAppVersion(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.AppVersion = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetChargeType(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.ChargeType = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetClusterType(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.ClusterType = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetCreateTime(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.CreateTime = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetId(v int32) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetImageVersion(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.ImageVersion = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetInstanceId(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.InstanceId = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetPrimaryUser(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.PrimaryUser = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetReplica(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.Replica = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetRiskList(v []*ListClusterHealthCheckTaskResponseBodyDataResultRiskList) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.RiskList = v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetScore(v int32) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.Score = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetSpec(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.Spec = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetStatus(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.Status = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetTotalItem(v int32) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.TotalItem = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetTotalRisk(v int32) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.TotalRisk = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetType(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.Type = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetUpdateTime(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.UpdateTime = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) SetVersionCode(v string) *ListClusterHealthCheckTaskResponseBodyDataResult {
	s.VersionCode = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResult) Validate() error {
	if s.RiskList != nil {
		for _, item := range s.RiskList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterHealthCheckTaskResponseBodyDataResultRiskList struct {
	// The description.
	//
	// example:
	//
	// {"desc":"The engine version is outdated and many features are missing. Upgrade to the latest version at the earliest opportunity."}
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The English description of the risk.
	//
	// example:
	//
	// Upgrade to the latest version to ensure that all features are available!
	DescriptionEn *string `json:"DescriptionEn,omitempty" xml:"DescriptionEn,omitempty"`
	// The ID.
	//
	// example:
	//
	// 3426
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Redundant field.
	//
	// example:
	//
	// null
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// Specifies whether to disable the notification feature if the risk item occurs.
	//
	// - true: Disables risk item notifications
	//
	// - false: Enables risk item notifications
	//
	// example:
	//
	// false
	Mute *bool `json:"Mute,omitempty" xml:"Mute,omitempty"`
	// Redundant field.
	//
	// example:
	//
	// null
	NoticeFeature *bool `json:"NoticeFeature,omitempty" xml:"NoticeFeature,omitempty"`
	// The ID of the user to which the cluster belongs.
	//
	// example:
	//
	// 123456
	PrimaryUser *string `json:"PrimaryUser,omitempty" xml:"PrimaryUser,omitempty"`
	// The risk code.
	//
	// example:
	//
	// 22020010001
	RiskCode *string `json:"RiskCode,omitempty" xml:"RiskCode,omitempty"`
	// The risk level. Valid values:
	//
	// - HIGH: High risk
	//
	// - MID: Medium risk
	//
	// - LOW: Low risk
	//
	// example:
	//
	// MID
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The risk name.
	//
	// example:
	//
	// Engine version too low
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
	// The English name of the risk.
	//
	// example:
	//
	// Low engine version
	RiskNameEn *string `json:"RiskNameEn,omitempty" xml:"RiskNameEn,omitempty"`
	// The type of the risk.
	//
	// example:
	//
	// Version risk
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
	// The current situation.
	//
	// example:
	//
	// {"desc":"The current engine version is too low with many missing features","links":[{"type":"url","value":"https://xxxx"","desc":"Version notes"}]}
	Situation *string `json:"Situation,omitempty" xml:"Situation,omitempty"`
	// The English description of the current situation.
	//
	// example:
	//
	// The current engine version is low, and there are many missing features
	SituationEn *string `json:"SituationEn,omitempty" xml:"SituationEn,omitempty"`
	// The suggestion.
	//
	// example:
	//
	// {"desc": "Upgrade to the latest version at the earliest opportunity.", "links":[{"type":"upgrade","desc":"Click to upgrade"}]}
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	// The English description of the suggestion.
	//
	// example:
	//
	// Upgrade to the latest version
	SuggestionEn *string `json:"SuggestionEn,omitempty" xml:"SuggestionEn,omitempty"`
	// The ID of the associated parent task.
	//
	// example:
	//
	// 1
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Redundant field.
	//
	// example:
	//
	// null
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// Redundant field.
	//
	// example:
	//
	// null
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s ListClusterHealthCheckTaskResponseBodyDataResultRiskList) String() string {
	return dara.Prettify(s)
}

func (s ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GoString() string {
	return s.String()
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetDescription() *string {
	return s.Description
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetDescriptionEn() *string {
	return s.DescriptionEn
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetId() *int32 {
	return s.Id
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetModule() *string {
	return s.Module
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetMute() *bool {
	return s.Mute
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetNoticeFeature() *bool {
	return s.NoticeFeature
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetPrimaryUser() *string {
	return s.PrimaryUser
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetRiskCode() *string {
	return s.RiskCode
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetRiskName() *string {
	return s.RiskName
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetRiskNameEn() *string {
	return s.RiskNameEn
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetRiskType() *string {
	return s.RiskType
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetSituation() *string {
	return s.Situation
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetSituationEn() *string {
	return s.SituationEn
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetSuggestion() *string {
	return s.Suggestion
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetSuggestionEn() *string {
	return s.SuggestionEn
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetType() *int32 {
	return s.Type
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) GetValues() *string {
	return s.Values
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetDescription(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.Description = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetDescriptionEn(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.DescriptionEn = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetId(v int32) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.Id = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetModule(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.Module = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetMute(v bool) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.Mute = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetNoticeFeature(v bool) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.NoticeFeature = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetPrimaryUser(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.PrimaryUser = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetRiskCode(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.RiskCode = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetRiskLevel(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.RiskLevel = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetRiskName(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.RiskName = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetRiskNameEn(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.RiskNameEn = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetRiskType(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.RiskType = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetSituation(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.Situation = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetSituationEn(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.SituationEn = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetSuggestion(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.Suggestion = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetSuggestionEn(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.SuggestionEn = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetTaskId(v int64) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.TaskId = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetType(v int32) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.Type = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) SetValues(v string) *ListClusterHealthCheckTaskResponseBodyDataResultRiskList {
	s.Values = &v
	return s
}

func (s *ListClusterHealthCheckTaskResponseBodyDataResultRiskList) Validate() error {
	return dara.Validate(s)
}
