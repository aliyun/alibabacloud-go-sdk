// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMOTokenUsageSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DescribeMOTokenUsageSummaryResponseBody
	GetMessage() *string
	SetRecords(v []*DescribeMOTokenUsageSummaryResponseBodyRecords) *DescribeMOTokenUsageSummaryResponseBody
	GetRecords() []*DescribeMOTokenUsageSummaryResponseBodyRecords
	SetRequestId(v string) *DescribeMOTokenUsageSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMOTokenUsageSummaryResponseBody
	GetSuccess() *bool
	SetSummary(v *DescribeMOTokenUsageSummaryResponseBodySummary) *DescribeMOTokenUsageSummaryResponseBody
	GetSummary() *DescribeMOTokenUsageSummaryResponseBodySummary
	SetUsageType(v string) *DescribeMOTokenUsageSummaryResponseBody
	GetUsageType() *string
}

type DescribeMOTokenUsageSummaryResponseBody struct {
	// example:
	//
	// success
	Message *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Records []*DescribeMOTokenUsageSummaryResponseBodyRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	Summary *DescribeMOTokenUsageSummaryResponseBodySummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
	// example:
	//
	// text
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s DescribeMOTokenUsageSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOTokenUsageSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMOTokenUsageSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMOTokenUsageSummaryResponseBody) GetRecords() []*DescribeMOTokenUsageSummaryResponseBodyRecords {
	return s.Records
}

func (s *DescribeMOTokenUsageSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMOTokenUsageSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMOTokenUsageSummaryResponseBody) GetSummary() *DescribeMOTokenUsageSummaryResponseBodySummary {
	return s.Summary
}

func (s *DescribeMOTokenUsageSummaryResponseBody) GetUsageType() *string {
	return s.UsageType
}

func (s *DescribeMOTokenUsageSummaryResponseBody) SetMessage(v string) *DescribeMOTokenUsageSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBody) SetRecords(v []*DescribeMOTokenUsageSummaryResponseBodyRecords) *DescribeMOTokenUsageSummaryResponseBody {
	s.Records = v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBody) SetRequestId(v string) *DescribeMOTokenUsageSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBody) SetSuccess(v bool) *DescribeMOTokenUsageSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBody) SetSummary(v *DescribeMOTokenUsageSummaryResponseBodySummary) *DescribeMOTokenUsageSummaryResponseBody {
	s.Summary = v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBody) SetUsageType(v string) *DescribeMOTokenUsageSummaryResponseBody {
	s.UsageType = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBody) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMOTokenUsageSummaryResponseBodyRecords struct {
	// example:
	//
	// sk-rds-ds5jjo08hyz1g9orhs3y56l5cy3l3shm
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// 8000
	CacheTokens *float64 `json:"CacheTokens,omitempty" xml:"CacheTokens,omitempty"`
	// example:
	//
	// 2026-04-21
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 10000
	InputTokens *float64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// qcy-apikey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// example:
	//
	// qwen-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// example:
	//
	// 10000
	OutputTokens *float64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 100
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// example:
	//
	// 100000
	TotalTokens *float64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
	// example:
	//
	// text
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s DescribeMOTokenUsageSummaryResponseBodyRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOTokenUsageSummaryResponseBodyRecords) GoString() string {
	return s.String()
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetCacheTokens() *float64 {
	return s.CacheTokens
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetDate() *string {
	return s.Date
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetInputTokens() *float64 {
	return s.InputTokens
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetKeyName() *string {
	return s.KeyName
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetModel() *string {
	return s.Model
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetOutputTokens() *float64 {
	return s.OutputTokens
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetRequestCount() *int64 {
	return s.RequestCount
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetTotalTokens() *float64 {
	return s.TotalTokens
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) GetUsageType() *string {
	return s.UsageType
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetApiKey(v string) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.ApiKey = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetCacheTokens(v float64) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.CacheTokens = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetDate(v string) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.Date = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetInputTokens(v float64) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.InputTokens = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetKeyName(v string) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.KeyName = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetModel(v string) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.Model = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetOutputTokens(v float64) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.OutputTokens = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetRequestCount(v int64) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.RequestCount = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetTotalTokens(v float64) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.TotalTokens = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) SetUsageType(v string) *DescribeMOTokenUsageSummaryResponseBodyRecords {
	s.UsageType = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodyRecords) Validate() error {
	return dara.Validate(s)
}

type DescribeMOTokenUsageSummaryResponseBodySummary struct {
	// example:
	//
	// 9000
	CacheTokens *float64 `json:"CacheTokens,omitempty" xml:"CacheTokens,omitempty"`
	// example:
	//
	// 10000
	InputTokens *float64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// example:
	//
	// 1000
	OutputTokens *float64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// example:
	//
	// 11000
	TotalTokens *float64 `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
}

func (s DescribeMOTokenUsageSummaryResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeMOTokenUsageSummaryResponseBodySummary) GoString() string {
	return s.String()
}

func (s *DescribeMOTokenUsageSummaryResponseBodySummary) GetCacheTokens() *float64 {
	return s.CacheTokens
}

func (s *DescribeMOTokenUsageSummaryResponseBodySummary) GetInputTokens() *float64 {
	return s.InputTokens
}

func (s *DescribeMOTokenUsageSummaryResponseBodySummary) GetOutputTokens() *float64 {
	return s.OutputTokens
}

func (s *DescribeMOTokenUsageSummaryResponseBodySummary) GetTotalTokens() *float64 {
	return s.TotalTokens
}

func (s *DescribeMOTokenUsageSummaryResponseBodySummary) SetCacheTokens(v float64) *DescribeMOTokenUsageSummaryResponseBodySummary {
	s.CacheTokens = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodySummary) SetInputTokens(v float64) *DescribeMOTokenUsageSummaryResponseBodySummary {
	s.InputTokens = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodySummary) SetOutputTokens(v float64) *DescribeMOTokenUsageSummaryResponseBodySummary {
	s.OutputTokens = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodySummary) SetTotalTokens(v float64) *DescribeMOTokenUsageSummaryResponseBodySummary {
	s.TotalTokens = &v
	return s
}

func (s *DescribeMOTokenUsageSummaryResponseBodySummary) Validate() error {
	return dara.Validate(s)
}
