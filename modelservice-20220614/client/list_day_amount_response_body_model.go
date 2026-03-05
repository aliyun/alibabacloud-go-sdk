// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDayAmountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDayAmounts(v []*ListDayAmountResponseBodyDayAmounts) *ListDayAmountResponseBody
	GetDayAmounts() []*ListDayAmountResponseBodyDayAmounts
	SetRequestId(v string) *ListDayAmountResponseBody
	GetRequestId() *string
}

type ListDayAmountResponseBody struct {
	DayAmounts []*ListDayAmountResponseBodyDayAmounts `json:"DayAmounts,omitempty" xml:"DayAmounts,omitempty" type:"Repeated"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDayAmountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDayAmountResponseBody) GoString() string {
	return s.String()
}

func (s *ListDayAmountResponseBody) GetDayAmounts() []*ListDayAmountResponseBodyDayAmounts {
	return s.DayAmounts
}

func (s *ListDayAmountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDayAmountResponseBody) SetDayAmounts(v []*ListDayAmountResponseBodyDayAmounts) *ListDayAmountResponseBody {
	s.DayAmounts = v
	return s
}

func (s *ListDayAmountResponseBody) SetRequestId(v string) *ListDayAmountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDayAmountResponseBody) Validate() error {
	if s.DayAmounts != nil {
		for _, item := range s.DayAmounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDayAmountResponseBodyDayAmounts struct {
	Amount *int32  `json:"Amount,omitempty" xml:"Amount,omitempty"`
	Date   *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s ListDayAmountResponseBodyDayAmounts) String() string {
	return dara.Prettify(s)
}

func (s ListDayAmountResponseBodyDayAmounts) GoString() string {
	return s.String()
}

func (s *ListDayAmountResponseBodyDayAmounts) GetAmount() *int32 {
	return s.Amount
}

func (s *ListDayAmountResponseBodyDayAmounts) GetDate() *string {
	return s.Date
}

func (s *ListDayAmountResponseBodyDayAmounts) SetAmount(v int32) *ListDayAmountResponseBodyDayAmounts {
	s.Amount = &v
	return s
}

func (s *ListDayAmountResponseBodyDayAmounts) SetDate(v string) *ListDayAmountResponseBodyDayAmounts {
	s.Date = &v
	return s
}

func (s *ListDayAmountResponseBodyDayAmounts) Validate() error {
	return dara.Validate(s)
}
