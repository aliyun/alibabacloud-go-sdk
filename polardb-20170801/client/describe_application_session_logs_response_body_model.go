// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSessionLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationSessionLogsResponseBody
	GetApplicationId() *string
	SetItems(v *DescribeApplicationSessionLogsResponseBodyItems) *DescribeApplicationSessionLogsResponseBody
	GetItems() *DescribeApplicationSessionLogsResponseBodyItems
	SetPageNumber(v int64) *DescribeApplicationSessionLogsResponseBody
	GetPageNumber() *int64
	SetPageRecordCount(v int64) *DescribeApplicationSessionLogsResponseBody
	GetPageRecordCount() *int64
	SetRequestId(v string) *DescribeApplicationSessionLogsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int64) *DescribeApplicationSessionLogsResponseBody
	GetTotalRecordCount() *int64
}

type DescribeApplicationSessionLogsResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string                                          `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Items         *DescribeApplicationSessionLogsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of log records returned on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int64 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of log records that meet the query conditions.
	//
	// example:
	//
	// 1
	TotalRecordCount *int64 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeApplicationSessionLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionLogsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionLogsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationSessionLogsResponseBody) GetItems() *DescribeApplicationSessionLogsResponseBodyItems {
	return s.Items
}

func (s *DescribeApplicationSessionLogsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApplicationSessionLogsResponseBody) GetPageRecordCount() *int64 {
	return s.PageRecordCount
}

func (s *DescribeApplicationSessionLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationSessionLogsResponseBody) GetTotalRecordCount() *int64 {
	return s.TotalRecordCount
}

func (s *DescribeApplicationSessionLogsResponseBody) SetApplicationId(v string) *DescribeApplicationSessionLogsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBody) SetItems(v *DescribeApplicationSessionLogsResponseBodyItems) *DescribeApplicationSessionLogsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBody) SetPageNumber(v int64) *DescribeApplicationSessionLogsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBody) SetPageRecordCount(v int64) *DescribeApplicationSessionLogsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBody) SetRequestId(v string) *DescribeApplicationSessionLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBody) SetTotalRecordCount(v int64) *DescribeApplicationSessionLogsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationSessionLogsResponseBodyItems struct {
	Items []*DescribeApplicationSessionLogsResponseBodyItemsItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s DescribeApplicationSessionLogsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionLogsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionLogsResponseBodyItems) GetItems() []*DescribeApplicationSessionLogsResponseBodyItemsItems {
	return s.Items
}

func (s *DescribeApplicationSessionLogsResponseBodyItems) SetItems(v []*DescribeApplicationSessionLogsResponseBodyItemsItems) *DescribeApplicationSessionLogsResponseBodyItems {
	s.Items = v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItems) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationSessionLogsResponseBodyItemsItems struct {
	Agent               *string `json:"Agent,omitempty" xml:"Agent,omitempty"`
	Api                 *string `json:"Api,omitempty" xml:"Api,omitempty"`
	Content             *string `json:"Content,omitempty" xml:"Content,omitempty"`
	DetailDurationMs    *int64  `json:"DetailDurationMs,omitempty" xml:"DetailDurationMs,omitempty"`
	DetailIsError       *bool   `json:"DetailIsError,omitempty" xml:"DetailIsError,omitempty"`
	DetailStatus        *string `json:"DetailStatus,omitempty" xml:"DetailStatus,omitempty"`
	Id                  *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Model               *string `json:"Model,omitempty" xml:"Model,omitempty"`
	ParentId            *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	Provider            *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	Role                *string `json:"Role,omitempty" xml:"Role,omitempty"`
	SessionId           *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StopReason          *string `json:"StopReason,omitempty" xml:"StopReason,omitempty"`
	Text                *string `json:"Text,omitempty" xml:"Text,omitempty"`
	Thinking            *string `json:"Thinking,omitempty" xml:"Thinking,omitempty"`
	ThinkingSignature   *string `json:"ThinkingSignature,omitempty" xml:"ThinkingSignature,omitempty"`
	Time                *string `json:"Time,omitempty" xml:"Time,omitempty"`
	Timestamp           *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	ToolName            *string `json:"ToolName,omitempty" xml:"ToolName,omitempty"`
	TotalTokens         *int64  `json:"TotalTokens,omitempty" xml:"TotalTokens,omitempty"`
	Type                *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UsageCacheRead      *int64  `json:"UsageCacheRead,omitempty" xml:"UsageCacheRead,omitempty"`
	UsageCacheWrite     *int64  `json:"UsageCacheWrite,omitempty" xml:"UsageCacheWrite,omitempty"`
	UsageCostCacheRead  *int64  `json:"UsageCostCacheRead,omitempty" xml:"UsageCostCacheRead,omitempty"`
	UsageCostCacheWrite *int64  `json:"UsageCostCacheWrite,omitempty" xml:"UsageCostCacheWrite,omitempty"`
	UsageCostInput      *int64  `json:"UsageCostInput,omitempty" xml:"UsageCostInput,omitempty"`
	UsageCostOutput     *int64  `json:"UsageCostOutput,omitempty" xml:"UsageCostOutput,omitempty"`
	UsageCostTotal      *int64  `json:"UsageCostTotal,omitempty" xml:"UsageCostTotal,omitempty"`
	UsageInput          *int64  `json:"UsageInput,omitempty" xml:"UsageInput,omitempty"`
	UsageOutput         *int64  `json:"UsageOutput,omitempty" xml:"UsageOutput,omitempty"`
	UsageTotalTokens    *int64  `json:"UsageTotalTokens,omitempty" xml:"UsageTotalTokens,omitempty"`
}

func (s DescribeApplicationSessionLogsResponseBodyItemsItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSessionLogsResponseBodyItemsItems) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetAgent() *string {
	return s.Agent
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetApi() *string {
	return s.Api
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetContent() *string {
	return s.Content
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetDetailDurationMs() *int64 {
	return s.DetailDurationMs
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetDetailIsError() *bool {
	return s.DetailIsError
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetDetailStatus() *string {
	return s.DetailStatus
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetId() *string {
	return s.Id
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetModel() *string {
	return s.Model
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetProvider() *string {
	return s.Provider
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetRole() *string {
	return s.Role
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetStopReason() *string {
	return s.StopReason
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetText() *string {
	return s.Text
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetThinking() *string {
	return s.Thinking
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetThinkingSignature() *string {
	return s.ThinkingSignature
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetTime() *string {
	return s.Time
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetToolName() *string {
	return s.ToolName
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetTotalTokens() *int64 {
	return s.TotalTokens
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetType() *string {
	return s.Type
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageCacheRead() *int64 {
	return s.UsageCacheRead
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageCacheWrite() *int64 {
	return s.UsageCacheWrite
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageCostCacheRead() *int64 {
	return s.UsageCostCacheRead
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageCostCacheWrite() *int64 {
	return s.UsageCostCacheWrite
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageCostInput() *int64 {
	return s.UsageCostInput
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageCostOutput() *int64 {
	return s.UsageCostOutput
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageCostTotal() *int64 {
	return s.UsageCostTotal
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageInput() *int64 {
	return s.UsageInput
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageOutput() *int64 {
	return s.UsageOutput
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) GetUsageTotalTokens() *int64 {
	return s.UsageTotalTokens
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetAgent(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Agent = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetApi(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Api = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetContent(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Content = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetDetailDurationMs(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.DetailDurationMs = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetDetailIsError(v bool) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.DetailIsError = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetDetailStatus(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.DetailStatus = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetId(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Id = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetModel(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Model = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetParentId(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.ParentId = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetProvider(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Provider = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetRole(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Role = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetSessionId(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.SessionId = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetStatus(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Status = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetStopReason(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.StopReason = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetText(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Text = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetThinking(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Thinking = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetThinkingSignature(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.ThinkingSignature = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetTime(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Time = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetTimestamp(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Timestamp = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetToolName(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.ToolName = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetTotalTokens(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.TotalTokens = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetType(v string) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.Type = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageCacheRead(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageCacheRead = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageCacheWrite(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageCacheWrite = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageCostCacheRead(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageCostCacheRead = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageCostCacheWrite(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageCostCacheWrite = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageCostInput(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageCostInput = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageCostOutput(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageCostOutput = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageCostTotal(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageCostTotal = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageInput(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageInput = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageOutput(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageOutput = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) SetUsageTotalTokens(v int64) *DescribeApplicationSessionLogsResponseBodyItemsItems {
	s.UsageTotalTokens = &v
	return s
}

func (s *DescribeApplicationSessionLogsResponseBodyItemsItems) Validate() error {
	return dara.Validate(s)
}
