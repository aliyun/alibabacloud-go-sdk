// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPromotionActivitiesForPartnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListPromotionActivitiesForPartnerResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListPromotionActivitiesForPartnerResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListPromotionActivitiesForPartnerResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListPromotionActivitiesForPartnerResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListPromotionActivitiesForPartnerResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListPromotionActivitiesForPartnerResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListPromotionActivitiesForPartnerResponseBody
	GetMaxResults() *int32
	SetModule(v *ListPromotionActivitiesForPartnerResponseBodyModule) *ListPromotionActivitiesForPartnerResponseBody
	GetModule() *ListPromotionActivitiesForPartnerResponseBodyModule
	SetNextToken(v string) *ListPromotionActivitiesForPartnerResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListPromotionActivitiesForPartnerResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListPromotionActivitiesForPartnerResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListPromotionActivitiesForPartnerResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListPromotionActivitiesForPartnerResponseBody
	GetSynchro() *bool
}

type ListPromotionActivitiesForPartnerResponseBody struct {
	// The access denied details.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Indicates whether retry is allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The application name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic message.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error arguments.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The number of entries per query.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The response data.
	Module *ListPromotionActivitiesForPartnerResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The token for the next query. This parameter is empty if no more results exist.
	//
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error code.
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// The root error message.
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListPromotionActivitiesForPartnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionActivitiesForPartnerResponseBody) GoString() string {
	return s.String()
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetModule() *ListPromotionActivitiesForPartnerResponseBodyModule {
	return s.Module
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListPromotionActivitiesForPartnerResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetAccessDeniedDetail(v string) *ListPromotionActivitiesForPartnerResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetAllowRetry(v bool) *ListPromotionActivitiesForPartnerResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetAppName(v string) *ListPromotionActivitiesForPartnerResponseBody {
	s.AppName = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetDynamicCode(v string) *ListPromotionActivitiesForPartnerResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetDynamicMessage(v string) *ListPromotionActivitiesForPartnerResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetErrorArgs(v []interface{}) *ListPromotionActivitiesForPartnerResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetMaxResults(v int32) *ListPromotionActivitiesForPartnerResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetModule(v *ListPromotionActivitiesForPartnerResponseBodyModule) *ListPromotionActivitiesForPartnerResponseBody {
	s.Module = v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetNextToken(v string) *ListPromotionActivitiesForPartnerResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetRequestId(v string) *ListPromotionActivitiesForPartnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetRootErrorCode(v string) *ListPromotionActivitiesForPartnerResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetRootErrorMsg(v string) *ListPromotionActivitiesForPartnerResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) SetSynchro(v bool) *ListPromotionActivitiesForPartnerResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPromotionActivitiesForPartnerResponseBodyModule struct {
	// The list of promotional activities.
	Activities []*ListPromotionActivitiesForPartnerResponseBodyModuleActivities `json:"Activities,omitempty" xml:"Activities,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPromotionActivitiesForPartnerResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionActivitiesForPartnerResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModule) GetActivities() []*ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	return s.Activities
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModule) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModule) SetActivities(v []*ListPromotionActivitiesForPartnerResponseBodyModuleActivities) *ListPromotionActivitiesForPartnerResponseBodyModule {
	s.Activities = v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModule) SetTotalCount(v int32) *ListPromotionActivitiesForPartnerResponseBodyModule {
	s.TotalCount = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModule) Validate() error {
	if s.Activities != nil {
		for _, item := range s.Activities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPromotionActivitiesForPartnerResponseBodyModuleActivities struct {
	// The activity code.
	//
	// example:
	//
	// acwfradoj5u
	ActivityCode *string `json:"ActivityCode,omitempty" xml:"ActivityCode,omitempty"`
	// The activity name.
	//
	// example:
	//
	// IP网段过滤统计
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The activity type.
	//
	// example:
	//
	// SCALE_IN
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	// The consumed quota.
	//
	// example:
	//
	// 100
	ConsumedQuota *int64 `json:"ConsumedQuota,omitempty" xml:"ConsumedQuota,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-02-14 11:57:51
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user who created the activity.
	//
	// example:
	//
	// onlinehoztestrolebasicallow1755461756261068111(300216315676902267)
	CreatedBy *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	// The eligibility configuration in JSON format.
	//
	// example:
	//
	// {}
	EligibilityConfig *string `json:"EligibilityConfig,omitempty" xml:"EligibilityConfig,omitempty"`
	// The end date.
	//
	// example:
	//
	// 2026-04-22 10:18:51 +0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The offer configuration in JSON format.
	//
	// example:
	//
	// {}
	OfferConfig *string `json:"OfferConfig,omitempty" xml:"OfferConfig,omitempty"`
	// The offer configuration summary.
	//
	// example:
	//
	// {}
	OfferConfigSummary *string `json:"OfferConfigSummary,omitempty" xml:"OfferConfigSummary,omitempty"`
	// The remaining quota.
	//
	// example:
	//
	// 10
	RemainingQuota *int64 `json:"RemainingQuota,omitempty" xml:"RemainingQuota,omitempty"`
	// The start date.
	//
	// example:
	//
	// 2026-05-12T16:00:00.000Z
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The activity status.
	//
	// example:
	//
	// FE_ABORTING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total quota.
	//
	// example:
	//
	// 10
	TotalQuota *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// The touchpoint configuration in JSON format.
	//
	// example:
	//
	// {}
	TouchpointConfig *string `json:"TouchpointConfig,omitempty" xml:"TouchpointConfig,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2025-10-11T21:13:38.164526965+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user who last updated the activity.
	//
	// example:
	//
	// 123414
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
	// The warning threshold.
	//
	// example:
	//
	// 0
	WarningThreshold *int32 `json:"WarningThreshold,omitempty" xml:"WarningThreshold,omitempty"`
}

func (s ListPromotionActivitiesForPartnerResponseBodyModuleActivities) String() string {
	return dara.Prettify(s)
}

func (s ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GoString() string {
	return s.String()
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetActivityCode() *string {
	return s.ActivityCode
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetActivityName() *string {
	return s.ActivityName
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetActivityType() *string {
	return s.ActivityType
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetConsumedQuota() *int64 {
	return s.ConsumedQuota
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetCreatedBy() *string {
	return s.CreatedBy
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetEligibilityConfig() *string {
	return s.EligibilityConfig
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetEndDate() *string {
	return s.EndDate
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetOfferConfig() *string {
	return s.OfferConfig
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetOfferConfigSummary() *string {
	return s.OfferConfigSummary
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetRemainingQuota() *int64 {
	return s.RemainingQuota
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetStartDate() *string {
	return s.StartDate
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetStatus() *string {
	return s.Status
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetTouchpointConfig() *string {
	return s.TouchpointConfig
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) GetWarningThreshold() *int32 {
	return s.WarningThreshold
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetActivityCode(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.ActivityCode = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetActivityName(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.ActivityName = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetActivityType(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.ActivityType = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetConsumedQuota(v int64) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.ConsumedQuota = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetCreateTime(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.CreateTime = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetCreatedBy(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.CreatedBy = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetEligibilityConfig(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.EligibilityConfig = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetEndDate(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.EndDate = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetOfferConfig(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.OfferConfig = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetOfferConfigSummary(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.OfferConfigSummary = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetRemainingQuota(v int64) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.RemainingQuota = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetStartDate(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.StartDate = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetStatus(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.Status = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetTotalQuota(v int64) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.TotalQuota = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetTouchpointConfig(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.TouchpointConfig = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetUpdateTime(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.UpdateTime = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetUpdatedBy(v string) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.UpdatedBy = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) SetWarningThreshold(v int32) *ListPromotionActivitiesForPartnerResponseBodyModuleActivities {
	s.WarningThreshold = &v
	return s
}

func (s *ListPromotionActivitiesForPartnerResponseBodyModuleActivities) Validate() error {
	return dara.Validate(s)
}
