// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveStyleLearningResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *SaveStyleLearningResultRequest
	GetAgentKey() *string
	SetAigcResult(v string) *SaveStyleLearningResultRequest
	GetAigcResult() *string
	SetCustomTextIdList(v []*int64) *SaveStyleLearningResultRequest
	GetCustomTextIdList() []*int64
	SetMaterialIdList(v []*int64) *SaveStyleLearningResultRequest
	GetMaterialIdList() []*int64
	SetRewriteResult(v string) *SaveStyleLearningResultRequest
	GetRewriteResult() *string
	SetStyleName(v string) *SaveStyleLearningResultRequest
	GetStyleName() *string
	SetTaskId(v string) *SaveStyleLearningResultRequest
	GetTaskId() *string
}

type SaveStyleLearningResultRequest struct {
	// example:
	//
	// xxxxx_p_efm
	AgentKey         *string  `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	AigcResult       *string  `json:"AigcResult,omitempty" xml:"AigcResult,omitempty"`
	CustomTextIdList []*int64 `json:"CustomTextIdList,omitempty" xml:"CustomTextIdList,omitempty" type:"Repeated"`
	MaterialIdList   []*int64 `json:"MaterialIdList,omitempty" xml:"MaterialIdList,omitempty" type:"Repeated"`
	RewriteResult    *string  `json:"RewriteResult,omitempty" xml:"RewriteResult,omitempty"`
	StyleName        *string  `json:"StyleName,omitempty" xml:"StyleName,omitempty"`
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SaveStyleLearningResultRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveStyleLearningResultRequest) GoString() string {
	return s.String()
}

func (s *SaveStyleLearningResultRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *SaveStyleLearningResultRequest) GetAigcResult() *string {
	return s.AigcResult
}

func (s *SaveStyleLearningResultRequest) GetCustomTextIdList() []*int64 {
	return s.CustomTextIdList
}

func (s *SaveStyleLearningResultRequest) GetMaterialIdList() []*int64 {
	return s.MaterialIdList
}

func (s *SaveStyleLearningResultRequest) GetRewriteResult() *string {
	return s.RewriteResult
}

func (s *SaveStyleLearningResultRequest) GetStyleName() *string {
	return s.StyleName
}

func (s *SaveStyleLearningResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *SaveStyleLearningResultRequest) SetAgentKey(v string) *SaveStyleLearningResultRequest {
	s.AgentKey = &v
	return s
}

func (s *SaveStyleLearningResultRequest) SetAigcResult(v string) *SaveStyleLearningResultRequest {
	s.AigcResult = &v
	return s
}

func (s *SaveStyleLearningResultRequest) SetCustomTextIdList(v []*int64) *SaveStyleLearningResultRequest {
	s.CustomTextIdList = v
	return s
}

func (s *SaveStyleLearningResultRequest) SetMaterialIdList(v []*int64) *SaveStyleLearningResultRequest {
	s.MaterialIdList = v
	return s
}

func (s *SaveStyleLearningResultRequest) SetRewriteResult(v string) *SaveStyleLearningResultRequest {
	s.RewriteResult = &v
	return s
}

func (s *SaveStyleLearningResultRequest) SetStyleName(v string) *SaveStyleLearningResultRequest {
	s.StyleName = &v
	return s
}

func (s *SaveStyleLearningResultRequest) SetTaskId(v string) *SaveStyleLearningResultRequest {
	s.TaskId = &v
	return s
}

func (s *SaveStyleLearningResultRequest) Validate() error {
	return dara.Validate(s)
}
