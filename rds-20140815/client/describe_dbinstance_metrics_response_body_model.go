// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceMetricsResponseBody
	GetDBInstanceName() *string
	SetItems(v []*DescribeDBInstanceMetricsResponseBodyItems) *DescribeDBInstanceMetricsResponseBody
	GetItems() []*DescribeDBInstanceMetricsResponseBodyItems
	SetRequestId(v string) *DescribeDBInstanceMetricsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeDBInstanceMetricsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeDBInstanceMetricsResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-bp1*****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// An array consisting of the Enhanced Monitoring metrics that are enabled for the instance.
	Items []*DescribeDBInstanceMetricsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 318C3754-F6D0-54BB-A55C-23EAA04708B7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of enhanced monitoring metrics that are enabled for the instance.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeDBInstanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMetricsResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceMetricsResponseBody) GetItems() []*DescribeDBInstanceMetricsResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceMetricsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeDBInstanceMetricsResponseBody) SetDBInstanceName(v string) *DescribeDBInstanceMetricsResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBody) SetItems(v []*DescribeDBInstanceMetricsResponseBodyItems) *DescribeDBInstanceMetricsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBody) SetRequestId(v string) *DescribeDBInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBody) SetTotalRecordCount(v int32) *DescribeDBInstanceMetricsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceMetricsResponseBodyItems struct {
	// The description of the enhanced monitoring metric.
	//
	// example:
	//
	// OS CPU utilization, equal to the number of OS-consumed CPUs divided by the total number of CPUs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The category of the enhanced monitoring metric. Valid values:
	//
	// 	- **os**: OS metric
	//
	// 	- **db**: database metric
	//
	// example:
	//
	// os
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The key of the group to which the enhanced monitoring metric belongs.
	//
	// example:
	//
	// os.cpu_usage
	GroupKey *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	// The name of the group to which the enhanced monitoring metric belongs.
	//
	// example:
	//
	// CPU Utilization Rate
	GroupKeyType *string `json:"GroupKeyType,omitempty" xml:"GroupKeyType,omitempty"`
	// The method that is used to aggregate the monitoring data of the enhanced monitoring metric. Valid values:
	//
	// 	- **avg**: The system calculates the average value of the enhanced monitoring metric.
	//
	// 	- **min**: The system calculates the minimum value of the enhanced monitoring metric.
	//
	// 	- **max**: The system calculates the maximum value of the enhanced monitoring metric.
	//
	// example:
	//
	// avg
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The key of the enhanced monitoring metric.
	//
	// example:
	//
	// os.cpu_usage.sys.avg
	MetricsKey *string `json:"MetricsKey,omitempty" xml:"MetricsKey,omitempty"`
	// The alias of the enhanced monitoring metric.
	//
	// example:
	//
	// os.cpu_usage.sys
	MetricsKeyAlias *string `json:"MetricsKeyAlias,omitempty" xml:"MetricsKeyAlias,omitempty"`
	// The serial number of the enhanced monitoring metric.
	//
	// example:
	//
	// 1
	SortRule *int32 `json:"SortRule,omitempty" xml:"SortRule,omitempty"`
	// The unit of the enhanced monitoring metric.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDBInstanceMetricsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceMetricsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetGroupKey() *string {
	return s.GroupKey
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetGroupKeyType() *string {
	return s.GroupKeyType
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetMethod() *string {
	return s.Method
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetMetricsKey() *string {
	return s.MetricsKey
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetMetricsKeyAlias() *string {
	return s.MetricsKeyAlias
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetSortRule() *int32 {
	return s.SortRule
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetDescription(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetDimension(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.Dimension = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetGroupKey(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.GroupKey = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetGroupKeyType(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.GroupKeyType = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetMethod(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.Method = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetMetricsKey(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.MetricsKey = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetMetricsKeyAlias(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.MetricsKeyAlias = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetSortRule(v int32) *DescribeDBInstanceMetricsResponseBodyItems {
	s.SortRule = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) SetUnit(v string) *DescribeDBInstanceMetricsResponseBodyItems {
	s.Unit = &v
	return s
}

func (s *DescribeDBInstanceMetricsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
