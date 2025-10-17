// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComputeResourceUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeComputeResourceUsageResponseBody
	GetCode() *int32
	SetData(v *DescribeComputeResourceUsageResponseBodyData) *DescribeComputeResourceUsageResponseBody
	GetData() *DescribeComputeResourceUsageResponseBodyData
	SetRequestId(v string) *DescribeComputeResourceUsageResponseBody
	GetRequestId() *string
}

type DescribeComputeResourceUsageResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried resource usage.
	Data *DescribeComputeResourceUsageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEAW
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComputeResourceUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeResourceUsageResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeComputeResourceUsageResponseBody) GetData() *DescribeComputeResourceUsageResponseBodyData {
	return s.Data
}

func (s *DescribeComputeResourceUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComputeResourceUsageResponseBody) SetCode(v int32) *DescribeComputeResourceUsageResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBody) SetData(v *DescribeComputeResourceUsageResponseBodyData) *DescribeComputeResourceUsageResponseBody {
	s.Data = v
	return s
}

func (s *DescribeComputeResourceUsageResponseBody) SetRequestId(v string) *DescribeComputeResourceUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeComputeResourceUsageResponseBodyData struct {
	// The AnalyticDB compute unit (ACU) usage of the cluster.
	AcuInfo []*DescribeComputeResourceUsageResponseBodyDataAcuInfo `json:"AcuInfo,omitempty" xml:"AcuInfo,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// amv-clusterxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The end time of the query. The time follows the ISO 8601 standard in the *yyyy-MM-ddTHH:mm:ssZ	- format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-06-07T02:37:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
	// The type of the resource group.
	//
	// example:
	//
	// interative
	ResourceGroupType *string `json:"ResourceGroupType,omitempty" xml:"ResourceGroupType,omitempty"`
	// The start time of the query. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-04-24T07:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeComputeResourceUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeResourceUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageResponseBodyData) GetAcuInfo() []*DescribeComputeResourceUsageResponseBodyDataAcuInfo {
	return s.AcuInfo
}

func (s *DescribeComputeResourceUsageResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeComputeResourceUsageResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeComputeResourceUsageResponseBodyData) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *DescribeComputeResourceUsageResponseBodyData) GetResourceGroupType() *string {
	return s.ResourceGroupType
}

func (s *DescribeComputeResourceUsageResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetAcuInfo(v []*DescribeComputeResourceUsageResponseBodyDataAcuInfo) *DescribeComputeResourceUsageResponseBodyData {
	s.AcuInfo = v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetDBClusterId(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetEndTime(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetResourceGroupName(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.ResourceGroupName = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetResourceGroupType(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.ResourceGroupType = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) SetStartTime(v string) *DescribeComputeResourceUsageResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyData) Validate() error {
	if s.AcuInfo != nil {
		for _, item := range s.AcuInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeComputeResourceUsageResponseBodyDataAcuInfo struct {
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

func (s DescribeComputeResourceUsageResponseBodyDataAcuInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeResourceUsageResponseBodyDataAcuInfo) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceUsageResponseBodyDataAcuInfo) GetName() *string {
	return s.Name
}

func (s *DescribeComputeResourceUsageResponseBodyDataAcuInfo) GetValues() []*string {
	return s.Values
}

func (s *DescribeComputeResourceUsageResponseBodyDataAcuInfo) SetName(v string) *DescribeComputeResourceUsageResponseBodyDataAcuInfo {
	s.Name = &v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyDataAcuInfo) SetValues(v []*string) *DescribeComputeResourceUsageResponseBodyDataAcuInfo {
	s.Values = v
	return s
}

func (s *DescribeComputeResourceUsageResponseBodyDataAcuInfo) Validate() error {
	return dara.Validate(s)
}
