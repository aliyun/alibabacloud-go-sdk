// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLicenseResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetLicenseResponseBody
	GetCode() *string
	SetData(v *GetLicenseResponseBodyData) *GetLicenseResponseBody
	GetData() *GetLicenseResponseBodyData
	SetHttpStatusCode(v int32) *GetLicenseResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetLicenseResponseBody
	GetSuccess() *string
}

type GetLicenseResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 68738E75-43C1-5AE5-9F3A-AFEF576D7B5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetLicenseResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLicenseResponseBody) GetData() *GetLicenseResponseBodyData {
	return s.Data
}

func (s *GetLicenseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLicenseResponseBody) GetSuccess() *string {
	return s.Success
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

func (s *GetLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLicenseResponseBodyData struct {
	ActivateRecord []*GetLicenseResponseBodyDataActivateRecord `json:"ActivateRecord,omitempty" xml:"ActivateRecord,omitempty" type:"Repeated"`
	// 代表资源一级ID的资源属性字段
	//
	// example:
	//
	// 2024-04-29 15:03:12
	ActivateTime    *string `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	AdaptiveMachine *string `json:"AdaptiveMachine,omitempty" xml:"AdaptiveMachine,omitempty"`
	AllDuration     *string `json:"AllDuration,omitempty" xml:"AllDuration,omitempty"`
	ApplicableSpecs *string `json:"ApplicableSpecs,omitempty" xml:"ApplicableSpecs,omitempty"`
	// 代表资源名称的资源属性字段
	//
	// example:
	//
	// 2024-04-29 15:01:15
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
	// 2024-05-01 00:00:00
	EffectTime *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty"`
	// example:
	//
	// 2025-05-01 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// 代表创建时间的资源属性字段
	//
	// example:
	//
	// XXX
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// ID
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// c31238fcb74e482588a72de90cd7dba3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// XXX
	LicenseCode     *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	LicenseSpecCode *string `json:"LicenseSpecCode,omitempty" xml:"LicenseSpecCode,omitempty"`
	// 代表资源组的资源属性字段
	LicenseSpecName *string `json:"LicenseSpecName,omitempty" xml:"LicenseSpecName,omitempty"`
	LicenseSpecType *string `json:"LicenseSpecType,omitempty" xml:"LicenseSpecType,omitempty"`
	// example:
	//
	// 256
	MemoryLimit *int32  `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	Proposal    *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
	// example:
	//
	// activated
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UnActivateAllDuration *string `json:"UnActivateAllDuration,omitempty" xml:"UnActivateAllDuration,omitempty"`
}

func (s GetLicenseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLicenseResponseBodyData) GetActivateRecord() []*GetLicenseResponseBodyDataActivateRecord {
	return s.ActivateRecord
}

func (s *GetLicenseResponseBodyData) GetActivateTime() *string {
	return s.ActivateTime
}

func (s *GetLicenseResponseBodyData) GetAdaptiveMachine() *string {
	return s.AdaptiveMachine
}

func (s *GetLicenseResponseBodyData) GetAllDuration() *string {
	return s.AllDuration
}

func (s *GetLicenseResponseBodyData) GetApplicableSpecs() *string {
	return s.ApplicableSpecs
}

func (s *GetLicenseResponseBodyData) GetBuyTime() *string {
	return s.BuyTime
}

func (s *GetLicenseResponseBodyData) GetCpuLimit() *int32 {
	return s.CpuLimit
}

func (s *GetLicenseResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetLicenseResponseBodyData) GetDuration() *string {
	return s.Duration
}

func (s *GetLicenseResponseBodyData) GetEffectTime() *string {
	return s.EffectTime
}

func (s *GetLicenseResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetLicenseResponseBodyData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *GetLicenseResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetLicenseResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLicenseResponseBodyData) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *GetLicenseResponseBodyData) GetLicenseSpecCode() *string {
	return s.LicenseSpecCode
}

func (s *GetLicenseResponseBodyData) GetLicenseSpecName() *string {
	return s.LicenseSpecName
}

func (s *GetLicenseResponseBodyData) GetLicenseSpecType() *string {
	return s.LicenseSpecType
}

func (s *GetLicenseResponseBodyData) GetMemoryLimit() *int32 {
	return s.MemoryLimit
}

func (s *GetLicenseResponseBodyData) GetProposal() *string {
	return s.Proposal
}

func (s *GetLicenseResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetLicenseResponseBodyData) GetUnActivateAllDuration() *string {
	return s.UnActivateAllDuration
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

func (s *GetLicenseResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetLicenseResponseBodyDataActivateRecord struct {
	// example:
	//
	// 2024-04-29 15:03:12
	ActivateTime *string `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	// example:
	//
	// 2024-04-29 15:01:15
	BuyTime  *string `json:"BuyTime,omitempty" xml:"BuyTime,omitempty"`
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2025-05-01 00:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// XXXX
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// example:
	//
	// 1040b91765d14ad1a40ff22589070342
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetLicenseResponseBodyDataActivateRecord) String() string {
	return dara.Prettify(s)
}

func (s GetLicenseResponseBodyDataActivateRecord) GoString() string {
	return s.String()
}

func (s *GetLicenseResponseBodyDataActivateRecord) GetActivateTime() *string {
	return s.ActivateTime
}

func (s *GetLicenseResponseBodyDataActivateRecord) GetBuyTime() *string {
	return s.BuyTime
}

func (s *GetLicenseResponseBodyDataActivateRecord) GetDuration() *string {
	return s.Duration
}

func (s *GetLicenseResponseBodyDataActivateRecord) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetLicenseResponseBodyDataActivateRecord) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *GetLicenseResponseBodyDataActivateRecord) GetOrderId() *string {
	return s.OrderId
}

func (s *GetLicenseResponseBodyDataActivateRecord) GetStatus() *string {
	return s.Status
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

func (s *GetLicenseResponseBodyDataActivateRecord) Validate() error {
	return dara.Validate(s)
}
