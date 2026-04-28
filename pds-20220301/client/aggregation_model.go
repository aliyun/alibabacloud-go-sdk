// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAggregation interface {
	dara.Model
	String() string
	GoString() string
	SetField(v []byte) *Aggregation
	GetField() []byte
	SetGroups(v []*AggregationsGroup) *Aggregation
	GetGroups() []*AggregationsGroup
	SetOperation(v []byte) *Aggregation
	GetOperation() []byte
	SetValue(v float64) *Aggregation
	GetValue() *float64
}

type Aggregation struct {
	Field     []byte               `json:"field,omitempty" xml:"field,omitempty"`
	Groups    []*AggregationsGroup `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	Operation []byte               `json:"operation,omitempty" xml:"operation,omitempty"`
	Value     *float64             `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Aggregation) String() string {
	return dara.Prettify(s)
}

func (s Aggregation) GoString() string {
	return s.String()
}

func (s *Aggregation) GetField() []byte {
	return s.Field
}

func (s *Aggregation) GetGroups() []*AggregationsGroup {
	return s.Groups
}

func (s *Aggregation) GetOperation() []byte {
	return s.Operation
}

func (s *Aggregation) GetValue() *float64 {
	return s.Value
}

func (s *Aggregation) SetField(v []byte) *Aggregation {
	s.Field = v
	return s
}

func (s *Aggregation) SetGroups(v []*AggregationsGroup) *Aggregation {
	s.Groups = v
	return s
}

func (s *Aggregation) SetOperation(v []byte) *Aggregation {
	s.Operation = v
	return s
}

func (s *Aggregation) SetValue(v float64) *Aggregation {
	s.Value = &v
	return s
}

func (s *Aggregation) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
