// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeDrdsInstanceMonitorResponseBodyData) *DescribeDrdsInstanceMonitorResponseBody
	GetData() []*DescribeDrdsInstanceMonitorResponseBodyData
	SetRequestId(v string) *DescribeDrdsInstanceMonitorResponseBody
	GetRequestId() *string
}

type DescribeDrdsInstanceMonitorResponseBody struct {
	// The result set of the query.
	Data []*DescribeDrdsInstanceMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2F7F8080-9132-4279-85D0-B7E5C4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBody) GetData() []*DescribeDrdsInstanceMonitorResponseBodyData {
	return s.Data
}

func (s *DescribeDrdsInstanceMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetData(v []*DescribeDrdsInstanceMonitorResponseBodyData) *DescribeDrdsInstanceMonitorResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBody) SetRequestId(v string) *DescribeDrdsInstanceMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceMonitorResponseBodyData struct {
	// The name of the metric.
	//
	// example:
	//
	// cpu
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The unit of the metric value.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The details of the monitoring data of the metric.
	Values []*DescribeDrdsInstanceMonitorResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) GetValues() []*DescribeDrdsInstanceMonitorResponseBodyDataValues {
	return s.Values
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetKey(v string) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetNodeNum(v int32) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.NodeNum = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetUnit(v string) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) SetValues(v []*DescribeDrdsInstanceMonitorResponseBodyDataValues) *DescribeDrdsInstanceMonitorResponseBodyData {
	s.Values = v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceMonitorResponseBodyDataValues struct {
	// The point in time when the value of the metric was collected. The value is in the UNIX timestamp format. The timestamp is displayed in UTC. Unit: ms.
	//
	// example:
	//
	// 1603163400000
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The value of the metric.
	//
	// example:
	//
	// 1.40
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataValues) GetDate() *int64 {
	return s.Date
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataValues) GetValue() *string {
	return s.Value
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataValues) SetDate(v int64) *DescribeDrdsInstanceMonitorResponseBodyDataValues {
	s.Date = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataValues) SetValue(v string) *DescribeDrdsInstanceMonitorResponseBodyDataValues {
	s.Value = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorResponseBodyDataValues) Validate() error {
	return dara.Validate(s)
}
