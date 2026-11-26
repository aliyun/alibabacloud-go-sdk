// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScriptHotWordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateScriptHotWordRequest
	GetDescription() *string
	SetHotWord(v string) *UpdateScriptHotWordRequest
	GetHotWord() *string
	SetHotWordId(v string) *UpdateScriptHotWordRequest
	GetHotWordId() *string
}

type UpdateScriptHotWordRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	HotWord     *string `json:"hotWord,omitempty" xml:"hotWord,omitempty"`
	HotWordId   *string `json:"hotWordId,omitempty" xml:"hotWordId,omitempty"`
}

func (s UpdateScriptHotWordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScriptHotWordRequest) GoString() string {
	return s.String()
}

func (s *UpdateScriptHotWordRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateScriptHotWordRequest) GetHotWord() *string {
	return s.HotWord
}

func (s *UpdateScriptHotWordRequest) GetHotWordId() *string {
	return s.HotWordId
}

func (s *UpdateScriptHotWordRequest) SetDescription(v string) *UpdateScriptHotWordRequest {
	s.Description = &v
	return s
}

func (s *UpdateScriptHotWordRequest) SetHotWord(v string) *UpdateScriptHotWordRequest {
	s.HotWord = &v
	return s
}

func (s *UpdateScriptHotWordRequest) SetHotWordId(v string) *UpdateScriptHotWordRequest {
	s.HotWordId = &v
	return s
}

func (s *UpdateScriptHotWordRequest) Validate() error {
	return dara.Validate(s)
}
