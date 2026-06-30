// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOperatorUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeModelOperatorUsageResponseBody
	GetEndTime() *string
	SetKeys(v []*DescribeModelOperatorUsageResponseBodyKeys) *DescribeModelOperatorUsageResponseBody
	GetKeys() []*DescribeModelOperatorUsageResponseBodyKeys
	SetPeriod(v int32) *DescribeModelOperatorUsageResponseBody
	GetPeriod() *int32
	SetRequestId(v string) *DescribeModelOperatorUsageResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeModelOperatorUsageResponseBody
	GetStartTime() *string
}

type DescribeModelOperatorUsageResponseBody struct {
	// example:
	//
	// 2026-06-02T00:00Z
	EndTime *string                                       `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Keys    []*DescribeModelOperatorUsageResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2026-06-01T00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeModelOperatorUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorUsageResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeModelOperatorUsageResponseBody) GetKeys() []*DescribeModelOperatorUsageResponseBodyKeys {
	return s.Keys
}

func (s *DescribeModelOperatorUsageResponseBody) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeModelOperatorUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelOperatorUsageResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeModelOperatorUsageResponseBody) SetEndTime(v string) *DescribeModelOperatorUsageResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeModelOperatorUsageResponseBody) SetKeys(v []*DescribeModelOperatorUsageResponseBodyKeys) *DescribeModelOperatorUsageResponseBody {
	s.Keys = v
	return s
}

func (s *DescribeModelOperatorUsageResponseBody) SetPeriod(v int32) *DescribeModelOperatorUsageResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeModelOperatorUsageResponseBody) SetRequestId(v string) *DescribeModelOperatorUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelOperatorUsageResponseBody) SetStartTime(v string) *DescribeModelOperatorUsageResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeModelOperatorUsageResponseBody) Validate() error {
	if s.Keys != nil {
		for _, item := range s.Keys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelOperatorUsageResponseBodyKeys struct {
	// example:
	//
	// request_count
	Name   *string                                             `json:"Name,omitempty" xml:"Name,omitempty"`
	Series []*DescribeModelOperatorUsageResponseBodyKeysSeries `json:"Series,omitempty" xml:"Series,omitempty" type:"Repeated"`
	// example:
	//
	// requests
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeModelOperatorUsageResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorUsageResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorUsageResponseBodyKeys) GetName() *string {
	return s.Name
}

func (s *DescribeModelOperatorUsageResponseBodyKeys) GetSeries() []*DescribeModelOperatorUsageResponseBodyKeysSeries {
	return s.Series
}

func (s *DescribeModelOperatorUsageResponseBodyKeys) GetUnit() *string {
	return s.Unit
}

func (s *DescribeModelOperatorUsageResponseBodyKeys) SetName(v string) *DescribeModelOperatorUsageResponseBodyKeys {
	s.Name = &v
	return s
}

func (s *DescribeModelOperatorUsageResponseBodyKeys) SetSeries(v []*DescribeModelOperatorUsageResponseBodyKeysSeries) *DescribeModelOperatorUsageResponseBodyKeys {
	s.Series = v
	return s
}

func (s *DescribeModelOperatorUsageResponseBodyKeys) SetUnit(v string) *DescribeModelOperatorUsageResponseBodyKeys {
	s.Unit = &v
	return s
}

func (s *DescribeModelOperatorUsageResponseBodyKeys) Validate() error {
	if s.Series != nil {
		for _, item := range s.Series {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelOperatorUsageResponseBodyKeysSeries struct {
	// example:
	//
	// 1
	ApiKeyId *int32 `json:"ApiKeyId,omitempty" xml:"ApiKeyId,omitempty"`
	// example:
	//
	// qwen3.6-plus
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// model
	Role   *string                                                   `json:"Role,omitempty" xml:"Role,omitempty"`
	Values []*DescribeModelOperatorUsageResponseBodyKeysSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeModelOperatorUsageResponseBodyKeysSeries) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorUsageResponseBodyKeysSeries) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeries) GetApiKeyId() *int32 {
	return s.ApiKeyId
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeries) GetName() *string {
	return s.Name
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeries) GetRole() *string {
	return s.Role
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeries) GetValues() []*DescribeModelOperatorUsageResponseBodyKeysSeriesValues {
	return s.Values
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeries) SetApiKeyId(v int32) *DescribeModelOperatorUsageResponseBodyKeysSeries {
	s.ApiKeyId = &v
	return s
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeries) SetName(v string) *DescribeModelOperatorUsageResponseBodyKeysSeries {
	s.Name = &v
	return s
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeries) SetRole(v string) *DescribeModelOperatorUsageResponseBodyKeysSeries {
	s.Role = &v
	return s
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeries) SetValues(v []*DescribeModelOperatorUsageResponseBodyKeysSeriesValues) *DescribeModelOperatorUsageResponseBodyKeysSeries {
	s.Values = v
	return s
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeries) Validate() error {
	if s.Values != nil {
		for _, item := range s.Values {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeModelOperatorUsageResponseBodyKeysSeriesValues struct {
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" type:"Repeated"`
}

func (s DescribeModelOperatorUsageResponseBodyKeysSeriesValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOperatorUsageResponseBodyKeysSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeriesValues) GetPoint() []*string {
	return s.Point
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeriesValues) SetPoint(v []*string) *DescribeModelOperatorUsageResponseBodyKeysSeriesValues {
	s.Point = v
	return s
}

func (s *DescribeModelOperatorUsageResponseBodyKeysSeriesValues) Validate() error {
	return dara.Validate(s)
}
