// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSlsQueryJoin interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*SlsJoinCondition) *SlsQueryJoin
	GetConditions() []*SlsJoinCondition
	SetType(v string) *SlsQueryJoin
	GetType() *string
}

type SlsQueryJoin struct {
	Conditions []*SlsJoinCondition `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Type       *string             `json:"type,omitempty" xml:"type,omitempty"`
}

func (s SlsQueryJoin) String() string {
	return dara.Prettify(s)
}

func (s SlsQueryJoin) GoString() string {
	return s.String()
}

func (s *SlsQueryJoin) GetConditions() []*SlsJoinCondition {
	return s.Conditions
}

func (s *SlsQueryJoin) GetType() *string {
	return s.Type
}

func (s *SlsQueryJoin) SetConditions(v []*SlsJoinCondition) *SlsQueryJoin {
	s.Conditions = v
	return s
}

func (s *SlsQueryJoin) SetType(v string) *SlsQueryJoin {
	s.Type = &v
	return s
}

func (s *SlsQueryJoin) Validate() error {
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
