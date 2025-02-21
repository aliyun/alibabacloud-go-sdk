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

type CreateEssOptJobRequest struct {
	// example:
	//
	// 1
	Duration  *int32                             `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ElecPrice []*CreateEssOptJobRequestElecPrice `json:"ElecPrice,omitempty" xml:"ElecPrice,omitempty" type:"Repeated"`
	// example:
	//
	// FIFTEEN_MIN
	Freq     *string                           `json:"Freq,omitempty" xml:"Freq,omitempty"`
	GenPrice []*CreateEssOptJobRequestGenPrice `json:"GenPrice,omitempty" xml:"GenPrice,omitempty" type:"Repeated"`
	Location *CreateEssOptJobRequestLocation   `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-02-12
	RunDate    *string                             `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	SystemData []*CreateEssOptJobRequestSystemData `json:"SystemData,omitempty" xml:"SystemData,omitempty" type:"Repeated"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// example:
	//
	// LOAD_ESS_SOLAR
	TopoType *string `json:"TopoType,omitempty" xml:"TopoType,omitempty"`
}

func (s CreateEssOptJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobRequest) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequest) SetDuration(v int32) *CreateEssOptJobRequest {
	s.Duration = &v
	return s
}

func (s *CreateEssOptJobRequest) SetElecPrice(v []*CreateEssOptJobRequestElecPrice) *CreateEssOptJobRequest {
	s.ElecPrice = v
	return s
}

func (s *CreateEssOptJobRequest) SetFreq(v string) *CreateEssOptJobRequest {
	s.Freq = &v
	return s
}

func (s *CreateEssOptJobRequest) SetGenPrice(v []*CreateEssOptJobRequestGenPrice) *CreateEssOptJobRequest {
	s.GenPrice = v
	return s
}

func (s *CreateEssOptJobRequest) SetLocation(v *CreateEssOptJobRequestLocation) *CreateEssOptJobRequest {
	s.Location = v
	return s
}

func (s *CreateEssOptJobRequest) SetModelVersion(v string) *CreateEssOptJobRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateEssOptJobRequest) SetRunDate(v string) *CreateEssOptJobRequest {
	s.RunDate = &v
	return s
}

func (s *CreateEssOptJobRequest) SetSystemData(v []*CreateEssOptJobRequestSystemData) *CreateEssOptJobRequest {
	s.SystemData = v
	return s
}

func (s *CreateEssOptJobRequest) SetTimeZone(v string) *CreateEssOptJobRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateEssOptJobRequest) SetTopoType(v string) *CreateEssOptJobRequest {
	s.TopoType = &v
	return s
}

type CreateEssOptJobRequestElecPrice struct {
	// example:
	//
	// 00:00:15
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// example:
	//
	// 0.5
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s CreateEssOptJobRequestElecPrice) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobRequestElecPrice) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequestElecPrice) SetDataTime(v string) *CreateEssOptJobRequestElecPrice {
	s.DataTime = &v
	return s
}

func (s *CreateEssOptJobRequestElecPrice) SetPrice(v string) *CreateEssOptJobRequestElecPrice {
	s.Price = &v
	return s
}

type CreateEssOptJobRequestGenPrice struct {
	// example:
	//
	// 00:00:15
	DataTime *string `json:"DataTime,omitempty" xml:"DataTime,omitempty"`
	// example:
	//
	// 0.3
	Price *string `json:"Price,omitempty" xml:"Price,omitempty"`
}

func (s CreateEssOptJobRequestGenPrice) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobRequestGenPrice) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequestGenPrice) SetDataTime(v string) *CreateEssOptJobRequestGenPrice {
	s.DataTime = &v
	return s
}

func (s *CreateEssOptJobRequestGenPrice) SetPrice(v string) *CreateEssOptJobRequestGenPrice {
	s.Price = &v
	return s
}

type CreateEssOptJobRequestLocation struct {
	// example:
	//
	// 10.123
	Altitude *float64 `json:"Altitude,omitempty" xml:"Altitude,omitempty"`
	// example:
	//
	// 40.027
	Latitude *float64 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// 120.042
	Longitude *float64 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
}

func (s CreateEssOptJobRequestLocation) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobRequestLocation) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequestLocation) SetAltitude(v float64) *CreateEssOptJobRequestLocation {
	s.Altitude = &v
	return s
}

func (s *CreateEssOptJobRequestLocation) SetLatitude(v float64) *CreateEssOptJobRequestLocation {
	s.Latitude = &v
	return s
}

func (s *CreateEssOptJobRequestLocation) SetLongitude(v float64) *CreateEssOptJobRequestLocation {
	s.Longitude = &v
	return s
}

type CreateEssOptJobRequestSystemData struct {
	HistoryData []map[string]interface{} `json:"HistoryData,omitempty" xml:"HistoryData,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SystemId     *string                `json:"SystemId,omitempty" xml:"SystemId,omitempty"`
	SystemParams map[string]interface{} `json:"SystemParams,omitempty" xml:"SystemParams,omitempty"`
	// example:
	//
	// ess
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
}

func (s CreateEssOptJobRequestSystemData) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobRequestSystemData) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobRequestSystemData) SetHistoryData(v []map[string]interface{}) *CreateEssOptJobRequestSystemData {
	s.HistoryData = v
	return s
}

func (s *CreateEssOptJobRequestSystemData) SetSystemId(v string) *CreateEssOptJobRequestSystemData {
	s.SystemId = &v
	return s
}

func (s *CreateEssOptJobRequestSystemData) SetSystemParams(v map[string]interface{}) *CreateEssOptJobRequestSystemData {
	s.SystemParams = v
	return s
}

func (s *CreateEssOptJobRequestSystemData) SetSystemType(v string) *CreateEssOptJobRequestSystemData {
	s.SystemType = &v
	return s
}

type CreateEssOptJobShrinkRequest struct {
	// example:
	//
	// 1
	Duration        *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ElecPriceShrink *string `json:"ElecPrice,omitempty" xml:"ElecPrice,omitempty"`
	// example:
	//
	// FIFTEEN_MIN
	Freq           *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	GenPriceShrink *string `json:"GenPrice,omitempty" xml:"GenPrice,omitempty"`
	LocationShrink *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-02-12
	RunDate          *string `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	SystemDataShrink *string `json:"SystemData,omitempty" xml:"SystemData,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// example:
	//
	// LOAD_ESS_SOLAR
	TopoType *string `json:"TopoType,omitempty" xml:"TopoType,omitempty"`
}

func (s CreateEssOptJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobShrinkRequest) SetDuration(v int32) *CreateEssOptJobShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetElecPriceShrink(v string) *CreateEssOptJobShrinkRequest {
	s.ElecPriceShrink = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetFreq(v string) *CreateEssOptJobShrinkRequest {
	s.Freq = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetGenPriceShrink(v string) *CreateEssOptJobShrinkRequest {
	s.GenPriceShrink = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetLocationShrink(v string) *CreateEssOptJobShrinkRequest {
	s.LocationShrink = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetModelVersion(v string) *CreateEssOptJobShrinkRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetRunDate(v string) *CreateEssOptJobShrinkRequest {
	s.RunDate = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetSystemDataShrink(v string) *CreateEssOptJobShrinkRequest {
	s.SystemDataShrink = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetTimeZone(v string) *CreateEssOptJobShrinkRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateEssOptJobShrinkRequest) SetTopoType(v string) *CreateEssOptJobShrinkRequest {
	s.TopoType = &v
	return s
}

type CreateEssOptJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateEssOptJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateEssOptJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobResponseBody) SetCode(v string) *CreateEssOptJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEssOptJobResponseBody) SetData(v *CreateEssOptJobResponseBodyData) *CreateEssOptJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateEssOptJobResponseBody) SetMessage(v string) *CreateEssOptJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEssOptJobResponseBody) SetRequestId(v string) *CreateEssOptJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEssOptJobResponseBody) SetSuccess(v string) *CreateEssOptJobResponseBody {
	s.Success = &v
	return s
}

type CreateEssOptJobResponseBodyData struct {
	// example:
	//
	// True
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 2024-12-22 00:00:21
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ""
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8c0ca18a-246a-4acd-80ca-e16d8ff5ef33
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 25
	Progress *int32                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *CreateEssOptJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateEssOptJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobResponseBodyData) SetCompleted(v bool) *CreateEssOptJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetCreateTime(v string) *CreateEssOptJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetError(v string) *CreateEssOptJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetJobId(v string) *CreateEssOptJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetProgress(v int32) *CreateEssOptJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetResponse(v *CreateEssOptJobResponseBodyDataResponse) *CreateEssOptJobResponseBodyData {
	s.Response = v
	return s
}

func (s *CreateEssOptJobResponseBodyData) SetStatus(v string) *CreateEssOptJobResponseBodyData {
	s.Status = &v
	return s
}

type CreateEssOptJobResponseBodyDataResponse struct {
	// example:
	//
	// {}
	DebugInfo interface{} `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	// example:
	//
	// LoadForecast
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// {}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateEssOptJobResponseBodyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *CreateEssOptJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *CreateEssOptJobResponseBodyDataResponse) SetJobType(v string) *CreateEssOptJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *CreateEssOptJobResponseBodyDataResponse) SetResult(v interface{}) *CreateEssOptJobResponseBodyDataResponse {
	s.Result = v
	return s
}

type CreateEssOptJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEssOptJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEssOptJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateEssOptJobResponse) GoString() string {
	return s.String()
}

func (s *CreateEssOptJobResponse) SetHeaders(v map[string]*string) *CreateEssOptJobResponse {
	s.Headers = v
	return s
}

func (s *CreateEssOptJobResponse) SetStatusCode(v int32) *CreateEssOptJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEssOptJobResponse) SetBody(v *CreateEssOptJobResponseBody) *CreateEssOptJobResponse {
	s.Body = v
	return s
}

type CreateLoadForecastJobRequest struct {
	// example:
	//
	// electricityMeter
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// FIFTEEN_MIN
	Freq        *string                                    `json:"Freq,omitempty" xml:"Freq,omitempty"`
	HistoryData []*CreateLoadForecastJobRequestHistoryData `json:"HistoryData,omitempty" xml:"HistoryData,omitempty" type:"Repeated"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-12-12
	RunDate *string `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	// example:
	//
	// load
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CreateLoadForecastJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadForecastJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobRequest) SetDeviceType(v string) *CreateLoadForecastJobRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetDuration(v int32) *CreateLoadForecastJobRequest {
	s.Duration = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetFreq(v string) *CreateLoadForecastJobRequest {
	s.Freq = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetHistoryData(v []*CreateLoadForecastJobRequestHistoryData) *CreateLoadForecastJobRequest {
	s.HistoryData = v
	return s
}

func (s *CreateLoadForecastJobRequest) SetModelVersion(v string) *CreateLoadForecastJobRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetRunDate(v string) *CreateLoadForecastJobRequest {
	s.RunDate = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetSystemType(v string) *CreateLoadForecastJobRequest {
	s.SystemType = &v
	return s
}

func (s *CreateLoadForecastJobRequest) SetTimeZone(v string) *CreateLoadForecastJobRequest {
	s.TimeZone = &v
	return s
}

type CreateLoadForecastJobRequestHistoryData struct {
	// example:
	//
	// 2025-12-12 00:00:00
	RunTime *string `json:"RunTime,omitempty" xml:"RunTime,omitempty"`
	// example:
	//
	// 1.0
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLoadForecastJobRequestHistoryData) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadForecastJobRequestHistoryData) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobRequestHistoryData) SetRunTime(v string) *CreateLoadForecastJobRequestHistoryData {
	s.RunTime = &v
	return s
}

func (s *CreateLoadForecastJobRequestHistoryData) SetValue(v float64) *CreateLoadForecastJobRequestHistoryData {
	s.Value = &v
	return s
}

type CreateLoadForecastJobShrinkRequest struct {
	// example:
	//
	// electricityMeter
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// FIFTEEN_MIN
	Freq              *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	HistoryDataShrink *string `json:"HistoryData,omitempty" xml:"HistoryData,omitempty"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-12-12
	RunDate *string `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	// example:
	//
	// load
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CreateLoadForecastJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadForecastJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobShrinkRequest) SetDeviceType(v string) *CreateLoadForecastJobShrinkRequest {
	s.DeviceType = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetDuration(v int32) *CreateLoadForecastJobShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetFreq(v string) *CreateLoadForecastJobShrinkRequest {
	s.Freq = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetHistoryDataShrink(v string) *CreateLoadForecastJobShrinkRequest {
	s.HistoryDataShrink = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetModelVersion(v string) *CreateLoadForecastJobShrinkRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetRunDate(v string) *CreateLoadForecastJobShrinkRequest {
	s.RunDate = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetSystemType(v string) *CreateLoadForecastJobShrinkRequest {
	s.SystemType = &v
	return s
}

func (s *CreateLoadForecastJobShrinkRequest) SetTimeZone(v string) *CreateLoadForecastJobShrinkRequest {
	s.TimeZone = &v
	return s
}

type CreateLoadForecastJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateLoadForecastJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreateLoadForecastJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadForecastJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobResponseBody) SetCode(v string) *CreateLoadForecastJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLoadForecastJobResponseBody) SetData(v *CreateLoadForecastJobResponseBodyData) *CreateLoadForecastJobResponseBody {
	s.Data = v
	return s
}

func (s *CreateLoadForecastJobResponseBody) SetMessage(v string) *CreateLoadForecastJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLoadForecastJobResponseBody) SetRequestId(v string) *CreateLoadForecastJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLoadForecastJobResponseBody) SetSuccess(v string) *CreateLoadForecastJobResponseBody {
	s.Success = &v
	return s
}

type CreateLoadForecastJobResponseBodyData struct {
	// example:
	//
	// True
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 2024-12-22 00:00:21
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ""
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8c0ca18a-246a-4acd-80ca-e16d8ff5ef33
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 25
	Progress *int32                                         `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *CreateLoadForecastJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateLoadForecastJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadForecastJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobResponseBodyData) SetCompleted(v bool) *CreateLoadForecastJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetCreateTime(v string) *CreateLoadForecastJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetError(v string) *CreateLoadForecastJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetJobId(v string) *CreateLoadForecastJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetProgress(v int32) *CreateLoadForecastJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetResponse(v *CreateLoadForecastJobResponseBodyDataResponse) *CreateLoadForecastJobResponseBodyData {
	s.Response = v
	return s
}

func (s *CreateLoadForecastJobResponseBodyData) SetStatus(v string) *CreateLoadForecastJobResponseBodyData {
	s.Status = &v
	return s
}

type CreateLoadForecastJobResponseBodyDataResponse struct {
	// example:
	//
	// {}
	DebugInfo interface{} `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	// example:
	//
	// PowerForecast
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// {}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreateLoadForecastJobResponseBodyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadForecastJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *CreateLoadForecastJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) SetJobType(v string) *CreateLoadForecastJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *CreateLoadForecastJobResponseBodyDataResponse) SetResult(v interface{}) *CreateLoadForecastJobResponseBodyDataResponse {
	s.Result = v
	return s
}

type CreateLoadForecastJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLoadForecastJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLoadForecastJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLoadForecastJobResponse) GoString() string {
	return s.String()
}

func (s *CreateLoadForecastJobResponse) SetHeaders(v map[string]*string) *CreateLoadForecastJobResponse {
	s.Headers = v
	return s
}

func (s *CreateLoadForecastJobResponse) SetStatusCode(v int32) *CreateLoadForecastJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLoadForecastJobResponse) SetBody(v *CreateLoadForecastJobResponseBody) *CreateLoadForecastJobResponse {
	s.Body = v
	return s
}

type CreatePowerForecastJobRequest struct {
	// example:
	//
	// electricityMeter
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// FIFTEEN_MIN
	Freq        *string                                     `json:"Freq,omitempty" xml:"Freq,omitempty"`
	HistoryData []*CreatePowerForecastJobRequestHistoryData `json:"HistoryData,omitempty" xml:"HistoryData,omitempty" type:"Repeated"`
	Location    *CreatePowerForecastJobRequestLocation      `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-02-12
	RunDate *string `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	// example:
	//
	// load
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CreatePowerForecastJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePowerForecastJobRequest) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobRequest) SetDeviceType(v string) *CreatePowerForecastJobRequest {
	s.DeviceType = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetDuration(v int32) *CreatePowerForecastJobRequest {
	s.Duration = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetFreq(v string) *CreatePowerForecastJobRequest {
	s.Freq = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetHistoryData(v []*CreatePowerForecastJobRequestHistoryData) *CreatePowerForecastJobRequest {
	s.HistoryData = v
	return s
}

func (s *CreatePowerForecastJobRequest) SetLocation(v *CreatePowerForecastJobRequestLocation) *CreatePowerForecastJobRequest {
	s.Location = v
	return s
}

func (s *CreatePowerForecastJobRequest) SetModelVersion(v string) *CreatePowerForecastJobRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetRunDate(v string) *CreatePowerForecastJobRequest {
	s.RunDate = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetSystemType(v string) *CreatePowerForecastJobRequest {
	s.SystemType = &v
	return s
}

func (s *CreatePowerForecastJobRequest) SetTimeZone(v string) *CreatePowerForecastJobRequest {
	s.TimeZone = &v
	return s
}

type CreatePowerForecastJobRequestHistoryData struct {
	// example:
	//
	// 2025-02-12 00:00:00
	RunTime *string `json:"RunTime,omitempty" xml:"RunTime,omitempty"`
	// example:
	//
	// 1.0
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePowerForecastJobRequestHistoryData) String() string {
	return tea.Prettify(s)
}

func (s CreatePowerForecastJobRequestHistoryData) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobRequestHistoryData) SetRunTime(v string) *CreatePowerForecastJobRequestHistoryData {
	s.RunTime = &v
	return s
}

func (s *CreatePowerForecastJobRequestHistoryData) SetValue(v float64) *CreatePowerForecastJobRequestHistoryData {
	s.Value = &v
	return s
}

type CreatePowerForecastJobRequestLocation struct {
	// example:
	//
	// 10.123
	Altitude *float64 `json:"Altitude,omitempty" xml:"Altitude,omitempty"`
	// example:
	//
	// 40.027
	Latitude *float64 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	// example:
	//
	// 120.042
	Longitude *float64 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
}

func (s CreatePowerForecastJobRequestLocation) String() string {
	return tea.Prettify(s)
}

func (s CreatePowerForecastJobRequestLocation) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobRequestLocation) SetAltitude(v float64) *CreatePowerForecastJobRequestLocation {
	s.Altitude = &v
	return s
}

func (s *CreatePowerForecastJobRequestLocation) SetLatitude(v float64) *CreatePowerForecastJobRequestLocation {
	s.Latitude = &v
	return s
}

func (s *CreatePowerForecastJobRequestLocation) SetLongitude(v float64) *CreatePowerForecastJobRequestLocation {
	s.Longitude = &v
	return s
}

type CreatePowerForecastJobShrinkRequest struct {
	// example:
	//
	// electricityMeter
	DeviceType *string `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// FIFTEEN_MIN
	Freq              *string `json:"Freq,omitempty" xml:"Freq,omitempty"`
	HistoryDataShrink *string `json:"HistoryData,omitempty" xml:"HistoryData,omitempty"`
	LocationShrink    *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// latest
	ModelVersion *string `json:"ModelVersion,omitempty" xml:"ModelVersion,omitempty"`
	// example:
	//
	// 2025-02-12
	RunDate *string `json:"RunDate,omitempty" xml:"RunDate,omitempty"`
	// example:
	//
	// load
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s CreatePowerForecastJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePowerForecastJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobShrinkRequest) SetDeviceType(v string) *CreatePowerForecastJobShrinkRequest {
	s.DeviceType = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetDuration(v int32) *CreatePowerForecastJobShrinkRequest {
	s.Duration = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetFreq(v string) *CreatePowerForecastJobShrinkRequest {
	s.Freq = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetHistoryDataShrink(v string) *CreatePowerForecastJobShrinkRequest {
	s.HistoryDataShrink = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetLocationShrink(v string) *CreatePowerForecastJobShrinkRequest {
	s.LocationShrink = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetModelVersion(v string) *CreatePowerForecastJobShrinkRequest {
	s.ModelVersion = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetRunDate(v string) *CreatePowerForecastJobShrinkRequest {
	s.RunDate = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetSystemType(v string) *CreatePowerForecastJobShrinkRequest {
	s.SystemType = &v
	return s
}

func (s *CreatePowerForecastJobShrinkRequest) SetTimeZone(v string) *CreatePowerForecastJobShrinkRequest {
	s.TimeZone = &v
	return s
}

type CreatePowerForecastJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreatePowerForecastJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s CreatePowerForecastJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePowerForecastJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobResponseBody) SetCode(v string) *CreatePowerForecastJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePowerForecastJobResponseBody) SetData(v *CreatePowerForecastJobResponseBodyData) *CreatePowerForecastJobResponseBody {
	s.Data = v
	return s
}

func (s *CreatePowerForecastJobResponseBody) SetMessage(v string) *CreatePowerForecastJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePowerForecastJobResponseBody) SetRequestId(v string) *CreatePowerForecastJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePowerForecastJobResponseBody) SetSuccess(v string) *CreatePowerForecastJobResponseBody {
	s.Success = &v
	return s
}

type CreatePowerForecastJobResponseBodyData struct {
	// example:
	//
	// True
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 2024-12-22 00:00:21
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ""
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 8c0ca18a-246a-4acd-80ca-e16d8ff5ef33
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 25
	Progress *int32                                          `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *CreatePowerForecastJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreatePowerForecastJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreatePowerForecastJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobResponseBodyData) SetCompleted(v bool) *CreatePowerForecastJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetCreateTime(v string) *CreatePowerForecastJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetError(v string) *CreatePowerForecastJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetJobId(v string) *CreatePowerForecastJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetProgress(v int32) *CreatePowerForecastJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetResponse(v *CreatePowerForecastJobResponseBodyDataResponse) *CreatePowerForecastJobResponseBodyData {
	s.Response = v
	return s
}

func (s *CreatePowerForecastJobResponseBodyData) SetStatus(v string) *CreatePowerForecastJobResponseBodyData {
	s.Status = &v
	return s
}

type CreatePowerForecastJobResponseBodyDataResponse struct {
	// example:
	//
	// {}
	DebugInfo interface{} `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	// example:
	//
	// LoadForecast
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// {}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CreatePowerForecastJobResponseBodyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePowerForecastJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *CreatePowerForecastJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) SetJobType(v string) *CreatePowerForecastJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *CreatePowerForecastJobResponseBodyDataResponse) SetResult(v interface{}) *CreatePowerForecastJobResponseBodyDataResponse {
	s.Result = v
	return s
}

type CreatePowerForecastJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePowerForecastJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePowerForecastJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePowerForecastJobResponse) GoString() string {
	return s.String()
}

func (s *CreatePowerForecastJobResponse) SetHeaders(v map[string]*string) *CreatePowerForecastJobResponse {
	s.Headers = v
	return s
}

func (s *CreatePowerForecastJobResponse) SetStatusCode(v int32) *CreatePowerForecastJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePowerForecastJobResponse) SetBody(v *CreatePowerForecastJobResponseBody) *CreatePowerForecastJobResponse {
	s.Body = v
	return s
}

type GetAivppAlgoJobRequest struct {
	// example:
	//
	// 5854bfa6-f002-43c2-8e1d-e9b2c28f9384
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAivppAlgoJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAivppAlgoJobRequest) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobRequest) SetJobId(v string) *GetAivppAlgoJobRequest {
	s.JobId = &v
	return s
}

type GetAivppAlgoJobResponseBody struct {
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAivppAlgoJobResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetAivppAlgoJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAivppAlgoJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobResponseBody) SetCode(v string) *GetAivppAlgoJobResponseBody {
	s.Code = &v
	return s
}

func (s *GetAivppAlgoJobResponseBody) SetData(v *GetAivppAlgoJobResponseBodyData) *GetAivppAlgoJobResponseBody {
	s.Data = v
	return s
}

func (s *GetAivppAlgoJobResponseBody) SetMessage(v string) *GetAivppAlgoJobResponseBody {
	s.Message = &v
	return s
}

func (s *GetAivppAlgoJobResponseBody) SetRequestId(v string) *GetAivppAlgoJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAivppAlgoJobResponseBody) SetSuccess(v string) *GetAivppAlgoJobResponseBody {
	s.Success = &v
	return s
}

type GetAivppAlgoJobResponseBodyData struct {
	// example:
	//
	// False
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// example:
	//
	// 2024-12-10 17:50:48
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// ""
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 5854bfa6-f002-43c2-8e1d-e9b2c28f9384
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 100
	Progress *int32                                   `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Response *GetAivppAlgoJobResponseBodyDataResponse `json:"Response,omitempty" xml:"Response,omitempty" type:"Struct"`
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetAivppAlgoJobResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetAivppAlgoJobResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobResponseBodyData) SetCompleted(v bool) *GetAivppAlgoJobResponseBodyData {
	s.Completed = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetCreateTime(v string) *GetAivppAlgoJobResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetError(v string) *GetAivppAlgoJobResponseBodyData {
	s.Error = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetJobId(v string) *GetAivppAlgoJobResponseBodyData {
	s.JobId = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetProgress(v int32) *GetAivppAlgoJobResponseBodyData {
	s.Progress = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetResponse(v *GetAivppAlgoJobResponseBodyDataResponse) *GetAivppAlgoJobResponseBodyData {
	s.Response = v
	return s
}

func (s *GetAivppAlgoJobResponseBodyData) SetStatus(v string) *GetAivppAlgoJobResponseBodyData {
	s.Status = &v
	return s
}

type GetAivppAlgoJobResponseBodyDataResponse struct {
	// example:
	//
	// {}
	DebugInfo interface{} `json:"DebugInfo,omitempty" xml:"DebugInfo,omitempty"`
	// example:
	//
	// PowerForecast
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// {}
	Result interface{} `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s GetAivppAlgoJobResponseBodyDataResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAivppAlgoJobResponseBodyDataResponse) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) SetDebugInfo(v interface{}) *GetAivppAlgoJobResponseBodyDataResponse {
	s.DebugInfo = v
	return s
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) SetJobType(v string) *GetAivppAlgoJobResponseBodyDataResponse {
	s.JobType = &v
	return s
}

func (s *GetAivppAlgoJobResponseBodyDataResponse) SetResult(v interface{}) *GetAivppAlgoJobResponseBodyDataResponse {
	s.Result = v
	return s
}

type GetAivppAlgoJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAivppAlgoJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAivppAlgoJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAivppAlgoJobResponse) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobResponse) SetHeaders(v map[string]*string) *GetAivppAlgoJobResponse {
	s.Headers = v
	return s
}

func (s *GetAivppAlgoJobResponse) SetStatusCode(v int32) *GetAivppAlgoJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAivppAlgoJobResponse) SetBody(v *GetAivppAlgoJobResponseBody) *GetAivppAlgoJobResponse {
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

type ListAivppResourcesRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
}

func (s ListAivppResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAivppResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListAivppResourcesRequest) SetCurrentPage(v int32) *ListAivppResourcesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAivppResourcesRequest) SetMaxResults(v int32) *ListAivppResourcesRequest {
	s.MaxResults = &v
	return s
}

type ListAivppResourcesResponseBody struct {
	// example:
	//
	// 200
	Code *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListAivppResourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 37cc36b4344b182d36b45d95ae4ef03952ee5c24733ba461
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAivppResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAivppResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAivppResourcesResponseBody) SetCode(v string) *ListAivppResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetData(v []*ListAivppResourcesResponseBodyData) *ListAivppResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListAivppResourcesResponseBody) SetMaxResults(v int32) *ListAivppResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetMessage(v string) *ListAivppResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetNextToken(v string) *ListAivppResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetRequestId(v string) *ListAivppResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetSuccess(v string) *ListAivppResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAivppResourcesResponseBody) SetTotalCount(v int32) *ListAivppResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListAivppResourcesResponseBodyData struct {
	Detail *string `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// example:
	//
	// 2034-03-09T17:47:11Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// i-bp154xh3gt3adb4xu1ue
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// api
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// 1
	LeftQuantity *string `json:"LeftQuantity,omitempty" xml:"LeftQuantity,omitempty"`
	// example:
	//
	// 231287932080007
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 1
	Quantity *string `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// example:
	//
	// LoadForecasting
	Specification *string `json:"Specification,omitempty" xml:"Specification,omitempty"`
	// example:
	//
	// 2024-09-08T01:16Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListAivppResourcesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ListAivppResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAivppResourcesResponseBodyData) SetDetail(v string) *ListAivppResourcesResponseBodyData {
	s.Detail = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetExpireTime(v string) *ListAivppResourcesResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetInstanceId(v string) *ListAivppResourcesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetInstanceType(v string) *ListAivppResourcesResponseBodyData {
	s.InstanceType = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetLeftQuantity(v string) *ListAivppResourcesResponseBodyData {
	s.LeftQuantity = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetOrderId(v string) *ListAivppResourcesResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetQuantity(v string) *ListAivppResourcesResponseBodyData {
	s.Quantity = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetSpecification(v string) *ListAivppResourcesResponseBodyData {
	s.Specification = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetStartTime(v string) *ListAivppResourcesResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetStatus(v string) *ListAivppResourcesResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAivppResourcesResponseBodyData) SetUserId(v string) *ListAivppResourcesResponseBodyData {
	s.UserId = &v
	return s
}

type ListAivppResourcesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAivppResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAivppResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAivppResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListAivppResourcesResponse) SetHeaders(v map[string]*string) *ListAivppResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListAivppResourcesResponse) SetStatusCode(v int32) *ListAivppResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAivppResourcesResponse) SetBody(v *ListAivppResourcesResponseBody) *ListAivppResourcesResponse {
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ActivateLicenseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ActivateLicenseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 创建储能运行优化任务
//
// @param tmpReq - CreateEssOptJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateEssOptJobResponse
func (client *Client) CreateEssOptJobWithOptions(tmpReq *CreateEssOptJobRequest, runtime *util.RuntimeOptions) (_result *CreateEssOptJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateEssOptJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ElecPrice)) {
		request.ElecPriceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ElecPrice, tea.String("ElecPrice"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.GenPrice)) {
		request.GenPriceShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.GenPrice, tea.String("GenPrice"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Location)) {
		request.LocationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Location, tea.String("Location"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SystemData)) {
		request.SystemDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SystemData, tea.String("SystemData"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.ElecPriceShrink)) {
		body["ElecPrice"] = request.ElecPriceShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Freq)) {
		body["Freq"] = request.Freq
	}

	if !tea.BoolValue(util.IsUnset(request.GenPriceShrink)) {
		body["GenPrice"] = request.GenPriceShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LocationShrink)) {
		body["Location"] = request.LocationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ModelVersion)) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RunDate)) {
		body["RunDate"] = request.RunDate
	}

	if !tea.BoolValue(util.IsUnset(request.SystemDataShrink)) {
		body["SystemData"] = request.SystemDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		body["TimeZone"] = request.TimeZone
	}

	if !tea.BoolValue(util.IsUnset(request.TopoType)) {
		body["TopoType"] = request.TopoType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateEssOptJob"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateEssOptJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateEssOptJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建储能运行优化任务
//
// @param request - CreateEssOptJobRequest
//
// @return CreateEssOptJobResponse
func (client *Client) CreateEssOptJob(request *CreateEssOptJobRequest) (_result *CreateEssOptJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateEssOptJobResponse{}
	_body, _err := client.CreateEssOptJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建用电负荷预测任务
//
// @param tmpReq - CreateLoadForecastJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateLoadForecastJobResponse
func (client *Client) CreateLoadForecastJobWithOptions(tmpReq *CreateLoadForecastJobRequest, runtime *util.RuntimeOptions) (_result *CreateLoadForecastJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateLoadForecastJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HistoryData)) {
		request.HistoryDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HistoryData, tea.String("HistoryData"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Freq)) {
		body["Freq"] = request.Freq
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryDataShrink)) {
		body["HistoryData"] = request.HistoryDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ModelVersion)) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RunDate)) {
		body["RunDate"] = request.RunDate
	}

	if !tea.BoolValue(util.IsUnset(request.SystemType)) {
		body["SystemType"] = request.SystemType
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		body["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLoadForecastJob"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreateLoadForecastJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreateLoadForecastJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建用电负荷预测任务
//
// @param request - CreateLoadForecastJobRequest
//
// @return CreateLoadForecastJobResponse
func (client *Client) CreateLoadForecastJob(request *CreateLoadForecastJobRequest) (_result *CreateLoadForecastJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateLoadForecastJobResponse{}
	_body, _err := client.CreateLoadForecastJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建发电功率预测任务
//
// @param tmpReq - CreatePowerForecastJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreatePowerForecastJobResponse
func (client *Client) CreatePowerForecastJobWithOptions(tmpReq *CreatePowerForecastJobRequest, runtime *util.RuntimeOptions) (_result *CreatePowerForecastJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreatePowerForecastJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.HistoryData)) {
		request.HistoryDataShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.HistoryData, tea.String("HistoryData"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Location)) {
		request.LocationShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Location, tea.String("Location"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceType)) {
		body["DeviceType"] = request.DeviceType
	}

	if !tea.BoolValue(util.IsUnset(request.Duration)) {
		body["Duration"] = request.Duration
	}

	if !tea.BoolValue(util.IsUnset(request.Freq)) {
		body["Freq"] = request.Freq
	}

	if !tea.BoolValue(util.IsUnset(request.HistoryDataShrink)) {
		body["HistoryData"] = request.HistoryDataShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LocationShrink)) {
		body["Location"] = request.LocationShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ModelVersion)) {
		body["ModelVersion"] = request.ModelVersion
	}

	if !tea.BoolValue(util.IsUnset(request.RunDate)) {
		body["RunDate"] = request.RunDate
	}

	if !tea.BoolValue(util.IsUnset(request.SystemType)) {
		body["SystemType"] = request.SystemType
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		body["TimeZone"] = request.TimeZone
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePowerForecastJob"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &CreatePowerForecastJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &CreatePowerForecastJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建发电功率预测任务
//
// @param request - CreatePowerForecastJobRequest
//
// @return CreatePowerForecastJobResponse
func (client *Client) CreatePowerForecastJob(request *CreatePowerForecastJobRequest) (_result *CreatePowerForecastJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePowerForecastJobResponse{}
	_body, _err := client.CreatePowerForecastJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询aivpp算法job
//
// @param request - GetAivppAlgoJobRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAivppAlgoJobResponse
func (client *Client) GetAivppAlgoJobWithOptions(request *GetAivppAlgoJobRequest, runtime *util.RuntimeOptions) (_result *GetAivppAlgoJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAivppAlgoJob"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetAivppAlgoJobResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetAivppAlgoJobResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询aivpp算法job
//
// @param request - GetAivppAlgoJobRequest
//
// @return GetAivppAlgoJobResponse
func (client *Client) GetAivppAlgoJob(request *GetAivppAlgoJobRequest) (_result *GetAivppAlgoJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAivppAlgoJobResponse{}
	_body, _err := client.GetAivppAlgoJobWithOptions(request, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &GetLicenseResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &GetLicenseResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
// 获取用户AIVPP资源列表
//
// @param request - ListAivppResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAivppResourcesResponse
func (client *Client) ListAivppResourcesWithOptions(request *ListAivppResourcesRequest, runtime *util.RuntimeOptions) (_result *ListAivppResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentPage)) {
		body["CurrentPage"] = request.CurrentPage
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAivppResources"),
		Version:     tea.String("2020-09-20"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListAivppResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListAivppResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取用户AIVPP资源列表
//
// @param request - ListAivppResourcesRequest
//
// @return ListAivppResourcesResponse
func (client *Client) ListAivppResources(request *ListAivppResourcesRequest) (_result *ListAivppResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAivppResourcesResponse{}
	_body, _err := client.ListAivppResourcesWithOptions(request, runtime)
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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListLicensesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListLicensesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &ListUserResourcesResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &ListUserResourcesResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
	if tea.BoolValue(util.IsUnset(client.SignatureVersion)) || !tea.BoolValue(util.EqualString(client.SignatureVersion, tea.String("v4"))) {
		_result = &UpdateLicenseDescriptionResponse{}
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	} else {
		_result = &UpdateLicenseDescriptionResponse{}
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = tea.Convert(_body, &_result)
		return _result, _err
	}

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
