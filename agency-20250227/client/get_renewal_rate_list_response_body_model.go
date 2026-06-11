// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenewalRateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRenewalRateListResponseBody
	GetCode() *string
	SetData(v []*GetRenewalRateListResponseBodyData) *GetRenewalRateListResponseBody
	GetData() []*GetRenewalRateListResponseBodyData
	SetMessage(v string) *GetRenewalRateListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRenewalRateListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetRenewalRateListResponseBody
	GetSuccess() *bool
}

type GetRenewalRateListResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data
	Data []*GetRenewalRateListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Message
	//
	// example:
	//
	// 成功
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9C14ADFE-DF0A-54D4-8BD5-45D0839246B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the invocation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRenewalRateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRenewalRateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetRenewalRateListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRenewalRateListResponseBody) GetData() []*GetRenewalRateListResponseBodyData {
	return s.Data
}

func (s *GetRenewalRateListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRenewalRateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRenewalRateListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetRenewalRateListResponseBody) SetCode(v string) *GetRenewalRateListResponseBody {
	s.Code = &v
	return s
}

func (s *GetRenewalRateListResponseBody) SetData(v []*GetRenewalRateListResponseBodyData) *GetRenewalRateListResponseBody {
	s.Data = v
	return s
}

func (s *GetRenewalRateListResponseBody) SetMessage(v string) *GetRenewalRateListResponseBody {
	s.Message = &v
	return s
}

func (s *GetRenewalRateListResponseBody) SetRequestId(v string) *GetRenewalRateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRenewalRateListResponseBody) SetSuccess(v bool) *GetRenewalRateListResponseBody {
	s.Success = &v
	return s
}

func (s *GetRenewalRateListResponseBody) Validate() error {
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

type GetRenewalRateListResponseBodyData struct {
	// Adjusted amount due for Gold customer renewal
	//
	// example:
	//
	// 100
	CustomerAdjustedRenewalAmountDue *float64 `json:"CustomerAdjustedRenewalAmountDue,omitempty" xml:"CustomerAdjustedRenewalAmountDue,omitempty"`
	// Rewind amount for new customer acquisition upgrade
	//
	// example:
	//
	// 100
	CustomerOtherBillAmount *float64 `json:"CustomerOtherBillAmount,omitempty" xml:"CustomerOtherBillAmount,omitempty"`
	// Amount due for Gold customer renewal
	//
	// example:
	//
	// 100
	FinalCustomerRenewalAmountDue *float64 `json:"FinalCustomerRenewalAmountDue,omitempty" xml:"FinalCustomerRenewalAmountDue,omitempty"`
	// Gold customer commission renewal rate
	//
	// example:
	//
	// 0.9
	FinalCustomerRenewalRate *float64 `json:"FinalCustomerRenewalRate,omitempty" xml:"FinalCustomerRenewalRate,omitempty"`
	// Renewed amount from customer acquisition
	//
	// example:
	//
	// 100
	FinalCustomerRenewedAmount *float64 `json:"FinalCustomerRenewedAmount,omitempty" xml:"FinalCustomerRenewedAmount,omitempty"`
	// Total rewind amount for new upgrades
	//
	// example:
	//
	// 100
	FinalOtherBillAmount *float64 `json:"FinalOtherBillAmount,omitempty" xml:"FinalOtherBillAmount,omitempty"`
	// Final amount due for renewal
	//
	// example:
	//
	// 100
	FinalRenewalAmountDue *float64 `json:"FinalRenewalAmountDue,omitempty" xml:"FinalRenewalAmountDue,omitempty"`
	// Final commission renewal rate
	//
	// example:
	//
	// 0.9
	FinalRenewalRate *float64 `json:"FinalRenewalRate,omitempty" xml:"FinalRenewalRate,omitempty"`
	// Final renewed amount
	//
	// example:
	//
	// 100
	FinalRenewedAmount *float64 `json:"FinalRenewedAmount,omitempty" xml:"FinalRenewedAmount,omitempty"`
	// Renewal amount due from channel partner acquisition
	//
	// example:
	//
	// 100
	FinalSubPartnerRenewalAmountDue *float64 `json:"FinalSubPartnerRenewalAmountDue,omitempty" xml:"FinalSubPartnerRenewalAmountDue,omitempty"`
	// Channel expansion commission renewal rate
	//
	// example:
	//
	// 0.85
	FinalSubPartnerRenewalRate *float64 `json:"FinalSubPartnerRenewalRate,omitempty" xml:"FinalSubPartnerRenewalRate,omitempty"`
	// Renewed amount from channel partner acquisition
	//
	// example:
	//
	// 100
	FinalSubPartnerRenewedAmount *float64 `json:"FinalSubPartnerRenewedAmount,omitempty" xml:"FinalSubPartnerRenewedAmount,omitempty"`
	// Fiscal year and quarter
	//
	// example:
	//
	// 2025Q4
	FiscalYearAndQuarter *string `json:"FiscalYearAndQuarter,omitempty" xml:"FiscalYearAndQuarter,omitempty"`
	// Partner PID
	//
	// example:
	//
	// P123423453
	MasterPid *string `json:"MasterPid,omitempty" xml:"MasterPid,omitempty"`
	// Partner PID name
	//
	// example:
	//
	// XX有限公司
	MasterPidName *string `json:"MasterPidName,omitempty" xml:"MasterPidName,omitempty"`
	// Gold customer commission renewal rate including special approval
	//
	// example:
	//
	// 0.7
	SpecialCustomerRenewRatio *float64 `json:"SpecialCustomerRenewRatio,omitempty" xml:"SpecialCustomerRenewRatio,omitempty"`
	// Renewal amount due including specially approved customer acquisition
	//
	// example:
	//
	// 100
	SpecialCustomerRenewalAmountDue *float64 `json:"SpecialCustomerRenewalAmountDue,omitempty" xml:"SpecialCustomerRenewalAmountDue,omitempty"`
	// Renewed amount including specially approved customer acquisition
	//
	// example:
	//
	// 100
	SpecialCustomerRenewedAmount *float64 `json:"SpecialCustomerRenewedAmount,omitempty" xml:"SpecialCustomerRenewedAmount,omitempty"`
	// Renewal rate including specially approved final quarter commission
	//
	// example:
	//
	// 0.7
	SpecialFinalRenewRatio *float64 `json:"SpecialFinalRenewRatio,omitempty" xml:"SpecialFinalRenewRatio,omitempty"`
	// Renewal amount due for final quarterly commission including special approvals
	//
	// example:
	//
	// 100
	SpecialFinalRenewalAmountDue *float64 `json:"SpecialFinalRenewalAmountDue,omitempty" xml:"SpecialFinalRenewalAmountDue,omitempty"`
	// Renewed amount including specially approved final quarter commission
	//
	// example:
	//
	// 100
	SpecialFinalRenewedAmount *float64 `json:"SpecialFinalRenewedAmount,omitempty" xml:"SpecialFinalRenewedAmount,omitempty"`
	// Channel expansion commission renewal rate including special approval
	//
	// example:
	//
	// 0.8
	SpecialSubPartnerRenewRatio *float64 `json:"SpecialSubPartnerRenewRatio,omitempty" xml:"SpecialSubPartnerRenewRatio,omitempty"`
	// Amount due for renewal including specially approved channel partner
	//
	// example:
	//
	// 100
	SpecialSubPartnerRenewalAmountDue *float64 `json:"SpecialSubPartnerRenewalAmountDue,omitempty" xml:"SpecialSubPartnerRenewalAmountDue,omitempty"`
	// Renewed amount including specially approved channel partner renewals
	//
	// example:
	//
	// 100
	SpecialSubPartnerRenewedAmount *float64 `json:"SpecialSubPartnerRenewedAmount,omitempty" xml:"SpecialSubPartnerRenewedAmount,omitempty"`
	// Adjusted amount due for channel partner renewal
	//
	// example:
	//
	// 100
	SubPartnerAdjustedRenewalAmountDue *float64 `json:"SubPartnerAdjustedRenewalAmountDue,omitempty" xml:"SubPartnerAdjustedRenewalAmountDue,omitempty"`
	// Channel partner new upgrade rewind amount
	//
	// example:
	//
	// 100
	SubPartnerOtherBillAmount *float64 `json:"SubPartnerOtherBillAmount,omitempty" xml:"SubPartnerOtherBillAmount,omitempty"`
}

func (s GetRenewalRateListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRenewalRateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRenewalRateListResponseBodyData) GetCustomerAdjustedRenewalAmountDue() *float64 {
	return s.CustomerAdjustedRenewalAmountDue
}

func (s *GetRenewalRateListResponseBodyData) GetCustomerOtherBillAmount() *float64 {
	return s.CustomerOtherBillAmount
}

func (s *GetRenewalRateListResponseBodyData) GetFinalCustomerRenewalAmountDue() *float64 {
	return s.FinalCustomerRenewalAmountDue
}

func (s *GetRenewalRateListResponseBodyData) GetFinalCustomerRenewalRate() *float64 {
	return s.FinalCustomerRenewalRate
}

func (s *GetRenewalRateListResponseBodyData) GetFinalCustomerRenewedAmount() *float64 {
	return s.FinalCustomerRenewedAmount
}

func (s *GetRenewalRateListResponseBodyData) GetFinalOtherBillAmount() *float64 {
	return s.FinalOtherBillAmount
}

func (s *GetRenewalRateListResponseBodyData) GetFinalRenewalAmountDue() *float64 {
	return s.FinalRenewalAmountDue
}

func (s *GetRenewalRateListResponseBodyData) GetFinalRenewalRate() *float64 {
	return s.FinalRenewalRate
}

func (s *GetRenewalRateListResponseBodyData) GetFinalRenewedAmount() *float64 {
	return s.FinalRenewedAmount
}

func (s *GetRenewalRateListResponseBodyData) GetFinalSubPartnerRenewalAmountDue() *float64 {
	return s.FinalSubPartnerRenewalAmountDue
}

func (s *GetRenewalRateListResponseBodyData) GetFinalSubPartnerRenewalRate() *float64 {
	return s.FinalSubPartnerRenewalRate
}

func (s *GetRenewalRateListResponseBodyData) GetFinalSubPartnerRenewedAmount() *float64 {
	return s.FinalSubPartnerRenewedAmount
}

func (s *GetRenewalRateListResponseBodyData) GetFiscalYearAndQuarter() *string {
	return s.FiscalYearAndQuarter
}

func (s *GetRenewalRateListResponseBodyData) GetMasterPid() *string {
	return s.MasterPid
}

func (s *GetRenewalRateListResponseBodyData) GetMasterPidName() *string {
	return s.MasterPidName
}

func (s *GetRenewalRateListResponseBodyData) GetSpecialCustomerRenewRatio() *float64 {
	return s.SpecialCustomerRenewRatio
}

func (s *GetRenewalRateListResponseBodyData) GetSpecialCustomerRenewalAmountDue() *float64 {
	return s.SpecialCustomerRenewalAmountDue
}

func (s *GetRenewalRateListResponseBodyData) GetSpecialCustomerRenewedAmount() *float64 {
	return s.SpecialCustomerRenewedAmount
}

func (s *GetRenewalRateListResponseBodyData) GetSpecialFinalRenewRatio() *float64 {
	return s.SpecialFinalRenewRatio
}

func (s *GetRenewalRateListResponseBodyData) GetSpecialFinalRenewalAmountDue() *float64 {
	return s.SpecialFinalRenewalAmountDue
}

func (s *GetRenewalRateListResponseBodyData) GetSpecialFinalRenewedAmount() *float64 {
	return s.SpecialFinalRenewedAmount
}

func (s *GetRenewalRateListResponseBodyData) GetSpecialSubPartnerRenewRatio() *float64 {
	return s.SpecialSubPartnerRenewRatio
}

func (s *GetRenewalRateListResponseBodyData) GetSpecialSubPartnerRenewalAmountDue() *float64 {
	return s.SpecialSubPartnerRenewalAmountDue
}

func (s *GetRenewalRateListResponseBodyData) GetSpecialSubPartnerRenewedAmount() *float64 {
	return s.SpecialSubPartnerRenewedAmount
}

func (s *GetRenewalRateListResponseBodyData) GetSubPartnerAdjustedRenewalAmountDue() *float64 {
	return s.SubPartnerAdjustedRenewalAmountDue
}

func (s *GetRenewalRateListResponseBodyData) GetSubPartnerOtherBillAmount() *float64 {
	return s.SubPartnerOtherBillAmount
}

func (s *GetRenewalRateListResponseBodyData) SetCustomerAdjustedRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.CustomerAdjustedRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetCustomerOtherBillAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.CustomerOtherBillAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalCustomerRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalCustomerRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalCustomerRenewalRate(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalCustomerRenewalRate = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalCustomerRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalCustomerRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalOtherBillAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalOtherBillAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalRenewalRate(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalRenewalRate = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalSubPartnerRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalSubPartnerRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalSubPartnerRenewalRate(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalSubPartnerRenewalRate = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFinalSubPartnerRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.FinalSubPartnerRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetFiscalYearAndQuarter(v string) *GetRenewalRateListResponseBodyData {
	s.FiscalYearAndQuarter = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetMasterPid(v string) *GetRenewalRateListResponseBodyData {
	s.MasterPid = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetMasterPidName(v string) *GetRenewalRateListResponseBodyData {
	s.MasterPidName = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialCustomerRenewRatio(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialCustomerRenewRatio = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialCustomerRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialCustomerRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialCustomerRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialCustomerRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialFinalRenewRatio(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialFinalRenewRatio = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialFinalRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialFinalRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialFinalRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialFinalRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialSubPartnerRenewRatio(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialSubPartnerRenewRatio = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialSubPartnerRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialSubPartnerRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSpecialSubPartnerRenewedAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.SpecialSubPartnerRenewedAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSubPartnerAdjustedRenewalAmountDue(v float64) *GetRenewalRateListResponseBodyData {
	s.SubPartnerAdjustedRenewalAmountDue = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) SetSubPartnerOtherBillAmount(v float64) *GetRenewalRateListResponseBodyData {
	s.SubPartnerOtherBillAmount = &v
	return s
}

func (s *GetRenewalRateListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
