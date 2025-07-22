// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataValue interface {
	dara.Model
	String() string
	GoString() string
	SetTimestamp(v string) *DataValue
	GetTimestamp() *string
	SetValue(v interface{}) *DataValue
	GetValue() interface{}
}

type DataValue struct {
	// The timestamp. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1681975870000
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 478.28
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataValue) String() string {
	return dara.Prettify(s)
}

func (s DataValue) GoString() string {
	return s.String()
}

func (s *DataValue) GetTimestamp() *string {
	return s.Timestamp
}

func (s *DataValue) GetValue() interface{} {
	return s.Value
}

func (s *DataValue) SetTimestamp(v string) *DataValue {
	s.Timestamp = &v
	return s
}

func (s *DataValue) SetValue(v interface{}) *DataValue {
	s.Value = v
	return s
}

func (s *DataValue) Validate() error {
	return dara.Validate(s)
}
