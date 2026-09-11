// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DescribeMetricListRequest
	GetAccountId() *string
	SetClientToken(v string) *DescribeMetricListRequest
	GetClientToken() *string
	SetDtsJobId(v string) *DescribeMetricListRequest
	GetDtsJobId() *string
	SetEndTime(v int64) *DescribeMetricListRequest
	GetEndTime() *int64
	SetEnv(v string) *DescribeMetricListRequest
	GetEnv() *string
	SetMetricName(v string) *DescribeMetricListRequest
	GetMetricName() *string
	SetMetricType(v string) *DescribeMetricListRequest
	GetMetricType() *string
	SetOwnerID(v string) *DescribeMetricListRequest
	GetOwnerID() *string
	SetParam(v string) *DescribeMetricListRequest
	GetParam() *string
	SetPeriod(v int64) *DescribeMetricListRequest
	GetPeriod() *int64
	SetResourceGroupId(v string) *DescribeMetricListRequest
	GetResourceGroupId() *string
	SetStartTime(v int64) *DescribeMetricListRequest
	GetStartTime() *int64
}

type DescribeMetricListRequest struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 12323344****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Ensures the idempotency of the request. Generate a parameter value from your client to make sure that the value is unique among different requests. **ClientToken*	- supports only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data migration or data synchronization task.
	//
	// example:
	//
	// k2gm967v16f****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The end timestamp, in milliseconds.
	//
	// example:
	//
	// 1642476194000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Default value: **ALIYUN**.
	//
	// example:
	//
	// ALIYUN
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// - **InternetOut**: outbound Internet traffic, in bytes.
	//
	// - **diskusage_utilization**: disk usage.
	//
	// - **IntranetInRate**: inbound internal network traffic, in bytes.
	//
	// - **InternetIn**: inbound Internet traffic, in bytes.
	//
	// - **cpu_total**: CPU utilization.
	//
	// - **memory_usedutilization**: memory utilization.
	//
	// - **IntranetOutRate**: outbound internal network traffic, in bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// Specifies whether to query a cluster or a node. Valid values:
	//
	// - **CLUSTER**: cluster.
	//
	// - **NODE**: node.
	//
	// example:
	//
	// NODE
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	OwnerID    *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The observation value. If **MetricType*	- is set to **NODE**, the value is **nodeid**.
	//
	// This parameter is required.
	//
	// example:
	//
	// nodeid
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The observation interval in seconds. The minimum interval is 15 seconds.
	//
	// example:
	//
	// 15
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The start timestamp, in milliseconds.
	//
	// example:
	//
	// 1642476144000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMetricListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricListRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DescribeMetricListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeMetricListRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeMetricListRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeMetricListRequest) GetEnv() *string {
	return s.Env
}

func (s *DescribeMetricListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMetricListRequest) GetMetricType() *string {
	return s.MetricType
}

func (s *DescribeMetricListRequest) GetOwnerID() *string {
	return s.OwnerID
}

func (s *DescribeMetricListRequest) GetParam() *string {
	return s.Param
}

func (s *DescribeMetricListRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *DescribeMetricListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeMetricListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeMetricListRequest) SetAccountId(v string) *DescribeMetricListRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeMetricListRequest) SetClientToken(v string) *DescribeMetricListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMetricListRequest) SetDtsJobId(v string) *DescribeMetricListRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeMetricListRequest) SetEndTime(v int64) *DescribeMetricListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMetricListRequest) SetEnv(v string) *DescribeMetricListRequest {
	s.Env = &v
	return s
}

func (s *DescribeMetricListRequest) SetMetricName(v string) *DescribeMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMetricListRequest) SetMetricType(v string) *DescribeMetricListRequest {
	s.MetricType = &v
	return s
}

func (s *DescribeMetricListRequest) SetOwnerID(v string) *DescribeMetricListRequest {
	s.OwnerID = &v
	return s
}

func (s *DescribeMetricListRequest) SetParam(v string) *DescribeMetricListRequest {
	s.Param = &v
	return s
}

func (s *DescribeMetricListRequest) SetPeriod(v int64) *DescribeMetricListRequest {
	s.Period = &v
	return s
}

func (s *DescribeMetricListRequest) SetResourceGroupId(v string) *DescribeMetricListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeMetricListRequest) SetStartTime(v int64) *DescribeMetricListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMetricListRequest) Validate() error {
	return dara.Validate(s)
}
