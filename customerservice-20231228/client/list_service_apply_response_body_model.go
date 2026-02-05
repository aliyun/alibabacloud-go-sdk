// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListServiceApplyResponseBody
	GetCode() *string
	SetData(v *ListServiceApplyResponseBodyData) *ListServiceApplyResponseBody
	GetData() *ListServiceApplyResponseBodyData
	SetHttpStatusCode(v int32) *ListServiceApplyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListServiceApplyResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListServiceApplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListServiceApplyResponseBody
	GetSuccess() *bool
}

type ListServiceApplyResponseBody struct {
	Code           *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Data           *ListServiceApplyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	HttpStatusCode *int32                            `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string                           `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string                           `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool                             `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListServiceApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListServiceApplyResponseBody) GetData() *ListServiceApplyResponseBodyData {
	return s.Data
}

func (s *ListServiceApplyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListServiceApplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListServiceApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceApplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListServiceApplyResponseBody) SetCode(v string) *ListServiceApplyResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceApplyResponseBody) SetData(v *ListServiceApplyResponseBodyData) *ListServiceApplyResponseBody {
	s.Data = v
	return s
}

func (s *ListServiceApplyResponseBody) SetHttpStatusCode(v int32) *ListServiceApplyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListServiceApplyResponseBody) SetMessage(v string) *ListServiceApplyResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceApplyResponseBody) SetRequestId(v string) *ListServiceApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceApplyResponseBody) SetSuccess(v bool) *ListServiceApplyResponseBody {
	s.Success = &v
	return s
}

func (s *ListServiceApplyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceApplyResponseBodyData struct {
	Extend   interface{}                             `json:"extend,omitempty" xml:"extend,omitempty"`
	List     []*ListServiceApplyResponseBodyDataList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	PageNum  *int32                                  `json:"pageNum,omitempty" xml:"pageNum,omitempty"`
	PageSize *int32                                  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Total    *int32                                  `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListServiceApplyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyData) GetExtend() interface{} {
	return s.Extend
}

func (s *ListServiceApplyResponseBodyData) GetList() []*ListServiceApplyResponseBodyDataList {
	return s.List
}

func (s *ListServiceApplyResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListServiceApplyResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServiceApplyResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListServiceApplyResponseBodyData) SetExtend(v interface{}) *ListServiceApplyResponseBodyData {
	s.Extend = v
	return s
}

func (s *ListServiceApplyResponseBodyData) SetList(v []*ListServiceApplyResponseBodyDataList) *ListServiceApplyResponseBodyData {
	s.List = v
	return s
}

func (s *ListServiceApplyResponseBodyData) SetPageNum(v int32) *ListServiceApplyResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *ListServiceApplyResponseBodyData) SetPageSize(v int32) *ListServiceApplyResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListServiceApplyResponseBodyData) SetTotal(v int32) *ListServiceApplyResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListServiceApplyResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceApplyResponseBodyDataList struct {
	ApplierId                       *string                                                  `json:"applierId,omitempty" xml:"applierId,omitempty"`
	ApplyCode                       *string                                                  `json:"applyCode,omitempty" xml:"applyCode,omitempty"`
	ApplyComponentDetails           [][]*string                                              `json:"applyComponentDetails,omitempty" xml:"applyComponentDetails,omitempty" type:"Repeated"`
	ApplyTime                       *int64                                                   `json:"applyTime,omitempty" xml:"applyTime,omitempty"`
	Appointments                    []*ListServiceApplyResponseBodyDataListAppointments      `json:"appointments,omitempty" xml:"appointments,omitempty" type:"Repeated"`
	BuyUrl                          *string                                                  `json:"buyUrl,omitempty" xml:"buyUrl,omitempty"`
	CreatorEmpId                    *string                                                  `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CustomerName                    *string                                                  `json:"customerName,omitempty" xml:"customerName,omitempty"`
	CycleService                    *bool                                                    `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExecutedCount                   *int64                                                   `json:"executedCount,omitempty" xml:"executedCount,omitempty"`
	FinishCount                     *int64                                                   `json:"finishCount,omitempty" xml:"finishCount,omitempty"`
	GmtCreate                       *string                                                  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                  `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                   `json:"id,omitempty" xml:"id,omitempty"`
	IsOneToOneExpertServiceByTime   *bool                                                    `json:"isOneToOneExpertServiceByTime,omitempty" xml:"isOneToOneExpertServiceByTime,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                    `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                  `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	PackDetails                     []map[string]interface{}                                 `json:"packDetails,omitempty" xml:"packDetails,omitempty" type:"Repeated"`
	PayOrders                       []*ListServiceApplyResponseBodyDataListPayOrders         `json:"payOrders,omitempty" xml:"payOrders,omitempty" type:"Repeated"`
	PayUrl                          *string                                                  `json:"payUrl,omitempty" xml:"payUrl,omitempty"`
	PerformanceOrders               []*ListServiceApplyResponseBodyDataListPerformanceOrders `json:"performanceOrders,omitempty" xml:"performanceOrders,omitempty" type:"Repeated"`
	PerformancePacks                []*ListServiceApplyResponseBodyDataListPerformancePacks  `json:"performancePacks,omitempty" xml:"performancePacks,omitempty" type:"Repeated"`
	ReneWalUrl                      *string                                                  `json:"reneWalUrl,omitempty" xml:"reneWalUrl,omitempty"`
	ServiceCode                     *string                                                  `json:"serviceCode,omitempty" xml:"serviceCode,omitempty"`
	ServiceName                     *string                                                  `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	ServiceReports                  []*ListServiceApplyResponseBodyDataListServiceReports    `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceTimeRange                []*int64                                                 `json:"serviceTimeRange,omitempty" xml:"serviceTimeRange,omitempty" type:"Repeated"`
	Status                          *string                                                  `json:"status,omitempty" xml:"status,omitempty"`
	StatusCode                      *int32                                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	StatusStr                       *string                                                  `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	TermOfValidity                  *string                                                  `json:"termOfValidity,omitempty" xml:"termOfValidity,omitempty"`
	TotalPack                       *int32                                                   `json:"totalPack,omitempty" xml:"totalPack,omitempty"`
	Type                            *string                                                  `json:"type,omitempty" xml:"type,omitempty"`
	UsePack                         *int64                                                   `json:"usePack,omitempty" xml:"usePack,omitempty"`
	UserRights                      *string                                                  `json:"userRights,omitempty" xml:"userRights,omitempty"`
}

func (s ListServiceApplyResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataList) GetApplierId() *string {
	return s.ApplierId
}

func (s *ListServiceApplyResponseBodyDataList) GetApplyCode() *string {
	return s.ApplyCode
}

func (s *ListServiceApplyResponseBodyDataList) GetApplyComponentDetails() [][]*string {
	return s.ApplyComponentDetails
}

func (s *ListServiceApplyResponseBodyDataList) GetApplyTime() *int64 {
	return s.ApplyTime
}

func (s *ListServiceApplyResponseBodyDataList) GetAppointments() []*ListServiceApplyResponseBodyDataListAppointments {
	return s.Appointments
}

func (s *ListServiceApplyResponseBodyDataList) GetBuyUrl() *string {
	return s.BuyUrl
}

func (s *ListServiceApplyResponseBodyDataList) GetCreatorEmpId() *string {
	return s.CreatorEmpId
}

func (s *ListServiceApplyResponseBodyDataList) GetCustomerName() *string {
	return s.CustomerName
}

func (s *ListServiceApplyResponseBodyDataList) GetCycleService() *bool {
	return s.CycleService
}

func (s *ListServiceApplyResponseBodyDataList) GetExecutedCount() *int64 {
	return s.ExecutedCount
}

func (s *ListServiceApplyResponseBodyDataList) GetFinishCount() *int64 {
	return s.FinishCount
}

func (s *ListServiceApplyResponseBodyDataList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataList) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataList) GetIsOneToOneExpertServiceByTime() *bool {
	return s.IsOneToOneExpertServiceByTime
}

func (s *ListServiceApplyResponseBodyDataList) GetMergeSolutionAndReporterOneStep() *bool {
	return s.MergeSolutionAndReporterOneStep
}

func (s *ListServiceApplyResponseBodyDataList) GetModifierEmpId() *string {
	return s.ModifierEmpId
}

func (s *ListServiceApplyResponseBodyDataList) GetPackDetails() []map[string]interface{} {
	return s.PackDetails
}

func (s *ListServiceApplyResponseBodyDataList) GetPayOrders() []*ListServiceApplyResponseBodyDataListPayOrders {
	return s.PayOrders
}

func (s *ListServiceApplyResponseBodyDataList) GetPayUrl() *string {
	return s.PayUrl
}

func (s *ListServiceApplyResponseBodyDataList) GetPerformanceOrders() []*ListServiceApplyResponseBodyDataListPerformanceOrders {
	return s.PerformanceOrders
}

func (s *ListServiceApplyResponseBodyDataList) GetPerformancePacks() []*ListServiceApplyResponseBodyDataListPerformancePacks {
	return s.PerformancePacks
}

func (s *ListServiceApplyResponseBodyDataList) GetReneWalUrl() *string {
	return s.ReneWalUrl
}

func (s *ListServiceApplyResponseBodyDataList) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *ListServiceApplyResponseBodyDataList) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServiceApplyResponseBodyDataList) GetServiceReports() []*ListServiceApplyResponseBodyDataListServiceReports {
	return s.ServiceReports
}

func (s *ListServiceApplyResponseBodyDataList) GetServiceTimeRange() []*int64 {
	return s.ServiceTimeRange
}

func (s *ListServiceApplyResponseBodyDataList) GetStatus() *string {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataList) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceApplyResponseBodyDataList) GetStatusStr() *string {
	return s.StatusStr
}

func (s *ListServiceApplyResponseBodyDataList) GetTermOfValidity() *string {
	return s.TermOfValidity
}

func (s *ListServiceApplyResponseBodyDataList) GetTotalPack() *int32 {
	return s.TotalPack
}

func (s *ListServiceApplyResponseBodyDataList) GetType() *string {
	return s.Type
}

func (s *ListServiceApplyResponseBodyDataList) GetUsePack() *int64 {
	return s.UsePack
}

func (s *ListServiceApplyResponseBodyDataList) GetUserRights() *string {
	return s.UserRights
}

func (s *ListServiceApplyResponseBodyDataList) SetApplierId(v string) *ListServiceApplyResponseBodyDataList {
	s.ApplierId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetApplyCode(v string) *ListServiceApplyResponseBodyDataList {
	s.ApplyCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetApplyComponentDetails(v [][]*string) *ListServiceApplyResponseBodyDataList {
	s.ApplyComponentDetails = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetApplyTime(v int64) *ListServiceApplyResponseBodyDataList {
	s.ApplyTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetAppointments(v []*ListServiceApplyResponseBodyDataListAppointments) *ListServiceApplyResponseBodyDataList {
	s.Appointments = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetBuyUrl(v string) *ListServiceApplyResponseBodyDataList {
	s.BuyUrl = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataList {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetCustomerName(v string) *ListServiceApplyResponseBodyDataList {
	s.CustomerName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetCycleService(v bool) *ListServiceApplyResponseBodyDataList {
	s.CycleService = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetExecutedCount(v int64) *ListServiceApplyResponseBodyDataList {
	s.ExecutedCount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetFinishCount(v int64) *ListServiceApplyResponseBodyDataList {
	s.FinishCount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetGmtModified(v string) *ListServiceApplyResponseBodyDataList {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetId(v int64) *ListServiceApplyResponseBodyDataList {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetIsOneToOneExpertServiceByTime(v bool) *ListServiceApplyResponseBodyDataList {
	s.IsOneToOneExpertServiceByTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetMergeSolutionAndReporterOneStep(v bool) *ListServiceApplyResponseBodyDataList {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataList {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPackDetails(v []map[string]interface{}) *ListServiceApplyResponseBodyDataList {
	s.PackDetails = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPayOrders(v []*ListServiceApplyResponseBodyDataListPayOrders) *ListServiceApplyResponseBodyDataList {
	s.PayOrders = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPayUrl(v string) *ListServiceApplyResponseBodyDataList {
	s.PayUrl = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPerformanceOrders(v []*ListServiceApplyResponseBodyDataListPerformanceOrders) *ListServiceApplyResponseBodyDataList {
	s.PerformanceOrders = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetPerformancePacks(v []*ListServiceApplyResponseBodyDataListPerformancePacks) *ListServiceApplyResponseBodyDataList {
	s.PerformancePacks = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetReneWalUrl(v string) *ListServiceApplyResponseBodyDataList {
	s.ReneWalUrl = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetServiceCode(v string) *ListServiceApplyResponseBodyDataList {
	s.ServiceCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetServiceName(v string) *ListServiceApplyResponseBodyDataList {
	s.ServiceName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetServiceReports(v []*ListServiceApplyResponseBodyDataListServiceReports) *ListServiceApplyResponseBodyDataList {
	s.ServiceReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetServiceTimeRange(v []*int64) *ListServiceApplyResponseBodyDataList {
	s.ServiceTimeRange = v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetStatus(v string) *ListServiceApplyResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetStatusCode(v int32) *ListServiceApplyResponseBodyDataList {
	s.StatusCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetStatusStr(v string) *ListServiceApplyResponseBodyDataList {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetTermOfValidity(v string) *ListServiceApplyResponseBodyDataList {
	s.TermOfValidity = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetTotalPack(v int32) *ListServiceApplyResponseBodyDataList {
	s.TotalPack = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetType(v string) *ListServiceApplyResponseBodyDataList {
	s.Type = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetUsePack(v int64) *ListServiceApplyResponseBodyDataList {
	s.UsePack = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) SetUserRights(v string) *ListServiceApplyResponseBodyDataList {
	s.UserRights = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataList) Validate() error {
	if s.Appointments != nil {
		for _, item := range s.Appointments {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PayOrders != nil {
		for _, item := range s.PayOrders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PerformanceOrders != nil {
		for _, item := range s.PerformanceOrders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PerformancePacks != nil {
		for _, item := range s.PerformancePacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceReports != nil {
		for _, item := range s.ServiceReports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceApplyResponseBodyDataListAppointments struct {
	HuhangId     *int64  `json:"huhangId,omitempty" xml:"huhangId,omitempty"`
	PurchaseCode *int32  `json:"purchaseCode,omitempty" xml:"purchaseCode,omitempty"`
	PurchaseDesc *string `json:"purchaseDesc,omitempty" xml:"purchaseDesc,omitempty"`
	SupportDays  *int32  `json:"supportDays,omitempty" xml:"supportDays,omitempty"`
	TravelDays   *int32  `json:"travelDays,omitempty" xml:"travelDays,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListAppointments) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListAppointments) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListAppointments) GetHuhangId() *int64 {
	return s.HuhangId
}

func (s *ListServiceApplyResponseBodyDataListAppointments) GetPurchaseCode() *int32 {
	return s.PurchaseCode
}

func (s *ListServiceApplyResponseBodyDataListAppointments) GetPurchaseDesc() *string {
	return s.PurchaseDesc
}

func (s *ListServiceApplyResponseBodyDataListAppointments) GetSupportDays() *int32 {
	return s.SupportDays
}

func (s *ListServiceApplyResponseBodyDataListAppointments) GetTravelDays() *int32 {
	return s.TravelDays
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetHuhangId(v int64) *ListServiceApplyResponseBodyDataListAppointments {
	s.HuhangId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetPurchaseCode(v int32) *ListServiceApplyResponseBodyDataListAppointments {
	s.PurchaseCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetPurchaseDesc(v string) *ListServiceApplyResponseBodyDataListAppointments {
	s.PurchaseDesc = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetSupportDays(v int32) *ListServiceApplyResponseBodyDataListAppointments {
	s.SupportDays = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListAppointments) SetTravelDays(v int32) *ListServiceApplyResponseBodyDataListAppointments {
	s.TravelDays = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListAppointments) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPayOrders struct {
	Amount               *string                `json:"amount,omitempty" xml:"amount,omitempty"`
	CompassCommodityCode *string                `json:"compassCommodityCode,omitempty" xml:"compassCommodityCode,omitempty"`
	CompassCommodityName *string                `json:"compassCommodityName,omitempty" xml:"compassCommodityName,omitempty"`
	CreatorEmpId         *string                `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate            *string                `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified          *string                `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                   *int64                 `json:"id,omitempty" xml:"id,omitempty"`
	ModifierEmpId        *string                `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Operate              map[string]interface{} `json:"operate,omitempty" xml:"operate,omitempty"`
	OrderDetail          interface{}            `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId              *int64                 `json:"orderId,omitempty" xml:"orderId,omitempty"`
	OriginalPrice        *float64               `json:"originalPrice,omitempty" xml:"originalPrice,omitempty"`
	PayAmount            *float64               `json:"payAmount,omitempty" xml:"payAmount,omitempty"`
	PayTime              *string                `json:"payTime,omitempty" xml:"payTime,omitempty"`
	ProductName          *string                `json:"productName,omitempty" xml:"productName,omitempty"`
	ReneWalUrl           *string                `json:"reneWalUrl,omitempty" xml:"reneWalUrl,omitempty"`
	ServiceContentMap    map[string]interface{} `json:"serviceContentMap,omitempty" xml:"serviceContentMap,omitempty"`
	Status               *int32                 `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr            *string                `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportDays          *int32                 `json:"supportDays,omitempty" xml:"supportDays,omitempty"`
	Uid                  *string                `json:"uid,omitempty" xml:"uid,omitempty"`
	Url                  *string                `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPayOrders) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPayOrders) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetAmount() *string {
	return s.Amount
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetCompassCommodityCode() *string {
	return s.CompassCommodityCode
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetCompassCommodityName() *string {
	return s.CompassCommodityName
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetCreatorEmpId() *string {
	return s.CreatorEmpId
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetModifierEmpId() *string {
	return s.ModifierEmpId
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetOperate() map[string]interface{} {
	return s.Operate
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetOrderDetail() interface{} {
	return s.OrderDetail
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetOriginalPrice() *float64 {
	return s.OriginalPrice
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetPayAmount() *float64 {
	return s.PayAmount
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetPayTime() *string {
	return s.PayTime
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetProductName() *string {
	return s.ProductName
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetReneWalUrl() *string {
	return s.ReneWalUrl
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetServiceContentMap() map[string]interface{} {
	return s.ServiceContentMap
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetStatusStr() *string {
	return s.StatusStr
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetSupportDays() *int32 {
	return s.SupportDays
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetUid() *string {
	return s.Uid
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetAmount(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Amount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetCompassCommodityCode(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.CompassCommodityCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetCompassCommodityName(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.CompassCommodityName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetId(v int64) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetOperate(v map[string]interface{}) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Operate = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetOrderDetail(v interface{}) *ListServiceApplyResponseBodyDataListPayOrders {
	s.OrderDetail = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetOrderId(v int64) *ListServiceApplyResponseBodyDataListPayOrders {
	s.OrderId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetOriginalPrice(v float64) *ListServiceApplyResponseBodyDataListPayOrders {
	s.OriginalPrice = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetPayAmount(v float64) *ListServiceApplyResponseBodyDataListPayOrders {
	s.PayAmount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetPayTime(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.PayTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetProductName(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.ProductName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetReneWalUrl(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.ReneWalUrl = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetServiceContentMap(v map[string]interface{}) *ListServiceApplyResponseBodyDataListPayOrders {
	s.ServiceContentMap = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetStatusStr(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetSupportDays(v int32) *ListServiceApplyResponseBodyDataListPayOrders {
	s.SupportDays = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetUid(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Uid = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) SetUrl(v string) *ListServiceApplyResponseBodyDataListPayOrders {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPayOrders) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrders struct {
	ApplyFileVOList                 []*ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList     `json:"applyFileVOList,omitempty" xml:"applyFileVOList,omitempty" type:"Repeated"`
	AppointmentCode                 *string                                                                     `json:"appointmentCode,omitempty" xml:"appointmentCode,omitempty"`
	AppointmentEndTime              *int64                                                                      `json:"appointmentEndTime,omitempty" xml:"appointmentEndTime,omitempty"`
	AppointmentId                   *string                                                                     `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	AppointmentPassTime             *int64                                                                      `json:"appointmentPassTime,omitempty" xml:"appointmentPassTime,omitempty"`
	AppointmentTime                 *int64                                                                      `json:"appointmentTime,omitempty" xml:"appointmentTime,omitempty"`
	CommodityDesc                   *string                                                                     `json:"commodityDesc,omitempty" xml:"commodityDesc,omitempty"`
	CreatorEmpId                    *string                                                                     `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CycleService                    *bool                                                                       `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExtList                         []*ListServiceApplyResponseBodyDataListPerformanceOrdersExtList             `json:"extList,omitempty" xml:"extList,omitempty" type:"Repeated"`
	GmtCreate                       *string                                                                     `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                                     `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                                      `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                                       `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                                     `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	NtmCommodityCode                *string                                                                     `json:"ntmCommodityCode,omitempty" xml:"ntmCommodityCode,omitempty"`
	OrderDetail                     interface{}                                                                 `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId                         *int64                                                                      `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PackCount                       *int32                                                                      `json:"packCount,omitempty" xml:"packCount,omitempty"`
	PackDetails                     []map[string]interface{}                                                    `json:"packDetails,omitempty" xml:"packDetails,omitempty" type:"Repeated"`
	PerformanceNodeDTOS             []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS `json:"performanceNodeDTOS,omitempty" xml:"performanceNodeDTOS,omitempty" type:"Repeated"`
	PerformancePacks                []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks    `json:"performancePacks,omitempty" xml:"performancePacks,omitempty" type:"Repeated"`
	PurchasePackCode                *int32                                                                      `json:"purchasePackCode,omitempty" xml:"purchasePackCode,omitempty"`
	ServiceApplyId                  *int64                                                                      `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	ServiceMonthReports             []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports `json:"serviceMonthReports,omitempty" xml:"serviceMonthReports,omitempty" type:"Repeated"`
	ServiceReports                  []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports      `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceSchemes                  []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes      `json:"serviceSchemes,omitempty" xml:"serviceSchemes,omitempty" type:"Repeated"`
	Status                          *int32                                                                      `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr                       *string                                                                     `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportTime                     []*int64                                                                    `json:"supportTime,omitempty" xml:"supportTime,omitempty" type:"Repeated"`
	TamEngineers                    []*ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers        `json:"tamEngineers,omitempty" xml:"tamEngineers,omitempty" type:"Repeated"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrders) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrders) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetApplyFileVOList() []*ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	return s.ApplyFileVOList
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetAppointmentCode() *string {
	return s.AppointmentCode
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetAppointmentEndTime() *int64 {
	return s.AppointmentEndTime
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetAppointmentPassTime() *int64 {
	return s.AppointmentPassTime
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetAppointmentTime() *int64 {
	return s.AppointmentTime
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetCommodityDesc() *string {
	return s.CommodityDesc
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetCreatorEmpId() *string {
	return s.CreatorEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetCycleService() *bool {
	return s.CycleService
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetExtList() []*ListServiceApplyResponseBodyDataListPerformanceOrdersExtList {
	return s.ExtList
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetMergeSolutionAndReporterOneStep() *bool {
	return s.MergeSolutionAndReporterOneStep
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetModifierEmpId() *string {
	return s.ModifierEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetNtmCommodityCode() *string {
	return s.NtmCommodityCode
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetOrderDetail() interface{} {
	return s.OrderDetail
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetPackCount() *int32 {
	return s.PackCount
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetPackDetails() []map[string]interface{} {
	return s.PackDetails
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetPerformanceNodeDTOS() []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	return s.PerformanceNodeDTOS
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetPerformancePacks() []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	return s.PerformancePacks
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetPurchasePackCode() *int32 {
	return s.PurchasePackCode
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetServiceMonthReports() []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	return s.ServiceMonthReports
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetServiceReports() []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	return s.ServiceReports
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetServiceSchemes() []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	return s.ServiceSchemes
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetStatusStr() *string {
	return s.StatusStr
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetSupportTime() []*int64 {
	return s.SupportTime
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) GetTamEngineers() []*ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	return s.TamEngineers
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetApplyFileVOList(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ApplyFileVOList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentEndTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentEndTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentPassTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentPassTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetAppointmentTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.AppointmentTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetCommodityDesc(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.CommodityDesc = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetCycleService(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.CycleService = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetExtList(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ExtList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetMergeSolutionAndReporterOneStep(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetNtmCommodityCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.NtmCommodityCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetOrderDetail(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.OrderDetail = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetOrderId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.OrderId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPackCount(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PackCount = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPackDetails(v []map[string]interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PackDetails = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPerformanceNodeDTOS(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PerformanceNodeDTOS = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPerformancePacks(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PerformancePacks = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetPurchasePackCode(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.PurchasePackCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetServiceMonthReports(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ServiceMonthReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetServiceReports(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ServiceReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetServiceSchemes(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.ServiceSchemes = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetStatusStr(v string) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetSupportTime(v []*int64) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.SupportTime = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) SetTamEngineers(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) *ListServiceApplyResponseBodyDataListPerformanceOrders {
	s.TamEngineers = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrders) Validate() error {
	if s.ApplyFileVOList != nil {
		for _, item := range s.ApplyFileVOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExtList != nil {
		for _, item := range s.ExtList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PerformanceNodeDTOS != nil {
		for _, item := range s.PerformanceNodeDTOS {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PerformancePacks != nil {
		for _, item := range s.PerformancePacks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceMonthReports != nil {
		for _, item := range s.ServiceMonthReports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceReports != nil {
		for _, item := range s.ServiceReports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceSchemes != nil {
		for _, item := range s.ServiceSchemes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TamEngineers != nil {
		for _, item := range s.TamEngineers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersApplyFileVOList) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersExtList struct {
	KeyCode *string     `json:"keyCode,omitempty" xml:"keyCode,omitempty"`
	Name    *string     `json:"name,omitempty" xml:"name,omitempty"`
	Value   interface{} `json:"value,omitempty" xml:"value,omitempty"`
	View    *string     `json:"view,omitempty" xml:"view,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) GetKeyCode() *string {
	return s.KeyCode
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) GetName() *string {
	return s.Name
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) GetValue() interface{} {
	return s.Value
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) GetView() *string {
	return s.View
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) SetKeyCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList {
	s.KeyCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) SetName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) SetValue(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList {
	s.Value = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) SetView(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList {
	s.View = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersExtList) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS struct {
	Display    *bool       `json:"display,omitempty" xml:"display,omitempty"`
	ExtendInfo interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Index      *int32      `json:"index,omitempty" xml:"index,omitempty"`
	NodeName   *string     `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	Status     *int32      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) GetDisplay() *bool {
	return s.Display
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) GetExtendInfo() interface{} {
	return s.ExtendInfo
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) GetIndex() *int32 {
	return s.Index
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) GetNodeName() *string {
	return s.NodeName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetDisplay(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.Display = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetExtendInfo(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.ExtendInfo = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetIndex(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.Index = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetNodeName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.NodeName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformanceNodeDTOS) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks struct {
	ApplyFileVOList                 []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList     `json:"applyFileVOList,omitempty" xml:"applyFileVOList,omitempty" type:"Repeated"`
	AppointmentCode                 *string                                                                                     `json:"appointmentCode,omitempty" xml:"appointmentCode,omitempty"`
	AppointmentEndTime              *int64                                                                                      `json:"appointmentEndTime,omitempty" xml:"appointmentEndTime,omitempty"`
	AppointmentId                   *string                                                                                     `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	AppointmentPassTime             *int64                                                                                      `json:"appointmentPassTime,omitempty" xml:"appointmentPassTime,omitempty"`
	AppointmentTime                 *int64                                                                                      `json:"appointmentTime,omitempty" xml:"appointmentTime,omitempty"`
	CommodityDesc                   *string                                                                                     `json:"commodityDesc,omitempty" xml:"commodityDesc,omitempty"`
	CreatorEmpId                    *string                                                                                     `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CycleService                    *bool                                                                                       `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExtList                         []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList             `json:"extList,omitempty" xml:"extList,omitempty" type:"Repeated"`
	GmtCreate                       *string                                                                                     `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                                                     `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                                                      `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                                                       `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                                                     `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	NtmCommodityCode                *string                                                                                     `json:"ntmCommodityCode,omitempty" xml:"ntmCommodityCode,omitempty"`
	OrderDetail                     interface{}                                                                                 `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId                         *int64                                                                                      `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PerformanceNodeDTOS             []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS `json:"performanceNodeDTOS,omitempty" xml:"performanceNodeDTOS,omitempty" type:"Repeated"`
	PurchasePackCode                *int32                                                                                      `json:"purchasePackCode,omitempty" xml:"purchasePackCode,omitempty"`
	ServiceApplyId                  *int64                                                                                      `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	ServiceMonthReports             []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports `json:"serviceMonthReports,omitempty" xml:"serviceMonthReports,omitempty" type:"Repeated"`
	ServiceReports                  []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports      `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceSchemes                  []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes      `json:"serviceSchemes,omitempty" xml:"serviceSchemes,omitempty" type:"Repeated"`
	Status                          *int32                                                                                      `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr                       *string                                                                                     `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportTime                     []*int64                                                                                    `json:"supportTime,omitempty" xml:"supportTime,omitempty" type:"Repeated"`
	TamEngineers                    []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers        `json:"tamEngineers,omitempty" xml:"tamEngineers,omitempty" type:"Repeated"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetApplyFileVOList() []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	return s.ApplyFileVOList
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetAppointmentCode() *string {
	return s.AppointmentCode
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetAppointmentEndTime() *int64 {
	return s.AppointmentEndTime
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetAppointmentPassTime() *int64 {
	return s.AppointmentPassTime
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetAppointmentTime() *int64 {
	return s.AppointmentTime
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetCommodityDesc() *string {
	return s.CommodityDesc
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetCreatorEmpId() *string {
	return s.CreatorEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetCycleService() *bool {
	return s.CycleService
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetExtList() []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList {
	return s.ExtList
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetMergeSolutionAndReporterOneStep() *bool {
	return s.MergeSolutionAndReporterOneStep
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetModifierEmpId() *string {
	return s.ModifierEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetNtmCommodityCode() *string {
	return s.NtmCommodityCode
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetOrderDetail() interface{} {
	return s.OrderDetail
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetPerformanceNodeDTOS() []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	return s.PerformanceNodeDTOS
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetPurchasePackCode() *int32 {
	return s.PurchasePackCode
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetServiceMonthReports() []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	return s.ServiceMonthReports
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetServiceReports() []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	return s.ServiceReports
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetServiceSchemes() []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	return s.ServiceSchemes
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetStatusStr() *string {
	return s.StatusStr
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetSupportTime() []*int64 {
	return s.SupportTime
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) GetTamEngineers() []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	return s.TamEngineers
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetApplyFileVOList(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ApplyFileVOList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentEndTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentEndTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentPassTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentPassTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetAppointmentTime(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.AppointmentTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetCommodityDesc(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.CommodityDesc = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetCycleService(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.CycleService = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetExtList(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ExtList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetMergeSolutionAndReporterOneStep(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetNtmCommodityCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.NtmCommodityCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetOrderDetail(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.OrderDetail = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetOrderId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.OrderId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetPerformanceNodeDTOS(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.PerformanceNodeDTOS = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetPurchasePackCode(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.PurchasePackCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetServiceMonthReports(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ServiceMonthReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetServiceReports(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ServiceReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetServiceSchemes(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.ServiceSchemes = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetStatusStr(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetSupportTime(v []*int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.SupportTime = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) SetTamEngineers(v []*ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks {
	s.TamEngineers = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacks) Validate() error {
	if s.ApplyFileVOList != nil {
		for _, item := range s.ApplyFileVOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExtList != nil {
		for _, item := range s.ExtList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PerformanceNodeDTOS != nil {
		for _, item := range s.PerformanceNodeDTOS {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceMonthReports != nil {
		for _, item := range s.ServiceMonthReports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceReports != nil {
		for _, item := range s.ServiceReports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceSchemes != nil {
		for _, item := range s.ServiceSchemes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TamEngineers != nil {
		for _, item := range s.TamEngineers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksApplyFileVOList) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList struct {
	KeyCode *string     `json:"keyCode,omitempty" xml:"keyCode,omitempty"`
	Name    *string     `json:"name,omitempty" xml:"name,omitempty"`
	Value   interface{} `json:"value,omitempty" xml:"value,omitempty"`
	View    *string     `json:"view,omitempty" xml:"view,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) GetKeyCode() *string {
	return s.KeyCode
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) GetName() *string {
	return s.Name
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) GetValue() interface{} {
	return s.Value
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) GetView() *string {
	return s.View
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) SetKeyCode(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList {
	s.KeyCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) SetName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) SetValue(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList {
	s.Value = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) SetView(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList {
	s.View = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksExtList) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS struct {
	Display    *bool       `json:"display,omitempty" xml:"display,omitempty"`
	ExtendInfo interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Index      *int32      `json:"index,omitempty" xml:"index,omitempty"`
	NodeName   *string     `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	Status     *int32      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) GetDisplay() *bool {
	return s.Display
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) GetExtendInfo() interface{} {
	return s.ExtendInfo
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) GetIndex() *int32 {
	return s.Index
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) GetNodeName() *string {
	return s.NodeName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetDisplay(v bool) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.Display = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetExtendInfo(v interface{}) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.ExtendInfo = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetIndex(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.Index = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetNodeName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.NodeName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksPerformanceNodeDTOS) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceMonthReports) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceReports) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksServiceSchemes) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers struct {
	CreatorEmpId  *string `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HrStatus      *string `json:"hrStatus,omitempty" xml:"hrStatus,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastName      *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	ModifierEmpId *string `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	NickNameEn    *string `json:"nickNameEn,omitempty" xml:"nickNameEn,omitempty"`
	RealmId       *int64  `json:"realmId,omitempty" xml:"realmId,omitempty"`
	Workid        *string `json:"workid,omitempty" xml:"workid,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetCreatorEmpId() *string {
	return s.CreatorEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetHrStatus() *string {
	return s.HrStatus
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetLastName() *string {
	return s.LastName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetModifierEmpId() *string {
	return s.ModifierEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetName() *string {
	return s.Name
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetNickNameEn() *string {
	return s.NickNameEn
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetRealmId() *int64 {
	return s.RealmId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) GetWorkid() *string {
	return s.Workid
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetHrStatus(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.HrStatus = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetLastName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.LastName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetNickNameEn(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.NickNameEn = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetRealmId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.RealmId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) SetWorkid(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers {
	s.Workid = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersPerformancePacksTamEngineers) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceMonthReports) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceReports) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersServiceSchemes) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers struct {
	CreatorEmpId  *string `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HrStatus      *string `json:"hrStatus,omitempty" xml:"hrStatus,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastName      *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	ModifierEmpId *string `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	NickNameEn    *string `json:"nickNameEn,omitempty" xml:"nickNameEn,omitempty"`
	RealmId       *int64  `json:"realmId,omitempty" xml:"realmId,omitempty"`
	Workid        *string `json:"workid,omitempty" xml:"workid,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetCreatorEmpId() *string {
	return s.CreatorEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetHrStatus() *string {
	return s.HrStatus
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetLastName() *string {
	return s.LastName
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetModifierEmpId() *string {
	return s.ModifierEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetName() *string {
	return s.Name
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetNickNameEn() *string {
	return s.NickNameEn
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetRealmId() *int64 {
	return s.RealmId
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) GetWorkid() *string {
	return s.Workid
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetHrStatus(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.HrStatus = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetLastName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.LastName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetName(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetNickNameEn(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.NickNameEn = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetRealmId(v int64) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.RealmId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) SetWorkid(v string) *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers {
	s.Workid = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformanceOrdersTamEngineers) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformancePacks struct {
	ApplyFileVOList                 []*ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList     `json:"applyFileVOList,omitempty" xml:"applyFileVOList,omitempty" type:"Repeated"`
	AppointmentCode                 *string                                                                    `json:"appointmentCode,omitempty" xml:"appointmentCode,omitempty"`
	AppointmentEndTime              *int64                                                                     `json:"appointmentEndTime,omitempty" xml:"appointmentEndTime,omitempty"`
	AppointmentId                   *string                                                                    `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	AppointmentPassTime             *int64                                                                     `json:"appointmentPassTime,omitempty" xml:"appointmentPassTime,omitempty"`
	AppointmentTime                 *int64                                                                     `json:"appointmentTime,omitempty" xml:"appointmentTime,omitempty"`
	CommodityDesc                   *string                                                                    `json:"commodityDesc,omitempty" xml:"commodityDesc,omitempty"`
	CreatorEmpId                    *string                                                                    `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	CycleService                    *bool                                                                      `json:"cycleService,omitempty" xml:"cycleService,omitempty"`
	ExtList                         []*ListServiceApplyResponseBodyDataListPerformancePacksExtList             `json:"extList,omitempty" xml:"extList,omitempty" type:"Repeated"`
	GmtCreate                       *string                                                                    `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified                     *string                                                                    `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id                              *int64                                                                     `json:"id,omitempty" xml:"id,omitempty"`
	MergeSolutionAndReporterOneStep *bool                                                                      `json:"mergeSolutionAndReporterOneStep,omitempty" xml:"mergeSolutionAndReporterOneStep,omitempty"`
	ModifierEmpId                   *string                                                                    `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	NtmCommodityCode                *string                                                                    `json:"ntmCommodityCode,omitempty" xml:"ntmCommodityCode,omitempty"`
	OrderDetail                     interface{}                                                                `json:"orderDetail,omitempty" xml:"orderDetail,omitempty"`
	OrderId                         *int64                                                                     `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PerformanceNodeDTOS             []*ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS `json:"performanceNodeDTOS,omitempty" xml:"performanceNodeDTOS,omitempty" type:"Repeated"`
	PurchasePackCode                *int32                                                                     `json:"purchasePackCode,omitempty" xml:"purchasePackCode,omitempty"`
	ServiceApplyId                  *int64                                                                     `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	ServiceMonthReports             []*ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports `json:"serviceMonthReports,omitempty" xml:"serviceMonthReports,omitempty" type:"Repeated"`
	ServiceReports                  []*ListServiceApplyResponseBodyDataListPerformancePacksServiceReports      `json:"serviceReports,omitempty" xml:"serviceReports,omitempty" type:"Repeated"`
	ServiceSchemes                  []*ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes      `json:"serviceSchemes,omitempty" xml:"serviceSchemes,omitempty" type:"Repeated"`
	Status                          *int32                                                                     `json:"status,omitempty" xml:"status,omitempty"`
	StatusStr                       *string                                                                    `json:"statusStr,omitempty" xml:"statusStr,omitempty"`
	SupportTime                     []*int64                                                                   `json:"supportTime,omitempty" xml:"supportTime,omitempty" type:"Repeated"`
	TamEngineers                    []*ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers        `json:"tamEngineers,omitempty" xml:"tamEngineers,omitempty" type:"Repeated"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacks) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacks) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetApplyFileVOList() []*ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	return s.ApplyFileVOList
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetAppointmentCode() *string {
	return s.AppointmentCode
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetAppointmentEndTime() *int64 {
	return s.AppointmentEndTime
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetAppointmentPassTime() *int64 {
	return s.AppointmentPassTime
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetAppointmentTime() *int64 {
	return s.AppointmentTime
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetCommodityDesc() *string {
	return s.CommodityDesc
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetCreatorEmpId() *string {
	return s.CreatorEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetCycleService() *bool {
	return s.CycleService
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetExtList() []*ListServiceApplyResponseBodyDataListPerformancePacksExtList {
	return s.ExtList
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetMergeSolutionAndReporterOneStep() *bool {
	return s.MergeSolutionAndReporterOneStep
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetModifierEmpId() *string {
	return s.ModifierEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetNtmCommodityCode() *string {
	return s.NtmCommodityCode
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetOrderDetail() interface{} {
	return s.OrderDetail
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetOrderId() *int64 {
	return s.OrderId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetPerformanceNodeDTOS() []*ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	return s.PerformanceNodeDTOS
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetPurchasePackCode() *int32 {
	return s.PurchasePackCode
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetServiceMonthReports() []*ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	return s.ServiceMonthReports
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetServiceReports() []*ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	return s.ServiceReports
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetServiceSchemes() []*ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	return s.ServiceSchemes
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetStatusStr() *string {
	return s.StatusStr
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetSupportTime() []*int64 {
	return s.SupportTime
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) GetTamEngineers() []*ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	return s.TamEngineers
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetApplyFileVOList(v []*ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ApplyFileVOList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentCode(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentEndTime(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentEndTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentPassTime(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentPassTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetAppointmentTime(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.AppointmentTime = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetCommodityDesc(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.CommodityDesc = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetCycleService(v bool) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.CycleService = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetExtList(v []*ListServiceApplyResponseBodyDataListPerformancePacksExtList) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ExtList = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetMergeSolutionAndReporterOneStep(v bool) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.MergeSolutionAndReporterOneStep = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetNtmCommodityCode(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.NtmCommodityCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetOrderDetail(v interface{}) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.OrderDetail = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetOrderId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.OrderId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetPerformanceNodeDTOS(v []*ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.PerformanceNodeDTOS = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetPurchasePackCode(v int32) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.PurchasePackCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetServiceMonthReports(v []*ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ServiceMonthReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetServiceReports(v []*ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ServiceReports = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetServiceSchemes(v []*ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.ServiceSchemes = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetStatusStr(v string) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.StatusStr = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetSupportTime(v []*int64) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.SupportTime = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) SetTamEngineers(v []*ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) *ListServiceApplyResponseBodyDataListPerformancePacks {
	s.TamEngineers = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacks) Validate() error {
	if s.ApplyFileVOList != nil {
		for _, item := range s.ApplyFileVOList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ExtList != nil {
		for _, item := range s.ExtList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PerformanceNodeDTOS != nil {
		for _, item := range s.PerformanceNodeDTOS {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceMonthReports != nil {
		for _, item := range s.ServiceMonthReports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceReports != nil {
		for _, item := range s.ServiceReports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ServiceSchemes != nil {
		for _, item := range s.ServiceSchemes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TamEngineers != nil {
		for _, item := range s.TamEngineers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksApplyFileVOList) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformancePacksExtList struct {
	KeyCode *string     `json:"keyCode,omitempty" xml:"keyCode,omitempty"`
	Name    *string     `json:"name,omitempty" xml:"name,omitempty"`
	Value   interface{} `json:"value,omitempty" xml:"value,omitempty"`
	View    *string     `json:"view,omitempty" xml:"view,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksExtList) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksExtList) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) GetKeyCode() *string {
	return s.KeyCode
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) GetName() *string {
	return s.Name
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) GetValue() interface{} {
	return s.Value
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) GetView() *string {
	return s.View
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) SetKeyCode(v string) *ListServiceApplyResponseBodyDataListPerformancePacksExtList {
	s.KeyCode = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) SetName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksExtList {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) SetValue(v interface{}) *ListServiceApplyResponseBodyDataListPerformancePacksExtList {
	s.Value = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) SetView(v string) *ListServiceApplyResponseBodyDataListPerformancePacksExtList {
	s.View = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksExtList) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS struct {
	Display    *bool       `json:"display,omitempty" xml:"display,omitempty"`
	ExtendInfo interface{} `json:"extendInfo,omitempty" xml:"extendInfo,omitempty"`
	Index      *int32      `json:"index,omitempty" xml:"index,omitempty"`
	NodeName   *string     `json:"nodeName,omitempty" xml:"nodeName,omitempty"`
	Status     *int32      `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) GetDisplay() *bool {
	return s.Display
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) GetExtendInfo() interface{} {
	return s.ExtendInfo
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) GetIndex() *int32 {
	return s.Index
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) GetNodeName() *string {
	return s.NodeName
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetDisplay(v bool) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.Display = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetExtendInfo(v interface{}) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.ExtendInfo = v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetIndex(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.Index = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetNodeName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.NodeName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksPerformanceNodeDTOS) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceMonthReports) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformancePacksServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceReports) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetFileName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetFileType(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetRemarke(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetStatus(v int32) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) SetUrl(v string) *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksServiceSchemes) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers struct {
	CreatorEmpId  *string `json:"creatorEmpId,omitempty" xml:"creatorEmpId,omitempty"`
	GmtCreate     *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified   *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	HrStatus      *string `json:"hrStatus,omitempty" xml:"hrStatus,omitempty"`
	Id            *int64  `json:"id,omitempty" xml:"id,omitempty"`
	LastName      *string `json:"lastName,omitempty" xml:"lastName,omitempty"`
	ModifierEmpId *string `json:"modifierEmpId,omitempty" xml:"modifierEmpId,omitempty"`
	Name          *string `json:"name,omitempty" xml:"name,omitempty"`
	NickNameEn    *string `json:"nickNameEn,omitempty" xml:"nickNameEn,omitempty"`
	RealmId       *int64  `json:"realmId,omitempty" xml:"realmId,omitempty"`
	Workid        *string `json:"workid,omitempty" xml:"workid,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetCreatorEmpId() *string {
	return s.CreatorEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetHrStatus() *string {
	return s.HrStatus
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetLastName() *string {
	return s.LastName
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetModifierEmpId() *string {
	return s.ModifierEmpId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetName() *string {
	return s.Name
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetNickNameEn() *string {
	return s.NickNameEn
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetRealmId() *int64 {
	return s.RealmId
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) GetWorkid() *string {
	return s.Workid
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetCreatorEmpId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.CreatorEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetHrStatus(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.HrStatus = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetLastName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.LastName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetModifierEmpId(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.ModifierEmpId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetName(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.Name = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetNickNameEn(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.NickNameEn = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetRealmId(v int64) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.RealmId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) SetWorkid(v string) *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers {
	s.Workid = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListPerformancePacksTamEngineers) Validate() error {
	return dara.Validate(s)
}

type ListServiceApplyResponseBodyDataListServiceReports struct {
	AppointmentId  *string `json:"appointmentId,omitempty" xml:"appointmentId,omitempty"`
	BatchGroup     *string `json:"batchGroup,omitempty" xml:"batchGroup,omitempty"`
	CustomerId     *string `json:"customerId,omitempty" xml:"customerId,omitempty"`
	FileName       *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	FileType       *int32  `json:"fileType,omitempty" xml:"fileType,omitempty"`
	GmtCreate      *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GmtModified    *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	Id             *int64  `json:"id,omitempty" xml:"id,omitempty"`
	Remarke        *string `json:"remarke,omitempty" xml:"remarke,omitempty"`
	ServiceApplyId *int64  `json:"serviceApplyId,omitempty" xml:"serviceApplyId,omitempty"`
	Status         *int32  `json:"status,omitempty" xml:"status,omitempty"`
	Url            *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s ListServiceApplyResponseBodyDataListServiceReports) String() string {
	return dara.Prettify(s)
}

func (s ListServiceApplyResponseBodyDataListServiceReports) GoString() string {
	return s.String()
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetAppointmentId() *string {
	return s.AppointmentId
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetBatchGroup() *string {
	return s.BatchGroup
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetCustomerId() *string {
	return s.CustomerId
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetFileName() *string {
	return s.FileName
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetFileType() *int32 {
	return s.FileType
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetId() *int64 {
	return s.Id
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetRemarke() *string {
	return s.Remarke
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetServiceApplyId() *int64 {
	return s.ServiceApplyId
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetStatus() *int32 {
	return s.Status
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) GetUrl() *string {
	return s.Url
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetAppointmentId(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.AppointmentId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetBatchGroup(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.BatchGroup = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetCustomerId(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.CustomerId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetFileName(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.FileName = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetFileType(v int32) *ListServiceApplyResponseBodyDataListServiceReports {
	s.FileType = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetGmtCreate(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.GmtCreate = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetGmtModified(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.GmtModified = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetId(v int64) *ListServiceApplyResponseBodyDataListServiceReports {
	s.Id = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetRemarke(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.Remarke = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetServiceApplyId(v int64) *ListServiceApplyResponseBodyDataListServiceReports {
	s.ServiceApplyId = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetStatus(v int32) *ListServiceApplyResponseBodyDataListServiceReports {
	s.Status = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) SetUrl(v string) *ListServiceApplyResponseBodyDataListServiceReports {
	s.Url = &v
	return s
}

func (s *ListServiceApplyResponseBodyDataListServiceReports) Validate() error {
	return dara.Validate(s)
}
