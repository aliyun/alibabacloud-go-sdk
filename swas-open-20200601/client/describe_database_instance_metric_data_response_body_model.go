// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDatabaseInstanceMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataFormat(v string) *DescribeDatabaseInstanceMetricDataResponseBody
	GetDataFormat() *string
	SetMetricData(v string) *DescribeDatabaseInstanceMetricDataResponseBody
	GetMetricData() *string
	SetMetricName(v string) *DescribeDatabaseInstanceMetricDataResponseBody
	GetMetricName() *string
	SetRequestId(v string) *DescribeDatabaseInstanceMetricDataResponseBody
	GetRequestId() *string
	SetUnit(v string) *DescribeDatabaseInstanceMetricDataResponseBody
	GetUnit() *string
}

type DescribeDatabaseInstanceMetricDataResponseBody struct {
	// The format of the metric data. Valid values:
	//
	// - `cpuusage&memusage`
	//
	// - `active_session&total_session`
	//
	// - `ins_size&data_size&log_size&tmp_size&other_size`
	//
	// - `io`
	//
	// example:
	//
	// cpuusage&memusage
	DataFormat *string `json:"DataFormat,omitempty" xml:"DataFormat,omitempty"`
	// The metric data details.
	//
	// example:
	//
	// [  {     \\"date"\\: " 2022-09-06T04:04:00Z",\\"value\\":\\"0.77&3.69\\"  } ]
	MetricData *string `json:"MetricData,omitempty" xml:"MetricData,omitempty"`
	// The name of the metric. Valid values:
	//
	// - `MySQL_MemCpuUsage`: The CPU utilization and memory usage of the instance, as a percentage of the total resources of the operating system.
	//
	// - `MySQL_DetailedSpaceUsage`: The total, data, log, temporary, and system disk space usage of the instance.
	//
	// - `MySQL_Sessions`: The number of active and total connections.
	//
	// - `MySQL_IOPS`: The IOPS of the instance.
	//
	// example:
	//
	// MySQL_MemCpuUsage
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The unit of the metric data.
	//
	// - %
	//
	// - int
	//
	// - MB
	//
	// example:
	//
	// %
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeDatabaseInstanceMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDatabaseInstanceMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) GetDataFormat() *string {
	return s.DataFormat
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) GetMetricData() *string {
	return s.MetricData
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) GetUnit() *string {
	return s.Unit
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetDataFormat(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.DataFormat = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetMetricData(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.MetricData = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetMetricName(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.MetricName = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetRequestId(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) SetUnit(v string) *DescribeDatabaseInstanceMetricDataResponseBody {
	s.Unit = &v
	return s
}

func (s *DescribeDatabaseInstanceMetricDataResponseBody) Validate() error {
	return dara.Validate(s)
}
