// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLicensesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListLicensesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListLicensesResponseBody
	GetCode() *string
	SetCurrentPage(v int32) *ListLicensesResponseBody
	GetCurrentPage() *int32
	SetHttpStatusCode(v int32) *ListLicensesResponseBody
	GetHttpStatusCode() *int32
	SetLicenseList(v []*ListLicensesResponseBodyLicenseList) *ListLicensesResponseBody
	GetLicenseList() []*ListLicensesResponseBodyLicenseList
	SetMessage(v string) *ListLicensesResponseBody
	GetMessage() *string
	SetPageSize(v int32) *ListLicensesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLicensesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLicensesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListLicensesResponseBody
	GetTotalCount() *int64
	SetTotalPage(v string) *ListLicensesResponseBody
	GetTotalPage() *string
	SetTotalPageCount(v string) *ListLicensesResponseBody
	GetTotalPageCount() *string
}

type ListLicensesResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	LicenseList    []*ListLicensesResponseBodyLicenseList `json:"LicenseList,omitempty" xml:"LicenseList,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 12
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// FAC892CE-5A94-5616-91DC-629B09CC6792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 42
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 4
	TotalPage *string `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	// example:
	//
	// 4
	TotalPageCount *string `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListLicensesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLicensesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLicensesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListLicensesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLicensesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListLicensesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListLicensesResponseBody) GetLicenseList() []*ListLicensesResponseBodyLicenseList {
	return s.LicenseList
}

func (s *ListLicensesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLicensesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLicensesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLicensesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLicensesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLicensesResponseBody) GetTotalPage() *string {
	return s.TotalPage
}

func (s *ListLicensesResponseBody) GetTotalPageCount() *string {
	return s.TotalPageCount
}

func (s *ListLicensesResponseBody) SetAccessDeniedDetail(v string) *ListLicensesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListLicensesResponseBody) SetCode(v string) *ListLicensesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLicensesResponseBody) SetCurrentPage(v int32) *ListLicensesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListLicensesResponseBody) SetHttpStatusCode(v int32) *ListLicensesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLicensesResponseBody) SetLicenseList(v []*ListLicensesResponseBodyLicenseList) *ListLicensesResponseBody {
	s.LicenseList = v
	return s
}

func (s *ListLicensesResponseBody) SetMessage(v string) *ListLicensesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLicensesResponseBody) SetPageSize(v int32) *ListLicensesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLicensesResponseBody) SetRequestId(v string) *ListLicensesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLicensesResponseBody) SetSuccess(v bool) *ListLicensesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLicensesResponseBody) SetTotalCount(v int64) *ListLicensesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLicensesResponseBody) SetTotalPage(v string) *ListLicensesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListLicensesResponseBody) SetTotalPageCount(v string) *ListLicensesResponseBody {
	s.TotalPageCount = &v
	return s
}

func (s *ListLicensesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLicensesResponseBodyLicenseList struct {
	// example:
	//
	// 2024-10-14 14:15:45
	ActivateTime    *string `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	AdaptiveMachine *int32  `json:"AdaptiveMachine,omitempty" xml:"AdaptiveMachine,omitempty"`
	AllDuration     *string `json:"AllDuration,omitempty" xml:"AllDuration,omitempty"`
	// example:
	//
	// 2024-10-14 13:17:20
	BuyTime *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	// example:
	//
	// 72
	CpuLimit *int32 `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// example:
	//
	// ""
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2024-10-14 14:15:45
	EffectTime *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	// example:
	//
	// 2025-10-14 14:15:45
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 451f14ee73604aesdfe4da606764ce09
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// aics_1000
	LicenseSpecCode *string `json:"LicenseSpecCode,omitempty" xml:"LicenseSpecCode,omitempty"`
	LicenseSpecName *string `json:"LicenseSpecName,omitempty" xml:"LicenseSpecName,omitempty"`
	// example:
	//
	// brainindustrial_aicsruntime_public_cn
	LicenseSpecType *string `json:"LicenseSpecType,omitempty" xml:"LicenseSpecType,omitempty"`
	// example:
	//
	// 256
	MemoryLimit *int32 `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// example:
	//
	// activated
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnActivateAllDuration *string `json:"UnActivateAllDuration,omitempty" xml:"UnActivateAllDuration,omitempty"`
}

func (s ListLicensesResponseBodyLicenseList) String() string {
	return dara.Prettify(s)
}

func (s ListLicensesResponseBodyLicenseList) GoString() string {
	return s.String()
}

func (s *ListLicensesResponseBodyLicenseList) GetActivateTime() *string {
	return s.ActivateTime
}

func (s *ListLicensesResponseBodyLicenseList) GetAdaptiveMachine() *int32 {
	return s.AdaptiveMachine
}

func (s *ListLicensesResponseBodyLicenseList) GetAllDuration() *string {
	return s.AllDuration
}

func (s *ListLicensesResponseBodyLicenseList) GetBuyTime() *string {
	return s.BuyTime
}

func (s *ListLicensesResponseBodyLicenseList) GetCpuLimit() *int32 {
	return s.CpuLimit
}

func (s *ListLicensesResponseBodyLicenseList) GetDescription() *string {
	return s.Description
}

func (s *ListLicensesResponseBodyLicenseList) GetDuration() *string {
	return s.Duration
}

func (s *ListLicensesResponseBodyLicenseList) GetEffectTime() *string {
	return s.EffectTime
}

func (s *ListLicensesResponseBodyLicenseList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListLicensesResponseBodyLicenseList) GetId() *string {
	return s.Id
}

func (s *ListLicensesResponseBodyLicenseList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListLicensesResponseBodyLicenseList) GetLicenseSpecCode() *string {
	return s.LicenseSpecCode
}

func (s *ListLicensesResponseBodyLicenseList) GetLicenseSpecName() *string {
	return s.LicenseSpecName
}

func (s *ListLicensesResponseBodyLicenseList) GetLicenseSpecType() *string {
	return s.LicenseSpecType
}

func (s *ListLicensesResponseBodyLicenseList) GetMemoryLimit() *int32 {
	return s.MemoryLimit
}

func (s *ListLicensesResponseBodyLicenseList) GetStatus() *string {
	return s.Status
}

func (s *ListLicensesResponseBodyLicenseList) GetUnActivateAllDuration() *string {
	return s.UnActivateAllDuration
}

func (s *ListLicensesResponseBodyLicenseList) SetActivateTime(v string) *ListLicensesResponseBodyLicenseList {
	s.ActivateTime = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetAdaptiveMachine(v int32) *ListLicensesResponseBodyLicenseList {
	s.AdaptiveMachine = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetAllDuration(v string) *ListLicensesResponseBodyLicenseList {
	s.AllDuration = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetBuyTime(v string) *ListLicensesResponseBodyLicenseList {
	s.BuyTime = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetCpuLimit(v int32) *ListLicensesResponseBodyLicenseList {
	s.CpuLimit = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetDescription(v string) *ListLicensesResponseBodyLicenseList {
	s.Description = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetDuration(v string) *ListLicensesResponseBodyLicenseList {
	s.Duration = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetEffectTime(v string) *ListLicensesResponseBodyLicenseList {
	s.EffectTime = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetExpireTime(v string) *ListLicensesResponseBodyLicenseList {
	s.ExpireTime = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetId(v string) *ListLicensesResponseBodyLicenseList {
	s.Id = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetInstanceId(v string) *ListLicensesResponseBodyLicenseList {
	s.InstanceId = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetLicenseSpecCode(v string) *ListLicensesResponseBodyLicenseList {
	s.LicenseSpecCode = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetLicenseSpecName(v string) *ListLicensesResponseBodyLicenseList {
	s.LicenseSpecName = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetLicenseSpecType(v string) *ListLicensesResponseBodyLicenseList {
	s.LicenseSpecType = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetMemoryLimit(v int32) *ListLicensesResponseBodyLicenseList {
	s.MemoryLimit = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetStatus(v string) *ListLicensesResponseBodyLicenseList {
	s.Status = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) SetUnActivateAllDuration(v string) *ListLicensesResponseBodyLicenseList {
	s.UnActivateAllDuration = &v
	return s
}

func (s *ListLicensesResponseBodyLicenseList) Validate() error {
	return dara.Validate(s)
}
