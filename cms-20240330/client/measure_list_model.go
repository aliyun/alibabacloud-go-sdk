// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeasureList interface {
	dara.Model
	String() string
	GoString() string
	SetGroupBy(v []*string) *MeasureList
	GetGroupBy() []*string
	SetMeasureCode(v string) *MeasureList
	GetMeasureCode() *string
	SetWindowSecs(v int32) *MeasureList
	GetWindowSecs() *int32
}

type MeasureList struct {
	// The grouping dimension.
	GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
	// The metric code.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_usage
	MeasureCode *string `json:"measureCode,omitempty" xml:"measureCode,omitempty"`
	// The query time window, in seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	WindowSecs *int32 `json:"windowSecs,omitempty" xml:"windowSecs,omitempty"`
}

func (s MeasureList) String() string {
	return dara.Prettify(s)
}

func (s MeasureList) GoString() string {
	return s.String()
}

func (s *MeasureList) GetGroupBy() []*string {
	return s.GroupBy
}

func (s *MeasureList) GetMeasureCode() *string {
	return s.MeasureCode
}

func (s *MeasureList) GetWindowSecs() *int32 {
	return s.WindowSecs
}

func (s *MeasureList) SetGroupBy(v []*string) *MeasureList {
	s.GroupBy = v
	return s
}

func (s *MeasureList) SetMeasureCode(v string) *MeasureList {
	s.MeasureCode = &v
	return s
}

func (s *MeasureList) SetWindowSecs(v int32) *MeasureList {
	s.WindowSecs = &v
	return s
}

func (s *MeasureList) Validate() error {
	return dara.Validate(s)
}
