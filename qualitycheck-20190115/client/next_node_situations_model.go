// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNextNodeSituations interface {
	dara.Model
	String() string
	GoString() string
	SetConditionGroup(v []*NextNodeSituationsConditionGroup) *NextNodeSituations
	GetConditionGroup() []*NextNodeSituationsConditionGroup
	SetType(v string) *NextNodeSituations
	GetType() *string
}

type NextNodeSituations struct {
	ConditionGroup []*NextNodeSituationsConditionGroup `json:"ConditionGroup,omitempty" xml:"ConditionGroup,omitempty" type:"Repeated"`
	Type           *string                             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s NextNodeSituations) String() string {
	return dara.Prettify(s)
}

func (s NextNodeSituations) GoString() string {
	return s.String()
}

func (s *NextNodeSituations) GetConditionGroup() []*NextNodeSituationsConditionGroup {
	return s.ConditionGroup
}

func (s *NextNodeSituations) GetType() *string {
	return s.Type
}

func (s *NextNodeSituations) SetConditionGroup(v []*NextNodeSituationsConditionGroup) *NextNodeSituations {
	s.ConditionGroup = v
	return s
}

func (s *NextNodeSituations) SetType(v string) *NextNodeSituations {
	s.Type = &v
	return s
}

func (s *NextNodeSituations) Validate() error {
	if s.ConditionGroup != nil {
		for _, item := range s.ConditionGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type NextNodeSituationsConditionGroup struct {
	Conditions []*JudgeNodeMetaDesc `json:"Conditions,omitempty" xml:"Conditions,omitempty" type:"Repeated"`
	Type       *string              `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s NextNodeSituationsConditionGroup) String() string {
	return dara.Prettify(s)
}

func (s NextNodeSituationsConditionGroup) GoString() string {
	return s.String()
}

func (s *NextNodeSituationsConditionGroup) GetConditions() []*JudgeNodeMetaDesc {
	return s.Conditions
}

func (s *NextNodeSituationsConditionGroup) GetType() *string {
	return s.Type
}

func (s *NextNodeSituationsConditionGroup) SetConditions(v []*JudgeNodeMetaDesc) *NextNodeSituationsConditionGroup {
	s.Conditions = v
	return s
}

func (s *NextNodeSituationsConditionGroup) SetType(v string) *NextNodeSituationsConditionGroup {
	s.Type = &v
	return s
}

func (s *NextNodeSituationsConditionGroup) Validate() error {
	if s.Conditions != nil {
		for _, item := range s.Conditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
