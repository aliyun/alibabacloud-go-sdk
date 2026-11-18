// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIqsUsageResult interface {
	dara.Model
	String() string
	GoString() string
	SetRecords(v []*GetIqsUsageResultRecords) *GetIqsUsageResult
	GetRecords() []*GetIqsUsageResultRecords
}

type GetIqsUsageResult struct {
	// The daily call volume records.
	Records []*GetIqsUsageResultRecords `json:"records,omitempty" xml:"records,omitempty" type:"Repeated"`
}

func (s GetIqsUsageResult) String() string {
	return dara.Prettify(s)
}

func (s GetIqsUsageResult) GoString() string {
	return s.String()
}

func (s *GetIqsUsageResult) GetRecords() []*GetIqsUsageResultRecords {
	return s.Records
}

func (s *GetIqsUsageResult) SetRecords(v []*GetIqsUsageResultRecords) *GetIqsUsageResult {
	s.Records = v
	return s
}

func (s *GetIqsUsageResult) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetIqsUsageResultRecords struct {
	// The API name.
	//
	// example:
	//
	// UnifiedSearch
	Api *string `json:"api,omitempty" xml:"api,omitempty"`
	// The queries per second (QPS) for tiered billing.
	//
	// example:
	//
	// 30
	BillingQps *int32 `json:"billingQps,omitempty" xml:"billingQps,omitempty"`
	// The metering date in the YYYYMMDD format.
	//
	// example:
	//
	// 20250716
	Date *string `json:"date,omitempty" xml:"date,omitempty"`
	// The search engine type.
	//
	// example:
	//
	// Generic
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// The number of failed calls.
	//
	// example:
	//
	// 0
	FailedCalls *int32 `json:"failedCalls,omitempty" xml:"failedCalls,omitempty"`
	// The matched metering tier. Valid values: 1001, 1002, 1003, and 1004.
	//
	// example:
	//
	// 1001
	LadderType *string `json:"ladderType,omitempty" xml:"ladderType,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 138608930******0
	MainAccountId *string `json:"mainAccountId,omitempty" xml:"mainAccountId,omitempty"`
	// The sub-account that initiates the call (RAM user or API key identifier).
	//
	// example:
	//
	// 2193843763***9272
	SubAccountId *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	// The number of successful calls.
	//
	// example:
	//
	// 5
	SuccessCalls *int32 `json:"successCalls,omitempty" xml:"successCalls,omitempty"`
	// The total number of calls.
	//
	// example:
	//
	// 5
	TotalBillStatCalls *int32 `json:"totalBillStatCalls,omitempty" xml:"totalBillStatCalls,omitempty"`
	// The total number of calls.
	//
	// example:
	//
	// 5
	TotalCalls *int32 `json:"totalCalls,omitempty" xml:"totalCalls,omitempty"`
	// The value-added package usage for calling enhanced APIs.
	//
	// example:
	//
	// 5
	ValueAddedAdvanced *int32 `json:"valueAddedAdvanced,omitempty" xml:"valueAddedAdvanced,omitempty"`
	// The number of value-added package calls used for the summarization feature.
	//
	// example:
	//
	// 5
	ValueAddedSummary *int32 `json:"valueAddedSummary,omitempty" xml:"valueAddedSummary,omitempty"`
}

func (s GetIqsUsageResultRecords) String() string {
	return dara.Prettify(s)
}

func (s GetIqsUsageResultRecords) GoString() string {
	return s.String()
}

func (s *GetIqsUsageResultRecords) GetApi() *string {
	return s.Api
}

func (s *GetIqsUsageResultRecords) GetBillingQps() *int32 {
	return s.BillingQps
}

func (s *GetIqsUsageResultRecords) GetDate() *string {
	return s.Date
}

func (s *GetIqsUsageResultRecords) GetEngineType() *string {
	return s.EngineType
}

func (s *GetIqsUsageResultRecords) GetFailedCalls() *int32 {
	return s.FailedCalls
}

func (s *GetIqsUsageResultRecords) GetLadderType() *string {
	return s.LadderType
}

func (s *GetIqsUsageResultRecords) GetMainAccountId() *string {
	return s.MainAccountId
}

func (s *GetIqsUsageResultRecords) GetSubAccountId() *string {
	return s.SubAccountId
}

func (s *GetIqsUsageResultRecords) GetSuccessCalls() *int32 {
	return s.SuccessCalls
}

func (s *GetIqsUsageResultRecords) GetTotalBillStatCalls() *int32 {
	return s.TotalBillStatCalls
}

func (s *GetIqsUsageResultRecords) GetTotalCalls() *int32 {
	return s.TotalCalls
}

func (s *GetIqsUsageResultRecords) GetValueAddedAdvanced() *int32 {
	return s.ValueAddedAdvanced
}

func (s *GetIqsUsageResultRecords) GetValueAddedSummary() *int32 {
	return s.ValueAddedSummary
}

func (s *GetIqsUsageResultRecords) SetApi(v string) *GetIqsUsageResultRecords {
	s.Api = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetBillingQps(v int32) *GetIqsUsageResultRecords {
	s.BillingQps = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetDate(v string) *GetIqsUsageResultRecords {
	s.Date = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetEngineType(v string) *GetIqsUsageResultRecords {
	s.EngineType = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetFailedCalls(v int32) *GetIqsUsageResultRecords {
	s.FailedCalls = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetLadderType(v string) *GetIqsUsageResultRecords {
	s.LadderType = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetMainAccountId(v string) *GetIqsUsageResultRecords {
	s.MainAccountId = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetSubAccountId(v string) *GetIqsUsageResultRecords {
	s.SubAccountId = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetSuccessCalls(v int32) *GetIqsUsageResultRecords {
	s.SuccessCalls = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetTotalBillStatCalls(v int32) *GetIqsUsageResultRecords {
	s.TotalBillStatCalls = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetTotalCalls(v int32) *GetIqsUsageResultRecords {
	s.TotalCalls = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetValueAddedAdvanced(v int32) *GetIqsUsageResultRecords {
	s.ValueAddedAdvanced = &v
	return s
}

func (s *GetIqsUsageResultRecords) SetValueAddedSummary(v int32) *GetIqsUsageResultRecords {
	s.ValueAddedSummary = &v
	return s
}

func (s *GetIqsUsageResultRecords) Validate() error {
	return dara.Validate(s)
}
