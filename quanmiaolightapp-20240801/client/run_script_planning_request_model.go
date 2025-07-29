// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunScriptPlanningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalNote(v string) *RunScriptPlanningRequest
	GetAdditionalNote() *string
	SetDialogueInScene(v bool) *RunScriptPlanningRequest
	GetDialogueInScene() *bool
	SetPlotConflict(v bool) *RunScriptPlanningRequest
	GetPlotConflict() *bool
	SetScriptName(v string) *RunScriptPlanningRequest
	GetScriptName() *string
	SetScriptShotCount(v int32) *RunScriptPlanningRequest
	GetScriptShotCount() *int32
	SetScriptSummary(v string) *RunScriptPlanningRequest
	GetScriptSummary() *string
	SetScriptTypeKeyword(v string) *RunScriptPlanningRequest
	GetScriptTypeKeyword() *string
}

type RunScriptPlanningRequest struct {
	// example:
	//
	// 故事尽可能狗血
	AdditionalNote  *string `json:"additionalNote,omitempty" xml:"additionalNote,omitempty"`
	DialogueInScene *bool   `json:"dialogueInScene,omitempty" xml:"dialogueInScene,omitempty"`
	PlotConflict    *bool   `json:"plotConflict,omitempty" xml:"plotConflict,omitempty"`
	// example:
	//
	// 都市战神
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// example:
	//
	// 3
	ScriptShotCount *int32 `json:"scriptShotCount,omitempty" xml:"scriptShotCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 在一个宁静的小镇上，每个家庭都在同一天收到一个神秘的、没有标记的包裹。
	ScriptSummary *string `json:"scriptSummary,omitempty" xml:"scriptSummary,omitempty"`
	// example:
	//
	// 现代，都市，爱情，玄幻
	ScriptTypeKeyword *string `json:"scriptTypeKeyword,omitempty" xml:"scriptTypeKeyword,omitempty"`
}

func (s RunScriptPlanningRequest) String() string {
	return dara.Prettify(s)
}

func (s RunScriptPlanningRequest) GoString() string {
	return s.String()
}

func (s *RunScriptPlanningRequest) GetAdditionalNote() *string {
	return s.AdditionalNote
}

func (s *RunScriptPlanningRequest) GetDialogueInScene() *bool {
	return s.DialogueInScene
}

func (s *RunScriptPlanningRequest) GetPlotConflict() *bool {
	return s.PlotConflict
}

func (s *RunScriptPlanningRequest) GetScriptName() *string {
	return s.ScriptName
}

func (s *RunScriptPlanningRequest) GetScriptShotCount() *int32 {
	return s.ScriptShotCount
}

func (s *RunScriptPlanningRequest) GetScriptSummary() *string {
	return s.ScriptSummary
}

func (s *RunScriptPlanningRequest) GetScriptTypeKeyword() *string {
	return s.ScriptTypeKeyword
}

func (s *RunScriptPlanningRequest) SetAdditionalNote(v string) *RunScriptPlanningRequest {
	s.AdditionalNote = &v
	return s
}

func (s *RunScriptPlanningRequest) SetDialogueInScene(v bool) *RunScriptPlanningRequest {
	s.DialogueInScene = &v
	return s
}

func (s *RunScriptPlanningRequest) SetPlotConflict(v bool) *RunScriptPlanningRequest {
	s.PlotConflict = &v
	return s
}

func (s *RunScriptPlanningRequest) SetScriptName(v string) *RunScriptPlanningRequest {
	s.ScriptName = &v
	return s
}

func (s *RunScriptPlanningRequest) SetScriptShotCount(v int32) *RunScriptPlanningRequest {
	s.ScriptShotCount = &v
	return s
}

func (s *RunScriptPlanningRequest) SetScriptSummary(v string) *RunScriptPlanningRequest {
	s.ScriptSummary = &v
	return s
}

func (s *RunScriptPlanningRequest) SetScriptTypeKeyword(v string) *RunScriptPlanningRequest {
	s.ScriptTypeKeyword = &v
	return s
}

func (s *RunScriptPlanningRequest) Validate() error {
	return dara.Validate(s)
}
