// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReversedDeductionHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryReversedDeductionHistoryResponseBody
	GetCode() *string
	SetData(v []*QueryReversedDeductionHistoryResponseBodyData) *QueryReversedDeductionHistoryResponseBody
	GetData() []*QueryReversedDeductionHistoryResponseBodyData
	SetMessage(v string) *QueryReversedDeductionHistoryResponseBody
	GetMessage() *string
	SetPageNo(v int32) *QueryReversedDeductionHistoryResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *QueryReversedDeductionHistoryResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryReversedDeductionHistoryResponseBody
	GetRequestId() *string
	SetTotal(v int32) *QueryReversedDeductionHistoryResponseBody
	GetTotal() *int32
}

type QueryReversedDeductionHistoryResponseBody struct {
	// example:
	//
	// 200
	Code *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*QueryReversedDeductionHistoryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// asda1231as
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryReversedDeductionHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryReversedDeductionHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReversedDeductionHistoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryReversedDeductionHistoryResponseBody) GetData() []*QueryReversedDeductionHistoryResponseBodyData {
	return s.Data
}

func (s *QueryReversedDeductionHistoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryReversedDeductionHistoryResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryReversedDeductionHistoryResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryReversedDeductionHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryReversedDeductionHistoryResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *QueryReversedDeductionHistoryResponseBody) SetCode(v string) *QueryReversedDeductionHistoryResponseBody {
	s.Code = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBody) SetData(v []*QueryReversedDeductionHistoryResponseBodyData) *QueryReversedDeductionHistoryResponseBody {
	s.Data = v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBody) SetMessage(v string) *QueryReversedDeductionHistoryResponseBody {
	s.Message = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBody) SetPageNo(v int32) *QueryReversedDeductionHistoryResponseBody {
	s.PageNo = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBody) SetPageSize(v int32) *QueryReversedDeductionHistoryResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBody) SetRequestId(v string) *QueryReversedDeductionHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBody) SetTotal(v int32) *QueryReversedDeductionHistoryResponseBody {
	s.Total = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBody) Validate() error {
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

type QueryReversedDeductionHistoryResponseBodyData struct {
	// example:
	//
	// 1.00
	OffsetAmount *string `json:"OffsetAmount,omitempty" xml:"OffsetAmount,omitempty"`
	// example:
	//
	// ACPN
	OperationSubmitType *string `json:"OperationSubmitType,omitempty" xml:"OperationSubmitType,omitempty"`
	// example:
	//
	// 2024-11-01 10:22:11 UTC+8
	OperationTime *string `json:"OperationTime,omitempty" xml:"OperationTime,omitempty"`
	// example:
	//
	// 132
	OperationUid *string `json:"OperationUid,omitempty" xml:"OperationUid,omitempty"`
}

func (s QueryReversedDeductionHistoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryReversedDeductionHistoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryReversedDeductionHistoryResponseBodyData) GetOffsetAmount() *string {
	return s.OffsetAmount
}

func (s *QueryReversedDeductionHistoryResponseBodyData) GetOperationSubmitType() *string {
	return s.OperationSubmitType
}

func (s *QueryReversedDeductionHistoryResponseBodyData) GetOperationTime() *string {
	return s.OperationTime
}

func (s *QueryReversedDeductionHistoryResponseBodyData) GetOperationUid() *string {
	return s.OperationUid
}

func (s *QueryReversedDeductionHistoryResponseBodyData) SetOffsetAmount(v string) *QueryReversedDeductionHistoryResponseBodyData {
	s.OffsetAmount = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBodyData) SetOperationSubmitType(v string) *QueryReversedDeductionHistoryResponseBodyData {
	s.OperationSubmitType = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBodyData) SetOperationTime(v string) *QueryReversedDeductionHistoryResponseBodyData {
	s.OperationTime = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBodyData) SetOperationUid(v string) *QueryReversedDeductionHistoryResponseBodyData {
	s.OperationUid = &v
	return s
}

func (s *QueryReversedDeductionHistoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
