// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConversationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConversations(v []*ListConversationsResponseBodyConversations) *ListConversationsResponseBody
	GetConversations() []*ListConversationsResponseBodyConversations
	SetPageNumber(v int32) *ListConversationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConversationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListConversationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListConversationsResponseBody
	GetTotalCount() *int64
}

type ListConversationsResponseBody struct {
	// The list of sessions.
	Conversations []*ListConversationsResponseBodyConversations `json:"Conversations,omitempty" xml:"Conversations,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// a2c26e67-5984-4935-984e-bcee52971993
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListConversationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListConversationsResponseBody) GetConversations() []*ListConversationsResponseBodyConversations {
	return s.Conversations
}

func (s *ListConversationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConversationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConversationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConversationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListConversationsResponseBody) SetConversations(v []*ListConversationsResponseBodyConversations) *ListConversationsResponseBody {
	s.Conversations = v
	return s
}

func (s *ListConversationsResponseBody) SetPageNumber(v int32) *ListConversationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListConversationsResponseBody) SetPageSize(v int32) *ListConversationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListConversationsResponseBody) SetRequestId(v string) *ListConversationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConversationsResponseBody) SetTotalCount(v int64) *ListConversationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListConversationsResponseBody) Validate() error {
	if s.Conversations != nil {
		for _, item := range s.Conversations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConversationsResponseBodyConversations struct {
	AbTestName *string `json:"AbTestName,omitempty" xml:"AbTestName,omitempty"`
	// The called number.
	//
	// example:
	//
	// 12356
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 135815884***
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// The session ID.
	//
	// example:
	//
	// 82b2eaae-ce5c-45f8-8231-f15b6b27e55c
	ConversationId *string   `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	DsReport       *string   `json:"DsReport,omitempty" xml:"DsReport,omitempty"`
	DsReportTitles []*string `json:"DsReportTitles,omitempty" xml:"DsReportTitles,omitempty" type:"Repeated"`
	// The reason for hanging up. Valid values:
	//
	//      1: Normal completion.
	//
	//      2: Bot hung up after unrecognized input.
	//
	//      3: Hung up due to silence timeout.
	//
	//      4: User hung up after unrecognized input.
	//
	//      5: User hung up without reason.
	//
	//      6: Transferred to human agent due to intent match.
	//
	//      7: Transferred to human agent due to unrecognized input.
	//
	//      8: No interaction from the user side.
	//
	//      9: System exception interruption.
	//
	//      10: Transferred to IVR due to intent match.
	//
	//      11: Transferred to IVR due to unrecognized input.
	//
	// example:
	//
	// 1
	EndReason *int32 `json:"EndReason,omitempty" xml:"EndReason,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1582266750353
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Indicates whether the last playback was completed when the session ended.
	//
	// example:
	//
	// true
	HasLastPlaybackCompleted *bool `json:"HasLastPlaybackCompleted,omitempty" xml:"HasLastPlaybackCompleted,omitempty"`
	// Indicates whether the session was transferred to a human agent.
	//
	// example:
	//
	// false
	HasToAgent *bool `json:"HasToAgent,omitempty" xml:"HasToAgent,omitempty"`
	// The number of conversation rounds.
	//
	// example:
	//
	// 2
	Rounds *int32 `json:"Rounds,omitempty" xml:"Rounds,omitempty"`
	// Indicates whether the session is in a sandbox environment.
	//
	// example:
	//
	// true
	SandBox *bool `json:"SandBox,omitempty" xml:"SandBox,omitempty"`
	// The skill group.
	//
	// example:
	//
	// skg-123
	SkillGroup *string `json:"SkillGroup,omitempty" xml:"SkillGroup,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1641625694286
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListConversationsResponseBodyConversations) String() string {
	return dara.Prettify(s)
}

func (s ListConversationsResponseBodyConversations) GoString() string {
	return s.String()
}

func (s *ListConversationsResponseBodyConversations) GetAbTestName() *string {
	return s.AbTestName
}

func (s *ListConversationsResponseBodyConversations) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *ListConversationsResponseBodyConversations) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *ListConversationsResponseBodyConversations) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListConversationsResponseBodyConversations) GetDsReport() *string {
	return s.DsReport
}

func (s *ListConversationsResponseBodyConversations) GetDsReportTitles() []*string {
	return s.DsReportTitles
}

func (s *ListConversationsResponseBodyConversations) GetEndReason() *int32 {
	return s.EndReason
}

func (s *ListConversationsResponseBodyConversations) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListConversationsResponseBodyConversations) GetHasLastPlaybackCompleted() *bool {
	return s.HasLastPlaybackCompleted
}

func (s *ListConversationsResponseBodyConversations) GetHasToAgent() *bool {
	return s.HasToAgent
}

func (s *ListConversationsResponseBodyConversations) GetRounds() *int32 {
	return s.Rounds
}

func (s *ListConversationsResponseBodyConversations) GetSandBox() *bool {
	return s.SandBox
}

func (s *ListConversationsResponseBodyConversations) GetSkillGroup() *string {
	return s.SkillGroup
}

func (s *ListConversationsResponseBodyConversations) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListConversationsResponseBodyConversations) SetAbTestName(v string) *ListConversationsResponseBodyConversations {
	s.AbTestName = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetCalledNumber(v string) *ListConversationsResponseBodyConversations {
	s.CalledNumber = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetCallingNumber(v string) *ListConversationsResponseBodyConversations {
	s.CallingNumber = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetConversationId(v string) *ListConversationsResponseBodyConversations {
	s.ConversationId = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetDsReport(v string) *ListConversationsResponseBodyConversations {
	s.DsReport = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetDsReportTitles(v []*string) *ListConversationsResponseBodyConversations {
	s.DsReportTitles = v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetEndReason(v int32) *ListConversationsResponseBodyConversations {
	s.EndReason = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetEndTime(v int64) *ListConversationsResponseBodyConversations {
	s.EndTime = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetHasLastPlaybackCompleted(v bool) *ListConversationsResponseBodyConversations {
	s.HasLastPlaybackCompleted = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetHasToAgent(v bool) *ListConversationsResponseBodyConversations {
	s.HasToAgent = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetRounds(v int32) *ListConversationsResponseBodyConversations {
	s.Rounds = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetSandBox(v bool) *ListConversationsResponseBodyConversations {
	s.SandBox = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetSkillGroup(v string) *ListConversationsResponseBodyConversations {
	s.SkillGroup = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) SetStartTime(v int64) *ListConversationsResponseBodyConversations {
	s.StartTime = &v
	return s
}

func (s *ListConversationsResponseBodyConversations) Validate() error {
	return dara.Validate(s)
}
