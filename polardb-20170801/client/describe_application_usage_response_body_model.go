// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationUsageResponseBody
	GetApplicationId() *string
	SetCode(v int32) *DescribeApplicationUsageResponseBody
	GetCode() *int32
	SetDailyUsage(v []*DescribeApplicationUsageResponseBodyDailyUsage) *DescribeApplicationUsageResponseBody
	GetDailyUsage() []*DescribeApplicationUsageResponseBodyDailyUsage
	SetDays(v int32) *DescribeApplicationUsageResponseBody
	GetDays() *int32
	SetMessage(v string) *DescribeApplicationUsageResponseBody
	GetMessage() *string
	SetModelUsage(v []*DescribeApplicationUsageResponseBodyModelUsage) *DescribeApplicationUsageResponseBody
	GetModelUsage() []*DescribeApplicationUsageResponseBodyModelUsage
	SetRequestId(v string) *DescribeApplicationUsageResponseBody
	GetRequestId() *string
	SetSessionSummary(v *DescribeApplicationUsageResponseBodySessionSummary) *DescribeApplicationUsageResponseBody
	GetSessionSummary() *DescribeApplicationUsageResponseBodySessionSummary
	SetSkillUsage(v *DescribeApplicationUsageResponseBodySkillUsage) *DescribeApplicationUsageResponseBody
	GetSkillUsage() *DescribeApplicationUsageResponseBodySkillUsage
	SetSummary(v *DescribeApplicationUsageResponseBodySummary) *DescribeApplicationUsageResponseBody
	GetSummary() *DescribeApplicationUsageResponseBodySummary
}

type DescribeApplicationUsageResponseBody struct {
	// The Hermes application ID.
	//
	// example:
	//
	// pa-123456
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The usage statistics grouped by UTC date.
	DailyUsage []*DescribeApplicationUsageResponseBodyDailyUsage `json:"DailyUsage,omitempty" xml:"DailyUsage,omitempty" type:"Repeated"`
	// The number of days covered by this statistical period.
	//
	// example:
	//
	// 30
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The response message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The usage statistics grouped by model.
	ModelUsage []*DescribeApplicationUsageResponseBodyModelUsage `json:"ModelUsage,omitempty" xml:"ModelUsage,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F45FFACC-1234-5678-90AB-1234567890AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The current session runtime and storage statistics.
	SessionSummary *DescribeApplicationUsageResponseBodySessionSummary `json:"SessionSummary,omitempty" xml:"SessionSummary,omitempty" type:"Struct"`
	// The aggregated statistics of skill activities.
	SkillUsage *DescribeApplicationUsageResponseBodySkillUsage `json:"SkillUsage,omitempty" xml:"SkillUsage,omitempty" type:"Struct"`
	// The aggregated usage within the query period.
	Summary *DescribeApplicationUsageResponseBodySummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
}

func (s DescribeApplicationUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationUsageResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationUsageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeApplicationUsageResponseBody) GetDailyUsage() []*DescribeApplicationUsageResponseBodyDailyUsage {
	return s.DailyUsage
}

func (s *DescribeApplicationUsageResponseBody) GetDays() *int32 {
	return s.Days
}

func (s *DescribeApplicationUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationUsageResponseBody) GetModelUsage() []*DescribeApplicationUsageResponseBodyModelUsage {
	return s.ModelUsage
}

func (s *DescribeApplicationUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationUsageResponseBody) GetSessionSummary() *DescribeApplicationUsageResponseBodySessionSummary {
	return s.SessionSummary
}

func (s *DescribeApplicationUsageResponseBody) GetSkillUsage() *DescribeApplicationUsageResponseBodySkillUsage {
	return s.SkillUsage
}

func (s *DescribeApplicationUsageResponseBody) GetSummary() *DescribeApplicationUsageResponseBodySummary {
	return s.Summary
}

func (s *DescribeApplicationUsageResponseBody) SetApplicationId(v string) *DescribeApplicationUsageResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationUsageResponseBody) SetCode(v int32) *DescribeApplicationUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationUsageResponseBody) SetDailyUsage(v []*DescribeApplicationUsageResponseBodyDailyUsage) *DescribeApplicationUsageResponseBody {
	s.DailyUsage = v
	return s
}

func (s *DescribeApplicationUsageResponseBody) SetDays(v int32) *DescribeApplicationUsageResponseBody {
	s.Days = &v
	return s
}

func (s *DescribeApplicationUsageResponseBody) SetMessage(v string) *DescribeApplicationUsageResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationUsageResponseBody) SetModelUsage(v []*DescribeApplicationUsageResponseBodyModelUsage) *DescribeApplicationUsageResponseBody {
	s.ModelUsage = v
	return s
}

func (s *DescribeApplicationUsageResponseBody) SetRequestId(v string) *DescribeApplicationUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationUsageResponseBody) SetSessionSummary(v *DescribeApplicationUsageResponseBodySessionSummary) *DescribeApplicationUsageResponseBody {
	s.SessionSummary = v
	return s
}

func (s *DescribeApplicationUsageResponseBody) SetSkillUsage(v *DescribeApplicationUsageResponseBodySkillUsage) *DescribeApplicationUsageResponseBody {
	s.SkillUsage = v
	return s
}

func (s *DescribeApplicationUsageResponseBody) SetSummary(v *DescribeApplicationUsageResponseBodySummary) *DescribeApplicationUsageResponseBody {
	s.Summary = v
	return s
}

func (s *DescribeApplicationUsageResponseBody) Validate() error {
	if s.DailyUsage != nil {
		for _, item := range s.DailyUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ModelUsage != nil {
		for _, item := range s.ModelUsage {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SessionSummary != nil {
		if err := s.SessionSummary.Validate(); err != nil {
			return err
		}
	}
	if s.SkillUsage != nil {
		if err := s.SkillUsage.Validate(); err != nil {
			return err
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationUsageResponseBodyDailyUsage struct {
	// The number of model API calls for the day.
	//
	// example:
	//
	// 48
	APICalls *int64 `json:"APICalls,omitempty" xml:"APICalls,omitempty"`
	// The number of tokens served from cache hits for the day.
	//
	// example:
	//
	// 1800
	CacheReadTokens *int64 `json:"CacheReadTokens,omitempty" xml:"CacheReadTokens,omitempty"`
	// The UTC date.
	//
	// example:
	//
	// 2026-07-24
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The number of input tokens for the day.
	//
	// example:
	//
	// 12000
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// The number of output tokens for the day.
	//
	// example:
	//
	// 3600
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The number of reasoning tokens for the day.
	//
	// example:
	//
	// 600
	ReasoningTokens *int64 `json:"ReasoningTokens,omitempty" xml:"ReasoningTokens,omitempty"`
	// The number of sessions for the day.
	//
	// example:
	//
	// 12
	Sessions *int64 `json:"Sessions,omitempty" xml:"Sessions,omitempty"`
}

func (s DescribeApplicationUsageResponseBodyDailyUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationUsageResponseBodyDailyUsage) GoString() string {
	return s.String()
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) GetAPICalls() *int64 {
	return s.APICalls
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) GetCacheReadTokens() *int64 {
	return s.CacheReadTokens
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) GetDate() *string {
	return s.Date
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) GetReasoningTokens() *int64 {
	return s.ReasoningTokens
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) GetSessions() *int64 {
	return s.Sessions
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) SetAPICalls(v int64) *DescribeApplicationUsageResponseBodyDailyUsage {
	s.APICalls = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) SetCacheReadTokens(v int64) *DescribeApplicationUsageResponseBodyDailyUsage {
	s.CacheReadTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) SetDate(v string) *DescribeApplicationUsageResponseBodyDailyUsage {
	s.Date = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) SetInputTokens(v int64) *DescribeApplicationUsageResponseBodyDailyUsage {
	s.InputTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) SetOutputTokens(v int64) *DescribeApplicationUsageResponseBodyDailyUsage {
	s.OutputTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) SetReasoningTokens(v int64) *DescribeApplicationUsageResponseBodyDailyUsage {
	s.ReasoningTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) SetSessions(v int64) *DescribeApplicationUsageResponseBodyDailyUsage {
	s.Sessions = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyDailyUsage) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationUsageResponseBodyModelUsage struct {
	// The number of API calls for this model.
	//
	// example:
	//
	// 320
	APICalls *int64 `json:"APICalls,omitempty" xml:"APICalls,omitempty"`
	// The number of tokens served from cache hits for this model.
	//
	// example:
	//
	// 12000
	CacheReadTokens *int64 `json:"CacheReadTokens,omitempty" xml:"CacheReadTokens,omitempty"`
	// The number of input tokens consumed by this model.
	//
	// example:
	//
	// 80000
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// The model identifier.
	//
	// example:
	//
	// qwen3-max
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The number of output tokens generated by this model.
	//
	// example:
	//
	// 24000
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The model provider.
	//
	// example:
	//
	// bailian
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// The number of reasoning tokens generated by this model.
	//
	// example:
	//
	// 4000
	ReasoningTokens *int64 `json:"ReasoningTokens,omitempty" xml:"ReasoningTokens,omitempty"`
	// The number of sessions that used this model.
	//
	// example:
	//
	// 80
	Sessions *int64 `json:"Sessions,omitempty" xml:"Sessions,omitempty"`
}

func (s DescribeApplicationUsageResponseBodyModelUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationUsageResponseBodyModelUsage) GoString() string {
	return s.String()
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) GetAPICalls() *int64 {
	return s.APICalls
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) GetCacheReadTokens() *int64 {
	return s.CacheReadTokens
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) GetModel() *string {
	return s.Model
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) GetProvider() *string {
	return s.Provider
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) GetReasoningTokens() *int64 {
	return s.ReasoningTokens
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) GetSessions() *int64 {
	return s.Sessions
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) SetAPICalls(v int64) *DescribeApplicationUsageResponseBodyModelUsage {
	s.APICalls = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) SetCacheReadTokens(v int64) *DescribeApplicationUsageResponseBodyModelUsage {
	s.CacheReadTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) SetInputTokens(v int64) *DescribeApplicationUsageResponseBodyModelUsage {
	s.InputTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) SetModel(v string) *DescribeApplicationUsageResponseBodyModelUsage {
	s.Model = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) SetOutputTokens(v int64) *DescribeApplicationUsageResponseBodyModelUsage {
	s.OutputTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) SetProvider(v string) *DescribeApplicationUsageResponseBodyModelUsage {
	s.Provider = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) SetReasoningTokens(v int64) *DescribeApplicationUsageResponseBodyModelUsage {
	s.ReasoningTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) SetSessions(v int64) *DescribeApplicationUsageResponseBodyModelUsage {
	s.Sessions = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodyModelUsage) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationUsageResponseBodySessionSummary struct {
	// The number of currently active sessions.
	//
	// example:
	//
	// 5
	ActiveSessions *int64 `json:"ActiveSessions,omitempty" xml:"ActiveSessions,omitempty"`
	// The total number of sessions in session storage.
	//
	// example:
	//
	// 120
	StoredSessions *int64 `json:"StoredSessions,omitempty" xml:"StoredSessions,omitempty"`
}

func (s DescribeApplicationUsageResponseBodySessionSummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationUsageResponseBodySessionSummary) GoString() string {
	return s.String()
}

func (s *DescribeApplicationUsageResponseBodySessionSummary) GetActiveSessions() *int64 {
	return s.ActiveSessions
}

func (s *DescribeApplicationUsageResponseBodySessionSummary) GetStoredSessions() *int64 {
	return s.StoredSessions
}

func (s *DescribeApplicationUsageResponseBodySessionSummary) SetActiveSessions(v int64) *DescribeApplicationUsageResponseBodySessionSummary {
	s.ActiveSessions = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySessionSummary) SetStoredSessions(v int64) *DescribeApplicationUsageResponseBodySessionSummary {
	s.StoredSessions = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySessionSummary) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationUsageResponseBodySkillUsage struct {
	// The number of distinct skills that have activity records.
	//
	// example:
	//
	// 8
	DistinctSkills *int64 `json:"DistinctSkills,omitempty" xml:"DistinctSkills,omitempty"`
	// The total number of skill-related operations.
	//
	// example:
	//
	// 60
	TotalActions *int64 `json:"TotalActions,omitempty" xml:"TotalActions,omitempty"`
	// The number of times skills were edited or managed.
	//
	// example:
	//
	// 10
	TotalEdits *int64 `json:"TotalEdits,omitempty" xml:"TotalEdits,omitempty"`
	// The number of times skills were loaded or viewed.
	//
	// example:
	//
	// 50
	TotalLoads *int64 `json:"TotalLoads,omitempty" xml:"TotalLoads,omitempty"`
}

func (s DescribeApplicationUsageResponseBodySkillUsage) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationUsageResponseBodySkillUsage) GoString() string {
	return s.String()
}

func (s *DescribeApplicationUsageResponseBodySkillUsage) GetDistinctSkills() *int64 {
	return s.DistinctSkills
}

func (s *DescribeApplicationUsageResponseBodySkillUsage) GetTotalActions() *int64 {
	return s.TotalActions
}

func (s *DescribeApplicationUsageResponseBodySkillUsage) GetTotalEdits() *int64 {
	return s.TotalEdits
}

func (s *DescribeApplicationUsageResponseBodySkillUsage) GetTotalLoads() *int64 {
	return s.TotalLoads
}

func (s *DescribeApplicationUsageResponseBodySkillUsage) SetDistinctSkills(v int64) *DescribeApplicationUsageResponseBodySkillUsage {
	s.DistinctSkills = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySkillUsage) SetTotalActions(v int64) *DescribeApplicationUsageResponseBodySkillUsage {
	s.TotalActions = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySkillUsage) SetTotalEdits(v int64) *DescribeApplicationUsageResponseBodySkillUsage {
	s.TotalEdits = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySkillUsage) SetTotalLoads(v int64) *DescribeApplicationUsageResponseBodySkillUsage {
	s.TotalLoads = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySkillUsage) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationUsageResponseBodySummary struct {
	// The number of model API calls.
	//
	// example:
	//
	// 480
	APICalls *int64 `json:"APICalls,omitempty" xml:"APICalls,omitempty"`
	// The number of tokens served from cache hits.
	//
	// example:
	//
	// 18000
	CacheReadTokens *int64 `json:"CacheReadTokens,omitempty" xml:"CacheReadTokens,omitempty"`
	// The number of input tokens.
	//
	// example:
	//
	// 120000
	InputTokens *int64 `json:"InputTokens,omitempty" xml:"InputTokens,omitempty"`
	// The number of output tokens.
	//
	// example:
	//
	// 36000
	OutputTokens *int64 `json:"OutputTokens,omitempty" xml:"OutputTokens,omitempty"`
	// The number of reasoning tokens.
	//
	// example:
	//
	// 6000
	ReasoningTokens *int64 `json:"ReasoningTokens,omitempty" xml:"ReasoningTokens,omitempty"`
	// The number of sessions.
	//
	// example:
	//
	// 120
	Sessions *int64 `json:"Sessions,omitempty" xml:"Sessions,omitempty"`
}

func (s DescribeApplicationUsageResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationUsageResponseBodySummary) GoString() string {
	return s.String()
}

func (s *DescribeApplicationUsageResponseBodySummary) GetAPICalls() *int64 {
	return s.APICalls
}

func (s *DescribeApplicationUsageResponseBodySummary) GetCacheReadTokens() *int64 {
	return s.CacheReadTokens
}

func (s *DescribeApplicationUsageResponseBodySummary) GetInputTokens() *int64 {
	return s.InputTokens
}

func (s *DescribeApplicationUsageResponseBodySummary) GetOutputTokens() *int64 {
	return s.OutputTokens
}

func (s *DescribeApplicationUsageResponseBodySummary) GetReasoningTokens() *int64 {
	return s.ReasoningTokens
}

func (s *DescribeApplicationUsageResponseBodySummary) GetSessions() *int64 {
	return s.Sessions
}

func (s *DescribeApplicationUsageResponseBodySummary) SetAPICalls(v int64) *DescribeApplicationUsageResponseBodySummary {
	s.APICalls = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySummary) SetCacheReadTokens(v int64) *DescribeApplicationUsageResponseBodySummary {
	s.CacheReadTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySummary) SetInputTokens(v int64) *DescribeApplicationUsageResponseBodySummary {
	s.InputTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySummary) SetOutputTokens(v int64) *DescribeApplicationUsageResponseBodySummary {
	s.OutputTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySummary) SetReasoningTokens(v int64) *DescribeApplicationUsageResponseBodySummary {
	s.ReasoningTokens = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySummary) SetSessions(v int64) *DescribeApplicationUsageResponseBodySummary {
	s.Sessions = &v
	return s
}

func (s *DescribeApplicationUsageResponseBodySummary) Validate() error {
	return dara.Validate(s)
}
