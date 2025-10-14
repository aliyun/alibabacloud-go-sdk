// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerLoadBalancerMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServerLoadBalancerMonitorResponseBody
	GetRequestId() *string
	SetServerLoadBalancerMonitorData(v []*DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) *DescribeServerLoadBalancerMonitorResponseBody
	GetServerLoadBalancerMonitorData() []*DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData
}

type DescribeServerLoadBalancerMonitorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// AAE90880-4970-4D81-A534-A6C0F3631F74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the monitoring data.
	ServerLoadBalancerMonitorData []*DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData `json:"ServerLoadBalancerMonitorData,omitempty" xml:"ServerLoadBalancerMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeServerLoadBalancerMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerLoadBalancerMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServerLoadBalancerMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServerLoadBalancerMonitorResponseBody) GetServerLoadBalancerMonitorData() []*DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	return s.ServerLoadBalancerMonitorData
}

func (s *DescribeServerLoadBalancerMonitorResponseBody) SetRequestId(v string) *DescribeServerLoadBalancerMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBody) SetServerLoadBalancerMonitorData(v []*DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) *DescribeServerLoadBalancerMonitorResponseBody {
	s.ServerLoadBalancerMonitorData = v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBody) Validate() error {
	if s.ServerLoadBalancerMonitorData != nil {
		for _, item := range s.ServerLoadBalancerMonitorData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData struct {
	// The total number of requests.
	//
	// example:
	//
	// 30
	Acc *int32 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// The business time of the log. Logs are collected every minute.
	//
	// example:
	//
	// 2024-09-15 16:00:00
	BizTime *string `json:"BizTime,omitempty" xml:"BizTime,omitempty"`
	// The ID of the node to which the ELB instance belongs.
	//
	// example:
	//
	// cn-wuxi-10
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The ID of the ELB instance.
	//
	// example:
	//
	// lb-5sc1s9zrui8lpb8u7cl4f****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the ELB instance.
	//
	// example:
	//
	// esk-edge-service-lb-a34****
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The specification of the ELB instance.
	//
	// example:
	//
	// elb.s2.medium
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	// The number of requests with HTTP 2xx status code returned.
	//
	// example:
	//
	// 25
	Reqs2xx *int32 `json:"Reqs2xx,omitempty" xml:"Reqs2xx,omitempty"`
	// The number of requests with HTTP 3xx status code returned.
	//
	// example:
	//
	// 0
	Reqs3xx *int32 `json:"Reqs3xx,omitempty" xml:"Reqs3xx,omitempty"`
	// The number of requests with HTTP 4xx status code returned.
	//
	// example:
	//
	// 5
	Reqs4xx *int32 `json:"Reqs4xx,omitempty" xml:"Reqs4xx,omitempty"`
	// The number of requests with HTTP 5xx status code returned.
	//
	// example:
	//
	// 0
	Reqs5xx *int32 `json:"Reqs5xx,omitempty" xml:"Reqs5xx,omitempty"`
	// The average response time. Unit: milliseconds.
	//
	// example:
	//
	// 1404
	RtAvg *int32 `json:"RtAvg,omitempty" xml:"RtAvg,omitempty"`
	// The virtual IP address (VIP) of the instance.
	//
	// example:
	//
	// 10.0****
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
	// The ID of the tunnel.
	//
	// example:
	//
	// 3018
	Vni *int32 `json:"Vni,omitempty" xml:"Vni,omitempty"`
}

func (s DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetAcc() *int32 {
	return s.Acc
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetBizTime() *string {
	return s.BizTime
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetReqs2xx() *int32 {
	return s.Reqs2xx
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetReqs3xx() *int32 {
	return s.Reqs3xx
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetReqs4xx() *int32 {
	return s.Reqs4xx
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetReqs5xx() *int32 {
	return s.Reqs5xx
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetRtAvg() *int32 {
	return s.RtAvg
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetVip() *string {
	return s.Vip
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) GetVni() *int32 {
	return s.Vni
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetAcc(v int32) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Acc = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetBizTime(v string) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.BizTime = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetEnsRegionId(v string) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetLoadBalancerId(v string) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetLoadBalancerName(v string) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetLoadBalancerSpec(v string) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.LoadBalancerSpec = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetReqs2xx(v int32) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Reqs2xx = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetReqs3xx(v int32) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Reqs3xx = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetReqs4xx(v int32) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Reqs4xx = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetReqs5xx(v int32) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Reqs5xx = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetRtAvg(v int32) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.RtAvg = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetVip(v string) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Vip = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) SetVni(v int32) *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Vni = &v
	return s
}

func (s *DescribeServerLoadBalancerMonitorResponseBodyServerLoadBalancerMonitorData) Validate() error {
	return dara.Validate(s)
}
