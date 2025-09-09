// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRDSPerformanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeRDSPerformanceResponseBodyData) *DescribeRDSPerformanceResponseBody
	GetData() []*DescribeRDSPerformanceResponseBodyData
	SetRequestId(v string) *DescribeRDSPerformanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRDSPerformanceResponseBody
	GetSuccess() *bool
}

type DescribeRDSPerformanceResponseBody struct {
	// The result set of the query.
	Data []*DescribeRDSPerformanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// DB53EC68-463C-4187-8D2B-C2AD8C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRDSPerformanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBody) GetData() []*DescribeRDSPerformanceResponseBodyData {
	return s.Data
}

func (s *DescribeRDSPerformanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRDSPerformanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRDSPerformanceResponseBody) SetData(v []*DescribeRDSPerformanceResponseBodyData) *DescribeRDSPerformanceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRDSPerformanceResponseBody) SetRequestId(v string) *DescribeRDSPerformanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBody) SetSuccess(v bool) *DescribeRDSPerformanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRDSPerformanceResponseBodyData struct {
	// The name of the monitoring metric.
	//
	// example:
	//
	// cpuusage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the node.
	//
	// >  This parameter is returned only when the storage type of the database is PolarDB for MySQL. If the storage type of the database is ApsaraDB RDS for MySQL, this parameter is not returned.
	//
	// example:
	//
	// pi-*************
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 1
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The unit of the monitoring metric.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The details of the monitoring metric data.
	Values []*DescribeRDSPerformanceResponseBodyDataValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeRDSPerformanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *DescribeRDSPerformanceResponseBodyData) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeRDSPerformanceResponseBodyData) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *DescribeRDSPerformanceResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *DescribeRDSPerformanceResponseBodyData) GetValues() []*DescribeRDSPerformanceResponseBodyDataValues {
	return s.Values
}

func (s *DescribeRDSPerformanceResponseBodyData) SetKey(v string) *DescribeRDSPerformanceResponseBodyData {
	s.Key = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetNodeName(v string) *DescribeRDSPerformanceResponseBodyData {
	s.NodeName = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetNodeNum(v int32) *DescribeRDSPerformanceResponseBodyData {
	s.NodeNum = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetUnit(v string) *DescribeRDSPerformanceResponseBodyData {
	s.Unit = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) SetValues(v []*DescribeRDSPerformanceResponseBodyDataValues) *DescribeRDSPerformanceResponseBodyData {
	s.Values = v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeRDSPerformanceResponseBodyDataValues struct {
	// The time point when the value of the monitoring metric was obtained. The value is in the UNIX timestamp format. The time is displayed in UTC. Unit: ms.
	//
	// example:
	//
	// 1603209660000
	Date *int64 `json:"Date,omitempty" xml:"Date,omitempty"`
	// The value of the monitoring metric.
	//
	// example:
	//
	// 0.58
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRDSPerformanceResponseBodyDataValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeRDSPerformanceResponseBodyDataValues) GoString() string {
	return s.String()
}

func (s *DescribeRDSPerformanceResponseBodyDataValues) GetDate() *int64 {
	return s.Date
}

func (s *DescribeRDSPerformanceResponseBodyDataValues) GetValue() *string {
	return s.Value
}

func (s *DescribeRDSPerformanceResponseBodyDataValues) SetDate(v int64) *DescribeRDSPerformanceResponseBodyDataValues {
	s.Date = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyDataValues) SetValue(v string) *DescribeRDSPerformanceResponseBodyDataValues {
	s.Value = &v
	return s
}

func (s *DescribeRDSPerformanceResponseBodyDataValues) Validate() error {
	return dara.Validate(s)
}
