// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecommendEventRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateRecommendEventRuleRequest
	GetLang() *string
	SetEventCode(v string) *CreateRecommendEventRuleRequest
	GetEventCode() *string
	SetEventName(v string) *CreateRecommendEventRuleRequest
	GetEventName() *string
	SetRecommendRuleIdsStr(v string) *CreateRecommendEventRuleRequest
	GetRecommendRuleIdsStr() *string
	SetRegId(v string) *CreateRecommendEventRuleRequest
	GetRegId() *string
	SetTaskId(v int64) *CreateRecommendEventRuleRequest
	GetTaskId() *int64
}

type CreateRecommendEventRuleRequest struct {
	// Set the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_ahqido8038
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 用户昵称文本审核检测结果
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
	// Strategy ID.
	//
	// example:
	//
	// [\\"100234\\",\\"100235\\"]
	RecommendRuleIdsStr *string `json:"recommendRuleIdsStr,omitempty" xml:"recommendRuleIdsStr,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 6770764
	TaskId *int64 `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateRecommendEventRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecommendEventRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRecommendEventRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateRecommendEventRuleRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *CreateRecommendEventRuleRequest) GetEventName() *string {
	return s.EventName
}

func (s *CreateRecommendEventRuleRequest) GetRecommendRuleIdsStr() *string {
	return s.RecommendRuleIdsStr
}

func (s *CreateRecommendEventRuleRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateRecommendEventRuleRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *CreateRecommendEventRuleRequest) SetLang(v string) *CreateRecommendEventRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateRecommendEventRuleRequest) SetEventCode(v string) *CreateRecommendEventRuleRequest {
	s.EventCode = &v
	return s
}

func (s *CreateRecommendEventRuleRequest) SetEventName(v string) *CreateRecommendEventRuleRequest {
	s.EventName = &v
	return s
}

func (s *CreateRecommendEventRuleRequest) SetRecommendRuleIdsStr(v string) *CreateRecommendEventRuleRequest {
	s.RecommendRuleIdsStr = &v
	return s
}

func (s *CreateRecommendEventRuleRequest) SetRegId(v string) *CreateRecommendEventRuleRequest {
	s.RegId = &v
	return s
}

func (s *CreateRecommendEventRuleRequest) SetTaskId(v int64) *CreateRecommendEventRuleRequest {
	s.TaskId = &v
	return s
}

func (s *CreateRecommendEventRuleRequest) Validate() error {
	return dara.Validate(s)
}
