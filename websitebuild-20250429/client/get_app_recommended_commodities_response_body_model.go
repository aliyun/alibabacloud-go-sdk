// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppRecommendedCommoditiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetAppRecommendedCommoditiesResponseBody
	GetAccessDeniedDetail() *string
	SetAllowRetry(v bool) *GetAppRecommendedCommoditiesResponseBody
	GetAllowRetry() *bool
	SetAppName(v string) *GetAppRecommendedCommoditiesResponseBody
	GetAppName() *string
	SetDynamicCode(v string) *GetAppRecommendedCommoditiesResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *GetAppRecommendedCommoditiesResponseBody
	GetDynamicMessage() *string
	SetErrorArgs(v []interface{}) *GetAppRecommendedCommoditiesResponseBody
	GetErrorArgs() []interface{}
	SetModule(v *GetAppRecommendedCommoditiesResponseBodyModule) *GetAppRecommendedCommoditiesResponseBody
	GetModule() *GetAppRecommendedCommoditiesResponseBodyModule
	SetRequestId(v string) *GetAppRecommendedCommoditiesResponseBody
	GetRequestId() *string
	SetRootErrorCode(v string) *GetAppRecommendedCommoditiesResponseBody
	GetRootErrorCode() *string
	SetRootErrorMsg(v string) *GetAppRecommendedCommoditiesResponseBody
	GetRootErrorMsg() *string
	SetSynchro(v bool) *GetAppRecommendedCommoditiesResponseBody
	GetSynchro() *bool
}

type GetAppRecommendedCommoditiesResponseBody struct {
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
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The dynamic error code.
	//
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message, which is used to replace the `%s` placeholder in the **ErrMessage*	- parameter.
	//
	// > If **ErrMessage*	- returns **The Value of Input Parameter %s is not valid*	- and **DynamicMessage*	- returns **DtsJobId**, the value of the request parameter **DtsJobId*	- is invalid.
	//
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error parameters.
	ErrorArgs []interface{} `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	// The data table module. Valid values:
	//
	// - ABTest: experiment data table
	//
	// - ExperimentTool: experiment tool table
	//
	// - DataDiagnosis: data modeling diagnostics
	Module *GetAppRecommendedCommoditiesResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// Indicates whether the request is processed synchronously.
	//
	// example:
	//
	// True
	Synchro *bool `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetAppRecommendedCommoditiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppRecommendedCommoditiesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetAllowRetry() *bool {
	return s.AllowRetry
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetAppName() *string {
	return s.AppName
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetErrorArgs() []interface{} {
	return s.ErrorArgs
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetModule() *GetAppRecommendedCommoditiesResponseBodyModule {
	return s.Module
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetRootErrorCode() *string {
	return s.RootErrorCode
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetRootErrorMsg() *string {
	return s.RootErrorMsg
}

func (s *GetAppRecommendedCommoditiesResponseBody) GetSynchro() *bool {
	return s.Synchro
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetAccessDeniedDetail(v string) *GetAppRecommendedCommoditiesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetAllowRetry(v bool) *GetAppRecommendedCommoditiesResponseBody {
	s.AllowRetry = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetAppName(v string) *GetAppRecommendedCommoditiesResponseBody {
	s.AppName = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetDynamicCode(v string) *GetAppRecommendedCommoditiesResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetDynamicMessage(v string) *GetAppRecommendedCommoditiesResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetErrorArgs(v []interface{}) *GetAppRecommendedCommoditiesResponseBody {
	s.ErrorArgs = v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetModule(v *GetAppRecommendedCommoditiesResponseBodyModule) *GetAppRecommendedCommoditiesResponseBody {
	s.Module = v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetRequestId(v string) *GetAppRecommendedCommoditiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetRootErrorCode(v string) *GetAppRecommendedCommoditiesResponseBody {
	s.RootErrorCode = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetRootErrorMsg(v string) *GetAppRecommendedCommoditiesResponseBody {
	s.RootErrorMsg = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) SetSynchro(v bool) *GetAppRecommendedCommoditiesResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppRecommendedCommoditiesResponseBodyModule struct {
	// The list of marketing commodities.
	Commodities []*GetAppRecommendedCommoditiesResponseBodyModuleCommodities `json:"Commodities,omitempty" xml:"Commodities,omitempty" type:"Repeated"`
}

func (s GetAppRecommendedCommoditiesResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetAppRecommendedCommoditiesResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetAppRecommendedCommoditiesResponseBodyModule) GetCommodities() []*GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	return s.Commodities
}

func (s *GetAppRecommendedCommoditiesResponseBodyModule) SetCommodities(v []*GetAppRecommendedCommoditiesResponseBodyModuleCommodities) *GetAppRecommendedCommoditiesResponseBodyModule {
	s.Commodities = v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModule) Validate() error {
	if s.Commodities != nil {
		for _, item := range s.Commodities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAppRecommendedCommoditiesResponseBodyModuleCommodities struct {
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The commodity code. Applicable to both resource plans and marketing commodities.
	//
	// example:
	//
	// rds
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extension fields, such as unsupportedReason.
	Extend map[string]*string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The order type. Valid values:
	//
	// - BUY: purchase.
	//
	// - UPGRADE: upgrade.
	//
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The sort priority. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The marketing commodity ID. Returned only for new purchases.
	//
	// example:
	//
	// 12345
	PromotionCommodityId *string `json:"PromotionCommodityId,omitempty" xml:"PromotionCommodityId,omitempty"`
	RecommendType        *string `json:"RecommendType,omitempty" xml:"RecommendType,omitempty"`
	// The redirect URL. Returned when redirection is required, such as for upgrades.
	//
	// example:
	//
	// https://ecs-workbench-disposable.aliyun.com/account/disposable/login/sst/1291612921555690/edvo2gevfh
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	// The commodity status.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetAppRecommendedCommoditiesResponseBodyModuleCommodities) String() string {
	return dara.Prettify(s)
}

func (s GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GoString() string {
	return s.String()
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetActionType() *string {
	return s.ActionType
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetDescription() *string {
	return s.Description
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetExtend() map[string]*string {
	return s.Extend
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetOrderType() *string {
	return s.OrderType
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetPriority() *int32 {
	return s.Priority
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetPromotionCommodityId() *string {
	return s.PromotionCommodityId
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetRecommendType() *string {
	return s.RecommendType
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetStatus() *string {
	return s.Status
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetTitle() *string {
	return s.Title
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetActionType(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.ActionType = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetCommodityCode(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.CommodityCode = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetDescription(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.Description = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetExtend(v map[string]*string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.Extend = v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetOrderType(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.OrderType = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetPriority(v int32) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.Priority = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetPromotionCommodityId(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.PromotionCommodityId = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetRecommendType(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.RecommendType = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetRedirectUrl(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.RedirectUrl = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetStatus(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.Status = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetTitle(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.Title = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) Validate() error {
	return dara.Validate(s)
}
