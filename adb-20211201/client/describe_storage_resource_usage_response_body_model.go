// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeStorageResourceUsageResponseBody
	GetCode() *int32
	SetData(v *DescribeStorageResourceUsageResponseBodyData) *DescribeStorageResourceUsageResponseBody
	GetData() *DescribeStorageResourceUsageResponseBodyData
	SetRequestId(v string) *DescribeStorageResourceUsageResponseBody
	GetRequestId() *string
}

type DescribeStorageResourceUsageResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried resource usage.
	Data *DescribeStorageResourceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEAW
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeStorageResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeStorageResourceUsageResponseBody) GetData() *DescribeStorageResourceUsageResponseBodyData {
	return s.Data
}

func (s *DescribeStorageResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStorageResourceUsageResponseBody) SetCode(v int32) *DescribeStorageResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBody) SetData(v *DescribeStorageResourceUsageResponseBodyData) *DescribeStorageResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeStorageResourceUsageResponseBody) SetRequestId(v string) *DescribeStorageResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeStorageResourceUsageResponseBodyData struct {
	// The AnalyticDB compute unit (ACU) usage of the cluster.
	AcuInfo []*DescribeStorageResourceUsageResponseBodyDataAcuInfo `json:"AcuInfo,omitempty" xml:"AcuInfo,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// amv-bp1bg858bo8c****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-23T01:06:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-08-22T01:06:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeStorageResourceUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageResponseBodyData) GetAcuInfo() []*DescribeStorageResourceUsageResponseBodyDataAcuInfo {
	return s.AcuInfo
}

func (s *DescribeStorageResourceUsageResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeStorageResourceUsageResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeStorageResourceUsageResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeStorageResourceUsageResponseBodyData) SetAcuInfo(v []*DescribeStorageResourceUsageResponseBodyDataAcuInfo) *DescribeStorageResourceUsageResponseBodyData {
	s.AcuInfo = v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyData) SetDBClusterId(v string) *DescribeStorageResourceUsageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyData) SetEndTime(v string) *DescribeStorageResourceUsageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyData) SetStartTime(v string) *DescribeStorageResourceUsageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeStorageResourceUsageResponseBodyDataAcuInfo struct {
	// The resource usage metric. Valid values:
	//
	// 	- `TotalAcuNumber`: the total number of ACUs.
	//
	// 	- `ReservedAcuNumber`: the number of ACUs for the reserved resources.
	//
	// example:
	//
	// TotalAcuNumber
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The values of the metric at specific points in time.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeStorageResourceUsageResponseBodyDataAcuInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageResourceUsageResponseBodyDataAcuInfo) GoString() string {
	return s.String()
}

func (s *DescribeStorageResourceUsageResponseBodyDataAcuInfo) GetName() *string {
	return s.Name
}

func (s *DescribeStorageResourceUsageResponseBodyDataAcuInfo) GetValues() []*string {
	return s.Values
}

func (s *DescribeStorageResourceUsageResponseBodyDataAcuInfo) SetName(v string) *DescribeStorageResourceUsageResponseBodyDataAcuInfo {
	s.Name = &v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyDataAcuInfo) SetValues(v []*string) *DescribeStorageResourceUsageResponseBodyDataAcuInfo {
	s.Values = v
	return s
}

func (s *DescribeStorageResourceUsageResponseBodyDataAcuInfo) Validate() error {
	return dara.Validate(s)
}
