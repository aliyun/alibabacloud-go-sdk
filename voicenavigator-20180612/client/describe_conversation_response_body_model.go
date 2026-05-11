// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeginTime(v int64) *DescribeConversationResponseBody
	GetBeginTime() *int64
	SetCallingNumber(v string) *DescribeConversationResponseBody
	GetCallingNumber() *string
	SetConversationId(v string) *DescribeConversationResponseBody
	GetConversationId() *string
	SetEffectiveAnswerCount(v int32) *DescribeConversationResponseBody
	GetEffectiveAnswerCount() *int32
	SetEndTime(v int64) *DescribeConversationResponseBody
	GetEndTime() *int64
	SetRequestId(v string) *DescribeConversationResponseBody
	GetRequestId() *string
	SetSkillGroupId(v string) *DescribeConversationResponseBody
	GetSkillGroupId() *string
	SetTransferredToAgent(v bool) *DescribeConversationResponseBody
	GetTransferredToAgent() *bool
	SetUserUtteranceCount(v int32) *DescribeConversationResponseBody
	GetUserUtteranceCount() *int32
}

type DescribeConversationResponseBody struct {
	// example:
	//
	// 1582103260434
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// example:
	//
	// 138106*****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// example:
	//
	// 2d5aa451-661f-4f08-b0c4-28eec78decc4
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 8
	EffectiveAnswerCount *int32 `json:"EffectiveAnswerCount,omitempty" xml:"EffectiveAnswerCount,omitempty"`
	// example:
	//
	// 1582103299434
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// ABABCBAC
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

func (s DescribeConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConversationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConversationResponseBody) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *DescribeConversationResponseBody) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *DescribeConversationResponseBody) GetConversationId() *string {
	return s.ConversationId
}

func (s *DescribeConversationResponseBody) GetEffectiveAnswerCount() *int32 {
	return s.EffectiveAnswerCount
}

func (s *DescribeConversationResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConversationResponseBody) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *DescribeConversationResponseBody) GetTransferredToAgent() *bool {
	return s.TransferredToAgent
}

func (s *DescribeConversationResponseBody) GetUserUtteranceCount() *int32 {
	return s.UserUtteranceCount
}

func (s *DescribeConversationResponseBody) SetBeginTime(v int64) *DescribeConversationResponseBody {
	s.BeginTime = &v
	return s
}

func (s *DescribeConversationResponseBody) SetCallingNumber(v string) *DescribeConversationResponseBody {
	s.CallingNumber = &v
	return s
}

func (s *DescribeConversationResponseBody) SetConversationId(v string) *DescribeConversationResponseBody {
	s.ConversationId = &v
	return s
}

func (s *DescribeConversationResponseBody) SetEffectiveAnswerCount(v int32) *DescribeConversationResponseBody {
	s.EffectiveAnswerCount = &v
	return s
}

func (s *DescribeConversationResponseBody) SetEndTime(v int64) *DescribeConversationResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeConversationResponseBody) SetRequestId(v string) *DescribeConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConversationResponseBody) SetSkillGroupId(v string) *DescribeConversationResponseBody {
	s.SkillGroupId = &v
	return s
}

func (s *DescribeConversationResponseBody) SetTransferredToAgent(v bool) *DescribeConversationResponseBody {
	s.TransferredToAgent = &v
	return s
}

func (s *DescribeConversationResponseBody) SetUserUtteranceCount(v int32) *DescribeConversationResponseBody {
	s.UserUtteranceCount = &v
	return s
}

func (s *DescribeConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
