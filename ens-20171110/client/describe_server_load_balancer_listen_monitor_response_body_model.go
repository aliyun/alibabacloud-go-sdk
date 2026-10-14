// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerLoadBalancerListenMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeServerLoadBalancerListenMonitorResponseBody
	GetRequestId() *string
	SetServerLoadBalancerMonitorData(v []*DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) *DescribeServerLoadBalancerListenMonitorResponseBody
	GetServerLoadBalancerMonitorData() []*DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData
}

type DescribeServerLoadBalancerListenMonitorResponseBody struct {
	// Id of the request。
	//
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of objects.
	ServerLoadBalancerMonitorData []*DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData `json:"ServerLoadBalancerMonitorData,omitempty" xml:"ServerLoadBalancerMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeServerLoadBalancerListenMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerLoadBalancerListenMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBody) GetServerLoadBalancerMonitorData() []*DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	return s.ServerLoadBalancerMonitorData
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBody) SetRequestId(v string) *DescribeServerLoadBalancerListenMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBody) SetServerLoadBalancerMonitorData(v []*DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) *DescribeServerLoadBalancerListenMonitorResponseBody {
	s.ServerLoadBalancerMonitorData = v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBody) Validate() error {
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

type DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData struct {
	// The total number of requests.
	//
	// example:
	//
	// 20
	Acc *int32 `json:"Acc,omitempty" xml:"Acc,omitempty"`
	// The log business time. The log collection granularity is 1 minute.
	//
	// example:
	//
	// 2024-05-16 15:00:00
	BizTime *string `json:"BizTime,omitempty" xml:"BizTime,omitempty"`
	// The ID of the edge node.
	//
	// example:
	//
	// cn-fuzhou-7
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The load balancing instance ID.
	//
	// example:
	//
	// lb-5rcvo1n1t3hykfhhjwjgqp****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The load balancing instance name.
	//
	// example:
	//
	// esk-edge-service-lb-8377****
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The load balancing instance specification.
	//
	// example:
	//
	// elb.s2.medium
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	// The request protocol, such as http, https, and tcp.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The number of 2xx status code responses.
	//
	// example:
	//
	// 10
	Reqs2xx *int32 `json:"Reqs2xx,omitempty" xml:"Reqs2xx,omitempty"`
	// The number of 3xx status code responses.
	//
	// example:
	//
	// 0
	Reqs3xx *int32 `json:"Reqs3xx,omitempty" xml:"Reqs3xx,omitempty"`
	// The number of 4xx status code responses.
	//
	// example:
	//
	// 0
	Reqs4xx *int32 `json:"Reqs4xx,omitempty" xml:"Reqs4xx,omitempty"`
	// The number of 5xx status code responses.
	//
	// example:
	//
	// 10
	Reqs5xx *int32 `json:"Reqs5xx,omitempty" xml:"Reqs5xx,omitempty"`
	// The average request latency in milliseconds.
	//
	// example:
	//
	// 1037
	RtAvg *int32 `json:"RtAvg,omitempty" xml:"RtAvg,omitempty"`
	// The VIP of the instance.
	//
	// example:
	//
	// 10.0****
	Vip *string `json:"Vip,omitempty" xml:"Vip,omitempty"`
	// The tunnel ID.
	//
	// example:
	//
	// 52497
	Vni *int32 `json:"Vni,omitempty" xml:"Vni,omitempty"`
	// The VIP port, such as 80, 8080, and 443.
	//
	// example:
	//
	// 80
	Vport *int32 `json:"Vport,omitempty" xml:"Vport,omitempty"`
}

func (s DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetAcc() *int32 {
	return s.Acc
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetBizTime() *string {
	return s.BizTime
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetProto() *string {
	return s.Proto
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetReqs2xx() *int32 {
	return s.Reqs2xx
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetReqs3xx() *int32 {
	return s.Reqs3xx
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetReqs4xx() *int32 {
	return s.Reqs4xx
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetReqs5xx() *int32 {
	return s.Reqs5xx
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetRtAvg() *int32 {
	return s.RtAvg
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetVip() *string {
	return s.Vip
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetVni() *int32 {
	return s.Vni
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) GetVport() *int32 {
	return s.Vport
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetAcc(v int32) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Acc = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetBizTime(v string) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.BizTime = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetEnsRegionId(v string) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetLoadBalancerId(v string) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetLoadBalancerName(v string) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetLoadBalancerSpec(v string) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.LoadBalancerSpec = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetProto(v string) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Proto = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetReqs2xx(v int32) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Reqs2xx = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetReqs3xx(v int32) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Reqs3xx = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetReqs4xx(v int32) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Reqs4xx = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetReqs5xx(v int32) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Reqs5xx = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetRtAvg(v int32) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.RtAvg = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetVip(v string) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Vip = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetVni(v int32) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Vni = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) SetVport(v int32) *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData {
	s.Vport = &v
	return s
}

func (s *DescribeServerLoadBalancerListenMonitorResponseBodyServerLoadBalancerMonitorData) Validate() error {
	return dara.Validate(s)
}
