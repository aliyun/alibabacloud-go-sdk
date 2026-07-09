// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInvoiceCandidateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListInvoiceCandidateResponseBody
	GetCurrentPage() *int32
	SetData(v []*ListInvoiceCandidateResponseBodyData) *ListInvoiceCandidateResponseBody
	GetData() []*ListInvoiceCandidateResponseBodyData
	SetMetadata(v interface{}) *ListInvoiceCandidateResponseBody
	GetMetadata() interface{}
	SetPageSize(v int32) *ListInvoiceCandidateResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListInvoiceCandidateResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListInvoiceCandidateResponseBody
	GetTotalCount() *int32
}

type ListInvoiceCandidateResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of data entries.
	Data []*ListInvoiceCandidateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The metadata of the response.
	//
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInvoiceCandidateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInvoiceCandidateResponseBody) GoString() string {
	return s.String()
}

func (s *ListInvoiceCandidateResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInvoiceCandidateResponseBody) GetData() []*ListInvoiceCandidateResponseBodyData {
	return s.Data
}

func (s *ListInvoiceCandidateResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ListInvoiceCandidateResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInvoiceCandidateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInvoiceCandidateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListInvoiceCandidateResponseBody) SetCurrentPage(v int32) *ListInvoiceCandidateResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListInvoiceCandidateResponseBody) SetData(v []*ListInvoiceCandidateResponseBodyData) *ListInvoiceCandidateResponseBody {
	s.Data = v
	return s
}

func (s *ListInvoiceCandidateResponseBody) SetMetadata(v interface{}) *ListInvoiceCandidateResponseBody {
	s.Metadata = v
	return s
}

func (s *ListInvoiceCandidateResponseBody) SetPageSize(v int32) *ListInvoiceCandidateResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListInvoiceCandidateResponseBody) SetRequestId(v string) *ListInvoiceCandidateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInvoiceCandidateResponseBody) SetTotalCount(v int32) *ListInvoiceCandidateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInvoiceCandidateResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInvoiceCandidateResponseBodyData struct {
	// The accepted offset amount.
	//
	// example:
	//
	// 0.01
	AcceptedOffsetAmount *string `json:"AcceptedOffsetAmount,omitempty" xml:"AcceptedOffsetAmount,omitempty"`
	// The account ID.
	//
	// example:
	//
	// 1990699401005016
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// The account name.
	//
	// example:
	//
	// 测试账号
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The billing cycle.
	//
	// example:
	//
	// 202506
	BillingCycle *int32 `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The business document number.
	//
	// example:
	//
	// 202506
	BusinessId *string `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The time when the business event occurred.
	//
	// example:
	//
	// 2025-06-01 00:00:00
	BusinessTime *string `json:"BusinessTime,omitempty" xml:"BusinessTime,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// pts
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The commodity name.
	//
	// example:
	//
	// 性能测试
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-06-91 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the invoice candidate.
	//
	// example:
	//
	// 12345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The invoice issuer.
	//
	// example:
	//
	// ALIYUN_SERVICE
	InvoiceIssuer *string `json:"InvoiceIssuer,omitempty" xml:"InvoiceIssuer,omitempty"`
	// The invoiceable amount.
	//
	// example:
	//
	// 0.01
	InvoiceableAmount *string `json:"InvoiceableAmount,omitempty" xml:"InvoiceableAmount,omitempty"`
	// The invoiced amount.
	//
	// example:
	//
	// 0
	InvoicedAmount *string `json:"InvoicedAmount,omitempty" xml:"InvoicedAmount,omitempty"`
	// The offset amount.
	//
	// example:
	//
	// 0
	OffsetAmount *string `json:"OffsetAmount,omitempty" xml:"OffsetAmount,omitempty"`
	// The product code.
	//
	// example:
	//
	// pts
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product name.
	//
	// example:
	//
	// 性能测试
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The ID of the resource owner account.
	//
	// example:
	//
	// 1990699401005016
	ResourceOwnerAccountId *int64 `json:"ResourceOwnerAccountId,omitempty" xml:"ResourceOwnerAccountId,omitempty"`
	// The name of the resource owner account.
	//
	// example:
	//
	// 测试账号
	ResourceOwnerAccountName *string `json:"ResourceOwnerAccountName,omitempty" xml:"ResourceOwnerAccountName,omitempty"`
	// The status of the invoice candidate.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total amount.
	//
	// example:
	//
	// 0.01
	TotalAmount *string `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
	// The type of the invoice candidate.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListInvoiceCandidateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInvoiceCandidateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInvoiceCandidateResponseBodyData) GetAcceptedOffsetAmount() *string {
	return s.AcceptedOffsetAmount
}

func (s *ListInvoiceCandidateResponseBodyData) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListInvoiceCandidateResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *ListInvoiceCandidateResponseBodyData) GetBillingCycle() *int32 {
	return s.BillingCycle
}

func (s *ListInvoiceCandidateResponseBodyData) GetBusinessId() *string {
	return s.BusinessId
}

func (s *ListInvoiceCandidateResponseBodyData) GetBusinessTime() *string {
	return s.BusinessTime
}

func (s *ListInvoiceCandidateResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListInvoiceCandidateResponseBodyData) GetCommodityName() *string {
	return s.CommodityName
}

func (s *ListInvoiceCandidateResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInvoiceCandidateResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListInvoiceCandidateResponseBodyData) GetInvoiceIssuer() *string {
	return s.InvoiceIssuer
}

func (s *ListInvoiceCandidateResponseBodyData) GetInvoiceableAmount() *string {
	return s.InvoiceableAmount
}

func (s *ListInvoiceCandidateResponseBodyData) GetInvoicedAmount() *string {
	return s.InvoicedAmount
}

func (s *ListInvoiceCandidateResponseBodyData) GetOffsetAmount() *string {
	return s.OffsetAmount
}

func (s *ListInvoiceCandidateResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *ListInvoiceCandidateResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *ListInvoiceCandidateResponseBodyData) GetResourceOwnerAccountId() *int64 {
	return s.ResourceOwnerAccountId
}

func (s *ListInvoiceCandidateResponseBodyData) GetResourceOwnerAccountName() *string {
	return s.ResourceOwnerAccountName
}

func (s *ListInvoiceCandidateResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListInvoiceCandidateResponseBodyData) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *ListInvoiceCandidateResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *ListInvoiceCandidateResponseBodyData) SetAcceptedOffsetAmount(v string) *ListInvoiceCandidateResponseBodyData {
	s.AcceptedOffsetAmount = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetAccountId(v int64) *ListInvoiceCandidateResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetAccountName(v string) *ListInvoiceCandidateResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetBillingCycle(v int32) *ListInvoiceCandidateResponseBodyData {
	s.BillingCycle = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetBusinessId(v string) *ListInvoiceCandidateResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetBusinessTime(v string) *ListInvoiceCandidateResponseBodyData {
	s.BusinessTime = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetCommodityCode(v string) *ListInvoiceCandidateResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetCommodityName(v string) *ListInvoiceCandidateResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetCreateTime(v string) *ListInvoiceCandidateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetId(v string) *ListInvoiceCandidateResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetInvoiceIssuer(v string) *ListInvoiceCandidateResponseBodyData {
	s.InvoiceIssuer = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetInvoiceableAmount(v string) *ListInvoiceCandidateResponseBodyData {
	s.InvoiceableAmount = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetInvoicedAmount(v string) *ListInvoiceCandidateResponseBodyData {
	s.InvoicedAmount = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetOffsetAmount(v string) *ListInvoiceCandidateResponseBodyData {
	s.OffsetAmount = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetProductCode(v string) *ListInvoiceCandidateResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetProductName(v string) *ListInvoiceCandidateResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetResourceOwnerAccountId(v int64) *ListInvoiceCandidateResponseBodyData {
	s.ResourceOwnerAccountId = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetResourceOwnerAccountName(v string) *ListInvoiceCandidateResponseBodyData {
	s.ResourceOwnerAccountName = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetStatus(v int32) *ListInvoiceCandidateResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetTotalAmount(v string) *ListInvoiceCandidateResponseBodyData {
	s.TotalAmount = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) SetType(v int32) *ListInvoiceCandidateResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListInvoiceCandidateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
