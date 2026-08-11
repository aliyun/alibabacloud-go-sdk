// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChartsValue interface {
	dara.Model
	String() string
	GoString() string
	SetX(v []*string) *ChartsValue
	GetX() []*string
	SetY(v []*ChartsValueY) *ChartsValue
	GetY() []*ChartsValueY
}

type ChartsValue struct {
	// The X-axis.
	X []*string `json:"X,omitempty" xml:"X,omitempty" type:"Repeated"`
	// The Y-axis.
	Y []*ChartsValueY `json:"Y,omitempty" xml:"Y,omitempty" type:"Repeated"`
}

func (s ChartsValue) String() string {
	return dara.Prettify(s)
}

func (s ChartsValue) GoString() string {
	return s.String()
}

func (s *ChartsValue) GetX() []*string {
	return s.X
}

func (s *ChartsValue) GetY() []*ChartsValueY {
	return s.Y
}

func (s *ChartsValue) SetX(v []*string) *ChartsValue {
	s.X = v
	return s
}

func (s *ChartsValue) SetY(v []*ChartsValueY) *ChartsValue {
	s.Y = v
	return s
}

func (s *ChartsValue) Validate() error {
	if s.Y != nil {
		for _, item := range s.Y {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ChartsValueY struct {
	// The name.
	//
	// example:
	//
	// cn-shanghai
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The QPS at the point in time.
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ChartsValueY) String() string {
	return dara.Prettify(s)
}

func (s ChartsValueY) GoString() string {
	return s.String()
}

func (s *ChartsValueY) GetName() *string {
	return s.Name
}

func (s *ChartsValueY) GetData() []*int64 {
	return s.Data
}

func (s *ChartsValueY) SetName(v string) *ChartsValueY {
	s.Name = &v
	return s
}

func (s *ChartsValueY) SetData(v []*int64) *ChartsValueY {
	s.Data = v
	return s
}

func (s *ChartsValueY) Validate() error {
	return dara.Validate(s)
}
