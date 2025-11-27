// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeAvailableMetricsResponseBody
	GetDBInstanceName() *string
	SetItems(v []*DescribeAvailableMetricsResponseBodyItems) *DescribeAvailableMetricsResponseBody
	GetItems() []*DescribeAvailableMetricsResponseBodyItems
	SetRequestId(v string) *DescribeAvailableMetricsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeAvailableMetricsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeAvailableMetricsResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-bp1*****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// Details of the Enhanced Monitoring metric.
	Items []*DescribeAvailableMetricsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 5CD61041-35F7-10F7-BE94-33A48B221218
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of enhanced monitoring metrics that are available for the instance.
	//
	// example:
	//
	// 4
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeAvailableMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMetricsResponseBody) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAvailableMetricsResponseBody) GetItems() []*DescribeAvailableMetricsResponseBodyItems {
	return s.Items
}

func (s *DescribeAvailableMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableMetricsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeAvailableMetricsResponseBody) SetDBInstanceName(v string) *DescribeAvailableMetricsResponseBody {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBody) SetItems(v []*DescribeAvailableMetricsResponseBodyItems) *DescribeAvailableMetricsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAvailableMetricsResponseBody) SetRequestId(v string) *DescribeAvailableMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBody) SetTotalRecordCount(v int32) *DescribeAvailableMetricsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBody) Validate() error {
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

type DescribeAvailableMetricsResponseBodyItems struct {
	// The description of the Enhanced Monitoring metric.
	//
	// example:
	//
	// OS CPU utilization, equal to the number of OS-consumed CPUs divided by the total number of CPUs
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The category of the Enhanced Monitoring metric. Valid values:
	//
	// 	- **os**: OS metric
	//
	// 	- **db**: database metric
	//
	// example:
	//
	// os
	Dimension *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	// The key of the group to which the Enhanced Monitoring metric belongs.
	//
	// example:
	//
	// os.cpu_usage
	GroupKey *string `json:"GroupKey,omitempty" xml:"GroupKey,omitempty"`
	// The name of the group to which the Enhanced Monitoring metric belongs.
	//
	// example:
	//
	// CPU Utilization Rate
	GroupKeyType *string `json:"GroupKeyType,omitempty" xml:"GroupKeyType,omitempty"`
	// The method that is used to aggregate the monitoring data of the Enhanced Monitoring metric. Valid values:
	//
	// 	- **avg**: The system calculates the average value of the Enhanced Monitoring metric.
	//
	// 	- **min**: The system calculates the minimum value of the Enhanced Monitoring metric.
	//
	// 	- **max**: The system calculates the maximum value of the Enhanced Monitoring metric.
	//
	// example:
	//
	// avg
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The key of the Enhanced Monitoring metric.
	//
	// example:
	//
	// os.cpu_usage.sys.avg
	MetricsKey *string `json:"MetricsKey,omitempty" xml:"MetricsKey,omitempty"`
	// The alias of the Enhanced Monitoring metric.
	//
	// example:
	//
	// cpu_sys_per_core
	MetricsKeyAlias *string `json:"MetricsKeyAlias,omitempty" xml:"MetricsKeyAlias,omitempty"`
	// The serial number of the Enhanced Monitoring metric.
	//
	// example:
	//
	// 1
	SortRule *int32 `json:"SortRule,omitempty" xml:"SortRule,omitempty"`
	// The unit of the Enhanced Monitoring metric.
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeAvailableMetricsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableMetricsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMetricsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *DescribeAvailableMetricsResponseBodyItems) GetDimension() *string {
	return s.Dimension
}

func (s *DescribeAvailableMetricsResponseBodyItems) GetGroupKey() *string {
	return s.GroupKey
}

func (s *DescribeAvailableMetricsResponseBodyItems) GetGroupKeyType() *string {
	return s.GroupKeyType
}

func (s *DescribeAvailableMetricsResponseBodyItems) GetMethod() *string {
	return s.Method
}

func (s *DescribeAvailableMetricsResponseBodyItems) GetMetricsKey() *string {
	return s.MetricsKey
}

func (s *DescribeAvailableMetricsResponseBodyItems) GetMetricsKeyAlias() *string {
	return s.MetricsKeyAlias
}

func (s *DescribeAvailableMetricsResponseBodyItems) GetSortRule() *int32 {
	return s.SortRule
}

func (s *DescribeAvailableMetricsResponseBodyItems) GetUnit() *string {
	return s.Unit
}

func (s *DescribeAvailableMetricsResponseBodyItems) SetDescription(v string) *DescribeAvailableMetricsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBodyItems) SetDimension(v string) *DescribeAvailableMetricsResponseBodyItems {
	s.Dimension = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBodyItems) SetGroupKey(v string) *DescribeAvailableMetricsResponseBodyItems {
	s.GroupKey = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBodyItems) SetGroupKeyType(v string) *DescribeAvailableMetricsResponseBodyItems {
	s.GroupKeyType = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBodyItems) SetMethod(v string) *DescribeAvailableMetricsResponseBodyItems {
	s.Method = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBodyItems) SetMetricsKey(v string) *DescribeAvailableMetricsResponseBodyItems {
	s.MetricsKey = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBodyItems) SetMetricsKeyAlias(v string) *DescribeAvailableMetricsResponseBodyItems {
	s.MetricsKeyAlias = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBodyItems) SetSortRule(v int32) *DescribeAvailableMetricsResponseBodyItems {
	s.SortRule = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBodyItems) SetUnit(v string) *DescribeAvailableMetricsResponseBodyItems {
	s.Unit = &v
	return s
}

func (s *DescribeAvailableMetricsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
