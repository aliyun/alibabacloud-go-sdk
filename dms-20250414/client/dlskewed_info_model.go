// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDLSkewedInfo interface {
	dara.Model
	String() string
	GoString() string
	SetSkewedColNames(v []*string) *DLSkewedInfo
	GetSkewedColNames() []*string
	SetSkewedColValueLocationMaps(v map[string]interface{}) *DLSkewedInfo
	GetSkewedColValueLocationMaps() map[string]interface{}
	SetSkewedColValues(v [][]*string) *DLSkewedInfo
	GetSkewedColValues() [][]*string
}

type DLSkewedInfo struct {
	SkewedColNames []*string `json:"SkewedColNames,omitempty" xml:"SkewedColNames,omitempty" type:"Repeated"`
	// example:
	//
	// {“col1”:"val1"}
	SkewedColValueLocationMaps map[string]interface{} `json:"SkewedColValueLocationMaps,omitempty" xml:"SkewedColValueLocationMaps,omitempty"`
	SkewedColValues            [][]*string            `json:"SkewedColValues,omitempty" xml:"SkewedColValues,omitempty" type:"Repeated"`
}

func (s DLSkewedInfo) String() string {
	return dara.Prettify(s)
}

func (s DLSkewedInfo) GoString() string {
	return s.String()
}

func (s *DLSkewedInfo) GetSkewedColNames() []*string {
	return s.SkewedColNames
}

func (s *DLSkewedInfo) GetSkewedColValueLocationMaps() map[string]interface{} {
	return s.SkewedColValueLocationMaps
}

func (s *DLSkewedInfo) GetSkewedColValues() [][]*string {
	return s.SkewedColValues
}

func (s *DLSkewedInfo) SetSkewedColNames(v []*string) *DLSkewedInfo {
	s.SkewedColNames = v
	return s
}

func (s *DLSkewedInfo) SetSkewedColValueLocationMaps(v map[string]interface{}) *DLSkewedInfo {
	s.SkewedColValueLocationMaps = v
	return s
}

func (s *DLSkewedInfo) SetSkewedColValues(v [][]*string) *DLSkewedInfo {
	s.SkewedColValues = v
	return s
}

func (s *DLSkewedInfo) Validate() error {
	return dara.Validate(s)
}
