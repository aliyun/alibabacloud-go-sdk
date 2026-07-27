// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinings interface {
	dara.Model
	String() string
	GoString() string
	SetConditions(v []*JoinConditions) *Joinings
	GetConditions() []*JoinConditions
	SetType(v string) *Joinings
	GetType() *string
}

type Joinings struct {
	Conditions []*JoinConditions `json:"conditions,omitempty" xml:"conditions,omitempty" type:"Repeated"`
	Type       *string           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s Joinings) String() string {
	return dara.Prettify(s)
}

func (s Joinings) GoString() string {
	return s.String()
}

func (s *Joinings) GetConditions() []*JoinConditions {
	return s.Conditions
}

func (s *Joinings) GetType() *string {
	return s.Type
}

func (s *Joinings) SetConditions(v []*JoinConditions) *Joinings {
	s.Conditions = v
	return s
}

func (s *Joinings) SetType(v string) *Joinings {
	s.Type = &v
	return s
}

func (s *Joinings) Validate() error {
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
