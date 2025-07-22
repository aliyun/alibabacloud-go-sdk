// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountTransactionDetailsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetBillNumber() *string
	SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetChannelTransactionNumber() *string
	SetCurrentPage(v int32) *GetFundAccountTransactionDetailsShrinkRequest
	GetCurrentPage() *int32
	SetEndTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest
	GetEndTime() *int64
	SetFundAccountId(v int64) *GetFundAccountTransactionDetailsShrinkRequest
	GetFundAccountId() *int64
	SetPageSize(v int32) *GetFundAccountTransactionDetailsShrinkRequest
	GetPageSize() *int32
	SetStartTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest
	GetStartTime() *int64
	SetTransactionChannelListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionChannelListShrink() *string
	SetTransactionDirection(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionDirection() *string
	SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionNumber() *int64
	SetTransactionType(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionType() *string
	SetTransactionTypeListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest
	GetTransactionTypeListShrink() *string
}

type GetFundAccountTransactionDetailsShrinkRequest struct {
	// example:
	//
	// 2023212312321
	BillNumber *string `json:"BillNumber,omitempty" xml:"BillNumber,omitempty"`
	// example:
	//
	// 20250312334312322
	ChannelTransactionNumber *string `json:"ChannelTransactionNumber,omitempty" xml:"ChannelTransactionNumber,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 1735664561000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 123221232
	FundAccountId *int64 `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1735664461000
	StartTime                    *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TransactionChannelListShrink *string `json:"TransactionChannelList,omitempty" xml:"TransactionChannelList,omitempty"`
	// example:
	//
	// IN
	TransactionDirection *string `json:"TransactionDirection,omitempty" xml:"TransactionDirection,omitempty"`
	// example:
	//
	// 543231231
	TransactionNumber *int64 `json:"TransactionNumber,omitempty" xml:"TransactionNumber,omitempty"`
	// example:
	//
	// CHARGE
	TransactionType           *string `json:"TransactionType,omitempty" xml:"TransactionType,omitempty"`
	TransactionTypeListShrink *string `json:"TransactionTypeList,omitempty" xml:"TransactionTypeList,omitempty"`
}

func (s GetFundAccountTransactionDetailsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountTransactionDetailsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetBillNumber() *string {
	return s.BillNumber
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetChannelTransactionNumber() *string {
	return s.ChannelTransactionNumber
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetFundAccountId() *int64 {
	return s.FundAccountId
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionChannelListShrink() *string {
	return s.TransactionChannelListShrink
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionDirection() *string {
	return s.TransactionDirection
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionNumber() *int64 {
	return s.TransactionNumber
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionType() *string {
	return s.TransactionType
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) GetTransactionTypeListShrink() *string {
	return s.TransactionTypeListShrink
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetBillNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.BillNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetChannelTransactionNumber(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.ChannelTransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetCurrentPage(v int32) *GetFundAccountTransactionDetailsShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetEndTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetFundAccountId(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetPageSize(v int32) *GetFundAccountTransactionDetailsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetStartTime(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionChannelListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionChannelListShrink = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionDirection(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionDirection = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionNumber(v int64) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionNumber = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionType(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionType = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) SetTransactionTypeListShrink(v string) *GetFundAccountTransactionDetailsShrinkRequest {
	s.TransactionTypeListShrink = &v
	return s
}

func (s *GetFundAccountTransactionDetailsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
