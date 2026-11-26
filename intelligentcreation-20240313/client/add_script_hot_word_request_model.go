// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddScriptHotWordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AddScriptHotWordRequest
	GetDescription() *string
	SetHotWord(v string) *AddScriptHotWordRequest
	GetHotWord() *string
	SetScriptRecordId(v string) *AddScriptHotWordRequest
	GetScriptRecordId() *string
}

type AddScriptHotWordRequest struct {
	Description    *string `json:"description,omitempty" xml:"description,omitempty"`
	HotWord        *string `json:"hotWord,omitempty" xml:"hotWord,omitempty"`
	ScriptRecordId *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
}

func (s AddScriptHotWordRequest) String() string {
	return dara.Prettify(s)
}

func (s AddScriptHotWordRequest) GoString() string {
	return s.String()
}

func (s *AddScriptHotWordRequest) GetDescription() *string {
	return s.Description
}

func (s *AddScriptHotWordRequest) GetHotWord() *string {
	return s.HotWord
}

func (s *AddScriptHotWordRequest) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *AddScriptHotWordRequest) SetDescription(v string) *AddScriptHotWordRequest {
	s.Description = &v
	return s
}

func (s *AddScriptHotWordRequest) SetHotWord(v string) *AddScriptHotWordRequest {
	s.HotWord = &v
	return s
}

func (s *AddScriptHotWordRequest) SetScriptRecordId(v string) *AddScriptHotWordRequest {
	s.ScriptRecordId = &v
	return s
}

func (s *AddScriptHotWordRequest) Validate() error {
	return dara.Validate(s)
}
