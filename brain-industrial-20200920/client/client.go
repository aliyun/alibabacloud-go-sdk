// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivateLicenseRequest struct {
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId     *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
}

func (s ActivateLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseRequest) GoString() string {
	return s.String()
}

func (s *ActivateLicenseRequest) SetFingerprint(v string) *ActivateLicenseRequest {
	s.Fingerprint = &v
	return s
}

func (s *ActivateLicenseRequest) SetId(v string) *ActivateLicenseRequest {
	s.Id = &v
	return s
}

func (s *ActivateLicenseRequest) SetInstanceId(v string) *ActivateLicenseRequest {
	s.InstanceId = &v
	return s
}

func (s *ActivateLicenseRequest) SetOrderId(v string) *ActivateLicenseRequest {
	s.OrderId = &v
	return s
}

type ActivateLicenseResponseBody struct {
	AccessDeniedDetail *string                          `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *ActivateLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode     *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *string                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActivateLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBody) SetAccessDeniedDetail(v string) *ActivateLicenseResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetCode(v string) *ActivateLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetData(v *ActivateLicenseResponseBodyData) *ActivateLicenseResponseBody {
	s.Data = v
	return s
}

func (s *ActivateLicenseResponseBody) SetHttpStatusCode(v int32) *ActivateLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetMessage(v string) *ActivateLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetRequestId(v string) *ActivateLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateLicenseResponseBody) SetSuccess(v string) *ActivateLicenseResponseBody {
	s.Success = &v
	return s
}

type ActivateLicenseResponseBodyData struct {
	ActivateRecord        []*ActivateLicenseResponseBodyDataActivateRecord `json:"ActivateRecord,omitempty" xml:"ActivateRecord,omitempty" type:"Repeated"`
	ActivateTime          *string                                          `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	AllDuration           *string                                          `json:"AllDuration,omitempty" xml:"AllDuration,omitempty"`
	ApplicableSpecs       *string                                          `json:"ApplicableSpecs,omitempty" xml:"ApplicableSpecs,omitempty"`
	BuyTime               *string                                          `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	CpuLimit              *int32                                           `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	Description           *string                                          `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration              *string                                          `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EffectTime            *string                                          `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	ExpireTime            *string                                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Fingerprint           *string                                          `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Id                    *string                                          `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId            *string                                          `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LicenseCode           *string                                          `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	LicenseSpecName       *string                                          `json:"LicenseSpecName,omitempty" xml:"LicenseSpecName,omitempty"`
	MemoryLimit           *int32                                           `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	Status                *string                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	UnActivateAllDuration *string                                          `json:"UnActivateAllDuration,omitempty" xml:"UnActivateAllDuration,omitempty"`
}

func (s ActivateLicenseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBodyData) SetActivateRecord(v []*ActivateLicenseResponseBodyDataActivateRecord) *ActivateLicenseResponseBodyData {
	s.ActivateRecord = v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetActivateTime(v string) *ActivateLicenseResponseBodyData {
	s.ActivateTime = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetAllDuration(v string) *ActivateLicenseResponseBodyData {
	s.AllDuration = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetApplicableSpecs(v string) *ActivateLicenseResponseBodyData {
	s.ApplicableSpecs = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetBuyTime(v string) *ActivateLicenseResponseBodyData {
	s.BuyTime = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetCpuLimit(v int32) *ActivateLicenseResponseBodyData {
	s.CpuLimit = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetDescription(v string) *ActivateLicenseResponseBodyData {
	s.Description = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetDuration(v string) *ActivateLicenseResponseBodyData {
	s.Duration = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetEffectTime(v string) *ActivateLicenseResponseBodyData {
	s.EffectTime = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetExpireTime(v string) *ActivateLicenseResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetFingerprint(v string) *ActivateLicenseResponseBodyData {
	s.Fingerprint = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetId(v string) *ActivateLicenseResponseBodyData {
	s.Id = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetInstanceId(v string) *ActivateLicenseResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetLicenseCode(v string) *ActivateLicenseResponseBodyData {
	s.LicenseCode = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetLicenseSpecName(v string) *ActivateLicenseResponseBodyData {
	s.LicenseSpecName = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetMemoryLimit(v int32) *ActivateLicenseResponseBodyData {
	s.MemoryLimit = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetStatus(v string) *ActivateLicenseResponseBodyData {
	s.Status = &v
	return s
}

func (s *ActivateLicenseResponseBodyData) SetUnActivateAllDuration(v string) *ActivateLicenseResponseBodyData {
	s.UnActivateAllDuration = &v
	return s
}

type ActivateLicenseResponseBodyDataActivateRecord struct {
	ActivateTime *string `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	BuyTime      *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExpireTime   *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	LicenseCode  *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ActivateLicenseResponseBodyDataActivateRecord) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponseBodyDataActivateRecord) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) SetActivateTime(v string) *ActivateLicenseResponseBodyDataActivateRecord {
	s.ActivateTime = &v
	return s
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) SetBuyTime(v string) *ActivateLicenseResponseBodyDataActivateRecord {
	s.BuyTime = &v
	return s
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) SetDuration(v string) *ActivateLicenseResponseBodyDataActivateRecord {
	s.Duration = &v
	return s
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) SetExpireTime(v string) *ActivateLicenseResponseBodyDataActivateRecord {
	s.ExpireTime = &v
	return s
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) SetLicenseCode(v string) *ActivateLicenseResponseBodyDataActivateRecord {
	s.LicenseCode = &v
	return s
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) SetOrderId(v string) *ActivateLicenseResponseBodyDataActivateRecord {
	s.OrderId = &v
	return s
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) SetStatus(v string) *ActivateLicenseResponseBodyDataActivateRecord {
	s.Status = &v
	return s
}

type ActivateLicenseResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ActivateLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ActivateLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivateLicenseResponse) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponse) SetHeaders(v map[string]*string) *ActivateLicenseResponse {
	s.Headers = v
	return s
}

func (s *ActivateLicenseResponse) SetStatusCode(v int32) *ActivateLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *ActivateLicenseResponse) SetBody(v *ActivateLicenseResponseBody) *ActivateLicenseResponse {
	s.Body = v
	return s
}

type GetLicenseRequest struct {
	Id         *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetLicenseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseRequest) GoString() string {
	return s.String()
}

func (s *GetLicenseRequest) SetId(v int64) *GetLicenseRequest {
	s.Id = &v
	return s
}

func (s *GetLicenseRequest) SetInstanceId(v string) *GetLicenseRequest {
	s.InstanceId = &v
	return s
}

type GetLicenseResponseBody struct {
	AccessDeniedDetail *string                     `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *GetLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode     *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *string                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLicenseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetLicenseResponseBody) SetAccessDeniedDetail(v string) *GetLicenseResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLicenseResponseBody) SetCode(v string) *GetLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *GetLicenseResponseBody) SetData(v *GetLicenseResponseBodyData) *GetLicenseResponseBody {
	s.Data = v
	return s
}

func (s *GetLicenseResponseBody) SetHttpStatusCode(v int32) *GetLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLicenseResponseBody) SetMessage(v string) *GetLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *GetLicenseResponseBody) SetRequestId(v string) *GetLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLicenseResponseBody) SetSuccess(v string) *GetLicenseResponseBody {
	s.Success = &v
	return s
}

type GetLicenseResponseBodyData struct {
	ActivateRecord []*GetLicenseResponseBodyDataActivateRecord `json:"ActivateRecord,omitempty" xml:"ActivateRecord,omitempty" type:"Repeated"`
	// 代表资源一级ID的资源属性字段
	ActivateTime    *string `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	AdaptiveMachine *string `json:"AdaptiveMachine,omitempty" xml:"AdaptiveMachine,omitempty"`
	AllDuration     *string `json:"AllDuration,omitempty" xml:"AllDuration,omitempty"`
	ApplicableSpecs *string `json:"ApplicableSpecs,omitempty" xml:"ApplicableSpecs,omitempty"`
	// 代表资源名称的资源属性字段
	BuyTime     *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	CpuLimit    *int32  `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration    *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EffectTime  *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	ExpireTime  *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// 代表创建时间的资源属性字段
	Fingerprint     *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LicenseCode     *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	LicenseSpecCode *string `json:"LicenseSpecCode,omitempty" xml:"LicenseSpecCode,omitempty"`
	// 代表资源组的资源属性字段
	LicenseSpecName       *string `json:"LicenseSpecName,omitempty" xml:"LicenseSpecName,omitempty"`
	LicenseSpecType       *string `json:"LicenseSpecType,omitempty" xml:"LicenseSpecType,omitempty"`
	MemoryLimit           *int32  `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	Proposal              *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnActivateAllDuration *string `json:"UnActivateAllDuration,omitempty" xml:"UnActivateAllDuration,omitempty"`
}

func (s GetLicenseResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLicenseResponseBodyData) SetActivateRecord(v []*GetLicenseResponseBodyDataActivateRecord) *GetLicenseResponseBodyData {
	s.ActivateRecord = v
	return s
}

func (s *GetLicenseResponseBodyData) SetActivateTime(v string) *GetLicenseResponseBodyData {
	s.ActivateTime = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetAdaptiveMachine(v string) *GetLicenseResponseBodyData {
	s.AdaptiveMachine = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetAllDuration(v string) *GetLicenseResponseBodyData {
	s.AllDuration = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetApplicableSpecs(v string) *GetLicenseResponseBodyData {
	s.ApplicableSpecs = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetBuyTime(v string) *GetLicenseResponseBodyData {
	s.BuyTime = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetCpuLimit(v int32) *GetLicenseResponseBodyData {
	s.CpuLimit = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetDescription(v string) *GetLicenseResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetDuration(v string) *GetLicenseResponseBodyData {
	s.Duration = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetEffectTime(v string) *GetLicenseResponseBodyData {
	s.EffectTime = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetExpireTime(v string) *GetLicenseResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetFingerprint(v string) *GetLicenseResponseBodyData {
	s.Fingerprint = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetId(v int64) *GetLicenseResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetInstanceId(v string) *GetLicenseResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetLicenseCode(v string) *GetLicenseResponseBodyData {
	s.LicenseCode = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetLicenseSpecCode(v string) *GetLicenseResponseBodyData {
	s.LicenseSpecCode = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetLicenseSpecName(v string) *GetLicenseResponseBodyData {
	s.LicenseSpecName = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetLicenseSpecType(v string) *GetLicenseResponseBodyData {
	s.LicenseSpecType = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetMemoryLimit(v int32) *GetLicenseResponseBodyData {
	s.MemoryLimit = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetProposal(v string) *GetLicenseResponseBodyData {
	s.Proposal = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetStatus(v string) *GetLicenseResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetLicenseResponseBodyData) SetUnActivateAllDuration(v string) *GetLicenseResponseBodyData {
	s.UnActivateAllDuration = &v
	return s
}

type GetLicenseResponseBodyDataActivateRecord struct {
	ActivateTime *string `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	BuyTime      *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	Duration     *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ExpireTime   *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	LicenseCode  *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLicenseResponseBodyDataActivateRecord) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseResponseBodyDataActivateRecord) GoString() string {
	return s.String()
}

func (s *GetLicenseResponseBodyDataActivateRecord) SetActivateTime(v string) *GetLicenseResponseBodyDataActivateRecord {
	s.ActivateTime = &v
	return s
}

func (s *GetLicenseResponseBodyDataActivateRecord) SetBuyTime(v string) *GetLicenseResponseBodyDataActivateRecord {
	s.BuyTime = &v
	return s
}

func (s *GetLicenseResponseBodyDataActivateRecord) SetDuration(v string) *GetLicenseResponseBodyDataActivateRecord {
	s.Duration = &v
	return s
}

func (s *GetLicenseResponseBodyDataActivateRecord) SetExpireTime(v string) *GetLicenseResponseBodyDataActivateRecord {
	s.ExpireTime = &v
	return s
}

func (s *GetLicenseResponseBodyDataActivateRecord) SetLicenseCode(v string) *GetLicenseResponseBodyDataActivateRecord {
	s.LicenseCode = &v
	return s
}

func (s *GetLicenseResponseBodyDataActivateRecord) SetOrderId(v string) *GetLicenseResponseBodyDataActivateRecord {
	s.OrderId = &v
	return s
}

func (s *GetLicenseResponseBodyDataActivateRecord) SetStatus(v string) *GetLicenseResponseBodyDataActivateRecord {
	s.Status = &v
	return s
}

type GetLicenseResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLicenseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetLicenseResponse) SetHeaders(v map[string]*string) *GetLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetLicenseResponse) SetStatusCode(v int32) *GetLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLicenseResponse) SetBody(v *GetLicenseResponseBody) *GetLicenseResponse {
	s.Body = v
	return s
}

type ListLicensesRequest struct {
	CurrentPage *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryStr    *string `json:"QueryStr,omitempty" xml:"QueryStr,omitempty"`
}

func (s ListLicensesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLicensesRequest) GoString() string {
	return s.String()
}

func (s *ListLicensesRequest) SetCurrentPage(v int32) *ListLicensesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListLicensesRequest) SetPageSize(v int32) *ListLicensesRequest {
	s.PageSize = &v
	return s
}

func (s *ListLicensesRequest) SetQueryStr(v string) *ListLicensesRequest {
	s.QueryStr = &v
	return s
}

type ListLicensesResponseBody struct {
	AccessDeniedDetail *string                                `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	CurrentPage        *int32                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	HttpStatusCode     *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	LicenseList        []*ListLicensesResponseBodyLicenseList `json:"LicenseList,omitempty" xml:"LicenseList,omitempty" type:"Repeated"`
	Message            *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageSize           *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount         *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	TotalPage          *string                                `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
	TotalPageCount     *string                                `json:"TotalPageCount,omitempty" xml:"TotalPageCount,omitempty"`
}

func (s ListLicensesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLicensesResponseBody) GoString() string {
	return s.String()
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

type ListLicensesResponseBodyLicenseList struct {
	ActivateTime          *string `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	AdaptiveMachine       *int32  `json:"AdaptiveMachine,omitempty" xml:"AdaptiveMachine,omitempty"`
	AllDuration           *string `json:"AllDuration,omitempty" xml:"AllDuration,omitempty"`
	BuyTime               *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	CpuLimit              *int32  `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration              *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EffectTime            *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	ExpireTime            *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Id                    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LicenseSpecCode       *string `json:"LicenseSpecCode,omitempty" xml:"LicenseSpecCode,omitempty"`
	LicenseSpecName       *string `json:"LicenseSpecName,omitempty" xml:"LicenseSpecName,omitempty"`
	LicenseSpecType       *string `json:"LicenseSpecType,omitempty" xml:"LicenseSpecType,omitempty"`
	MemoryLimit           *int32  `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnActivateAllDuration *string `json:"UnActivateAllDuration,omitempty" xml:"UnActivateAllDuration,omitempty"`
}

func (s ListLicensesResponseBodyLicenseList) String() string {
	return tea.Prettify(s)
}

func (s ListLicensesResponseBodyLicenseList) GoString() string {
	return s.String()
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

type ListLicensesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLicensesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLicensesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLicensesResponse) GoString() string {
	return s.String()
}

func (s *ListLicensesResponse) SetHeaders(v map[string]*string) *ListLicensesResponse {
	s.Headers = v
	return s
}

func (s *ListLicensesResponse) SetStatusCode(v int32) *ListLicensesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLicensesResponse) SetBody(v *ListLicensesResponseBody) *ListLicensesResponse {
	s.Body = v
	return s
}

type ListUserResourcesRequest struct {
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
}

func (s ListUserResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListUserResourcesRequest) SetCommodityCode(v string) *ListUserResourcesRequest {
	s.CommodityCode = &v
	return s
}

type ListUserResourcesResponseBody struct {
	AccessDeniedDetail *string                              `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               []*ListUserResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode     *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message            *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *string                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBody) SetAccessDeniedDetail(v string) *ListUserResourcesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetCode(v string) *ListUserResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetData(v []*ListUserResourcesResponseBodyData) *ListUserResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListUserResourcesResponseBody) SetHttpStatusCode(v int32) *ListUserResourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetMessage(v string) *ListUserResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetRequestId(v string) *ListUserResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserResourcesResponseBody) SetSuccess(v string) *ListUserResourcesResponseBody {
	s.Success = &v
	return s
}

type ListUserResourcesResponseBodyData struct {
	ChargeType    *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	EndDate       *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	InstanceId    *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	Region        *string `json:"region,omitempty" xml:"region,omitempty"`
	StartDate     *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
	StatusMsg     *string `json:"statusMsg,omitempty" xml:"statusMsg,omitempty"`
}

func (s ListUserResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListUserResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponseBodyData) SetChargeType(v string) *ListUserResourcesResponseBodyData {
	s.ChargeType = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetCommodityCode(v string) *ListUserResourcesResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetEndDate(v string) *ListUserResourcesResponseBodyData {
	s.EndDate = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetInstanceId(v string) *ListUserResourcesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetRegion(v string) *ListUserResourcesResponseBodyData {
	s.Region = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetStartDate(v string) *ListUserResourcesResponseBodyData {
	s.StartDate = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetStatus(v string) *ListUserResourcesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUserResourcesResponseBodyData) SetStatusMsg(v string) *ListUserResourcesResponseBodyData {
	s.StatusMsg = &v
	return s
}

type ListUserResourcesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListUserResourcesResponse) SetHeaders(v map[string]*string) *ListUserResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListUserResourcesResponse) SetStatusCode(v int32) *ListUserResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserResourcesResponse) SetBody(v *ListUserResourcesResponseBody) *ListUserResourcesResponse {
	s.Body = v
	return s
}

type UpdateLicenseDescriptionRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// brainindustrial_aicsruntime_public_cn-mdu3ps3kw04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateLicenseDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateLicenseDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateLicenseDescriptionRequest) SetDescription(v string) *UpdateLicenseDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateLicenseDescriptionRequest) SetInstanceId(v string) *UpdateLicenseDescriptionRequest {
	s.InstanceId = &v
	return s
}

type UpdateLicenseDescriptionResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 68738E75-43C1-5AE5-9F3A-AFEF576D7B5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLicenseDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateLicenseDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLicenseDescriptionResponseBody) SetAccessDeniedDetail(v string) *UpdateLicenseDescriptionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetCode(v string) *UpdateLicenseDescriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetHttpStatusCode(v int32) *UpdateLicenseDescriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetMessage(v string) *UpdateLicenseDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetRequestId(v string) *UpdateLicenseDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetSuccess(v string) *UpdateLicenseDescriptionResponseBody {
	s.Success = &v
	return s
}

type UpdateLicenseDescriptionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLicenseDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLicenseDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateLicenseDescriptionResponse) GoString() string {
	return s.String()
}

func (s *UpdateLicenseDescriptionResponse) SetHeaders(v map[string]*string) *UpdateLicenseDescriptionResponse {
	s.Headers = v
	return s
}

func (s *UpdateLicenseDescriptionResponse) SetStatusCode(v int32) *UpdateLicenseDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLicenseDescriptionResponse) SetBody(v *UpdateLicenseDescriptionResponseBody) *UpdateLicenseDescriptionResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("brain-industrial"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 激活License
//
// @param request - ActivateLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicenseWithOptions(request *ActivateLicenseRequest, runtime *util.RuntimeOptions) (_result *ActivateLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Fingerprint)) {
		body["Fingerprint"] = request.Fingerprint
	}

	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderId)) {
		body["OrderId"] = request.OrderId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ActivateLicense"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 激活License
//
// @param request - ActivateLicenseRequest
//
// @return ActivateLicenseResponse
func (client *Client) ActivateLicense(request *ActivateLicenseRequest) (_result *ActivateLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivateLicenseResponse{}
	_body, _err := client.ActivateLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// License详情
//
// @param request - GetLicenseRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLicenseResponse
func (client *Client) GetLicenseWithOptions(request *GetLicenseRequest, runtime *util.RuntimeOptions) (_result *GetLicenseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Id)) {
		body["Id"] = request.Id
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLicense"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLicenseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// License详情
//
// @param request - GetLicenseRequest
//
// @return GetLicenseResponse
func (client *Client) GetLicense(request *GetLicenseRequest) (_result *GetLicenseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLicenseResponse{}
	_body, _err := client.GetLicenseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// License列表
//
// @param request - ListLicensesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListLicensesResponse
func (client *Client) ListLicensesWithOptions(request *ListLicensesRequest, runtime *util.RuntimeOptions) (_result *ListLicensesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryStr)) {
		body["QueryStr"] = request.QueryStr
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLicenses"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListLicensesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// License列表
//
// @param request - ListLicensesRequest
//
// @return ListLicensesResponse
func (client *Client) ListLicenses(request *ListLicensesRequest) (_result *ListLicensesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListLicensesResponse{}
	_body, _err := client.ListLicensesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取用户资源列表
//
// @param request - ListUserResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserResourcesResponse
func (client *Client) ListUserResourcesWithOptions(request *ListUserResourcesRequest, runtime *util.RuntimeOptions) (_result *ListUserResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CommodityCode)) {
		body["CommodityCode"] = request.CommodityCode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserResources"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取用户资源列表
//
// @param request - ListUserResourcesRequest
//
// @return ListUserResourcesResponse
func (client *Client) ListUserResources(request *ListUserResourcesRequest) (_result *ListUserResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserResourcesResponse{}
	_body, _err := client.ListUserResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新license描述
//
// @param request - UpdateLicenseDescriptionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateLicenseDescriptionResponse
func (client *Client) UpdateLicenseDescriptionWithOptions(request *UpdateLicenseDescriptionRequest, runtime *util.RuntimeOptions) (_result *UpdateLicenseDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateLicenseDescription"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateLicenseDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新license描述
//
// @param request - UpdateLicenseDescriptionRequest
//
// @return UpdateLicenseDescriptionResponse
func (client *Client) UpdateLicenseDescription(request *UpdateLicenseDescriptionRequest) (_result *UpdateLicenseDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateLicenseDescriptionResponse{}
	_body, _err := client.UpdateLicenseDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
