// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveStyleLearningResultShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SaveStyleLearningResultShrinkRequest
	GetAgentKey() *string
	SetAigcResult(v string) *SaveStyleLearningResultShrinkRequest
	GetAigcResult() *string
	SetCustomTextIdListShrink(v string) *SaveStyleLearningResultShrinkRequest
	GetCustomTextIdListShrink() *string
	SetMaterialIdListShrink(v string) *SaveStyleLearningResultShrinkRequest
	GetMaterialIdListShrink() *string
	SetRewriteResult(v string) *SaveStyleLearningResultShrinkRequest
	GetRewriteResult() *string
	SetStyleName(v string) *SaveStyleLearningResultShrinkRequest
	GetStyleName() *string
	SetTaskId(v string) *SaveStyleLearningResultShrinkRequest
	GetTaskId() *string
}

type SaveStyleLearningResultShrinkRequest struct {
	// example:
	//
	// xxxxx_p_efm
	AgentKey               *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AigcResult             *string `json:"AigcResult,omitempty" xml:"AigcResult,omitempty"`
	CustomTextIdListShrink *string `json:"CustomTextIdList,omitempty" xml:"CustomTextIdList,omitempty"`
	MaterialIdListShrink   *string `json:"MaterialIdList,omitempty" xml:"MaterialIdList,omitempty"`
	RewriteResult          *string `json:"RewriteResult,omitempty" xml:"RewriteResult,omitempty"`
	StyleName              *string `json:"StyleName,omitempty" xml:"StyleName,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SaveStyleLearningResultShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveStyleLearningResultShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveStyleLearningResultShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SaveStyleLearningResultShrinkRequest) GetAigcResult() *string {
	return s.AigcResult
}

func (s *SaveStyleLearningResultShrinkRequest) GetCustomTextIdListShrink() *string {
	return s.CustomTextIdListShrink
}

func (s *SaveStyleLearningResultShrinkRequest) GetMaterialIdListShrink() *string {
	return s.MaterialIdListShrink
}

func (s *SaveStyleLearningResultShrinkRequest) GetRewriteResult() *string {
	return s.RewriteResult
}

func (s *SaveStyleLearningResultShrinkRequest) GetStyleName() *string {
	return s.StyleName
}

func (s *SaveStyleLearningResultShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SaveStyleLearningResultShrinkRequest) SetAgentKey(v string) *SaveStyleLearningResultShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveStyleLearningResultShrinkRequest) SetAigcResult(v string) *SaveStyleLearningResultShrinkRequest {
	s.AigcResult = &v
	return s
}

func (s *SaveStyleLearningResultShrinkRequest) SetCustomTextIdListShrink(v string) *SaveStyleLearningResultShrinkRequest {
	s.CustomTextIdListShrink = &v
	return s
}

func (s *SaveStyleLearningResultShrinkRequest) SetMaterialIdListShrink(v string) *SaveStyleLearningResultShrinkRequest {
	s.MaterialIdListShrink = &v
	return s
}

func (s *SaveStyleLearningResultShrinkRequest) SetRewriteResult(v string) *SaveStyleLearningResultShrinkRequest {
	s.RewriteResult = &v
	return s
}

func (s *SaveStyleLearningResultShrinkRequest) SetStyleName(v string) *SaveStyleLearningResultShrinkRequest {
	s.StyleName = &v
	return s
}

func (s *SaveStyleLearningResultShrinkRequest) SetTaskId(v string) *SaveStyleLearningResultShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *SaveStyleLearningResultShrinkRequest) Validate() error {
	return dara.Validate(s)
}
