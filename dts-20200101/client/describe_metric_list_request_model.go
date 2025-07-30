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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data migration or synchronization task.
	//
	// example:
	//
	// k2gm967v16f****
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	// The timestamp that indicates the end of the time range to query. Unit: milliseconds.
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
	// This parameter is required.
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
	// NODE
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	OwnerID    *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The monitored object. If the **MetricType*	- parameter is set to **NODE**, set this parameter to the ID of the node that is monitored.
	//
	// This parameter is required.
	//
	// example:
	//
	// nodeid
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The monitoring interval. Unit: seconds. The minimum value is 15.
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
	// The timestamp that indicates the beginning of the time range to query. Unit: milliseconds.
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
