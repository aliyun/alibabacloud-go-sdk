// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSubscriptionPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeSubscriptionPerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) *DescribeSubscriptionPerformanceResponseBody
	GetPerformanceKeys() *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys
	SetReplicaId(v string) *DescribeSubscriptionPerformanceResponseBody
	GetReplicaId() *string
	SetRequestId(v string) *DescribeSubscriptionPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeSubscriptionPerformanceResponseBody
	GetStartTime() *string
}

type DescribeSubscriptionPerformanceResponseBody struct {
	EndTime         *string                                                     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PerformanceKeys *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	ReplicaId       *string                                                     `json:"ReplicaId,omitempty" xml:"ReplicaId,omitempty"`
	RequestId       *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StartTime       *string                                                     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSubscriptionPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSubscriptionPerformanceResponseBody) GetPerformanceKeys() *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeSubscriptionPerformanceResponseBody) GetReplicaId() *string {
	return s.ReplicaId
}

func (s *DescribeSubscriptionPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSubscriptionPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetEndTime(v string) *DescribeSubscriptionPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetPerformanceKeys(v *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) *DescribeSubscriptionPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetReplicaId(v string) *DescribeSubscriptionPerformanceResponseBody {
	s.ReplicaId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetRequestId(v string) *DescribeSubscriptionPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBody) SetStartTime(v string) *DescribeSubscriptionPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		if err := s.PerformanceKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSubscriptionPerformanceResponseBodyPerformanceKeys struct {
	PerformanceKey []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey `json:"PerformanceKey,omitempty" xml:"PerformanceKey,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) GetPerformanceKey() []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey {
	return s.PerformanceKey
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) SetPerformanceKey(v []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys {
	s.PerformanceKey = v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeys) Validate() error {
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

type DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey struct {
	Key               *string                                                                                    `json:"Key,omitempty" xml:"Key,omitempty"`
	PerformanceValues *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues `json:"PerformanceValues,omitempty" xml:"PerformanceValues,omitempty" type:"Struct"`
	Unit              *string                                                                                    `json:"Unit,omitempty" xml:"Unit,omitempty"`
	ValueFormat       *string                                                                                    `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) GetKey() *string {
	return s.Key
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) GetPerformanceValues() *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues {
	return s.PerformanceValues
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) GetUnit() *string {
	return s.Unit
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) SetKey(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Key = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) SetPerformanceValues(v *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.PerformanceValues = v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) SetUnit(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Unit = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) SetValueFormat(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.ValueFormat = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKey) Validate() error {
	if s.PerformanceValues != nil {
		if err := s.PerformanceValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues struct {
	PerformanceValue []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue `json:"PerformanceValue,omitempty" xml:"PerformanceValue,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) GetPerformanceValue() []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	return s.PerformanceValue
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) SetPerformanceValue(v []*DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues {
	s.PerformanceValue = v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValues) Validate() error {
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

type DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue struct {
	Date  *string `json:"Date,omitempty" xml:"Date,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) GetDate() *string {
	return s.Date
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) GetValue() *string {
	return s.Value
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) SetDate(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	s.Date = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) SetValue(v string) *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue {
	s.Value = &v
	return s
}

func (s *DescribeSubscriptionPerformanceResponseBodyPerformanceKeysPerformanceKeyPerformanceValuesPerformanceValue) Validate() error {
	return dara.Validate(s)
}
