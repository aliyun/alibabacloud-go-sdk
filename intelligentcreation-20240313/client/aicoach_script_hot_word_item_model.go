// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAICoachScriptHotWordItem interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v string) *AICoachScriptHotWordItem
	GetCreatorId() *string
	SetCreatorName(v string) *AICoachScriptHotWordItem
	GetCreatorName() *string
	SetDescription(v string) *AICoachScriptHotWordItem
	GetDescription() *string
	SetGmtCreate(v string) *AICoachScriptHotWordItem
	GetGmtCreate() *string
	SetGmtModified(v string) *AICoachScriptHotWordItem
	GetGmtModified() *string
	SetHotWord(v string) *AICoachScriptHotWordItem
	GetHotWord() *string
	SetHotWordId(v string) *AICoachScriptHotWordItem
	GetHotWordId() *string
	SetModifierId(v string) *AICoachScriptHotWordItem
	GetModifierId() *string
	SetModifierName(v string) *AICoachScriptHotWordItem
	GetModifierName() *string
	SetScriptRecordId(v string) *AICoachScriptHotWordItem
	GetScriptRecordId() *string
}

type AICoachScriptHotWordItem struct {
	// example:
	//
	// 849843
	CreatorId *string `json:"creatorId,omitempty" xml:"creatorId,omitempty"`
	// example:
	//
	// username
	CreatorName *string `json:"creatorName,omitempty" xml:"creatorName,omitempty"`
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 2026-11-13 12:01:12
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 2026-11-13 12:01:12
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// test
	HotWord *string `json:"hotWord,omitempty" xml:"hotWord,omitempty"`
	// example:
	//
	// 34532
	HotWordId *string `json:"hotWordId,omitempty" xml:"hotWordId,omitempty"`
	// example:
	//
	// 892340
	ModifierId *string `json:"modifierId,omitempty" xml:"modifierId,omitempty"`
	// example:
	//
	// username
	ModifierName *string `json:"modifierName,omitempty" xml:"modifierName,omitempty"`
	// example:
	//
	// 8491340
	ScriptRecordId *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
}

func (s AICoachScriptHotWordItem) String() string {
	return dara.Prettify(s)
}

func (s AICoachScriptHotWordItem) GoString() string {
	return s.String()
}

func (s *AICoachScriptHotWordItem) GetCreatorId() *string {
	return s.CreatorId
}

func (s *AICoachScriptHotWordItem) GetCreatorName() *string {
	return s.CreatorName
}

func (s *AICoachScriptHotWordItem) GetDescription() *string {
	return s.Description
}

func (s *AICoachScriptHotWordItem) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *AICoachScriptHotWordItem) GetGmtModified() *string {
	return s.GmtModified
}

func (s *AICoachScriptHotWordItem) GetHotWord() *string {
	return s.HotWord
}

func (s *AICoachScriptHotWordItem) GetHotWordId() *string {
	return s.HotWordId
}

func (s *AICoachScriptHotWordItem) GetModifierId() *string {
	return s.ModifierId
}

func (s *AICoachScriptHotWordItem) GetModifierName() *string {
	return s.ModifierName
}

func (s *AICoachScriptHotWordItem) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *AICoachScriptHotWordItem) SetCreatorId(v string) *AICoachScriptHotWordItem {
	s.CreatorId = &v
	return s
}

func (s *AICoachScriptHotWordItem) SetCreatorName(v string) *AICoachScriptHotWordItem {
	s.CreatorName = &v
	return s
}

func (s *AICoachScriptHotWordItem) SetDescription(v string) *AICoachScriptHotWordItem {
	s.Description = &v
	return s
}

func (s *AICoachScriptHotWordItem) SetGmtCreate(v string) *AICoachScriptHotWordItem {
	s.GmtCreate = &v
	return s
}

func (s *AICoachScriptHotWordItem) SetGmtModified(v string) *AICoachScriptHotWordItem {
	s.GmtModified = &v
	return s
}

func (s *AICoachScriptHotWordItem) SetHotWord(v string) *AICoachScriptHotWordItem {
	s.HotWord = &v
	return s
}

func (s *AICoachScriptHotWordItem) SetHotWordId(v string) *AICoachScriptHotWordItem {
	s.HotWordId = &v
	return s
}

func (s *AICoachScriptHotWordItem) SetModifierId(v string) *AICoachScriptHotWordItem {
	s.ModifierId = &v
	return s
}

func (s *AICoachScriptHotWordItem) SetModifierName(v string) *AICoachScriptHotWordItem {
	s.ModifierName = &v
	return s
}

func (s *AICoachScriptHotWordItem) SetScriptRecordId(v string) *AICoachScriptHotWordItem {
	s.ScriptRecordId = &v
	return s
}

func (s *AICoachScriptHotWordItem) Validate() error {
	return dara.Validate(s)
}
