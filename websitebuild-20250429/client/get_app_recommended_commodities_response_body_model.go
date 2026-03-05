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
	// spring-cloud-b
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// ERROR-oo1
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// example:
	//
	// SYSTEM_ERROR
	DynamicMessage *string                                         `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrorArgs      []interface{}                                   `json:"ErrorArgs,omitempty" xml:"ErrorArgs,omitempty" type:"Repeated"`
	Module         *GetAppRecommendedCommoditiesResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
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
	// example:
	//
	// rds
	CommodityCode *string            `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Extend        map[string]*string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// example:
	//
	// DESC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 12345
	PromotionCommodityId *string `json:"PromotionCommodityId,omitempty" xml:"PromotionCommodityId,omitempty"`
	// example:
	//
	// https://ecs-workbench-disposable.aliyun.com/account/disposable/login/sst/1291612921555690/edvo2gevfh
	RedirectUrl *string `json:"RedirectUrl,omitempty" xml:"RedirectUrl,omitempty"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAppRecommendedCommoditiesResponseBodyModuleCommodities) String() string {
	return dara.Prettify(s)
}

func (s GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GoString() string {
	return s.String()
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetCommodityCode() *string {
	return s.CommodityCode
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

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetRedirectUrl() *string {
	return s.RedirectUrl
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) GetStatus() *string {
	return s.Status
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetCommodityCode(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.CommodityCode = &v
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

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetRedirectUrl(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.RedirectUrl = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) SetStatus(v string) *GetAppRecommendedCommoditiesResponseBodyModuleCommodities {
	s.Status = &v
	return s
}

func (s *GetAppRecommendedCommoditiesResponseBodyModuleCommodities) Validate() error {
	return dara.Validate(s)
}
