// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBProxyPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBProxyPerformanceResponseBody
	GetDBInstanceId() *string
	SetDBProxyEngineType(v string) *DescribeDBProxyPerformanceResponseBody
	GetDBProxyEngineType() *string
	SetEndTime(v string) *DescribeDBProxyPerformanceResponseBody
	GetEndTime() *string
	SetPerformanceKeys(v *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) *DescribeDBProxyPerformanceResponseBody
	GetPerformanceKeys() *DescribeDBProxyPerformanceResponseBodyPerformanceKeys
	SetRequestId(v string) *DescribeDBProxyPerformanceResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDBProxyPerformanceResponseBody
	GetStartTime() *string
}

type DescribeDBProxyPerformanceResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// lsmexxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// An internal parameter. You do not need to specify this parameter.
	//
	// example:
	//
	// normal
	DBProxyEngineType *string `json:"DBProxyEngineType,omitempty" xml:"DBProxyEngineType,omitempty"`
	// The end time of the query.
	//
	// example:
	//
	// 2019-09-21T18:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance list.
	PerformanceKeys *DescribeDBProxyPerformanceResponseBodyPerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// DD31056F-A0CE-41D7-AD39-689B6ABAE982
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start time of the query.
	//
	// example:
	//
	// 2019-09-19T01:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDBProxyPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBProxyPerformanceResponseBody) GetDBProxyEngineType() *string {
	return s.DBProxyEngineType
}

func (s *DescribeDBProxyPerformanceResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDBProxyPerformanceResponseBody) GetPerformanceKeys() *DescribeDBProxyPerformanceResponseBodyPerformanceKeys {
	return s.PerformanceKeys
}

func (s *DescribeDBProxyPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBProxyPerformanceResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDBProxyPerformanceResponseBody) SetDBInstanceId(v string) *DescribeDBProxyPerformanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetDBProxyEngineType(v string) *DescribeDBProxyPerformanceResponseBody {
	s.DBProxyEngineType = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetEndTime(v string) *DescribeDBProxyPerformanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetPerformanceKeys(v *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) *DescribeDBProxyPerformanceResponseBody {
	s.PerformanceKeys = v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetRequestId(v string) *DescribeDBProxyPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) SetStartTime(v string) *DescribeDBProxyPerformanceResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBody) Validate() error {
	if s.PerformanceKeys != nil {
		if err := s.PerformanceKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBProxyPerformanceResponseBodyPerformanceKeys struct {
	PerformanceKey []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey `json:"PerformanceKey,omitempty" xml:"PerformanceKey,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) GetPerformanceKey() []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey {
	return s.PerformanceKey
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) SetPerformanceKey(v []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) *DescribeDBProxyPerformanceResponseBodyPerformanceKeys {
	s.PerformanceKey = v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeys) Validate() error {
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

type DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey struct {
	// The performance parameter.
	//
	// example:
	//
	// cpu_ratio
	Key    *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Node   *string `json:"Node,omitempty" xml:"Node,omitempty"`
	Server *string `json:"Server,omitempty" xml:"Server,omitempty"`
	// The service dimension.
	//
	// example:
	//
	// reserve_3
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The format in which the value of the performance metric is returned.
	//
	// example:
	//
	// docker_container_cpu
	ValueFormat *string `json:"ValueFormat,omitempty" xml:"ValueFormat,omitempty"`
	// The performance metrics.
	Values *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) GetKey() *string {
	return s.Key
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) GetNode() *string {
	return s.Node
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) GetServer() *string {
	return s.Server
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) GetService() *string {
	return s.Service
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) GetValueFormat() *string {
	return s.ValueFormat
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) GetValues() *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues {
	return s.Values
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) SetKey(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Key = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) SetNode(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Node = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) SetServer(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Server = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) SetService(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Service = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) SetValueFormat(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.ValueFormat = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) SetValues(v *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey {
	s.Values = v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKey) Validate() error {
	if s.Values != nil {
		if err := s.Values.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues struct {
	PerformanceValue []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue `json:"PerformanceValue,omitempty" xml:"PerformanceValue,omitempty" type:"Repeated"`
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues) GetPerformanceValue() []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue {
	return s.PerformanceValue
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues) SetPerformanceValue(v []*DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues {
	s.PerformanceValue = v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValues) Validate() error {
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

type DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue struct {
	// The date and time when the value of the performance metric was recorded. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-10T09:00:00Z
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The value of the performance metric.
	//
	// example:
	//
	// 2.83
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) GoString() string {
	return s.String()
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) GetDate() *string {
	return s.Date
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) GetValue() *string {
	return s.Value
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) SetDate(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue {
	s.Date = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) SetValue(v string) *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue {
	s.Value = &v
	return s
}

func (s *DescribeDBProxyPerformanceResponseBodyPerformanceKeysPerformanceKeyValuesPerformanceValue) Validate() error {
	return dara.Validate(s)
}
