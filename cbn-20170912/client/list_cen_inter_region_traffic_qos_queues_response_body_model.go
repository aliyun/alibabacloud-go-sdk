// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCenInterRegionTrafficQosQueuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListCenInterRegionTrafficQosQueuesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCenInterRegionTrafficQosQueuesResponseBody
	GetRequestId() *string
	SetTrafficQosQueues(v []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) *ListCenInterRegionTrafficQosQueuesResponseBody
	GetTrafficQosQueues() []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues
}

type ListCenInterRegionTrafficQosQueuesResponseBody struct {
	// A pagination token. It is used in the next request to retrieve a new page of results. Valid values:
	//
	// - If **NextToken*	- is empty, no more results are returned.
	//
	// - If NextToken is not empty, the value of **NextToken*	- is used for the next query.
	//
	// example:
	//
	// 0151fa6aa1ed****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1D1E15D2-416D-54F3-BDD9-BC27DE4C6352
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the queues of the QoS policy.
	TrafficQosQueues []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues `json:"TrafficQosQueues,omitempty" xml:"TrafficQosQueues,omitempty" type:"Repeated"`
}

func (s ListCenInterRegionTrafficQosQueuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosQueuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) GetTrafficQosQueues() []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	return s.TrafficQosQueues
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) SetNextToken(v string) *ListCenInterRegionTrafficQosQueuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) SetRequestId(v string) *ListCenInterRegionTrafficQosQueuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) SetTrafficQosQueues(v []*ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) *ListCenInterRegionTrafficQosQueuesResponseBody {
	s.TrafficQosQueues = v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBody) Validate() error {
	if s.TrafficQosQueues != nil {
		for _, item := range s.TrafficQosQueues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues struct {
	// The bandwidth allocated to the queue. This parameter is returned only if you allocate bandwidth to the queue by absolute value.
	//
	// For example, a value of **1*	- indicates that traffic that matches the queue can use up to 1 Mbit/s of the inter-region connection bandwidth.
	//
	// example:
	//
	// 1
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The Differentiated Services Code Point (DSCP) values that are matched by the queue.
	Dscps []*int32 `json:"Dscps,omitempty" xml:"Dscps,omitempty" type:"Repeated"`
	// The actual bandwidth of the queue.
	//
	// example:
	//
	// 1.35
	EffectiveBandwidth *string `json:"EffectiveBandwidth,omitempty" xml:"EffectiveBandwidth,omitempty"`
	// The percentage of the inter-region connection bandwidth that can be used by the queue. This parameter is returned only if you allocate bandwidth to the queue by percentage.
	//
	// For example, a value of **1*	- indicates that traffic that matches the queue can use up to 1% of the inter-region connection bandwidth.
	//
	// example:
	//
	// 1
	RemainBandwidthPercent *int32 `json:"RemainBandwidthPercent,omitempty" xml:"RemainBandwidthPercent,omitempty"`
	// The status of the queue.
	//
	// - **Creating**: The queue is being created.
	//
	// - **Active**: The queue is running.
	//
	// - **Deleting**: The queue is being deleted.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-fv2qq9yqrsjowp****
	TrafficQosPolicyId *string `json:"TrafficQosPolicyId,omitempty" xml:"TrafficQosPolicyId,omitempty"`
	// The description of the queue in the QoS policy.
	//
	// example:
	//
	// qosQueueDescription
	TrafficQosQueueDescription *string `json:"TrafficQosQueueDescription,omitempty" xml:"TrafficQosQueueDescription,omitempty"`
	// The ID of the queue in the QoS policy.
	//
	// example:
	//
	// qos-queue-siakjb2nn9gz5z****
	TrafficQosQueueId *string `json:"TrafficQosQueueId,omitempty" xml:"TrafficQosQueueId,omitempty"`
	// The name of the queue in the QoS policy.
	//
	// example:
	//
	// qosQueueName
	TrafficQosQueueName *string `json:"TrafficQosQueueName,omitempty" xml:"TrafficQosQueueName,omitempty"`
	// The ID of the inter-region connection.
	//
	// example:
	//
	// tr-attach-nzrcv25d7ezt23****
	TransitRouterAttachmentId *string `json:"TransitRouterAttachmentId,omitempty" xml:"TransitRouterAttachmentId,omitempty"`
	// The ID of the transit router.
	//
	// example:
	//
	// tr-p0wwagjv6fvxt4b7y****
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) String() string {
	return dara.Prettify(s)
}

func (s ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GoString() string {
	return s.String()
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetDscps() []*int32 {
	return s.Dscps
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetEffectiveBandwidth() *string {
	return s.EffectiveBandwidth
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetRemainBandwidthPercent() *int32 {
	return s.RemainBandwidthPercent
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetStatus() *string {
	return s.Status
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTrafficQosPolicyId() *string {
	return s.TrafficQosPolicyId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTrafficQosQueueDescription() *string {
	return s.TrafficQosQueueDescription
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTrafficQosQueueId() *string {
	return s.TrafficQosQueueId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTrafficQosQueueName() *string {
	return s.TrafficQosQueueName
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTransitRouterAttachmentId() *string {
	return s.TransitRouterAttachmentId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetBandwidth(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.Bandwidth = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetDscps(v []*int32) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.Dscps = v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetEffectiveBandwidth(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.EffectiveBandwidth = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetRemainBandwidthPercent(v int32) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.RemainBandwidthPercent = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetStatus(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.Status = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTrafficQosPolicyId(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TrafficQosPolicyId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTrafficQosQueueDescription(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TrafficQosQueueDescription = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTrafficQosQueueId(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TrafficQosQueueId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTrafficQosQueueName(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TrafficQosQueueName = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTransitRouterAttachmentId(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TransitRouterAttachmentId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) SetTransitRouterId(v string) *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues {
	s.TransitRouterId = &v
	return s
}

func (s *ListCenInterRegionTrafficQosQueuesResponseBodyTrafficQosQueues) Validate() error {
	return dara.Validate(s)
}
