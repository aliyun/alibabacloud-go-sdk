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
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// False
	AllowRetry *bool `json:"AllowRetry,omitempty" xml:"AllowRetry,omitempty"`
	// example:
	//
	// or
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                          `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *GetAppPublishStatusResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6C6B99AC-39EC-5350-874C-204128C905E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// SYSTEM.ERROR
	RootErrorCode *string `json:"RootErrorCode,omitempty" xml:"RootErrorCode,omitempty"`
	RootErrorMsg  *string `json:"RootErrorMsg,omitempty" xml:"RootErrorMsg,omitempty"`
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
	// example:
	//
	// true
	CanQuickRevert *string `json:"CanQuickRevert,omitempty" xml:"CanQuickRevert,omitempty"`
	// example:
	//
	// PRE_CHECK
	CurrentStep *string `json:"CurrentStep,omitempty" xml:"CurrentStep,omitempty"`
	// example:
	//
	// PC,WEAPP
	DeployChannel *string `json:"DeployChannel,omitempty" xml:"DeployChannel,omitempty"`
	// example:
	//
	// /bak->serverless.handler(2020091300200279)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// abc
	ErrorStep *string `json:"ErrorStep,omitempty" xml:"ErrorStep,omitempty"`
	IsFinish  *bool   `json:"IsFinish,omitempty" xml:"IsFinish,omitempty"`
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// common notify successfully.
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 20
	Percent *int32 `json:"Percent,omitempty" xml:"Percent,omitempty"`
	// example:
	//
	// 123
	PublishNumber *string `json:"PublishNumber,omitempty" xml:"PublishNumber,omitempty"`
	// example:
	//
	// 123
	PublishOrderId *int64 `json:"PublishOrderId,omitempty" xml:"PublishOrderId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ss.SSSZ
	//
	// example:
	//
	// 123123
	PublishTime *string `json:"PublishTime,omitempty" xml:"PublishTime,omitempty"`
	// example:
	//
	// 865181640657408
	SiteId *string   `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	Steps  []*string `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
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

func (s *GetAppPublishStatusResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
