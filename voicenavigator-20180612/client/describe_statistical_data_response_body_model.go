// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStatisticalDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConversationTotalNum(v int64) *DescribeStatisticalDataResponseBody
	GetConversationTotalNum() *int64
	SetRequestId(v string) *DescribeStatisticalDataResponseBody
	GetRequestId() *string
	SetResolvedQuestionTotalNum(v int64) *DescribeStatisticalDataResponseBody
	GetResolvedQuestionTotalNum() *int64
	SetStatisticalDataReports(v []*DescribeStatisticalDataResponseBodyStatisticalDataReports) *DescribeStatisticalDataResponseBody
	GetStatisticalDataReports() []*DescribeStatisticalDataResponseBodyStatisticalDataReports
	SetTotalDialoguePassRate(v string) *DescribeStatisticalDataResponseBody
	GetTotalDialoguePassRate() *string
	SetTotalKnowledgeHitRate(v string) *DescribeStatisticalDataResponseBody
	GetTotalKnowledgeHitRate() *string
	SetTotalResolutionRate(v string) *DescribeStatisticalDataResponseBody
	GetTotalResolutionRate() *string
	SetTotalValidAnswerRate(v string) *DescribeStatisticalDataResponseBody
	GetTotalValidAnswerRate() *string
}

type DescribeStatisticalDataResponseBody struct {
	// example:
	//
	// 100
	ConversationTotalNum *int64 `json:"ConversationTotalNum,omitempty" xml:"ConversationTotalNum,omitempty"`
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 80
	ResolvedQuestionTotalNum *int64                                                       `json:"ResolvedQuestionTotalNum,omitempty" xml:"ResolvedQuestionTotalNum,omitempty"`
	StatisticalDataReports   []*DescribeStatisticalDataResponseBodyStatisticalDataReports `json:"StatisticalDataReports,omitempty" xml:"StatisticalDataReports,omitempty" type:"Repeated"`
	// example:
	//
	// 80.00%
	TotalDialoguePassRate *string `json:"TotalDialoguePassRate,omitempty" xml:"TotalDialoguePassRate,omitempty"`
	// example:
	//
	// 80.00%
	TotalKnowledgeHitRate *string `json:"TotalKnowledgeHitRate,omitempty" xml:"TotalKnowledgeHitRate,omitempty"`
	// example:
	//
	// 80.00%
	TotalResolutionRate *string `json:"TotalResolutionRate,omitempty" xml:"TotalResolutionRate,omitempty"`
	// example:
	//
	// 80.00%
	TotalValidAnswerRate *string `json:"TotalValidAnswerRate,omitempty" xml:"TotalValidAnswerRate,omitempty"`
}

func (s DescribeStatisticalDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticalDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataResponseBody) GetConversationTotalNum() *int64 {
	return s.ConversationTotalNum
}

func (s *DescribeStatisticalDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStatisticalDataResponseBody) GetResolvedQuestionTotalNum() *int64 {
	return s.ResolvedQuestionTotalNum
}

func (s *DescribeStatisticalDataResponseBody) GetStatisticalDataReports() []*DescribeStatisticalDataResponseBodyStatisticalDataReports {
	return s.StatisticalDataReports
}

func (s *DescribeStatisticalDataResponseBody) GetTotalDialoguePassRate() *string {
	return s.TotalDialoguePassRate
}

func (s *DescribeStatisticalDataResponseBody) GetTotalKnowledgeHitRate() *string {
	return s.TotalKnowledgeHitRate
}

func (s *DescribeStatisticalDataResponseBody) GetTotalResolutionRate() *string {
	return s.TotalResolutionRate
}

func (s *DescribeStatisticalDataResponseBody) GetTotalValidAnswerRate() *string {
	return s.TotalValidAnswerRate
}

func (s *DescribeStatisticalDataResponseBody) SetConversationTotalNum(v int64) *DescribeStatisticalDataResponseBody {
	s.ConversationTotalNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetRequestId(v string) *DescribeStatisticalDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetResolvedQuestionTotalNum(v int64) *DescribeStatisticalDataResponseBody {
	s.ResolvedQuestionTotalNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetStatisticalDataReports(v []*DescribeStatisticalDataResponseBodyStatisticalDataReports) *DescribeStatisticalDataResponseBody {
	s.StatisticalDataReports = v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetTotalDialoguePassRate(v string) *DescribeStatisticalDataResponseBody {
	s.TotalDialoguePassRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetTotalKnowledgeHitRate(v string) *DescribeStatisticalDataResponseBody {
	s.TotalKnowledgeHitRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetTotalResolutionRate(v string) *DescribeStatisticalDataResponseBody {
	s.TotalResolutionRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) SetTotalValidAnswerRate(v string) *DescribeStatisticalDataResponseBody {
	s.TotalValidAnswerRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBody) Validate() error {
	if s.StatisticalDataReports != nil {
		for _, item := range s.StatisticalDataReports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStatisticalDataResponseBodyStatisticalDataReports struct {
	// example:
	//
	// 80.00%
	DialoguePassRate *string `json:"DialoguePassRate,omitempty" xml:"DialoguePassRate,omitempty"`
	// example:
	//
	// 80.00%
	KnowledgeHitRate *string `json:"KnowledgeHitRate,omitempty" xml:"KnowledgeHitRate,omitempty"`
	// example:
	//
	// 80.00%
	ResolutionRate *string `json:"ResolutionRate,omitempty" xml:"ResolutionRate,omitempty"`
	// example:
	//
	// 80
	ResolvedQuestionNum *int32 `json:"ResolvedQuestionNum,omitempty" xml:"ResolvedQuestionNum,omitempty"`
	// example:
	//
	// 19:00:00
	StatisticalDate *string `json:"StatisticalDate,omitempty" xml:"StatisticalDate,omitempty"`
	// example:
	//
	// 100
	TotalConversationNum *int32 `json:"TotalConversationNum,omitempty" xml:"TotalConversationNum,omitempty"`
	// example:
	//
	// 80.0
	ValidAnswerRate *string `json:"ValidAnswerRate,omitempty" xml:"ValidAnswerRate,omitempty"`
}

func (s DescribeStatisticalDataResponseBodyStatisticalDataReports) String() string {
	return dara.Prettify(s)
}

func (s DescribeStatisticalDataResponseBodyStatisticalDataReports) GoString() string {
	return s.String()
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) GetDialoguePassRate() *string {
	return s.DialoguePassRate
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) GetKnowledgeHitRate() *string {
	return s.KnowledgeHitRate
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) GetResolutionRate() *string {
	return s.ResolutionRate
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) GetResolvedQuestionNum() *int32 {
	return s.ResolvedQuestionNum
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) GetStatisticalDate() *string {
	return s.StatisticalDate
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) GetTotalConversationNum() *int32 {
	return s.TotalConversationNum
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) GetValidAnswerRate() *string {
	return s.ValidAnswerRate
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetDialoguePassRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.DialoguePassRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetKnowledgeHitRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.KnowledgeHitRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetResolutionRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.ResolutionRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetResolvedQuestionNum(v int32) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.ResolvedQuestionNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetStatisticalDate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.StatisticalDate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetTotalConversationNum(v int32) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.TotalConversationNum = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) SetValidAnswerRate(v string) *DescribeStatisticalDataResponseBodyStatisticalDataReports {
	s.ValidAnswerRate = &v
	return s
}

func (s *DescribeStatisticalDataResponseBodyStatisticalDataReports) Validate() error {
	return dara.Validate(s)
}
