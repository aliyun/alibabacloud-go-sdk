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
	// The details of the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Specifies whether retries are allowed.
	//
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// The name of the frontend application.
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
	// The dynamic message. This parameter is not used and can be ignored.
	//
	// example:
	//
	// abc
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The maximum number of entries to return for a single query. Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The data table module. Valid values:
	//
	// - ABTest: experiment data table
	//
	// - ExperimentTool: experiment tool table
	//
	// - DataDiagnosis: data diagnostics
	Module *ListAppPublishHistoryResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The token that is used to start the next query. If no next query exists, this parameter is empty.
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
	// The exception message.
	//
	// example:
	//
	// System exception
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
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
	// The ID of the current release order.
	//
	// example:
	//
	// 123
	CurrentPublishOrderId *int64 `json:"CurrentPublishOrderId,omitempty" xml:"CurrentPublishOrderId,omitempty"`
	// Specifies whether to display the history of applying alert templates to application groups. Valid values:
	//
	// - true: Displayed.
	//
	// - false: Not displayed. Default value: false.
	History []*ListAppPublishHistoryResponseBodyModuleHistory `json:"History,omitempty" xml:"History,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
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
	ArtifactDigest *string `json:"ArtifactDigest,omitempty" xml:"ArtifactDigest,omitempty"`
	AuditId        *string `json:"AuditId,omitempty" xml:"AuditId,omitempty"`
	BindingId      *string `json:"BindingId,omitempty" xml:"BindingId,omitempty"`
	// Specifies whether quick rollback is supported.
	//
	// example:
	//
	// true
	CanQuickRevert  *string `json:"CanQuickRevert,omitempty" xml:"CanQuickRevert,omitempty"`
	CandidateId     *string `json:"CandidateId,omitempty" xml:"CandidateId,omitempty"`
	CandidateStatus *string `json:"CandidateStatus,omitempty" xml:"CandidateStatus,omitempty"`
	// The commit hash of the code.
	//
	// example:
	//
	// 890243d5977e74d80eebd4734f7b05f626f84cb1
	CommitHash *string `json:"CommitHash,omitempty" xml:"CommitHash,omitempty"`
	// The current step number.
	//
	// example:
	//
	// PRE_CHECK
	CurrentStep *string `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	// The code of the current release step.
	//
	// example:
	//
	// INIT
	CurrentStepCode *string `json:"CurrentStepCode,omitempty" xml:"CurrentStepCode,omitempty"`
	// The deployment channel.
	//
	// example:
	//
	// PC
	DeployChannel *string `json:"DeployChannel,omitempty" xml:"DeployChannel,omitempty"`
	// The application description.
	//
	// example:
	//
	// PSKB
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The release process.
	//
	// example:
	//
	// error
	ErrorStep *string `json:"ErrorStep,omitempty" xml:"ErrorStep,omitempty"`
	// The code of the release step where the error occurred.
	//
	// example:
	//
	// BUILDING
	ErrorStepCode *string `json:"ErrorStepCode,omitempty" xml:"ErrorStepCode,omitempty"`
	// Indicates whether the operation is completed.
	IsFinish *bool `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	// Indicates whether the resource allocation for the financial unit is successful. Valid values:
	//
	// - true: The allocation is successful.
	//
	// - false: The allocation fails.
	//
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The error message.
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The sort order. Valid values: ASC and DESC.
	//
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The code of the release order type.
	//
	// example:
	//
	// DEPLOY
	OrderTypeCode *string `json:"OrderTypeCode,omitempty" xml:"OrderTypeCode,omitempty"`
	// The transcoding progress.
	//
	// example:
	//
	// true
	Percent               *int32  `json:"Percent,omitempty" xml:"Percent,omitempty"`
	PlatformAppid         *string `json:"PlatformAppid,omitempty" xml:"PlatformAppid,omitempty"`
	PluginDeclarationHash *string `json:"PluginDeclarationHash,omitempty" xml:"PluginDeclarationHash,omitempty"`
	PluginState           *string `json:"PluginState,omitempty" xml:"PluginState,omitempty"`
	PrivacyDraftId        *string `json:"PrivacyDraftId,omitempty" xml:"PrivacyDraftId,omitempty"`
	PrivacyRevision       *string `json:"PrivacyRevision,omitempty" xml:"PrivacyRevision,omitempty"`
	PublishDiagnostics    *string `json:"PublishDiagnostics,omitempty" xml:"PublishDiagnostics,omitempty"`
	PublishEnv            *string `json:"PublishEnv,omitempty" xml:"PublishEnv,omitempty"`
	// The release number.
	//
	// example:
	//
	// 123
	PublishNumber *string `json:"PublishNumber,omitempty" xml:"PublishNumber,omitempty"`
	// The release order ID.
	//
	// example:
	//
	// 123
	PublishOrderId *int64 `json:"PublishOrderId,omitempty" xml:"PublishOrderId,omitempty"`
	// The release time in the UTC format yyyy-MM-dd\"T\"HH:mm:ss.SSS\"Z\".
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 2026
	PublishTime    *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	ReadbackStatus *string `json:"ReadbackStatus,omitempty" xml:"ReadbackStatus,omitempty"`
	// The list of release step codes.
	StepCodes []*string `json:"StepCodes,omitempty" xml:"StepCodes,omitempty" type:"Repeated"`
	// The specific configurations of the component.
	Steps []*string `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	// The release sub-channel.
	//
	// example:
	//
	// wx0123456789abcdef
	Subchannel         *string `json:"Subchannel,omitempty" xml:"Subchannel,omitempty"`
	SubmitIntentStatus *string `json:"SubmitIntentStatus,omitempty" xml:"SubmitIntentStatus,omitempty"`
	UpdatedAt          *int64  `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	UploadAttemptId    *string `json:"UploadAttemptId,omitempty" xml:"UploadAttemptId,omitempty"`
	UploadLeaseState   *string `json:"UploadLeaseState,omitempty" xml:"UploadLeaseState,omitempty"`
	UploadSequence     *int64  `json:"UploadSequence,omitempty" xml:"UploadSequence,omitempty"`
	UploadTicketState  *string `json:"UploadTicketState,omitempty" xml:"UploadTicketState,omitempty"`
}

func (s ListAppPublishHistoryResponseBodyModuleHistory) String() string {
	return dara.Prettify(s)
}

func (s ListAppPublishHistoryResponseBodyModuleHistory) GoString() string {
	return s.String()
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetArtifactDigest() *string {
	return s.ArtifactDigest
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetAuditId() *string {
	return s.AuditId
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetBindingId() *string {
	return s.BindingId
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetCanQuickRevert() *string {
	return s.CanQuickRevert
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetCandidateId() *string {
	return s.CandidateId
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetCandidateStatus() *string {
	return s.CandidateStatus
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetCommitHash() *string {
	return s.CommitHash
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetCurrentStepCode() *string {
	return s.CurrentStepCode
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

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetErrorStepCode() *string {
	return s.ErrorStepCode
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

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetOrderTypeCode() *string {
	return s.OrderTypeCode
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPercent() *int32 {
	return s.Percent
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPlatformAppid() *string {
	return s.PlatformAppid
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPluginDeclarationHash() *string {
	return s.PluginDeclarationHash
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPluginState() *string {
	return s.PluginState
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPrivacyDraftId() *string {
	return s.PrivacyDraftId
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPrivacyRevision() *string {
	return s.PrivacyRevision
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPublishDiagnostics() *string {
	return s.PublishDiagnostics
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetPublishEnv() *string {
	return s.PublishEnv
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

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetReadbackStatus() *string {
	return s.ReadbackStatus
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetStepCodes() []*string {
	return s.StepCodes
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetSteps() []*string {
	return s.Steps
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetSubchannel() *string {
	return s.Subchannel
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetSubmitIntentStatus() *string {
	return s.SubmitIntentStatus
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetUploadAttemptId() *string {
	return s.UploadAttemptId
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetUploadLeaseState() *string {
	return s.UploadLeaseState
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetUploadSequence() *int64 {
	return s.UploadSequence
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) GetUploadTicketState() *string {
	return s.UploadTicketState
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetArtifactDigest(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.ArtifactDigest = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetAuditId(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.AuditId = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetBindingId(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.BindingId = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetCanQuickRevert(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.CanQuickRevert = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetCandidateId(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.CandidateId = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetCandidateStatus(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.CandidateStatus = &v
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

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetCurrentStepCode(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.CurrentStepCode = &v
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

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetErrorStepCode(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.ErrorStepCode = &v
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

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetOrderTypeCode(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.OrderTypeCode = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPercent(v int32) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.Percent = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPlatformAppid(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PlatformAppid = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPluginDeclarationHash(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PluginDeclarationHash = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPluginState(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PluginState = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPrivacyDraftId(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PrivacyDraftId = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPrivacyRevision(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PrivacyRevision = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPublishDiagnostics(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PublishDiagnostics = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetPublishEnv(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.PublishEnv = &v
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

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetReadbackStatus(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.ReadbackStatus = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetStepCodes(v []*string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.StepCodes = v
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

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetSubmitIntentStatus(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.SubmitIntentStatus = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetUpdatedAt(v int64) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.UpdatedAt = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetUploadAttemptId(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.UploadAttemptId = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetUploadLeaseState(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.UploadLeaseState = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetUploadSequence(v int64) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.UploadSequence = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) SetUploadTicketState(v string) *ListAppPublishHistoryResponseBodyModuleHistory {
	s.UploadTicketState = &v
	return s
}

func (s *ListAppPublishHistoryResponseBodyModuleHistory) Validate() error {
	return dara.Validate(s)
}
