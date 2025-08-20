// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeClusterResourceUsageResponseBody
	GetCode() *int32
	SetData(v *DescribeClusterResourceUsageResponseBodyData) *DescribeClusterResourceUsageResponseBody
	GetData() *DescribeClusterResourceUsageResponseBodyData
	SetRequestId(v string) *DescribeClusterResourceUsageResponseBody
	GetRequestId() *string
}

type DescribeClusterResourceUsageResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried resource usage.
	Data *DescribeClusterResourceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEAW
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClusterResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeClusterResourceUsageResponseBody) GetData() *DescribeClusterResourceUsageResponseBodyData {
	return s.Data
}

func (s *DescribeClusterResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterResourceUsageResponseBody) SetCode(v int32) *DescribeClusterResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBody) SetData(v *DescribeClusterResourceUsageResponseBodyData) *DescribeClusterResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeClusterResourceUsageResponseBody) SetRequestId(v string) *DescribeClusterResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterResourceUsageResponseBodyData struct {
	// The AnalyticDB compute unit (ACU) usage of the cluster.
	AcuInfo []*DescribeClusterResourceUsageResponseBodyDataAcuInfo `json:"AcuInfo,omitempty" xml:"AcuInfo,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// amv-uf6dj23rt5zo9s9d
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-23T02:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-03-14T03:42:15Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClusterResourceUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageResponseBodyData) GetAcuInfo() []*DescribeClusterResourceUsageResponseBodyDataAcuInfo {
	return s.AcuInfo
}

func (s *DescribeClusterResourceUsageResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeClusterResourceUsageResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeClusterResourceUsageResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeClusterResourceUsageResponseBodyData) SetAcuInfo(v []*DescribeClusterResourceUsageResponseBodyDataAcuInfo) *DescribeClusterResourceUsageResponseBodyData {
	s.AcuInfo = v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyData) SetDBClusterId(v string) *DescribeClusterResourceUsageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyData) SetEndTime(v string) *DescribeClusterResourceUsageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyData) SetStartTime(v string) *DescribeClusterResourceUsageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterResourceUsageResponseBodyDataAcuInfo struct {
	// The resource usage metric. Valid values:
	//
	// 	- `TotalAcuNumber`: the total number of ACUs.
	//
	// 	- `ReservedAcuNumber`: the number of ACUs for the reserved resources.
	//
	// 	- `ReservedAcuUsageNumber`: the number of ACUs for the reserved resources that are used.
	//
	// example:
	//
	// TotalAcuNumber
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the metric at specific points in time.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeClusterResourceUsageResponseBodyDataAcuInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResourceUsageResponseBodyDataAcuInfo) GoString() string {
	return s.String()
}

func (s *DescribeClusterResourceUsageResponseBodyDataAcuInfo) GetName() *string {
	return s.Name
}

func (s *DescribeClusterResourceUsageResponseBodyDataAcuInfo) GetValues() []*string {
	return s.Values
}

func (s *DescribeClusterResourceUsageResponseBodyDataAcuInfo) SetName(v string) *DescribeClusterResourceUsageResponseBodyDataAcuInfo {
	s.Name = &v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyDataAcuInfo) SetValues(v []*string) *DescribeClusterResourceUsageResponseBodyDataAcuInfo {
	s.Values = v
	return s
}

func (s *DescribeClusterResourceUsageResponseBodyDataAcuInfo) Validate() error {
	return dara.Validate(s)
}
