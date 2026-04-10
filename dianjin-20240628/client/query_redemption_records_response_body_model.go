// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRedemptionRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRedemptionRecordsResponseBody
	GetCode() *string
	SetData(v *QueryRedemptionRecordsResponseBodyData) *QueryRedemptionRecordsResponseBody
	GetData() *QueryRedemptionRecordsResponseBodyData
	SetMessage(v string) *QueryRedemptionRecordsResponseBody
	GetMessage() *string
	SetRetryAble(v bool) *QueryRedemptionRecordsResponseBody
	GetRetryAble() *bool
	SetSuccess(v bool) *QueryRedemptionRecordsResponseBody
	GetSuccess() *bool
}

type QueryRedemptionRecordsResponseBody struct {
	// example:
	//
	// 0
	Code *string                                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *QueryRedemptionRecordsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// false
	RetryAble *bool `json:"retryAble,omitempty" xml:"retryAble,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryRedemptionRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRedemptionRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRedemptionRecordsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRedemptionRecordsResponseBody) GetData() *QueryRedemptionRecordsResponseBodyData {
	return s.Data
}

func (s *QueryRedemptionRecordsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryRedemptionRecordsResponseBody) GetRetryAble() *bool {
	return s.RetryAble
}

func (s *QueryRedemptionRecordsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRedemptionRecordsResponseBody) SetCode(v string) *QueryRedemptionRecordsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBody) SetData(v *QueryRedemptionRecordsResponseBodyData) *QueryRedemptionRecordsResponseBody {
	s.Data = v
	return s
}

func (s *QueryRedemptionRecordsResponseBody) SetMessage(v string) *QueryRedemptionRecordsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBody) SetRetryAble(v bool) *QueryRedemptionRecordsResponseBody {
	s.RetryAble = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBody) SetSuccess(v bool) *QueryRedemptionRecordsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRedemptionRecordsResponseBodyData struct {
	Items []*QueryRedemptionRecordsResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 0
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"totalPages,omitempty" xml:"totalPages,omitempty"`
}

func (s QueryRedemptionRecordsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryRedemptionRecordsResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRedemptionRecordsResponseBodyData) GetItems() []*QueryRedemptionRecordsResponseBodyDataItems {
	return s.Items
}

func (s *QueryRedemptionRecordsResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *QueryRedemptionRecordsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRedemptionRecordsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *QueryRedemptionRecordsResponseBodyData) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *QueryRedemptionRecordsResponseBodyData) SetItems(v []*QueryRedemptionRecordsResponseBodyDataItems) *QueryRedemptionRecordsResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyData) SetPage(v int32) *QueryRedemptionRecordsResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyData) SetPageSize(v int32) *QueryRedemptionRecordsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyData) SetTotal(v int64) *QueryRedemptionRecordsResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyData) SetTotalPages(v int32) *QueryRedemptionRecordsResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRedemptionRecordsResponseBodyDataItems struct {
	AllowedModels []*string `json:"allowedModels,omitempty" xml:"allowedModels,omitempty" type:"Repeated"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	EffectiveAt *string `json:"effectiveAt,omitempty" xml:"effectiveAt,omitempty"`
	// example:
	//
	// 2024-01-31T23:59:59Z
	ExpireAt *string `json:"expireAt,omitempty" xml:"expireAt,omitempty"`
	// example:
	//
	// 2024-01-01T00:00:00Z
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// a1b2c3d4e5f6...
	KeyHash *string `json:"keyHash,omitempty" xml:"keyHash,omitempty"`
	// example:
	//
	// AR123233333
	OutBizNo *string `json:"outBizNo,omitempty" xml:"outBizNo,omitempty"`
	// example:
	//
	// 500
	QuotaBalance *int64 `json:"quotaBalance,omitempty" xml:"quotaBalance,omitempty"`
	// example:
	//
	// AVAILABLE
	QuotaStatus *string `json:"quotaStatus,omitempty" xml:"quotaStatus,omitempty"`
	// example:
	//
	// ORD20240101000001
	RedemptionOrderNo *string `json:"redemptionOrderNo,omitempty" xml:"redemptionOrderNo,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 37624
	TemplateId *int64 `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// example:
	//
	// 10001
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryRedemptionRecordsResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s QueryRedemptionRecordsResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetAllowedModels() []*string {
	return s.AllowedModels
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetEffectiveAt() *string {
	return s.EffectiveAt
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetId() *int64 {
	return s.Id
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetKeyHash() *string {
	return s.KeyHash
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetOutBizNo() *string {
	return s.OutBizNo
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetQuotaBalance() *int64 {
	return s.QuotaBalance
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetQuotaStatus() *string {
	return s.QuotaStatus
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetRedemptionOrderNo() *string {
	return s.RedemptionOrderNo
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetStatus() *string {
	return s.Status
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) GetTenantId() *int64 {
	return s.TenantId
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetAllowedModels(v []*string) *QueryRedemptionRecordsResponseBodyDataItems {
	s.AllowedModels = v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetEffectiveAt(v string) *QueryRedemptionRecordsResponseBodyDataItems {
	s.EffectiveAt = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetExpireAt(v string) *QueryRedemptionRecordsResponseBodyDataItems {
	s.ExpireAt = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetGmtCreate(v string) *QueryRedemptionRecordsResponseBodyDataItems {
	s.GmtCreate = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetId(v int64) *QueryRedemptionRecordsResponseBodyDataItems {
	s.Id = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetKeyHash(v string) *QueryRedemptionRecordsResponseBodyDataItems {
	s.KeyHash = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetOutBizNo(v string) *QueryRedemptionRecordsResponseBodyDataItems {
	s.OutBizNo = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetQuotaBalance(v int64) *QueryRedemptionRecordsResponseBodyDataItems {
	s.QuotaBalance = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetQuotaStatus(v string) *QueryRedemptionRecordsResponseBodyDataItems {
	s.QuotaStatus = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetRedemptionOrderNo(v string) *QueryRedemptionRecordsResponseBodyDataItems {
	s.RedemptionOrderNo = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetStatus(v string) *QueryRedemptionRecordsResponseBodyDataItems {
	s.Status = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetTemplateId(v int64) *QueryRedemptionRecordsResponseBodyDataItems {
	s.TemplateId = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) SetTenantId(v int64) *QueryRedemptionRecordsResponseBodyDataItems {
	s.TenantId = &v
	return s
}

func (s *QueryRedemptionRecordsResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
