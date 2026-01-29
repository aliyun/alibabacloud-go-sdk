// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEvaluateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryEvaluateListResponseBody
	GetCode() *string
	SetData(v *QueryEvaluateListResponseBodyData) *QueryEvaluateListResponseBody
	GetData() *QueryEvaluateListResponseBodyData
	SetMessage(v string) *QueryEvaluateListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryEvaluateListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryEvaluateListResponseBody
	GetSuccess() *bool
}

type QueryEvaluateListResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryEvaluateListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// Successful！
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D35FF10E-1B2E-4ABA-8401-0AE17725F50B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryEvaluateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEvaluateListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryEvaluateListResponseBody) GetData() *QueryEvaluateListResponseBodyData {
	return s.Data
}

func (s *QueryEvaluateListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryEvaluateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEvaluateListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryEvaluateListResponseBody) SetCode(v string) *QueryEvaluateListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEvaluateListResponseBody) SetData(v *QueryEvaluateListResponseBodyData) *QueryEvaluateListResponseBody {
	s.Data = v
	return s
}

func (s *QueryEvaluateListResponseBody) SetMessage(v string) *QueryEvaluateListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEvaluateListResponseBody) SetRequestId(v string) *QueryEvaluateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEvaluateListResponseBody) SetSuccess(v bool) *QueryEvaluateListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryEvaluateListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEvaluateListResponseBodyData struct {
	// The data returned.
	EvaluateList *QueryEvaluateListResponseBodyDataEvaluateList `json:"EvaluateList,omitempty" xml:"EvaluateList,omitempty" type:"Struct"`
	// The ID of the host.
	//
	// example:
	//
	// cn
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	// The number of the page returned.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The invoiced amount that meets the query conditions. Unit: Cent.
	//
	// example:
	//
	// 12344
	TotalInvoiceAmount *int64 `json:"TotalInvoiceAmount,omitempty" xml:"TotalInvoiceAmount,omitempty"`
	// The invoiceable amount that meets the query conditions. Unit: Cent.
	//
	// example:
	//
	// 12344
	TotalUnAppliedInvoiceAmount *int64 `json:"TotalUnAppliedInvoiceAmount,omitempty" xml:"TotalUnAppliedInvoiceAmount,omitempty"`
}

func (s QueryEvaluateListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryEvaluateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponseBodyData) GetEvaluateList() *QueryEvaluateListResponseBodyDataEvaluateList {
	return s.EvaluateList
}

func (s *QueryEvaluateListResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *QueryEvaluateListResponseBodyData) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryEvaluateListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryEvaluateListResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryEvaluateListResponseBodyData) GetTotalInvoiceAmount() *int64 {
	return s.TotalInvoiceAmount
}

func (s *QueryEvaluateListResponseBodyData) GetTotalUnAppliedInvoiceAmount() *int64 {
	return s.TotalUnAppliedInvoiceAmount
}

func (s *QueryEvaluateListResponseBodyData) SetEvaluateList(v *QueryEvaluateListResponseBodyDataEvaluateList) *QueryEvaluateListResponseBodyData {
	s.EvaluateList = v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetHostId(v string) *QueryEvaluateListResponseBodyData {
	s.HostId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetPageNum(v int32) *QueryEvaluateListResponseBodyData {
	s.PageNum = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetPageSize(v int32) *QueryEvaluateListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetTotalCount(v int32) *QueryEvaluateListResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetTotalInvoiceAmount(v int64) *QueryEvaluateListResponseBodyData {
	s.TotalInvoiceAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) SetTotalUnAppliedInvoiceAmount(v int64) *QueryEvaluateListResponseBodyData {
	s.TotalUnAppliedInvoiceAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyData) Validate() error {
	if s.EvaluateList != nil {
		if err := s.EvaluateList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEvaluateListResponseBodyDataEvaluateList struct {
	Evaluate []*QueryEvaluateListResponseBodyDataEvaluateListEvaluate `json:"Evaluate,omitempty" xml:"Evaluate,omitempty" type:"Repeated"`
}

func (s QueryEvaluateListResponseBodyDataEvaluateList) String() string {
	return dara.Prettify(s)
}

func (s QueryEvaluateListResponseBodyDataEvaluateList) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponseBodyDataEvaluateList) GetEvaluate() []*QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	return s.Evaluate
}

func (s *QueryEvaluateListResponseBodyDataEvaluateList) SetEvaluate(v []*QueryEvaluateListResponseBodyDataEvaluateListEvaluate) *QueryEvaluateListResponseBodyDataEvaluateList {
	s.Evaluate = v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateList) Validate() error {
	if s.Evaluate != nil {
		for _, item := range s.Evaluate {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryEvaluateListResponseBodyDataEvaluateListEvaluate struct {
	// The billing cycle.
	//
	// example:
	//
	// 202002
	BillCycle *string `json:"BillCycle,omitempty" xml:"BillCycle,omitempty"`
	// The ID of the bill.
	//
	// example:
	//
	// 234543254325
	BillId *int64 `json:"BillId,omitempty" xml:"BillId,omitempty"`
	// The time.
	//
	// example:
	//
	// 2018-10-10 18:05:44
	BizTime *string `json:"BizTime,omitempty" xml:"BizTime,omitempty"`
	// The market type in the invoice. Valid values:
	//
	// 	- ALIYUN: Alibaba Cloud
	//
	// 	- MARKETPLACE: Alibaba Cloud Marketplace
	//
	// example:
	//
	// ALIYUN
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The invoiceable amount.
	//
	// example:
	//
	// 123213
	CanInvoiceAmount *int64 `json:"CanInvoiceAmount,omitempty" xml:"CanInvoiceAmount,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2018-10-10 18:05:44
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2018-10-10 18:05:44
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the invoice.
	//
	// example:
	//
	// 1325321532
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The invoiced amount.
	//
	// example:
	//
	// 10000
	InvoicedAmount *int64 `json:"InvoicedAmount,omitempty" xml:"InvoicedAmount,omitempty"`
	// The ID of the item.
	//
	// example:
	//
	// 23453245
	ItemId *int64 `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	// The name of the object to be invoiced.
	//
	// example:
	//
	// Refund of a voucher with denomination marked
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// If a refund is issued due to an order such as an unsubscription order or a configuration downgrade order, the refund amount is used to offset the amount of the invoice. The value is consistent with the value of the **OffsetCostAmount*	- parameter.
	//
	// example:
	//
	// 500
	OffsetAcceptAmount *int64 `json:"OffsetAcceptAmount,omitempty" xml:"OffsetAcceptAmount,omitempty"`
	// The refund amount used to offset the amount of the invoice. If a refund is issued due to an order such as an unsubscription order or a configuration downgrade order, the refund amount is used to offset the amount of the invoice. The value is consistent with the value of the **OffsetAcceptAmount*	- parameter.
	//
	// example:
	//
	// 500
	OffsetCostAmount *int64 `json:"OffsetCostAmount,omitempty" xml:"OffsetCostAmount,omitempty"`
	// The ID of the external object.
	//
	// example:
	//
	// 12341
	OpId *string `json:"OpId,omitempty" xml:"OpId,omitempty"`
	// The original amount.
	//
	// example:
	//
	// -10000
	OriginalAmount *int64 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The ID of the external order.
	//
	// example:
	//
	// 124324213421
	OutBizId *string `json:"OutBizId,omitempty" xml:"OutBizId,omitempty"`
	// The balance.
	//
	// example:
	//
	// -10000
	PresentAmount *int64 `json:"PresentAmount,omitempty" xml:"PresentAmount,omitempty"`
	// The status of the invoiceable amount.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of orders that are queried. Valid values:
	//
	// 	- 1: the orders in which the invoiceable amount is negative.
	//
	// 	- 2: the orders in which the invoiceable amount is positive.
	//
	// 	- 3: the orders in which the invoiceable amount is not 0.
	//
	// 	- 4: the orders in which the amount that has been invoiced is greater than 0.
	//
	// >  By default, this parameter is left empty. If this parameter is left empty, all orders are queried.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 2738543
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// test
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s QueryEvaluateListResponseBodyDataEvaluateListEvaluate) String() string {
	return dara.Prettify(s)
}

func (s QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GoString() string {
	return s.String()
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetBillCycle() *string {
	return s.BillCycle
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetBillId() *int64 {
	return s.BillId
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetBizTime() *string {
	return s.BizTime
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetBizType() *string {
	return s.BizType
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetCanInvoiceAmount() *int64 {
	return s.CanInvoiceAmount
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetId() *int64 {
	return s.Id
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetInvoicedAmount() *int64 {
	return s.InvoicedAmount
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetItemId() *int64 {
	return s.ItemId
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetName() *string {
	return s.Name
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetOffsetAcceptAmount() *int64 {
	return s.OffsetAcceptAmount
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetOffsetCostAmount() *int64 {
	return s.OffsetCostAmount
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetOpId() *string {
	return s.OpId
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetOriginalAmount() *int64 {
	return s.OriginalAmount
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetOutBizId() *string {
	return s.OutBizId
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetPresentAmount() *int64 {
	return s.PresentAmount
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetStatus() *int32 {
	return s.Status
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetType() *int32 {
	return s.Type
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) GetUserNick() *string {
	return s.UserNick
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetBillCycle(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.BillCycle = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetBillId(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.BillId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetBizTime(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.BizTime = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetBizType(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.BizType = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetCanInvoiceAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.CanInvoiceAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetGmtCreate(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.GmtCreate = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetGmtModified(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.GmtModified = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetId(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.Id = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetInvoicedAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.InvoicedAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetItemId(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.ItemId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetName(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.Name = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOffsetAcceptAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OffsetAcceptAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOffsetCostAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OffsetCostAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOpId(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OpId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOriginalAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OriginalAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetOutBizId(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.OutBizId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetPresentAmount(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.PresentAmount = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetStatus(v int32) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.Status = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetType(v int32) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.Type = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetUserId(v int64) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.UserId = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) SetUserNick(v string) *QueryEvaluateListResponseBodyDataEvaluateListEvaluate {
	s.UserNick = &v
	return s
}

func (s *QueryEvaluateListResponseBodyDataEvaluateListEvaluate) Validate() error {
	return dara.Validate(s)
}
