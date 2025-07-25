// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTotalCount(v int32) *DescribeCloudGtmSummaryResponseBody
	GetInstanceTotalCount() *int32
	SetMonitorTaskTotalCount(v int32) *DescribeCloudGtmSummaryResponseBody
	GetMonitorTaskTotalCount() *int32
	SetMonitorTaskTotalQuota(v int32) *DescribeCloudGtmSummaryResponseBody
	GetMonitorTaskTotalQuota() *int32
	SetRequestId(v string) *DescribeCloudGtmSummaryResponseBody
	GetRequestId() *string
}

type DescribeCloudGtmSummaryResponseBody struct {
	// The total number of instances within the current account.
	//
	// example:
	//
	// 10
	InstanceTotalCount *int32 `json:"InstanceTotalCount,omitempty" xml:"InstanceTotalCount,omitempty"`
	// The total number of configured health check tasks.
	//
	// example:
	//
	// 20
	MonitorTaskTotalCount *int32 `json:"MonitorTaskTotalCount,omitempty" xml:"MonitorTaskTotalCount,omitempty"`
	// The quota on the number of health check tasks.
	//
	// example:
	//
	// 101
	MonitorTaskTotalQuota *int32 `json:"MonitorTaskTotalQuota,omitempty" xml:"MonitorTaskTotalQuota,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 199C3699-9A7B-41A1-BB5A-F1E862D3CB38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudGtmSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmSummaryResponseBody) GetInstanceTotalCount() *int32 {
	return s.InstanceTotalCount
}

func (s *DescribeCloudGtmSummaryResponseBody) GetMonitorTaskTotalCount() *int32 {
	return s.MonitorTaskTotalCount
}

func (s *DescribeCloudGtmSummaryResponseBody) GetMonitorTaskTotalQuota() *int32 {
	return s.MonitorTaskTotalQuota
}

func (s *DescribeCloudGtmSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmSummaryResponseBody) SetInstanceTotalCount(v int32) *DescribeCloudGtmSummaryResponseBody {
	s.InstanceTotalCount = &v
	return s
}

func (s *DescribeCloudGtmSummaryResponseBody) SetMonitorTaskTotalCount(v int32) *DescribeCloudGtmSummaryResponseBody {
	s.MonitorTaskTotalCount = &v
	return s
}

func (s *DescribeCloudGtmSummaryResponseBody) SetMonitorTaskTotalQuota(v int32) *DescribeCloudGtmSummaryResponseBody {
	s.MonitorTaskTotalQuota = &v
	return s
}

func (s *DescribeCloudGtmSummaryResponseBody) SetRequestId(v string) *DescribeCloudGtmSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}
