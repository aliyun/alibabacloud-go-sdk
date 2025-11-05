// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerQuotaRecordListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CustomerQuotaRecordListResponseBody
	GetCode() *string
	SetData(v []*CustomerQuotaRecordListResponseBodyData) *CustomerQuotaRecordListResponseBody
	GetData() []*CustomerQuotaRecordListResponseBodyData
	SetMsg(v string) *CustomerQuotaRecordListResponseBody
	GetMsg() *string
	SetPageNo(v int32) *CustomerQuotaRecordListResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *CustomerQuotaRecordListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *CustomerQuotaRecordListResponseBody
	GetRequestId() *string
	SetTotal(v int32) *CustomerQuotaRecordListResponseBody
	GetTotal() *int32
}

type CustomerQuotaRecordListResponseBody struct {
	// Status code of returning result, 200 means success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Listed data of returning result
	Data []*CustomerQuotaRecordListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of returning data
	//
	// example:
	//
	// SUCCESS
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Current page number
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Record number on each page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID of request
	//
	// example:
	//
	// 2103a0ae16849855284594613d874e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total volume
	//
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s CustomerQuotaRecordListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CustomerQuotaRecordListResponseBody) GoString() string {
	return s.String()
}

func (s *CustomerQuotaRecordListResponseBody) GetCode() *string {
	return s.Code
}

func (s *CustomerQuotaRecordListResponseBody) GetData() []*CustomerQuotaRecordListResponseBodyData {
	return s.Data
}

func (s *CustomerQuotaRecordListResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *CustomerQuotaRecordListResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *CustomerQuotaRecordListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CustomerQuotaRecordListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CustomerQuotaRecordListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *CustomerQuotaRecordListResponseBody) SetCode(v string) *CustomerQuotaRecordListResponseBody {
	s.Code = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetData(v []*CustomerQuotaRecordListResponseBodyData) *CustomerQuotaRecordListResponseBody {
	s.Data = v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetMsg(v string) *CustomerQuotaRecordListResponseBody {
	s.Msg = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetPageNo(v int32) *CustomerQuotaRecordListResponseBody {
	s.PageNo = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetPageSize(v int32) *CustomerQuotaRecordListResponseBody {
	s.PageSize = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetRequestId(v string) *CustomerQuotaRecordListResponseBody {
	s.RequestId = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) SetTotal(v int32) *CustomerQuotaRecordListResponseBody {
	s.Total = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBody) Validate() error {
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

type CustomerQuotaRecordListResponseBodyData struct {
	// The way to submit the quota adjustment operation. API/ACPN
	//
	// example:
	//
	// ACPN
	OperationSubmitType *string `json:"OperationSubmitType,omitempty" xml:"OperationSubmitType,omitempty"`
	// The time of submit the quota adjustment operation.
	//
	// example:
	//
	// 2023-12-15 10:34:36 UTC+8
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// Operation Type Enum</br>
	//
	// all All types</br>
	//
	// quota_create Create quota</br>
	//
	// quota_amount_adjust Adjust the amount of quota</br>
	//
	// example:
	//
	// quota_amount_adjust
	OperationTypeCode *string `json:"OperationTypeCode,omitempty" xml:"OperationTypeCode,omitempty"`
	// The description of submitted quota adjustment operation.
	//
	// example:
	//
	// Quota Adjustment
	OperationTypeDesc *string `json:"OperationTypeDesc,omitempty" xml:"OperationTypeDesc,omitempty"`
	// The UID of operator(Partner\\"s UID).
	//
	// example:
	//
	// 5113766248601929
	OperationUid *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
	// Updated quota amount
	//
	// example:
	//
	// 121.00
	UpdateAfterAmount *string `json:"UpdateAfterAmount,omitempty" xml:"UpdateAfterAmount,omitempty"`
	// The difference amount between updated quota and original quota.
	//
	// example:
	//
	// -100.00
	UpdateAmount *string `json:"UpdateAmount,omitempty" xml:"UpdateAmount,omitempty"`
	// Original quota amount
	//
	// example:
	//
	// 221.00
	UpdateBeforeAmount *string `json:"UpdateBeforeAmount,omitempty" xml:"UpdateBeforeAmount,omitempty"`
}

func (s CustomerQuotaRecordListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CustomerQuotaRecordListResponseBodyData) GoString() string {
	return s.String()
}

func (s *CustomerQuotaRecordListResponseBodyData) GetOperationSubmitType() *string {
	return s.OperationSubmitType
}

func (s *CustomerQuotaRecordListResponseBodyData) GetOperationTime() *string {
	return s.OperationTime
}

func (s *CustomerQuotaRecordListResponseBodyData) GetOperationTypeCode() *string {
	return s.OperationTypeCode
}

func (s *CustomerQuotaRecordListResponseBodyData) GetOperationTypeDesc() *string {
	return s.OperationTypeDesc
}

func (s *CustomerQuotaRecordListResponseBodyData) GetOperationUid() *string {
	return s.OperationUid
}

func (s *CustomerQuotaRecordListResponseBodyData) GetUpdateAfterAmount() *string {
	return s.UpdateAfterAmount
}

func (s *CustomerQuotaRecordListResponseBodyData) GetUpdateAmount() *string {
	return s.UpdateAmount
}

func (s *CustomerQuotaRecordListResponseBodyData) GetUpdateBeforeAmount() *string {
	return s.UpdateBeforeAmount
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationSubmitType(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationSubmitType = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationTime(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationTime = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationTypeCode(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationTypeCode = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationTypeDesc(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationTypeDesc = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetOperationUid(v string) *CustomerQuotaRecordListResponseBodyData {
	s.OperationUid = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetUpdateAfterAmount(v string) *CustomerQuotaRecordListResponseBodyData {
	s.UpdateAfterAmount = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetUpdateAmount(v string) *CustomerQuotaRecordListResponseBodyData {
	s.UpdateAmount = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) SetUpdateBeforeAmount(v string) *CustomerQuotaRecordListResponseBodyData {
	s.UpdateBeforeAmount = &v
	return s
}

func (s *CustomerQuotaRecordListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
