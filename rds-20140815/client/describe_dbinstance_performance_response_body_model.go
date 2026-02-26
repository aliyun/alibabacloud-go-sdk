// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponseBody
	GetDBInstanceId() *string
	SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody
	GetEndTime() *string
	SetEngine(v string) *DescribeDBInstancePerformanceResponseBody
	GetEngine() *string
	SetPerformanceKeys(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) *DescribeDBInstancePerformanceResponseBody
	GetPerformanceKeys() *DescribeDBInstancePerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBInstancePerformanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The end time of the query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2012-06-19T15:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The database engine of the instance.
	//
	// example:
	//
	// MySQL
	Engine          *string                                                   `json:"Engine,omitempty" xml:"Engine,omitempty"`
	PerformanceKeys *DescribeDBInstancePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A5409D02-D661-4BF3-8F3D-0A814D0574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time must be in UTC.
	//
	// example:
	//
	// 2012-06-10T15:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstancePerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBInstancePerformanceResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBInstancePerformanceResponseBody) GetPerformanceKeys() *DescribeDBInstancePerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeDBInstancePerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstancePerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBInstancePerformanceResponseBody) SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetEngine(v string) *DescribeDBInstancePerformanceResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetPerformanceKeys(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) *DescribeDBInstancePerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		if err := s.PerformanceKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeys struct {
	PerformanceKey []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey `json:"PerformanceKey,omitempty" xml:"PerformanceKey,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) GetPerformanceKey() []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	return s.PerformanceKey
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) SetPerformanceKey(v []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) *DescribeDBInstancePerformanceResponseBodyPerformanceKeys {
	s.PerformanceKey = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) Validate() error {
	if s.PerformanceKey != nil {
		for _, item := range s.PerformanceKey {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey struct {
	Key         *string                                                                       `json:"Key,omitempty" xml:"Key,omitempty"`
	Unit        *string                                                                       `json:"Unit,omitempty" xml:"Unit,omitempty"`
	ValueFormat *string                                                                       `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
	Values      *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) GetKey() *string {
	return s.Key
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) GetValues() *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues {
	return s.Values
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetKey(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetUnit(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Unit = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetValueFormat(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.ValueFormat = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetValues(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Values = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) Validate() error {
	if s.Values != nil {
		if err := s.Values.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues struct {
	PerformanceValue []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue `json:"PerformanceValue,omitempty" xml:"PerformanceValue,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues) GetPerformanceValue() []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue {
	return s.PerformanceValue
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues) SetPerformanceValue(v []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues {
	s.PerformanceValue = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValues) Validate() error {
	if s.PerformanceValue != nil {
		for _, item := range s.PerformanceValue {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue struct {
	Date  *string `json:"Date,omitempty" xml:"Date,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) GetDate() *string {
	return s.Date
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) SetDate(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue {
	s.Date = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) SetValue(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) Validate() error {
	return dara.Validate(s)
}
