// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMetricListResponseBody
	GetCode() *string
	SetDataPoints(v []*DescribeMetricListResponseBodyDataPoints) *DescribeMetricListResponseBody
	GetDataPoints() []*DescribeMetricListResponseBodyDataPoints
	SetDynamicMessage(v string) *DescribeMetricListResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeMetricListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeMetricListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *DescribeMetricListResponseBody
	GetHttpStatusCode() *int32
	SetMetricName(v string) *DescribeMetricListResponseBody
	GetMetricName() *string
	SetMetricType(v string) *DescribeMetricListResponseBody
	GetMetricType() *string
	SetParam(v string) *DescribeMetricListResponseBody
	GetParam() *string
	SetPeriod(v int64) *DescribeMetricListResponseBody
	GetPeriod() *int64
	SetRequestId(v string) *DescribeMetricListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMetricListResponseBody
	GetSuccess() *bool
}

type DescribeMetricListResponseBody struct {
	// The error code returned by the backend service. The number is incremented.
	//
	// example:
	//
	// 403
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The monitoring statistics.
	DataPoints []*DescribeMetricListResponseBodyDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Repeated"`
	// The dynamic part in the error message. This parameter is used to replace the %s variable in the **ErrMessage*	- parameter.
	//
	// example:
	//
	// Type
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code returned for an exception.
	//
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// 	- **InternetOut**: the outbound traffic over the Internet. Unit: byte.
	//
	// 	- **diskusage_utilization**: the disk usage.
	//
	// 	- **IntranetInRate**: the inbound traffic over the internal network. Unit: byte.
	//
	// 	- **InternetIn**: the inbound traffic from the Internet. Unit: byte.
	//
	// 	- **cpu_total**: the CPU utilization.
	//
	// 	- **memory_usedutilization**: the memory usage.
	//
	// 	- **IntranetOutRate**: the outbound traffic over the internal network. Unit: byte.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// Indicates whether the metrics of the cluster or a node are queried. Valid values:
	//
	// 	- **CLUSTER**: The metrics of the cluster are queried.
	//
	// 	- **NODE**: The metrics of a node are queried.
	//
	// example:
	//
	// CLUSTER
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// The monitored object.
	//
	// 	- If the **MetricType*	- parameter is set to **NODE**, the value of this parameter is the ID of the node that is monitored.****
	//
	// 	- If the **MetricType*	- parameter is set to **CLUSTER**, the value of this parameter is the ID of the dedicated cluster. You can obtain the ID by calling the ListDedicatedCluster operation.
	//
	// example:
	//
	// ecs-jhjnjjn
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The monitoring interval. Unit: seconds. Minimum value: 15.
	//
	// example:
	//
	// 15
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMetricListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMetricListResponseBody) GetDataPoints() []*DescribeMetricListResponseBodyDataPoints {
	return s.DataPoints
}

func (s *DescribeMetricListResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeMetricListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeMetricListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeMetricListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeMetricListResponseBody) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricListResponseBody) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeMetricListResponseBody) GetParam() *string {
	return s.Param
}

func (s *DescribeMetricListResponseBody) GetPeriod() *int64 {
	return s.Period
}

func (s *DescribeMetricListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMetricListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMetricListResponseBody) SetCode(v string) *DescribeMetricListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetDataPoints(v []*DescribeMetricListResponseBodyDataPoints) *DescribeMetricListResponseBody {
	s.DataPoints = v
	return s
}

func (s *DescribeMetricListResponseBody) SetDynamicMessage(v string) *DescribeMetricListResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetErrCode(v string) *DescribeMetricListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetErrMessage(v string) *DescribeMetricListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetHttpStatusCode(v int32) *DescribeMetricListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetMetricName(v string) *DescribeMetricListResponseBody {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetMetricType(v string) *DescribeMetricListResponseBody {
	s.MetricType = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetParam(v string) *DescribeMetricListResponseBody {
	s.Param = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetPeriod(v int64) *DescribeMetricListResponseBody {
	s.Period = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetRequestId(v string) *DescribeMetricListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMetricListResponseBody) SetSuccess(v bool) *DescribeMetricListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMetricListResponseBody) Validate() error {
	if s.DataPoints != nil {
		for _, item := range s.DataPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMetricListResponseBodyDataPoints struct {
	// The statistical value.
	//
	// example:
	//
	// 15.25
	Statistics *float32 `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// The timestamp of the record. Unit: milliseconds.
	//
	// example:
	//
	// 1650872310000
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s DescribeMetricListResponseBodyDataPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListResponseBodyDataPoints) GoString() string {
	return s.String()
}

func (s *DescribeMetricListResponseBodyDataPoints) GetStatistics() *float32 {
	return s.Statistics
}

func (s *DescribeMetricListResponseBodyDataPoints) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *DescribeMetricListResponseBodyDataPoints) SetStatistics(v float32) *DescribeMetricListResponseBodyDataPoints {
	s.Statistics = &v
	return s
}

func (s *DescribeMetricListResponseBodyDataPoints) SetTimestamp(v int64) *DescribeMetricListResponseBodyDataPoints {
	s.Timestamp = &v
	return s
}

func (s *DescribeMetricListResponseBodyDataPoints) Validate() error {
	return dara.Validate(s)
}
