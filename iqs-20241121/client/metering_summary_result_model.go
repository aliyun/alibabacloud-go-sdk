// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeteringSummaryResult interface {
	dara.Model
	String() string
	GoString() string
	SetMeteringSummaryList(v []*MeteringSummaryResultMeteringSummaryList) *MeteringSummaryResult
	GetMeteringSummaryList() []*MeteringSummaryResultMeteringSummaryList
}

type MeteringSummaryResult struct {
	MeteringSummaryList []*MeteringSummaryResultMeteringSummaryList `json:"meteringSummaryList,omitempty" xml:"meteringSummaryList,omitempty" type:"Repeated"`
}

func (s MeteringSummaryResult) String() string {
	return dara.Prettify(s)
}

func (s MeteringSummaryResult) GoString() string {
	return s.String()
}

func (s *MeteringSummaryResult) GetMeteringSummaryList() []*MeteringSummaryResultMeteringSummaryList {
	return s.MeteringSummaryList
}

func (s *MeteringSummaryResult) SetMeteringSummaryList(v []*MeteringSummaryResultMeteringSummaryList) *MeteringSummaryResult {
	s.MeteringSummaryList = v
	return s
}

func (s *MeteringSummaryResult) Validate() error {
	if s.MeteringSummaryList != nil {
		for _, item := range s.MeteringSummaryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MeteringSummaryResultMeteringSummaryList struct {
	BillingFunctionItem *string `json:"billingFunctionItem,omitempty" xml:"billingFunctionItem,omitempty"`
	BillingItem         *string `json:"billingItem,omitempty" xml:"billingItem,omitempty"`
	FailCount           *int32  `json:"failCount,omitempty" xml:"failCount,omitempty"`
	SuccessCount        *int32  `json:"successCount,omitempty" xml:"successCount,omitempty"`
}

func (s MeteringSummaryResultMeteringSummaryList) String() string {
	return dara.Prettify(s)
}

func (s MeteringSummaryResultMeteringSummaryList) GoString() string {
	return s.String()
}

func (s *MeteringSummaryResultMeteringSummaryList) GetBillingFunctionItem() *string {
	return s.BillingFunctionItem
}

func (s *MeteringSummaryResultMeteringSummaryList) GetBillingItem() *string {
	return s.BillingItem
}

func (s *MeteringSummaryResultMeteringSummaryList) GetFailCount() *int32 {
	return s.FailCount
}

func (s *MeteringSummaryResultMeteringSummaryList) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *MeteringSummaryResultMeteringSummaryList) SetBillingFunctionItem(v string) *MeteringSummaryResultMeteringSummaryList {
	s.BillingFunctionItem = &v
	return s
}

func (s *MeteringSummaryResultMeteringSummaryList) SetBillingItem(v string) *MeteringSummaryResultMeteringSummaryList {
	s.BillingItem = &v
	return s
}

func (s *MeteringSummaryResultMeteringSummaryList) SetFailCount(v int32) *MeteringSummaryResultMeteringSummaryList {
	s.FailCount = &v
	return s
}

func (s *MeteringSummaryResultMeteringSummaryList) SetSuccessCount(v int32) *MeteringSummaryResultMeteringSummaryList {
	s.SuccessCount = &v
	return s
}

func (s *MeteringSummaryResultMeteringSummaryList) Validate() error {
	return dara.Validate(s)
}
