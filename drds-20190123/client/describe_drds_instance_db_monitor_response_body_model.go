// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceDbMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeDrdsInstanceDbMonitorResponseBodyData) *DescribeDrdsInstanceDbMonitorResponseBody
	GetData() []*DescribeDrdsInstanceDbMonitorResponseBodyData
	SetRequestId(v string) *DescribeDrdsInstanceDbMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsInstanceDbMonitorResponseBody
	GetSuccess() *bool
}

type DescribeDrdsInstanceDbMonitorResponseBody struct {
	// The list of monitoring data.
	Data []*DescribeDrdsInstanceDbMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2F7F8080-9132-4279-85D0-B7E5C4******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) GetData() []*DescribeDrdsInstanceDbMonitorResponseBodyData {
	return s.Data
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetData(v []*DescribeDrdsInstanceDbMonitorResponseBodyData) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetRequestId(v string) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceDbMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceDbMonitorResponseBodyData struct {
	// The name of the monitoring metric.
	//
	// example:
	//
	// qps
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The unit of the monitoring metric.
	//
	// example:
	//
	// qps
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The details about the value of monitoring data.
	Values []*DescribeDrdsInstanceDbMonitorResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) GetValues() []*DescribeDrdsInstanceDbMonitorResponseBodyDataValues {
	return s.Values
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetKey(v string) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetUnit(v string) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) SetValues(v []*DescribeDrdsInstanceDbMonitorResponseBodyDataValues) *DescribeDrdsInstanceDbMonitorResponseBodyData {
	s.Values = v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceDbMonitorResponseBodyDataValues struct {
	// The time point when the value of monitoring data was obtained. The value is in the UNIX timestamp format. Unit: ms.
	//
	// example:
	//
	// 1603162805000
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The data value.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceDbMonitorResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataValues) GetDate() *int64 {
	return s.Date
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataValues) GetValue() *string {
	return s.Value
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataValues) SetDate(v int64) *DescribeDrdsInstanceDbMonitorResponseBodyDataValues {
	s.Date = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataValues) SetValue(v string) *DescribeDrdsInstanceDbMonitorResponseBodyDataValues {
	s.Value = &v
	return s
}

func (s *DescribeDrdsInstanceDbMonitorResponseBodyDataValues) Validate() error {
	return dara.Validate(s)
}
