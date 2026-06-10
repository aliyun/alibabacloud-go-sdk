// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppPublishHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListAppPublishHistoryResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *ListAppPublishHistoryResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *ListAppPublishHistoryResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *ListAppPublishHistoryResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *ListAppPublishHistoryResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *ListAppPublishHistoryResponseBody
	GetErrorArgs() []interface{}
	SetMaxResults(v int32) *ListAppPublishHistoryResponseBody
	GetMaxResults() *int32
	SetModule(v *ListAppPublishHistoryResponseBodyModule) *ListAppPublishHistoryResponseBody
	GetModule() *ListAppPublishHistoryResponseBodyModule
	SetNextToken(v string) *ListAppPublishHistoryResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAppPublishHistoryResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *ListAppPublishHistoryResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *ListAppPublishHistoryResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *ListAppPublishHistoryResponseBody
	GetSynchro() *bool
}

type ListAppPublishHistoryResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// is retry allowed
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// frontend application Name.
	//
	// example:
	//
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// dynamic message, not currently used. Please ignore.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// fault parameters
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// Number of results per query.
	//
	// Value range: 10 to 100. Default Value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Data Table module.
	//
	// - ABTest: experiment Data Table
	//
	// - ExperimentTool: experiment tool table
	//
	// - DataDiagnosis: Data Diagnosis
	Module *ListAppPublishHistoryResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Token for the start of the next query. It is empty if there is no next query.
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
	// error code
	//
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	// abnormal message
	//
	// example:
	//
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// Reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s ListAppPublishHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppPublishHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppPublishHistoryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListAppPublishHistoryResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *ListAppPublishHistoryResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *ListAppPublishHistoryResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *ListAppPublishHistoryResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *ListAppPublishHistoryResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *ListAppPublishHistoryResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppPublishHistoryResponseBody) GetModule() *ListAppPublishHistoryResponseBodyModule {
	return s.Module
}

func (s *ListAppPublishHistoryResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppPublishHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppPublishHistoryResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *ListAppPublishHistoryResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *ListAppPublishHistoryResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *ListAppPublishHistoryResponseBody) SetAccessDeniedDetail(v string) *ListAppPublishHistoryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetAllowRetry(v bool) *ListAppPublishHistoryResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetAppName(v string) *ListAppPublishHistoryResponseBody {
	s.AppName = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetDynamicCode(v string) *ListAppPublishHistoryResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetDynamicMessage(v string) *ListAppPublishHistoryResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetErrorArgs(v []interface{}) *ListAppPublishHistoryResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetMaxResults(v int32) *ListAppPublishHistoryResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetModule(v *ListAppPublishHistoryResponseBodyModule) *ListAppPublishHistoryResponseBody {
	s.Module = v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetNextToken(v string) *ListAppPublishHistoryResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetRequestId(v string) *ListAppPublishHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetRootErrorCode(v string) *ListAppPublishHistoryResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetRootErrorMsg(v string) *ListAppPublishHistoryResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) SetSynchro(v bool) *ListAppPublishHistoryResponseBody {
	s.Synchro = &v
	return s
}

func (s *ListAppPublishHistoryResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAppPublishHistoryResponseBodyModule struct {
	// Current publish order ID
	//
	// example:
	//
	// 123
	CurrentPublishOrderId *int64 `json:"CurrentPublishOrderId,omitempty" xml:"CurrentPublishOrderId,omitempty"`
	// Indicates whether to display the history of applying the alert template to application groups. Valid values:
	//
	// - true: Display.
	//
	// - false (default): Do not display.
	History []*ListAppPublishHistoryResponseBodyModuleHistory `json:"History,omitempty" xml:"History,omitempty" type:"Repeated"`
	// Page number. Default value is 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Paging size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total count.
	//
	// example:
	//
	// 0
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListAppPublishHistoryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s ListAppPublishHistoryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *ListAppPublishHistoryResponseBodyModule) GetCurrentPublishOrderId() *int64 {
	return s.CurrentPublishOrderId
}

func (s *ListAppPublishHistoryResponseBodyModule) GetHistory() []*ListAppPublishHistoryResponseBodyModuleHistory {
	return s.History
}

func (s *ListAppPublishHistoryResponseBodyModule) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAppPublishHistoryResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppPublishHistoryResponseBodyModule) GetTotal() *int32 {
	return s.Total
}

func (s *ListAppPublishHistoryResponseBodyModule) SetCurrentPublishOrderId(v int64) *ListAppPublishHistoryResponseBodyModule {
	s.CurrentPublishOrderId = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModule) SetHistory(v []*ListAppPublishHistoryResponseBodyModuleHistory) *ListAppPublishHistoryResponseBodyModule {
	s.History = v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModule) SetPageNum(v int32) *ListAppPublishHistoryResponseBodyModule {
	s.PageNum = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModule) SetPageSize(v int32) *ListAppPublishHistoryResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModule) SetTotal(v int32) *ListAppPublishHistoryResponseBodyModule {
	s.Total = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModule) Validate() error {
	if s.History != nil {
		for _, item := range s.History {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppPublishHistoryResponseBodyModuleHistory struct {
	// Indicates whether quick rollback is supported.
	//
	// example:
	//
	// true
	CanQuickRevert *string `json:"CanQuickRevert,omitempty" xml:"CanQuickRevert,omitempty"`
	CommitHash     *string `json:"CommitHash,omitempty" xml:"CommitHash,omitempty"`
	// Current step.
	//
	// example:
	//
	// PRE_CHECK
	CurrentStep *string `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	// Deployment channel.
	//
	// example:
	//
	// PC
	DeployChannel *string `json:"DeployChannel,omitempty" xml:"DeployChannel,omitempty"`
	// Application description.
	//
	// example:
	//
	// PSKB
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Publishing procedure.
	//
	// example:
	//
	// error
	ErrorStep *string `json:"ErrorStep,omitempty" xml:"ErrorStep,omitempty"`
	// Indicates whether the process is complete.
	IsFinish *bool `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	// Indicates whether resource allocation to the cost center succeeded.
	//
	// - true indicates success.
	//
	// - false indicates failure.
	//
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// Error message.
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Sorting type: ASC or DESC.
	//
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// Transcoding progress.
	//
	// example:
	//
	// true
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// Publish number.
	//
	// example:
	//
	// 123
	PublishNumber *string `json:"PublishNumber,omitempty" xml:"PublishNumber,omitempty"`
	// Publish order ID.
	//
	// example:
	//
	// 123
	PublishOrderId *int64 `json:"PublishOrderId,omitempty" xml:"PublishOrderId,omitempty"`
	// Published At.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// Specific widget configuration.
	Steps      []*string `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	Subchannel *string   `json:"Subchannel,omitempty" xml:"Subchannel,omitempty"`
}

func (s ListAppPublishHistoryResponseBodyModuleHistory) String() string {
	return dara.Prettify(s)
}

func (s ListAppPublishHistoryResponseBodyModuleHistory) GoString() string {
	return s.String()
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetCanQuickRevert() *string {
	return s.CanQuickRevert
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetCommitHash() *string {
	return s.CommitHash
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetDeployChannel() *string {
	return s.DeployChannel
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetDescription() *string {
	return s.Description
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetErrorStep() *string {
	return s.ErrorStep
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetMsg() *string {
	return s.Msg
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetOrderType() *string {
	return s.OrderType
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPercent() *int32 {
	return s.Percent
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPublishNumber() *string {
	return s.PublishNumber
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPublishOrderId() *int64 {
	return s.PublishOrderId
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPublishTime() *string {
	return s.PublishTime
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetSteps() []*string {
	return s.Steps
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetSubchannel() *string {
	return s.Subchannel
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetCanQuickRevert(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.CanQuickRevert = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetCommitHash(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.CommitHash = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetCurrentStep(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.CurrentStep = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetDeployChannel(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.DeployChannel = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetDescription(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.Description = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetErrorStep(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.ErrorStep = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetIsFinish(v bool) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.IsFinish = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetIsSuccess(v bool) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.IsSuccess = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetMsg(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.Msg = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetOrderType(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.OrderType = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPercent(v int32) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.Percent = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPublishNumber(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PublishNumber = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPublishOrderId(v int64) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PublishOrderId = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPublishTime(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PublishTime = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetSteps(v []*string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.Steps = v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetSubchannel(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.Subchannel = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) Validate() error {
	return dara.Validate(s)
}
