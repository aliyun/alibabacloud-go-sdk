// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotationMissionDataSourceType(v int64) *ListIntentionsRequest
	GetAnnotationMissionDataSourceType() *int64
	SetBotId(v string) *ListIntentionsRequest
	GetBotId() *string
	SetEnvironment(v int64) *ListIntentionsRequest
	GetEnvironment() *int64
	SetInstanceId(v string) *ListIntentionsRequest
	GetInstanceId() *string
	SetIntentId(v int64) *ListIntentionsRequest
	GetIntentId() *int64
	SetPageIndex(v int32) *ListIntentionsRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *ListIntentionsRequest
	GetPageSize() *int32
	SetScriptId(v string) *ListIntentionsRequest
	GetScriptId() *string
	SetUserNick(v string) *ListIntentionsRequest
	GetUserNick() *string
}

type ListIntentionsRequest struct {
	// Source of annotation data.
	//
	// - 1: Outbound call
	//
	// - 2: Navigation
	//
	// example:
	//
	// 1
	AnnotationMissionDataSourceType *int64 `json:"AnnotationMissionDataSourceType,omitempty" xml:"AnnotationMissionDataSourceType,omitempty"`
	// Bot ID
	//
	// > You can obtain this parameter from the DescribeScript API. The corresponding parameter is ChatbotId.
	//
	// example:
	//
	// chatbot-cn-n7QmzrUnNe
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// Environment
	//
	// - 1: Private Cloud
	//
	// - 2: Public Cloud
	//
	// example:
	//
	// 2
	Environment *int64 `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// Instance ID
	//
	// example:
	//
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// Number of entries
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Intent name
	//
	// > You can use a keyword to perform a fuzzy query as a filter condition.
	//
	// example:
	//
	// 知道
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s ListIntentionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntentionsRequest) GoString() string {
	return s.String()
}

func (s *ListIntentionsRequest) GetAnnotationMissionDataSourceType() *int64 {
	return s.AnnotationMissionDataSourceType
}

func (s *ListIntentionsRequest) GetBotId() *string {
	return s.BotId
}

func (s *ListIntentionsRequest) GetEnvironment() *int64 {
	return s.Environment
}

func (s *ListIntentionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIntentionsRequest) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListIntentionsRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListIntentionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntentionsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListIntentionsRequest) GetUserNick() *string {
	return s.UserNick
}

func (s *ListIntentionsRequest) SetAnnotationMissionDataSourceType(v int64) *ListIntentionsRequest {
	s.AnnotationMissionDataSourceType = &v
	return s
}

func (s *ListIntentionsRequest) SetBotId(v string) *ListIntentionsRequest {
	s.BotId = &v
	return s
}

func (s *ListIntentionsRequest) SetEnvironment(v int64) *ListIntentionsRequest {
	s.Environment = &v
	return s
}

func (s *ListIntentionsRequest) SetInstanceId(v string) *ListIntentionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListIntentionsRequest) SetIntentId(v int64) *ListIntentionsRequest {
	s.IntentId = &v
	return s
}

func (s *ListIntentionsRequest) SetPageIndex(v int32) *ListIntentionsRequest {
	s.PageIndex = &v
	return s
}

func (s *ListIntentionsRequest) SetPageSize(v int32) *ListIntentionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListIntentionsRequest) SetScriptId(v string) *ListIntentionsRequest {
	s.ScriptId = &v
	return s
}

func (s *ListIntentionsRequest) SetUserNick(v string) *ListIntentionsRequest {
	s.UserNick = &v
	return s
}

func (s *ListIntentionsRequest) Validate() error {
	return dara.Validate(s)
}
