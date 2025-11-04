// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotword interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *Hotword
	GetLanguage() *string
	SetText(v string) *Hotword
	GetText() *string
	SetTranspositionResultList(v []*TranspositionResult) *Hotword
	GetTranspositionResultList() []*TranspositionResult
	SetWeight(v int32) *Hotword
	GetWeight() *int32
}

type Hotword struct {
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// hello
	Text                    *string                `json:"Text,omitempty" xml:"Text,omitempty"`
	TranspositionResultList []*TranspositionResult `json:"TranspositionResultList,omitempty" xml:"TranspositionResultList,omitempty" type:"Repeated"`
	Weight                  *int32                 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s Hotword) String() string {
	return dara.Prettify(s)
}

func (s Hotword) GoString() string {
	return s.String()
}

func (s *Hotword) GetLanguage() *string {
	return s.Language
}

func (s *Hotword) GetText() *string {
	return s.Text
}

func (s *Hotword) GetTranspositionResultList() []*TranspositionResult {
	return s.TranspositionResultList
}

func (s *Hotword) GetWeight() *int32 {
	return s.Weight
}

func (s *Hotword) SetLanguage(v string) *Hotword {
	s.Language = &v
	return s
}

func (s *Hotword) SetText(v string) *Hotword {
	s.Text = &v
	return s
}

func (s *Hotword) SetTranspositionResultList(v []*TranspositionResult) *Hotword {
	s.TranspositionResultList = v
	return s
}

func (s *Hotword) SetWeight(v int32) *Hotword {
	s.Weight = &v
	return s
}

func (s *Hotword) Validate() error {
	if s.TranspositionResultList != nil {
		for _, item := range s.TranspositionResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
