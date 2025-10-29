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
	Api                *string `json:"api,omitempty" xml:"api,omitempty"`
	BillingQps         *int32  `json:"billingQps,omitempty" xml:"billingQps,omitempty"`
	Date               *string `json:"date,omitempty" xml:"date,omitempty"`
	EngineType         *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	FailedCalls        *int32  `json:"failedCalls,omitempty" xml:"failedCalls,omitempty"`
	LadderType         *string `json:"ladderType,omitempty" xml:"ladderType,omitempty"`
	MainAccountId      *string `json:"mainAccountId,omitempty" xml:"mainAccountId,omitempty"`
	SubAccountId       *string `json:"subAccountId,omitempty" xml:"subAccountId,omitempty"`
	SuccessCalls       *int32  `json:"successCalls,omitempty" xml:"successCalls,omitempty"`
	TotalCalls         *int32  `json:"totalCalls,omitempty" xml:"totalCalls,omitempty"`
	ValueAddedAdvanced *int32  `json:"valueAddedAdvanced,omitempty" xml:"valueAddedAdvanced,omitempty"`
	ValueAddedSummary  *int32  `json:"valueAddedSummary,omitempty" xml:"valueAddedSummary,omitempty"`
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
