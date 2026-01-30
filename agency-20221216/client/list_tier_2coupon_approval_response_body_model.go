// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTier2CouponApprovalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListTier2CouponApprovalResponseBody
	GetCode() *string
	SetData(v []*ListTier2CouponApprovalResponseBodyData) *ListTier2CouponApprovalResponseBody
	GetData() []*ListTier2CouponApprovalResponseBodyData
	SetMessage(v string) *ListTier2CouponApprovalResponseBody
	GetMessage() *string
	SetPageNo(v int32) *ListTier2CouponApprovalResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *ListTier2CouponApprovalResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListTier2CouponApprovalResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListTier2CouponApprovalResponseBody
	GetTotal() *int32
}

type ListTier2CouponApprovalResponseBody struct {
	// example:
	//
	// 200
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListTier2CouponApprovalResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListTier2CouponApprovalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTier2CouponApprovalResponseBody) GoString() string {
	return s.String()
}

func (s *ListTier2CouponApprovalResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListTier2CouponApprovalResponseBody) GetData() []*ListTier2CouponApprovalResponseBodyData {
	return s.Data
}

func (s *ListTier2CouponApprovalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTier2CouponApprovalResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListTier2CouponApprovalResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTier2CouponApprovalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTier2CouponApprovalResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListTier2CouponApprovalResponseBody) SetCode(v string) *ListTier2CouponApprovalResponseBody {
	s.Code = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBody) SetData(v []*ListTier2CouponApprovalResponseBodyData) *ListTier2CouponApprovalResponseBody {
	s.Data = v
	return s
}

func (s *ListTier2CouponApprovalResponseBody) SetMessage(v string) *ListTier2CouponApprovalResponseBody {
	s.Message = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBody) SetPageNo(v int32) *ListTier2CouponApprovalResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBody) SetPageSize(v int32) *ListTier2CouponApprovalResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBody) SetRequestId(v string) *ListTier2CouponApprovalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBody) SetTotal(v int32) *ListTier2CouponApprovalResponseBody {
	s.Total = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBody) Validate() error {
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

type ListTier2CouponApprovalResponseBodyData struct {
	// example:
	//
	// d54ca949-9b88-4514-add3-c6029c4027f4
	ApplicationSheetId *string `json:"ApplicationSheetId,omitempty" xml:"ApplicationSheetId,omitempty"`
	// example:
	//
	// 2026-01-21 11:24
	ApplicationTime *string `json:"ApplicationTime,omitempty" xml:"ApplicationTime,omitempty"`
	// example:
	//
	// 1
	ApprovalStatus *string `json:"ApprovalStatus,omitempty" xml:"ApprovalStatus,omitempty"`
	// example:
	//
	// test account
	T2PartnerName *string `json:"T2PartnerName,omitempty" xml:"T2PartnerName,omitempty"`
	// example:
	//
	// 5248516956402795
	T2PartnerUid *int64 `json:"T2PartnerUid,omitempty" xml:"T2PartnerUid,omitempty"`
	// example:
	//
	// 200
	TotalAmount *string `json:"TotalAmount,omitempty" xml:"TotalAmount,omitempty"`
}

func (s ListTier2CouponApprovalResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTier2CouponApprovalResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTier2CouponApprovalResponseBodyData) GetApplicationSheetId() *string {
	return s.ApplicationSheetId
}

func (s *ListTier2CouponApprovalResponseBodyData) GetApplicationTime() *string {
	return s.ApplicationTime
}

func (s *ListTier2CouponApprovalResponseBodyData) GetApprovalStatus() *string {
	return s.ApprovalStatus
}

func (s *ListTier2CouponApprovalResponseBodyData) GetT2PartnerName() *string {
	return s.T2PartnerName
}

func (s *ListTier2CouponApprovalResponseBodyData) GetT2PartnerUid() *int64 {
	return s.T2PartnerUid
}

func (s *ListTier2CouponApprovalResponseBodyData) GetTotalAmount() *string {
	return s.TotalAmount
}

func (s *ListTier2CouponApprovalResponseBodyData) SetApplicationSheetId(v string) *ListTier2CouponApprovalResponseBodyData {
	s.ApplicationSheetId = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBodyData) SetApplicationTime(v string) *ListTier2CouponApprovalResponseBodyData {
	s.ApplicationTime = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBodyData) SetApprovalStatus(v string) *ListTier2CouponApprovalResponseBodyData {
	s.ApprovalStatus = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBodyData) SetT2PartnerName(v string) *ListTier2CouponApprovalResponseBodyData {
	s.T2PartnerName = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBodyData) SetT2PartnerUid(v int64) *ListTier2CouponApprovalResponseBodyData {
	s.T2PartnerUid = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBodyData) SetTotalAmount(v string) *ListTier2CouponApprovalResponseBodyData {
	s.TotalAmount = &v
	return s
}

func (s *ListTier2CouponApprovalResponseBodyData) Validate() error {
	return dara.Validate(s)
}
