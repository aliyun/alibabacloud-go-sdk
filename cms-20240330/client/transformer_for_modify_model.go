// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformerForModify interface {
	dara.Model
	String() string
	GoString() string
	SetActions(v []*TransformAction) *TransformerForModify
	GetActions() []*TransformAction
	SetDescription(v string) *TransformerForModify
	GetDescription() *string
	SetFilterSetting(v *FilterSetting) *TransformerForModify
	GetFilterSetting() *FilterSetting
	SetQuitAfterMatch(v bool) *TransformerForModify
	GetQuitAfterMatch() *bool
	SetSortId(v int32) *TransformerForModify
	GetSortId() *int32
	SetTransformerName(v string) *TransformerForModify
	GetTransformerName() *string
}

type TransformerForModify struct {
	Actions        []*TransformAction `json:"actions,omitempty" xml:"actions,omitempty" type:"Repeated"`
	Description    *string            `json:"description,omitempty" xml:"description,omitempty"`
	FilterSetting  *FilterSetting     `json:"filterSetting,omitempty" xml:"filterSetting,omitempty"`
	QuitAfterMatch *bool              `json:"quitAfterMatch,omitempty" xml:"quitAfterMatch,omitempty"`
	SortId         *int32             `json:"sortId,omitempty" xml:"sortId,omitempty"`
	// This parameter is required.
	TransformerName *string `json:"transformerName,omitempty" xml:"transformerName,omitempty"`
}

func (s TransformerForModify) String() string {
	return dara.Prettify(s)
}

func (s TransformerForModify) GoString() string {
	return s.String()
}

func (s *TransformerForModify) GetActions() []*TransformAction {
	return s.Actions
}

func (s *TransformerForModify) GetDescription() *string {
	return s.Description
}

func (s *TransformerForModify) GetFilterSetting() *FilterSetting {
	return s.FilterSetting
}

func (s *TransformerForModify) GetQuitAfterMatch() *bool {
	return s.QuitAfterMatch
}

func (s *TransformerForModify) GetSortId() *int32 {
	return s.SortId
}

func (s *TransformerForModify) GetTransformerName() *string {
	return s.TransformerName
}

func (s *TransformerForModify) SetActions(v []*TransformAction) *TransformerForModify {
	s.Actions = v
	return s
}

func (s *TransformerForModify) SetDescription(v string) *TransformerForModify {
	s.Description = &v
	return s
}

func (s *TransformerForModify) SetFilterSetting(v *FilterSetting) *TransformerForModify {
	s.FilterSetting = v
	return s
}

func (s *TransformerForModify) SetQuitAfterMatch(v bool) *TransformerForModify {
	s.QuitAfterMatch = &v
	return s
}

func (s *TransformerForModify) SetSortId(v int32) *TransformerForModify {
	s.SortId = &v
	return s
}

func (s *TransformerForModify) SetTransformerName(v string) *TransformerForModify {
	s.TransformerName = &v
	return s
}

func (s *TransformerForModify) Validate() error {
	if s.Actions != nil {
		for _, item := range s.Actions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FilterSetting != nil {
		if err := s.FilterSetting.Validate(); err != nil {
			return err
		}
	}
	return nil
}
