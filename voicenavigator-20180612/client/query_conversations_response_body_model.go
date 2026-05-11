// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConversationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConversations(v []*QueryConversationsResponseBodyConversations) *QueryConversationsResponseBody
	GetConversations() []*QueryConversationsResponseBodyConversations
	SetPageNumber(v int32) *QueryConversationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryConversationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryConversationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *QueryConversationsResponseBody
	GetTotalCount() *int64
}

type QueryConversationsResponseBody struct {
	Conversations []*QueryConversationsResponseBodyConversations `json:"Conversations,omitempty" xml:"Conversations,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryConversationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryConversationsResponseBody) GetConversations() []*QueryConversationsResponseBodyConversations {
	return s.Conversations
}

func (s *QueryConversationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryConversationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryConversationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryConversationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryConversationsResponseBody) SetConversations(v []*QueryConversationsResponseBodyConversations) *QueryConversationsResponseBody {
	s.Conversations = v
	return s
}

func (s *QueryConversationsResponseBody) SetPageNumber(v int32) *QueryConversationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryConversationsResponseBody) SetPageSize(v int32) *QueryConversationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryConversationsResponseBody) SetRequestId(v string) *QueryConversationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryConversationsResponseBody) SetTotalCount(v int64) *QueryConversationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryConversationsResponseBody) Validate() error {
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

type QueryConversationsResponseBodyConversations struct {
	// example:
	//
	// 1582183381000
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 02811111111
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 8
	EffectiveAnswerCount *int32 `json:"EffectiveAnswerCount,omitempty" xml:"EffectiveAnswerCount,omitempty"`
	// example:
	//
	// 1582183481000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// AAA
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
	// example:
	//
	// true
	TransferredToAgent *bool `json:"TransferredToAgent,omitempty" xml:"TransferredToAgent,omitempty"`
	// example:
	//
	// 10
	UserUtteranceCount *int32 `json:"UserUtteranceCount,omitempty" xml:"UserUtteranceCount,omitempty"`
}

func (s QueryConversationsResponseBodyConversations) String() string {
	return dara.Prettify(s)
}

func (s QueryConversationsResponseBodyConversations) GoString() string {
	return s.String()
}

func (s *QueryConversationsResponseBodyConversations) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *QueryConversationsResponseBodyConversations) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *QueryConversationsResponseBodyConversations) GetConversationId() *string {
	return s.ConversationId
}

func (s *QueryConversationsResponseBodyConversations) GetEffectiveAnswerCount() *int32 {
	return s.EffectiveAnswerCount
}

func (s *QueryConversationsResponseBodyConversations) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryConversationsResponseBodyConversations) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *QueryConversationsResponseBodyConversations) GetTransferredToAgent() *bool {
	return s.TransferredToAgent
}

func (s *QueryConversationsResponseBodyConversations) GetUserUtteranceCount() *int32 {
	return s.UserUtteranceCount
}

func (s *QueryConversationsResponseBodyConversations) SetBeginTime(v int64) *QueryConversationsResponseBodyConversations {
	s.BeginTime = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetCallingNumber(v string) *QueryConversationsResponseBodyConversations {
	s.CallingNumber = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetConversationId(v string) *QueryConversationsResponseBodyConversations {
	s.ConversationId = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetEffectiveAnswerCount(v int32) *QueryConversationsResponseBodyConversations {
	s.EffectiveAnswerCount = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetEndTime(v int64) *QueryConversationsResponseBodyConversations {
	s.EndTime = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetSkillGroupId(v string) *QueryConversationsResponseBodyConversations {
	s.SkillGroupId = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetTransferredToAgent(v bool) *QueryConversationsResponseBodyConversations {
	s.TransferredToAgent = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) SetUserUtteranceCount(v int32) *QueryConversationsResponseBodyConversations {
	s.UserUtteranceCount = &v
	return s
}

func (s *QueryConversationsResponseBodyConversations) Validate() error {
	return dara.Validate(s)
}
