// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillFileDetectResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSkillFileDetectResultResponseBody
	GetCode() *string
	SetHashKey(v string) *GetSkillFileDetectResultResponseBody
	GetHashKey() *string
	SetMessage(v string) *GetSkillFileDetectResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSkillFileDetectResultResponseBody
	GetRequestId() *string
	SetResult(v int32) *GetSkillFileDetectResultResponseBody
	GetResult() *int32
	SetScore(v int32) *GetSkillFileDetectResultResponseBody
	GetScore() *int32
}

type GetSkillFileDetectResultResponseBody struct {
	// The result code. A value of "200" indicates success. Other values indicate an error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The unique identifier for the detection task.
	//
	// example:
	//
	// 8cfb1102cc5f56fe752f6e9b8c6f7f3d
	HashKey *string `json:"HashKey,omitempty" xml:"HashKey,omitempty"`
	// The response message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F976EF8-C10A-57DC-917C-BB7BEB508FFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The file detection result. Valid values:
	//
	// 0: safe file
	//
	// 1: suspicious file
	//
	// 3: The detection is in progress.
	//
	// example:
	//
	// 0
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// The file detection score. The score ranges correspond to the following threat levels:
	//
	// 0–60: safe
	//
	// 61–70: risk
	//
	// 71–80: suspicious
	//
	// 81–100: malicious
	//
	// example:
	//
	// 10
	Score *int32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s GetSkillFileDetectResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillFileDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillFileDetectResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSkillFileDetectResultResponseBody) GetHashKey() *string {
	return s.HashKey
}

func (s *GetSkillFileDetectResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSkillFileDetectResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillFileDetectResultResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *GetSkillFileDetectResultResponseBody) GetScore() *int32 {
	return s.Score
}

func (s *GetSkillFileDetectResultResponseBody) SetCode(v string) *GetSkillFileDetectResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetSkillFileDetectResultResponseBody) SetHashKey(v string) *GetSkillFileDetectResultResponseBody {
	s.HashKey = &v
	return s
}

func (s *GetSkillFileDetectResultResponseBody) SetMessage(v string) *GetSkillFileDetectResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetSkillFileDetectResultResponseBody) SetRequestId(v string) *GetSkillFileDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillFileDetectResultResponseBody) SetResult(v int32) *GetSkillFileDetectResultResponseBody {
	s.Result = &v
	return s
}

func (s *GetSkillFileDetectResultResponseBody) SetScore(v int32) *GetSkillFileDetectResultResponseBody {
	s.Score = &v
	return s
}

func (s *GetSkillFileDetectResultResponseBody) Validate() error {
	return dara.Validate(s)
}
