// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMarkResult interface {
	dara.Model
	String() string
	GoString() string
	SetIsNeedVoteJudge(v bool) *MarkResult
	GetIsNeedVoteJudge() *bool
	SetMarkResult(v string) *MarkResult
	GetMarkResult() *string
	SetMarkResultId(v string) *MarkResult
	GetMarkResultId() *string
	SetMarkTime(v string) *MarkResult
	GetMarkTime() *string
	SetMarkTitle(v string) *MarkResult
	GetMarkTitle() *string
	SetProgress(v string) *MarkResult
	GetProgress() *string
	SetQuestionId(v string) *MarkResult
	GetQuestionId() *string
	SetResultType(v string) *MarkResult
	GetResultType() *string
	SetUserMarkResultId(v string) *MarkResult
	GetUserMarkResultId() *string
	SetVersion(v string) *MarkResult
	GetVersion() *string
}

type MarkResult struct {
	// Indicates whether voting is required. Valid values:
	//
	// - True: Yes
	//
	// - False: No
	//
	// example:
	//
	// False
	IsNeedVoteJudge *bool `json:"IsNeedVoteJudge,omitempty" xml:"IsNeedVoteJudge,omitempty"`
	// Question result.
	//
	// example:
	//
	// b
	MarkResult *string `json:"MarkResult,omitempty" xml:"MarkResult,omitempty"`
	// Question ID.
	//
	// example:
	//
	// 1500***849089597440
	MarkResultId *string `json:"MarkResultId,omitempty" xml:"MarkResultId,omitempty"`
	// Annotation time.
	//
	// example:
	//
	// Mon Mar 07 17:02:48 CST 2022
	MarkTime *string `json:"MarkTime,omitempty" xml:"MarkTime,omitempty"`
	// Question name.
	//
	// example:
	//
	// 单选
	MarkTitle *string `json:"MarkTitle,omitempty" xml:"MarkTitle,omitempty"`
	// Progress. The return value is either None or data of JSON type. It includes:
	//
	// - Total: total number of results.
	//
	// - Finished: number of completed results.
	//
	// example:
	//
	// None
	Progress *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// Attached question.
	//
	// example:
	//
	// 1
	QuestionId *string `json:"QuestionId,omitempty" xml:"QuestionId,omitempty"`
	// Result type. Valid values:
	//
	// - RADIO: Radio
	//
	// - SLOT: Segment
	//
	// - INPUT: Fill-in-the-blank
	//
	// - CHECKBOX: Multiple Choice
	//
	// example:
	//
	// RADIO
	ResultType *string `json:"ResultType,omitempty" xml:"ResultType,omitempty"`
	// User annotation result ID.
	//
	// example:
	//
	// 1500***849358032896
	UserMarkResultId *string `json:"UserMarkResultId,omitempty" xml:"UserMarkResultId,omitempty"`
	// Version.
	//
	// example:
	//
	// 1646643768468
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s MarkResult) String() string {
	return dara.Prettify(s)
}

func (s MarkResult) GoString() string {
	return s.String()
}

func (s *MarkResult) GetIsNeedVoteJudge() *bool {
	return s.IsNeedVoteJudge
}

func (s *MarkResult) GetMarkResult() *string {
	return s.MarkResult
}

func (s *MarkResult) GetMarkResultId() *string {
	return s.MarkResultId
}

func (s *MarkResult) GetMarkTime() *string {
	return s.MarkTime
}

func (s *MarkResult) GetMarkTitle() *string {
	return s.MarkTitle
}

func (s *MarkResult) GetProgress() *string {
	return s.Progress
}

func (s *MarkResult) GetQuestionId() *string {
	return s.QuestionId
}

func (s *MarkResult) GetResultType() *string {
	return s.ResultType
}

func (s *MarkResult) GetUserMarkResultId() *string {
	return s.UserMarkResultId
}

func (s *MarkResult) GetVersion() *string {
	return s.Version
}

func (s *MarkResult) SetIsNeedVoteJudge(v bool) *MarkResult {
	s.IsNeedVoteJudge = &v
	return s
}

func (s *MarkResult) SetMarkResult(v string) *MarkResult {
	s.MarkResult = &v
	return s
}

func (s *MarkResult) SetMarkResultId(v string) *MarkResult {
	s.MarkResultId = &v
	return s
}

func (s *MarkResult) SetMarkTime(v string) *MarkResult {
	s.MarkTime = &v
	return s
}

func (s *MarkResult) SetMarkTitle(v string) *MarkResult {
	s.MarkTitle = &v
	return s
}

func (s *MarkResult) SetProgress(v string) *MarkResult {
	s.Progress = &v
	return s
}

func (s *MarkResult) SetQuestionId(v string) *MarkResult {
	s.QuestionId = &v
	return s
}

func (s *MarkResult) SetResultType(v string) *MarkResult {
	s.ResultType = &v
	return s
}

func (s *MarkResult) SetUserMarkResultId(v string) *MarkResult {
	s.UserMarkResultId = &v
	return s
}

func (s *MarkResult) SetVersion(v string) *MarkResult {
	s.Version = &v
	return s
}

func (s *MarkResult) Validate() error {
	return dara.Validate(s)
}
