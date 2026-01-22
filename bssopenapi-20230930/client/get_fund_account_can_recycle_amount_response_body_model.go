// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFundAccountCanRecycleAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableAmount(v string) *GetFundAccountCanRecycleAmountResponseBody
	GetAvailableAmount() *string
	SetCurrency(v string) *GetFundAccountCanRecycleAmountResponseBody
	GetCurrency() *string
	SetMetadata(v interface{}) *GetFundAccountCanRecycleAmountResponseBody
	GetMetadata() interface{}
	SetRecycleFromFundAccountId(v string) *GetFundAccountCanRecycleAmountResponseBody
	GetRecycleFromFundAccountId() *string
	SetRecycleToFundAccountList(v []*GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) *GetFundAccountCanRecycleAmountResponseBody
	GetRecycleToFundAccountList() []*GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList
	SetRequestId(v string) *GetFundAccountCanRecycleAmountResponseBody
	GetRequestId() *string
	SetTransferAmount(v string) *GetFundAccountCanRecycleAmountResponseBody
	GetTransferAmount() *string
}

type GetFundAccountCanRecycleAmountResponseBody struct {
	// example:
	//
	// 300
	AvailableAmount *string `json:"AvailableAmount,omitempty" xml:"AvailableAmount,omitempty"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 1232122132
	RecycleFromFundAccountId *string                                                               `json:"RecycleFromFundAccountId,omitempty" xml:"RecycleFromFundAccountId,omitempty"`
	RecycleToFundAccountList []*GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList `json:"RecycleToFundAccountList,omitempty" xml:"RecycleToFundAccountList,omitempty" type:"Repeated"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 200
	TransferAmount *string `json:"TransferAmount,omitempty" xml:"TransferAmount,omitempty"`
}

func (s GetFundAccountCanRecycleAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanRecycleAmountResponseBody) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanRecycleAmountResponseBody) GetAvailableAmount() *string {
	return s.AvailableAmount
}

func (s *GetFundAccountCanRecycleAmountResponseBody) GetCurrency() *string {
	return s.Currency
}

func (s *GetFundAccountCanRecycleAmountResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *GetFundAccountCanRecycleAmountResponseBody) GetRecycleFromFundAccountId() *string {
	return s.RecycleFromFundAccountId
}

func (s *GetFundAccountCanRecycleAmountResponseBody) GetRecycleToFundAccountList() []*GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	return s.RecycleToFundAccountList
}

func (s *GetFundAccountCanRecycleAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFundAccountCanRecycleAmountResponseBody) GetTransferAmount() *string {
	return s.TransferAmount
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetAvailableAmount(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.AvailableAmount = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetCurrency(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.Currency = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetMetadata(v interface{}) *GetFundAccountCanRecycleAmountResponseBody {
	s.Metadata = v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetRecycleFromFundAccountId(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.RecycleFromFundAccountId = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetRecycleToFundAccountList(v []*GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) *GetFundAccountCanRecycleAmountResponseBody {
	s.RecycleToFundAccountList = v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetRequestId(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) SetTransferAmount(v string) *GetFundAccountCanRecycleAmountResponseBody {
	s.TransferAmount = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBody) Validate() error {
	if s.RecycleToFundAccountList != nil {
		for _, item := range s.RecycleToFundAccountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList struct {
	// example:
	//
	// 122323121
	FundAccountId *string `json:"FundAccountId,omitempty" xml:"FundAccountId,omitempty"`
	// example:
	//
	// 云某的账户
	FundAccountName *string `json:"FundAccountName,omitempty" xml:"FundAccountName,omitempty"`
	// example:
	//
	// 183221321
	FundAccountOwnerAccountId *string `json:"FundAccountOwnerAccountId,omitempty" xml:"FundAccountOwnerAccountId,omitempty"`
	// example:
	//
	// 120
	MaxRecyclableAmount *string `json:"MaxRecyclableAmount,omitempty" xml:"MaxRecyclableAmount,omitempty"`
	// example:
	//
	// 120
	OriginalTransferRemainAmount *string `json:"OriginalTransferRemainAmount,omitempty" xml:"OriginalTransferRemainAmount,omitempty"`
}

func (s GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) String() string {
	return dara.Prettify(s)
}

func (s GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) GoString() string {
	return s.String()
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) GetFundAccountId() *string {
	return s.FundAccountId
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) GetFundAccountName() *string {
	return s.FundAccountName
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) GetFundAccountOwnerAccountId() *string {
	return s.FundAccountOwnerAccountId
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) GetMaxRecyclableAmount() *string {
	return s.MaxRecyclableAmount
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) GetOriginalTransferRemainAmount() *string {
	return s.OriginalTransferRemainAmount
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetFundAccountId(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.FundAccountId = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetFundAccountName(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.FundAccountName = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetFundAccountOwnerAccountId(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.FundAccountOwnerAccountId = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetMaxRecyclableAmount(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.MaxRecyclableAmount = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) SetOriginalTransferRemainAmount(v string) *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList {
	s.OriginalTransferRemainAmount = &v
	return s
}

func (s *GetFundAccountCanRecycleAmountResponseBodyRecycleToFundAccountList) Validate() error {
	return dara.Validate(s)
}
