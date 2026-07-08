// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppPublishStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppPublishStatusResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppPublishStatusResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppPublishStatusResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppPublishStatusResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppPublishStatusResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppPublishStatusResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppPublishStatusResponseBodyModule) *GetAppPublishStatusResponseBody
	GetModule() *GetAppPublishStatusResponseBodyModule
	SetRequestId(v string) *GetAppPublishStatusResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppPublishStatusResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppPublishStatusResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppPublishStatusResponseBody
	GetSynchro() *bool
}

type GetAppPublishStatusResponseBody struct {
	// The detailed reason why access is denied.
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
	// or
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
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The response data.
	Module *GetAppPublishStatusResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// 系统异常
	RootErrorMsg *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppPublishStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppPublishStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppPublishStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppPublishStatusResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppPublishStatusResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppPublishStatusResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppPublishStatusResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppPublishStatusResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppPublishStatusResponseBody) GetModule() *GetAppPublishStatusResponseBodyModule {
	return s.Module
}

func (s *GetAppPublishStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppPublishStatusResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppPublishStatusResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppPublishStatusResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppPublishStatusResponseBody) SetAccessDeniedDetail(v string) *GetAppPublishStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetAllowRetry(v bool) *GetAppPublishStatusResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetAppName(v string) *GetAppPublishStatusResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetDynamicCode(v string) *GetAppPublishStatusResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetDynamicMessage(v string) *GetAppPublishStatusResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetErrorArgs(v []interface{}) *GetAppPublishStatusResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetModule(v *GetAppPublishStatusResponseBodyModule) *GetAppPublishStatusResponseBody {
	s.Module = v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetRequestId(v string) *GetAppPublishStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetRootErrorCode(v string) *GetAppPublishStatusResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetRootErrorMsg(v string) *GetAppPublishStatusResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppPublishStatusResponseBody) SetSynchro(v bool) *GetAppPublishStatusResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppPublishStatusResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppPublishStatusResponseBodyModule struct {
	// Indicates whether quick rollback is supported.
	//
	// example:
	//
	// true
	CanQuickRevert *string `json:"CanQuickRevert,omitempty" xml:"CanQuickRevert,omitempty"`
	// The current step of the task.
	//
	// example:
	//
	// PRE_CHECK
	CurrentStep *string `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	// The deployment channel.
	//
	// example:
	//
	// PC,WEAPP
	DeployChannel *string `json:"DeployChannel,omitempty" xml:"DeployChannel,omitempty"`
	// The application description.
	//
	// example:
	//
	// /bak->serverless.handler(2020091300200279)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The publish process.
	//
	// example:
	//
	// abc
	ErrorStep *string `json:"ErrorStep,omitempty" xml:"ErrorStep,omitempty"`
	// Indicates whether the task is complete.
	IsFinish *bool `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - `true`: The call is successful.
	//
	// - `false`: The call failed.
	//
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The additional description.
	//
	// example:
	//
	// common notify successfully.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The sort type. Valid values: ASC and DESC.
	//
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The task completion percentage.
	//
	// example:
	//
	// 20
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// The publish number.
	//
	// example:
	//
	// 123
	PublishNumber *string `json:"PublishNumber,omitempty" xml:"PublishNumber,omitempty"`
	// The publish order ID.
	//
	// example:
	//
	// 123
	PublishOrderId *int64 `json:"PublishOrderId,omitempty" xml:"PublishOrderId,omitempty"`
	// The scheduled publish time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 123123
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// The site ID. You can obtain the site ID by calling the [ListSites](~~ListSites~~) operation.
	//
	// example:
	//
	// 865181640657408
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The error level. Valid values: FATAL, ERROR, WARNING, and CRITICAL.
	Steps      []*string `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	Subchannel *string   `json:"Subchannel,omitempty" xml:"Subchannel,omitempty"`
}

func (s GetAppPublishStatusResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppPublishStatusResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppPublishStatusResponseBodyModule) GetCanQuickRevert() *string {
	return s.CanQuickRevert
}

func (s *GetAppPublishStatusResponseBodyModule) GetCurrentStep() *string {
	return s.CurrentStep
}

func (s *GetAppPublishStatusResponseBodyModule) GetDeployChannel() *string {
	return s.DeployChannel
}

func (s *GetAppPublishStatusResponseBodyModule) GetDescription() *string {
	return s.Description
}

func (s *GetAppPublishStatusResponseBodyModule) GetErrorStep() *string {
	return s.ErrorStep
}

func (s *GetAppPublishStatusResponseBodyModule) GetIsFinish() *bool {
	return s.IsFinish
}

func (s *GetAppPublishStatusResponseBodyModule) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetAppPublishStatusResponseBodyModule) GetMsg() *string {
	return s.Msg
}

func (s *GetAppPublishStatusResponseBodyModule) GetOrderType() *string {
	return s.OrderType
}

func (s *GetAppPublishStatusResponseBodyModule) GetPercent() *int32 {
	return s.Percent
}

func (s *GetAppPublishStatusResponseBodyModule) GetPublishNumber() *string {
	return s.PublishNumber
}

func (s *GetAppPublishStatusResponseBodyModule) GetPublishOrderId() *int64 {
	return s.PublishOrderId
}

func (s *GetAppPublishStatusResponseBodyModule) GetPublishTime() *string {
	return s.PublishTime
}

func (s *GetAppPublishStatusResponseBodyModule) GetSiteId() *string {
	return s.SiteId
}

func (s *GetAppPublishStatusResponseBodyModule) GetSteps() []*string {
	return s.Steps
}

func (s *GetAppPublishStatusResponseBodyModule) GetSubchannel() *string {
	return s.Subchannel
}

func (s *GetAppPublishStatusResponseBodyModule) SetCanQuickRevert(v string) *GetAppPublishStatusResponseBodyModule {
	s.CanQuickRevert = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetCurrentStep(v string) *GetAppPublishStatusResponseBodyModule {
	s.CurrentStep = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetDeployChannel(v string) *GetAppPublishStatusResponseBodyModule {
	s.DeployChannel = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetDescription(v string) *GetAppPublishStatusResponseBodyModule {
	s.Description = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetErrorStep(v string) *GetAppPublishStatusResponseBodyModule {
	s.ErrorStep = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetIsFinish(v bool) *GetAppPublishStatusResponseBodyModule {
	s.IsFinish = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetIsSuccess(v bool) *GetAppPublishStatusResponseBodyModule {
	s.IsSuccess = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetMsg(v string) *GetAppPublishStatusResponseBodyModule {
	s.Msg = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetOrderType(v string) *GetAppPublishStatusResponseBodyModule {
	s.OrderType = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetPercent(v int32) *GetAppPublishStatusResponseBodyModule {
	s.Percent = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetPublishNumber(v string) *GetAppPublishStatusResponseBodyModule {
	s.PublishNumber = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetPublishOrderId(v int64) *GetAppPublishStatusResponseBodyModule {
	s.PublishOrderId = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetPublishTime(v string) *GetAppPublishStatusResponseBodyModule {
	s.PublishTime = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetSiteId(v string) *GetAppPublishStatusResponseBodyModule {
	s.SiteId = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetSteps(v []*string) *GetAppPublishStatusResponseBodyModule {
	s.Steps = v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) SetSubchannel(v string) *GetAppPublishStatusResponseBodyModule {
	s.Subchannel = &v
	return s
}

func (s *GetAppPublishStatusResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
