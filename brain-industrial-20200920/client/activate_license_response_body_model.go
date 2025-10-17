// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ActivateLicenseResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ActivateLicenseResponseBody
	GetCode() *string
	SetData(v *ActivateLicenseResponseBodyData) *ActivateLicenseResponseBody
	GetData() *ActivateLicenseResponseBodyData
	SetHttpStatusCode(v int32) *ActivateLicenseResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ActivateLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *ActivateLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ActivateLicenseResponseBody
	GetSuccess() *string
}

type ActivateLicenseResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ActivateLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 65308A66-8764-53EE-8D4A-201E86CA88C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActivateLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ActivateLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *ActivateLicenseResponseBody) GetData() *ActivateLicenseResponseBodyData {
	return s.Data
}

func (s *ActivateLicenseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ActivateLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ActivateLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateLicenseResponseBody) GetSuccess() *string {
	return s.Success
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

func (s *ActivateLicenseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ActivateLicenseResponseBodyData struct {
	ActivateRecord []*ActivateLicenseResponseBodyDataActivateRecord `json:"ActivateRecord,omitempty" xml:"ActivateRecord,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-04-29 15:03:12
	ActivateTime    *string `json:"ActivateTime,omitempty" xml:"ActivateTime,omitempty"`
	AllDuration     *string `json:"AllDuration,omitempty" xml:"AllDuration,omitempty"`
	ApplicableSpecs *string `json:"ApplicableSpecs,omitempty" xml:"ApplicableSpecs,omitempty"`
	// example:
	//
	// 2024-04-29 15:00:15
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
	// example:
	//
	// XXXX
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// Id
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// c39128fcb74e482546a72de90cdsaba3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// XXX
	LicenseCode     *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	LicenseSpecName *string `json:"LicenseSpecName,omitempty" xml:"LicenseSpecName,omitempty"`
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

func (s ActivateLicenseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ActivateLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBodyData) GetActivateRecord() []*ActivateLicenseResponseBodyDataActivateRecord {
	return s.ActivateRecord
}

func (s *ActivateLicenseResponseBodyData) GetActivateTime() *string {
	return s.ActivateTime
}

func (s *ActivateLicenseResponseBodyData) GetAllDuration() *string {
	return s.AllDuration
}

func (s *ActivateLicenseResponseBodyData) GetApplicableSpecs() *string {
	return s.ApplicableSpecs
}

func (s *ActivateLicenseResponseBodyData) GetBuyTime() *string {
	return s.BuyTime
}

func (s *ActivateLicenseResponseBodyData) GetCpuLimit() *int32 {
	return s.CpuLimit
}

func (s *ActivateLicenseResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ActivateLicenseResponseBodyData) GetDuration() *string {
	return s.Duration
}

func (s *ActivateLicenseResponseBodyData) GetEffectTime() *string {
	return s.EffectTime
}

func (s *ActivateLicenseResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ActivateLicenseResponseBodyData) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *ActivateLicenseResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ActivateLicenseResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ActivateLicenseResponseBodyData) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *ActivateLicenseResponseBodyData) GetLicenseSpecName() *string {
	return s.LicenseSpecName
}

func (s *ActivateLicenseResponseBodyData) GetMemoryLimit() *int32 {
	return s.MemoryLimit
}

func (s *ActivateLicenseResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ActivateLicenseResponseBodyData) GetUnActivateAllDuration() *string {
	return s.UnActivateAllDuration
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

func (s *ActivateLicenseResponseBodyData) Validate() error {
	if s.ActivateRecord != nil {
		for _, item := range s.ActivateRecord {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ActivateLicenseResponseBodyDataActivateRecord struct {
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
	// XXX
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// example:
	//
	// XXX
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// activated
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ActivateLicenseResponseBodyDataActivateRecord) String() string {
	return dara.Prettify(s)
}

func (s ActivateLicenseResponseBodyDataActivateRecord) GoString() string {
	return s.String()
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) GetActivateTime() *string {
	return s.ActivateTime
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) GetBuyTime() *string {
	return s.BuyTime
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) GetDuration() *string {
	return s.Duration
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) GetOrderId() *string {
	return s.OrderId
}

func (s *ActivateLicenseResponseBodyDataActivateRecord) GetStatus() *string {
	return s.Status
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

func (s *ActivateLicenseResponseBodyDataActivateRecord) Validate() error {
	return dara.Validate(s)
}
