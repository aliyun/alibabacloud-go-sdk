// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeys) *DescribeDBInstancePerformanceResponseBody
	GetPerformanceKeys() *DescribeDBInstancePerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeDBInstancePerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBInstancePerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBInstancePerformanceResponseBody struct {
	// The end of the queried time range. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-06-13T11:58Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The details of performance metrics.
	PerformanceKeys *DescribeDBInstancePerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4A300BC7-6D8F-527F-A2DB-A7768D26E9AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the queried time range. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-06-13T10:58Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
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

func (s *DescribeDBInstancePerformanceResponseBody) SetEndTime(v string) *DescribeDBInstancePerformanceResponseBody {
	s.EndTime = &v
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
	// The performance metrics that are returned.
	//
	// example:
	//
	// CpuUsage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The details of the performance metric values.
	PerformanceValues *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues `json:"PerformanceValues,omitempty" xml:"PerformanceValues,omitempty" type:"Struct"`
	// The unit of the performance metric.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The format of the performance metric value. If the performance metric contains multiple fields, the fields are separated with ampersands ( &).
	//
	// For example, if you query disk space usage, the returned value of the **ValueFormat*	- parameter is **ins_size\\&data_size\\&log_size**.
	//
	// example:
	//
	// cpu_usage
	ValueFormat *string `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
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

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) GetPerformanceValues() *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues {
	return s.PerformanceValues
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetKey(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) SetPerformanceValues(v *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.PerformanceValues = v
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

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKey) Validate() error {
	if s.PerformanceValues != nil {
		if err := s.PerformanceValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues struct {
	PerformanceValue []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue `json:"PerformanceValue,omitempty" xml:"PerformanceValue,omitempty" type:"Repeated"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) GetPerformanceValue() []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	return s.PerformanceValue
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) SetPerformanceValue(v []*DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues {
	s.PerformanceValue = v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) Validate() error {
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

type DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue struct {
	// The date and time when the metric value was generated.
	//
	// example:
	//
	// 2022-06-13T10:58:00Z
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The value of the performance metric.
	//
	// example:
	//
	// 0.23
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) GetDate() *string {
	return s.Date
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) GetValue() *string {
	return s.Value
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) SetDate(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	s.Date = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) SetValue(v string) *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	s.Value = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) Validate() error {
	return dara.Validate(s)
}
